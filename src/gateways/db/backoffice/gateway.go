package backoffice

import (
	"github.com/postech-soat2-grupo16/clientes-api/entities"
	"gorm.io/gorm"
)

type Gateway struct {
	repository *gorm.DB
}

func NewGateway(repository *gorm.DB) *Gateway {
	return &Gateway{repository: repository}
}

func (g *Gateway) Save(request entities.BackofficeRequest) (*entities.BackofficeRequest, error) {
	result := g.repository.Create(&request)
	if result.Error != nil {
		return nil, result.Error
	}

	return &request, nil
}

func (g *Gateway) GetAll(conds ...interface{}) (*[]entities.BackofficeRequest, error) {
	requests := &[]entities.BackofficeRequest{}
	result := g.repository.Find(requests, conds...)
	if result.Error != nil {
		return nil, result.Error
	}

	return requests, nil
}

func (g *Gateway) GetByID(requestID uint32) (*entities.BackofficeRequest, error) {
	request := &entities.BackofficeRequest{}
	result := g.repository.First(request, requestID)
	if result.Error != nil {
		return nil, result.Error
	}

	return request, nil
}

func (g *Gateway) Update(request entities.BackofficeRequest) (*entities.BackofficeRequest, error) {
	result := g.repository.Save(&request)
	if result.Error != nil {
		return nil, result.Error
	}

	return &request, nil
}
