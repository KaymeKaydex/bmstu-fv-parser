-- +goose Up
-- +goose StatementBegin
create table working_out_items
(
    id                    int not null
        constraint working_out_items_pkey
            primary key,
    branch_id             int,
    points_general        int,
    points_extra          int,
    title                 varchar(255),
    description           text,
    url                   varchar(255),
    date_begin            timestamp with time zone,
    date_end              timestamp with time zone,
    date_registration_end timestamp with time zone,
    count_user_max        int,
    count_users           int,
    address_id            int,
    address_title         text,
    address               text,
    is_allow_join         boolean,
    users_count_all       int,

    created_at            date
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE working_out_items CASCADE;
-- +goose StatementEnd
