package product

import "context"

type service interface {
	ListProducts(ctx context.Context) error
}

type svc struct {
	//

}

func NewService() service {
	return &svc{
		// Service constructor
	}
}

func (s *svc) ListProducts(ctx context.Context) error {
	// 1. get the products from the database
	// 2. return the products
	return nil
}
