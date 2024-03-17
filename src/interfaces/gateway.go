package interfaces

import (
	"github.com/postech-soat2-grupo16/clientes-api/entities"
)

type ClienteGatewayI interface {
	Save(cliente entities.Cliente) (*entities.Cliente, error)
	Update(cliente entities.Cliente) (*entities.Cliente, error)
	Delete(clienteID uint32) error
	GetByID(clienteID uint32) (*entities.Cliente, error)
	GetAll(conds ...interface{}) (*[]entities.Cliente, error)
}

type BackofficeGatewayI interface {
	Save(request entities.BackofficeRequest) (*entities.BackofficeRequest, error)
	GetAll(conds ...interface{}) (*[]entities.BackofficeRequest, error)
	GetByID(requestID uint32) (*entities.BackofficeRequest, error)
	Update(request entities.BackofficeRequest) (*entities.BackofficeRequest, error)
}

type NotificationGatewayI interface {
	ClientSubscriber(cliente *entities.Cliente) error
}
