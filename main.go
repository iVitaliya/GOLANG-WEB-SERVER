package main // https://www.youtube.com/watch?v=3y3g1RBCLs4&list=PL5dTjWUk_cPbExff-KfZ18abspIgfcfmt&index=5

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := routes.Rnd.Template(w, http.StatusOK, []string{"static/home.tpl"}, nil)
	if err != nil {
		routes.Logger.Error(fmt.Sprint(err))
	}

	// https://www.youtube.com/watch?v=3y3g1RBCLs4&list=PL5dTjWUk_cPbExff-KfZ18abspIgfcfmt&index=5
}

func main() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
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

	<-stopChan
	routes.Logger.Debug("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel(
		routes.Log(routes.DebugSeverity, "Server gracefully stopped")
	)
}
