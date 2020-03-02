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

// Body swagger api body info
type Body struct {
	Head
	Info     Info     `json:"info"`
	Host     string   `json:"host"`     // http host
	BasePath string   `json:"basePath"` // 根级别
	Tags     []Tag    `json:"tags"`
	Schemes  []string `json:"schemes"` // http/https
}
