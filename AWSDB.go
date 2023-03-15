package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"io"
	"os"
)

var AWSDB = GetAWSDB()

var (
	DBPORT     string
	USER       string
	PASSWORD   string
	DBNAME     string
	HOST       string
	SECRET_KEY string
)

func AuthDetails() {

	jsonFile, _ := os.Open("db_info.json")
	byteValues, _ := io.ReadAll(jsonFile)

	var data map[string]string
	json.Unmarshal(byteValues, &data)
	USER = data["USER"]
	PASSWORD = data["PASSWORD"]
	HOST = data["HOST"]
	DBPORT = data["DBPORT"]
	DBNAME = data["DBNAME"]
	SECRET_KEY = data["SECRET_KEY"]

}

func GetAWSDB() *pgxpool.Pool {
	AuthDetails()

	qStr := fmt.Sprintf("postgresql://%v:%v@%v:%v/%v", USER, PASSWORD, HOST, DBPORT, DBNAME)

	db, err := pgxpool.Connect(context.Background(), qStr)

	if err != nil {
		panic(err)
	}

	return db
}
