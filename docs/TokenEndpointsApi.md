# {{classname}}

All URIs are relative to *https://core-api.stardust.gg/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TokenBurnPost**](TokenEndpointsApi.md#TokenBurnPost) | **Post** /token/burn | Burn Token
[**TokenGetGet**](TokenEndpointsApi.md#TokenGetGet) | **Get** /token/get | Get Token
[**TokenMintBulkPost**](TokenEndpointsApi.md#TokenMintBulkPost) | **Post** /token/mint-bulk | Mint Tokens
[**TokenMutatePut**](TokenEndpointsApi.md#TokenMutatePut) | **Put** /token/mutate | Mutate Token
[**TokenPropsRemoveDelete**](TokenEndpointsApi.md#TokenPropsRemoveDelete) | **Delete** /token/props-remove | Remove Token Property
[**TokenTransferPost**](TokenEndpointsApi.md#TokenTransferPost) | **Post** /token/transfer | Transfer Tokens

# **TokenBurnPost**
> Empty TokenBurnPost(ctx, body)
Burn Token

Burns token on-chain. Cannot be reversed.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdTokenBurnRequest**](SdTokenBurnRequest.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TokenGetGet**
> []SdTokenGetResponse TokenGetGet(ctx, tokenIds)
Get Token

Some of the details of this token are based upon the Template that it was created from (using token/mint)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenIds** | **string**| Stringify Array of token ids (ex. &#x27;[3589, 3580]&#x27;) | 

### Return type

[**[]SdTokenGetResponse**](array.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TokenMintBulkPost**
> []int32 TokenMintBulkPost(ctx, body)
Mint Tokens

Mint tokens to a player.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdTokenMintBulkRequest**](SdTokenMintBulkRequest.md)|  | 

### Return type

[**[]int32**](array.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TokenMutatePut**
> Empty TokenMutatePut(ctx, body)
Mutate Token

Mutates a Property of a Token

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdTokenMutateRequest**](SdTokenMutateRequest.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TokenPropsRemoveDelete**
> Empty TokenPropsRemoveDelete(ctx, tokenId, props)
Remove Token Property

Removes a Tokens Property from Your Game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tokenId** | **string**| Token Id returned from token/create (ex. 5) | 
  **props** | **string**| Stringify Array of token mutable property names ex: &#x27;[\&quot;exp\&quot;, \&quot;health\&quot;]&#x27; } | 

### Return type

[**Empty**](Empty.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TokenTransferPost**
> Empty TokenTransferPost(ctx, body)
Transfer Tokens

Use this Endpoint to Facilitate Moving Tokens from one Player to Another

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdTokenTransferRequest**](SdTokenTransferRequest.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

