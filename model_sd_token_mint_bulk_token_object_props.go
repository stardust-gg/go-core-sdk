/*
 * Stardust API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2023-02-22T22:32:07Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package stardust_core_sdk

// Internal metadata
type SdTokenMintBulkTokenObjectProps struct {
	// Specifies which properties you want to add (ex. {prop1: 5, prop2: 6, prop3: 7})
	Mutable *interface{} `json:"mutable"`
	// Specifies which properties you want to add (ex. {prop1: 5, prop2: 6, prop3: 7})
	Immutable *interface{} `json:"immutable,omitempty"`
}
