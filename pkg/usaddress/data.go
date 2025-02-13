package usaddress

var streetSuffixes = map[string]string{
	"ALLEY": "ALY", "ALY": "ALY",
	"AVENUE": "AVE", "AVE": "AVE",
	"BOULEVARD": "BLVD", "BLVD": "BLVD",
	"CIRCLE": "CIR", "CIR": "CIR",
	"COURT": "CT", "CT": "CT",
	"DRIVE": "DR", "DR": "DR",
	"EXPRESSWAY": "EXPY", "EXPY": "EXPY",
	"FREEWAY": "FWY", "FWY": "FWY",
	"HIGHWAY": "HWY", "HWY": "HWY",
	"LANE": "LN", "LN": "LN",
	"PASS":  "PASS",
	"PLACE": "PL", "PL": "PL",
	"ROAD": "RD", "RD": "RD",
	"STREET": "ST", "ST": "ST",
	"TERRACE": "TER", "TER": "TER",
	"TRAIL": "TRL", "TRL": "TRL",
	"WAY": "WAY",
}

var directionalMap = map[string]string{
	"NORTH":     "N",
	"SOUTH":     "S",
	"EAST":      "E",
	"WEST":      "W",
	"NORTHEAST": "NE",
	"NORTHWEST": "NW",
	"SOUTHEAST": "SE",
	"SOUTHWEST": "SW",
	"NE":        "NE",
	"NW":        "NW",
	"SE":        "SE",
	"SW":        "SW",
}

var stateNamesAndAbbreviations = map[string]string{
	"ALABAMA": "AL", "AL": "AL",
	"ALASKA": "AK", "AK": "AK",
	"ARIZONA": "AZ", "AZ": "AZ",
	"ARKANSAS": "AR", "AR": "AR",
	"CALIFORNIA": "CA", "CA": "CA",
	"COLORADO": "CO", "CO": "CO",
	"CONNECTICUT": "CT", "CT": "CT",
	"DELAWARE": "DE", "DE": "DE",
	"FLORIDA": "FL", "FL": "FL",
	"GEORGIA": "GA", "GA": "GA",
	"HAWAII": "HI", "HI": "HI",
	"IDAHO": "ID", "ID": "ID",
	"ILLINOIS": "IL", "IL": "IL",
	"INDIANA": "IN", "IN": "IN",
	"IOWA": "IA", "IA": "IA",
	"KANSAS": "KS", "KS": "KS",
	"KENTUCKY": "KY", "KY": "KY",
	"LOUISIANA": "LA", "LA": "LA",
	"MAINE": "ME", "ME": "ME",
	"MARYLAND": "MD", "MD": "MD",
	"MASSACHUSETTS": "MA", "MA": "MA",
	"MICHIGAN": "MI", "MI": "MI",
	"MINNESOTA": "MN", "MN": "MN",
	"MISSISSIPPI": "MS", "MS": "MS",
	"MISSOURI": "MO", "MO": "MO",
	"MONTANA": "MT", "MT": "MT",
	"NEBRASKA": "NE", "NE": "NE",
	"NEVADA": "NV", "NV": "NV",
	"NEW HAMPSHIRE": "NH", "NH": "NH",
	"NEW JERSEY": "NJ", "NJ": "NJ",
	"NEW MEXICO": "NM", "NM": "NM",
	"NEW YORK": "NY", "NY": "NY",
	"NORTH CAROLINA": "NC", "NC": "NC",
	"NORTH DAKOTA": "ND", "ND": "ND",
	"OHIO": "OH", "OH": "OH",
	"OKLAHOMA": "OK", "OK": "OK",
	"OREGON": "OR", "OR": "OR",
	"PENNSYLVANIA": "PA", "PA": "PA",
	"RHODE ISLAND": "RI", "RI": "RI",
	"SOUTH CAROLINA": "SC", "SC": "SC",
	"SOUTH DAKOTA": "SD", "SD": "SD",
	"TENNESSEE": "TN", "TN": "TN",
	"TEXAS": "TX", "TX": "TX",
	"UTAH": "UT", "UT": "UT",
	"VERMONT": "VT", "VT": "VT",
	"VIRGINIA": "VA", "VA": "VA",
	"WASHINGTON": "WA", "WA": "WA",
	"WEST VIRGINIA": "WV", "WV": "WV",
	"WISCONSIN": "WI", "WI": "WI",
	"WYOMING": "WY", "WY": "WY",
}

var stateAbbreviations = map[string]bool{
	"AL": true, "AK": true, "AZ": true, "AR": true, "CA": true, "CO": true, "CT": true,
	"DE": true, "FL": true, "GA": true, "HI": true, "ID": true, "IL": true, "IN": true,
	"IA": true, "KS": true, "KY": true, "LA": true, "ME": true, "MD": true, "MA": true,
	"MI": true, "MN": true, "MS": true, "MO": true, "MT": true, "NE": true, "NV": true,
	"NH": true, "NJ": true, "NM": true, "NY": true, "NC": true, "ND": true, "OH": true,
	"OK": true, "OR": true, "PA": true, "RI": true, "SC": true, "SD": true, "TN": true,
	"TX": true, "UT": true, "VT": true, "VA": true, "WA": true, "WV": true, "WI": true,
	"WY": true, "DC": true,
	"AS": true, "GU": true, "MP": true, "PR": true, "VI": true,
}

var secondaryUnitDesignators = map[string]string{
	"APARTMENT":  "APT",
	"APT":        "APT",
	"BUILDING":   "BLDG",
	"BLDG":       "BLDG",
	"FLOOR":      "FL",
	"FL":         "FL",
	"SUITE":      "STE",
	"STE":        "STE",
	"UNIT":       "UNIT",
	"ROOM":       "RM",
	"RM":         "RM",
	"DEPARTMENT": "DEPT",
	"DEPT":       "DEPT",
	"#":          "#",
}
