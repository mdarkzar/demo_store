package template_test

import (
	"store/bimport"
	"store/internal/entity/global"
	"store/internal/entity/notification"
	"store/internal/entity/product"
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

func TestCreate(t *testing.T) {
	r := assert.New(t)

	type fields struct {
		ri rimport.TestRepositoryImports
		bi *bimport.TestBridgeImports
		ts *transaction.MockSession
	}
	type args struct {
		UserID int
		Name   string
		Price  float64
		StID   int
	}

	const (
		id = 1
	)

	var arg1 args
	faker.FakeData(&arg1)

	var userData user.User
	faker.FakeData(&userData)

	tests := []struct {
		name         string
		prepare      func(f *fields)
		args         args
		expectedData int
		err          error
	}{
		{
			name: "успешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.User.EXPECT().FindByID(f.ts, arg1.UserID).Return(userData, nil),
					f.ri.MockRepository.Product.EXPECT().Create(f.ts, arg1.UserID, arg1.Name, arg1.Price, arg1.StID).Return(id, nil),
					f.bi.TestBridge.Notification.EXPECT().SendAll(f.ts, notification.CreateProductTitle, notification.CreateProductBody(userData.Login, arg1.Name, arg1.Price)).Return(nil),
				)
			},
			expectedData: id,
			args:         arg1,
			err:          nil,
		},
		{
			name: "успешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.User.EXPECT().FindByID(f.ts, arg1.UserID).Return(userData, nil),
					f.ri.MockRepository.Product.EXPECT().Create(f.ts, arg1.UserID, arg1.Name, arg1.Price, arg1.StID).Return(id, global.ErrDBUnvailable),
				)
			},
			expectedData: id,
			args:         arg1,
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

			ui := uimport.NewUsecaseImports(testLogger, f.ri.RepositoryImports(), f.bi.BridgeImports())

			data, err := ui.Usecase.Product.Create(f.ts, tt.args.UserID, tt.args.Name, tt.args.Price, tt.args.StID)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, data)

		})
	}
}

func TestRemove(t *testing.T) {
	r := assert.New(t)

	type fields struct {
		ri rimport.TestRepositoryImports
		bi *bimport.TestBridgeImports
		ts *transaction.MockSession
	}
	type args struct {
		ProductID int
		UserID    int
	}

	var arg1 args
	faker.FakeData(&arg1)

	var result product.Product
	faker.FakeData(&result)
	result.ID = arg1.ProductID

	var userData user.User
	faker.FakeData(&userData)

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
					f.ri.MockRepository.Product.EXPECT().FindByID(f.ts, arg1.ProductID).Return(result, nil),
					f.ri.MockRepository.User.EXPECT().FindByID(f.ts, arg1.UserID).Return(userData, nil),
					f.ri.MockRepository.Product.EXPECT().Remove(f.ts, arg1.ProductID).Return(nil),
					f.bi.TestBridge.Notification.EXPECT().SendAll(f.ts, notification.DeleteProductTitle, notification.DeleteProductBody(arg1.ProductID, userData.Login, result.Name, result.Price)).Return(nil),
				)
			},
			args: arg1,
			err:  nil,
		},
		{
			name: "неуспешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.Product.EXPECT().FindByID(f.ts, arg1.ProductID).Return(result, nil),
					f.ri.MockRepository.User.EXPECT().FindByID(f.ts, arg1.UserID).Return(userData, nil),
					f.ri.MockRepository.Product.EXPECT().Remove(f.ts, arg1.ProductID).Return(global.ErrDBUnvailable),
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

			err := ui.Usecase.Product.Remove(f.ts, tt.args.UserID, tt.args.ProductID)
			r.Equal(tt.err, err)

		})
	}
}

func TestFindByID(t *testing.T) {
	r := assert.New(t)

	type fields struct {
		ri rimport.TestRepositoryImports
		bi *bimport.TestBridgeImports
		ts *transaction.MockSession
	}
	type args struct {
		ProductID int
	}

	var arg1 args
	faker.FakeData(&arg1)

	var result product.Product
	faker.FakeData(&result)

	tests := []struct {
		name         string
		prepare      func(f *fields)
		args         args
		expectedData product.Product
		err          error
	}{
		{
			name: "успешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.Product.EXPECT().FindByID(f.ts, arg1.ProductID).Return(result, nil),
				)
			},
			args:         arg1,
			expectedData: result,
			err:          nil,
		},
		{
			name: "успешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.Product.EXPECT().FindByID(f.ts, arg1.ProductID).Return(product.Product{}, global.ErrDBUnvailable),
				)
			},
			args:         arg1,
			expectedData: product.Product{},
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

			ui := uimport.NewUsecaseImports(testLogger, f.ri.RepositoryImports(), f.bi.BridgeImports())

			result, err := ui.Usecase.Product.FindByID(f.ts, tt.args.ProductID)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, result)

		})
	}
}

func TestLoadAll(t *testing.T) {
	r := assert.New(t)

	type fields struct {
		ri rimport.TestRepositoryImports
		bi *bimport.TestBridgeImports
		ts *transaction.MockSession
	}
	type args struct {
	}

	var result []product.Product
	faker.FakeData(&result)

	tests := []struct {
		name         string
		prepare      func(f *fields)
		args         args
		expectedData []product.Product
		err          error
	}{
		{
			name: "успешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.Product.EXPECT().LoadAll(f.ts).Return(result, nil),
				)
			},
			expectedData: result,
			err:          nil,
		},
		{
			name: "успешный результат",
			prepare: func(f *fields) {
				gomock.InOrder(
					f.ri.MockRepository.Product.EXPECT().LoadAll(f.ts).Return(nil, global.ErrDBUnvailable),
				)
			},
			expectedData: nil,
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

			ui := uimport.NewUsecaseImports(testLogger, f.ri.RepositoryImports(), f.bi.BridgeImports())

			result, err := ui.Usecase.Product.LoadAll(f.ts)
			r.Equal(tt.err, err)
			r.Equal(tt.expectedData, result)

		})
	}
}
