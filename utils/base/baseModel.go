package base


// 定义model interface
type IModel interface {
	// 获取表明
	TableName() string
	// 获取主键值
	PkVal() int
}
