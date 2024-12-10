-- +goose Up

		`CREATE TYPE weapon AS (
			damage INTEGER,
			reach INTEGER,
			name TEXT,
			description TEXT
		)`,-- +goose Down
DROP TABLE weapons;
