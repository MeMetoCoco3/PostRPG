-- +goose Up
INSERT INTO map (x, y, terrain, character)
SELECT
    x_series.x,
    y_series.y,
    'GRASS' AS terrain,
    NULL AS character
FROM
    generate_series(0, 10) AS x_series(x),
    generate_series(0, 10) AS y_series(y);

-- +goose Down
DELETE FROM map;
