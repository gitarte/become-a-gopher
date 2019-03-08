package main

import (
	"database/sql"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// List -
func List(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		files, err := ioutil.ReadDir("/files/")
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			log.Error().Msg(err.Error())
			return
		}

		result := make([]File, 1, 100)
		for i, f := range files {
			result = append(result, File{i, f.Name()})
		}

		c.JSON(http.StatusOK, result)
	}
}

// Read -
func Read(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		file := c.Param("file")
		data, err := ioutil.ReadFile("/files/" + file)
		if err != nil {
			c.String(http.StatusNotFound, err.Error())
			return
		}
		c.String(http.StatusOK, string(data))
	}
}

// GetAccounts -
func GetAccounts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		key, _ := url.QueryUnescape(c.Query("key"))
		val, _ := url.QueryUnescape(c.Query("val"))

		query := ""
		var param interface{}
		switch key {
		case "id":
			query = "SELECT id,username,password,email FROM account WHERE id = $1;"
			param, _ = strconv.Atoi(val)
		case "username":
			query = "SELECT id,username,password,email FROM account WHERE username = $1;"
			param = val
		case "password":
			query = "SELECT id,username,password,email FROM account WHERE password = $1;"
			param = val
		case "email":
			query = "SELECT id,username,password,email FROM account WHERE email = $1;"
			param = val
		default:
			query = "SELECT id,username,password,email FROM account;"
		}

		result := []Account{}

		var rows *sql.Rows
		var err error
		if param == nil {
			rows, err = db.Query(query)
		} else {
			rows, err = db.Query(query, param)
		}
		if err != nil {
			log.Debug().Msg(err.Error())
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		defer rows.Close()
		for rows.Next() {
			var a Account
			err := rows.Scan(&a.ID, &a.Username, &a.Password, &a.Email)
			if err != nil {
				log.Debug().Msg(err.Error())
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			result = append(result, a)
		}
		c.JSON(http.StatusOK, result)
	}
}

// PutAccounts -
func PutAccounts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

// PostAccounts -
func PostAccounts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var account Account
		if err := c.BindJSON(&account); err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		sql := "INSERT INTO account (username,password,email) VALUES ($1,$2,$3) RETURNING id;"
		err := db.QueryRow(sql, account.Username, account.Password, account.Email).Scan(&account.ID)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		//	return results and yes: 201 instead of 200!
		c.JSON(http.StatusCreated, account)
	}
}

// DeleteAccounts -
func DeleteAccounts(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		sql := "DELETE FROM account WHERE id=$1 RETURNING id;"
		_, err = db.Query(sql, id)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		//	return results and yes: 201 instead of 200!
		c.String(http.StatusOK, "OK")
	}
}
