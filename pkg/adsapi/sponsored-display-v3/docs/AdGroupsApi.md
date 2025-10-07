# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveAdGroup**](AdGroupsApi.md#ArchiveAdGroup) | **Delete** /sd/adGroups/{adGroupId} | Sets the ad group status to archived.
[**CreateAdGroups**](AdGroupsApi.md#CreateAdGroups) | **Post** /sd/adGroups | Creates one or more ad groups.
[**GetAdGroup**](AdGroupsApi.md#GetAdGroup) | **Get** /sd/adGroups/{adGroupId} | Gets a requested ad group.
[**GetAdGroupResponseEx**](AdGroupsApi.md#GetAdGroupResponseEx) | **Get** /sd/adGroups/extended/{adGroupId} | Gets extended information for a requested ad group.
[**ListAdGroups**](AdGroupsApi.md#ListAdGroups) | **Get** /sd/adGroups | Gets a list of ad groups.
[**ListAdGroupsEx**](AdGroupsApi.md#ListAdGroupsEx) | **Get** /sd/adGroups/extended | Gets a list of ad groups with extended fields.
[**UpdateAdGroups**](AdGroupsApi.md#UpdateAdGroups) | **Put** /sd/adGroups | Updates on or more ad groups.

# **ArchiveAdGroup**
> AdGroupResponse ArchiveAdGroup(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, adGroupId)
Sets the ad group status to archived.

This operation is equivalent to an update operation that sets the status field to 'archived'. Note that setting the status field to 'archived' is permanent and can't be undone. See [Developer Notes](https://advertising.amazon.com/API/docs/en-us/info/developer-notes#archiving) for more information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **adGroupId** | **int64**| The identifier of the requested ad group. | 

### Return type

[**AdGroupResponse**](AdGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateAdGroups**
> []AdGroupResponse CreateAdGroups(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Creates one or more ad groups.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***AdGroupsApiCreateAdGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiCreateAdGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []CreateAdGroup**](CreateAdGroup.md)| An array of AdGroup objects. For each object, specify required fields and their values. Required fields are &#x60;campaignId&#x60;, &#x60;name&#x60;, &#x60;state&#x60;, and &#x60;defaultBid&#x60;. Maximum length of the array is 100 objects. Note - when using landingPageType of OFF_AMAZON_LINK or STORES within productAds, only 1 adGroup is supported. | 

### Return type

[**[]AdGroupResponse**](AdGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdGroup**
> AdGroup GetAdGroup(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, adGroupId)
Gets a requested ad group.

Returns an AdGroup object for a requested campaign. Note that the AdGroup object is designed for performance, with a small set of commonly used ad group fields to reduce size. If the extended set of fields is required, use the campaign operations that return the AdGroupResponseEx object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **adGroupId** | **int64**| The identifier of the requested ad group. | 

### Return type

[**AdGroup**](AdGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAdGroupResponseEx**
> AdGroupResponseEx GetAdGroupResponseEx(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, adGroupId)
Gets extended information for a requested ad group.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **adGroupId** | **int64**| The identifier of the requested ad group. | 

### Return type

[**AdGroupResponseEx**](AdGroupResponseEx.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAdGroups**
> []AdGroup ListAdGroups(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of ad groups.

Gets an array of AdGroup objects for a requested set of Sponsored Display ad groups. Note that the AdGroup object is designed for performance, and includes a small set of commonly used fields to reduce size. If the extended set of fields is required, use the ad group operations that return the AdGroupResponseEx object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***AdGroupsApiListAdGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiListAdGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Optional. Sets a cursor into the requested set of campaigns. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0. | 
 **count** | **optional.Int32**| Optional. Sets the number of AdGroup objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten ad groups set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten ad groups, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size. | 
 **stateFilter** | **optional.String**| Optional. The returned array is filtered to include only ad groups with state set to one of the values in the specified comma-delimited list. | [default to enabled, paused, archived]
 **campaignIdFilter** | **optional.String**| Optional. The returned array is filtered to include only ad groups associated with the campaign identifiers in the specified comma-delimited list. | 
 **adGroupIdFilter** | **optional.String**| Optional. The returned array is filtered to include only ad groups with an identifier specified in the comma-delimited list. | 
 **name** | **optional.String**| Optional. The returned array includes only ad groups with the specified name. | 

### Return type

[**[]AdGroup**](AdGroup.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAdGroupsEx**
> []AdGroupResponseEx ListAdGroupsEx(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of ad groups with extended fields.

Gets an array of AdGroupResponseEx objects for a set of requested ad groups.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***AdGroupsApiListAdGroupsExOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiListAdGroupsExOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Optional. Sets a cursor into the requested set of ad groups. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0. | 
 **count** | **optional.Int32**| Optional. Sets the number of Campaign objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten campaigns set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten campaigns, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size. | 
 **stateFilter** | **optional.String**| Optional. The returned array is filtered to include only campaigns with state set to one of the values in the comma-delimited list. | [default to enabled, paused, archived]
 **campaignIdFilter** | **optional.String**| Optional. The returned array is filtered to include only ad groups associated with the campaign identifiers in the comma-delimited list. | 
 **adGroupIdFilter** | **optional.String**| Optional. The returned array is filtered to include only ad groups with an identifier specified in the comma-delimited list. | 
 **name** | **optional.String**| Optional. The returned array includes only ad groups with the specified name. | 

### Return type

[**[]AdGroupResponseEx**](AdGroupResponseEx.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAdGroups**
> []AdGroupResponse UpdateAdGroups(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Updates on or more ad groups.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***AdGroupsApiUpdateAdGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdGroupsApiUpdateAdGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []UpdateAdGroup**](UpdateAdGroup.md)| An array of AdGroup objects. For each object, specify an ad group identifier and mutable fields with their updated values. The mutable fields are &#x27;name&#x27;, &#x27;defaultBid&#x27;, and &#x27;state&#x27;. Maximum length of the array is 100 objects. | 

### Return type

[**[]AdGroupResponse**](AdGroupResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

