package workflow

import (
	"context"
	"errors"
	"fmt"
	tb "github.com/tigerbeetledb/tigerbeetle-go"
	tb_types "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"
	"log"
	"math/rand"
	"strconv"
	"time"
)

const LedgerID = 1
const BankAccountId = 1
const CurrencyCodeID = 718

func getDBClient() *tb.Client {
	client, err := tb.NewClient(0, []string{"3000"}, 1)

	if err != nil {
		panic("Can't create TigerBeetle client created" + err.Error())
		return nil
	}

	return &client
}

func GetBankAccountId() tb_types.Uint128 {
	strId := strconv.FormatInt(BankAccountId, 10)

	id, err := tb_types.HexStringToUint128(strId)

	if err != nil {
		panic("Can't get Bank Account ID" + strId)
	}

	return id
}

func GetRandomId() tb_types.Uint128 {
	curTime := time.Now().UnixNano()

	leftPart := []byte(strconv.FormatInt(curTime, 10))
	rightPart := []byte(strconv.FormatUint(rand.Uint64(), 10))

	byteSlice := append(leftPart, rightPart...)

	var bytesArray [16]byte

	copy(bytesArray[:], byteSlice[:16])

	id := tb_types.BytesToUint128(bytesArray)

	return id
}

func Authorize(ctx context.Context, accountId tb_types.Uint128, amount uint64) (tb_types.Uint128, error) {
	clientRef := getDBClient()
	client := *clientRef

	batch := make([]tb_types.Transfer, 1)

	transferId := GetRandomId()

	batch[0] = tb_types.Transfer{
		ID:              transferId,
		DebitAccountID:  accountId,
		CreditAccountID: GetBankAccountId(),
		Ledger:          LedgerID,
		Code:            CurrencyCodeID,
		Amount:          amount,
		Flags:           tb_types.TransferFlags{Pending: true}.ToUint16(),
	}

	_, err := client.CreateTransfers(batch)

	if err != nil {
		log.Printf("Can't authorize %s for Account %s", amount, accountId)
		return tb_types.Uint128{}, err
	}

	// hanlde transfer result

	return transferId, nil
}

func Present(ctx context.Context, accountId tb_types.Uint128, transferId *tb_types.Uint128, amount uint64) error {
	clientRef := getDBClient()
	client := *clientRef

	var transfers []tb_types.Transfer

	if transferId != nil {
		foundTransfers, err := client.LookupTransfers([]tb_types.Uint128{*transferId})

		if err != nil {
			log.Printf("Can't find any trasnfers for %s with error %s", accountId, err)
			return err
		}

		transfers = foundTransfers
	}

	batch := make([]tb_types.Transfer, 1)

	if len(transfers) == 0 {
		// No transfers that fit our original one, let's try to use money from the account
		batch[0] = tb_types.Transfer{
			ID:              GetRandomId(),
			DebitAccountID:  accountId,
			CreditAccountID: GetBankAccountId(),
			Ledger:          LedgerID,
			Code:            CurrencyCodeID,
			Amount:          amount,
		}
	} else {
		// Try to make a PostPending transaction
		batch[0] = tb_types.Transfer{
			ID:              GetRandomId(),
			DebitAccountID:  accountId,
			CreditAccountID: GetBankAccountId(),
			Ledger:          LedgerID,
			Code:            CurrencyCodeID,
			Amount:          amount,
			Flags:           tb_types.TransferFlags{PostPendingTransfer: true}.ToUint16(),
			PendingID:       *transferId,
		}
	}

	res, err := client.CreateTransfers(batch)

	if err != nil {
		log.Printf("Can't make presentmanet %s %s %s", accountId, *transferId, amount)
		return err
	}

	if len(res) == 0 {
		log.Printf("Presentment result is empty %s %s %s", accountId, *transferId, amount)
		return errors.New("Presentment result is empty")
	}

	switch res[0].Result {
	case tb_types.TransferOK:
		return nil
	default:
		error_mesage := fmt.Sprintf("Presentment result issue %s %s", res[0].Result)
		return errors.New(error_mesage)
	}
}

