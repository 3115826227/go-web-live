package application

import (
	"context"
	"github.com/3115826227/go-web-live/internal/constant"
	"github.com/3115826227/go-web-live/internal/db/infrastructure/dbclient"
	"github.com/3115826227/go-web-live/internal/db/tables"
	"github.com/3115826227/go-web-live/internal/dtos"
	"github.com/3115826227/go-web-live/internal/handle/requests"
	"github.com/3115826227/go-web-live/internal/utils"
	"gorm.io/gorm"
	"time"
)

func AddSession(ctx context.Context, accountId, name, description string) error {
	var now = time.Now().Unix()
	var session = tables.Session{
		AccountId:   accountId,
		Name:        name,
		Description: description,
	}
	session.CreateTimestamp, session.UpdateTimestamp = now, now
	session.ID = utils.GenerateSerialNumber()
	return dbclient.GetDBClient().AddSession(session)
}

func UpdateSession(ctx context.Context, accountId, name, description string) error {
	var session = tables.Session{
		AccountId:   accountId,
		Name:        name,
		Description: description,
	}
	session.UpdateTimestamp = time.Now().Unix()
	return dbclient.GetDBClient().UpdateSession(session)
}

func GetSessionById(ctx context.Context, sessionId string) (dtos.SessionOrigin, bool, error) {
	session, err := dbclient.GetDBClient().GetSessionById(sessionId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dtos.SessionOrigin{}, false, nil
		}
		return dtos.SessionOrigin{}, false, err
	}
	var userTotal int64
	userTotal, err = dbclient.GetDBClient().GetUserRelationTotal(session.ID, constant.SessionBizType)
	if err != nil {
		return dtos.SessionOrigin{}, true, err
	}
	var s = dtos.SessionOrigin{
		SessionID:        session.ID,
		SessionOrigin:    session.AccountId,
		SessionUserTotal: userTotal,
		SessionReports:   session.Reports,
	}
	return s, true, nil
}

func QuerySession(ctx context.Context, req requests.PageCommonReq, accountId string) ([]dtos.Session, int64, error) {
	sessions, total, err := dbclient.GetDBClient().GetSessions(req.Page, req.PageSize, accountId)
	if err != nil {
		return nil, 0, err
	}
	var list = make([]dtos.Session, 0)
	for _, session := range sessions {
		var userTotal int64
		userTotal, err = dbclient.GetDBClient().GetUserRelationTotal(session.ID, constant.SessionBizType)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, dtos.Session{
			SessionID:        session.ID,
			SessionOrigin:    session.AccountId,
			SessionUserTotal: userTotal,
		})
	}
	return list, total, nil
}

func QuerySessionDetail(ctx context.Context, sessionId string) (dtos.SessionDetail, bool, error) {
	session, err := dbclient.GetDBClient().GetSessionById(sessionId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dtos.SessionDetail{}, false, nil
		}
		return dtos.SessionDetail{}, false, err
	}
	var userTotal int64
	userTotal, err = dbclient.GetDBClient().GetUserRelationTotal(session.ID, constant.SessionBizType)
	if err != nil {
		return dtos.SessionDetail{}, true, err
	}
	var detail = dtos.SessionDetail{
		Session: dtos.Session{
			SessionID:        session.ID,
			SessionOrigin:    session.AccountId,
			SessionUserTotal: userTotal,
		},
		PermissionSendMessage: session.PermissionSendMessage,
	}
	return detail, true, nil
}

func UpdateSessionSendMessagePermission(ctx context.Context, sessionId string, permission bool) error {
	return dbclient.GetDBClient().UpdateSessionSendMessagePermission(sessionId, permission)
}
