package metric_docs

const mainQuery = `
SELECT MAX(metric_name) AS metric_name,
       metric,
	   LOWER(asset_type) AS asset_type,
       STRING_AGG(ma.body, E'\n\n' ORDER BY "order") AS body,
       MAX(ma.last_modified)::TIMESTAMP WITHOUT TIME ZONE AS last_modified,
       JSONB_AGG(cl.payload::JSONB) AS logs
FROM covid19.metric_reference AS mr
  LEFT JOIN covid19.metric_asset_to_metric AS matm ON mr.metric = matm.metric_id
  LEFT JOIN covid19.metric_asset           AS ma   ON ma.id = matm.asset_id
  LEFT JOIN covid19.change_log_to_metric   AS cltm ON mr.metric = cltm.metric_id
  LEFT JOIN (
      SELECT id,
			 JSONB_BUILD_OBJECT(
			 	 'id', id::TEXT,
				 'date', date::TEXT,
				 'expiry', expiry::TEXT,
				 'heading', heading
			 ) AS payload
      FROM covid19.change_log
    ) AS cl   ON cltm.log_id = cl.id
WHERE mr.metric ILIKE ${metric_token}
GROUP BY metric, asset_type;
`
