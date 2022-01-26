package dbclient

import (
	"github.com/3115826227/go-web-live/internal/db/tables"
	"github.com/3115826227/go-web-live/internal/errors"
	"gorm.io/gorm"
)

func addUserToken(c *Client, userToken tables.UserToken) error {
	return c.client.CreateObject(&userToken)
}

func getUserTokenByToken(c *Client, token string) (tables.UserToken, errors.Error) {
	var userToken tables.UserToken
	if err := c.client.GetDB().Where("token = ?", token).First(&userToken).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return tables.UserToken{}, errors.NewCommonError(errors.CodeUnLoginError)
		}
		return tables.UserToken{}, errors.NewCommonError(errors.CodeInternalError)
	}
	return userToken, nil
}

func deleteUserTokenByAccountId(c *Client, accountId string) error {
	return c.client.GetDB().Where("account_id = ?", accountId).Delete(&tables.UserToken{}).Error
}

func addUser(c *Client, user tables.User) errors.Error {
	if _, err := getUserByLoginName(c, user.LoginName); err != nil {
		if err.Code() == errors.CodeLoginNameNotExistError {
			if err1 := c.client.CreateObject(&user); err1 != nil {
				return errors.NewCommonError(errors.CodeInternalError)
			}
			return nil
		}
		return err
	}
	return errors.NewCommonError(errors.CodeLoginNameExistError)
}

func updateUser(c *Client, user tables.User) errors.Error {
	if err := c.client.GetDB().Save(&user).Error; err != nil {
		return errors.NewCommonError(errors.CodeInternalError)
	}
	return nil
}

func getUserByLoginName(c *Client, loginName string) (tables.User, errors.Error) {
	var u tables.User
	if err := c.client.GetDB().Where("login_name = ?", loginName).First(&u).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return tables.User{}, errors.NewCommonError(errors.CodeLoginNameNotExistError)
		}
		return tables.User{}, errors.NewCommonError(errors.CodeInternalError)
	}
	return u, nil
}

func getUsersByIds(c *Client, ids ...string) (map[string]tables.User, errors.Error) {
	var userMap map[string]tables.User
	var users []tables.User
	if err := c.client.GetDB().Where("account_id in (?)", ids).Find(&users).Error; err != nil {
		return nil, errors.NewCommonError(errors.CodeInternalError)
	}
	userMap = make(map[string]tables.User)
	for _, u := range users {
		userMap[u.AccountId] = u
	}
	return userMap, nil
}

func addVisitor(c *Client, visitor tables.Visitor) errors.Error {
	if err := c.client.CreateObject(&visitor); err != nil {
		return errors.NewCommonError(errors.CodeInternalError)
	}
	return nil
}

func updateVisitor(c *Client, visitor tables.Visitor) errors.Error {
	if err := c.client.GetDB().Save(&visitor).Error; err != nil {
		return errors.NewCommonError(errors.CodeInternalError)
	}
	return nil
}

func getVisitorsByIds(c *Client, ids ...string) (map[string]tables.Visitor, errors.Error) {
	var visitors []tables.Visitor
	if err := c.client.GetDB().Where("account_id in (?)", ids).Find(&visitors).Error; err != nil {
		return nil, errors.NewCommonError(errors.CodeInternalError)
	}
	visitorMap := make(map[string]tables.Visitor)
	for _, u := range visitors {
		visitorMap[u.AccountId] = u
	}
	return visitorMap, nil
}

func getVisitorByIp(c *Client, ip string) (tables.Visitor, errors.Error) {
	var visitor tables.Visitor
	if err := c.client.GetDB().Where("ip = ?", ip).First(&visitor).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return tables.Visitor{}, errors.NewCommonError(errors.CodeResourceNotExistError)
		}
		return tables.Visitor{}, errors.NewCommonError(errors.CodeInternalError)
	}
	return visitor, nil
}
