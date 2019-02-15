// Copyright 2018 The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	moovhttp "github.com/moov-io/base/http"
	"github.com/moov-io/ofac"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
)

var (
	errNoCompanyId = errors.New("no Company ID found")
)

type Company struct {
	ID string `json:"id"`
	// Federal Data
	SDN       *ofac.SDN                 `json:"sdn"`
	Addresses []*ofac.Address           `json:"addresses"`
	Alts      []*ofac.AlternateIdentity `json:"alts"`
	Comments  []*ofac.SDNComments       `json:"comments"`
	// Metadata
	Status *CompanyStatus `json:"status"`
}

type CompanyBlockStatus string

const (
	// CompanyUnsafe companies have been manually marked to block all transactions with
	CompanyUnsafe CompanyBlockStatus = "unsafe"
	// Exception companies have been manually marked to allow transactions with
	CompanyException CompanyBlockStatus = "exception"
)

// CompanyStatus represents a userId's manual override of an SDN
type CompanyStatus struct {
	UserId    string             `json:"userId"`
	Note      string             `json:"note"`
	Status    CompanyBlockStatus `json:"block"`
	CreatedAt time.Time          `json:"createdAt"`
}

type companyWatchResponse struct {
	WatchID string `json:"watchId"`
}

func addCompanyRoutes(logger log.Logger, r *mux.Router, searcher *searcher, companyRepo companyRepository, watchRepo *sqliteWatchRepository) {
	r.Methods("GET").Path("/companies/{companyId}").HandlerFunc(getCompany(logger, searcher, companyRepo))
	r.Methods("PUT").Path("/companies/{companyId}").HandlerFunc(updateCompanyStatus(logger, searcher, companyRepo))

	r.Methods("POST").Path("/companies/{companyId}/watch").HandlerFunc(addCompanyWatch(logger, searcher, watchRepo))
	r.Methods("DELETE").Path("/companies/{companyId}/watch/{watchId}").HandlerFunc(removeCompanyWatch(logger, searcher, watchRepo))

	r.Methods("POST").Path("/companies/watch").HandlerFunc(addCompanyNameWatch(logger, searcher, watchRepo))
	r.Methods("DELETE").Path("/companies/watch/{watchId}").HandlerFunc(removeCompanyNameWatch(logger, searcher, watchRepo))
}

func getCompanyId(w http.ResponseWriter, r *http.Request) string {
	v, ok := mux.Vars(r)["companyId"]
	if !ok || v == "" {
		moovhttp.Problem(w, errNoCompanyId)
		return ""
	}
	return v
}

func getCompanyById(id string, searcher *searcher, repo companyRepository) (*Company, error) {
	sdn := searcher.FindSDN(id)
	if sdn == nil {
		return nil, fmt.Errorf("Company %s not found", id)
	}
	if strings.EqualFold(sdn.SDNType, "individual") {
		// SDN is an individual, so they aren't a company/trust/organization
		return nil, nil
	}
	if repo == nil {
		return nil, errors.New("nil companyRepository provided - logic bug")
	}
	status, err := repo.getCompanyStatus(sdn.EntityID)
	if err != nil {
		return nil, fmt.Errorf("problem reading Company=%s block status: %v", id, err)
	}
	return &Company{
		ID:        id,
		SDN:       sdn,
		Addresses: searcher.FindAddresses(100, id),
		Alts:      searcher.FindAlts(100, id),
		Status:    status,
	}, nil
}

// companyRepository holds the current status (i.e. unsafe or exception) for a given company and
// is expected to save metadata about each time the status is changed.
type companyRepository interface {
	getCompanyStatus(companyId string) (*CompanyStatus, error)
	upsertCompanyStatus(companyId string, status *CompanyStatus) error
}

type sqliteCompanyRepository struct {
	db *sql.DB
}

func (r *sqliteCompanyRepository) close() error {
	return r.db.Close()
}

