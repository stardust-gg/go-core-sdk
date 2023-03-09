/*
 * Stardust API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2023-02-22T22:32:07Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package stardust_core_sdk

type SdGameMutateRequest struct {
	// The games description
	Description string `json:"description,omitempty"`
	// The games news
	News string `json:"news,omitempty"`
	// Game properties
	Props *interface{} `json:"props,omitempty"`
	// Test mode will enable any test mode features. Currently in development.
	TestMode bool `json:"testMode,omitempty"`
}
