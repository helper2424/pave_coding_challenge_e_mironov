package ledger

import (
	//"encore.app/graphql/model"
	tb_types "github.com/tigerbeetledb/tigerbeetle-go/pkg/types"
	//"reflect"
	"testing"
	//"context"
)

//func TestAuthorize(t *testing.T) {
//	type args struct {
//		ctx context.Context
//		p   *AuthorizeParams
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := Authorize(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
//				t.Errorf("Authorize() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

//func TestCreateNewAccount(t *testing.T) {
//	type args struct {
//		ctx context.Context
//		p   *CreateNewAccountParams
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := CreateNewAccount(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
//				t.Errorf("CreateNewAccount() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

//func TestGetAccount(t *testing.T) {
//	type args struct {
//		ctx context.Context
//		p   *GetAccountParams
//	}
//	tests := []struct {
//		name    string
//		args    args
//		want    *model.Account
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, err := GetAccount(tt.args.ctx, tt.args.p)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetAccount() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetAccount() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestGetBalance(t *testing.T) {
	type args struct {
		account *tb_types.Account
	}

	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "with nil",
			args: args{account: nil},
			want: 0,
		},
		{
			name: "with credits only",
			args: args{account: &tb_types.Account{CreditsPosted: 100}},
			want: 100,
		},
		{
			name: "with credits and debits",
			args: args{account: &tb_types.Account{CreditsPosted: 100, DebitsPosted: 10}},
			want: 90,
		},
		{
			name: "with credits and debits",
			args: args{account: &tb_types.Account{CreditsPosted: 100, DebitsPosted: 10, DebitsPending: 10}},
			want: 80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBalance(tt.args.account); got != tt.want {
				t.Errorf("GetBalance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWorkflowId(t *testing.T) {
	type args struct {
		accountId string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Default",
			args: args{accountId: "1111"},
			want: "account-workflow-1111",
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWorkflowId(tt.args.accountId); got != tt.want {
				t.Errorf("GetWorkflowId() = %v, want %v", got, tt.want)
			}
		})
	}
}

//func TestPresent(t *testing.T) {
//	type args struct {
//		ctx Context
//		p   *PresentParams
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := Present(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
//				t.Errorf("Present() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

//func TestSerializeToModel(t *testing.T) {
//	type args struct {
//		account *tb_types.Account
//	}
//	tests := []struct {
//		name string
//		args args
//		want *model.Account
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := SerializeToModel(tt.args.account); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("SerializeToModel() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

//func TestService_Authorize(t *testing.T) {
//	type fields struct {
//		client Client
//		worker Worker
//	}
//	type args struct {
//		ctx   Context
//		input *AuthorizeParams
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Service{
//				client: tt.fields.client,
//				worker: tt.fields.worker,
//			}
//			if err := s.Authorize(tt.args.ctx, tt.args.input); (err != nil) != tt.wantErr {
//				t.Errorf("Authorize() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

//func TestService_CreateNewAccount(t *testing.T) {
//	type fields struct {
//		client Client
//		worker Worker
//	}
//	type args struct {
//		ctx   Context
//		input *CreateNewAccountParams
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Service{
//				client: tt.fields.client,
//				worker: tt.fields.worker,
//			}
//			if err := s.CreateNewAccount(tt.args.ctx, tt.args.input); (err != nil) != tt.wantErr {
//				t.Errorf("CreateNewAccount() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}

//func TestService_GetAccount(t *testing.T) {
//	type fields struct {
//		client Client
//		worker Worker
//	}
//	type args struct {
//		ctx   Context
//		input *GetAccountParams
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    *model.Account
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Service{
//				client: tt.fields.client,
//				worker: tt.fields.worker,
//			}
//			got, err := s.GetAccount(tt.args.ctx, tt.args.input)
//			if (err != nil) != tt.wantErr {
//				t.Errorf("GetAccount() error = %v, wantErr %v", err, tt.wantErr)
//				return
//			}
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetAccount() got = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestService_Present(t *testing.T) {
//	type fields struct {
//		client Client
//		worker Worker
//	}
//	type args struct {
//		ctx   Context
//		input *PresentParams
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Service{
//				client: tt.fields.client,
//				worker: tt.fields.worker,
//			}
//			if err := s.Present(tt.args.ctx, tt.args.input); (err != nil) != tt.wantErr {
//				t.Errorf("Present() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestService_StartAccountWorkflow(t *testing.T) {
//	type fields struct {
//		client Client
//		worker Worker
//	}
//	type args struct {
//		ctx   Context
//		input *StartAccountWorkflowParams
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			s := &Service{
//				client: tt.fields.client,
//				worker: tt.fields.worker,
//			}
//			if err := s.StartAccountWorkflow(tt.args.ctx, tt.args.input); (err != nil) != tt.wantErr {
//				t.Errorf("StartAccountWorkflow() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
//
//func TestStartAccountWorkflow(t *testing.T) {
//	type args struct {
//		ctx Context
//		p   *StartAccountWorkflowParams
//	}
//	tests := []struct {
//		name    string
//		args    args
//		wantErr bool
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if err := StartAccountWorkflow(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
//				t.Errorf("StartAccountWorkflow() error = %v, wantErr %v", err, tt.wantErr)
//			}
//		})
//	}
//}
