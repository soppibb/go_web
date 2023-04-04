package products

import "github.com/soppibb/go_web/día_2/paquetes/internal/domain"

type Service interface {
	Create(product *domain.Product) error
	GetAll() ([]domain.Product, error)
}
