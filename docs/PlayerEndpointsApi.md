# {{classname}}

All URIs are relative to *https://core-api.stardust.gg/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlayerCountGet**](PlayerEndpointsApi.md#PlayerCountGet) | **Get** /player/count | Get Player Count
[**PlayerCreatePost**](PlayerEndpointsApi.md#PlayerCreatePost) | **Post** /player/create | Create Player
[**PlayerGetAllGet**](PlayerEndpointsApi.md#PlayerGetAllGet) | **Get** /player/get-all | Get All Players
[**PlayerGetGet**](PlayerEndpointsApi.md#PlayerGetGet) | **Get** /player/get | Get Player
[**PlayerGetIdGet**](PlayerEndpointsApi.md#PlayerGetIdGet) | **Get** /player/get-id | Get Player ID
[**PlayerGetIdsGet**](PlayerEndpointsApi.md#PlayerGetIdsGet) | **Get** /player/get-ids | Get All Player IDs
[**PlayerGetInventoryGet**](PlayerEndpointsApi.md#PlayerGetInventoryGet) | **Get** /player/get-inventory | Get Player Inventory
[**PlayerMutatePut**](PlayerEndpointsApi.md#PlayerMutatePut) | **Put** /player/mutate | Mutate Player
[**PlayerRemoveDelete**](PlayerEndpointsApi.md#PlayerRemoveDelete) | **Delete** /player/remove | Remove Player
[**PlayerWalletGetGet**](PlayerEndpointsApi.md#PlayerWalletGetGet) | **Get** /player/wallet-get | Get Player Wallet
[**PlayerWithdrawPost**](PlayerEndpointsApi.md#PlayerWithdrawPost) | **Post** /player/withdraw | Withdraw From Player

# **PlayerCountGet**
> SdPlayerCountResponse PlayerCountGet(ctx, )
Get Player Count

Get Player count within a game

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SdPlayerCountResponse**](SDPlayerCountResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlayerCreatePost**
> SdPlayerCreateResponse PlayerCreatePost(ctx, body)
Create Player

Create a Player for a game. Returns their player id which can be used to reference them later in Stardust's system

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdPlayerCreateRequest**](SdPlayerCreateRequest.md)|  | 

### Return type

[**SdPlayerCreateResponse**](SDPlayerCreateResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlayerGetAllGet**
> []map[string]string PlayerGetAllGet(ctx, optional)
Get All Players

Get the List of All Players in Game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PlayerEndpointsApiPlayerGetAllGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlayerEndpointsApiPlayerGetAllGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **start** | **optional.String**| position in list | 
 **filter** | **optional.String**| Find a substring within the Player unique ID field | 
 **limit** | **optional.String**| maximum items returned in list | 

### Return type

[**[]map[string]string**](array.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlayerGetGet**
> map[string]string PlayerGetGet(ctx, playerId)
Get Player

Get Details of a Player Within a Game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **string**| Player id | 

### Return type

[**map[string]string**](map.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlayerGetIdGet**
> SdPlayerGetIdResponse PlayerGetIdGet(ctx, uniqueId)
Get Player ID

Get a Player's ID via their Unique ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uniqueId** | **string**| Player&#x27;s Unique ID | 

### Return type

[**SdPlayerGetIdResponse**](SDPlayerGetIdResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlayerGetIdsGet**
> []SdPlayerGetIdsResponse PlayerGetIdsGet(ctx, )
Get All Player IDs

Get All Player IDs for a Given Game

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SdPlayerGetIdsResponse**](array.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlayerGetInventoryGet**
> []SdPlayerGetInventoryResponse PlayerGetInventoryGet(ctx, playerId, optional)
Get Player Inventory

Get a players inventory and all the items it holds

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **string**| Player id | 
 **optional** | ***PlayerEndpointsApiPlayerGetInventoryGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PlayerEndpointsApiPlayerGetInventoryGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| position in list | 
 **tokenIds** | **optional.String**| Comma-Separated String of token ids (ex. &#x27;3589, 3580&#x27;) | 
 **limit** | **optional.String**| maximum items returned in list | 

### Return type

[**[]SdPlayerGetInventoryResponse**](array.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlayerMutatePut**
> Empty PlayerMutatePut(ctx, body)
Mutate Player

Change player data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdPlayerMutateRequest**](SdPlayerMutateRequest.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlayerRemoveDelete**
> Empty PlayerRemoveDelete(ctx, playerId)
Remove Player

Removes (hides) a player from your game. This is not permanent.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **string**| Player Id returned from player/create a UUID, i.e. 802760b0-2bb5-4fba-9237-895ed02cf8d8 | 

### Return type

[**Empty**](Empty.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlayerWalletGetGet**
> SdPlayerWalletGetResponse PlayerWalletGetGet(ctx, playerId)
Get Player Wallet

Get player's wallet within a game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playerId** | **string**| Player id | 

### Return type

[**SdPlayerWalletGetResponse**](SDPlayerWalletGetResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlayerWithdrawPost**
> Empty PlayerWithdrawPost(ctx, body)
Withdraw From Player

Withdraw a Player's Tokens from their Stardust Wallet

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdPlayerWithdrawRequest**](SdPlayerWithdrawRequest.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

