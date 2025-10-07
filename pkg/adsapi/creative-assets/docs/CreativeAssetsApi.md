# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssetsBatchRegister**](CreativeAssetsApi.md#AssetsBatchRegister) | **Post** /assets/batchRegister | Batch register uploaded assets
[**GetAsset**](CreativeAssetsApi.md#GetAsset) | **Get** /assets | Retrieve an asset
[**GetAssetsBatchRegister**](CreativeAssetsApi.md#GetAssetsBatchRegister) | **Get** /assets/batchRegister/{requestId} | Retrieves status of the batch asset registration request, uniquely identified by requestId.
[**GetUploadLocation**](CreativeAssetsApi.md#GetUploadLocation) | **Post** /assets/upload | Create an upload location
[**RegisterAsset**](CreativeAssetsApi.md#RegisterAsset) | **Post** /assets/register | Register an uploaded asset
[**SearchAssets**](CreativeAssetsApi.md#SearchAssets) | **Post** /assets/search | Search assets

# **AssetsBatchRegister**
> BatchRegisterAssetResponse AssetsBatchRegister(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Batch register uploaded assets

This is an asynchronous api that provides clients an identifier for their batch registration request. They need to check for the status of their request by calling the `GET /assets/batchRegister/{requestId}` api with the identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AssetsBatchRegisterRequest**](AssetsBatchRegisterRequest.md)| Batch Register Uploaded Assets. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**BatchRegisterAssetResponse**](BatchRegisterAssetResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.assetsbatchregisterrequest.v1+json
 - **Accept**: application/vnd.assetsbatchregisterresponse.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAsset**
> InlineResponse200 GetAsset(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, assetId, optional)
Retrieve an asset

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **assetId** | **string**| The assetId | 
 **optional** | ***CreativeAssetsApiGetAssetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreativeAssetsApiGetAssetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **version** | **optional.String**| The versionId of the asset, if not included all versions will return. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.creativeassetsgetresponse.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAssetsBatchRegister**
> GetBatchRegisterAssetStatusOutput GetAssetsBatchRegister(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, requestId)
Retrieves status of the batch asset registration request, uniquely identified by requestId.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **requestId** | **string**| Batch asset registration request id. It is returned in response of POST /assets/batchRegister API. | 

### Return type

[**GetBatchRegisterAssetStatusOutput**](GetBatchRegisterAssetStatusOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.creativeassetsgetbatchregisterresponse.v3+json, application/vnd.assetsbatchregisterresponse.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUploadLocation**
> InlineResponse2002 GetUploadLocation(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Create an upload location

Creates an ephemeral resource (upload location) to upload Assets to Creative Assets tool. The upload location is short lived and expires in 15 minutes.The upload location only supports PUT HTTP Method to upload the asset content. If the upload location expires, API user will get `403` Forbidden response. * All ad specs - sizes and policies can be found [here](https://advertising.amazon.com/resources/ad-specs/?ref_=a20m_us_hnav_spcs)  * Program specific links 1. **Stores** - [here](https://advertising.amazon.com/resources/ad-specs/stores?ref_=a20m_us_spcs_stcrgd) 2. **SB/SBV/sponsored ads** - [here](https://advertising.amazon.com/resources/ad-policy/sponsored-ads-policies?ref_=a20m_us_spcs_sbv_spcs_spadcap) See [Creating assets](guides/creative-asset/creating-assets) to understand the call flow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AssetsUploadBody**](AssetsUploadBody.md)| Make sure to include file extension along with filename . (ie. &quot;filename.mp4&quot;) | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.creativeassetsuploadresponse.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterAsset**
> InlineResponse2001 RegisterAsset(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Register an uploaded asset

The API should be called once the asset is uploaded to the location provided by the /asset/upload API endpoint. See [Creating assets](guides/creative-asset/creating-assets) to understand the call flow.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AssetsRegisterBody**](AssetsRegisterBody.md)| Note - **asinList**, **versionInfo**, and **tags** are optional in the request body. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.creativeassetsregisterresponse.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchAssets**
> InlineResponse2003 SearchAssets(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Search assets

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CaSearchRequestCommon**](CaSearchRequestCommon.md)| text - optional, this field matches asset name, asset name prefix, tags and ASINs associated with the assets

filterCriteria - optional, this is used to filter results

sortCriteria - optional, this is used to get sorted results

pageCriteria - optional, this is used for pagination
 | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.creativeassetssearchassetsresponse.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

