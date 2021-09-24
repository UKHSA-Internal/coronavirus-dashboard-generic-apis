CREATE SCHEMA covid19;
--=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~

create table if not exists covid19.page
(
    id            uuid                  not null,
    title         varchar(120)          not null,
    uri           varchar(150)          not null,
    data_category boolean default false not null,
    constraint page_pkey
        primary key (id),
    constraint page_title_key
        unique (title)
);

create index if not exists page_lower_idx
    on covid19.page (lower(title::text));


CREATE TABLE IF NOT EXISTS covid19.area_reference
(
    id                SERIAL                  NOT NULL UNIQUE,
    area_type         VARCHAR(15)             NOT NULL,
    area_code         VARCHAR(12)             NOT NULL,
    area_name         VARCHAR(120)            NOT NULL,
    unique_ref varchar(26) default "substring"(((now())::character varying)::text, 0, 26) not null,

    PRIMARY KEY (area_type, area_code),
    UNIQUE (area_type, area_code),
    constraint unq_area_reference_ref
        unique (unique_ref)
);

CREATE UNIQUE INDEX IF NOT EXISTS arearef_type_code_idx
    ON covid19.area_reference
    USING BTREE (area_type, area_code);

CREATE UNIQUE INDEX IF NOT EXISTS arearef_id_idx
    ON covid19.area_reference
    USING BTREE (id);

CREATE INDEX IF NOT EXISTS arearef_namelower_idx
    ON covid19.area_reference
    USING BTREE (LOWER(area_name));

CREATE INDEX IF NOT EXISTS arearef_area_code_initial
    ON covid19.area_reference
    USING BTREE (SUBSTRING(area_code, 1));

--=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~

create table if not exists covid19.metric_reference
(
    id            serial                not null,
    metric        varchar(120)          not null,
    released      boolean default false not null,
    metric_name   varchar(150),
    source_metric boolean default false not null,
    category      uuid,
    deprecated    date,
    constraint metric_reference_pkey
        primary key (id),
    constraint metric_reference_metric_key
        unique (metric),
    constraint fk_metric_category
        foreign key (category) references covid19.page
            on delete cascade
            deferrable initially deferred
);

create index if not exists metricref_metrics_idx
    on covid19.metric_reference (metric);

create index if not exists metricref_releasedmetrics_idx
    on covid19.metric_reference (metric, released)
    where (released = true);

create index if not exists metricref_sourcemetric_idx
    on covid19.metric_reference (source_metric);

create index if not exists metricref_metric_id_idx
    on covid19.metric_reference (metric, id);

create index if not exists metricref_id_metric_idx
    on covid19.metric_reference (id, metric);

--=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~

CREATE TABLE IF NOT EXISTS covid19.release_reference
(
    id         SERIAL          NOT NULL UNIQUE PRIMARY KEY,
    timestamp  TIMESTAMP       NOT NULL UNIQUE,
    released   BOOLEAN         NOT NULL DEFAULT FALSE
);

CREATE INDEX IF NOT EXISTS releaseref_timestamp_idx
    ON covid19.release_reference
    USING BTREE (timestamp);

CREATE INDEX IF NOT EXISTS releaseref_releasedtimestamp_idx
    ON covid19.release_reference
    USING BTREE (timestamp, released) WHERE released = TRUE;

--=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~

CREATE TABLE IF NOT EXISTS covid19.area_priorities
(
    area_type         VARCHAR(15)             NOT NULL UNIQUE,
    priority          NUMERIC                 NOT NULL,

    PRIMARY KEY (area_type, priority)
    );

--=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~

CREATE TABLE IF NOT EXISTS covid19.time_series
(
    hash               VARCHAR(24)  NOT NULL,
    partition_id       VARCHAR(26)  NOT NULL,
    release_id         INT          NOT NULL,
    area_id            INT          NOT NULL,
    metric_id          INT          NOT NULL,
    date               DATE         NOT NULL,
    payload            JSONB        DEFAULT '{"value": null}',

    CONSTRAINT unique_partition UNIQUE (hash, partition_id),
    PRIMARY KEY (hash, area_id, metric_id, release_id, partition_id)

    )
    PARTITION BY LIST ( partition_id );

CREATE INDEX IF NOT EXISTS timeseries_hash_idx
    ON covid19.time_series USING BTREE (hash);

CREATE INDEX IF NOT EXISTS timeseries_releaseid_idx
    ON covid19.time_series USING BTREE (release_id);


CREATE INDEX IF NOT EXISTS timeseries_payload_idx
    ON covid19.time_series USING GIN (payload jsonb_path_ops);

CREATE INDEX IF NOT EXISTS timeseries_payload_notnull_idx
    ON covid19.time_series USING GIN (payload)
    WHERE (payload -> 'value') NOTNULL;

