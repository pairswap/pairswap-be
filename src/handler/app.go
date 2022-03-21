package handler

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sisu-network/lib/log"
	"github.com/sisu-network/pairswap-be/src/config"
	"github.com/sisu-network/pairswap-be/src/store"
)

type App struct {
	Router *mux.Router
	Db     *store.DBStores

	Config config.AppConfig
}

func NewApp() *App {
	cfg := config.NewDefaultAppConfig()
	db, err := store.NewDBStores(cfg.DB)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	app := &App{Router: r, Db: db, Config: cfg}
	app.registerMiddlewares()
	app.registerHandlers()
	return app
}

func (a *App) Start() {
	log.Infof("Server listening at port %d ...", a.Config.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", a.Config.Port), a.Router); err != nil {
		panic(err)
	}
}

func (a *App) registerHandlers() {
	supportFormHandler := NewSupportFormHandler(a.Db)
	supportFormRoute := &Route{
		Path:    "/support",
		Method:  http.MethodPost,
		Handler: supportFormHandler.HandleSubmitSupportForm,
	}

	a.Router.HandleFunc(supportFormRoute.Path, supportFormRoute.Handler).Methods(supportFormRoute.Method)
}

func (a *App) registerMiddlewares() {
	a.Router.Use(customCORSHeader())
}

func customCORSHeader() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, req)
		})
	}
}
