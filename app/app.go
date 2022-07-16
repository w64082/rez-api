package app

import "database/sql"

type App struct {
	DbHandle *sql.DB
}

var Container = App{}
