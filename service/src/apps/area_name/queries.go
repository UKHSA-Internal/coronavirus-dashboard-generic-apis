package area_name

const areaQuery = `
SELECT
	area_code    AS "areaCode",
	area_name    AS "areaName",
	ar.area_type AS "areaType"
FROM covid19.area_reference AS ar
WHERE id IN (
	  	 SELECT parent_id
	  	 FROM covid19.area_reference AS ar2
		     JOIN covid19.area_relation AS pl2 ON pl2.child_id = ar2.id
	  	 WHERE area_type = $1
	  	   AND area_name = $2
	  )
   OR ( area_type = $1 AND area_name = $2 )
`
