package utils

var (
	AreaTypes = map[string]string{
		"postcode":  "postcode",
		"msoa":      "msoa",
		"nhstrust":  "nhsTrust",
		"nhsregion": "nhsRegion",
		"utla":      "utla",
		"ltla":      "ltla",
		"region":    "region",
		"nation":    "nation",
		"overview":  "overview",
		"la":        "ANY('{utla,ltla}'::VARCHAR[])",
	}

	ReleaseCategories = map[string]string{
		"msoa":      "MSOA",
		"nhsTrust":  "MAIN",
		"nhsRegion": "MAIN",
		"utla":      "MAIN",
		"ltla":      "MAIN",
		"region":    "MAIN",
		"nation":    "MAIN",
		"overview":  "MAIN",
	}

	AreaPartitions = map[string]string{
		"msoa":      "msoa",
		"nhsTrust":  "nhstrust",
		"nhsRegion": "other",
		"utla":      "utla",
		"ltla":      "ltla",
		"region":    "other",
		"nation":    "other",
		"overview":  "other",
	}
)
