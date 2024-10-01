-- +goose Up
-- +goose StatementBegin
CREATE TABLE cookies(
    name text,
    value text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE cookies;
-- +goose StatementEnd
