-- +migrate Up
ALTER TABLE notes ADD INDEX vaccine_type_number_of_vaccination_index(vaccine_type, number_of_vaccination);

-- +migrate Down
ALTER TABLE notes DROP INDEX vaccine_type_number_of_vaccination_index;