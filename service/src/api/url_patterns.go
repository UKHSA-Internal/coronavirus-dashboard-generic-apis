package api

import (
	"generic_apis/apps/areaByType"
	"generic_apis/apps/code"
	"generic_apis/apps/healthcheck"
	"generic_apis/apps/pageArea"
	"generic_apis/apps/soa"
)

var urlPatterns = []routeEntry{
	{
		"healthcheck",
		`/generic/healthcheck`,
		[]string{},
		healthcheck.Handler,
	},
	{
		"code",
		`/generic/code/{area_type:[a-zA-Z]{4,10}}/{area_code:[a-zA-Z0-9+%\s]{3,12}}`,
		[]string{},
		code.Handler,
	},
	{
		"soa",
		`/generic/soa/{area_type:[a-zA-Z]{4,10}}/{area_code:[a-zA-Z0-9]{3,10}}/{metric:[a-zA-Z2860]{5,120}}`,
		[]string{"date", `200\d-[01]\d-[0123]\d`},
		soa.Handler,
	},
	{
		"area",
		`/generic/area/{area_type:[a-zA-Z]{4,10}}`,
		[]string{},
		areaByType.Handler,
	},
	{
		"page_areas",
		`/generic/page_areas/{page:[a-zA-Z]{3,12}}`,
		[]string{},
		pageArea.Handler,
	},
	{
		"page_areas_with_type",
		`/generic/page_areas/{page:[a-zA-Z]{3,12}}/{area_type:[a-zA-Z]{2,12}}`,
		[]string{},
		pageArea.Handler,
	},
} // routes
