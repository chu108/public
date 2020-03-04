package myswagger

// DocSwagger ...
type DocSwagger struct {
	client *APIBody
}

// NewDoc 新建一个swagger doc
func NewDoc() *DocSwagger {

	doc := &DocSwagger{}
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
	return nil
}
