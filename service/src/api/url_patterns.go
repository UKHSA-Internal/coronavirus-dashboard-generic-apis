package api

import (
	"time"

	"generic_apis/apps/announcements"
	"generic_apis/apps/area_by_type"
	"generic_apis/apps/area_name"
	"generic_apis/apps/change_logs"
	"generic_apis/apps/code"
	"generic_apis/apps/log_banners"
	"generic_apis/apps/metric_areas"
	"generic_apis/apps/metric_availability"
	"generic_apis/apps/metric_docs"
	"generic_apis/apps/metric_props"
	"generic_apis/apps/metric_search"
	"generic_apis/apps/page_area"
	"generic_apis/apps/soa"
	"generic_apis/apps/utils"
)

var UrlPatterns = &[]utils.RouteEntry{
	{
		"code",
		`/generic/code/{area_type:[a-zA-Z]{4,10}}/{area_code:[a-zA-Z0-9+%\s]{3,12}}`,
		[]string{},
		code.Handler,
		time.Minute * 5,
	},
	{
		"area_name",
		`/generic/area/{area_type:[a-zA-Z]{4,10}}/area_name/{area_name:[A-Za-z0-9,'.\s()-]{5,120}}`,
		[]string{},
		area_name.Handler,
		time.Hour * 12,
	},
	{
		"soa",
		`/generic/soa/{area_type:[a-zA-Z]{4,10}}/{area_code:[a-zA-Z0-9]{3,10}}/{metric:[a-zA-Z2860]{5,120}}`,
		[]string{"date", `200\d-[01]\d-[0123]\d`},
		soa.Handler,
		0,
	},
	{
		"area",
		`/generic/area/{area_type:[a-zA-Z]{4,10}}`,
		[]string{},
		area_by_type.Handler,
		time.Hour * 12,
	},
	{
		"page_areas",
		`/generic/page_areas/{page:[a-zA-Z]{3,12}}`,
		[]string{},
		page_area.Handler,
		time.Minute * 4,
	},
	{
		"metric_search",
		`/generic/metrics`,
		[]string{"search", `[a-zA-Z2860\s]{2,120}`, "category", `[a-zA-Z]{2,120}`, "tags", `[a-zA-Z0-9,]{2,40}`},
		metric_search.Handler,
		time.Minute * 30,
	},
	{
		"metric_search_with_category",
		`/generic/metrics`,
		[]string{"search", `[a-zA-Z2860\s]{2,120}`, "category", `[a-zA-Z]{2,120}`},
		metric_search.Handler,
		time.Hour * 2,
	},
	{
		"metric_search_category_only",
		`/generic/metrics`,
		[]string{"category", `[a-zA-Z]{2,120}`},
		metric_search.Handler,
		time.Hour * 2,
	},
	{
		"metric_search_with_tags",
		`/generic/metrics`,
		[]string{"search", `[a-zA-Z2860\s]{2,120}`, "tags", `[a-zA-Z0-9,]{2,40}`},
		metric_search.Handler,
		time.Hour * 2,
	},
	{
		"metric_search_tags_only",
		`/generic/metrics`,
		[]string{"tags", `[a-zA-Z0-9\s,]{2,40}`},
		metric_search.Handler,
		time.Hour * 2,
	},
	{
		"metric_search_with_category_and_tag",
		`/generic/metrics`,
		[]string{"search", `[a-zA-Z2860\s]{2,120}`, "category", `[a-zA-Z]{2,120}`, "tags", `[a-zA-Z0-9,]{2,40}`},
		metric_search.Handler,
		time.Hour * 2,
	},
	{
		"metric_search_by_category_and_tag",
		`/generic/metrics`,
		[]string{"category", `[a-zA-Z]{2,120}`, "tags", `[a-zA-Z0-9\s,]{2,40}`},
		metric_search.Handler,
		time.Hour * 2,
	},
	{
		"metric_props",
		`/generic/metrics/props`,
		[]string{`by`},
		metric_props.Handler,
		time.Hour * 2,
	},
	{
		"page_areas_with_type",
		`/generic/page_areas/{page:[a-zA-Z]{3,12}}/{area_type:[a-zA-Z]{2,12}}`,
		[]string{},
		page_area.Handler,
		time.Hour * 1,
	},
	{
		"metric_availability_by_area_type",
		`/generic/metric_availability/{area_type:[a-zA-Z]{2,12}}`,
		[]string{"date", `202\d-[01]\d-[0123]\d`},
		metric_availability.Handler,
		time.Minute * 5,
	},
	{
		"metric_availability_by_area",
		`/generic/metric_availability/{area_type:[a-zA-Z]{2,12}}/{area_code:[a-zA-Z0-9]{3,10}}`,
		[]string{"date", `202\d-[01]\d-[0123]\d`},
		metric_availability.Handler,
		time.Minute * 5,
	},
	{
		"change_logs",
		`/generic/change_logs`,
		[]string{},
		change_logs.Handler,
		0,
	},
	{
		"change_log_item",
		`/generic/change_logs/log/{id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}`,
		[]string{},
		change_logs.Handler,
		0,
	},
	{
		"change_logs_components",
		`/generic/change_logs/components/{component:dates|types|titles}`,
		[]string{},
		change_logs.Handler,
		0,
	},
	{
		"change_logs_feed",
		`/generic/change_logs/{type:rss|atom}.xml`,
		[]string{},
		change_logs.FeedHandler,
		0,
	},
	{
		"change_logs_single_month",
		`/generic/change_logs/{date:202\d-[01]\d}`,
		[]string{},
		change_logs.Handler,
		0,
	},
	{
		"log_banners",
		`/generic/log_banners/{date:202\d-[01]\d-[0123]\d}/{page:[A-Za-z\s:\-']{5,40}}/{area_type:[a-zA-Z]{2,12}}/{area_name:[A-Za-z0-9,'.\s()-]{5,120}}`,
		[]string{},
		log_banners.Handler,
		0,
	},
	{
		"announcements",
		`/generic/announcements`,
		[]string{},
		announcements.Handler,
		0,
	},
	{
		"announcement_item",
		`/generic/announcements/{id:[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}}`,
		[]string{},
		announcements.Handler,
		0,
	},
	{
		"announcements_feed",
		`/generic/announcements/{type:rss|atom}.xml`,
		[]string{},
		announcements.FeedHandler,
		0,
	},
	{
		"latest_announcement",
		`/generic/announcements/latest`,
		[]string{},
		announcements.Handler,
		0,
	},
	{
		"metric_doc",
		`/generic/metrics/{metric:[a-zA-Z2860\s]{2,120}}/doc`,
		[]string{},
		metric_docs.Handler,
		time.Minute * 30,
	},
	{
		"metric_areas",
		`/generic/metrics/{metric:[a-zA-Z2860\s]{2,120}}/areas/{date:202\d-[01]\d-[0123]\d}`,
		[]string{},
		metric_areas.Handler,
		time.Minute * 30,
	},
} // routes
