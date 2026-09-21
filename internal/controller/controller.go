package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/cris329/aprendiz-postgres/internal/model"
	"github.com/cris329/aprendiz-postgres/internal/service"
)

type AprendizController struct {
	svc *service.AprendizService
}

func NewAprendizController(svc *service.AprendizService) *AprendizController {
	return &AprendizController{svc: svc}
}

func (c *AprendizController) Create(ctx *gin.Context) {
	var a model.Aprendiz
	if err := ctx.ShouldBindJSON(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.svc.CreateAprendiz(&a); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, a)
}

func (c *AprendizController) GetAll(ctx *gin.Context) {
	aprendices, err := c.svc.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, aprendices)
}

func (c *AprendizController) GetByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}
	a, err := c.svc.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no encontrado"})
		return
	}
	ctx.JSON(http.StatusOK, a)
}

func (c *AprendizController) Update(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}
	var a model.Aprendiz
	if err := ctx.ShouldBindJSON(&a); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.svc.Update(uint(id), &a); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, a)
}

func (c *AprendizController) Delete(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "id invalido"})
		return
	}
	if err := c.svc.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "eliminado"})
}
