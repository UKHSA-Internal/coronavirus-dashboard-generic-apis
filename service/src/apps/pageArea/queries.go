package pageArea

const query = `
	SELECT MIN(ap.priority), MIN(ar.area_type) AS area_type, area_name, area_code
	FROM covid19.page AS pg
		JOIN covid19.page_area_reference AS par ON par.category_id = pg.id
		JOIN covid19.area_reference 	 AS ar  ON ar.id = par.area_id
		JOIN covid19.area_priorities 	 AS ap  ON ap.area_type = ar.area_type
	WHERE LOWER(pg.title) = $1
	  AND (
		   LOWER(pg.title) != 'vaccinations'
	    OR ar.area_type = 'nation'
	    OR area_code ~ '^[^WN]'
      )`

const areaTypeFilter = ` AND ar.area_type = %s`

const queryWrapper = `
SELECT area_type AS "areaType", area_name AS "areaName", area_code AS "areaCode"
FROM (%s) AS d
ORDER BY area_name`

const queryExtras = `
	GROUP BY area_code, area_name
`
