package utils

const processTimestampQuery = `
SELECT MAX(timestamp)::DATE::TEXT AS date
FROM covid19.release_reference AS rr
	JOIN covid19.release_category AS rc ON rc.release_id = rr.id
WHERE released IS TRUE
  AND process_name = $1
`

const timestampQuery = `
SELECT MAX(timestamp)::DATE::TEXT AS date
FROM covid19.release_reference AS rr
WHERE released IS TRUE
`
