-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
  id        uuid  PRIMARY KEY,
  name      text  NOT NULL,
  email     text  NOT NULL UNIQUE,
  password  text  NOT NULL,
  salt      text  NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
