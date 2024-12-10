-- +goose Up
CREATE TABLE skills (
    id SERIAL PRIMARY KEY,
    details JSONB NOT NULL,
    
    CONSTRAINT skill_schema 
    CHECK (
        jsonb_matches_schema(
            '{
                "type": "object",
                "required": ["name", "damage", "distance", "payment", "pay_with"],
                "properties": {
                    "name": {"type": "string", "minLength": 1, "maxLength": 50},
                    "damage": {"type": "integer", "minimum": 0},
                    "distance": {"type": "integer", "minimum": 0},
                    "payment": {"type": "array"},
                    "pay_with": {"type": "array"}
                }
            }'::jsonb, 
            details
        )
    )
);

-- +goose Down
DROP TABLE skills;
