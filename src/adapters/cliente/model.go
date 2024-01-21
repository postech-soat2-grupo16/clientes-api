package cliente

import "github.com/postech-soat2-grupo16/clientes-api/entities"

type Cliente struct {
	ID        uint32 `json:"id"`
	CPF       string `json:"cpf"`
	Email     string `json:"email"`
	Nome      string `json:"nome"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
}

func FromUseCaseEntity(cliente *entities.Cliente) *Cliente {
	return &Cliente{
		ID:        cliente.ID,
		CPF:       cliente.CPF,
		Email:     cliente.Email,
		Nome:      cliente.Name,
		CreatedAt: cliente.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: cliente.UpdatedAt.Format("2006-01-02 15:04:05"),
		DeletedAt: cliente.DeletedAt.Time.Format("2006-01-02 15:04:05"),
	}
}
