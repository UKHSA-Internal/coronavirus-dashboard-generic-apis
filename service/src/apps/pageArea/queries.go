package pageArea

const query = `
SELECT area_type AS "areaType", area_name AS "areaName", area_code AS "areaCode"
FROM covid19.page AS pg
	JOIN covid19.page_area_reference AS par ON par.category_id = pg.id
	JOIN covid19.area_reference AS ar ON ar.id = par.area_id
WHERE LOWER(pg.title) = $1`

const areaTypeFilter = ` AND ar.area_type = $2`
