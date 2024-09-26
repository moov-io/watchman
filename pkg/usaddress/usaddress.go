package usaddress

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// Address represents a standardized address
type Address struct {
	PrimaryNumber   string // House or building number
	StreetPredir    string // Pre-directional (e.g., N, S)
	StreetName      string
	StreetSuffix    string
	StreetPostdir   string // Post-directional (e.g., NW)
	SecondaryUnit   string // e.g., APT 101
	City            string
	State           string
	ZIPCode         string
	Plus4           string // Optional ZIP+4 code
	POBox           string // For PO Box addresses
	RuralRoute      string // For rural route addresses
	HighwayContract string // For highway contract route addresses
}

// String formats the Address into a standardized string representation.
func (a Address) String() string {
	var lines []string

	// Line 1: Delivery Address Line
	var deliveryLine string

	// Handle PO Box, Rural Route, or Highway Contract first
	if a.POBox != "" {
		deliveryLine = a.POBox
	} else if a.RuralRoute != "" {
		deliveryLine = a.RuralRoute
	} else if a.HighwayContract != "" {
		deliveryLine = a.HighwayContract
	} else {
		// Build street address
		var streetComponents []string

		if a.PrimaryNumber != "" {
			streetComponents = append(streetComponents, a.PrimaryNumber)
		}

		if a.StreetPredir != "" {
			streetComponents = append(streetComponents, a.StreetPredir)
		}

		if a.StreetName != "" {
			streetComponents = append(streetComponents, a.StreetName)
		}

		if a.StreetSuffix != "" {
			streetComponents = append(streetComponents, a.StreetSuffix)
		}

		if a.StreetPostdir != "" {
			streetComponents = append(streetComponents, a.StreetPostdir)
		}

		deliveryLine = strings.Join(streetComponents, " ")

		if a.SecondaryUnit != "" {
			deliveryLine = fmt.Sprintf("%s %s", deliveryLine, a.SecondaryUnit)
		}
	}

	if deliveryLine != "" {
		lines = append(lines, deliveryLine)
	}

	// Line 2: City, State ZIP Code
	var lastLineComponents []string

	if a.City != "" {
		lastLineComponents = append(lastLineComponents, a.City)
	}

	if a.State != "" {
		lastLineComponents = append(lastLineComponents, a.State)
	}

	if a.ZIPCode != "" {
		zipCode := a.ZIPCode
		if a.Plus4 != "" {
			zipCode = fmt.Sprintf("%s-%s", zipCode, a.Plus4)
		}
		lastLineComponents = append(lastLineComponents, zipCode)
	}

	lastLine := strings.Join(lastLineComponents, ", ")

	if lastLine != "" {
		lines = append(lines, lastLine)
	}

	return strings.Join(lines, "\n")
}

// Validate checks if the address is valid and returns an error explaining why it's invalid.
func (a Address) Validate() error {
	var errs []string

	// At least one of POBox, RuralRoute, HighwayContract, or street address must be present
	if a.POBox == "" && a.RuralRoute == "" && a.HighwayContract == "" {
		if a.PrimaryNumber == "" || a.StreetName == "" {
			errs = append(errs, "address must contain either a PO Box, rural route, highway contract, or a valid street address with a primary number and street name")
		}
	}

	// Validate ZIP Code
	if a.ZIPCode != "" {
		matched, _ := regexp.MatchString(`^\d{5}$`, a.ZIPCode)
		if !matched {
			errs = append(errs, "ZIP Code must be a 5-digit number")
		}
	} else {
		errs = append(errs, "ZIP Code is required")
	}

	// Validate Plus4 (if provided)
	if a.Plus4 != "" {
		matched, _ := regexp.MatchString(`^\d{4}$`, a.Plus4)
		if !matched {
			errs = append(errs, "Plus4 must be a 4-digit number")
		}
	}

	// Validate State
	if a.State != "" {
		if !stateAbbreviations[a.State] {
			errs = append(errs, fmt.Sprintf("State '%s' is not a valid US state abbreviation", a.State))
		}
	} else {
		errs = append(errs, "State is required")
	}

	// Validate City
	if a.City == "" {
		errs = append(errs, "City is required")
	}

	if len(errs) > 0 {
		return errors.New(strings.Join(errs, "; "))
	}
	return nil
}

