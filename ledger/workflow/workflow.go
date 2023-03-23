package workflow

import (
	"container/list"
	tb_types "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"
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
	authChannel := workflow.GetSignalChannel(ctx, "Authorize")
	presentChannel := workflow.GetSignalChannel(ctx, "Present")
	selector := workflow.NewSelector(ctx)

	pendingQueue := list.New()

	selector.AddReceive(authChannel, func(c workflow.ReceiveChannel, _ bool) {
		var signal AuthorizeSignal
		c.Receive(ctx, &signal)

		var transferId tb_types.Uint128
		err := workflow.ExecuteActivity(ctx, Authorize, accountId, signal.Amount).Get(ctx, &transferId)

		if err != nil {
			log.Printf("Can't authorize %s with amount %s", accountId, signal.Amount)
			return
		}

		pendingQueueItem := pendingQueue.PushBack(PendingTransfer{
			Id: transferId,
			Amount: signal.Amount,
		})

		selector.AddFuture(workflow.NewTimer(ctx, authDuration), func(f workflow.Future) {
			workflow.ExecuteActivity(ctx, Release, transferId)
			pendingQueue.Remove(pendingQueueItem)
		})

	})

	selector.AddReceive(presentChannel, func(c workflow.ReceiveChannel, _ bool) {
		var signal PresentSignal
		c.Receive(ctx, &signal)

		var pendingTransferId *tb_types.Uint128 = nil

		for temp := pendingQueue.Front(); temp != nil; temp = temp.Next() {
			value := temp.Value.(PendingTransfer)
			if value.Amount == signal.Amount {
				*pendingTransferId = value.Id
			}
		}

		workflow.ExecuteActivity(ctx, Present, pendingTransferId, signal.Amount)

	})

	selector.Select(ctx)

	return nil
}


func CreateAccountWorkflow(ctx workflow.Context, accountId tb_types.Uint128, amount uint64) error {
	err := workflow.ExecuteActivity(ctx, CreateAccount, accountId).Get(ctx, nil)

	if err != nil {
		panic("Issue with creating new account" + accountId.String())
	}

	var bankAccount tb_types.Account

	bankAccountIdUint128, nil := tb_types.HexStringToUint128(strconv.FormatUint(BankAccountId, 10))
	err = workflow.ExecuteActivity(ctx, GetAccount, bankAccountIdUint128).Get(ctx, &bankAccount)

	if err != nil {
		err := workflow.ExecuteActivity(ctx, CreateBankAccount).Get(ctx, nil)

		if err != nil {
			panic("Can't create bank Account")
			return err
		}
	}

	return workflow.ExecuteActivity(ctx, FullFillAccount, accountId, amount).Get(ctx, nil)
}