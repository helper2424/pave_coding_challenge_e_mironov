package workflow

import (
	"container/list"
	tb_types "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"
	"go.temporal.io/sdk/temporal"
	"log"
	"strconv"
	"time"

	"go.temporal.io/sdk/workflow"
)

type PendingTransfer struct {
	Id    tb_types.Uint128
	Amount uint64
}

type AuthorizeSignal struct {
	Amount uint64
}

type PresentSignal struct {
	Amount uint64
}

func AccountWorkflow(ctx workflow.Context, accountId tb_types.Uint128, authDuration time.Duration) error {
	// Check that we don't try using bank account

	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	// Check that account exists

	authChannel := workflow.GetSignalChannel(ctx, "Authorize")
	presentChannel := workflow.GetSignalChannel(ctx, "Present")
	selector := workflow.NewSelector(ctx)

	pendingQueue := list.New()

	selector.AddReceive(authChannel, func(c workflow.ReceiveChannel, _ bool) {
		var signal AuthorizeSignal
		c.Receive(ctx, &signal)

		var authTransferId tb_types.Uint128
		workflow.ExecuteActivity(ctx, GetRandomId, accountId, signal.Amount).Get(ctx, &authTransferId)

		err := workflow.ExecuteActivity(ctx, Authorize, accountId, authTransferId, signal.Amount).Get(ctx, nil)

		if err != nil {
			log.Printf("Can't authorize %s with amount %d, error %s", accountId.String(), signal.Amount, err)
			return
		}

		log.Printf("Authorization transfer ID %s queue: %s", authTransferId.String(), pendingQueue)

		pendingQueueItem := pendingQueue.PushBack(PendingTransfer{
			Id:     authTransferId,
			Amount: signal.Amount,
		})

		selector.AddFuture(workflow.NewTimer(ctx, authDuration), func(f workflow.Future) {
			log.Printf("Call Release with transfer ID %s", authTransferId)
			var releaseTransferId tb_types.Uint128
			workflow.ExecuteActivity(ctx, GetRandomId, accountId, signal.Amount).Get(ctx, &releaseTransferId)

			workflow.ExecuteActivity(ctx, Release, accountId, releaseTransferId, authTransferId).Get(ctx, nil)
			pendingQueue.Remove(pendingQueueItem)
			log.Printf("%s removed from Queue %s", authTransferId.String(), pendingQueueItem)
		})

	})

	selector.AddReceive(presentChannel, func(c workflow.ReceiveChannel, _ bool) {
		var signal PresentSignal
		c.Receive(ctx, &signal)

		log.Printf("Received signal %s", signal)

		var pendingTransferId tb_types.Uint128
		var pendingQueueItem *list.Element

		log.Printf("Qeueu lookup %s", pendingQueue)

		for temp := pendingQueue.Front(); temp != nil; temp = temp.Next() {
			log.Printf("Qeueu lookup %s", temp)

			value := temp.Value.(PendingTransfer)
			if value.Amount == signal.Amount {
				pendingTransferId = value.Id
				pendingQueueItem = temp

				log.Printf("Found item in queue %s", pendingTransferId)

				break
			}
		}

		log.Printf("Call Present with %s %s %d", accountId.String(), pendingTransferId.String(), signal.Amount)

		var presentTransferId tb_types.Uint128
		workflow.ExecuteActivity(ctx, GetRandomId, accountId, signal.Amount).Get(ctx, &presentTransferId)

		workflow.ExecuteActivity(ctx, Present, accountId, presentTransferId, pendingTransferId, signal.Amount).Get(ctx, nil)

		if pendingQueueItem != nil {
			pendingQueue.Remove(pendingQueueItem)
			log.Printf("Removed item %s from queue %s", pendingQueueItem, pendingQueue)
		}
	})

	for {
		selector.Select(ctx)
	}
}


func CreateAccountWorkflow(ctx workflow.Context, accountId tb_types.Uint128, amount uint64) error {
	// Check that we not trying to create a Bank Account
	workflowOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, workflowOptions)

	err := workflow.ExecuteActivity(ctx, CreateAccount, accountId).Get(ctx, nil)

	if err != nil {
		panic("Issue with creating new account" + accountId.String())
	}

	var bankAccount tb_types.Account

	bankAccountIdUint128, nil := tb_types.HexStringToUint128(strconv.FormatUint(BankAccountId, 10))

	 ao := workflow.ActivityOptions{
		 	RetryPolicy: &temporal.RetryPolicy{
		 		MaximumAttempts: 1,
			},
			StartToCloseTimeout: 10 * time.Second,
		}
	getAccountCtx := workflow.WithActivityOptions(ctx, ao)

	err = workflow.ExecuteActivity(getAccountCtx, GetAccount, bankAccountIdUint128, false).Get(ctx, &bankAccount)

	if err != nil {
		err := workflow.ExecuteActivity(ctx, CreateBankAccount).Get(ctx, nil)

		if err != nil {
			panic("Can't create bank Account")
		}
	}

	return workflow.ExecuteActivity(ctx, FullFillAccount, accountId, amount).Get(ctx, nil)
}

func GetAccountWorkflow(ctx workflow.Context, accountId tb_types.Uint128) (tb_types.Account, error) {
	// Check that we not trying to create a Bank Account
	workflowOptions := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, workflowOptions)

	var account tb_types.Account

	ao := workflow.ActivityOptions{
		RetryPolicy: &temporal.RetryPolicy{
			MaximumAttempts: 1,
		},
		StartToCloseTimeout: 10 * time.Second,
	}
	getAccountCtx := workflow.WithActivityOptions(ctx, ao)

	err := workflow.ExecuteActivity(getAccountCtx, GetAccount, accountId).Get(ctx, &account)

	return account, err
}
