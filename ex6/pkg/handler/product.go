package handler

import (
	"net/http"
	"strconv"

	"github.com/alexvux/dwarves-go23/ex6/pkg/model"
	"github.com/alexvux/dwarves-go23/ex6/pkg/repo"
	"github.com/alexvux/dwarves-go23/ex6/pkg/util"
	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	products, err := repo.GetAllProducts()
	if err != nil {
		util.BindJSONWithError(c, http.StatusNotFound, err)
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, products)
}

func AddProduct(c *gin.Context) {
	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		util.BindJSONWithError(c, http.StatusBadRequest, err)
		return
	}
	if err := repo.AddProduct(product); err != nil {
		util.BindJSONWithError(c, http.StatusBadRequest, err)
		return
	}
	util.BindJSONWithMessage(c, http.StatusCreated, "add product successfully")
}

func UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.BindJSONWithError(c, http.StatusBadRequest, err)
		return
	}

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		util.BindJSONWithError(c, http.StatusBadRequest, err)
		return
	}

	if err := repo.UpdateProduct(id, product); err != nil {
		util.BindJSONWithError(c, http.StatusBadRequest, err)
		return
	}
	util.BindJSONWithMessage(c, http.StatusOK, "update product successfully")
}

func DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		util.BindJSONWithError(c, http.StatusBadRequest, err)
		return
	}

	if err := repo.DeleteProduct(id); err != nil {
		util.BindJSONWithError(c, http.StatusBadRequest, err)
		return
	}
	util.BindJSONWithMessage(c, http.StatusOK, "delete product successfully")
}
