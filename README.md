# Go serve swagger

обработчик http запросов для отображения swagger документации

### Установка

```shell
go get github.com/Gwinkamp/go-serve-swagger
```

### Использование

Необходимо передать путь до файла со swagger спецификацией обработчику следующим образом:

```go
package main

import (
	swagger "github.com/Gwinkamp/go-serve-swagger"
	"net/http"
)


func main() {
	httpMux := http.NewServeMux()
	
	httpMux.Handle(
		"/swagger/",
		http.StripPrefix("/swagger", swagger.Handler("path/to/swagger.json")),
	)

	if err := http.ListenAndServe(":5050", httpMux); err != nil {
		panic(err)
	}
}
```

После этого можно открывать в браузере `http://localhost:5050/swagger/` и вы увидете swagger документацию.