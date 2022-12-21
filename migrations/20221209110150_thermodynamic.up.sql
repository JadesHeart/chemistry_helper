CREATE TABLE thermodynamic_characteristics
(
    id           bigserial not null primary key,
    el_name      varchar not null unique,
    first_param  varchar not null,
    second_param varchar not null,
    third_param  varchar not null,
    fourth_param varchar
)