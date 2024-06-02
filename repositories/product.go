package repositories

import (
	"dbo/dbo-is-backend/models"
	"dbo/dbo-is-backend/pkg/utils"
	"fmt"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

type ProductRepositoryInterface interface {
	CreateProduct(product *models.Product) error
	GetProducts(pagination utils.Pagination, where map[string]string) ([]models.Product, int64, error)
	GetProductById(id string) (models.Product, error)
	UpdateProduct(product *models.ProductUpdate) error
	DeleteProduct(product *models.ProductUpdate) (err error)
}

func NewProductRepository(db *gorm.DB) ProductRepositoryInterface {
	return &productRepository{
		db: db,
	}
}

func (pr *productRepository) CreateProduct(product *models.Product) error {
	return pr.db.Create(product).Error
}

func (pr *productRepository) GetProducts(pagination utils.Pagination, where map[string]string) ([]models.Product, int64, error) {
	var count int64
	var err error
	var sortField, sortDirection string
	var products []models.Product

	queryBuilder := pr.db.Model(&models.Product{}).Where("status <> ?", models.StatusDeleted)

	if id, ok := where["id"]; ok && id != "" {
		queryBuilder = queryBuilder.Where("id = ?", id)
	}

	if name, ok := where["name"]; ok && name != "" {
		name := fmt.Sprintf("%%%s%%", name)
		queryBuilder = queryBuilder.Where(`"name" ILIKE ?`, name)
	}

	if pagination.SortField != "" {
		if pagination.SortField == "name" {
			sortField = `INITCAP("name")`
		} else if pagination.SortField == "price" {
			sortField = "price"
		} else if pagination.SortField == "stock" {
			sortField = "stock"
		} else {
			sortField = "created_at"
		}
	} else {
		sortField = "created_at"
	}

	if pagination.SortDirection != "" {
		sortDirection = pagination.SortDirection
	} else {
		sortDirection = models.SortDirectionDESC.String()
	}

	err = queryBuilder.Count(&count).Error
	if err != nil {
		return nil, count, err
	}

	offset := (pagination.Page - 1) * pagination.Limit
	orderBy := fmt.Sprintf("%s %s", sortField, sortDirection)
	limitBuilder := queryBuilder.Limit(pagination.Limit).Offset(offset).Order(orderBy)

	result := limitBuilder.Find(&products)
	if result.Error != nil {
		return nil, count, result.Error
	}

	return products, count, nil
}

func (pr *productRepository) GetProductById(id string) (models.Product, error) {
	var product models.Product
	if err := pr.db.Where(&models.Product{ID: id}).First(&product).Error; err != nil {
		return product, err
	}
	return product, nil
}

func (pr *productRepository) UpdateProduct(product *models.ProductUpdate) error {
	return pr.db.
		Model(&models.Product{ID: product.ID}).
		Updates(
			map[string]interface{}{
				"name":       product.Name,
				"price":      product.Price,
				"stock":      product.Stock,
				"status":     product.Status,
				"updated_at": gorm.Expr("now()"),
				"updated_by": product.UpdatedBy,
			},
		).Error
}

func (pr *productRepository) DeleteProduct(product *models.ProductUpdate) (err error) {
	return pr.db.
		Model(&models.Product{ID: product.ID}).
		Updates(
			map[string]interface{}{
				"status":     models.StatusDeleted.String(),
				"updated_at": gorm.Expr("now()"),
				"updated_by": product.UpdatedBy,
			},
		).Error
}
