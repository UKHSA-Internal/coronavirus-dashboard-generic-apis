CREATE TABLE IF NOT EXISTS covid19.time_series_p2021_3_5_utla
    PARTITION OF covid19.time_series ( partition_id )
        FOR VALUES IN ('2021_3_5|utla');

CREATE TABLE IF NOT EXISTS covid19.time_series_p2021_3_5_msoa
    PARTITION OF covid19.time_series ( partition_id )
        FOR VALUES IN ('2021_3_5|msoa');

CREATE TABLE IF NOT EXISTS covid19.time_series_p2021_3_5_nhstrust
    PARTITION OF covid19.time_series ( partition_id )
        FOR VALUES IN ('2021_3_5|nhstrust');

CREATE TABLE IF NOT EXISTS covid19.time_series_p2021_3_5_other
    PARTITION OF covid19.time_series ( partition_id )
        FOR VALUES IN ('2021_3_5|other');

CREATE TABLE IF NOT EXISTS covid19.time_series_p2021_3_5_ltla
    PARTITION OF covid19.time_series ( partition_id )
        FOR VALUES IN ('2021_3_5|ltla');
