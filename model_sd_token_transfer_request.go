/*
 * Stardust API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2023-02-22T22:32:07Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package stardust_core_sdk

type SdTokenTransferRequest struct {
	// The Player's id, can be found with Player/getId(s). (ex. CzySggxVQz51jciGRFDY7d5BER2fav6TNEnPGjusPJPd)
	FromPlayerId string `json:"fromPlayerId"`
	// The Player's id, can be found with Player/getId(s). (ex. 53ywNSVp46QpiA6S86DLLfeKVfjcSAFxHR2L9j8tnte2)
	ToPlayerId string `json:"toPlayerId"`
	// An array of objects of which tokens to transfer (ex. [{tokenId: 5, amount: \"3\"}])
	TokenObjects []SdTokenBurnRequestTokenObjects `json:"tokenObjects"`
}