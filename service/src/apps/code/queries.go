package code

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
	  	   AND area_code = $2
	  )
   OR ( area_type = $1 AND area_code = $2 )
`

const postcodeQuery = `
SELECT postcode, 
	   area_code AS "areaCode", 
	   area_name AS "areaName", 
	   ar.area_type AS "areaType"
FROM covid19.area_reference AS ar
  JOIN covid19.postcode_lookup AS pl ON pl.area_id = ar.id
  JOIN covid19.area_priorities AS ap ON ap.area_type = ar.area_type
WHERE UPPER(REPLACE(postcode, ' ', '')) = $2
  AND priority >= (
	SELECT priority 
	FROM covid19.area_priorities
	WHERE area_type = $1
	LIMIT 1
  )
`
