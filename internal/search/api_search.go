package search

import (
	"cmp"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/moov-io/base/log"
	"github.com/moov-io/base/strx"
	"github.com/moov-io/base/telemetry"
	"github.com/moov-io/watchman/internal/api"
	"github.com/moov-io/watchman/internal/norm"
	"github.com/moov-io/watchman/internal/postalpool"
	"github.com/moov-io/watchman/internal/prepare"
	"github.com/moov-io/watchman/pkg/address"
	"github.com/moov-io/watchman/pkg/search"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/otel/attribute"
)

type Controller interface {
	AppendRoutes(router *mux.Router) *mux.Router
}

func NewController(logger log.Logger, service Service, addressParsingPool *postalpool.Service) Controller {
	return &controller{
		logger:             logger,
		service:            service,
		addressParsingPool: addressParsingPool,
	}
}

type controller struct {
	logger             log.Logger
	service            Service
	addressParsingPool *postalpool.Service
}

func (c *controller) AppendRoutes(router *mux.Router) *mux.Router {
	router.
		Name("Search.v2").
		Methods("GET").
		Path("/v2/search").
		HandlerFunc(c.search)

	router.
		Name("ListInfo.v2").
		Methods("GET").
		Path("/v2/listinfo").
		HandlerFunc(c.listinfo)

	return router
}

func (c *controller) listinfo(w http.ResponseWriter, r *http.Request) {
	stats := c.service.LatestStats()

	api.JsonResponse(w, stats)
}

func (c *controller) search(w http.ResponseWriter, r *http.Request) {
	ctx, span := telemetry.StartSpan(r.Context(), "api-search")
	defer span.End()

	queryParams := &api.QueryParams{Values: r.URL.Query()}

	debug := strx.Yes(queryParams.Get("debug"))

	req, err := readSearchRequest(ctx, c.addressParsingPool, queryParams)
	if err != nil {
		err = c.logger.Error().LogErrorf("problem reading v2 search request: %w", err).Err()
		api.ErrorResponse(w, err)
		return
	}

	opts := SearchOpts{
		Limit:          extractSearchLimit(queryParams),
		MinMatch:       extractSearchMinMatch(queryParams),
		RequestID:      queryParams.Get("requestID"),
		Debug:          debug,
		DebugSourceIDs: strings.Split(queryParams.Get("debugSourceIDs"), ","),
	}

	span.SetAttributes(
		attribute.String("request_id", opts.RequestID),
		attribute.String("entity.type", string(req.Type)),
	)

	// Check we don't have extra query params
	if extra := queryParams.UnusedQueryParams(); len(extra) > 0 {
		err = c.logger.Error().LogErrorf("extra/unused query parameters in request: %v", strings.Join(extra, ",")).Err()
		api.ErrorResponse(w, err)
		return
	}

	// Perform the search
	entities, err := c.service.Search(ctx, req, opts)
	if err != nil {
		err = c.logger.Error().LogErrorf("problem with v2 search: %v", err).Err()
		api.ErrorResponse(w, err)
		return
	}

	err = api.JsonResponse(w, search.SearchResponse{
		Query:    req,
		Entities: entities,
	})
	if err != nil {
		span.RecordError(err)
	}
}

var (
	softResultsLimit, hardResultsLimit = 10, 100
)

func extractSearchLimit(q *api.QueryParams) int {
	limit := softResultsLimit
	if v := q.Get("limit"); v != "" {
		n, _ := strconv.Atoi(v)
		if n > 0 {
			limit = n
		}
	}
	if limit > hardResultsLimit {
		limit = hardResultsLimit
	}
	if limit < 0 {
		limit = softResultsLimit
	}
	return limit
}

func extractSearchMinMatch(q *api.QueryParams) float64 {
	if v := q.Get("minMatch"); v != "" {
		n, _ := strconv.ParseFloat(v, 64)
		return n
	}
	return 0.00
}

