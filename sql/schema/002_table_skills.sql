-- +goose Up
CREATE TABLE skills (
    id SERIAL PRIMARY KEY,
    details JSONB NOT NULL,
    
    CONSTRAINT skill_schema CHECK (
        details ? 'name' AND
        details ? 'damage' AND
        details ? 'reach' AND
        details ? 'payment' AND
        details ? 'pay_with' AND
        
        length(details->>'name') BETWEEN 1 AND 50 AND
        (details->>'damage')::integer >= 0 AND
        (details->>'distance')::integer >= 0 AND

        jsonb_typeof(details->'payment') = 'array' AND
        jsonb_typeof(details->'pay_with') = 'array'
    )
);

-- +goose Down
DROP TABLE skills;
