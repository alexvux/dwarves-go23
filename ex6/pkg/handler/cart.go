package handler

import (
	"net/http"
	"strconv"

	"github.com/alexvux/dwarves-go23/ex6/pkg/model"
	"github.com/alexvux/dwarves-go23/ex6/pkg/repo"
	"github.com/alexvux/dwarves-go23/ex6/pkg/util"
	"github.com/gin-gonic/gin"
)

func AddItemToCart(c *gin.Context) {
	var item model.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		util.BindJSONWithError(c, http.StatusBadRequest, err)
		return
	}
	if err := repo.AddItemToCart(item); err != nil {
		util.BindJSONWithError(c, http.StatusBadRequest, err)
		return
	}
	util.BindJSONWithMessage(c, http.StatusCreated, "add item to cart successfully")
}

func DeleteItemFromCart(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.BindJSONWithError(c, http.StatusBadRequest, err)
		return
	}

	if err := repo.DeleteItemFromCart(id); err != nil {
		util.BindJSONWithError(c, http.StatusBadRequest, err)
		return
	}
	util.BindJSONWithMessage(c, http.StatusOK, "delete item from cart successfully")
}

func Checkout(c *gin.Context) {
	reciept, err := repo.Checkout()
	if err != nil {
		util.BindJSONWithError(c, http.StatusBadRequest, err)
		return
	}
	c.String(http.StatusOK, reciept)
}
