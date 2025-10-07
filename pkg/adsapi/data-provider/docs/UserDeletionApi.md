# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2DpUsersPatch**](UserDeletionApi.md#V2DpUsersPatch) | **Patch** /v2/dp/users | Deletes user data originally sourced from the client. The API call rate is limited to 1 transactions per second (TPS). Calls exceeding this rate are throttled. Payload size is limited to 1000 users or 1MB. Calls with a more than 1000 users or 1MB will receive a 413 response.

# **V2DpUsersPatch**
> InlineResponse202 V2DpUsersPatch(ctx, body, authorization, amazonAdvertisingAPIClientID, contentType)
Deletes user data originally sourced from the client. The API call rate is limited to 1 transactions per second (TPS). Calls exceeding this rate are throttled. Payload size is limited to 1000 users or 1MB. Calls with a more than 1000 users or 1MB will receive a 413 response.

Deletes user data sourced from data providers. Deletes users scoped either to an advertiser or for the data provider. The SLA for data deletion is 30 days.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DpUsersBody**](DpUsersBody.md)|  | 
  **authorization** | **string**| A valid access token.  | 
  **amazonAdvertisingAPIClientID** | **string**| The client identifier. | 
  **contentType** | **string**| The Content-Type is &#x60;application/json&#x60;. | 

### Return type

[**InlineResponse202**](inline_response_202.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

