package swagger

import (
	"embed"
	"io/fs"
	"mime"
	"net/http"
	"os"
	"strings"
)

var FS embed.FS

// route запроса на получение спецификации swagger в формате json
const swaggerPath = "/swagger.json"

// Handler формирует http.HandlerFunc обработчик запросов на получение спецификации swagger
func Handler(pathToSpec string) http.HandlerFunc {
	if _, err := os.Stat(pathToSpec); os.IsNotExist(err) {
		panic(err)
	}

	_ = mime.AddExtensionType(".svg", "image/svg+xml")

	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, swaggerPath) {
			http.ServeFile(w, r, pathToSpec)
		}
		swaggerUI, err := fs.Sub(FS, "swagger-ui")
		if err != nil {
			panic(err)
		}
		http.FileServer(http.FS(swaggerUI)).ServeHTTP(w, r)
	}
}
