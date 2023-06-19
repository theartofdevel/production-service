package product

import (
	pb_prod_products "github.com/theartofdevel/production-service-contracts/gen/go/prod_service/products/v1"
	"production_service/internal/domain/policy/product"
	"production_service/internal/domain/product/model"
	"production_service/pkg/utils/pointer"
)

func NewCreateProductInput(req *pb_prod_products.CreateProductRequest) (product.CreateProductInput, error) {
	// validate here
	return product.NewCreateProductInput(
		req.Name,
		req.Description,
		req.ImageId,
		req.Price,
		req.CurrencyId,
		req.Rating,
		req.CategoryId,
	)
}

func NewCreateProductResponse(out product.CreateProductOutput) *pb_prod_products.CreateProductResponse {
	return &pb_prod_products.CreateProductResponse{
		Product: NewProductPB(out.Product),
	}
}

func NewProductPB(ent model.Product) *pb_prod_products.Product {
	pbProduct := &pb_prod_products.Product{
		Id:          ent.ID,
		Name:        ent.Name,
		Description: ent.Description,
		ImageId:     ent.ImageID,
		Price:       ent.Price.String(),
		CurrencyId:  ent.CurrencyID,
		Rating:      ent.Rating,
		CategoryId:  ent.CategoryID,
		CreatedAt:   ent.CreatedAt.UnixMilli(),
		UpdatedAt:   nil,
	}

	// TODO: updated_at возвращается как строка с 0

	if ent.UpdatedAt != nil {
		pbProduct.UpdatedAt = pointer.Pointer(ent.UpdatedAt.UnixMilli())
	}

	return pbProduct
}
