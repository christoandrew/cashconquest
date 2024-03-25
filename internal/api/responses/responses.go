package responses

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreateUserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type LoginResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Token string `json:"token"`
}

type AccountResponse struct {
	ID          int     `json:"id"`
	AccountName string  `json:"name"`
	Currency    string  `json:"currency"`
	Balance     float64 `json:"balance"`
	AccountType string  `json:"account_type"`
	Category    string  `json:"category"`
}

type CategoryResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Context     string `json:"context"`
	ContextType string `json:"context_type"`
}

type TransactionResponse struct {
	TransactionID   int     `json:"id"`
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	Date            int64   `json:"date"`
	Description     string  `json:"description"`
	AccountID       int     `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Category        string  `json:"category"`
}
