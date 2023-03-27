package workflow

import (
	"context"
	tb "github.com/tigerbeetledb/tigerbeetle-go"
	"github.com/tigerbeetledb/tigerbeetle-go/pkg/types"
	"go.temporal.io/sdk/workflow"
	"reflect"
	"testing"
	"time"
)

func TestAccountWorkflow(t *testing.T) {
	type args struct {
		ctx          workflow.Context
		accountId    types.Uint128
		authDuration time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AccountWorkflow(tt.args.ctx, tt.args.accountId, tt.args.authDuration); (err != nil) != tt.wantErr {
				t.Errorf("AccountWorkflow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAuthorize(t *testing.T) {
	type args struct {
		ctx        context.Context
		accountId  types.Uint128
		transferId types.Uint128
		amount     uint64
	}
	tests := []struct {
		name    string
		args    args
		want    types.Uint128
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Authorize(tt.args.ctx, tt.args.accountId, tt.args.transferId, tt.args.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("Authorize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Authorize() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateAccount(t *testing.T) {
	type args struct {
		ctx       context.Context
		accountId types.Uint128
		isBank    bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateAccount(tt.args.ctx, tt.args.accountId, tt.args.isBank); (err != nil) != tt.wantErr {
				t.Errorf("CreateAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateAccountWorkflow(t *testing.T) {
	type args struct {
		ctx       workflow.Context
		accountId types.Uint128
		amount    uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateAccountWorkflow(tt.args.ctx, tt.args.accountId, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("CreateAccountWorkflow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateBankAccount(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateBankAccount(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("CreateBankAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFullFillAccount(t *testing.T) {
	type args struct {
		ctx       context.Context
		accountId types.Uint128
		amount    uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := FullFillAccount(tt.args.ctx, tt.args.accountId, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("FullFillAccount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetAccount(t *testing.T) {
	type args struct {
		ctx context.Context
		id  types.Uint128
	}
	tests := []struct {
		name    string
		args    args
		want    types.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAccount(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetAccountWorkflow(t *testing.T) {
	type args struct {
		ctx       workflow.Context
		accountId types.Uint128
	}
	tests := []struct {
		name    string
		args    args
		want    types.Account
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAccountWorkflow(tt.args.ctx, tt.args.accountId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountWorkflow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountWorkflow() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBankAccountId(t *testing.T) {
	tests := []struct {
		name string
		want types.Uint128
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBankAccountId(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBankAccountId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRandomId(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		want    types.Uint128
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRandomId(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRandomId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandomId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPresent(t *testing.T) {
	type args struct {
		ctx               context.Context
		accountId         types.Uint128
		transferId        types.Uint128
		pendingTransferId types.Uint128
		amount            uint64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Present(tt.args.ctx, tt.args.accountId, tt.args.transferId, tt.args.pendingTransferId, tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Present() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRelease(t *testing.T) {
	type args struct {
		ctx               context.Context
		accountId         types.Uint128
		transferId        types.Uint128
		pendingTransferId types.Uint128
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Release(tt.args.ctx, tt.args.accountId, tt.args.transferId, tt.args.pendingTransferId); (err != nil) != tt.wantErr {
				t.Errorf("Release() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getDBClient(t *testing.T) {
	tests := []struct {
		name string
		want *tb.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDBClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDBClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
