package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"regexp"
	"strconv"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type UrlPair struct {
	Id      int    `json:"id,omitempty"`
	Active  string `json:"active_link"`
	History string `json:"history_link"`
}

func (u *UrlPair) Validate() error {
	fmt.Println(u)

	if u.Active == "" || u.History == "" {
		return errors.New("body can't be empty")
	}

	exp := regexp.MustCompile("/")
	line1 := exp.Split(u.Active, -1)
	line2 := exp.Split(u.History, -1)

	if len(line1) == 1 || len(line2) == 1 {
		return errors.New("enter valid url")
	}

	exp = regexp.MustCompile(" ")
	line1 = exp.Split(u.Active, -1)
	line2 = exp.Split(u.History, -1)

	if len(line1) != 1 || len(line2) != 1 {
		return errors.New("urls must not contain white spaces")
	}
	return nil
}

func ValidateQuery(r *http.Request) (error, int, int, int) {
	pageStr := r.URL.Query().Get("page")
	perPageStr := r.URL.Query().Get("per_page")

	if pageStr == "" {
		return errors.New("enter page number"), 400, 0, 0
	}
	if perPageStr == "" {
		return errors.New("enter items per page number"), 400, 0, 0
	}

	pageInt, err1 := strconv.Atoi(pageStr)
	perPageInt, err2 := strconv.Atoi(perPageStr)

	if err1 != nil || err2 != nil {
		return errors.New("page number and Items per page must be a number"), 400, 0, 0
	}

	return nil, 200, pageInt, perPageInt
}

func DataControl(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		vars := mux.Vars(r)

		if vars["id"] != "" {
			id, _ := strconv.Atoi(vars["id"])

			qStr := fmt.Sprintf(`SELECT * FROM urls WHERE id=%d`, id)

			row := AWSDB.QueryRow(context.Background(), qStr)
			var pair UrlPair
			err := row.Scan(&pair.Id, &pair.Active, &pair.History)

			if err != nil {
				json.NewEncoder(w).Encode(Response{
					Message: "This ID does not exist",
					Status:  400,
				})
			} else {
				json.NewEncoder(w).Encode(pair)
			}
			return
		}

		err, status, page, items := ValidateQuery(r)

		if err != nil {
			json.NewEncoder(w).Encode(Response{
				Message: err.Error(),
				Status:  status,
			})
		}

		qStr := fmt.Sprintf(`SELECT * FROM urls LIMIT %d OFFSET %d`, items, (page-1)*items)

		rows, err := AWSDB.Query(context.Background(), qStr)

		if err != nil {
			fmt.Println(err)
		}

		data := make([]UrlPair, 0, items)

		for rows.Next() {
			var pair UrlPair
			err := rows.Scan(&pair.Id, &pair.Active, &pair.History)
			if err != nil {
				fmt.Println(err.Error())
			}
			data = append(data, pair)
		}
		if len(data) != 0 {
			json.NewEncoder(w).Encode(data)
		} else {
			json.NewEncoder(w).Encode(Response{
				Message: "There is no data on this page",
				Status:  400,
			})
		}

	}

	if r.Method == "POST" {
		var newPair UrlPair
		err := json.NewDecoder(r.Body).Decode(&newPair)
		if err != nil {
			fmt.Println(err.Error())
		}

		err = newPair.Validate()
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		qStr := fmt.Sprintf(`INSERT into urls(active, history) VALUES ('%v','%v')`, newPair.Active, newPair.History)

		rows, _ := AWSDB.Exec(context.Background(), qStr)
		n := rows.RowsAffected()

		var resp Response

		if n == 0 {
			resp.Message = "The link already exists"
			resp.Status = 400
		} else {
			resp.Message = "The link added"
			resp.Status = 200
		}

		json.NewEncoder(w).Encode(resp)
	}

	if r.Method == "PATCH" {
		var updated UrlPair
		json.NewDecoder(r.Body).Decode(&updated)

		if updated.Id == 0 {
			json.NewEncoder(w).Encode(Response{
				Message: "Provide Id of the record",
				Status:  400,
			})
			return
		}

		err := updated.Validate()

		if err != nil {
			json.NewEncoder(w).Encode(Response{
				Message: err.Error(),
				Status:  400,
			})
			return
		}
		fmt.Println(updated)

		qStr := fmt.Sprintf(`UPDATE urls SET active='%v',history='%v' WHERE id=%v`,
			updated.Active, updated.History, updated.Id,
		)

		_, err = AWSDB.Exec(context.Background(), qStr)

		var message string
		var status int

		if err != nil {
			message = "This active or/and history link(s) is(are) already occupied"
			status = 400
		} else {
			message = fmt.Sprintf(`Record Id=%v is updated`, updated.Id)
			status = 200
		}

		json.NewEncoder(w).Encode(Response{
			Message: message,
			Status:  status,
		})
	}

	if r.Method == "DELETE" {
		var deleting UrlPair
		json.NewDecoder(r.Body).Decode(&deleting)

		if deleting.Id == 0 {
			json.NewEncoder(w).Encode(Response{
				Message: "Enter numerical non-zero Id to be deleted",
				Status:  400,
			})
			return
		}

		qStr := fmt.Sprintf(`DELETE from urls WHERE id=%d`, deleting.Id)

		row, err := AWSDB.Exec(context.Background(), qStr)

		var message string
		var status int

		if err != nil {
			message = err.Error()
			status = 400
		} else if row.RowsAffected() == 0 {
			message = "This Id does not exist"
			status = 400

		} else {
			message = fmt.Sprintf(`Record Id=%v was deleted`, deleting.Id)
			status = 200
		}
		json.NewEncoder(w).Encode(Response{
			Message: message,
			Status:  status,
		})
	}
}
