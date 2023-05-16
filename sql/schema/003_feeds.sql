-- +goose Up

CREATE TABLE feeds (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name TEXT NOT NULL,
  url TEXT UNIQUE NOT NULL,
  -- user id in this table that refers to an id from users table
  -- biz logic - if we try to create an rss feed with an id that doesn't exist in the users table it should give an error
  --            - when a user is deleted in the users table the users relevant feeds should also get deleted
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);



-- +goose Down
DROP TABLE feeds;