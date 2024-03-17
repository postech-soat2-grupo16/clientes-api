package api

import (
	"github.com/aws/aws-sdk-go/service/sns"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/postech-soat2-grupo16/clientes-api/controllers"
	"github.com/postech-soat2-grupo16/clientes-api/external"
	bg "github.com/postech-soat2-grupo16/clientes-api/gateways/db/backoffice"
	cg "github.com/postech-soat2-grupo16/clientes-api/gateways/db/cliente"
	ng "github.com/postech-soat2-grupo16/clientes-api/gateways/notification"

	"github.com/postech-soat2-grupo16/clientes-api/usecases/backoffice"
	"github.com/postech-soat2-grupo16/clientes-api/usecases/cliente"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

func SetupDB() *gorm.DB {
	dialector := external.GetPostgresDialector()
	db := external.NewORM(dialector)

	return db
}

func SetupNotification() *sns.SNS {
	return external.GetSnsClient()
}

func SetupRouter(db *gorm.DB, notification *sns.SNS) *chi.Mux {
	r := chi.NewRouter()
	r.Use(commonMiddleware)

	mapRoutes(r, db, notification)

	return r
}

func mapRoutes(r *chi.Mux, orm *gorm.DB, notification *sns.SNS) {
	// Swagger
	r.Get("/swagger/*", httpSwagger.Handler())

	// Injections
	// Gateways
	clienteGateway := cg.NewGateway(orm)
	backofficeGateway := bg.NewGateway(orm)
	notificationGateway := ng.NewGateway(notification)
	// Use cases
	clienteUseCase := cliente.NewUseCase(clienteGateway, notificationGateway)
	backOfficeUseCase := backoffice.NewUseCase(backofficeGateway)
	// Handlers
	controllers.NewClienteController(clienteUseCase, r)
	controllers.NewBackofficeController(backOfficeUseCase, r)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
