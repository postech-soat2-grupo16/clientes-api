package interfaces

import (
	"github.com/postech-soat2-grupo16/clientes-api/entities"
)

type ClienteGatewayI interface {
	Save(cliente entities.Cliente) (*entities.Cliente, error)
	Update(cliente entities.Cliente) (*entities.Cliente, error)
	Delete(clienteID uint32) error
	GetByID(clienteID uint32) (*entities.Cliente, error)
	GetAll(conds ...interface{}) ([]entities.Cliente, error)
}
