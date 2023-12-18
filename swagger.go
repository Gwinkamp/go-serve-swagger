package swagger

import (
	"embed"
	"fmt"
	"github.com/gorilla/mux"
	"io/fs"
	"net/http"
	"os"
)

//go:embed swagger-ui
var FS embed.FS

// New создает новый маршрутизатор для swagger-ui
func New(specPath string, swaggerPath string) *mux.Router {
	if _, err := os.Stat(specPath); os.IsNotExist(err) {
		panic(err)
	}

	router := mux.NewRouter()

	fileSys, err := fs.Sub(FS, "swagger-ui")
	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.FS(fileSys))

	if swaggerPath == "" {
		swaggerPath = "/swagger"
	}

	router.Handle(
		fmt.Sprintf("%s/", swaggerPath),
		http.StripPrefix(swaggerPath, fileServer),
	).Methods(http.MethodGet)

	router.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, specPath)
	})

	return router
}
