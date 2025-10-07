# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InsightsGetAudiencesOverlappingAudiences**](AudienceInsightsApi.md#InsightsGetAudiencesOverlappingAudiences) | **Get** /insights/audiences/{audienceId}/overlappingAudiences | Retrieves the top audiences that overlap with the provided audience.

# **InsightsGetAudiencesOverlappingAudiences**
> InlineResponse200 InsightsGetAudiencesOverlappingAudiences(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, audienceId, adType, optional)
Retrieves the top audiences that overlap with the provided audience.

  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **audienceId** | **string**| The identifier of an audience. | 
  **adType** | **string**| The advertising program. | 
 **optional** | ***AudienceInsightsApiInsightsGetAudiencesOverlappingAudiencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AudienceInsightsApiInsightsGetAudiencesOverlappingAudiencesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **advertiserId** | **optional.String**| The identifier of the advertiser you&#x27;d like to retrieve overlapping audiences for. This parameter is required for the DSP adType, but is optional for the SD adType. | 
 **minimumOverlapAffinity** | **optional.Float64**| If specified, the affinities of all returned overlapping audiences will be at least the provided affinity. | 
 **maximumOverlapAffinity** | **optional.Float64**| If specified, the affinities of all returned overlapping audiences will be at most the provided affinity. | 
 **audienceCategory** | [**optional.Interface of []string**](string.md)| If specified, the categories of all returned overlapping audiences will be one of the provided categories. | 
 **maxResults** | **optional.Int32**| Sets the maximum number of overlapping audiences in the response. This parameter is supported only for request to return &#x60;application/vnd.insightsaudiencesoverlap.v2+json&#x60;. | 
 **nextToken** | **optional.String**| Token to be used to request additional overlapping audiences. If not provided, the top 30 overlapping audiences are returned. Note: subsequent calls must be made using the same parameters as used in previous requests. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.insightsaudiencesoverlap.v2+json, application/vnd.insightserror.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

