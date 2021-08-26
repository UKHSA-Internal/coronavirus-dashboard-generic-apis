package metric_docs

const mainQuery = `
SELECT MAX(metric_name) AS metric_name,
       metric,
	   LOWER(asset_type) AS asset_type,
       STRING_AGG(ma.body, E'\n\n' ORDER BY "order") AS body,
       MAX(ma.last_modified)::TIMESTAMP WITHOUT TIME ZONE AS last_modified
FROM covid19.metric_reference AS mr
  JOIN covid19.metric_asset_to_metric AS matm ON mr.metric = matm.metric_id
  JOIN covid19.metric_asset           AS ma   ON ma.id = matm.asset_id
WHERE LOWER(mr.metric) = LOWER(${metric_token})
GROUP BY metric, asset_type;
`
