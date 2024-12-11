-- +goose Up

CREATE TABLE weapons(
	id SERIAL PRIMARY KEY,
	details JSONB NOT NULL,

	CONSTRAINT weapon_schema CHECK(
		details ? 'name' AND
		details ? 'damage' AND
		details ? 'reach' AND
		details ? 'description' AND

		length(details->>'name') BETWEEN 1 AND 50 AND
		length(details->>'description') BETWEEN 1 AND 100 AND 
		(details->>'damage')::integer >= 0 AND
		(details->>'reach')::integer>= 0
	)
);

-- +goose Down
DROP TABLE weapons;
