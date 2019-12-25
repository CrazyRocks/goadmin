package bean

type SessionUser struct {
	Id       int    `form:"id" json:"id"`             // 主键
	Username string `form:"username" json:"username"` // 登录名/11111
	Mobile   string `form:"mobile" json:"mobile"`     // 真实姓名
}
