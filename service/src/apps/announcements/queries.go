package announcements

const latestDataQuery = `
WITH latest_release AS (
    SELECT MAX(rr.timestamp)::DATE
    FROM covid19.release_reference AS rr
    WHERE rr.released IS TRUE
)
SELECT id::TEXT,
       launch::DATE::TEXT,
       expire::DATE::TEXT,
       COALESCE(date, launch::DATE)::TEXT AS date,
       body
FROM covid19.announcement AS an
WHERE
    (
        (
                an.deploy_with_release IS TRUE
            AND an.launch::DATE <= (SELECT * FROM latest_release)
        )
      OR (
                an.deploy_with_release IS FALSE
            AND an.launch <= NOW()
        )
    )
  AND (
        (
                an.remove_with_release IS TRUE
            AND an.expire::DATE > (SELECT * FROM latest_release)
        )
      OR (
                an.remove_with_release IS FALSE
            AND an.expire > NOW()
        )
    )
ORDER BY an.launch DESC, an.expire DESC;
`

const allDataQuery = `
WITH latest_release AS (
    SELECT MAX(rr.timestamp)::DATE
    FROM covid19.release_reference AS rr
    WHERE rr.released IS TRUE
)
SELECT id::TEXT,
       launch::DATE::TEXT,
       expire::DATE::TEXT,
       (
         (
                an.remove_with_release IS TRUE
            AND an.expire::DATE < (SELECT * FROM latest_release)
         )
         OR (
                an.remove_with_release IS FALSE
            AND an.expire < NOW()
         )
       ) AS has_expired,
       COALESCE(date, launch::DATE)::TEXT AS date,
       body
FROM covid19.announcement AS an
WHERE
    (
        (
                an.deploy_with_release IS TRUE
            AND an.launch::DATE <= (SELECT * FROM latest_release)
        )
      OR (
                an.deploy_with_release IS FALSE
            AND an.launch <= NOW()
        )
    )
ORDER BY an.launch DESC, an.expire DESC;
`

const allDataQueryFeed = `
WITH
    latest_release AS (
		SELECT MAX(rr.timestamp)::DATE
		FROM covid19.release_reference AS rr
		WHERE rr.released IS TRUE
	),
    last_update_ts AS (
	       SELECT MAX(launch)
	       FROM covid19.announcement
	       WHERE (
                    deploy_with_release IS TRUE
                AND launch::DATE <= (SELECT * FROM latest_release)
            )
	        OR (
                deploy_with_release IS FALSE
            AND launch <= NOW()
        )
    )
SELECT id::TEXT AS guid,
       launch AS "pubDate",
	   COALESCE(date, launch::DATE) AS date,
       body AS description,
	   (SELECT * FROM last_update_ts) AS last_update
FROM covid19.announcement AS an
WHERE
    (
        (
                an.deploy_with_release IS TRUE
            AND an.launch::DATE <= (SELECT * FROM latest_release)
        )
      OR (
                an.deploy_with_release IS FALSE
            AND an.launch <= NOW()
        )
    )
ORDER BY an.launch DESC, an.expire DESC;
`

const itemQuery = `
WITH
    latest_release AS (
		SELECT MAX(rr.timestamp)::DATE
		FROM covid19.release_reference AS rr
		WHERE rr.released IS TRUE
	)
SELECT id::TEXT AS guid,
       launch::DATE::TEXT,
       expire::DATE::TEXT,
	   COALESCE(date, launch::DATE) AS date,
       (
         (
                an.remove_with_release IS TRUE
            AND an.expire::DATE < (SELECT * FROM latest_release)
         )
         OR (
                an.remove_with_release IS FALSE
            AND an.expire < NOW()
         )
       ) AS has_expired,
       body
FROM covid19.announcement AS an
WHERE id::TEXT = $1
ORDER BY an.launch DESC, an.expire DESC;
`
