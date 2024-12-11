-- +goose Up
CREATE TABLE characters (
    id SERIAL PRIMARY KEY,
    details JSONB NOT NULL,
    
    CONSTRAINT valid_character CHECK (
        details ? 'health' AND
        details ? 'mana' AND
        details ? 'stamina' AND
        details ? 'strength' AND
        details ? 'role' AND
        details ? 'name' AND
        details ? 'icon' AND
        
        (details->>'health')::integer >= 0 AND
        (details->>'health')::integer <= 1000 AND
        
        (details->>'mana')::integer >= 0 AND
        (details->>'mana')::integer <= 1000 AND
        
        (details->>'stamina')::integer >= 0 AND
        (details->>'stamina')::integer <= 1000 AND
        
        (details->>'strength')::integer >= 0 AND
        (details->>'strength')::integer <= 100 AND
        
        (details->>'role')::integer BETWEEN 0 AND 2 AND
        
        length(details->>'name') BETWEEN 1 AND 50 AND
        
        length(details->>'icon') = 1
    )
);

-- +goose Down
DROP TABLE characters;
