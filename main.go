package main // https://www.youtube.com/watch?v=3y3g1RBCLs4&list=PL5dTjWUk_cPbExff-KfZ18abspIgfcfmt&index=5

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/iVitaliya/GOLANG-WEB-SERVER/database"
	"github.com/iVitaliya/GOLANG-WEB-SERVER/routes"
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
)

func init() {
	routes.Rnd = renderer.New()
	sess, err := mgo.Dial(database.HostName)
	if err != nil {
		routes.Logger.Fatal("Database connection errorred: " + err.Error())
	}

	sess.SetMode(mgo.Monotonic, true)
	database.Db = sess.DB(database.DbName)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler())
	r.Mount("/todo", routes.TodoHandlers())

	srv := &http.Server{
		Addr:         database.Ports.ToDo,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		routes.Logger.Info("Listening on port " + database.Ports.ToDo)
		if err := srv.ListenAndServe(); err != nil {
			routes.Logger.Debug("listen:" + fmt.Sprint(err))
		}
	}()
}
