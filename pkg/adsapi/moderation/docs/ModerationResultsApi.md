# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ModerationResults**](ModerationResultsApi.md#ModerationResults) | **Post** /moderation/results | 

# **ModerationResults**
> ModerationResultsResponse ModerationResults(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)


API to get the moderation results for the ad. Currently this API supports only SponsoredBrands, SponsoredProducts and SponsoredDisplay ad type.  **Requires one of these permissions**: [\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModerationResultsRequest**](ModerationResultsRequest.md)| The ad identifier along with adProgram for which the moderation results are required. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**ModerationResultsResponse**](ModerationResultsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.moderationresultsrequest.v4.1+json
 - **Accept**: application/vnd.moderationresultsresponse.v4.0+json, application/vnd.moderationresultsbadrequesterror.v4.0+json, application/vnd.moderationresultsaccessdeniederror.v4.0+json, application/vnd.moderationresultsnotfounderror.v4.0+json, application/vnd.moderationresultsthrottlingerror.v4.0+json, application/vnd.moderationresultsinternalservererror.v4.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

