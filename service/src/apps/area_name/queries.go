package area_name

const areaQuery = `
SELECT
   area_code    AS "areaCode",
   area_name    AS "areaName",
   ar.area_type AS "areaType"
FROM covid19.area_reference AS ar
WHERE area_type = LOWER($1)
  AND area_name ILIKE LOWER($2)
LIMIT 1 OFFSET 0;
`
