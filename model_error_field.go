/*
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package godaddy

type ErrorField struct {
	// Short identifier for the error, suitable for indicating the specific error within client code
	Code string `json:"code"`
	// Human-readable, English description of the problem with the contents of the field
	Message string `json:"message,omitempty"`
	// <ul> <li style='margin-left: 12px;'>JSONPath referring to a field containing an error</li> <strong style='margin-left: 12px;'>OR</strong> <li style='margin-left: 12px;'>JSONPath referring to a field that refers to an object containing an error, with more detail in `pathRelated`</li> </ul>
	Path string `json:"path"`
	// JSONPath referring to a field containing an error, which is referenced by `path`
	PathRelated string `json:"pathRelated,omitempty"`
}