func Release(ctx context.Context, accountId tb_types.Uint128, transferId tb_types.Uint128) error {
	clientRef := getDBClient()
	client := *clientRef

	batch := make([]tb_types.Transfer, 1)

	batch[0] = tb_types.Transfer{
		ID:              GetRandomId(),
		DebitAccountID:  accountId,
		CreditAccountID: GetBankAccountId(),
		Ledger:          LedgerID,
		Code:            CurrencyCodeID,
		Flags:           tb_types.TransferFlags{VoidPendingTransfer: true}.ToUint16(),
		PendingID: 		 transferId,
	}

	res, err := client.CreateTransfers(batch)

	if err != nil {
		log.Printf("Can't make Release %s %s %s", accountId, transferId)
		return err
	}

	if len(res) == 0 {
		log.Printf("Release result is empty %s %s %s", accountId, transferId)
		return errors.New("Release result is empty")
	}

	switch res[0].Result {
	case tb_types.TransferOK:
		return nil
	default:
		error_mesage := fmt.Sprintf("Release result issue %s %s", res[0].Result)
		return errors.New(error_mesage)
	}
}

func CreateAccount(ctx context.Context, accountId tb_types.Uint128) error {
	clientRef := getDBClient()
	client := *clientRef

	accountRes, err := client.CreateAccounts([]tb_types.Account{
		{
			ID:       accountId,
			UserData: tb_types.Uint128{},
			Reserved: [48]uint8{},
			Ledger:   LedgerID,
			Code:     CurrencyCodeID,
			Flags:    tb_types.AccountFlags{DebitsMustNotExceedCredits: true}.ToUint16(),
		},
	})

	if err != nil {
		log.Printf("Error creating accounts: %s", err)
		return err
	}

	if len(accountRes) == 0 {
		error_message := fmt.Sprintf("Error creating accounts, the result array is empty %s %s", accountRes, accountId)
		log.Printf(error_message)
		return errors.New(error_message)
	}

	result := accountRes[0]

	switch result.Result {
	case tb_types.AccountOK:
		log.Println("Account succesfully create with id %s", accountId)
	case tb_types.AccountExists:
		log.Println("Account already exists with id", accountId)
	default:
		error_message := fmt.Sprintf("Error creating account %d: %s", result.Index, result.Result)
		log.Printf(error_message)

		return errors.New(error_message)
	}

	return nil
}

func CreateBankAccount(ctx context.Context) error {
	value, _ := tb_types.HexStringToUint128(strconv.FormatUint(BankAccountId, 10))
	return CreateAccount(ctx, value)
}

func GetAccount(ctx context.Context, id tb_types.Uint128) (tb_types.Account, error) {
	clientRef := getDBClient()
	client := *clientRef

	accounts, err := client.LookupAccounts([]tb_types.Uint128{id})

	if err != nil {
		error_message := fmt.Sprintf("Can' take accoutns for ids", id)
		return tb_types.Account{}, errors.New(error_message)
	}

	if len(accounts) == 0 {
		return tb_types.Account{}, errors.New("No accounts")
	}

	return accounts[0], nil
}

func FullFillAccount(ctx context.Context, accountId tb_types.Uint128, amount uint64) error {
	clientRef := getDBClient()
	client := *clientRef

	batch := make([]tb_types.Transfer, 1)

	batch[0] = tb_types.Transfer{
		ID:              GetRandomId(),
		DebitAccountID:  GetBankAccountId(),
		CreditAccountID: accountId,
		Ledger:          LedgerID,
		Code:            CurrencyCodeID,
		Amount:          amount,
	}

	res, err := client.CreateTransfers(batch)

	log.Printf("Transfer basic credits %s", res)

	if err != nil {
		log.Printf("Can't setup default balance %s for %s; %s", amount, accountId.String(), res)
		return err
	}

	return nil
}


