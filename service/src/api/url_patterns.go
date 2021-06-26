package api

import (
	"generic_apis/apps/area_by_type"
	"generic_apis/apps/code"
	"generic_apis/apps/healthcheck"
	"generic_apis/apps/metric_availability"
	"generic_apis/apps/metric_search"
	"generic_apis/apps/page_area"
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
		area_by_type.Handler,
	},
	{
		"page_areas",
		`/generic/page_areas/{page:[a-zA-Z]{3,12}}`,
		[]string{},
		page_area.Handler,
	},
	{
		"metric_search",
		`/generic/metrics`,
		[]string{"search", `[a-zA-Z2860\s]{2,120}`, "category", `[a-zA-Z]{2,120}`, "tags", `[a-zA-Z0-9,]{2,40}`},
		metric_search.Handler,
	},
	{
		"metric_search_with_category",
		`/generic/metrics`,
		[]string{"search", `[a-zA-Z2860\s]{2,120}`, "category", `[a-zA-Z]{2,120}`},
		metric_search.Handler,
	},
	{
		"metric_search_with_tags",
		`/generic/metrics`,
		[]string{"search", `[a-zA-Z2860\s]{2,120}`, "tags", `[a-zA-Z0-9,]{2,40}`},
		metric_search.Handler,
	},
	{
		"metric_search_with_category_and_tag",
		`/generic/metrics`,
		[]string{"search", `[a-zA-Z2860\s]{2,120}`, "category", `[a-zA-Z]{2,120}`, "tags", `[a-zA-Z0-9,]{2,40}`},
		metric_search.Handler,
	},
	{
		"page_areas_with_type",
		`/generic/page_areas/{page:[a-zA-Z]{3,12}}/{area_type:[a-zA-Z]{2,12}}`,
		[]string{},
		page_area.Handler,
	},
	{
		"metric_availability_by_area_type",
		`/generic/metric_availability/{area_type:[a-zA-Z]{2,12}}`,
		[]string{"date", `200\d-[01]\d-[0123]\d`},
		metric_availability.Handler,
	},
	{
		"metric_availability_by_area",
		`/generic/metric_availability/{area_type:[a-zA-Z]{2,12}}/{area_code:[a-zA-Z0-9]{3,10}}`,
		[]string{"date", `200\d-[01]\d-[0123]\d`},
		metric_availability.Handler,
	},
} // routes
