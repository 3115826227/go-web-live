package dbclient

import "github.com/3115826227/go-web-live/internal/db/tables"

func addUserToken(c *Client, userToken tables.UserToken) error {
	return c.client.CreateObject(&userToken)
}

func getUserTokenByToken(c *Client, token string) (userToken tables.UserToken, err error) {
	err = c.client.GetDB().Where("token = ?", token).First(&userToken).Error
	return
}

func deleteUserTokenByAccountId(c *Client, accountId string) error {
	return c.client.GetDB().Where("account_id = ?", accountId).Delete(&tables.UserToken{}).Error
}

func addUser(c *Client, user tables.User) error {
	return c.client.CreateObject(&user)
}

func getUserByLoginName(c *Client, loginName string) (tables.User, error) {
	var u tables.User
	if err := c.client.GetDB().Where("login_name = ?", loginName).First(&u).Error; err != nil {
		return tables.User{}, err
	}
	return u, nil
}

func getUsersByIds(c *Client, ids ...string) (userMap map[string]tables.User, err error) {
	var users []tables.User
	if err = c.client.GetDB().Where("account_id in (?)", ids).Find(&users).Error; err != nil {
		return
	}
	userMap = make(map[string]tables.User)
	for _, u := range users {
		userMap[u.AccountId] = u
	}
	return
}

func addVisitor(c *Client, visitor tables.Visitor) error {
	return c.client.CreateObject(&visitor)
}

func getVisitorsByIds(c *Client, ids ...string) (visitorMap map[string]tables.Visitor, err error) {
	var visitors []tables.Visitor
	if err = c.client.GetDB().Where("account_id in (?)", ids).Find(&visitors).Error; err != nil {
		return
	}
	visitorMap = make(map[string]tables.Visitor)
	for _, u := range visitors {
		visitorMap[u.AccountId] = u
	}
	return
}

func getVisitorByIp(c *Client, ip string) (tables.Visitor, error) {
	var visitor tables.Visitor
	if err := c.client.GetDB().Where("ip = ?", ip).First(&visitor).Error; err != nil {
		return tables.Visitor{}, err
	}
	return visitor, nil
}
