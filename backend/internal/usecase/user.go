package usecase

import (
	"fmt"
	"store/bimport"
	"store/internal/entity/global"
	"store/internal/entity/jwt"
	"store/internal/entity/user"
	"store/internal/rimport"
	"store/internal/transaction"
	"store/tools/passfunc"

	"github.com/sirupsen/logrus"
)

type UserUsecase struct {
	log *logrus.Logger

	rimport.RepositoryImports
	*bimport.BridgeImports
}

func NewUserUsecase(
	log *logrus.Logger,
	ri rimport.RepositoryImports,
	b *bimport.BridgeImports,
) *UserUsecase {
	return &UserUsecase{
		log:               log,
		RepositoryImports: ri,
		BridgeImports:     b,
	}
}

func (u *UserUsecase) logPrefix() string {
	return "[user_usecase]"
}

func (u *UserUsecase) Auth(ts transaction.Session, login, password string) (jwtToken string, err error) {
	lf := logrus.Fields{"login": login}

	userData, err := u.Repository.User.FindByLogin(ts, login)
	switch err {
	case global.ErrNoData:
		// сделать авторегистрацию

		passHash, err := passfunc.BcryptCreatePassword(password)
		if err != nil {
			u.log.WithFields(lf).Errorln(u.logPrefix(),
				fmt.Sprintf("не удалось создать пароль; ошибка: %v", err),
			)
			return "", global.ErrInternalError
		}

		userID, err := u.Repository.User.Create(ts, login, passHash)
		if err != nil {
			u.log.WithFields(lf).Errorln(u.logPrefix(),
				fmt.Sprintf("не удалось создать пользователя; ошибка: %v", err),
			)
			return "", global.ErrInternalError
		}

		userData, err = u.Repository.User.FindByID(ts, userID)
		if err != nil {
			u.log.WithFields(lf).Errorln(u.logPrefix(),
				fmt.Sprintf("не удалось загрузить пользователя; ошибка: %v", err),
			)
			return "", global.ErrInternalError
		}
	case nil:
		if passwordCorrect := passfunc.BcryptCheckPassword(password, userData.Password); !passwordCorrect {
			return "", user.ErrLoginOrPasswordIncorrect
		}
	default:
		u.log.WithFields(lf).Errorln(u.logPrefix(),
			fmt.Sprintf("не удалось найти пользователя; ошибка: %v", err),
		)
		return "", global.ErrInternalError
	}

	now := u.Bridge.Date.Now()

	jwtToken, err = jwt.NewJwtToken(userData.ID, userData.Login, now)
	if err != nil {
		u.log.WithFields(lf).Errorln(u.logPrefix(),
			fmt.Sprintf("не удалось создать jwt token; ошибка: %v", err),
		)
		return "", global.ErrInternalError
	}

	return jwtToken, nil
}

func (u *UserUsecase) FindUser(userID int) (user.User, error) {
	lf := logrus.Fields{"userID": userID}

	userData, exists := u.Repository.UserCache.Get(userID)
	if !exists {
		ts := u.SessionManager.CreateSession()
		if err := ts.Start(); err != nil {
			u.log.WithFields(lf).Errorln(u.logPrefix(),
				fmt.Sprintf("не удалось открыть бд сессию; ошибка: %v", err),
			)
			return userData, global.ErrInternalError
		}
		defer ts.Rollback()

		userDBData, err := u.Repository.User.FindByID(ts, userID)
		if err != nil {
			u.log.WithFields(lf).Errorln(u.logPrefix(),
				fmt.Sprintf("не удалось найти пользователя; ошибка: %v", err),
			)
			return user.User{}, global.ErrInternalError
		}
		// сохранить в кэше для последущих запросов
		u.Repository.UserCache.Add(userID, userDBData)
		return userDBData, nil
	}

	return userData, nil
}
