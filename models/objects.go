package models

import "os"

//#region type Enviroment
type Enviroment struct {
	DatabaseHost string
	DatabasePort string
	DatabaseUser string
	DatabasePass string
	DatabaseName string
}

func (env *Enviroment) Load() {
	env.DatabaseHost = os.Getenv("DATABASE_HOST")
	env.DatabasePort = os.Getenv("DATABASE_PORT")
	env.DatabaseUser = os.Getenv("DATABASE_USER")
	env.DatabasePass = os.Getenv("DATABASE_PASS")
	env.DatabaseName = os.Getenv("DATABASE_NAME")
}

//#endregion
