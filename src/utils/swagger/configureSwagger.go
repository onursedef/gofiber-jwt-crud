package swagger

import (
	"os"

	"backend.go/docs"
)

func Configure() {
	docs.SwaggerInfo.Title = os.Getenv("APP_NAME")
	docs.SwaggerInfo.Description = os.Getenv("APP_DESCRIPTION")
	docs.SwaggerInfo.Version = os.Getenv("APP_VERSION")
	docs.SwaggerInfo.Host = os.Getenv("APP_URL")
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
