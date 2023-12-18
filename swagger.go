package swagger

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var Router *mux.Router

// New создает новый маршрутизатор для swagger-ui
func New(specPath string, swaggerPath string) *mux.Router {
	if _, err := os.Stat(specPath); os.IsNotExist(err) {
		panic(err)
	}

	Router = mux.NewRouter()
	fs := http.FileServer(http.Dir("swagger-ui"))

	if swaggerPath == "" {
		swaggerPath = "/swagger"
	}

	Router.Handle(
		fmt.Sprintf("%s/", swaggerPath),
		http.StripPrefix(swaggerPath, fs),
	).Methods(http.MethodGet)

	Router.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, specPath)
	})

	return Router
}
