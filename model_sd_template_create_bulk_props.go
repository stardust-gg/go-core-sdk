/*
 * Stardust API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2023-02-22T22:32:07Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package stardust_core_sdk

type SdTemplateCreateBulkProps struct {
	// Properties for immutable on the Template
	Immutable *interface{} `json:"immutable,omitempty"`
	// Properties for mutable on the Template
	Mutable *interface{} `json:"mutable,omitempty"`
	// Properties for mutable on the Item
	Mutable *interface{} `json:"$mutable,omitempty"`
}
