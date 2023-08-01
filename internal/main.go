package internal

import (
	"go-mini/docs"
	"go-mini/infras"
	"go-mini/internal/domains/run"
	"go-mini/internal/handlers"
	"go-mini/internal/log"
	"go-mini/internal/middleware"
	"net/http"

	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupAndServe() {
	db, err := infras.GetDB()
	logger := log.GetLogger()
	if err != nil {
		logger.Fatal("Error opening database:", err)
	}
	logger.Info("Connected to DB")
	defer db.Close()

	r := chi.NewRouter()
	logger.Info("setting up middleware")
	r.Use(middleware.JsonMiddleware)
	docs.SwaggerInfo.Title = "Run Tracker"
	docs.SwaggerInfo.Version = "v1"
	conf := httpSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.Get("/swagger/*", httpSwagger.Handler(conf))
	logger.Info("setting up routes")
	r.Route("/v1/runs", func(r chi.Router) {

		repo := run.NewRunRepository(db, logger)
		service := run.NewRunService(repo)
		handler := handlers.NewRunHandler(*service)

		r.Get("/", handler.HandleGetAll)
		r.Post("/", handler.HandlePost)
		r.Put("/{id}", handler.HandleUpdate)
		r.Delete("/{id}", handler.HandleDelete)
	})
	logger.Info("server up and running")
	logger.Fatal(http.ListenAndServe(":8080", r))

}
