package handler

import (
	"net/http"
	"strconv"
	"vue-golang-payment-app/backend-api/db"
)

// context handleRawConn
func GetLists(c Context) {
	res, err := db.SelectAllItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, res)
}

// context getItem
func GetItem(c Context) {
	identifier, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	res, err := db.SelectItem(int64(identifier))
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	c.JSON(http.StatusOK, res)
}
