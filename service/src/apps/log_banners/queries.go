package log_banners

const mainQuery = `
SELECT cl.date::TEXT,
       cl.high_priority,
       t.tag           AS type,
       cl.heading,
       cl.body         AS body
FROM covid19.change_log AS cl
  LEFT JOIN covid19.tag                  AS t ON t.id = cl.type_id
  LEFT JOIN covid19.change_log_to_page   AS cltp ON cltp.log_id = cl.id
  LEFT JOIN covid19.page                 AS p ON p.id = cltp.page_id
WHERE cl.display_banner IS TRUE
  AND date = ${date_token}::DATE
  AND LOWER(p.title) = LOWER(${page_token})
  AND (
	   cl.area ISNULL
	OR LOWER(${area_type_token}) || '::' || LOWER(${area_name_token}) IN (
      	 SELECT LOWER(area_type) || '::' || LOWER(area_name)
      	 FROM covid19.area_reference AS ar
      	 WHERE ar.area_type IN (
      	     SELECT SPLIT_PART(n, '::', 1)
      	     FROM UNNEST(cl.area::TEXT[]) AS n
      	 ) AND ar.area_code ~* (
      	   SELECT string_agg(SPLIT_PART(n, '::', 2), '|')
      	   FROM UNNEST(cl.area::TEXT[]) AS n
      	 )
      )
    )
ORDER BY date DESC;
`

// var urlParamFilters = map[string]string{
// 	"page":      `LOWER(p.title) = LOWER(${token_id})`,
// 	"date":      `${token_id}`,
// 	"area_name": `{area_name_token}`,
// 	"area_type": `{area_type_token}`,
// }
