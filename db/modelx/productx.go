package modelx

import (
	"database/sql"

	"github.com/rmsj/stock/db/models"
	. "github.com/vattle/sqlboiler/queries/qm"
)

// ProductJSON is a product with all it's relationship
type ProductJSON struct {
	*models.Product
	TaxClass    *models.TaxClass
	WeightClass *models.WeightClass
	LengthClass *models.LengthClass
}

// ProductJSONBySKU finds product by SKU
func ProductJSONBySKU(db *sql.DB, sku string) (product *ProductJSON, err error) {

	dbProduct, err := models.Products(db, Where("sku = ?", sku)).One()
	if err != nil {
		return
	}
	taxClass, err := dbProduct.TaxClass(db).One()
	if err != nil {
		return
	}
	weightClass, err := dbProduct.WeightClass(db).One()
	if err != nil {
		return
	}
	lengthClass, err := dbProduct.LengthClass(db).One()
	if err != nil {
		return
	}

	product = &ProductJSON{
		dbProduct,
		taxClass,
		weightClass,
		lengthClass,
	}

	return
}

// ProductBySKU finds product by SKU
func ProductBySKU(db *sql.DB, sku string) (product *models.Product, err error) {

	product, err = models.Products(db, Where("sku = ?", sku)).One()
	if err != nil {
		return
	}

	return
}

// ProductSetStock removes product from stock - because of order, etc.
func ProductStock(db *sql.DB, sku string) (product *models.Product, err error) {

	product, err = models.Products(db, Where("sku = ?", sku)).One()
	if err != nil {
		return
	}

	return
}

// ProductExists checks if SKU for product is in use already
func ProductExists(db *sql.DB, sku string) (exists bool, err error) {
	exists, err = models.Users(db, Where("sku = ?", sku)).Exists()
	return
}
