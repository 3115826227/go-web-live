package tables

// 用户登录token表
type UserToken struct {
	// 用户登录token
	Token string `gorm:"column:token;unique"`
	// 用户唯一Id
	AccountId string `gorm:"column:account_id;pk"`
	// 存储的序列化数据
	Value string `gorm:"column:value"`
}

func (table *UserToken) TableName() string {
	return "wl_user_token"
}

// 注册用户表
type User struct {
	CommonIntField
	// 用户唯一id 长度默认为8位
	AccountId string `gorm:"column:account_id;pk"`
	// 用户名
	Username string `gorm:"column:username;unique"`
	// 登录名称
	LoginName string `gorm:"column:login_name;unique"`
	// 登录密码 sha256加密后的值
	Password string `gorm:"column:password"`
}

func (table *User) TableName() string {
	return "wl_user"
}

// 访客用户表
type Visitor struct {
	CommonIntField
	// 访客用户唯一id 长度默认为8位
	AccountId string `gorm:"column:account_id;pk"`
	// 访客IP地址 唯一
	IP string `gorm:"column:ip;unique_index"`
	// 访客用户名
	Username string `gorm:"column:username;unique"`
}

func (table *Visitor) TableName() string {
	return "wl_visitor"
}
