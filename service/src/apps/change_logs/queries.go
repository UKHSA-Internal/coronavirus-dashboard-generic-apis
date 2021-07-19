package change_logs

const (
	queryToken      = "{token_id}"
	paginationToken = "{#pagination#}"
	filtersToken    = "{filters}"
)

const simpleQuery = `
WITH
    data AS (
        SELECT MAX(cl.id::TEXT) AS id,
               cl.date::TEXT,
               cl.heading,
               MAX(cl.body)    AS body,
               MAX(cl.details) AS details,
               cl.high_priority,
               cl.display_banner,
               t.tag 		   AS type,
               CASE
                   WHEN MAX(p.title) NOTNULL
                       THEN JSONB_BUILD_OBJECT(
                            'title', MAX(p.title),
                            'uri', MAX(p.uri)
                       )
                   ELSE '[]'::JSONB
               END             AS page,
               CASE
                   WHEN MAX(mr.metric) NOTNULL
                       THEN JSONB_AGG(
                           JSONB_BUILD_OBJECT(
                               'metric', mr.metric,
                               'metric_name', mr.metric_name
                           )
                       )
                   ELSE '[]'::JSONB
               END             AS metrics
        FROM covid19.change_log AS cl
          LEFT JOIN covid19.change_log_to_metric  AS cltm ON cltm.log_id = cl.id
          LEFT JOIN covid19.tag                   AS t ON t.id = cl.type_id
          LEFT JOIN covid19.metric_reference      AS mr   ON mr.metric = cltm.metric_id
          LEFT JOIN covid19.change_log_to_page    AS cltp ON cltp.log_id = cl.id
          LEFT JOIN covid19.page                  AS p    ON p.id = cltp.page_id
        {filters}
        GROUP BY cl.date, cl.heading, t.tag, cl.high_priority, cl.display_banner
        ORDER BY cl.date DESC
    ),
    payload AS (
        SELECT *
        FROM DATA
        {#pagination#}
    )
SELECT JSONB_AGG(payload.*) AS data,
       (SELECT COUNT(*) FROM data)::INT AS "total_length",
       (SELECT CEIL(COUNT(*) / 20.0) FROM data)::INT AS total_pages
FROM payload;
`

const searchQuery = `
WITH
    data AS (
        SELECT *
        FROM (
            SELECT MAX(cl.id::TEXT) AS id,
                   cl.date::TEXT,
                   cl.high_priority,
                   cl.display_banner,
                   t.tag           AS type,
                   MAX(cl.details) AS details,
                   CASE
                       WHEN MAX(p.title) NOTNULL
                           THEN JSONB_BUILD_OBJECT(
                                'title', MAX(p.title),
                                'uri', MAX(p.uri)
                           )
                       ELSE '[]'::JSONB
                   END             AS page,
                   CASE
                       WHEN MAX(mr.metric) NOTNULL
                           THEN JSONB_AGG(
                               JSONB_BUILD_OBJECT(
                                   'metric', mr.metric,
                                   'metric_name', mr.metric_name
                               )
                           )
                       ELSE '[]'::JSONB
                   END             AS metrics,
                   cl.heading,
                   MAX(cl.body)    AS body,
                   ROUND(
                       ts_rank(to_tsvector('english'::REGCONFIG, MAX(cl.body)), plainto_tsquery('english', ${token_id}))::NUMERIC +
                       ts_rank(to_tsvector('english'::REGCONFIG, cl.heading),  plainto_tsquery('english', ${token_id}))::NUMERIC * 2,
                       5
                   )::FLOAT        AS rank
            FROM covid19.change_log AS cl
              LEFT JOIN covid19.tag                  AS t ON t.id = cl.type_id
              LEFT JOIN covid19.change_log_to_metric AS cltm ON cltm.log_id = cl.id
              LEFT JOIN covid19.metric_reference     AS mr ON mr.metric = cltm.metric_id
              LEFT JOIN covid19.change_log_to_page   AS cltp ON cltp.log_id = cl.id
              LEFT JOIN covid19.page                 AS p ON p.id = cltp.page_id
            WHERE to_tsvector('english'::REGCONFIG, cl.body) @@ plainto_tsquery('english', ${token_id})
               OR to_tsvector('english'::REGCONFIG, cl.heading) @@ plainto_tsquery('english', ${token_id})
            GROUP BY cl.heading, cl.date, t.tag, cl.high_priority, cl.display_banner
        ) AS df
        {filters}
        ORDER BY rank DESC, df.date DESC
    ),
    payload AS (
        SELECT *
        FROM DATA
        {#pagination#}
    )
SELECT JSONB_AGG(payload.*) AS data,
       (SELECT COUNT(*) FROM data)::INT AS "total_length",
       (SELECT CEIL(COUNT(*) / 20.0) FROM data)::INT AS total_pages
FROM payload;
`

const recordMonths = `
SELECT DISTINCT date_trunc('month', cl.date)::DATE::TEXT AS date
FROM covid19.change_log AS cl
ORDER BY date DESC
`

const filtersQuery = `WHERE %s`
const paginationQuery = "LIMIT 20 OFFSET %d"

var queryParamFilters = map[string]string{
	"search": `rank > 0.01`,
	"title":  `LOWER(p.title) = LOWER(${token_id})`,
	"type":   `LOWER(p.tag) = LOWER(${token_id})`,
	"date":   `date::DATE BETWEEN ${token_id}::DATE AND ${token_id}::DATE + INTERVAL '1 month'`,
}