CREATE INDEX IF NOT EXISTS timeseries_timestamp_idx
    ON covid19.time_series USING BTREE (partition_id);

CREATE INDEX IF NOT EXISTS timeseries_area_selfjoin_idx
    ON covid19.time_series
    USING BTREE (partition_id, area_id, date);

CREATE INDEX IF NOT EXISTS timeseries_response_order_idx
    ON covid19.time_series
    USING BTREE (area_id DESC, date DESC);

CREATE INDEX IF NOT EXISTS arearef_area_code_initial
    ON covid19.area_reference
    USING BTREE (SUBSTRING(area_code, 1));

CREATE INDEX IF NOT EXISTS timeseries_recorddate_idx
    ON covid19.time_series
    USING BTREE (partition_id);

CREATE INDEX IF NOT EXISTS timeseries_metric_idx
    ON covid19.time_series
    USING BTREE (metric_id);

CREATE INDEX IF NOT EXISTS timeseries_area_idx
    ON covid19.time_series
    USING BTREE (area_id);

--=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~=~

CREATE TYPE RELEASE_PROCESSES AS ENUM (
    'MAIN',
    'MSOA',
    'VACCINATION',
    'AGE DEMOGRAPHICS: CASE - EVENT DATE',
    'AGE-DEMOGRAPHICS: DEATH28DAYS - EVENT DATE',
    'AGE-DEMOGRAPHICS: VACCINATION - EVENT DATE',
    'MSOA: VACCINATION - EVENT DATE',
    'POSITIVITY & PEOPLE TESTED'
);

CREATE TABLE IF NOT EXISTS covid19.release_category
(
    release_id       INT                   NOT NULL,
    process_name     RELEASE_PROCESSES     NOT NULL,

    PRIMARY KEY (release_id, process_name),

    FOREIGN KEY ( release_id )
    REFERENCES covid19.release_reference ( id )
    ON UPDATE CASCADE
    ON DELETE CASCADE
    );

CREATE INDEX IF NOT EXISTS release_cat_release_idx
    ON covid19.release_category
    USING BTREE ( release_id );

CREATE INDEX IF NOT EXISTS release_cat_release_proc_idx
    ON covid19.release_category
    USING BTREE ( release_id,  process_name );






create table if not exists covid19.tag
(
    id          uuid        not null,
    association varchar(30) not null,
    tag         varchar(40) not null,
    constraint tag_pkey
        primary key (id)
);

create table if not exists covid19.metric_tag
(
    id        uuid         not null,
    metric_id varchar(120) not null,
    tag_id    uuid         not null,
    constraint metric_tag_pkey
        primary key (tag_id, metric_id, id),
    constraint metric_tag_tag_id_id_dabd969c_uniq
        unique (tag_id, id),
    constraint metric_tag_metric_id_8702beb6_fk_metric_reference_metric
        foreign key (metric_id) references covid19.metric_reference (metric)
            deferrable initially deferred,
    constraint metric_tag_tag_id_b55c2437_fk_tag_id
        foreign key (tag_id) references covid19.tag
            deferrable initially deferred
);

create index if not exists metric_tag_metric_id_8702beb6
    on covid19.metric_tag (metric_id);

create index if not exists metric_tag_tag_id_b55c2437
    on covid19.metric_tag (tag_id);


create table if not exists covid19.metric_asset
(
    id            uuid                                               not null,
    label         varchar(255)                                       not null,
    body          text                                               not null,
    last_modified timestamp with time zone default CURRENT_TIMESTAMP not null,
    released      boolean                  default false             not null,
    constraint metric_asset_se_pkey
        primary key (id)
);

create index if not exists metric_asset_se_released_idx
    on covid19.metric_asset (released);

create table if not exists covid19.metric_asset_to_metric
(
    id         uuid         not null,
    metric_id  varchar(120) not null,
    asset_id   uuid         not null,
    asset_type varchar(50)  not null,
    "order"    integer,
    constraint metric_asset_to_metric_pkey
        primary key (id),
    constraint metric_asset_to_metric_metric_id_fkey
        foreign key (metric_id) references covid19.metric_reference (metric)
            on delete cascade
            deferrable initially deferred,
    constraint metric_asset_to_metric_asset_id_fkey
        foreign key (asset_id) references covid19.metric_asset
            on delete cascade
            deferrable initially deferred,
    constraint metric_asset_to_metric_order_check
        check (("order" IS NULL) OR ("order" > 0))
);

create unique index if not exists unique_metric_asset
    on covid19.metric_asset_to_metric (metric_id, asset_id);

create unique index if not exists unique_metric_asset_order
    on covid19.metric_asset_to_metric (metric_id, asset_type, "order");

create index if not exists idx_metric_asset_order
    on covid19.metric_asset_to_metric (metric_id, asset_id, "order");

