package metric_areas

const mainQuery = `
SELECT area_type, last_update
FROM (
         SELECT area_type, MAX(date)::DATE::TEXT AS last_update
         FROM covid19.time_series_p${partition_date}_other AS ts
                  JOIN covid19.area_reference AS ar ON ar.id = ts.area_id
         WHERE date > (NOW() - INTERVAL '30 days')
           AND metric_id IN (SELECT id
                             FROM covid19.metric_reference AS mr
                             WHERE metric ILIKE ${metric_token}
                               AND released IS TRUE)
         GROUP BY area_type
         UNION
         (
             SELECT area_type, MAX(date)::DATE::TEXT AS last_update
             FROM covid19.time_series_p${partition_date}_utla AS ts
                      JOIN covid19.area_reference AS ar ON ar.id = ts.area_id
             WHERE date > (NOW() - INTERVAL '30 days')
               AND metric_id IN (SELECT id
                                 FROM covid19.metric_reference AS mr
                                 WHERE metric ILIKE ${metric_token}
                                   AND released IS TRUE)
             GROUP BY area_type
         )
         UNION
         (
             SELECT area_type, MAX(date)::DATE::TEXT AS last_update
             FROM covid19.time_series_p${partition_date}_ltla AS ts
                      JOIN covid19.area_reference AS ar ON ar.id = ts.area_id
             WHERE date > (NOW() - INTERVAL '30 days')
               AND metric_id IN (SELECT id
                                 FROM covid19.metric_reference AS mr
                                 WHERE metric ILIKE ${metric_token}
                                   AND released IS TRUE)
             GROUP BY area_type
         )
         UNION
         (
             SELECT area_type, MAX(date)::DATE::TEXT AS last_update
             FROM covid19.time_series_p${partition_date}_nhstrust AS ts
                      JOIN covid19.area_reference AS ar ON ar.id = ts.area_id
             WHERE date > (NOW() - INTERVAL '30 days')
               AND metric_id IN (SELECT id
                                 FROM covid19.metric_reference AS mr
                                 WHERE metric ILIKE ${metric_token}
                                   AND released IS TRUE)
             GROUP BY area_type
         )
         UNION
         (
             SELECT area_type, MAX(date)::DATE::TEXT AS last_update
             FROM covid19.time_series_p${partition_date}_msoa AS ts
                      JOIN covid19.area_reference AS ar ON ar.id = ts.area_id
             WHERE date > (NOW() - INTERVAL '30 days')
               AND metric_id IN (SELECT id
                                 FROM covid19.metric_reference AS mr
                                 WHERE metric ILIKE ${metric_token}
                                   AND released IS TRUE)
             GROUP BY area_type
         )
     ) AS metrics
ORDER BY area_type;
`
