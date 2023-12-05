package main

import (
	"net/http"

	"section2/packages/controllers"
)

func main() {

	controllers.RegisterControllers()

	http.ListenAndServe(":3000", nil)
}
