package db
//
//import (
//	"context"
//	"errors"
//	"fmt"
//
//	//"encoding/binary"
//	tb "github.com/tigerbeetledb/tigerbeetle-go"
//	tb_types "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"
//	"go4.org/syncutil"
//	"log"
//	//"math/rand"
//	"strconv"
//)
//
//const LedgerID = 1
//const BankAccountId = 1
//const CurrencyCodeID = 718
//
//var (
//	once syncutil.Once
//	lastId uint64 = 100
//)
//
////encore:service
//type Service struct {
//	client tb.Client
//}
//
//func initService() (*Service, error) {
//	// TODO: Move port to config
//	client, err := tb.NewClient(0, []string{"3000"}, 1)
//
//	if err != nil {
//		log.Println("Can't create TigerBeetle client created", err)
//		return nil, err
//	}
//
//	return &Service{client: client}, nil
//}
//
//func (s *Service) Shutdown(force context.Context) {
//	s.client.Close()
//}
//
//
//
//
//func uint128(value string) tb_types.Uint128 {
//	x, err := tb_types.HexStringToUint128(value)
//	if err != nil {
//		panic(err)
//	}
//	return x
//}
//
//// Tigerbeetle tema recommends using random IDS, so let's do it.
//// In ideal case it should be 128 bit BigInteger random ID
//// but this apporach required additional effort for implementation.
//// As this is a toy ledger and it shouldn't have more than ~4 billions users
//// We can use uint32 random generator
//func GetRandomId() tb_types.Uint128 {
//	//var b [16]byte
//	//for i := 0; i < 16; i++ {
//	//	b[i] = 0
//	//}
//	//
//	//binary.BigEndian.PutUint64(b[0:8], rand.Uint64())
//	//return tb_types.BytesToUint128(b)
//
//	//value, err := tb_types.HexStringToUint128(strconv.FormatUint(rand.Uint64(), 10))
//
//	value, err := tb_types.HexStringToUint128(strconv.FormatUint(lastId, 10))
//	lastId += 1
//
//	if err != nil {
//		log.Println("Can't generate ID with error %s", err)
//	}
//
//	return value
//}
//
////encore:api private
//func (s *Service) CreateAccount(ctx context.Context, id tb_types.Uint128) ([]tb_types.AccountEventResult, error) {
//	return s.client.CreateAccounts([]tb_types.Account{
//		{
//			ID:       id,
//			UserData: tb_types.Uint128{},
//			Reserved: [48]uint8{},
//			Ledger:   LedgerID,
//			Code:     CurrencyCodeID,
//			Flags:    tb_types.AccountFlags{DebitsMustNotExceedCredits: true}.ToUint16(),
//		},
//	})
//}
//
////encore:api private
//func (s *Service) CreateBankAccount(ctx context.Context) error {
//	results, err := s.client.CreateAccounts([]tb_types.Account{
//		{
//			ID:       GetBankAccountId(),
//			UserData: tb_types.Uint128{},
//			Reserved: [48]uint8{},
//			Ledger:   LedgerID,
//			Code:     CurrencyCodeID,
//			Flags:    tb_types.AccountFlags{CreditsMustNotExceedDebits: true}.ToUint16(),
//		},
//	})
//
//	log.Printf("Creating bank account %s %s", result, err)
//
//	if len(results) == 0 {
//		panic("Bank Account creation result is empty")
//	}
//
//	result := results[0]
//
//	switch result.Result {
//	case tb_types.AccountOK:
//		log.Println("Bank Account was succesfully created with")
//	case tb_types.AccountExists:
//		log.Println("Bank Account already exists with id")
//	default:
//		error_message := fmt.Sprintf("Error creating Bank Account %d: %s", result.Index, result.Result)
//		log.Printf(error_message)
//
//		return errors.New(error_message)
//	}
//	return nil
//}
//
////encore:api private
//func (s *Service) GetAccounts(ctx context.Context, ids []tb_types.Uint128) error {
//
//}
//
//func GetBalance(account *tb_types.Account) uint64 {
//	return account.CreditsPosted - account.DebitsPosted - account.DebitsPending
//}