create table if not exists covid19.change_log
(
    id                uuid                                   not null,
    date              date                                   not null,
    expiry            date,
    heading           varchar(150)                           not null,
    body              text                                   not null,
    details           text,
    type_id           uuid                                   not null,
    high_priority     boolean                  default false not null,
    display_banner    boolean                  default false not null,
    area              varchar(50)[],
    timestamp_created timestamp with time zone default now() not null,
    constraint change_log_pkey
        primary key (id),
    constraint unique_change_log
        unique (date, heading, type_id),
    constraint change_log_type_id_fkey
        foreign key (type_id) references covid19.tag
            on delete cascade
            deferrable initially deferred
);

create index if not exists idx_changelog_date_order
    on covid19.change_log (date desc);

create index if not exists idx_changelog_type_id
    on covid19.change_log (type_id);

create index if not exists idx_changelog_heading
    on covid19.change_log using gin (to_tsvector('english'::regconfig, heading::text));

create index if not exists idx_changelog_body
    on covid19.change_log using gin (to_tsvector('english'::regconfig, body));

create table if not exists covid19.change_log_to_metric
(
    id        uuid         not null,
    log_id    uuid         not null,
    metric_id varchar(120) not null,
    constraint change_log_to_metric_pkey
        primary key (id),
    constraint change_log_to_metric_log_id_fkey
        foreign key (log_id) references covid19.change_log
            on delete cascade
            deferrable initially deferred,
    constraint change_log_to_metric_metric_id_fkey
        foreign key (metric_id) references covid19.metric_reference (metric)
            on delete cascade
            deferrable initially deferred
);

create unique index if not exists idx_changelog2metric
    on covid19.change_log_to_metric (log_id, metric_id);

create table if not exists covid19.change_log_to_page
(
    id      uuid not null,
    log_id  uuid not null,
    page_id uuid not null,
    constraint change_log_to_page_pkey
        primary key (id),
    constraint change_log_to_page_log_id_fkey
        foreign key (log_id) references covid19.change_log
            on delete cascade
            deferrable initially deferred,
    constraint change_log_to_page_page_id_fkey
        foreign key (page_id) references covid19.page
            on delete cascade
            deferrable initially deferred
);

create unique index if not exists idx_changelog2page
    on covid19.change_log_to_page (log_id, page_id);

create table if not exists covid19.announcement
(
    id                  uuid                     not null,
    launch              timestamp with time zone not null,
    expire              timestamp with time zone not null,
    date                date,
    deploy_with_release boolean default true     not null,
    remove_with_release boolean default true     not null,
    body                varchar(400)             not null,
    constraint announcement_pkey
        primary key (id),
    constraint chk__anc_exp_gt_launch
        check (launch < expire),
    constraint chk__anc_date_bw_launch_exp
        check (((launch)::date <= date) AND (date <= (expire)::date))
);

create index if not exists idx__anc_launch
    on covid19.announcement (launch);

create index if not exists idx__anc_expire
    on covid19.announcement (expire);

create index if not exists idx__anc_launch_expire
    on covid19.announcement (launch desc, expire desc);



create table if not exists covid19.area_relation
(
    child_id  integer not null,
    parent_id integer not null,
    constraint area_relation_pkey
        primary key (child_id, parent_id),
    constraint fk_child_area_rel_id
        foreign key (child_id) references covid19.area_reference (id)
            on update cascade on delete cascade,
    constraint fk_parent_area_rel_id
        foreign key (parent_id) references covid19.area_reference (id)
            on update cascade on delete cascade
);

create index if not exists msoarels_parent_child_idx
    on covid19.area_relation (child_id, parent_id);



create table if not exists covid19.page_area_type_reference
(
    page      varchar(120) not null,
    area_type varchar(15)  not null,
    constraint page_areatype_pkey
        primary key (page, area_type),
    constraint page_areatype_page_fk
        foreign key (page) references covid19.page (title)
            on update cascade on delete cascade
);

create index if not exists pageareatype_ref_page_aretype_idx
    on covid19.page_area_type_reference (page, area_type);

create index if not exists pageareatype_ref_pagelower_aretype_idx
    on covid19.page_area_type_reference (lower(page::text), area_type);


CREATE TABLE IF NOT EXISTS covid19.postcode_lookup
(
    postcode            VARCHAR(10)     NOT NULL,
    area_id             INT             NOT NULL,

    PRIMARY KEY (postcode, area_id),

    CONSTRAINT fk_area
        FOREIGN KEY (area_id)
            REFERENCES covid19.area_reference ( id )
            ON DELETE CASCADE

);

CREATE INDEX IF NOT EXISTS area_lookup_trimmedpostcode_idx
    ON covid19.postcode_lookup USING BTREE (UPPER(REPLACE(postcode, ' ', '')));

