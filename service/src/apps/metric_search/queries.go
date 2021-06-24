package metric_search

const query = `
SELECT metric,
       MAX(metric_name)           AS metric_name,
       MAX(title)                 AS category,
       JSON_AGG(tag ORDER BY tag) AS tags
FROM covid19.metric_reference AS mr
	JOIN covid19.page 			  AS pg ON mr.category = pg.id
	JOIN covid19.metric_tag       AS mt ON mr.metric = mt.metric_id
	JOIN covid19.tag              AS tg ON mt.tag_id = tg.id
WHERE mr.metric ~* ('.*' || $1 || '.*')
  AND mr.released IS TRUE
GROUP BY metric
ORDER BY metric`

const metricTemplate = "%s"
