package handlers

import (
	"net/http"
	"strconv"

	"github.com/christo-andrew/maybe-go/internal/api/requests"
	"github.com/christo-andrew/maybe-go/internal/api/serializers"
	"github.com/christo-andrew/maybe-go/internal/models"
	"github.com/christo-andrew/maybe-go/pkg/database/scopes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllTransactionsHandler GetAllTransactions godoc
// @Summary Get all transactions
// @Description Retrieve all transactions
// @Produce json
// @Success 200 {array} responses.TransactionResponse
// @Router /transactions [get]
// @Tags transactions
func GetAllTransactionsHandler(c *gin.Context, db *gorm.DB) {
	serializer := serializers.NewTransactionSerializer(getAllTransactions(db), true)
	c.JSON(http.StatusOK, serializer.Serialize())
}

func getAllTransactions(db *gorm.DB) []models.Transaction {
	var transactions []models.Transaction
	scopes.GetAllTransactions(db).Find(&transactions)
	return transactions
}

// GetTransactionHandler GetTransaction godoc
// @Summary Get a transaction
// @Description Retrieve a transaction
// @Produce json
// @Param id path string true "Transaction ID"
// @Success 200 {object} responses.TransactionResponse
// @Router /transactions/{id} [get]
// @Tags transactions
func GetTransactionHandler(c *gin.Context, db *gorm.DB) {
	transactionId, _ := strconv.Atoi(c.Param("id"))
	response := serializers.NewTransactionSerializer(getTransaction(transactionId, db), false).Serialize()
	c.JSON(http.StatusOK, response)
}

func getTransaction(transactionID int, db *gorm.DB) models.Transaction {
	var transaction models.Transaction
	scopes.GetTransactionById(transactionID, db).Find(&transaction)
	return transaction
}

// CreateAccountTransactionHandler CreateTransaction godoc
// @Summary Create a transaction
// @Description Create a transaction
// @Accept json
// @Produce json
// @Param transaction body requests.CreateTransactionRequest true "Create Transaction Request"
// @Success 201 {object} responses.TransactionResponse
// @Failure 400 {object} responses.ErrorResponse
// @Router /transactions/create [post]
// @Tags transactions
func CreateAccountTransactionHandler(c *gin.Context, db *gorm.DB) {
	var transactionRequest requests.CreateTransactionRequest
	err := c.ShouldBindJSON(&transactionRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaction, err := createTransaction(&transactionRequest, db)
	response := serializers.NewTransactionSerializer(transaction, false).Serialize()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, response)
}

func createTransaction(transactionRequest *requests.CreateTransactionRequest, db *gorm.DB) (*models.Transaction, error) {
	transaction := transactionRequest.Transaction()
	result := db.Create(transaction)

	return transaction, result.Error
}
