# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchTaxonomy**](DiscoveryApi.md#FetchTaxonomy) | **Post** /audiences/taxonomy/list | Browse the taxonomy of audience categories
[**ListAudiences**](DiscoveryApi.md#ListAudiences) | **Post** /audiences/list | Gets audience segments based on filters

# **FetchTaxonomy**
> FetchTaxonomyResponseV1 FetchTaxonomy(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Browse the taxonomy of audience categories

Returns a list of audience categories for a given category path  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Advertising-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**FetchTaxonomyRequestBodyV1**](FetchTaxonomyRequestBodyV1.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***DiscoveryApiFetchTaxonomyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiscoveryApiFetchTaxonomyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **advertiserId** | **optional.**| The advertiser associated with the advertising account. This parameter is required for the DSP adType, but optional for the SD adType. | 
 **nextToken** | **optional.**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.**| Sets the maximum number of categories in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | 

### Return type

[**FetchTaxonomyResponseV1**](FetchTaxonomyResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAudiences**
> ListAudiencesResponseV1 ListAudiences(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets audience segments based on filters

Returns a list of audience segments for an advertiser. The result set can be filtered by providing an array of Filter objects. Each item in the resulting set will match all specified filters.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Advertising-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***DiscoveryApiListAudiencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DiscoveryApiListAudiencesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ListAudiencesRequestBodyV1**](ListAudiencesRequestBodyV1.md)|  | 
 **advertiserId** | **optional.**| The advertiser to retrieve segments for. This parameter is required for the DSP adType, but optional for the SD adType. | 
 **canTarget** | **optional.**| When set to true, only targetable audience segments will be returned. | [default to false]
 **nextToken** | **optional.**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.**| Sets the maximum number of audiences in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | 

### Return type

[**ListAudiencesResponseV1**](ListAudiencesResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

