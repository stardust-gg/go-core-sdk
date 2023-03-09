# {{classname}}

All URIs are relative to *https://core-api.stardust.gg/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TemplateCountGet**](TemplateEndpointsApi.md#TemplateCountGet) | **Get** /template/count | Get Template Count
[**TemplateCreatePost**](TemplateEndpointsApi.md#TemplateCreatePost) | **Post** /template/create | Create Template
[**TemplateGetAllGet**](TemplateEndpointsApi.md#TemplateGetAllGet) | **Get** /template/get-all | Get All Templates
[**TemplateGetGet**](TemplateEndpointsApi.md#TemplateGetGet) | **Get** /template/get | Get Template
[**TemplateGetTokensGet**](TemplateEndpointsApi.md#TemplateGetTokensGet) | **Get** /template/get-tokens | Get Template Tokens
[**TemplateMutatePut**](TemplateEndpointsApi.md#TemplateMutatePut) | **Put** /template/mutate | Mutate Template
[**TemplatePropsRemoveDelete**](TemplateEndpointsApi.md#TemplatePropsRemoveDelete) | **Delete** /template/props-remove | Remove Template Property
[**TemplateRemoveDelete**](TemplateEndpointsApi.md#TemplateRemoveDelete) | **Delete** /template/remove | Remove Template

# **TemplateCountGet**
> SdTemplateCountResponse TemplateCountGet(ctx, optional)
Get Template Count

Get a Template's Count Within a Game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***TemplateEndpointsApiTemplateCountGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TemplateEndpointsApiTemplateCountGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **optional.String**| Find a substring within the Template name field | 

### Return type

[**SdTemplateCountResponse**](SDTemplateCountResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TemplateCreatePost**
> SdTemplateCreateResponse TemplateCreatePost(ctx, body)
Create Template

Adds a New Token Template that lets Players Acquire that Token using the Token/Mint Endpoint

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdTemplateCreateRequest**](SdTemplateCreateRequest.md)|  | 

### Return type

[**SdTemplateCreateResponse**](SDTemplateCreateResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TemplateGetAllGet**
> []SdTemplateGetAllResponse TemplateGetAllGet(ctx, start, limit, optional)
Get All Templates

Get All of the Templates Within a Game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **start** | **string**| position in list | 
  **limit** | **string**| maximum templates returned in list | 
 **optional** | ***TemplateEndpointsApiTemplateGetAllGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TemplateEndpointsApiTemplateGetAllGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **optional.String**| Find a substring within the Template name field | 

### Return type

[**[]SdTemplateGetAllResponse**](array.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TemplateGetGet**
> SdTemplateGetResponse TemplateGetGet(ctx, templateId)
Get Template

Get the Details of a Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**| Template Id such as 3589) | 

### Return type

[**SdTemplateGetResponse**](SDTemplateGetResponse.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TemplateGetTokensGet**
> []SdTemplateGetTokensResponse TemplateGetTokensGet(ctx, templateId, optional)
Get Template Tokens

Get a List of All Minted Tokens from a Given Template

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**| The ID of the Template | 
 **optional** | ***TemplateEndpointsApiTemplateGetTokensGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TemplateEndpointsApiTemplateGetTokensGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **start** | **optional.String**| position in list | 
 **limit** | **optional.String**| maximum templates returned in list | 
 **includeDeleted** | **optional.String**| If true tokens from deleted players are included in the response | 

### Return type

[**[]SdTemplateGetTokensResponse**](array.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TemplateMutatePut**
> Empty TemplateMutatePut(ctx, body)
Mutate Template

Mutates a Property of a Template, Which in Turn will Affect the Inherited Property on All of the Tokens Created from this Template (via token/mint)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SdTemplateMutateRequest**](SdTemplateMutateRequest.md)|  | 

### Return type

[**Empty**](Empty.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TemplatePropsRemoveDelete**
> Empty TemplatePropsRemoveDelete(ctx, templateId, props)
Remove Template Property

Removes a Templates Property from Your Game

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**| Template Id returned from template/create (ex. 5) | 
  **props** | **string**| Stringify Array of template mutable property names ex: &#x27;[\&quot;exp\&quot;, \&quot;health\&quot;]&#x27; } | 

### Return type

[**Empty**](Empty.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TemplateRemoveDelete**
> Empty TemplateRemoveDelete(ctx, templateId)
Remove Template

Removes a Template from Your Game. If Players have Instances of this Template from the token/mint Command, Their Tokens will NOT be Removed

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **templateId** | **string**| Template Id returned from template/create (ex. 5) | 

### Return type

[**Empty**](Empty.md)

### Authorization

[api_key](../README.md#api_key)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

