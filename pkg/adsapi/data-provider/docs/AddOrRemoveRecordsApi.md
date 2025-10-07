# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V2DpAudiencePatch**](AddOrRemoveRecordsApi.md#V2DpAudiencePatch) | **Patch** /v2/dp/audience | Associates or disassociates a record with an audience. Note that the API call rate is limited to 100 transactions per second (TPS). Calls exceeding this rate are throttled. Payload size is limited to a maximum of 6MB or 2000 records. Calls with a payload exceeding limit receive a 413 response.

# **V2DpAudiencePatch**
> InlineResponse2002 V2DpAudiencePatch(ctx, body, authorization, amazonAdvertisingAPIClientId, contentType)
Associates or disassociates a record with an audience. Note that the API call rate is limited to 100 transactions per second (TPS). Calls exceeding this rate are throttled. Payload size is limited to a maximum of 6MB or 2000 records. Calls with a payload exceeding limit receive a 413 response.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DpAudienceBody**](DpAudienceBody.md)|  | 
  **authorization** | **string**| A valid access token. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **contentType** | **string**| The &#x60;Content-Type&#x60; is application/json. | 

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

