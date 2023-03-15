package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func PopulateDB() {

	jsonFile, _ := os.Open("links.json")
	byteValues, _ := io.ReadAll(jsonFile)

	var arr []UrlPair
	json.Unmarshal(byteValues, &arr)

	for _, val := range arr {
		qStr := fmt.Sprintf(`INSERT into urls(active, history) VALUES('%v','%v')`, val.Active, val.History)
		row, err := AWSDB.Exec(context.Background(), qStr)

		fmt.Printf(`%v - %v`, row.RowsAffected(), err)

	}

}
