package main

import (
	"avito-backend-trainee-assignment-2023/internal/api"
	"avito-backend-trainee-assignment-2023/pkg/logging"
)

var logger = logging.GetLogger()

func main() {
	api.StartAPIServer()
}
