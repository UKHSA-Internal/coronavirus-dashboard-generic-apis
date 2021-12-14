package metric_availability

const query = `
SELECT metric, MAX(date)::DATE::TEXT AS last_update, MAX(deprecated) AS deprecated
FROM  %s AS ts
JOIN covid19.area_reference AS ar ON ar.id = ts.area_id
JOIN covid19.metric_reference AS mr ON mr.id = ts.metric_id
WHERE area_type = $1 %s
  AND date > (SELECT MAX(date) FROM %s) - INTERVAL '30 days'
  AND mr.released IS TRUE
GROUP BY metric
ORDER BY metric;
`

const areaCodeFilter = `AND area_code = $2`

const queryTable = "covid19.time_series_p%s_%s"