func (r *sqliteCompanyRepository) getCompanyStatus(companyId string) (*CompanyStatus, error) {
	if companyId == "" {
		return nil, errors.New("getCompanyStatus: no Company.ID")
	}
	query := `select user_id, note, status, created_at from company_status where company_id = ? and deleted_at is null order by created_at desc limit 1;`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(companyId)

	var status CompanyStatus
	err = row.Scan(&status.UserId, &status.Note, &status.Status, &status.CreatedAt)
	if err != nil && !strings.Contains(err.Error(), "no rows in result set") {
		return nil, fmt.Errorf("getCompanyStatus: %v", err)
	}
	if status.UserId == "" {
		return nil, nil // not found
	}
	return &status, nil
}

func (r *sqliteCompanyRepository) upsertCompanyStatus(companyId string, status *CompanyStatus) error {
	query := `insert or replace into company_status (company_id, user_id, note, status, created_at) values (?, ?, ?, ?, ?);`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(companyId, status.UserId, status.Note, status.Status, status.CreatedAt)
	return err
}

func getCompany(logger log.Logger, searcher *searcher, companyRepo companyRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)
		id := getCompanyId(w, r)
		if id == "" {
			return
		}
		company, err := getCompanyById(id, searcher, companyRepo)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(company)
	}
}

type companyStatusRequest struct {
	Notes string `json:"notes,omitempty"`

	// Status represents a manual exception or unsafe designation
	Status string `json:"status"`
}

func updateCompanyStatus(logger log.Logger, searcher *searcher, companyRepo companyRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		companyId, userId := getCompanyId(w, r), moovhttp.GetUserId(r)
		if companyId == "" {
			return
		}
		if userId == "" {
			moovhttp.Problem(w, errNoUserId)
			return
		}

		var req companyStatusRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}

		status := CompanyBlockStatus(strings.ToLower(strings.TrimSpace(req.Status)))
		switch status {
		case CompanyUnsafe, CompanyException:
			companyStatus := &CompanyStatus{
				UserId:    userId,
				Note:      req.Notes,
				Status:    status,
				CreatedAt: time.Now(),
			}
			if err := companyRepo.upsertCompanyStatus(companyId, companyStatus); err != nil {
				moovhttp.Problem(w, err)
				return
			}
			w.WriteHeader(http.StatusOK)
			return
		default:
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "unknown status"}`))
		}
	}
}

func addCompanyWatch(logger log.Logger, searcher *searcher, repo watchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		var req watchRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		if req.AuthToken == "" {
			moovhttp.Problem(w, errNoAuthToken)
			return
		}
		webhook, err := validateWebhook(req.Webhook)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}
		req.Webhook = webhook

		companyId := getCompanyId(w, r)
		if companyId == "" {
			return
		}
		watchId, err := repo.addCompanyWatch(companyId, req)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(companyWatchResponse{watchId})
	}
}

func removeCompanyWatch(logger log.Logger, searcher *searcher, repo watchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		companyId, watchId := getCompanyId(w, r), getWatchId(w, r)
		if companyId == "" || watchId == "" {
			return
		}
		if err := repo.removeCompanyWatch(companyId, watchId); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func addCompanyNameWatch(logger log.Logger, searcher *searcher, repo watchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		name := r.URL.Query().Get("name")
		if name == "" {
			moovhttp.Problem(w, errNoNameParam)
			return
		}

		var req watchRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		if req.AuthToken == "" {
			moovhttp.Problem(w, errNoAuthToken)
			return
		}
		webhook, err := validateWebhook(req.Webhook)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}
		watchId, err := repo.addCompanyNameWatch(name, webhook, req.AuthToken)
		if err != nil {
			moovhttp.Problem(w, err)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(companyWatchResponse{watchId})
	}
}

func removeCompanyNameWatch(logger log.Logger, searcher *searcher, repo watchRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w = wrapResponseWriter(logger, w, r)

		watchId := getWatchId(w, r)
		if watchId == "" {
			return
		}
		if err := repo.removeCompanyNameWatch(watchId); err != nil {
			moovhttp.Problem(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