func readSearchRequest(ctx context.Context, addressParsingPool *postalpool.Service, q *api.QueryParams) (search.Entity[search.Value], error) {
	var err error
	var req search.Entity[search.Value]

	req.Name = strings.TrimSpace(q.Get("name"))
	req.Type = search.EntityType(strings.TrimSpace(strings.ToLower(q.Get("type"))))
	req.Source = cmp.Or(search.SourceList(q.Get("source")), search.SourceAPIRequest)

	switch req.Type {
	case search.EntityPerson:
		req.Person = &search.Person{
			Name:      req.Name,
			AltNames:  q.GetAll("altNames"),
			Gender:    search.Gender(prepare.NormalizeGender(q.Get("gender"))),
			BirthDate: readDate(q.Get("birthDate")),
			DeathDate: readDate(q.Get("deathDate")),
			Titles:    q.GetAll("titles"),
			// GovernmentIDs []GovernmentID `json:"governmentIDs"` // TODO(adam):
		}

	case search.EntityBusiness:
		req.Business = &search.Business{
			Name:      req.Name,
			AltNames:  q.GetAll("altNames"),
			Created:   readDate(q.Get("created")),
			Dissolved: readDate(q.Get("dissolved")),
			// Identifier []Identifier `json:"identifier"`
		}

	case search.EntityOrganization:
		req.Organization = &search.Organization{
			Name:      req.Name,
			AltNames:  q.GetAll("altNames"),
			Created:   readDate(q.Get("created")),
			Dissolved: readDate(q.Get("dissolved")),
			// Identifier []Identifier `json:"identifier"`
		}

	case search.EntityAircraft:
		req.Aircraft = &search.Aircraft{
			Name:         req.Name,
			AltNames:     q.GetAll("altNames"),
			Type:         search.AircraftType(q.Get("aircraftType")),
			Flag:         q.Get("flag"),
			Built:        readDate("built"),
			ICAOCode:     q.Get("icaoCode"),
			Model:        q.Get("model"),
			SerialNumber: q.Get("serialNumber"),
		}

	case search.EntityVessel:
		req.Vessel = &search.Vessel{
			Name:      req.Name,
			AltNames:  q.GetAll("altNames"),
			IMONumber: q.Get("imoNumber"),
			Type:      search.VesselType(q.Get("vesselType")),
			Flag:      q.Get("flag"),
			Built:     readDate("built"),
			Model:     q.Get("model"),
			// Tonnage:  // TODO(adam):
			MMSI:     q.Get("mmsi"),
			CallSign: q.Get("callSign"),
			Owner:    q.Get("owner"),
		}
		if v := strings.TrimSpace(q.Get("tonnage")); v != "" {
			req.Vessel.Tonnage, err = readInt(v)
			if err != nil {
				return req, fmt.Errorf("reading vessel tonnage: %w", err)
			}
		}
		if v := strings.TrimSpace(q.Get("grossRegisteredTonnage")); v != "" {
			req.Vessel.GrossRegisteredTonnage, err = readInt(v)
			if err != nil {
				return req, fmt.Errorf("reading vessel GrossRegisteredTonnage: %w", err)
			}
		}
	}

	// contact info // TODO(adam): normalize
	req.Contact.EmailAddresses = readStrings(q.GetAll("email"), q.GetAll("emailAddress"), q.GetAll("emailAddresses"))
	req.Contact.PhoneNumbers = readStrings(q.GetAll("phone"), q.GetAll("phoneNumber"), q.GetAll("phoneNumbers"))
	req.Contact.FaxNumbers = readStrings(q.GetAll("fax"), q.GetAll("faxNumber"), q.GetAll("faxNumbers"))
	req.Contact.Websites = readStrings(q.GetAll("website"), q.GetAll("websites"))

	addresses := readStrings(q.GetAll("address"), q.GetAll("addresses"))
	req.Addresses = readAddresses(ctx, addressParsingPool, addresses)

	cryptoAddresses := readStrings(q.GetAll("cryptoAddress"), q.GetAll("cryptoAddresses"))
	req.CryptoAddresses = readCryptoCurrencyAddresses(cryptoAddresses)

	// TODO(adam):
	// Affiliations   []Affiliation    `json:"affiliations"`
	// SanctionsInfo  *SanctionsInfo   `json:"sanctionsInfo"`
	// HistoricalInfo []HistoricalInfo `json:"historicalInfo"`

	return req.Normalize(), nil
}

var (
	allowedDateFormats = []string{"2006-01-02", "2006-01", "2006"}
)

func readDate(input string) *time.Time {
	if input == "" {
		return nil
	}

	for _, format := range allowedDateFormats {
		tt, err := time.Parse(format, input)
		if err == nil {
			return &tt
		}
	}
	return nil
}

func readInt(input string) (int, error) {
	n, err := strconv.ParseInt(input, 10, 20) // -524,288 to 523,767
	if err != nil {
		return 0, err
	}
	return int(n), nil
}

func readStrings(inputs ...[]string) []string {
	var out []string
	for _, items := range inputs {
		for _, item := range items {
			out = append(out, strings.TrimSpace(item))
		}
	}
	return out
}

func readAddresses(ctx context.Context, addressParsingPool *postalpool.Service, inputs []string) []search.Address {
	out := make([]search.Address, len(inputs))

	for idx, input := range inputs {
		// Prefer the pool if it's defined
		if addressParsingPool != nil {
			addr, err := addressParsingPool.ParseAddress(ctx, input)
			if err == nil {
				out[idx] = addr
			}
		} else {
			// Fallback to standard parsing
			out[idx] = address.ParseAddress(ctx, input)
		}

		// Normalize the country
		out[idx].Country = norm.Country(out[idx].Country)
	}

	return out
}

func readCryptoCurrencyAddresses(inputs []string) []search.CryptoAddress {
	var out []search.CryptoAddress
	for _, input := range inputs {
		// Query param looks like: cryptoAddress=XBT:x123456
		parts := strings.Split(input, ":")
		if len(parts) == 2 {
			out = append(out, search.CryptoAddress{
				Currency: strings.ToUpper(parts[0]),
				Address:  parts[1],
			})
		}
	}
	return out
}
