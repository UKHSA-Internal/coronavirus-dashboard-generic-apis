package metric_search

const query = `
SELECT metric,
       MAX(metric_name)           AS metric_name,
       MAX(title)                 AS category,
       CASE
           WHEN MAX(tag) NOTNULL THEN JSONB_AGG(tag ORDER BY tag)
           ELSE '[]'::JSONB
       END AS tags
FROM covid19.metric_reference AS mr
	LEFT OUTER JOIN covid19.page 		AS pg ON mr.category = pg.id
	LEFT OUTER JOIN covid19.metric_tag  AS mt ON mr.metric = mt.metric_id
	LEFT OUTER JOIN covid19.tag         AS tg ON mt.tag_id = tg.id
WHERE (
		   mr.metric ~* ('.*' || REGEXP_REPLACE($1, '\s+', '', 'g') || '.*')
	    OR mr.metric_name ~* ('.*' || $1 || '.*')
      )
  AND mr.released IS TRUE
GROUP BY metric
ORDER BY metric`

const metricTemplate = "%s"
