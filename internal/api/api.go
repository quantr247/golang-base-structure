package api

import (
	"golang-base-structure/config"
	"golang-base-structure/internal/dto"
	"golang-base-structure/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type API struct {
	cfg                *config.Config
	applicationUseCase usecase.ApplicationUseCase
	userUseCase        usecase.UserUseCase
	transactionUseCase usecase.TransactionUseCase
}

func NewAPI(
	cfg *config.Config,
	applicationUseCase usecase.ApplicationUseCase,
	userUseCase usecase.UserUseCase,
	transactionUseCase usecase.TransactionUseCase,
) http.Handler {
	a := API{cfg, applicationUseCase, userUseCase, transactionUseCase}

	router := gin.Default()
	router.GET("/application", a.getApplication)
	router.GET("/user/:id", a.getUserByID)
	router.POST("/transaction", a.postTransaction)

	return router
}

func (a *API) getApplication(c *gin.Context) {
	application, err := a.applicationUseCase.GetApplication(c.Request.Context(), nil)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "application not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, application)
}

func (a *API) getUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := a.userUseCase.GetUserByID(c.Request.Context(), id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func (a *API) postTransaction(c *gin.Context) {
	var transactionReq *dto.TransactionRequestDTO
	if err := c.BindJSON(&transactionReq); err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	transactionRes, err := a.transactionUseCase.Payment(c.Request.Context(), transactionReq)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "payment error"})
		return
	}
	c.IndentedJSON(http.StatusCreated, transactionRes)
}
