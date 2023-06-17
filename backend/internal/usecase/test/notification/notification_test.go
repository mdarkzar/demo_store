package template_test

import (
	"store/bimport"
	"store/internal/entity/global"
	"store/internal/entity/notification"
	"store/internal/entity/user"
	"store/internal/rimport"
	"store/internal/transaction"
	"store/tools/logger"
	"store/uimport"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	testLogger = logger.NewNoFileLogger("test")
)

func TestSendUser(t *testing.T) {
	r := assert.New(t)

	type fields struct {
		ri rimport.TestRepositoryImports
		bi *bimport.TestBridgeImports
		ts *transaction.MockSession
	}
	type args struct {
		UserID  int
		Title   string
		Message string
	}

	const (
		id = 1
	)

	var arg1 args
	faker.FakeData(&arg1)

	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		err     error
	}{
		{
			name: "успешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.Notification.EXPECT().Create(f.ts, arg1.Title, arg1.Message).Return(id, nil),
					f.ri.MockRepository.Notification.EXPECT().CreateUserMessage(f.ts, arg1.UserID, id).Return(nil),
				)
			},
			args: arg1,
			err:  nil,
		},
		{
			name: "успешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.Notification.EXPECT().Create(f.ts, arg1.Title, arg1.Message).Return(id, global.ErrDBUnvailable),
				)
			},
			args: arg1,
			err:  global.ErrInternalError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				ri: rimport.NewTestRepositoryImports(ctrl),
				ts: transaction.NewMockSession(ctrl),
				bi: bimport.NewTestBridgeImports(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			ui := uimport.NewUsecaseImports(testLogger, f.ri.RepositoryImports(), f.bi.BridgeImports())

			err := ui.Usecase.Notification.SendUser(f.ts, tt.args.UserID, tt.args.Title, tt.args.Message)
			r.Equal(tt.err, err)

		})
	}
}

func TestSendAll(t *testing.T) {
	r := assert.New(t)

	type fields struct {
		ri rimport.TestRepositoryImports
		bi *bimport.TestBridgeImports
		ts *transaction.MockSession
	}
	type args struct {
		Title   string
		Message string
	}

	const (
		id = 1
	)

	var arg1 args
	faker.FakeData(&arg1)

	userList := []user.User{
		{ID: 1},
		{ID: 2},
	}

	tests := []struct {
		name    string
		prepare func(f *fields)
		args    args
		err     error
	}{
		{
			name: "успешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.User.EXPECT().LoadAll(f.ts).Return(userList, nil),
					f.ri.MockRepository.Notification.EXPECT().Create(f.ts, arg1.Title, arg1.Message).Return(id, nil),
					f.ri.MockRepository.Notification.EXPECT().CreateUserMessage(f.ts, userList[0].ID, id).Return(nil),
					f.ri.MockRepository.Notification.EXPECT().CreateUserMessage(f.ts, userList[1].ID, id).Return(nil),
				)
			},
			args: arg1,
			err:  nil,
		},
		{
			name: "1 неуспешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.User.EXPECT().LoadAll(f.ts).Return(userList, nil),
					f.ri.MockRepository.Notification.EXPECT().Create(f.ts, arg1.Title, arg1.Message).Return(id, nil),
					f.ri.MockRepository.Notification.EXPECT().CreateUserMessage(f.ts, userList[0].ID, id).Return(global.ErrDBUnvailable),
					f.ri.MockRepository.Notification.EXPECT().CreateUserMessage(f.ts, userList[1].ID, id).Return(nil),
				)
			},
			args: arg1,
			err:  nil,
		},
		{
			name: "нет получаетелей",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.User.EXPECT().LoadAll(f.ts).Return([]user.User{}, global.ErrNoData),
				)
			},
			args: arg1,
			err:  global.ErrInternalError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				ri: rimport.NewTestRepositoryImports(ctrl),
				ts: transaction.NewMockSession(ctrl),
				bi: bimport.NewTestBridgeImports(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			ui := uimport.NewUsecaseImports(testLogger, f.ri.RepositoryImports(), f.bi.BridgeImports())

			err := ui.Usecase.Notification.SendAll(f.ts, tt.args.Title, tt.args.Message)
			r.Equal(tt.err, err)

		})
	}
}

func TestLoadUserMessages(t *testing.T) {
	r := assert.New(t)

	type fields struct {
		ri rimport.TestRepositoryImports
		bi *bimport.TestBridgeImports
		ts *transaction.MockSession
	}
	type args struct {
		UserID  int
		Title   string
		Message string
	}

	var arg1 args
	faker.FakeData(&arg1)

	messageList := []notification.Notification{
		{ID: 1},
		{ID: 2},
	}

	tests := []struct {
		name     string
		prepare  func(f *fields)
		args     args
		err      error
		expected []notification.Notification
	}{
		{
			name: "успешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.Notification.EXPECT().FindUserMessages(f.ts, arg1.UserID).Return(messageList, nil),
					f.ri.MockRepository.Notification.EXPECT().Delete(f.ts, messageList[0].ID, arg1.UserID).Return(nil),
					f.ri.MockRepository.Notification.EXPECT().Delete(f.ts, messageList[1].ID, arg1.UserID).Return(nil),
				)
			},
			args:     arg1,
			expected: messageList,
			err:      nil,
		},
		{
			name: "частичный успешный результат результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.Notification.EXPECT().FindUserMessages(f.ts, arg1.UserID).Return(messageList, nil),
					f.ri.MockRepository.Notification.EXPECT().Delete(f.ts, messageList[0].ID, arg1.UserID).Return(global.ErrDBUnvailable),
					f.ri.MockRepository.Notification.EXPECT().Delete(f.ts, messageList[1].ID, arg1.UserID).Return(nil),
				)
			},
			args:     arg1,
			err:      nil,
			expected: messageList,
		},
		{
			name: "нет сообщений",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.Notification.EXPECT().FindUserMessages(f.ts, arg1.UserID).Return(nil, global.ErrNoData),
				)
			},
			args:     arg1,
			err:      nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				ri: rimport.NewTestRepositoryImports(ctrl),
				ts: transaction.NewMockSession(ctrl),
				bi: bimport.NewTestBridgeImports(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			ui := uimport.NewUsecaseImports(testLogger, f.ri.RepositoryImports(), f.bi.BridgeImports())

			messages, err := ui.Usecase.Notification.LoadUserMessages(f.ts, tt.args.UserID)
			r.Equal(tt.err, err)
			r.Equal(tt.expected, messages)

		})
	}
}
