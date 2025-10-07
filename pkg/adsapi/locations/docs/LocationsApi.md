# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListLocations**](LocationsApi.md#ListLocations) | **Post** /locations/list | Gets location objects based on one or more filters.

# **ListLocations**
> InlineResponse200 ListLocations(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets location objects based on one or more filters.

Note: This endpoint is currently limited to US only. Gets a list of location objects after filtering on at least one of **locationId**, **name**, **category**. Each item in the resulting set will match all specified filters.   **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***LocationsApiListLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationsApiListLocationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ListLocationsRequestBodyV1**](ListLocationsRequestBodyV1.md)|  | 
 **nextToken** | **optional.**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.**| Sets the number of locations in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,2000] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch the next 20 items. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

