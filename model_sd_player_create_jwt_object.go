/*
 * Stardust API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2023-02-22T22:32:07Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package stardust_core_sdk

type SdPlayerCreateJwtObject struct {
	// Your provider JWT name
	Type_ string `json:"type"`
	// idToken of the user session
	IdToken string `json:"idToken,omitempty"`
	// Access token of the user session
	AccessToken string `json:"accessToken,omitempty"`
}
