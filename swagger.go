package swagger

import (
	"embed"
	"fmt"
	"github.com/gorilla/mux"
	"io/fs"
	"net/http"
	"os"
)

var (
	Router *mux.Router
	FS     embed.FS
)

// New создает новый маршрутизатор для swagger-ui
func New(specPath string, swaggerPath string) *mux.Router {
	if _, err := os.Stat(specPath); os.IsNotExist(err) {
		panic(err)
	}

	Router = mux.NewRouter()

	fileSys, err := fs.Sub(FS, "swagger-ui")
	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.FS(fileSys))

	if swaggerPath == "" {
		swaggerPath = "/swagger"
	}

	Router.Handle(
		fmt.Sprintf("%s/", swaggerPath),
		http.StripPrefix(swaggerPath, fileServer),
	).Methods(http.MethodGet)

	Router.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, specPath)
	})

	return Router
}
