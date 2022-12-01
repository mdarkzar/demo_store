package restapi

import (
	"store/internal/entity/global"
	"store/internal/transaction"

	"github.com/gin-gonic/gin"
)

type LoadDataFunc[T any] func(ts transaction.Session) (T, error)

type ReturnErrFunc func(ts transaction.Session) error

type ReturnResultFunc func(ts transaction.Session) (gin.H, error)

func (e *RestAPI) ReturnSuccessNull(
	c *gin.Context,
	f ReturnErrFunc,
) {
	ts := e.SessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		e.log.Errorln(e.logPrefix(), "не удалось открыть бд сессию", err)
		e.errorResponse(c, global.ErrInternalError)
		return
	}
	defer ts.Rollback()

	err := f(ts)
	if err != nil {
		e.errorResponse(c, err)
		return
	}

	if err := ts.Commit(); err != nil {
		e.log.Errorln(e.logPrefix(), "не удалось сохранить бд сессию", err)
		e.errorResponse(c, global.ErrInternalError)
		return
	}

	e.returnResult(c, "ok")
}

func (e *RestAPI) ReturnResult(
	c *gin.Context,
	f ReturnResultFunc,
) {
	ts := e.SessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		e.log.Errorln(e.logPrefix(), "не удалось открыть бд сессию", err)
		e.errorResponse(c, global.ErrInternalError)
		return
	}
	defer ts.Rollback()

	result, err := f(ts)
	if err != nil {
		e.errorResponse(c, err)
		return
	}

	if err := ts.Commit(); err != nil {
		e.log.Errorln(e.logPrefix(), "не удалось сохранить бд сессию", err)
		e.errorResponse(c, global.ErrInternalError)
		return
	}

	e.returnResult(c, result)
}

func (e *RestAPI) ReturnResultWithoutCommit(
	c *gin.Context,
	f ReturnResultFunc,
) {
	ts := e.SessionManager.CreateSession()
	if err := ts.Start(); err != nil {
		e.log.Errorln(e.logPrefix(), "не удалось открыть бд сессию", err)
		e.errorResponse(c, global.ErrInternalError)
		return
	}
	defer ts.Rollback()

	result, err := f(ts)
	if err != nil {
		e.errorResponse(c, err)
		return
	}

	e.returnResult(c, result)
}
