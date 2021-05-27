package soa

const query = `
SELECT
	area_code AS "areaCode",
	area_name AS "areaName",
	area_type AS "areaType",
	DATE(date)::TEXT AS "date",
	(payload -> 'rollingSum') AS "rollingSum",
	(payload -> 'rollingRate') AS "rollingRate",
	(payload -> 'change') AS "change",
	(payload -> 'direction') AS "direction",
	(payload -> 'changePercentage') AS "changePercentage"
FROM %s AS ts
	JOIN covid19.area_reference AS ar ON ts.area_id = ar.id
WHERE area_code = $1
  AND date = ( SELECT MAX(date) FROM %s )
`

const queryTable = "covid19.time_series_p%s_%s"

const timestampQuery = `
SELECT DATE(MAX(timestamp))::TEXT AS date
FROM covid19.release_reference AS rr
	JOIN covid19.release_category AS rc ON rc.release_id = rr.id
WHERE released IS TRUE
  AND process_name = $1
`
