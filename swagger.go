package swagger

import (
	"embed"
	"io"
	"io/fs"
	"net/http"
	"os"
	"strconv"
)

//go:embed swagger-ui
var FS embed.FS

// Handler создает обработчик запросов для получения swagger документации
func Handler(specPath string) http.HandlerFunc {
	specFile, err := os.Open(specPath)
	if err != nil {
		panic(err)
	}
	defer specFile.Close()

	spec, err := io.ReadAll(specFile)
	if err != nil {
		panic(err)
	}

	fileSys, err := fs.Sub(FS, "swagger-ui")
	if err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.FS(fileSys))

	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/swagger.json" {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Content-Length", strconv.Itoa(len(spec)))
			w.Write(spec)
			return
		}
		fileServer.ServeHTTP(w, r)
	}
}
