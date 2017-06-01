
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
INSERT INTO `geo_zone_matrix` (`country_id`, region_id, `geo_zone_id`) VALUES
(153, 2344, 1),
(153, 2345, 1),
(153, 2346, 1),
(153, 2347, 1),
(153, 2348, 1),
(153, 2349, 1),
(153, 2350, 1),
(153, 2351, 1),
(153, 2352, 1),
(153, 2353, 1),
(153, 2354, 1),
(153, 2355, 1),
(153, 2356, 1),
(153, 2357, 1),
(153, 2358, 1),
(153, 2359, 1),
(153, 2360, 1),
(153, 2361, 1),
(153, 2362, 1);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DELETE FROM geo_zone_matrix;
