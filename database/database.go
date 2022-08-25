package database

import mgo "gopkg.in/mgo.v2"

type IDatabases struct {
	ToDo string
}

type IPorts struct {
	ToDo string
}

var (
	Db          *mgo.Database
	Collections *IDatabases = &IDatabases{
		ToDo: "todo",
	}
	Ports *IPorts = &IPorts{
		ToDo: ":9000",
	}
)

const (
	HostName string = "GOLANGWEBSERVER.ranked.repl.co"
	DbName   string = "web_server"
)
