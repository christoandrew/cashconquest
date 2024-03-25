package scopes

import (
	"gorm.io/gorm"
)

func GetAccountTransactions(db *gorm.DB, accountId int) *gorm.DB {
	return db.Preload("TransactionType").Preload("Category").Where("account_id = ?", accountId)
}

func GetRecentTransactions(db *gorm.DB, accountId int, limit int) *gorm.DB {
	return GetAccountTransactions(db, accountId).Order("date desc").Limit(limit)
}

func GetAccountTransactionsByCategory(accountId int, categoryId int, db *gorm.DB) *gorm.DB {
	return db.Where("account_id = ? AND category_id = ?", accountId, categoryId)
}
