-- +goose Up
-- +goose StatementBegin
CREATE TABLE personal_notes(
    project_id integer,
    notes varchar(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE personal_notes;
-- +goose StatementEnd
