package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/postech-soat2-grupo16/clientes-api/adapters/backoffice"
	"github.com/postech-soat2-grupo16/clientes-api/interfaces"
)

type BackofficeController struct {
	useCase interfaces.BackofficeUseCase
}

func NewBackofficeController(useCase interfaces.BackofficeUseCase, r *chi.Mux) *BackofficeController {
	controller := BackofficeController{useCase: useCase}
	r.Route("/backoffice", func(r chi.Router) {
		r.Get("/", controller.ListRequests())
		r.Post("/request-data-deletion", controller.RequestToDeleteClientData())
		r.Patch("/{id}/process-data-deletion", controller.ProcessDataDeletion())
	})
	return &controller
}

// @Summary	Get all requests
// @Tags		Backoffice
// @ID			get-all-requests
// @Produce	json
// @Success	200	{object}	backoffice.BackofficeRequest
// @Failure	500
// @Router		/backoffice [get]
func (c *BackofficeController) ListRequests() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requests, err := c.useCase.ListRequests()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		requestsJSON := make([]backoffice.BackofficeRequest, 0)
		for _, request := range *requests {
			requestsJSON = append(requestsJSON, backoffice.FromUseCaseEntity(&request))
		}

		json.NewEncoder(w).Encode(requestsJSON)
	}
}

// @Summary	Request to delete client data
// @Tags		Backoffice
// @ID			request-data-deletion
// @Accept		json
// @Produce		json
// @Param		data	body		backoffice.BackofficeRequest	true	"Backoffice data"
// @Success	200	{object}	backoffice.BackofficeRequest
// @Failure	400
// @Failure	500
// @Router		/backoffice/request-data-deletion [post]
func (c *BackofficeController) RequestToDeleteClientData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request backoffice.BackofficeRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		requestCreated, err := c.useCase.RequestToDeleteClientData(request.Nome, request.Endereco, request.Telefone)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(backoffice.FromUseCaseEntity(requestCreated))
	}
}

// @Summary	Process data deletion
// @Tags		Backoffice
// @ID			process-data-deletion
// @Produce		json
// @Param		id	path	int	true	"Request ID"
// @Success	200
// @Failure	400
// @Failure	500
// @Router		/backoffice/{id}/process-data-deletion [patch]
func (c *BackofficeController) ProcessDataDeletion() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestID := chi.URLParam(r, "id")
		if requestID == "" {
			http.Error(w, "Invalid request ID", http.StatusBadRequest)
			return
		}
		requestIDUint32, err := strconv.ParseUint(requestID, 10, 32)
		if err != nil {
			http.Error(w, "Invalid request ID", http.StatusBadRequest)
			return
		}

		err = c.useCase.ProcessDataDeletion(uint32(requestIDUint32))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
