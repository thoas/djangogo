package application

import (
	"github.com/thoas/djangogo/auth"
	"github.com/thoas/djangogo/auth/models"
	"github.com/thoas/djangogo/serializers"
	"github.com/thoas/djangogo/sessions"
	"net/http"
)

type Request struct {
	Request     *http.Request
	User        *models.User
	Session     *sessions.Session
	Application *Application
}

type Handler func(app *Application, w http.ResponseWriter, r *Request)

func (h Handler) ServeHTTP(app *Application, w http.ResponseWriter, r *http.Request) {
	var session *sessions.Session

	c, err := r.Cookie(app.CookieName)

	if err == nil {
		session = sessions.NewSession(c.Value, &serializers.PickleSerializer{}, app.Store)
	}

	req := &Request{r, &models.User{}, session, app}

	user, err := req.CurrentUser()

	if err == nil {
		req.User = user
	}

	h(app, w, req)
}

func (req *Request) CurrentUser() (*models.User, error) {
	user := models.User{}

	if req.Session != nil {
		value, err := req.Session.Get(auth.SESSION_AUTH_KEY)

		if err != nil {
			return nil, err
		}

		req.Application.DB.Where("id = ?", value).First(&user)
	}

	return &user, nil
}
