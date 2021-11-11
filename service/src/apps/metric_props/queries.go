package metric_props

const byCategory = `
SELECT category,
       JSONB_AGG(
           JSONB_BUILD_OBJECT(
               'metric', metric,
               'metric_name', metric_name,
               'tag', tags
           )
           ORDER BY metric
       ) AS payload
FROM (
	SELECT title AS category,
		   metric,
           metric_name,
           CASE
               WHEN MAX(tag) NOTNULL THEN JSONB_AGG(tag ORDER BY tag)
               ELSE '[]'::JSONB
           END AS tags
    FROM covid19.metric_reference AS mr
		LEFT OUTER JOIN covid19.page AS pg ON mr.category = pg.id
		LEFT OUTER JOIN covid19.metric_tag AS mt ON mr.metric = mt.metric_id
		LEFT OUTER JOIN covid19.tag AS tg ON mt.tag_id = tg.id
    WHERE mr.released IS TRUE
	  AND title NOTNULL
    GROUP BY title, metric, metric_name
) AS df
GROUP BY category;`

const byTag = `
SELECT tag,
       JSONB_AGG(
           JSONB_BUILD_OBJECT(
               'metric', metric,
               'metric_name', metric_name,
               'category', category,
               'tag', array(
                   SELECT DISTINCT tg.tag
                   FROM covid19.tag AS tg
                   LEFT OUTER JOIN covid19.metric_tag AS mt ON tg.id = mt.tag_id
                   WHERE metric_id = metric
               )
		   )
           ORDER BY metric
       ) AS payload
FROM (
    SELECT tag,
           title AS category,
           metric,
           metric_name
    FROM covid19.metric_reference AS mr
             LEFT OUTER JOIN covid19.page AS pg ON mr.category = pg.id
             LEFT OUTER JOIN covid19.metric_tag AS mt ON mr.metric = mt.metric_id
             LEFT OUTER JOIN covid19.tag AS tg ON mt.tag_id = tg.id
    WHERE mr.released IS TRUE
      AND tag NOTNULL
    GROUP BY tag, title, metric, metric_name
) AS df
GROUP BY tag;`
