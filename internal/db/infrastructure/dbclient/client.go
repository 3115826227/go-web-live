package dbclient

import (
	"github.com/3115826227/go-web-live/internal/constant"
	"github.com/3115826227/go-web-live/internal/db/infrastructure/gormclient"
	"github.com/3115826227/go-web-live/internal/db/interfaces"
	"github.com/3115826227/go-web-live/internal/db/tables"
	"github.com/3115826227/go-web-live/internal/log"
)

type Client struct {
	client interfaces.GORMClient
	lc     log.Logging
}

var (
	dbClient interfaces.DBClient
)

func GetDBClient() interfaces.DBClient {
	return dbClient
}

func InitDB(dataSource string, lc log.Logging) error {
	client, err := gormclient.NewSQLiteClient(dataSource, lc)
	if err != nil {
		return err
	}
	if err = client.InitTable(
		&tables.User{},
		&tables.UserToken{},
		&tables.Visitor{},
		&tables.LiveRoom{},
		&tables.Session{},
		&tables.Message{},
		&tables.UserRelation{},
	); err != nil {
		return err
	}
	dbClient = &Client{
		client: client,
		lc:     lc,
	}
	return nil
}

func (c *Client) Close() {
	c.client.Close()
}

func (c *Client) AddUserToken(userToken tables.UserToken) error {
	return addUserToken(c, userToken)
}

func (c *Client) GetUserTokenByToken(token string) (tables.UserToken, error) {
	return getUserTokenByToken(c, token)
}

func (c *Client) DeleteUserTokenByAccountId(accountId string) error {
	return deleteUserTokenByAccountId(c, accountId)
}

func (c *Client) AddUser(user tables.User) error {
	return addUser(c, user)
}

func (c *Client) GetUserByLoginName(loginName string) (tables.User, error) {
	return getUserByLoginName(c, loginName)
}

func (c *Client) GetUserByIds(accountIds ...string) (map[string]tables.User, error) {
	return getUsersByIds(c, accountIds...)
}

func (c *Client) AddVisitor(visitor tables.Visitor) error {
	return addVisitor(c, visitor)
}

func (c *Client) GetVisitorByIds(accountIds ...string) (map[string]tables.Visitor, error) {
	return getVisitorsByIds(c, accountIds...)
}

func (c *Client) GetVisitorByIp(ip string) (tables.Visitor, error) {
	return getVisitorByIp(c, ip)
}

func (c *Client) AddLiveRoom(liveRoom tables.LiveRoom) error {
	return addLiveRoom(c, liveRoom)
}

func (c *Client) GetLiveRooms(page, pageSize int64, accountId string) ([]tables.LiveRoom, int64, error) {
	return getLiveRooms(c, page, pageSize, accountId)
}

func (c *Client) GetLiveRoomById(liveRoomId string) (tables.LiveRoom, error) {
	return getLiveRoomById(c, liveRoomId)
}

func (c *Client) GetLiveRoomByAccountId(accountId string) (tables.LiveRoom, error) {
	return getLiveRoomByAccountId(c, accountId)
}

func (c *Client) UpdateLiveRoomStatus(accountId string, status constant.LiveRoomStatus) error {
	return updateLiveRoomStatus(c, accountId, status)
}

func (c *Client) UpdateLiveRoomSendMessagePermission(accountId string, permission bool) error {
	return updateLiveRoomSendMessagePermission(c, accountId, permission)
}

func (c *Client) AddSession(session tables.Session) error {
	return addSession(c, session)
}

func (c *Client) UpdateSession(session tables.Session) error {
	return updateSession(c, session)
}

func (c *Client) GetSessions(page, pageSize int64, accountId string) ([]tables.Session, int64, error) {
	return getSessions(c, page, pageSize, accountId)
}

func (c *Client) GetSessionById(sessionId string) (tables.Session, error) {
	return getSessionById(c, sessionId)
}

func (c *Client) UpdateSessionSendMessagePermission(sessionId string, permission bool) error {
	return updateSessionSendMessagePermission(c, sessionId, permission)
}

func (c *Client) AddMessage(msg tables.Message) error {
	return addMessage(c, msg)
}

func (c *Client) GetMessagesByBiz(bizId string, bizType constant.BizType, timestamp int64, page, pageSize int64) ([]tables.Message, error) {
	return getMessages(c, bizId, bizType, timestamp, page, pageSize)
}

func (c *Client) GetUserRelationTotal(bizId string, bizType constant.BizType) (int64, error) {
	return getUserRelationTotal(c, bizId, bizType)
}

func (c *Client) GetUserRelations(bizId string, bizType constant.BizType, page, pageSize int64) ([]tables.UserRelation, int64, error) {
	return getUserRelations(c, bizId, bizType, page, pageSize)
}

func (c *Client) AddUserRelation(relation tables.UserRelation) error {
	return addUserRelation(c, relation)
}

func (c *Client) GetUserRelation(bizId string, bizType constant.BizType, accountId string) (relation tables.UserRelation, err error) {
	return getUserRelation(c, bizId, bizType, accountId)
}

func (c *Client) DeleteUserRelation(bizId string, bizType constant.BizType, accountId string) error {
	return deleteUserRelation(c, bizId, bizType, accountId)
}

func (c *Client) UpdateUserSendMessagePermission(bizId string, bizType constant.BizType, accountId string, permission bool) error {
	return updateUserSendMessagePermission(c, bizId, bizType, accountId, permission)
}
