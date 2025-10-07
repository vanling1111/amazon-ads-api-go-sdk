# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PartnerOpportunitiesApplicationStatus**](PartnerOpportunitiesApi.md#PartnerOpportunitiesApplicationStatus) | **Post** /partnerOpportunities/{partnerOpportunityId}/applicationStatus | 
[**PartnerOpportunitiesApply**](PartnerOpportunitiesApi.md#PartnerOpportunitiesApply) | **Post** /partnerOpportunities/{partnerOpportunityId}/apply | 
[**PartnerOpportunitiesGetOpportunityFile**](PartnerOpportunitiesApi.md#PartnerOpportunitiesGetOpportunityFile) | **Get** /partnerOpportunities/{partnerOpportunityId}/file | 
[**PartnerOpportunitiesListOpportunities**](PartnerOpportunitiesApi.md#PartnerOpportunitiesListOpportunities) | **Get** /partnerOpportunities | 
[**PartnerOpportunitiesSummarizeOpportunities**](PartnerOpportunitiesApi.md#PartnerOpportunitiesSummarizeOpportunities) | **Get** /partnerOpportunities/summary | 

# **PartnerOpportunitiesApplicationStatus**
> PartnerOpportunitiesApplicationStatusResponseDtoV1 PartnerOpportunitiesApplicationStatus(ctx, body, partnerOpportunityId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIManagerAccount)


Retrieves the current status of applied recommendations.  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"MasterAccount_Manager\",\"MasterAccount_Viewer\",\"ManagerAccount_Dev\",\"MA_ReadOnly\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PartnerOpportunitiesApplicationStatusRequestDtoV1**](PartnerOpportunitiesApplicationStatusRequestDtoV1.md)|  | 
  **partnerOpportunityId** | **string**|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
  **amazonAdvertisingAPIManagerAccount** | **string**| &#x27;Partner Network Account ID&#x27; which is accessible from Partner Network under the [&#x27;User settings&#x27;](https://advertising.amazon.com/partner-network/settings) link in the upper right corner. | 

### Return type

[**PartnerOpportunitiesApplicationStatusResponseDtoV1**](PartnerOpportunitiesApplicationStatusResponseDtoV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.partneropportunity.v1+json, application/vnd.partneropportunity.v1.1+json, application/vnd.partneropportunity.v1.2+json
 - **Accept**: application/vnd.partneropportunity.v1+json, application/vnd.partneropportunity.v1.1+json, application/vnd.partneropportunity.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerOpportunitiesApply**
> PartnerOpportunitiesApply(ctx, body, partnerOpportunityId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIManagerAccount)


Applies a given set of recommendations. Application may be asynchronous. Application status may be checked using the applicationStatus operation. Note that all required parameters are retrieved from opportunity data.  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"MasterAccount_Manager\",\"MasterAccount_Viewer\",\"ManagerAccount_Dev\",\"MA_ReadOnly\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PartnerOpportunitiesApplyRequestDtoV1**](PartnerOpportunitiesApplyRequestDtoV1.md)|  | 
  **partnerOpportunityId** | **string**|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
  **amazonAdvertisingAPIManagerAccount** | **string**| &#x27;Partner Network Account ID&#x27; which is accessible from Partner Network under the [&#x27;User settings&#x27;](https://advertising.amazon.com/partner-network/settings) link in the upper right corner. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.partneropportunity.v1+json, application/vnd.partneropportunity.v1.1+json, application/vnd.partneropportunity.v1.2+json
 - **Accept**: application/vnd.partneropportunity.v1+json, application/vnd.partneropportunity.v1.1+json, application/vnd.partneropportunity.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerOpportunitiesGetOpportunityFile**
> PartnerOpportunitiesGetOpportunityFile(ctx, partnerOpportunityId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIManagerAccount, optional)


Gets a 307 - TEMPORARY_REDIRECT to an opportunity data file.  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"MasterAccount_Manager\",\"MasterAccount_Viewer\",\"ManagerAccount_Dev\",\"MA_ReadOnly\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **partnerOpportunityId** | **string**| The opportunity ID for which the file URL is desired. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
  **amazonAdvertisingAPIManagerAccount** | **string**| &#x27;Partner Network Account ID&#x27; which is accessible from Partner Network under the [&#x27;User settings&#x27;](https://advertising.amazon.com/partner-network/settings) link in the upper right corner. | 
 **optional** | ***PartnerOpportunitiesApiPartnerOpportunitiesGetOpportunityFileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerOpportunitiesApiPartnerOpportunitiesGetOpportunityFileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fileFormat** | **optional.String**| Specify the desired file format for the opportunity data. Default is CSV. | [default to CSV]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerOpportunitiesListOpportunities**
> PartnerOpportunitiesOpportunitiesPageV1 PartnerOpportunitiesListOpportunities(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIManagerAccount, optional)


Gets a list of opportunities specific to the partner making the request.  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"MasterAccount_Manager\",\"MasterAccount_Viewer\",\"ManagerAccount_Dev\",\"MA_ReadOnly\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
  **amazonAdvertisingAPIManagerAccount** | **string**| &#x27;Partner Network Account ID&#x27; which is accessible from Partner Network under the [&#x27;User settings&#x27;](https://advertising.amazon.com/partner-network/settings) link in the upper right corner. | 
 **optional** | ***PartnerOpportunitiesApiPartnerOpportunitiesListOpportunitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerOpportunitiesApiPartnerOpportunitiesListOpportunitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **maxResults** | **optional.Float64**| The maximum number of results to return in a single page. | 
 **nextToken** | **optional.String**| An obfuscated cursor value that indicates which &#x27;page&#x27; of results should be returned next. | 
 **locale** | [**optional.Interface of Locale**](.md)| The desired locale for opportunity data to be presented in. The &#x60;title&#x60;, &#x60;description&#x60;, and &#x60;callToAction&#x60; fields of the response items support localization. | 
 **retrieveTranslationKeys** | **optional.Bool**|  | [default to false]
 **advertiserId** | [**optional.Interface of []string**](string.md)| Filter for opportunities using advertiserId. Each advertiserId must be fewer than 50 characters. | 
 **profileId** | [**optional.Interface of []float64**](float64.md)| Filter for opportunities using profileId. Each profileId must be fewer than 50 characters. | 
 **audience** | [**optional.Interface of []string**](string.md)| Filter for opportunities with these audience values. * PARTNER_MANAGED_ADVERTISERS - Recommendation relates to advertisers the partner manages. * PARTNER_MANAGED_AD_BUSINESS - Recommendation relates to other partners the partner interacts with. * PARTNER - Recommendation relates to you, the partner. | 
 **objectiveType** | [**optional.Interface of []string**](string.md)| Filter for opportunities with these objectiveType values. * AD_API_ENDPOINT_ADOPTION - Recommendation relates to adopting a new API endpoint. * ADVERTISER_INSIGHTS - Recommendation relates to advertiser insights. * AMAZON_ACCOUNT_TEAM_RECOMMENDATIONS - Recommendation is provided by the Amazon Ads Account Team. * BENCHMARKING_INSIGHTS - Recommendation relates to performance insights comparing against similar partners, brands, campaigns, or ASINs. * CAMPAIGN_OPTIMIZATION - Recommendation relates to optimizing campaigns. * CATEGORY_INSIGHTS - Recommendation relates to advertising insights across product categories.. * CLICK_CREDITS - Recommendation relates to available click credits. * DEALS - Recommendation relates to deals. * MARKETPLACE_EXPANSION - Recommendation relates to expanding to new marketplaces. * NEW_TO_BRAND_INSIGHTS - Recommendation relates to new to brand advertising insights. * PARTNER_GROWTH - Recommendation relates to growing your business as a partner. * PATH_TO_PURCHASE_INSIGHTS - Recommendation relates to path to purchase insights. * READY_TO_LAUNCH_CAMPAIGNS - Recommendation relates to ready to launch campaigns. * RETAIL_INSIGHTS - Recommendation related to retail insights about products you manage. * SHARE_OF_VOICE_INSIGHTS - Recommendation relates to share of voice for a particular audience. * UNLAUNCHED_ASINS - Recommendation relates to ASINs you manage that are not enrolled in advertising campaigns.  | 
 **product** | [**optional.Interface of []string**](string.md)| Filter for opportunities with these product values. * AMAZON_DSP - Recommendation relates to the Amazon DSP. * AMAZON_LIVE - Recommendation relates to Amazon&#x27;s Live Show and Tell program. * CROSS_PRODUCT - Recommendation relates to multiple Amazon Ads products including leveraging insights between products * POSTS - Deprecated * SPONSORED_BRANDS - Recommendation relates to Sponsored Brands. * SPONSORED_BRANDS_VIDEO - Recommendation relates to Sponsored Brands Video. * SPONSORED_DISPLAY - Recommendation relates to Sponsored Display. * SPONSORED_DISPLAY_VIDEO - Recommendation relates to Sponsored Display Video. * SPONSORED_PRODUCTS - Recommendation relates to Sponsored Products. * SPONSORED_TV - Recommendation relates to Sponsored TV. * STORES - Recommendation relates to building a storefront page on Amazon. * VIDEO_ADS - Deprecated value, replaced by SPONSORED_BRANDS_VIDEO and SPONSORED_DISPLAY_VIDEO values. | 

### Return type

[**PartnerOpportunitiesOpportunitiesPageV1**](PartnerOpportunitiesOpportunitiesPageV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.partneropportunity.v1+json, application/vnd.partneropportunity.v1.1+json, application/vnd.partneropportunity.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PartnerOpportunitiesSummarizeOpportunities**
> PartnerOpportunitiesOpportunitiesSummaryV1 PartnerOpportunitiesSummarizeOpportunities(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIManagerAccount, optional)


Gets aggregated information about all opportunities specific to the partner making the request. Supported since V1.1.  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"MasterAccount_Manager\",\"MasterAccount_Viewer\",\"ManagerAccount_Dev\",\"MA_ReadOnly\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x27;Login with Amazon&#x27; account. | 
  **amazonAdvertisingAPIManagerAccount** | **string**| &#x27;Partner Network Account ID&#x27; which is accessible from Partner Network under the [&#x27;User settings&#x27;](https://advertising.amazon.com/partner-network/settings) link in the upper right corner. | 
 **optional** | ***PartnerOpportunitiesApiPartnerOpportunitiesSummarizeOpportunitiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PartnerOpportunitiesApiPartnerOpportunitiesSummarizeOpportunitiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **audience** | [**optional.Interface of []string**](string.md)| Filter for opportunities with these audience values. * PARTNER_MANAGED_ADVERTISERS - Recommendation relates to advertisers the partner manages. * PARTNER_MANAGED_AD_BUSINESS - Recommendation relates to other partners the partner interacts with. * PARTNER - Recommendation relates to you, the partner. | 
 **objectiveType** | [**optional.Interface of []string**](string.md)| Filter for opportunities with these objectiveType values. * AD_API_ENDPOINT_ADOPTION - Recommendation relates to adopting a new API endpoint. * ADVERTISER_INSIGHTS - Recommendation relates to advertiser insights. * AMAZON_ACCOUNT_TEAM_RECOMMENDATIONS - Recommendation is provided by the Amazon Ads Account Team. * BENCHMARKING_INSIGHTS - Recommendation relates to performance insights comparing against similar partners, brands, campaigns, or ASINs. * CAMPAIGN_OPTIMIZATION - Recommendation relates to optimizing campaigns. * CATEGORY_INSIGHTS - Recommendation relates to advertising insights across product categories.. * CLICK_CREDITS - Recommendation relates to available click credits. * DEALS - Recommendation relates to deals. * MARKETPLACE_EXPANSION - Recommendation relates to expanding to new marketplaces. * NEW_TO_BRAND_INSIGHTS - Recommendation relates to new to brand advertising insights. * PARTNER_GROWTH - Recommendation relates to growing your business as a partner. * PATH_TO_PURCHASE_INSIGHTS - Recommendation relates to path to purchase insights. * READY_TO_LAUNCH_CAMPAIGNS - Recommendation relates to ready to launch campaigns. * RETAIL_INSIGHTS - Recommendation related to retail insights about products you manage. * SHARE_OF_VOICE_INSIGHTS - Recommendation relates to share of voice for a particular audience. * UNLAUNCHED_ASINS - Recommendation relates to ASINs you manage that are not enrolled in advertising campaigns.  | 
 **product** | [**optional.Interface of []string**](string.md)| Filter for opportunities with these product values. * AMAZON_DSP - Recommendation relates to the Amazon DSP. * AMAZON_LIVE - Recommendation relates to Amazon&#x27;s Live Show and Tell program. * CROSS_PRODUCT - Recommendation relates to multiple Amazon Ads products including leveraging insights between products * POSTS - Deprecated * SPONSORED_BRANDS - Recommendation relates to Sponsored Brands. * SPONSORED_BRANDS_VIDEO - Recommendation relates to Sponsored Brands Video. * SPONSORED_DISPLAY - Recommendation relates to Sponsored Display. * SPONSORED_DISPLAY_VIDEO - Recommendation relates to Sponsored Display Video. * SPONSORED_PRODUCTS - Recommendation relates to Sponsored Products. * SPONSORED_TV - Recommendation relates to Sponsored TV. * STORES - Recommendation relates to building a storefront page on Amazon. * VIDEO_ADS - Deprecated value, replaced by SPONSORED_BRANDS_VIDEO and SPONSORED_DISPLAY_VIDEO values. | 

### Return type

[**PartnerOpportunitiesOpportunitiesSummaryV1**](PartnerOpportunitiesOpportunitiesSummaryV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.partneropportunity.v1+json, application/vnd.partneropportunity.v1.1+json, application/vnd.partneropportunity.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

