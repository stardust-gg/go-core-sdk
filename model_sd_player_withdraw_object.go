/*
 * Stardust API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2023-02-22T22:32:07Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package stardust_core_sdk

type SdPlayerWithdrawObject struct {
	TokenId int32 `json:"tokenId"`
	// u64 Number as String, min: 0, max: 9223372036854775807
	Amount string `json:"amount"`
}