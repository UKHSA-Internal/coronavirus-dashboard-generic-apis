package soa

const query = `
SELECT
	area_code 		 AS "areaCode",
	area_name 		 AS "areaName",
	area_type 		 AS "areaType",
	DATE(date)::TEXT AS "date",
	metric,
	payload
FROM %s AS ts
	JOIN covid19.area_reference   AS ar ON ts.area_id = ar.id
	JOIN covid19.metric_reference AS mr ON ts.metric_id = mr.id
WHERE area_code = $1
  AND metric = $2
  AND date = ( SELECT MAX(date) FROM %s )
`

const queryTable = "covid19.time_series_p%s_%s"
