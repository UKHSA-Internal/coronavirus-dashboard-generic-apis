package utils

const timestampQuery = `
SELECT DATE(MAX(timestamp))::TEXT AS date
FROM covid19.release_reference AS rr
	JOIN covid19.release_category AS rc ON rc.release_id = rr.id
WHERE released IS TRUE
  AND process_name = $1
`
