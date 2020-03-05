package myswagger

import (
	"strings"

	"github.com/xxjwxc/public/tools"
)

// docSwagger ...
type docSwagger struct {
	client *APIBody
}

// NewDoc 新建一个swagger doc
func NewDoc() *docSwagger {
	doc := &docSwagger{}
	doc.client = &APIBody{
		Head:     Head{Swagger: version},
		Info:     info,
		Host:     host,
		BasePath: basePath,
		// Tags
		Schemes: schemes,
		// Patchs
		// SecurityDefinitions
		// Definitions
		ExternalDocs: externalDocs,
	}
	doc.client.Patchs = make(map[string]map[string]Param)
	return doc
}

// AddTag add tag (排他)
func (doc *docSwagger) AddTag(tag Tag) {
	for _, v := range doc.client.Tags {
		if v.Name == tag.Name { // find it
			return
		}
	}

	doc.client.Tags = append(doc.client.Tags, tag)
}

// AddDefinitions 添加 通用结构体定义
func (doc *docSwagger) AddDefinitions(key string, def Definition) {
	// for k := range doc.client.Definitions {
	// 	if k == key { // find it
	// 		return
	// 	}
	// }
	if doc.client.Definitions == nil {
		doc.client.Definitions = make(map[string]Definition)
	}

	doc.client.Definitions[key] = def
}

// AddPatch ... API 路径 paths 和操作在 API 规范的全局部分定义
func (doc *docSwagger) AddPatch(url string, p Param, metheds ...string) {
	if doc.client.Patchs[url] == nil {
		doc.client.Patchs[url] = make(map[string]Param)
	}
	if len(p.Consumes) == 0 {
		p.Consumes = reqCtxType
	}
	if len(p.Produces) == 0 {
		p.Produces = respCtxType
	}
	if p.Responses == nil {
		p.Responses = map[string]map[string]string{
			"400": {"description": "v"},
			"404": {"description": "not found"},
			"405": {"description": "Validation exception"},
		}
	}

	for _, v := range metheds {
		doc.client.Patchs[url][strings.ToLower(v)] = p
	}
}

// GetAPIString 获取返回数据
func (doc *docSwagger) GetAPIString() string {
	return tools.GetJSONStr(doc.client, true)
}
