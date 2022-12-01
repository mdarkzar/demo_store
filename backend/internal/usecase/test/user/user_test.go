package template_test

import (
	"store/bimport"
	"store/internal/entity/global"
	"store/internal/entity/jwt"
	"store/internal/entity/user"
	"store/internal/rimport"
	"store/internal/transaction"
	"store/tools/logger"
	"store/tools/passfunc"
	"store/uimport"
	"testing"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	testLogger = logger.NewNoFileLogger("test")
)

func TestAuth(t *testing.T) {
	r := assert.New(t)

	type fields struct {
		ri rimport.TestRepositoryImports
		bi *bimport.TestBridgeImports
		ts *transaction.MockSession
	}
	type args struct {
		login    string
		password string
	}

	const (
		userID = 1
	)

	var (
		login    string
		password string
	)

	faker.FakeData(&login)
	faker.FakeData(&password)

	passHash, err := passfunc.BcryptCreatePassword(password)
	r.NoError(err)

	userData := user.User{
		ID:       userID,
		Login:    login,
		Password: passHash,
	}

	now := time.Now()

	jwtToken, err := jwt.NewJwtToken(userData.ID, userData.Login, now)
	r.NoError(err)

	tests := []struct {
		name         string
		prepare      func(f *fields)
		args         args
		expectedData string
		err          error
	}{
		{
			name: "успешный результат: авторегистрация",
			prepare: func(f *fields) {

				gomock.InOrder(
					f.ri.MockRepository.User.EXPECT().FindByLogin(f.ts, login).Return(user.User{}, global.ErrNoData),
					f.ri.MockRepository.User.EXPECT().Create(f.ts, login, gomock.Any()).Return(userID, nil),
					f.ri.MockRepository.User.EXPECT().FindByID(f.ts, userID).Return(userData, nil),
					f.bi.TestBridge.Date.EXPECT().Now().Return(now),
				)
			},
			args:         args{login: login, password: password},
			expectedData: jwtToken,
			err:          nil,
		},
		{
			name: "успешный результат: существует",
			prepare: func(f *fields) {

				gomock.InOrder(
					f.ri.MockRepository.User.EXPECT().FindByLogin(f.ts, login).Return(userData, nil),
					f.bi.TestBridge.Date.EXPECT().Now().Return(now),
				)
			},
			args:         args{login: login, password: password},
			expectedData: jwtToken,
			err:          nil,
		},
		{
			name: "неуспешный результат: неверный пароль",
			prepare: func(f *fields) {

				userData2 := userData
				userData2.Password = "wrongpassword"
				gomock.InOrder(
					f.ri.MockRepository.User.EXPECT().FindByLogin(f.ts, login).Return(userData2, nil),
				)
			},
			args:         args{login: login, password: password},
			expectedData: "",
			err:          user.ErrLoginOrPasswordIncorrect,
		},
		{
			name: "неуспешный результат: внутреняя ошибка",
			prepare: func(f *fields) {

				gomock.InOrder(
					f.ri.MockRepository.User.EXPECT().FindByLogin(f.ts, login).Return(userData, global.ErrDBUnvailable),
				)
			},
			args:         args{login: login, password: password},
			expectedData: "",
			err:          global.ErrInternalError,
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

			sm := transaction.NewMockSessionManager(ctrl)
			ui := uimport.NewUsecaseImports(testLogger, f.ri.RepositoryImports(), f.bi.BridgeImports(), sm)

			data, err := ui.Usecase.User.Auth(f.ts, tt.args.login, tt.args.password)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)

		})
	}
}

func TestFindUser(t *testing.T) {
	r := assert.New(t)

	type fields struct {
		ri rimport.TestRepositoryImports
		bi *bimport.TestBridgeImports
		ts *transaction.MockSession
	}
	type args struct {
		UserID int
	}

	var arg1 args
	faker.FakeData(&arg1)

	var user1 user.User
	faker.FakeData(&user1)

	tests := []struct {
		name         string
		prepare      func(f *fields)
		args         args
		expectedData user.User
		err          error
	}{
		{
			name: "успешный результат: есть в кэше",
			prepare: func(f *fields) {

				gomock.InOrder(
					f.ri.MockRepository.UserCache.EXPECT().Get(arg1.UserID).Return(user1, true),
				)
			},
			args:         arg1,
			expectedData: user1,
			err:          nil,
		},
		{
			name: "успешный результат: нет в кэше, есть в бд",
			prepare: func(f *fields) {

				gomock.InOrder(
					f.ri.MockRepository.UserCache.EXPECT().Get(arg1.UserID).Return(user.User{}, false),
					f.ri.SessionManager.EXPECT().CreateSession().Return(f.ts),
					f.ts.EXPECT().Start().Return(nil),
					f.ri.MockRepository.User.EXPECT().FindByID(f.ts, arg1.UserID).Return(user1, nil),
					f.ri.MockRepository.UserCache.EXPECT().Add(arg1.UserID, user1),
					f.ts.EXPECT().Rollback().Return(nil),
				)
			},
			args:         arg1,
			expectedData: user1,
			err:          nil,
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

			sm := transaction.NewMockSessionManager(ctrl)
			ui := uimport.NewUsecaseImports(testLogger, f.ri.RepositoryImports(), f.bi.BridgeImports(), sm)

			data, err := ui.Usecase.User.FindUser(tt.args.UserID)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)

		})
	}
}
