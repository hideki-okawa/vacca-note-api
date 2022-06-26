-- +migrate Up
ALTER TABLE notes ADD second_vaccine_type VARCHAR(1) NOT NULL;

-- +migrate Down
ALTER TABLE notes DROP COLUMN second_vaccine_type;