# SdTemplateCreateRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the template (ex. Bronze Axe) | [default to null]
**Cap** | **string** | u96 Number as String (ex. 200000000)  | [default to null]
**ContractType** | **string** | The type of custom contract to use for this template. Default will use a shared contract. | [optional] [default to null]
**Type_** | **string** | FT is a currency where every instance is the same, NFT is where every token instance differes (ex. FT) | [default to null]
**OwnerAddress** | **string** | Blockchain address to set as owner of the custom contract. Required if contractType is set. | [optional] [default to null]
**Props** | [***SdTemplateCreateRequestProps**](SDTemplateCreateRequest_props.md) |  | [default to null]
**PublicContractMetadata** | [***interface{}**](interface{}.md) | Returned to marketplaces as contract metadata | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

