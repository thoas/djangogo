package application

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/thoas/djangogo/sessions/store"
	"net/http"
)

type Application struct {
	DB         gorm.DB
	CookieName string
	Options    *Option
	Store      store.Store
}

type Option struct {
	Database map[string]string
	Session  map[string]string
}

func New(cookieName string, options *Option) (*Application, error) {
	app := &Application{CookieName: cookieName}

	DB, err := gorm.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		options.Database["USER"], options.Database["PASSWORD"], options.Database["NAME"]))

	if err != nil {
		return nil, err
	}

	DB.LogMode(true)
	DB.DB()
	DB.DB().Ping()
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(1000)

	app.DB = DB
	app.Options = options

	store, err := store.NewRedisStore(1000, "tcp",
		fmt.Sprintf(":%s",
			options.Session["PORT"]), "", options.Session["DATABASE"], options.Session["PREFIX"])

	if err != nil {
		return nil, err
	}

	app.Store = store

	return app, nil
}

func (app *Application) ServeHTTP(h Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		h.ServeHTTP(app, w, req)
	})
}

func (app *Application) Run(port int) {
	router := mux.NewRouter()

	router.Handle("/me", app.ServeHTTP(MeHandler))
	router.Handle("/{username}", app.ServeHTTP(UserHandler))

	n := negroni.Classic()
	n.UseHandler(router)

	n.Run(fmt.Sprintf(":%d", port))
}
