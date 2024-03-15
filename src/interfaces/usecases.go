package interfaces

import (
	"github.com/postech-soat2-grupo16/clientes-api/entities"
)

type ClienteUseCase interface {
	List(cpf string) (*[]entities.Cliente, error)
	Create(email, cpf, nome string) (*entities.Cliente, error)
	GetByID(clienteID uint32) (*entities.Cliente, error)
	Update(clienteID uint32, email, cpf, nome string) (*entities.Cliente, error)
	Delete(clienteID uint32) error
}

type BackofficeUseCase interface {
	RequestToDeleteClientData(nome, endereco, telefone string) (*entities.BackofficeRequest, error)
	ListRequests() (*[]entities.BackofficeRequest, error)
	ProcessDataDeletion(requestID uint32) error
}
