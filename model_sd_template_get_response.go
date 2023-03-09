/*
 * Stardust API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2023-02-22T22:32:07Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package stardust_core_sdk

type SdTemplateGetResponse struct {
	// Game ID Number (unsigned 32 bit integer)
	GameId int32 `json:"gameId"`
	Id int32 `json:"_id"`
	// u96 Number as String, min: 0, max: 39614081257132168796771975167
	Cap string `json:"cap"`
	// The type of custom contract bieng used for this template.
	ContractType string `json:"contractType,omitempty"`
	// u96 Number as String, min: 0, max: 39614081257132168796771975167
	TotalSupply string `json:"totalSupply"`
	// The name of the template
	Name string `json:"name"`
	Type_ string `json:"type"`
	Props *SdTemplateGetAllResponseProps `json:"props"`
	// Returned to marketplaces as contract metadata
	PublicContractMetadata *interface{} `json:"publicContractMetadata,omitempty"`
	// Inherited by tokens, and returned to marketplaces as token metadata
	PublicTokenMetadata *interface{} `json:"publicTokenMetadata,omitempty"`
	Fees []SdGameGetResponseFees `json:"fees,omitempty"`
}
