# {{classname}}

All URIs are relative to *https://core-api.stardust.gg/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GameGetGet**](GameEndpointsApi.md#GameGetGet) | **Get** /game/get | Get Game
[**GameMutatePut**](GameEndpointsApi.md#GameMutatePut) | **Put** /game/mutate | Mutate Game

# **GameGetGet**
> SdGameGetResponse GameGetGet(ctx, )
Get Game

Get the Details of Your Game

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**SdGameGetResponse**](SDGameGetResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GameMutatePut**
> Empty GameMutatePut(ctx, body)
Mutate Game

Change a games data

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdGameMutateRequest**](SdGameMutateRequest.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

