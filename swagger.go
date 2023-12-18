package swagger

import (
	"embed"
	"io/fs"
	"net/http"
	"os"
)

//go:embed swagger-ui
var FS embed.FS

// New создает новый маршрутизатор для swagger-ui
func New(specPath string) *http.ServeMux {
	if _, err := os.Stat(specPath); os.IsNotExist(err) {
		panic(err)
	}

	mux := http.NewServeMux()

	fileSys, err := fs.Sub(FS, "swagger-ui")
	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.FS(fileSys))

	mux.Handle("/swagger/", http.StripPrefix("/swagger", fileServer))

	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, specPath)
	})

	return mux
}
