package ledger

import (
	"context"
	"encore.app/graphql/model"
	"encore.app/ledger/workflow"
	encore "encore.dev"
	"fmt"
	tb_types "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"
	"log"
	"strconv"
	"time"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func SerializeToModel(account *tb_types.Account) (model.Account) {
	return model.Account{
		ID:      account.ID.String(),
		Credits: strconv.FormatUint(GetBalance(account), 10),
	}
}

func GetBalance(account *tb_types.Account) uint64 {
	return account.CreditsPosted - account.DebitsPosted - account.DebitsPending
}

// Use an environment-specific task queue so we can use the same
// Temporal Cluster for all cloud environments.
var (
	envName   = encore.Meta().Environment.Name
	TaskQueue = envName + "-accounts"
)

//encore:service
type Service struct {
	client client.Client
	worker worker.Worker
}

func initService() (*Service, error) {
	c, err := client.Dial(client.Options{})
	if err != nil {
		return nil, fmt.Errorf("create temporal client: %v", err)
	}

	w := worker.New(c, TaskQueue, worker.Options{})

	w.RegisterWorkflow(workflow.AccountWorkflow)
	w.RegisterActivity(workflow.Authorize)
	w.RegisterActivity(workflow.Present)
	w.RegisterActivity(workflow.Release)

	w.RegisterWorkflow(workflow.CreateAccountWorkflow)
	w.RegisterActivity(workflow.CreateAccount)
	w.RegisterActivity(workflow.CreateBankAccount)
	w.RegisterActivity(workflow.FullFillAccount)

	err = w.Start()
	if err != nil {
		c.Close()
		return nil, fmt.Errorf("start temporal worker: %v", err)
	}

	return &Service{client: c, worker: w}, nil
}

func (s *Service) Shutdown(force context.Context) {
	s.client.Close()
	s.worker.Stop()
}

type StartAccountWorkflowParams struct {
	AccountId string
}

type AuthorizeParams struct {
	AccountId string
	Amount uint64
}

type PresentParams struct {
	AccountId string
	Amount    uint64
}

type CreateNewAccountParams struct {
	AccountId string
	Amount    uint64
}


//encore:api private
func (s *Service) StartAccountWorkflow(ctx context.Context, input *StartAccountWorkflowParams) error {
	accId, err := tb_types.HexStringToUint128(input.AccountId)

	if err != nil {
		return err
	}

	options := client.StartWorkflowOptions{
		ID:        GetWorkflowId(input.AccountId),
		TaskQueue: TaskQueue,
	}

	// Migrate duration to Encore config
	duration := time.Second * 100

	we, err := s.client.ExecuteWorkflow(ctx, options, workflow.AccountWorkflow, accId, duration)
	if err != nil {
		return err
	}

	log.Printf("Started workflow with ID %s and run ID %s", we.GetID(), we.GetRunID())

	return nil
}

//encore:api private
func (s *Service) Authorize(ctx context.Context, input *AuthorizeParams) error {
	err := s.client.SignalWorkflow(context.Background(), GetWorkflowId(input.AccountId), "", "Authorize", workflow.AuthorizeSignal{Amount: input.Amount})

	if err != nil {
		log.Printf("Authorize error for %s with Amount %s", input.AccountId, input.Amount)
		return err
	}

	return nil
}

//encore:api private
func (s *Service) Present(ctx context.Context, input *PresentParams) error {
	err := s.client.SignalWorkflow(context.Background(), GetWorkflowId(input.AccountId), "", "Present", workflow.PresentSignal{Amount: input.Amount})

	if err != nil {
		log.Printf("Present error for %s with Amount %s", input.AccountId, input.Amount)
		return err
	}

	return nil
}

//encore:api private
func (s *Service) CreateNewAccount(ctx context.Context, input *CreateNewAccountParams) error {
	options := client.StartWorkflowOptions{
		ID:        "create-new-account-workflow-" + input.AccountId,
		TaskQueue: "create-account-queue",
	}

	accId, _ := tb_types.HexStringToUint128(input.AccountId)

	we, err := s.client.ExecuteWorkflow(ctx, options, workflow.CreateAccountWorkflow, accId, input.Amount)
	if err != nil {
		return err
	}

	log.Printf("Started workflow with ID %s and run ID %s", we.GetID(), we.GetRunID())

	we.Get(ctx, nil)

	// Run as Workflow!
	return s.StartAccountWorkflow(ctx, &StartAccountWorkflowParams{AccountId: input.AccountId})

	return nil
}

func GetWorkflowId(accountId string) string {
	return "account-workflow-" + accountId
}