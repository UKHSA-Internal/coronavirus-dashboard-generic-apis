package areaByType

const query = `
SELECT area_name AS "areaName", area_code AS "areaCode"
FROM covid19.area_reference
WHERE area_type = $1
ORDER BY area_name ASC;
`
