package backoffice

import "github.com/postech-soat2-grupo16/clientes-api/entities"

type BackofficeRequest struct {
	ID        uint32 `json:"id"`
	Nome      string `json:"nome"`
	Endereco  string `json:"endereco"`
	Telefone  string `json:"telefone"`
	Processed bool   `json:"processed"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func FromUseCaseEntity(request *entities.BackofficeRequest) BackofficeRequest {
	return BackofficeRequest{
		ID:        request.ID,
		Nome:      request.Nome,
		Endereco:  request.Endereco,
		Telefone:  request.Telefone,
		Processed: request.Processed,
		CreatedAt: request.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: request.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: request.DeletedAt.Time.Format("2006-01-02 15:04:05"),
	}
}
