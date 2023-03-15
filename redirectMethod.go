package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Object struct {
	Status  int    `json:"status"`
	Working string `json:"active_link"`
}

func Redirect(w http.ResponseWriter, r *http.Request) {

	url := r.URL.Query().Get("url")
	if url == "" {
		json.NewEncoder(w).Encode(Response{
			Message: "Enter link into query params",
			Status:  400,
		})
		return
	}
	var obj Object

	// Check cache first
	value, status := PerformGet(&Box, url)

	if status != -1 {
		obj.Status = status
		obj.Working = value
		fmt.Println("From Cache")
	} else {
		qStr := fmt.Sprintf(`SELECT * FROM checklink('%v')`, url)

		row := AWSDB.QueryRow(context.Background(), qStr)

		row.Scan(&obj.Status, &obj.Working)

		// Store or Update data in the cache
		PerformAdd(&Box, url, obj.Working)
		fmt.Println("From DB")
	}

	json.NewEncoder(w).Encode(obj)
}
