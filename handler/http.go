package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ridhozhr10/gdsc-backend-toko/entities"
	"github.com/ridhozhr10/gdsc-backend-toko/service/product"
)

type GinHttp struct {
	ProductService product.Service
}

func NewGinHttp(
	productService product.Service,
) *GinHttp {
	return &GinHttp{
		ProductService: productService,
	}
}

func (gh *GinHttp) GetAllProduct(c *gin.Context) {
	data, err := gh.ProductService.GetAllProduct()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": data,
	})
}

func (gh *GinHttp) GetOneProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}
	data, err := gh.ProductService.GetOneProduct(idInt)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": data,
	})
}

func (gh *GinHttp) CreateProduct(c *gin.Context) {
	req := entities.Product{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}
	data, err := gh.ProductService.CreateProduct(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": data,
	})
}

func (gh *GinHttp) UpdateProduct(c *gin.Context) {
	req := entities.Product{}
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}
	req.ID = uint(idInt)
	data, err := gh.ProductService.UpdateProduct(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": data,
	})
}

func (gh *GinHttp) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}
	if err := gh.ProductService.DeleteProduct(idInt); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"ok":  false,
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"ok":  true,
		"msg": "success",
	})
}
