# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetHistory**](DefaultApi.md#GetHistory) | **Post** /history | History of entity changes.

# **GetHistory**
> InlineResponse200 GetHistory(ctx, body)
History of entity changes.

Returns history of changes for provided event sources that match the filters and time ranges specified. Only events that belong to the authenticated Advertiser can be queried. All times will be in UTC Epoch format. This API accepts identifiers in either the alphamumeric format (default), or the numeric format. If numeric IDs are supplied, then numeric IDs will be returned otherwise, alphanumeric IDs are returned. This API only returns change history for Sponsored Products and Sponsored Brands campaigns. Changes for Sponsored Display campaigns are not returned. Also, it doesn't return information about who made the change.  **API Versioning:** This API supports versioning through the Accept header. Use `Accept: application/vnd.historyresponse.v1.1+json` to access version 1.1 features. **THEME event type requires version 1.1 or higher.** Without the proper Accept header, THEME requests will return a 400 error.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**HistoryBody**](HistoryBody.md)| HistoryQuery | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

