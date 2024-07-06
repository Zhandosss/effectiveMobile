-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS works (
    id UUID NOT NULL DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    start_time TIMESTAMPTZ NOT NULL,
    end_time TIMESTAMPTZ,
    time_period_in_minute VARCHAR(255) NOT NULL DEFAULT '00h00m', --time adds only after work is done
    primary key (id),
    foreign key (user_id) references users(id)
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS works;
-- +goose StatementEnd