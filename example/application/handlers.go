package application

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/thoas/djangogo/auth/models"
	"net/http"
)

var MeHandler Handler = func(app *Application, w http.ResponseWriter, r *Request) {
	fmt.Fprintf(w, fmt.Sprintf("Welcome %s", r.User.Username))
}

var UserHandler Handler = func(app *Application, w http.ResponseWriter, r *Request) {
	params := mux.Vars(r.Request)

	user := models.User{}

	app.DB.Where("username = ?", params["username"]).First(&user)

	fmt.Fprintf(w, fmt.Sprintf("Welcome %s", user.Username))
}
