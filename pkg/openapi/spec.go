package openapi

// ApplicationJSON is a media type
const ApplicationJSON = "application/json"

// Schema defines an api object
type Schema map[string]interface{}

// Properties is a key value mapping from string to property
type Properties map[string]interface{}

// FieldProperty is a free form property
type FieldProperty map[string]interface{}

// Property of an object
type Property struct {
	Type        string `json:"type"`
	Format      string `json:"format,omitempty"`
	Description string `json:"description"`
}

// SchemaRef references a schema from within the components tree
type SchemaRef struct {
	Ref string `json:"$ref"`
}

// MediaType is the body of a respones object, e.g. a JSON payload
// with a schema. All schemas here are referenced.
type MediaType struct {
	Schema SchemaRef `json:"schema"`
}

// Response encode a mapping between a status code and a
// media type object
type Response struct {
	Description string               `json:"description"`
	Content     map[string]MediaType `json:"content,omitempty"`
}

// A PathItem describes an api endpoint in a mapping
// of HTTP verb to sdescription
type PathItem map[string]struct {
	Description string              `json:"description"`
	Responses   map[string]Response `json:"responses"`
}

// SecurityScheme describes a security scheme
type SecurityScheme struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Name        string `json:"name"`
	In          string `json:"in"` // query, header cookie
	// Scheme       string `json:"scheme"` // RFC7235
	BearerFormat string `json:"bearerFormat"`
}

// Components describe the components object
type Components struct {
	Schemas   map[string]Schema   `json:"schemas"`
	Responses map[string]Response `json:"responses"`
	// Parameters      map[string]Parameter      `json:"parameters"`
	// Examples        map[string]Example        `json:"examples"`
	// RequestBodies   map[string]RequestBody    `json:"requestBodies"`
	// Headers         map[string]Headers        `json:"headers"`
	SecuritySchemes map[string]SecurityScheme `json:"securitySchemes"`
}

// Tag object for metadata
type Tag struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

// Server is a description of an API server
type Server struct {
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}

// SecuritySpec is a polymorphic type
type SecuritySpec map[string][]interface{}

// Security describes the auth methods of the API
type Security []SecuritySpec

// Path is a mapping of http verb to response
type Path map[string]struct {
	Description string              `json:"description"`
	Responses   map[string]Response `json:"responses"`
}

// License information
type License struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Info about the API
type Info struct {
	Title       string  `json:"title"`
	Version     string  `json:"version"`
	License     License `json:"license"`
	Description string  `json:"description"`
}

// Spec is the OpenAPI document describing the b3scale API
type Spec struct {
	OpenAPI    string          `json:"openapi"` // Version
	Info       Info            `json:"info"`
	Paths      map[string]Path `json:"paths"`
	Components Components      `json:"components"`
	Tags       []Tag           `json:"tags"`
	Servers    []Server        `json:"servers,omitempty"`
	Security   Security        `json:"security"`
}
