package swagger

import (
	"embed"
	"github.com/gorilla/mux"
	"io/fs"
	"net/http"
	"os"
)

//go:embed swagger-ui
var FS embed.FS

// New создает новый маршрутизатор для swagger-ui
func New(specPath string, opts struct{ SwaggerPath string }) *mux.Router {
	if _, err := os.Stat(specPath); os.IsNotExist(err) {
		panic(err)
	}

	router := mux.NewRouter()

	fileSys, err := fs.Sub(FS, "swagger-ui")
	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.FS(fileSys))

	if opts.SwaggerPath == "" {
		opts.SwaggerPath = "/swagger"
	}

	router.Handle(opts.SwaggerPath, fileServer)

	router.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, specPath)
	})

	return router
}
