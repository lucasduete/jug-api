package controller

import (
	"net/http"
	"fmt"
)

func (app *App) NotFound(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "ERROUUUUUU")
}
