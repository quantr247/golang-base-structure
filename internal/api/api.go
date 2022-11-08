package api

import (
	"golang-base-structure/config"
	"golang-base-structure/internal/common"
	"golang-base-structure/internal/dto"
	"golang-base-structure/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	_ "golang-base-structure/cmd/server/statik"

	"github.com/rakyll/statik/fs"
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

	statikFS, err := fs.New()
	if err != nil {
		zap.S().Panic("Can't create statik", zap.Error(err))
	}

	router := gin.Default()

	// swagger:route GET /application application getApplication
	// Get application list
	// responses:
	//
	//	200: GetApplicationResponseDTO
	router.GET("/application", a.getApplication)

	// swagger:route GET /user/{id} user getUserByID
	// Get user by id
	// responses:
	//
	//	200: UserResponseDTO
	router.GET("/user/:id", a.getUserByID)

	// swagger:route POST /transaction transaction postTransaction
	// Create payment transaction
	//
	// responses:
	//
	//	200: TransactionResponseDTO
	router.POST("/transaction", a.postTransaction)
	router.StaticFS("/swagger-ui/", statikFS)

	return router
}

func (a *API) getApplication(c *gin.Context) {
	application, err := a.applicationUseCase.GetApplication(c.Request.Context(), nil)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": common.ParseError(err).Message()})
		return
	}
	c.IndentedJSON(http.StatusOK, application)
}

func (a *API) getUserByID(c *gin.Context) {
	id := c.Param("id")

	user, err := a.userUseCase.GetUserByID(c.Request.Context(), id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": common.ParseError(err).Message()})
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
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": common.ParseError(err).Message()})
		return
	}
	c.IndentedJSON(http.StatusCreated, transactionRes)
}
