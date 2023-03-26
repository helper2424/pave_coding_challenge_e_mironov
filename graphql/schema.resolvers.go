package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"
	"errors"
	"fmt"

	"encore.app/graphql/generated"
	"encore.app/graphql/model"
	"encore.app/ledger"
)

// CreateAccount is the resolver for the createAccount field.
func (r *mutationResolver) CreateAccount(ctx context.Context, input model.CreateAccountInput) (*model.MutationResult, error) {
	err := ledger.CreateNewAccount(ctx, &ledger.CreateNewAccountParams{AccountId: input.ID, Amount: uint64(input.InitialAmount)})

	if err != nil {
		return &model.MutationResult{1}, err
	}

	return &model.MutationResult{Status: 0}, nil
}

// Authorize is the resolver for the authorize field.
func (r *mutationResolver) Authorize(ctx context.Context, input model.AuthorizeInput) (*model.MutationResult, error) {
	if input.Amount <= 0 {
		error_message := fmt.Sprintf("The amount should be positive")
		return &model.MutationResult{Status: 1}, errors.New(error_message)
	}

	err := ledger.Authorize(ctx, &ledger.AuthorizeParams{
		Amount:    uint64(input.Amount),
		AccountId: input.AccountID,
	})

	if err != nil {
		return &model.MutationResult{Status: 1}, err
	}

	return &model.MutationResult{Status: 0}, nil
}

// Present is the resolver for the present field.
func (r *mutationResolver) Present(ctx context.Context, input model.PresentInput) (*model.MutationResult, error) {
	if input.Amount <= 0 {
		error_message := fmt.Sprintf("The amount should be positive")
		return &model.MutationResult{Status: 1}, errors.New(error_message)
	}

	err := ledger.Present(ctx, &ledger.PresentParams{
		Amount:    uint64(input.Amount),
		AccountId: input.AccountID,
	})

	if err != nil {
		return &model.MutationResult{1}, err
	}

	return &model.MutationResult{Status: 0}, nil
}

// Accounts is the resolver for the accounts field.
func (r *queryResolver) Account(ctx context.Context, input model.AccountsInput) (*model.Account, error) {
	account, err := ledger.GetAccount(ctx, &ledger.GetAccountParams{AccountId: input.AccountID})

	if err != nil {
		return nil, err
	}

	return account, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
