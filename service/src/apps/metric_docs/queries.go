package metric_docs

const mainQuery = `
SELECT MAX(metric_name) AS metric_name,
       metric,
	   MAX(deprecated) AS deprecated,
       asset_type,
       STRING_AGG(body, E'\n\n' ORDER BY "order") AS body,
       MAX(last_modified) AS last_modified,
       MAX(logs::TEXT)::JSONB AS logs,
       MAX(tags) AS tags,
       MAX(category) AS category
FROM (
	SELECT DISTINCT ma.body AS body,
           MAX(metric_name) AS metric_name,
           metric,
		   MAX(deprecated::TEXT) AS deprecated,
           LOWER(asset_type) AS asset_type,
           MAX("order") AS "order",
           MAX(ma.last_modified)::TIMESTAMP WITHOUT TIME ZONE AS last_modified,
           JSONB_AGG(DISTINCT cl.payload::JSONB) AS logs,
	       STRING_AGG(DISTINCT tg.tag, ',') AS tags,
           MAX(pg.title) AS category
	FROM covid19.metric_reference AS mr
	  LEFT JOIN covid19.metric_asset_to_metric AS matm ON mr.metric = matm.metric_id
	  LEFT JOIN covid19.metric_asset           AS ma   ON ma.id = matm.asset_id
	  LEFT JOIN covid19.change_log_to_metric   AS cltm ON mr.metric = cltm.metric_id
	  LEFT JOIN (
		  SELECT cl_inner.id,
				 JSONB_BUILD_OBJECT(
					 'id', cl_inner.id::TEXT,
					 'heading', heading,
					 'date', cl_inner.date::TEXT,
					 'expiry', expiry::TEXT,
					 'type', tag
				 ) AS payload
		  FROM covid19.change_log AS cl_inner
		  LEFT JOIN covid19.tag AS t ON t.id = cl_inner.type_id
		  WHERE
			  cl_inner.date <= (
				  SELECT MAX(timestamp)::DATE
				  FROM covid19.release_reference AS rr
				  WHERE rr.released IS TRUE
			  )
		  ORDER BY date DESC
	  ) AS cl ON cltm.log_id = cl.id
	LEFT JOIN covid19.metric_tag AS mtg ON mr.metric = mtg.metric_id
	LEFT JOIN covid19.tag AS tg ON mtg.tag_id = tg.id
	LEFT JOIN covid19.page AS pg ON mr.category = pg.id
	WHERE mr.metric ILIKE ${metric_token}
	GROUP BY metric, asset_type, body
) AS dd
GROUP BY dd.metric, dd.asset_type;
`
