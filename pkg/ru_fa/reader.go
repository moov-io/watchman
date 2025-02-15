package ru_fa

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"time"

	"github.com/moov-io/watchman/pkg/download"

	"github.com/PuerkitoBio/goquery"
	"github.com/mozillazg/go-unidecode"
)

// ---------- Constants and Helper Functions ----------

const (
	urlToDownload = "https://ru.ruwiki.ru/wiki/%D0%A1%D0%BF%D0%B8%D1%81%D0%BE%D0%BA_%D0%B8%D0%BD%D0%BE%D1%81%D1%82%D1%80%D0%B0%D0%BD%D0%BD%D1%8B%D1%85_%D0%B0%D0%B3%D0%B5%D0%BD%D1%82%D0%BE%D0%B2_(%D0%A0%D0%BE%D1%81%D1%81%D0%B8%D1%8F)"
	htmlFilename  = "page.html"
	dateLayout    = "02.01.2006" // Adjust if necessary.
)

// normalizeName replaces « and » with a standard double quote.
func normalizeName(s string) string {
	s = strings.ReplaceAll(s, "«", "\"")
	s = strings.ReplaceAll(s, "»", "\"")
	return s
}

// transliterate converts a Russian string into Latin using go-unidecode.
func transliterate(s string) string {
	return unidecode.Unidecode(s)
}

// matchHeaders returns true if the two slices match element-by-element (after trimming).
func matchHeaders(actual, expected []string) bool {
	if len(actual) != len(expected) {
		return false
	}
	for i := range actual {
		if strings.TrimSpace(actual[i]) != expected[i] {
			return false
		}
	}
	return true
}

// isInclusionDateFuture returns true if dateStr (if non-empty) parsed in loc is after now.
func isInclusionDateFuture(dateStr string, now time.Time, loc *time.Location) bool {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" {
		return false
	}
	t, err := time.ParseInLocation(dateLayout, dateStr, loc)
	if err != nil {
		return false
	}
	return t.After(now)
}

// isExclusionDatePast returns true if dateStr (if non-empty) parsed in loc is before now.
func isExclusionDatePast(dateStr string, now time.Time, loc *time.Location) bool {
	dateStr = strings.TrimSpace(dateStr)
	if dateStr == "" {
		return false
	}
	t, err := time.ParseInLocation(dateLayout, dateStr, loc)
	if err != nil {
		return false
	}
	return t.Before(now)
}

// ---------- Expected Table Headers ----------

var headerNonCommercial = []string{"№", "Наименование", "Адрес", "ИНН", "Реестр.№", "Дата включения в реестр", "Дата исключения из реестра"}
var headerMassMedia = []string{"№", "Название", "Дата включения в реестр[1]", "Дата исключения из реестра"}
var headerMediaIndividual = []string{"№", "Имя", "Род деятельности", "Дата включения в реестр[1]", "Дата исключения из реестра"}
var headerForeignAgentIndividual = []string{"№", "", "", "Сведения об иностранных источниках", "Информация об осуществлении политической деятельности и (или) целенаправленном сборе сведений", "Дата включения в реестр[2]"}
var headerUnregisteredAssociation = []string{"№", "Наименование", "Дата включения в реестр", "Цели деятельности объединения", "Сведения об источниках формирования денежных средств и (или) иного имущества объединения, в том числе сведения об иностранных источниках поступления (планируемого поступления) денежных средств и иного имущества"}

// ---------- Struct Definitions ----------

type NonCommercialOrganization struct {
	Number        string
	Name          string
	NameTranslit  string
	Address       string
	INN           string
	RegistryNo    string
	InclusionDate string
	ExclusionDate string
}

type MassMedia struct {
	Number        string
	Name          string
	NameTranslit  string
	InclusionDate string
	ExclusionDate string
}

type MediaIndividual struct {
	Number        string
	Name          string
	NameTranslit  string
	Activity      string
	InclusionDate string
	ExclusionDate string
}

// In this table, the second column is Name and the third is Activity.
type ForeignAgentIndividual struct {
	Number                string
	Name                  string
	NameTranslit          string
	Activity              string
	ForeignSources        string
	PoliticalActivityInfo string
	InclusionDate         string
}

type UnregisteredAssociation struct {
	Number           string
	Name             string
	NameTranslit     string
	InclusionDate    string
	AssociationGoals string
	SourceInfo       string
}

// Tables holds all table data.
type Tables struct {
	NonCommercialOrgs        []NonCommercialOrganization
	MassMedias               []MassMedia
	MediaIndividuals         []MediaIndividual
	ForeignAgentIndividuals  []ForeignAgentIndividual
	UnregisteredAssociations []UnregisteredAssociation
}

