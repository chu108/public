package mydoc

// ElementInfo 结构信息
type ElementInfo struct {
	Name     string     // 参数名
	Type     string     // 类型
	TypeRef  StructInfo // 类型定义
	Requierd bool       // 是否必须
	Note     string     // 注释
	Default  string     // 默认值
}

// StructInfo struct define
type StructInfo struct {
	element []ElementInfo // 结构体元素
	Note    string        // 注释
}
