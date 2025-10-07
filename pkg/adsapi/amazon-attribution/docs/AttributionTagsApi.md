# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPublisherAttributionTagTemplate**](AttributionTagsApi.md#GetPublisherAttributionTagTemplate) | **Get** /attribution/tags/macroTag | Gets a list of attribution tags for third-party publisher campaigns that support macros.
[**GetPublisherMacroAttributionTag**](AttributionTagsApi.md#GetPublisherMacroAttributionTag) | **Get** /attribution/tags/nonMacroTemplateTag | Gets a list of attribution tags for third-party publisher campaigns that do not support macros.

# **GetPublisherAttributionTagTemplate**
> AttributionTagResponse GetPublisherAttributionTagTemplate(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, publisherIds, optional)
Gets a list of attribution tags for third-party publisher campaigns that support macros.

Third-party publishers, such as Google Ads, Facebook, Microsoft Ads, and Pinterest support tags that include macro parameters. Using macro parameters, campaign tracking information is dynamically inserted into the click-through URL when an ad is clicked. This resource is a tag pre-populated with campaign, ad group, and ad level publisher macros with the values associated with your campaign. <br/><br/> For example, a Google Ads macro tag is \"?maas=maas_adg_api_123456789_1_99&ref_=aa_maas&tag=maas&aa_campaignid={campaignid}&aa_adgroupid={adgroupid}&aa_creativeid=ad-{creative}_{targetid}\"

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **publisherIds** | [**[]string**](string.md)| A list of publisher identifiers for which to request tags. | 
 **optional** | ***AttributionTagsApiGetPublisherAttributionTagTemplateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AttributionTagsApiGetPublisherAttributionTagTemplateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **advertiserIds** | [**optional.Interface of []int64**](int64.md)| List of advertiser identifiers for which to request tags. If no values are passed, all advertisers are returned. | 

### Return type

[**AttributionTagResponse**](AttributionTagResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPublisherMacroAttributionTag**
> AttributionTagResponse GetPublisherMacroAttributionTag(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, publisherIds, optional)
Gets a list of attribution tags for third-party publisher campaigns that do not support macros.

Some third-party publishers do not support tags that include macro parameters. In this case, the attribution tag includes a set of '**insertValue**' placeholder values. Replace these placeholder values with your campaign, ad group, and ad identifiers to create unique ad-level tags.<br/><br/> For example: \"?maas=maas_adg_api_123456789_static_9_99&ref_=aa_maas&tag=maas&aa_campaignid={**insertCampaignId**}&aa_adgroupid={**insertAdGroupId**}&aa_creativeid={**insertAdiD**}\"<br/><br/> An example of an integrator nonMacro tag with filled campaign, ad group, and ad ID values is \"?maas=maas_adg_api_123456789_static_9_99&ref_=aa_maas&tag=maas&aa_campaignid=**12345**&aa_adgroupid=**5678**&aa_creativeid=**1357**\"

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **publisherIds** | [**[]string**](string.md)| A list of publisher identifiers for which to request tags. | 
 **optional** | ***AttributionTagsApiGetPublisherMacroAttributionTagOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AttributionTagsApiGetPublisherMacroAttributionTagOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **advertiserIds** | [**optional.Interface of []string**](string.md)| List of advertiser identifiers for which to request tags. If no values are passed, all advertisers are returned. | 

### Return type

[**AttributionTagResponse**](AttributionTagResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

