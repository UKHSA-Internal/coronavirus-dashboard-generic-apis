package change_log

const simpleQuery = `
SELECT cl.date,
       cl.heading,
       MAX(cl.body)    AS body,
       MAX(cl.details) AS details,
       cl.high_priority,
	   cl.display_banner,
       t.tag 		   AS type,
       MAX(p.title)    AS page_title,
       JSONB_AGG(
           jsonb_build_object(
               'metric', mr.metric,
               'metric_name', mr.metric_name
           )
       )               AS metrics
FROM covid19.change_log AS cl
  LEFT JOIN covid19.change_log_to_metric  AS cltm ON cltm.log_id = cl.id
  LEFT JOIN covid19.tag                   AS t ON t.id = cl.type_id
  LEFT JOIN covid19.metric_reference      AS mr   ON mr.metric = cltm.metric_id
  LEFT JOIN covid19.change_log_to_page    AS cltp ON cltp.log_id = cl.id
  LEFT JOIN covid19.page                  AS p    ON p.id = cltp.page_id
{filters}
GROUP BY cl.date, cl.heading, t.tag, cl.high_priority, cl.display_banner
ORDER BY cl.date DESC
{pagination};
`

const searchQuery = `
SELECT *
FROM (
         SELECT cl.date,
                MAX(cl.details)  AS details,
                cl.high_priority,
                cl.display_banner,
                t.tag           AS type,
                MAX(cl.details) AS details,
                JSONB_AGG(
                    jsonb_build_object(
                        'metric', mr.metric,
                        'metric_name', mr.metric_name
                    )
                )               AS metrics,
                MAX(p.title)    As page_title,
                ts_headline(
                    cl.heading,
                    plainto_tsquery('english', ${token_id})
                )::TEXT         AS heading,
                ts_headline(
                    MAX(cl.body),
                    plainto_tsquery('english', ${token_id})
                )::TEXT         AS body,
                ROUND(
                    ts_rank(to_tsvector('english', MAX(cl.body)), plainto_tsquery('english', ${token_id}))::NUMERIC +
                    ts_rank(to_tsvector('english', cl.heading),  plainto_tsquery('english', ${token_id}))::NUMERIC * 2,
                    5
                )               AS rank
         FROM covid19.change_log AS cl
           LEFT JOIN covid19.tag                  AS t ON t.id = cl.type_id
           LEFT JOIN covid19.change_log_to_metric AS cltm ON cltm.log_id = cl.id
           LEFT JOIN covid19.metric_reference     AS mr ON mr.metric = cltm.metric_id
           LEFT JOIN covid19.change_log_to_page   AS cltp ON cltp.log_id = cl.id
           LEFT JOIN covid19.page                 AS p ON p.id = cltp.page_id
         GROUP BY cl.heading, cl.date, t.tag, cl.high_priority, cl.display_banner
    ) AS df
    {filters}
ORDER BY rank DESC, df.date DESC
{pagination};
`

// const recordMonths = `
// SELECT DISTINCT date_trunc('month', cl.date) AS date
// FROM covid19.change_log AS cl
// ORDER BY date DESC
// `

const filtersQuery = `WHERE %s`
const paginationQuery = "LIMIT 20 OFFSET %d"

var queryParamFilters = map[string]string{
	"search": `rank > 0.01`,
	"type":   `t.tag = ${token_id}`,
	"date":   `date BETWEEN ${token_id}::DATE AND ${token_id}::DATE + INTERVAL '1 month'`,
}
