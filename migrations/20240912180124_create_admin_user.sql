-- +goose Up
-- +goose StatementBegin
INSERT INTO users (email, hashed_password, is_admin, created_at, updated_at, is_premium) VALUES ('nikos.vare@gmail.com', 'password', TRUE, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, TRUE);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM user WHERE email = 'nikos.vare@gmail.com';
-- +goose StatementEnd
