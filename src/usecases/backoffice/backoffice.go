package backoffice

import (
	"github.com/postech-soat2-grupo16/clientes-api/entities"
	"github.com/postech-soat2-grupo16/clientes-api/interfaces"
)

type UseCase struct {
	backofficeGateway interfaces.BackofficeGatewayI
}

func NewUseCase(backofficeGateway interfaces.BackofficeGatewayI) *UseCase {
	return &UseCase{
		backofficeGateway: backofficeGateway,
	}
}

func (p *UseCase) ListRequests() (*[]entities.BackofficeRequest, error) {
	return p.backofficeGateway.GetAll()
}

func (p *UseCase) RequestToDeleteClientData(nome, endereco, telefone string) (*entities.BackofficeRequest, error) {
	request := entities.BackofficeRequest{
		Nome:     nome,
		Endereco: endereco,
		Telefone: telefone,
	}
	result, err := p.backofficeGateway.Save(request)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *UseCase) ProcessDataDeletion(requestID uint32) error {
	request, err := p.backofficeGateway.GetByID(requestID)
	if err != nil {
		return err
	}
	request.Processed = true
	_, err = p.backofficeGateway.Update(*request)
	if err != nil {
		return err
	}
	return nil
}