func Read(files download.Files) (Tables, error) {
	for filename, file := range files {
		switch strings.ToLower(filepath.Base(filename)) {
		case htmlFilename:
			tables, err := htmlFile(file)
			if err != nil {
				return tables, fmt.Errorf("%s: %v", htmlFilename, err)
			}
			return tables, nil
		default:
			var tables Tables
			file.Close()
			return tables, fmt.Errorf("error: file %s does not have a handler for processing", filename)
		}
	}

	var tables Tables
	return tables, fmt.Errorf("error: file %s not found", htmlFilename)
}

func htmlFile(f io.ReadCloser) (Tables, error) {
	defer f.Close()
	var tables Tables

	// Load Moscow timezone.
	moscowLoc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return tables, fmt.Errorf("loading timezone: %v", err)
	}
	now := time.Now().In(moscowLoc)

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		return tables, fmt.Errorf("parsing HTML: %v", err)
	}

	// Iterate over every table.
	doc.Find("table").Each(func(i int, tableSel *goquery.Selection) {
		var rows [][]string
		tableSel.Find("tr").Each(func(j int, rowSel *goquery.Selection) {
			var row []string
			rowSel.Find("th, td").Each(func(k int, cellSel *goquery.Selection) {
				row = append(row, strings.TrimSpace(cellSel.Text()))
			})
			if len(row) > 0 {
				rows = append(rows, row)
			}
		})
		if len(rows) < 2 {
			return
		}
		header := rows[0]
		switch {
		case matchHeaders(header, headerNonCommercial):
			for _, row := range rows[1:] {
				if len(row) < 7 {
					continue
				}
				if isInclusionDateFuture(row[5], now, moscowLoc) || isExclusionDatePast(row[6], now, moscowLoc) {
					continue
				}
				entry := NonCommercialOrganization{
					Number:        row[0],
					Name:          row[1],
					NameTranslit:  transliterate(normalizeName(row[1])),
					Address:       row[2],
					INN:           row[3],
					RegistryNo:    row[4],
					InclusionDate: row[5],
					ExclusionDate: row[6],
				}
				tables.NonCommercialOrgs = append(tables.NonCommercialOrgs, entry)
			}
		case matchHeaders(header, headerMassMedia):
			for _, row := range rows[1:] {
				if len(row) < 4 {
					continue
				}
				if isInclusionDateFuture(row[2], now, moscowLoc) || isExclusionDatePast(row[3], now, moscowLoc) {
					continue
				}
				entry := MassMedia{
					Number:        row[0],
					Name:          row[1],
					NameTranslit:  transliterate(normalizeName(row[1])),
					InclusionDate: row[2],
					ExclusionDate: row[3],
				}
				tables.MassMedias = append(tables.MassMedias, entry)
			}
		case matchHeaders(header, headerMediaIndividual):
			for _, row := range rows[1:] {
				if len(row) < 5 {
					continue
				}
				if isInclusionDateFuture(row[3], now, moscowLoc) || isExclusionDatePast(row[4], now, moscowLoc) {
					continue
				}
				entry := MediaIndividual{
					Number:        row[0],
					Name:          row[1],
					NameTranslit:  transliterate(normalizeName(row[1])),
					Activity:      row[2],
					InclusionDate: row[3],
					ExclusionDate: row[4],
				}
				tables.MediaIndividuals = append(tables.MediaIndividuals, entry)
			}
		case matchHeaders(header, headerForeignAgentIndividual):
			// In this table, the second column is Name and the third is Activity.
			for _, row := range rows[1:] {
				if len(row) < 6 {
					continue
				}
				if isInclusionDateFuture(row[5], now, moscowLoc) {
					continue
				}
				entry := ForeignAgentIndividual{
					Number:                row[0],
					Name:                  row[1],
					NameTranslit:          transliterate(normalizeName(row[1])),
					Activity:              row[2],
					ForeignSources:        row[3],
					PoliticalActivityInfo: row[4],
					InclusionDate:         row[5],
				}
				tables.ForeignAgentIndividuals = append(tables.ForeignAgentIndividuals, entry)
			}
		case matchHeaders(header, headerUnregisteredAssociation):
			for _, row := range rows[1:] {
				if len(row) < 5 {
					continue
				}
				if isInclusionDateFuture(row[2], now, moscowLoc) {
					continue
				}
				entry := UnregisteredAssociation{
					Number:           row[0],
					Name:             row[1],
					NameTranslit:     transliterate(normalizeName(row[1])),
					InclusionDate:    row[2],
					AssociationGoals: row[3],
					SourceInfo:       row[4],
				}
				tables.UnregisteredAssociations = append(tables.UnregisteredAssociations, entry)
			}
		}
	})

	return tables, nil
}
