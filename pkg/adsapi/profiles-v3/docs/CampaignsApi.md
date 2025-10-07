# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveCampaign**](CampaignsApi.md#ArchiveCampaign) | **Delete** /sd/campaigns/{campaignId} | Sets the campaign status to archived.
[**CreateCampaigns**](CampaignsApi.md#CreateCampaigns) | **Post** /sd/campaigns | Creates one or more campaigns.
[**GetCampaign**](CampaignsApi.md#GetCampaign) | **Get** /sd/campaigns/{campaignId} | Gets a requested campaign.
[**GetCampaignResponseEx**](CampaignsApi.md#GetCampaignResponseEx) | **Get** /sd/campaigns/extended/{campaignId} | Gets extended information for a requested campaign.
[**ListCampaigns**](CampaignsApi.md#ListCampaigns) | **Get** /sd/campaigns | Gets a list of campaigns.
[**ListCampaignsEx**](CampaignsApi.md#ListCampaignsEx) | **Get** /sd/campaigns/extended | Gets a list of campaigns with extended fields.
[**UpdateCampaigns**](CampaignsApi.md#UpdateCampaigns) | **Put** /sd/campaigns | Updates one or more campaigns.

# **ArchiveCampaign**
> CampaignResponse ArchiveCampaign(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, campaignId)
Sets the campaign status to archived.

This operation is equivalent to an update operation that sets the status field to 'archived'. Note that setting the status field to 'archived' is permanent and can't be undone. See [Developer Notes](https://advertising.amazon.com/API/docs/en-us/info/developer-notes#archiving) for more information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **campaignId** | **int64**| The identifier of the campaign. | 

### Return type

[**CampaignResponse**](CampaignResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCampaigns**
> []CampaignResponse CreateCampaigns(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Creates one or more campaigns.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***CampaignsApiCreateCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiCreateCampaignsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []CreateCampaign**](CreateCampaign.md)| An array of Campaign objects. For each object, specify required fields and their values. Required fields are &#x60;name&#x60;, &#x60;tactic&#x60;, &#x60;state&#x60;, and &#x60;startDate&#x60;. Maximum length of the array is 100 objects. If you don&#x27;t specify a &#x60;budget&#x60;, it will be set as the [default budget for your region](https://advertising.amazon.com/API/docs/en-us/concepts/limits#default-budgets). Campaign names must be unique across SD, SB, and SP.
  If you are using Optimization rules, the following campaign budget must be at least:
  - 5x the value of any COST_PER_ORDER threshold.
  - 10x the value of any COST_PER_THOUSAND_VIEWABLE_IMPRESSIONS threshold.
  - 20x the value of any COST_PER_CLICK threshold.
 | 

### Return type

[**[]CampaignResponse**](CampaignResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaign**
> Campaign GetCampaign(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, campaignId)
Gets a requested campaign.

Returns a Campaign object for a requested campaign. Note that the Campaign object is designed for performance, with a small set of commonly used campaign fields to reduce size. If the extended set of fields is required, use the campaign operations that return the CampaignResponseEx object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **campaignId** | **int64**| The identifier of the requested campaign. | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignResponseEx**
> CampaignResponseEx GetCampaignResponseEx(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, campaignId)
Gets extended information for a requested campaign.

Returns a CampaignResponseEx object for a requested campaign. The CampaignResponseEx includes the extended set of available fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **campaignId** | **int64**| The identifier of the requested campaign. | 

### Return type

[**CampaignResponseEx**](CampaignResponseEx.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCampaigns**
> []Campaign ListCampaigns(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of campaigns.

Gets an array of Campaign objects for a requested set of Sponsored Display campaigns. Note that the Campaign object is designed for performance, and includes a small set of commonly used fields to reduce size. If the extended set of fields is required, use the campaign operations that return the CampaignResponseEx object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***CampaignsApiListCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiListCampaignsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Optional. Sets a cursor into the requested set of campaigns. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0. | [default to 0]
 **count** | **optional.Int32**| Optional. Sets the number of Campaign objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten campaigns set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten campaigns, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size. | 
 **stateFilter** | **optional.String**| Optional. The returned array is filtered to include only campaigns with state set to one of the values in the specified comma-delimited list. | [default to enabled, paused, archived]
 **name** | **optional.String**| Optional. The returned array includes only campaign with the specified name using an exact string match. | 
 **campaignIdFilter** | **optional.String**| Optional. The returned array includes only campaigns with identifiers matching those specified in the comma-delimited string. | 
 **portfolioIdFilter** | **optional.String**| Optional. The returned array includes only campaigns associated with Portfolio identifiers matching those specified in the comma-delimited string. | 

### Return type

[**[]Campaign**](Campaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListCampaignsEx**
> []CampaignResponseEx ListCampaignsEx(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of campaigns with extended fields.

Gets an array of CampaignResponseEx objects for a set of requested campaigns.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***CampaignsApiListCampaignsExOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiListCampaignsExOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Optional. Sets a cursor into the requested set of campaigns. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0. | 
 **count** | **optional.Int32**| Optional. Sets the number of Campaign objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten campaigns set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten campaigns, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size. | 
 **stateFilter** | **optional.String**| Optional. The returned array is filtered to include only campaigns with state set to one of the values in the specified comma-delimited list. | [default to enabled, paused, archived]
 **name** | **optional.String**| Optional. The returned array includes only campaign with the specified name using an exact string match. | 
 **campaignIdFilter** | **optional.String**| Optional. The returned array includes only campaigns with identifiers matching those specified in the comma-delimited string. | 
 **portfolioIdFilter** | **optional.String**| Optional. The returned array includes only campaigns associated with Portfolio identifiers matching those specified in the comma-delimited string. | 

### Return type

[**[]CampaignResponseEx**](CampaignResponseEx.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCampaigns**
> []CampaignResponse UpdateCampaigns(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Updates one or more campaigns.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***CampaignsApiUpdateCampaignsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CampaignsApiUpdateCampaignsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []UpdateCampaign**](UpdateCampaign.md)| An array of Campaign objects. For each object, specify a campaign identifier and mutable fields with their updated values. The mutable fields are &#x60;name&#x60;, &#x60;state&#x60;, &#x60;budget&#x60;, &#x60;startDate&#x60;, and &#x60;endDate&#x60;. Maximum length of the array is 100 objects. | 

### Return type

[**[]CampaignResponse**](CampaignResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

