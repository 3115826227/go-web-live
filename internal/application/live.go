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

func OpenLiveRoom(ctx context.Context, accountId string) error {
	var now = time.Now().Unix()
	var liveRoom = tables.LiveRoom{
		AccountId: accountId,
	}
	liveRoom.CreateTimestamp, liveRoom.UpdateTimestamp = now, now
	liveRoom.ID = utils.GenerateSerialNumber()
	return dbclient.GetDBClient().AddLiveRoom(liveRoom)
}

func GetLiveRoomIdByAccountId(ctx context.Context, accountId string) (string, bool, error) {
	live, err := dbclient.GetDBClient().GetLiveRoomByAccountId(accountId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", false, nil
		}
		return "", false, err
	}
	return live.ID, true, nil
}

func GetLiveRoomByAccountId(ctx context.Context, accountId string) (dtos.LiveOrigin, bool, error) {
	live, err := dbclient.GetDBClient().GetLiveRoomByAccountId(accountId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dtos.LiveOrigin{}, false, nil
		}
		return dtos.LiveOrigin{}, false, err
	}
	var userTotal int64
	userTotal, err = dbclient.GetDBClient().GetUserRelationTotal(live.ID, constant.LiveRoomBizType)
	if err != nil {
		return dtos.LiveOrigin{}, true, err
	}
	var l = dtos.LiveOrigin{
		LiveRoomID:           live.ID,
		LiveRoomOrigin:       live.AccountId,
		LiveRoomStatus:       live.Status,
		LiveRoomUserTotal:    userTotal,
		LiveRoomMaxUserTotal: live.MaxUserTotal,
		LiveRoomReports:      live.Reports,
	}
	return l, true, nil
}

func QueryLive(ctx context.Context, req requests.PageCommonReq, accountId string) ([]dtos.Live, int64, error) {
	lives, total, err := dbclient.GetDBClient().GetLiveRooms(req.Page, req.PageSize, accountId)
	if err != nil {
		return nil, 0, err
	}
	var list = make([]dtos.Live, 0)
	for _, live := range lives {
		var userTotal int64
		userTotal, err = dbclient.GetDBClient().GetUserRelationTotal(live.ID, constant.LiveRoomBizType)
		if err != nil {
			return nil, 0, err
		}
		list = append(list, dtos.Live{
			LiveRoomID:        live.ID,
			LiveRoomOrigin:    live.AccountId,
			LiveRoomStatus:    live.Status,
			LiveRoomUserTotal: userTotal,
		})
	}
	return list, total, nil
}

func QueryLiveDetail(ctx context.Context, liveRoomId string) (dtos.LiveDetail, bool, error) {
	live, err := dbclient.GetDBClient().GetLiveRoomById(liveRoomId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dtos.LiveDetail{}, false, nil
		}
		return dtos.LiveDetail{}, false, err
	}
	var userTotal int64
	userTotal, err = dbclient.GetDBClient().GetUserRelationTotal(live.ID, constant.LiveRoomBizType)
	if err != nil {
		return dtos.LiveDetail{}, true, err
	}
	var detail = dtos.LiveDetail{
		Live: dtos.Live{
			LiveRoomID:        live.ID,
			LiveRoomOrigin:    live.AccountId,
			LiveRoomStatus:    live.Status,
			LiveRoomUserTotal: userTotal,
		},
		PermissionSendMessage: live.PermissionSendMessage,
	}
	return detail, true, nil
}

func UpdateLiveRoomStatus(ctx context.Context, accountId string, status constant.LiveRoomStatus) error {
	return dbclient.GetDBClient().UpdateLiveRoomStatus(accountId, status)
}

func UpdateLiveRoomSendMessagePermission(ctx context.Context, accountId string, permission bool) error {
	return dbclient.GetDBClient().UpdateLiveRoomSendMessagePermission(accountId, permission)
}
