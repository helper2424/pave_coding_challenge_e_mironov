// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Account struct {
	ID      string `json:"id"`
	Credits string `json:"credits"`
}

type AccountsInput struct {
	AccountID string `json:"accountId"`
}

type AuthorizeInput struct {
	AccountID string `json:"accountId"`
	Amount    int    `json:"amount"`
}

type CreateAccountInput struct {
	ID            string `json:"id"`
	InitialAmount int    `json:"initialAmount"`
}

type MutationResult struct {
	Status int `json:"status"`
}

type PresentInput struct {
	AccountID string `json:"accountId"`
	Amount    int    `json:"amount"`
}
