package interfaces

import (
	"github.com/3115826227/go-web-live/internal/constant"
	"github.com/3115826227/go-web-live/internal/db/tables"
)

type DBClient interface {
	Close()

	// 添加用户登录信息
	AddUserToken(userToken tables.UserToken) error
	// 根据token查询用户登录信息
	GetUserTokenByToken(token string) (tables.UserToken, error)
	// 根据用户id删除用户登录信息
	DeleteUserTokenByAccountId(accountId string) error
	// 添加用户
	AddUser(user tables.User) error
	// 根据登录名称查找用户
	GetUserByLoginName(loginName string) (tables.User, error)
	// 根据用户id列表查询用户
	GetUserByIds(accountIds ...string) (map[string]tables.User, error)
	// 添加访客
	AddVisitor(visitor tables.Visitor) error
	// 根据id列表查询访客
	GetVisitorByIds(accountIds ...string) (map[string]tables.Visitor, error)
	// 根据ip查询访客
	GetVisitorByIp(ip string) (tables.Visitor, error)
	// 添加用户直播间
	AddLiveRoom(liveRoom tables.LiveRoom) error
	// 查询直播间列表
	GetLiveRooms(page, pageSize int64, accountId string) ([]tables.LiveRoom, int64, error)
	// 通过直播间id查询直播间
	GetLiveRoomById(liveRoomId string) (tables.LiveRoom, error)
	// 通过用户id查询直播间
	GetLiveRoomByAccountId(accountId string) (tables.LiveRoom, error)
	// 更新直播间状态
	UpdateLiveRoomStatus(accountId string, status constant.LiveRoomStatus) error
	// 更新直播间禁言/解除禁言状态
	UpdateLiveRoomSendMessagePermission(accountId string, permission bool) error
	// 添加会话
	AddSession(session tables.Session) error
	// 更新会话信息
	UpdateSession(session tables.Session) error
	// 查询会话列表
	GetSessions(page, pageSize int64, accountId string) ([]tables.Session, int64, error)
	// 根据id查询会话
	GetSessionById(sessionId string) (tables.Session, error)
	// 更新会话禁言/解除禁言状态
	UpdateSessionSendMessagePermission(sessionId string, permission bool) error
	// 添加消息
	AddMessage(msg tables.Message) error
	// 通过业务id和类型查询消息列表
	GetMessagesByBiz(bizId string, bizType constant.BizType, timestamp int64, page, pageSize int64) ([]tables.Message, error)
	// 获取关联用户人数
	GetUserRelationTotal(bizId string, bizType constant.BizType) (int64, error)
	// 获取关联用户列表
	GetUserRelations(bizId string, bizType constant.BizType, page, pageSize int64) ([]tables.UserRelation, int64, error)
	// 添加关联用户
	AddUserRelation(relation tables.UserRelation) error
	// 查询单个关联用户
	GetUserRelation(bizId string, bizType constant.BizType, accountId string) (relation tables.UserRelation, err error)
	// 删除关联用户
	DeleteUserRelation(bizId string, bizType constant.BizType, accountId string) error
	// 更新某个用户禁言/解除禁言
	UpdateUserSendMessagePermission(bizId string, bizType constant.BizType, accountId string, permission bool) error
}
