package metric_docs

const mainQuery = `
SELECT MAX(metric_name) AS metric_name,
       metric,
       asset_type,
       STRING_AGG(body, E'\n\n' ORDER BY "order") AS body,
       MAX(last_modified) AS last_modified,
       MAX(logs::TEXT)::JSONB AS logs
FROM (
	SELECT DISTINCT ma.body AS body,
           MAX(metric_name) AS metric_name,
           metric,
           LOWER(asset_type) AS asset_type,
           MAX("order") AS "order",
           MAX(ma.last_modified)::TIMESTAMP WITHOUT TIME ZONE AS last_modified,
           JSONB_AGG(DISTINCT cl.payload::JSONB) AS logs
	FROM covid19.metric_reference AS mr
	  LEFT JOIN covid19.metric_asset_to_metric AS matm ON mr.metric = matm.metric_id
	  LEFT JOIN covid19.metric_asset           AS ma   ON ma.id = matm.asset_id
	  LEFT JOIN covid19.change_log_to_metric   AS cltm ON mr.metric = cltm.metric_id
	  LEFT JOIN (
		  SELECT cl_inner.id,
				 JSONB_BUILD_OBJECT(
					 'id', cl_inner.id::TEXT,
					 'heading', heading,
					 'date', date::TEXT,
					 'expiry', expiry::TEXT,
					 'type', tag
				 ) AS payload
		  FROM covid19.change_log AS cl_inner
		  LEFT JOIN covid19.tag AS t ON t.id = cl_inner.type_id
		  ORDER BY date DESC
	  ) AS cl ON cltm.log_id = cl.id
	WHERE mr.metric ILIKE ${metric_token}
	GROUP BY metric, asset_type, body
) AS dd
GROUP BY dd.metric, dd.asset_type;
`
