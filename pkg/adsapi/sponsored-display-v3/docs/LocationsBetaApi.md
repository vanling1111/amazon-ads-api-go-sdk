# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveLocations**](LocationsBetaApi.md#ArchiveLocations) | **Post** /sd/locations/delete | Sets the &#x60;state&#x60; of each Location clause given to &#x60;archived&#x60;.
[**CreateLocations**](LocationsBetaApi.md#CreateLocations) | **Post** /sd/locations | Creates one or more locations associated with an ad group.
[**ListLocations**](LocationsBetaApi.md#ListLocations) | **Get** /sd/locations | Gets a list of locations associated with ad groups.

# **ArchiveLocations**
> []ArchiveLocationResponse ArchiveLocations(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Sets the `state` of each Location clause given to `archived`.

This is a bulk operation that accepts up to a limit of 1000 Location Expression Ids at a time. This deletion operation is equivalent to using the `updateTargetingClauses` operation to set the `state` property of a Location clause to `archived`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ArchiveLocationRequest**](ArchiveLocationRequest.md)| A list of up to 1000 Location Expression Ids for archival. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**[]ArchiveLocationResponse**](ArchiveLocationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, a-pplication/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateLocations**
> []Location CreateLocations(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Creates one or more locations associated with an ad group.

This resource is not available when productAds have ASIN or SKU fields and only available for advertisers that do not sell products on Amazon.   See [Developer Guide](https://advertising.amazon.com/API/docs/en-us/guides/sponsored-display/non-amazon-sellers/get-started)  Locations optimize Ad Groups for delivery to users that have an association with those locations. For example, an Ad Group might contain the following:  - A Targeting Clause representing an audience of users that viewed a shoe  - A Location representing Seattle, Washington, USA. - A Location representing New York, New York, USA. In this case, delivery of the Targeting Clause will be optimized for New York and Seattle.   You can discover predefined Locations to use in your ad group by calling the GET /locations API. The table below lists  several example Locations. | Location | Description | |---------------------------|-------------| | location=amzn1.ad-geo.XHvCjcKHXsKUwos= | Optimize the ad group for the specified location (either a 'city', 'state', 'dma', 'postal code', or 'country').|  Using locations is optional. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***LocationsBetaApiCreateLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationsBetaApiCreateLocationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []CreateLocation**](CreateLocation.md)| A list of up to 100 Locations for creation for each call.  1000 locations can be added for each ad group. | 

### Return type

[**[]Location**](Location.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, a-pplication/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLocations**
> []Location ListLocations(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of locations associated with ad groups.

Gets a list of Sponsored Display Location objects. This resource is not available when productAds have ASIN or SKU fields and only available for advertisers that do not sell products on Amazon. See [Developer Guide](https://advertising.amazon.com/API/docs/en-us/guides/sponsored-display/non-amazon-sellers/get-started)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***LocationsBetaApiListLocationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a LocationsBetaApiListLocationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Optional. 0-indexed record offset for the result set. Defaults to 0. | 
 **count** | **optional.Int32**| Optional. Number of records to include in the paged response. Defaults to max page size. | 
 **stateFilter** | **optional.String**| Optional. Restricts results to those with state within the specified comma-separated list. Must be one of: &#x60;enabled&#x60;. | [default to enabled]
 **adGroupIdFilter** | **optional.String**| Optional list of comma separated adGroupIds. Restricts results to locations with the specified &#x60;adGroupId&#x60;. | 
 **campaignIdFilter** | **optional.String**| Optional list of comma separated campaignIds. Restricts results to locations with the specified &#x60;campaignId&#x60;. | 

### Return type

[**[]Location**](Location.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

