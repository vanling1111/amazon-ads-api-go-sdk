# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2DpAudiencemetadataAudienceIdGet**](MetadataApi.md#V2DpAudiencemetadataAudienceIdGet) | **Get** /v2/dp/audiencemetadata/{audienceId} | Gets metadata for an audience specified by identifier. Note that the API call rate is limited to 1 transaction per second (TPS). Calls exceeding this rate are throttled.
[**V2DpAudiencemetadataAudienceIdPut**](MetadataApi.md#V2DpAudiencemetadataAudienceIdPut) | **Put** /v2/dp/audiencemetadata/{audienceId} | Updates metadata of an existing audience specified by identifier. Note that the API call rate is limited to 1 transaction per second (TPS). Calls exceeding this rate are throttled.
[**V2DpAudiencemetadataPost**](MetadataApi.md#V2DpAudiencemetadataPost) | **Post** /v2/dp/audiencemetadata/ | Creates a new data provider audience. Note that the API call rate is limited to 1 transaction per second (TPS). Calls exceeding this rate are throttled.

# **V2DpAudiencemetadataAudienceIdGet**
> InlineResponse2001 V2DpAudiencemetadataAudienceIdGet(ctx, authorization, amazonAdvertisingAPIClientId, contentType, audienceId)
Gets metadata for an audience specified by identifier. Note that the API call rate is limited to 1 transaction per second (TPS). Calls exceeding this rate are throttled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| A valid access token. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **contentType** | **string**| The &#x60;Content-Type&#x60; is application/json. | 
  **audienceId** | **int64**| The audience identifier. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2DpAudiencemetadataAudienceIdPut**
> InlineResponse200 V2DpAudiencemetadataAudienceIdPut(ctx, body, authorization, amazonAdvertisingAPIClientId, contentType, audienceId)
Updates metadata of an existing audience specified by identifier. Note that the API call rate is limited to 1 transaction per second (TPS). Calls exceeding this rate are throttled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AudiencemetadataAudienceIdBody**](AudiencemetadataAudienceIdBody.md)|  | 
  **authorization** | **string**| A valid access token. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **contentType** | **string**| The &#x60;Content-Type&#x60; is application/json. | 
  **audienceId** | **int64**| The audience identifier. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V2DpAudiencemetadataPost**
> InlineResponse200 V2DpAudiencemetadataPost(ctx, body, authorization, amazonAdvertisingAPIClientId, contentType)
Creates a new data provider audience. Note that the API call rate is limited to 1 transaction per second (TPS). Calls exceeding this rate are throttled.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DpAudiencemetadataBody**](DpAudiencemetadataBody.md)|  | 
  **authorization** | **string**| A valid access token. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **contentType** | **string**| The &#x60;Content-Type&#x60; is application/json. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

