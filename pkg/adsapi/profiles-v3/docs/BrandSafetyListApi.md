# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrandSafetyDenyListDomains**](BrandSafetyListApi.md#CreateBrandSafetyDenyListDomains) | **Post** /sd/brandSafety/deny | Creates one or more domains to add to a Brand Safety Deny List. 
[**DeleteBrandSafetyDenyList**](BrandSafetyListApi.md#DeleteBrandSafetyDenyList) | **Delete** /sd/brandSafety/deny | Archives all of the domains in the Brand Safety Deny List. 
[**GetRequestResults**](BrandSafetyListApi.md#GetRequestResults) | **Get** /sd/brandSafety/{requestId}/results | Gets the results for the given request
[**GetRequestStatus**](BrandSafetyListApi.md#GetRequestStatus) | **Get** /sd/brandSafety/{requestId}/status | Gets the status of the given request
[**ListDomains**](BrandSafetyListApi.md#ListDomains) | **Get** /sd/brandSafety/deny | Gets a list of websites/apps that are on the advertiser&#x27;s Brand Safety Deny List.
[**ListRequestStatus**](BrandSafetyListApi.md#ListRequestStatus) | **Get** /sd/brandSafety/status | List status of all requests

# **CreateBrandSafetyDenyListDomains**
> BrandSafetyUpdateResponse CreateBrandSafetyDenyListDomains(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Creates one or more domains to add to a Brand Safety Deny List. 

Creates one or more domains to add to a Brand Safety Deny List. The Brand Safety Deny List is at the advertiser level. It can take up to 15 minutes from the time a domain is added to the time it is reflected in the deny list. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BrandSafetyPostRequest**](BrandSafetyPostRequest.md)| An array of Brand Safety List Domain objects. For each object, specify required fields and their values. Maximum length of the array is 10,000 objects.
 | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**BrandSafetyUpdateResponse**](BrandSafetyUpdateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteBrandSafetyDenyList**
> BrandSafetyUpdateResponse DeleteBrandSafetyDenyList(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Archives all of the domains in the Brand Safety Deny List. 

Archives all of the domains in the Brand Safety Deny List. It can take several hours from the time a domain is deleted to the time it is reflected in the deny list. You can check the status of the delete request by calling GET /sd/brandSafety/{requestId}/status. If the status is \"COMPLETED\", you can call GET /sd/brandSafety/deny to validate that your deny list has been successfully deleted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**BrandSafetyUpdateResponse**](BrandSafetyUpdateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRequestResults**
> BrandSafetyRequestResultsResponse GetRequestResults(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, requestId, optional)
Gets the results for the given request

When a user adds domains to their Brand Safety Deny List, the request is processed asynchronously, and a requestId is provided to the user. This requestId can be used to view the request results for up to 90 days from when the request was submitted. The results provide the status of each domain in the given request. Request results may contain multiple pages. This endpoint will only be available once the request has completed processing. To see the status of the request you can call GET /sd/brandSafety/{requestId}/status. Note that this endpoint only lists the results of POST requests to /sd/brandSafety/deny - it does not reflect the results of DELETE requests. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **requestId** | **string**| The ID of the request previously submitted. | 
 **optional** | ***BrandSafetyListApiGetRequestResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandSafetyListApiGetRequestResultsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **startIndex** | **optional.Int32**| Optional. Sets a cursor into the requested set of results. Use in conjunction with the count parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0.  | 
 **count** | **optional.Int32**| Optional. Sets the number of results in the returned array. Use in conjunction with the startIndex parameter to control pagination. For example, to return the first 1000 results set startIndex&#x3D;0 and count&#x3D;1000. To return the next 1000 results, set startIndex&#x3D;1000 and count&#x3D;1000, and so on. Defaults to max page size(1000).  | 

### Return type

[**BrandSafetyRequestResultsResponse**](BrandSafetyRequestResultsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetRequestStatus**
> BrandSafetyRequestStatusResponse GetRequestStatus(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, requestId)
Gets the status of the given request

When a user modifies their Brand Safety Deny List, the request is processed asynchronously, and a requestId is provided to the user. This requestId can be used to check the status of the request for up to 90 days from when the request was submitted. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **requestId** | **string**| The ID of the request previously submitted. | 

### Return type

[**BrandSafetyRequestStatusResponse**](BrandSafetyRequestStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDomains**
> BrandSafetyGetResponse ListDomains(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of websites/apps that are on the advertiser's Brand Safety Deny List.

Gets an array of websites/apps that are on the advertiser's Brand Safety Deny List. It can take up to 15 minutes from the time a domain is added/deleted to the time it is reflected in the deny list.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***BrandSafetyListApiListDomainsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandSafetyListApiListDomainsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Optional. Sets a cursor into the requested set of domains. Use in conjunction with the count parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0.  | 
 **count** | **optional.Int32**| Optional. Sets the number of domain objects in the returned array. Use in conjunction with the startIndex parameter to control pagination. For example, to return the first 1000 domains set startIndex&#x3D;0 and count&#x3D;1000. To return the next 1000 domains, set startIndex&#x3D;1000 and count&#x3D;1000, and so on. Defaults to max page size(1000).  | 

### Return type

[**BrandSafetyGetResponse**](BrandSafetyGetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRequestStatus**
> BrandSafetyListRequestStatusResponse ListRequestStatus(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
List status of all requests

List status of all Brand Safety List requests. The list will contain requests that were submitted in the past 90 days. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**BrandSafetyListRequestStatusResponse**](BrandSafetyListRequestStatusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

