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
		healthcheck.Handler,
	},
	{
		"code",
		`/generic/code/{area_type:[a-zA-Z]{4,10}}/{area_code:[a-zA-Z0-9+%\s]{3,12}}`,
		code.Handler,
	},
	{
		"soa",
		`/generic/soa/{area_type:[a-zA-Z]{4,10}}/{area_code:[a-zA-Z0-9]{3,10}}`,
		soa.Handler,
	},
	{
		"area",
		`/generic/area/{area_type:[a-zA-Z]{4,10}}`,
		areaByType.Handler,
	},
	{
		"page_areas",
		`/generic/page_areas/{page:[a-zA-Z]{3,12}}`,
		pageArea.Handler,
	},
	{
		"page_areas_with_type",
		`/generic/page_areas/{page:[a-zA-Z]{3,12}}/{area_type:[a-zA-Z]{2,12}}`,
		pageArea.Handler,
	},
} // routes
