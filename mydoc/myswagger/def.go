package myswagger

// Head Swagger 版本
type Head struct {
	Swagger string `json:"swagger"`
}

// Info 指定 API 的 info-title
type Info struct {
	Description string `json:"description"`
	Version     string `json:"version"`
	Title       string `json:"title"`
}

// ExternalDocs tags of group
type ExternalDocs struct {
	Description string `json:"description"` // 描述
	URL         string `json:"url"`         // url addr
}

// Tag group of tags
type Tag struct {
	Name         string       `json:"name"`         // tags name
	Description  string       `json:"description"`  // 描述
	ExternalDocs ExternalDocs `json:"externalDocs"` // doc group of tags
}

// Schema 引用
type Schema struct {
	Ref string `json:"$ref"` // 主体模式和响应主体模式中引用
}

// Element 元素定义
type Element struct {
	In          string `json:"in"`          // 入参
	Name        string `json:"name"`        // 参数名字
	Description string `json:"description"` // 描述
	Required    bool   `json:"required"`    // 是否必须
	Schema      Schema `json:"schema"`      // 引用
}

// Param API 路径 paths 和操作在 API 规范的全局部分定义
type Param struct {
	Tags        []Tag       `json:"tags"`        // 分组标记
	Summary     string      `json:"summary"`     // 摘要
	Description string      `json:"description"` // 描述
	OperationID string      `json:"operationId"` // 操作id
	Consumes    []string    `json:"consumes"`    // Parameter content type
	Produces    []string    `json:"produces"`    // Response content type
	Parameters  []Element   `json:"parameters"`  // 请求参数
	Responses   interface{} `json:"responses"`   // 返回参数
}

// Body swagger api body info
type Body struct {
	Head
	Info     Info                        `json:"info"`
	Host     string                      `json:"host"`     // http host
	BasePath string                      `json:"basePath"` // 根级别
	Tags     []Tag                       `json:"tags"`
	Schemes  []string                    `json:"schemes"` // http/https
	Patchs   map[string]map[string]Param `json:"paths"`   // API 路径
}
