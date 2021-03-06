package application

import (
	"context"
	"github.com/3115826227/go-web-live/internal/db/infrastructure/dbclient"
	"github.com/3115826227/go-web-live/internal/db/tables"
	"github.com/3115826227/go-web-live/internal/dtos"
	"github.com/3115826227/go-web-live/internal/errors"
	"github.com/3115826227/go-web-live/internal/utils"
	"gorm.io/gorm"
	"time"
)

func UserRegister(ctx context.Context, ul dtos.UserRegister) errors.Error {
	var now = time.Now().Unix()
	var user = tables.User{
		AccountId: utils.GenerateSerialNumber(),
		Username:  ul.Username,
		LoginName: ul.LoginName,
		Password:  ul.Password,
	}
	user.CreateTimestamp, user.UpdateTimestamp = now, now
	return dbclient.GetDBClient().AddUser(user)
}

func UserLogin(ctx context.Context, loginName, password string) (dtos.User, errors.Error) {
	user, err := dbclient.GetDBClient().GetUserByLoginName(loginName)
	if err != nil {
		return dtos.User{}, err
	}
	if user.Password != password {
		err = errors.NewCommonError(errors.CodePasswordFaultError)
		return dtos.User{}, err
	}
	var u = dtos.User{
		AccountId: user.AccountId,
		Username:  user.Username,
	}
	return u, nil
}

func GetUserByLoginName(ctx context.Context, loginName string) (dtos.User, bool, error) {
	user, err := dbclient.GetDBClient().GetUserByLoginName(loginName)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return dtos.User{}, false, nil
		}
		return dtos.User{}, false, err
	}
	var u = dtos.User{
		AccountId: user.AccountId,
		Username:  user.Username,
		Visitor:   false,
	}
	return u, true, nil
}

func GetUserById(ctx context.Context, accountId string) (dtos.User, errors.Error) {
	mp, err := dbclient.GetDBClient().GetUserByIds(accountId)
	if err != nil {
		return dtos.User{}, err
	}
	var user = dtos.User{
		AccountId: mp[accountId].AccountId,
		Username:  mp[accountId].Username,
	}
	return user, nil
}

func GetUserByIds(ctx context.Context, accountIds []string) (userMap map[string]dtos.User, err errors.Error) {
	var mp map[string]tables.User
	if mp, err = dbclient.GetDBClient().GetUserByIds(accountIds...); err != nil {
		return
	}
	userMap = make(map[string]dtos.User)
	for _, u := range mp {
		userMap[u.AccountId] = dtos.User{
			AccountId: u.AccountId,
			Username:  u.Username,
		}
	}
	return
}

func AddVisitor(ctx context.Context, accountId, username, ip string) errors.Error {
	var now = time.Now().Unix()
	var visitor = tables.Visitor{
		AccountId: accountId,
		Username:  username,
		IP:        ip,
	}
	visitor.CreateTimestamp, visitor.UpdateTimestamp = now, now
	return dbclient.GetDBClient().AddVisitor(visitor)
}

func GetVisitorByIp(ctx context.Context, ip string) (dtos.User, bool, errors.Error) {
	visitor, err := dbclient.GetDBClient().GetVisitorByIp(ip)
	if err != nil {
		if err.Code() == errors.CodeResourceNotExistError {
			return dtos.User{}, false, nil
		}
		return dtos.User{}, false, err
	}
	var u = dtos.User{
		AccountId: visitor.AccountId,
		Username:  visitor.Username,
		Visitor:   true,
	}
	return u, true, nil
}

func GetVisitorByIds(ctx context.Context, accountIds []string) (userMap map[string]dtos.User, err errors.Error) {
	var mp map[string]tables.Visitor
	if mp, err = dbclient.GetDBClient().GetVisitorByIds(accountIds...); err != nil {
		return
	}
	userMap = make(map[string]dtos.User)
	for _, u := range mp {
		userMap[u.AccountId] = dtos.User{
			AccountId: u.AccountId,
			Username:  u.Username,
			Visitor:   true,
		}
	}
	return
}
