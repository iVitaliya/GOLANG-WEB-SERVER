package main

import (
	"fmt"
	"main/database"
	"main/routes"

	"github.com/iVitaliya/GOLANG-WEB-SERVER/database"
	"github.com/iVitaliya/GOLANG-WEB-SERVER/routes"
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
)

func init() {
	routes.Rnd = renderer.New()
	sess, err := mgo.Dial(database.HostName)
	if err != nil {
		routes.Log(routes.FatalSeverity, "Database connection errorred: ", err.Error())
	}

	sess.SetMode(mgo.Monotonic, true)
}

func main() {
	fmt.Println("Hello, World!")
}
