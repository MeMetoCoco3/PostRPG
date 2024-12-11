-- +goose Up
CREATE TABLE terrains (
    id SERIAL PRIMARY KEY,
    matrix INTEGER[][] 
    CONSTRAINT terrain_schema CHECK(
      matrix IS NOT NULL AND
      array_length(matrix, 1) > 8 AND
      array_length(matrix, 2) > 8
    )
);

-- +goose Down
DROP TABLE terrains;
