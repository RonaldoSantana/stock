
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE product ADD CONSTRAINT fk_product_tax_class FOREIGN KEY (tax_class_id) REFERENCES tax_class(id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE product DROP FOREIGN KEY fk_product_tax_class;
