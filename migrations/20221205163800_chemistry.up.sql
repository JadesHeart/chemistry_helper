CREATE TABLE lowercase_characteristics
(
    id                                      bigserial not null primary key,
    el_name                                 varchar   not null unique,
    electronic_configuration                varchar   not null,
    stability_oxidation_state_configuration varchar   not null,
    melting_point                           varchar   not null,
    boiling_point                           varchar   not null,
    chemical_compounds                      varchar not null
)