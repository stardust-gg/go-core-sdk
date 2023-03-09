/*
 * Stardust API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2023-02-22T22:32:07Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package stardust_core_sdk

type SdPlayerWithdrawRequest struct {
	// Blockchain Address
	Address string `json:"address"`
	// The Player's id, can be found with Player/getId(s). Also returned from player/create (ex. CzySggxVQz51jciGRFDY7d5BER2fav6TNEnPGjusPJPd)
	PlayerId string `json:"playerId"`
	// array of Token objects
	TokenObjects []SdPlayerWithdrawObject `json:"tokenObjects"`
}
