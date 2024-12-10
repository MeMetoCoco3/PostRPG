-- +goose Up
CREATE TABLE characters (
    id SERIAL PRIMARY KEY,
    details JSONB NOT NULL,
    
    CONSTRAINT character_schema 
    CHECK (
        jsonb_matches_schema(
            '{
                "type": "object",
                "required": ["health", "mana", "stamina", "strength", "role", "name", "icon"],
                "properties": {
                    "name": {"type": "string", "minLength": 1, "maxLength": 50},
                    "health": {"type": "integer", "minimum": 0},
                    "mana":{"type": "integer", "minimum": 0},
                    "stamina":{"type": "integer", "minimum": 0},
                    "strength": {"type": "integer", "minimum": 0},
                    "role": {"type": "integer", "minimum": 0, "maximum": 2},
                    "icon": {"type": "string",  "minLength": 1, "maxLength": 1}
                }
            }'::jsonb, 
            details
        )
    )
);

-- +goose Down
DROP TABLE characters;
