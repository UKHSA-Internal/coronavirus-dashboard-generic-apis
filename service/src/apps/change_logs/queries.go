package change_logs

const (
	queryToken = "{token_id}"

	paginationToken = "{#pagination#}"

	filtersToken = "{filters}"
)

// Component queries
const (
	recordMonths = `SELECT DISTINCT date_trunc('month', cl.date)::DATE::TEXT AS date ` +
		`FROM covid19.change_log AS cl ` +
		`ORDER BY date DESC;`

	recordTypes = `SELECT tag ` +
		`FROM covid19.tag AS t ` +
		`WHERE t.association = 'CHANGE LOGS' ` +
		`ORDER BY tag DESC;`

	recordTitles = `SELECT title ` +
		`FROM covid19.page AS p ` +
		`ORDER BY title DESC;`
)

// filter queries parts
const (
	filtersQuery = `WHERE %s`

	paginationQuery = `LIMIT 20 OFFSET %d`

	releaseFilter = `date <= (` +
		`SELECT MAX(timestamp)::DATE ` +
		`FROM covid19.release_reference AS rr ` +
		`WHERE rr.released IS TRUE)`
)

// filter parts
const (
	searchFilter = `(to_tsvector('english'::REGCONFIG, cl.body) @@ (SELECT t::TSQUERY FROM search_token) ` +
		`OR to_tsvector('english'::REGCONFIG, cl.heading) @@ (SELECT t::TSQUERY FROM search_token))`

	titleFilter = `LOWER(p.title) = LOWER(${token_id})`

	typeFilter = `LOWER(t.tag) = LOWER(${token_id})`

	dateFilter = `date::DATE BETWEEN ${token_id}::DATE AND ${token_id}::DATE + INTERVAL '1 month'`

	metricFilter = `(mr.metric ISNULL OR (mr.released IS TRUE AND (mr.deprecated ISNULL OR mr.deprecated <= cl.date)))`

	recordFilter = `cl.id = ${token_id}`
)

var queryParamFilters = map[string]string{
	"search": searchFilter,
	"title":  titleFilter,
	"type":   typeFilter,
	"date":   dateFilter,
	"record": recordFilter,
}

var componentQueries = map[string]string{
	"titles": recordTitles,
	"types":  recordTypes,
	"dates":  recordMonths,
}

// Query for getting all logs up to and
// including the latest release date.
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
                       THEN ARRAY_AGG(DISTINCT (
                           JSONB_BUILD_OBJECT('metric', mr.metric, 'metric_name', mr.metric_name)
                       ))
                   ELSE '{}'::JSONB[]
               END AS metrics
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

// Query for getting the same data as `simpleQuery`
// with tokenised text search.
// Tokenisation is performed as a part of the query.
const searchQuery = `
WITH
	search_token AS (
		SELECT regexp_replace(
			plainto_tsquery('english'::REGCONFIG, ${token_id})::TEXT,
			'[''](.*?)([''])', '\1:*',
			'g'
		) AS t
	),
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
                           THEN ARRAY_AGG(DISTINCT (
                               JSONB_BUILD_OBJECT('metric', mr.metric, 'metric_name', mr.metric_name)
                           ))
                       ELSE '{}'::JSONB[]
                   END AS metrics,
                   cl.heading,
                   MAX(cl.body)    AS body,
                   ROUND(
                       ts_rank(
                           to_tsvector('english'::REGCONFIG, MAX(cl.body)),
						   (SELECT t::TSQUERY FROM search_token)
                       )::NUMERIC +
                       ts_rank(
                           to_tsvector('english'::REGCONFIG, cl.heading),
						   (SELECT t::TSQUERY FROM search_token)
					   )::NUMERIC * 2,
                       5
                   )::FLOAT        AS rank
            FROM covid19.change_log AS cl
              LEFT JOIN covid19.tag                  AS t ON t.id = cl.type_id
              LEFT JOIN covid19.change_log_to_metric AS cltm ON cltm.log_id = cl.id
              LEFT JOIN covid19.metric_reference     AS mr ON mr.metric = cltm.metric_id
              LEFT JOIN covid19.change_log_to_page   AS cltp ON cltp.log_id = cl.id
              LEFT JOIN covid19.page                 AS p ON p.id = cltp.page_id
            {filters}
            GROUP BY cl.heading, cl.date, t.tag, cl.high_priority, cl.display_banner
        ) AS df
        WHERE rank > 0
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

const recordQuery = `
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
                       THEN ARRAY_AGG(DISTINCT (
                           JSONB_BUILD_OBJECT('metric', mr.metric, 'metric_name', mr.metric_name)
                       ))
                   ELSE '{}'::JSONB[]
               END AS metrics
        FROM covid19.change_log AS cl
          LEFT JOIN covid19.change_log_to_metric  AS cltm ON cltm.log_id = cl.id
          LEFT JOIN covid19.tag                   AS t ON t.id = cl.type_id
          LEFT JOIN covid19.metric_reference      AS mr   ON mr.metric = cltm.metric_id
          LEFT JOIN covid19.change_log_to_page    AS cltp ON cltp.log_id = cl.id
          LEFT JOIN covid19.page                  AS p    ON p.id = cltp.page_id
		{filters}
        GROUP BY cl.date, cl.heading, t.tag, cl.high_priority, cl.display_banner
        ORDER BY cl.date DESC
    )
SELECT JSONB_AGG(data.*) AS data
FROM data;
`

const feedQuery = `
SELECT cl.id::TEXT          AS guid,
       cl.timestamp_created AS date,
       cl.heading           AS title,
       cl.body              AS description,
       UPPER(t.tag)         AS category
FROM covid19.change_log AS cl
  LEFT JOIN covid19.tag AS t ON t.id = cl.type_id
WHERE
	cl.date <= (
		SELECT MAX(timestamp)::DATE
		FROM covid19.release_reference AS rr
		WHERE rr.released IS TRUE
	)
ORDER BY cl.date DESC;
`
