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

const partitionDatePlaceholder = "{{partition_date}}"

const byAreaType = `
SELECT area_type,
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
    SELECT MAX(title) AS category,
           metric,
           MAX(metric_name) AS metric_name,
           area_type
    FROM (SELECT area_id, metric_id FROM covid19.time_series_p{{partition_date}}_other GROUP BY area_id, metric_id) AS ts
        JOIN covid19.metric_reference AS mr ON ts.metric_id = mr.id
        JOIN covid19.page 			  AS pg ON mr.category = pg.id
        JOIN covid19.area_reference   AS ar ON ar.id = ts.area_id
    WHERE mr.released IS TRUE
      AND mr.id IN (SELECT DISTINCT metric_id FROM covid19.time_series_p{{partition_date}}_other)
    GROUP BY metric, area_type
    UNION
    (
        SELECT MAX(title)       AS category,
               metric,
               MAX(metric_name) AS metric_name,
               area_type
        FROM (SELECT area_id, metric_id FROM covid19.time_series_p{{partition_date}}_utla GROUP BY area_id, metric_id) AS ts
            JOIN covid19.metric_reference AS mr ON ts.metric_id = mr.id
            JOIN covid19.page 			  AS pg ON mr.category = pg.id
            JOIN covid19.area_reference   AS ar ON ar.id = ts.area_id
        WHERE mr.released IS TRUE
          AND mr.id IN (SELECT DISTINCT metric_id FROM covid19.time_series_p{{partition_date}}_utla)
        GROUP BY metric, area_type
    )
    UNION
    (
        SELECT MAX(title)       AS category,
               metric,
               MAX(metric_name) AS metric_name,
               area_type
        FROM (SELECT area_id, metric_id FROM covid19.time_series_p{{partition_date}}_ltla GROUP BY area_id, metric_id) AS ts
            JOIN covid19.metric_reference AS mr ON ts.metric_id = mr.id
            JOIN covid19.page 			  AS pg ON mr.category = pg.id
            JOIN covid19.area_reference   AS ar ON ar.id = ts.area_id
        WHERE mr.released IS TRUE
          AND mr.id IN (SELECT DISTINCT metric_id FROM covid19.time_series_p{{partition_date}}_ltla)
        GROUP BY metric, area_type
    )
    UNION
    (
        SELECT MAX(title)       AS category,
               metric,
               MAX(metric_name) AS metric_name,
               area_type
        FROM (SELECT area_id, metric_id FROM covid19.time_series_p{{partition_date}}_nhstrust GROUP BY area_id, metric_id) AS ts
            JOIN covid19.metric_reference AS mr ON ts.metric_id = mr.id
            JOIN covid19.page 			  AS pg ON mr.category = pg.id
            JOIN covid19.area_reference   AS ar ON ar.id = ts.area_id
        WHERE mr.released IS TRUE
          AND mr.id IN (SELECT DISTINCT metric_id FROM covid19.time_series_p{{partition_date}}_nhstrust)
        GROUP BY metric, area_type
    )
    UNION
    (
        SELECT MAX(title)       AS category,
               metric,
               MAX(metric_name) AS metric_name,
               area_type
        FROM (SELECT area_id, metric_id FROM covid19.time_series_p{{partition_date}}_msoa GROUP BY area_id, metric_id) AS ts
            JOIN covid19.metric_reference AS mr ON ts.metric_id = mr.id
            JOIN covid19.page 			  AS pg ON mr.category = pg.id
            JOIN covid19.area_reference   AS ar ON ar.id = ts.area_id
        WHERE mr.released IS TRUE
          AND mr.id IN (SELECT DISTINCT metric_id FROM covid19.time_series_p{{partition_date}}_utla)
        GROUP BY metric, area_type
    )
) AS df
GROUP BY area_type;`
