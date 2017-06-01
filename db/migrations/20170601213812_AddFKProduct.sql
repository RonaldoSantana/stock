
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE product ADD CONSTRAINT fk_product_length_class FOREIGN KEY (length_class_id) REFERENCES length_class(id);
ALTER TABLE product ADD CONSTRAINT fk_product_weight_class FOREIGN KEY (weight_class_id) REFERENCES weight_class(id);
ALTER TABLE product ADD CONSTRAINT fk_product_supplier FOREIGN KEY (supplier_id) REFERENCES supplier(id);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE product DROP FOREIGN KEY fk_product_length_class;
ALTER TABLE product DROP FOREIGN KEY fk_product_weight_class;
ALTER TABLE product DROP FOREIGN KEY fk_product_supplier;
