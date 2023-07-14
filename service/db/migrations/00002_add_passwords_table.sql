-- +goose Up
-- +goose StatementBegin
CREATE TABLE passwords(
  id        uuid PRIMARY KEY, 
  user_id   uuid REFERENCES users(id),
  username  text NOT NULL,
  password  text NOT NULL,
  name      text NOT NULL, 
  website   text NOT NULL
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE passwords;
-- +goose StatementEnd