// StandardizeAddress accepts an address string and returns a standard Address struct
func StandardizeAddress(addressStr string) Address {
	addressStr = strings.ReplaceAll(addressStr, "\r", "")
	addressStr = strings.ToUpper(addressStr)

	lines := strings.Split(addressStr, "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	lines = filterEmptyStrings(lines)

	// Filter out lines starting with 'C/O' or '℅' BEFORE removing punctuation
	var filteredLines []string
	for _, line := range lines {
		if strings.HasPrefix(line, "C/O") || strings.HasPrefix(line, "℅") {
			continue
		}
		filteredLines = append(filteredLines, line)
	}
	lines = filteredLines

	// Now remove punctuation
	for i := range lines {
		lines[i] = removePunctuation(lines[i])
	}

	addr := Address{}
	deliveryLines := []string{}
	lastLine := ""

	if len(lines) >= 2 {
		// Assume last line is city/state/ZIP
		lastLine = lines[len(lines)-1]
		// Collect delivery lines
		for i := len(lines) - 2; i >= 0; i-- {
			if hasNumbers(lines[i]) || isPOBoxLine(lines[i]) {
				deliveryLines = append([]string{lines[i]}, deliveryLines...)
			}
		}
	} else if len(lines) == 1 {
		// Try to split into delivery line and last line
		parts := strings.SplitN(lines[0], ",", 2)
		if len(parts) == 2 {
			deliveryLines = append(deliveryLines, strings.TrimSpace(parts[0]))
			lastLine = strings.TrimSpace(parts[1])
		} else {
			deliveryLines = append(deliveryLines, lines[0])
		}
	}

	// Parse the last line to get City, State, ZIP
	parseCityStateZip(lastLine, &addr)

	// Parse delivery lines
	for _, line := range deliveryLines {
		if isPOBoxLine(line) {
			addr.POBox = line
		} else {
			parseDeliveryLine(line, &addr)
		}
	}

	return addr
}

func parseCityStateZip(line string, addr *Address) {
	if line == "" {
		return
	}

	components := strings.Fields(line)
	if len(components) == 0 {
		return
	}

	// Assume ZIP code is the last element
	zipCode := strings.TrimSpace(components[len(components)-1])
	if isZipCode(zipCode) {
		addr.ZIPCode = zipCode
		components = components[:len(components)-1]
	}

	// Handle Plus4
	if strings.Contains(addr.ZIPCode, "-") {
		zipParts := strings.Split(addr.ZIPCode, "-")
		addr.ZIPCode = zipParts[0]
		addr.Plus4 = zipParts[1]
	}

	if len(components) >= 2 && isMilitaryState(components[len(components)-1]) {
		addr.State = components[len(components)-1]
		components = components[:len(components)-1]
		addr.City = strings.Join(components, " ")
	} else {
		// Existing logic for state and city
		stateFound := false
		for i := len(components); i > 0; i-- {
			possibleState := strings.Join(components[i-1:], " ")
			abbr := standardizeState(possibleState)
			if abbr != "" {
				addr.State = abbr
				addr.City = strings.Join(components[:i-1], " ")
				stateFound = true
				break
			}
		}
		if !stateFound {
			addr.City = strings.Join(components, " ")
		}
	}

	// Trim trailing commas and spaces from City
	addr.City = strings.TrimSpace(strings.TrimSuffix(addr.City, ","))
}

func parseDeliveryLine(line string, addr *Address) {
	components := strings.Fields(line)
	i := 0

	if i < len(components) {
		// Handle Military Addresses
		if isMilitaryAddress(components[i]) {
			addr.POBox = strings.Join(components, " ")
			return
		}

		// Handle PO Box
		if isPOBox(components[i]) {
			addr.POBox = parsePOBox(components[i:])
			return
		}

		// Handle Rural Route
		if isRuralRoute(components[i]) {
			addr.RuralRoute = strings.Join(components, " ")
			return
		}

		// Handle Highway Contract Route
		if isHighwayContract(components[i]) {
			addr.HighwayContract = strings.Join(components, " ")
			return
		}

		// Handle Primary Number (may be hyphenated or fractional)
		if isPrimaryNumber(components[i]) {
			addr.PrimaryNumber = components[i]
			i++
			// Check for fractional part
			if i < len(components) && isFraction(components[i]) {
				addr.PrimaryNumber += " " + components[i]
				i++
			}
		}

		// Handle Pre-directional
		if i < len(components) && directionalMap[components[i]] != "" {
			addr.StreetPredir = directionalMap[components[i]]
			i++
		}

		// Parse Street Name
		for i < len(components) && !isStreetSuffix(components[i]) && !isSecondaryUnitDesignator(components[i]) && directionalMap[components[i]] == "" {
			addr.StreetName += components[i] + " "
			i++
		}
		addr.StreetName = strings.TrimSpace(addr.StreetName)

		// Parse Street Suffix
		if i < len(components) && isStreetSuffix(components[i]) {
			addr.StreetSuffix = standardizeStreetSuffix(components[i])
			i++
		}

		// Handle Post-directional
		if i < len(components) && directionalMap[components[i]] != "" {
			addr.StreetPostdir = directionalMap[components[i]]
			i++
		}

		// After parsing Street Suffix and Post-directional
		if i < len(components) {
			unit := ""
			// Collect remaining components as SecondaryUnit
			for ; i < len(components); i++ {
				unit += components[i] + " "
			}
			addr.SecondaryUnit = strings.TrimSpace(unit)
		}
	}
}

func hasNumbers(s string) bool {
	return regexp.MustCompile(`\d`).MatchString(s)
}

func removePunctuation(s string) string {
	re := regexp.MustCompile(`[^\w\s#&-/]`)
	s = re.ReplaceAllString(s, "")
	// Ensure there's a space after '#' if followed by a digit
	s = regexp.MustCompile(`#(\d)`).ReplaceAllString(s, "# $1")
	return s
}

func isZipCode(s string) bool {
	s = strings.TrimSpace(s)
	matched, _ := regexp.MatchString(`^\d{5}(-\d{4})?$`, s)
	return matched
}

func isPOBox(s string) bool {
	return s == "PO" || s == "P.O" || s == "POBOX" || s == "P.O.BOX"
}

func isPOBoxLine(s string) bool {
	s = strings.TrimSpace(s)
	return strings.HasPrefix(s, "PO BOX") || strings.HasPrefix(s, "P.O BOX") || strings.HasPrefix(s, "P.O. BOX")
}

func parsePOBox(components []string) string {
	return strings.Join(components, " ")
}

func isRuralRoute(s string) bool {
	return s == "RR" || s == "RFD" || s == "RD"
}

func isHighwayContract(s string) bool {
	return s == "HC"
}

func isPrimaryNumber(s string) bool {
	matched, _ := regexp.MatchString(`^\d+([- ]\d+)?(\s*\d+\/\d+)?$`, s)
	return matched
}

func isFraction(s string) bool {
	matched, _ := regexp.MatchString(`^\d+\/\d+$`, s)
	return matched
}

func isStreetSuffix(s string) bool {
	_, exists := streetSuffixes[s]
	return exists
}

func standardizeStreetSuffix(s string) string {
	if val, exists := streetSuffixes[s]; exists {
		return val
	}
	return s
}

func isSecondaryUnitDesignator(s string) bool {
	_, exists := secondaryUnitDesignators[s]
	return exists
}

func standardizeSecondaryUnitDesignator(s string) string {
	if val, exists := secondaryUnitDesignators[s]; exists {
		return val
	}
	return s
}

func standardizeState(s string) string {
	s = strings.TrimSpace(strings.ToUpper(s))
	if val, exists := stateNamesAndAbbreviations[s]; exists {
		return val
	}
	return ""
}

func filterEmptyStrings(slice []string) []string {
	var filtered []string
	for _, str := range slice {
		if strings.TrimSpace(str) != "" {
			filtered = append(filtered, str)
		}
	}
	return filtered
}

func isMilitaryState(s string) bool {
	militaryStates := map[string]bool{
		"AE": true, "AA": true, "AP": true,
	}
	return militaryStates[s]
}

func isMilitaryAddress(s string) bool {
	return s == "PSC" || s == "UNIT" || s == "CMR"
}
