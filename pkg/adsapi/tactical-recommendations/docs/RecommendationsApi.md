# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyRecommendations**](RecommendationsApi.md#ApplyRecommendations) | **Post** /recommendations/apply | Applies one or more recommendations.
[**ListRecommendations**](RecommendationsApi.md#ListRecommendations) | **Post** /recommendations/list | Retrieves a list of recommendations.
[**UpdateRecommendation**](RecommendationsApi.md#UpdateRecommendation) | **Put** /recommendations/{recommendationId} | Updates a recommendation.

# **ApplyRecommendations**
> ApplyRecommendationsResponse ApplyRecommendations(ctx, body, amazonAdvertisingAPIClientId, optional)
Applies one or more recommendations.

Applies one or more recommendations.  When applying recommendations with recommendationType NEW_CAMPAIGN or NEW_VIDEO_CAMPAIGN, this will return HTTP status code ACCEPTED to indicate that it has been accepted for processing. You can then use POST /recommendations/list to check for the status of these recommendations.  A recommendation's status will be APPLY_SUCCESS if it has been successfully applied. Otherwise, the status will be APPLY_FAILED.   **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplyRecommendationsRequest**](ApplyRecommendationsRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***RecommendationsApiApplyRecommendationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecommendationsApiApplyRecommendationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**ApplyRecommendationsResponse**](ApplyRecommendationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.applyRecommendationsRequest.v1+json
 - **Accept**: application/vnd.applyRecommendationsResponse.v1+json, application/vnd.recommendationApiError.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListRecommendations**
> ListRecommendationsResponse ListRecommendations(ctx, body, amazonAdvertisingAPIClientId, optional)
Retrieves a list of recommendations.

Retrieves a paginated list of recommendations with optional filtering.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ListRecommendationsRequest**](ListRecommendationsRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***RecommendationsApiListRecommendationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecommendationsApiListRecommendationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. account. | 

### Return type

[**ListRecommendationsResponse**](ListRecommendationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.listRecommendationsRequest.v1+json
 - **Accept**: application/vnd.listRecommendationsResponse.v1+json, application/vnd.recommendationApiError.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateRecommendation**
> Recommendation UpdateRecommendation(ctx, body, amazonAdvertisingAPIClientId, recommendationId, optional)
Updates a recommendation.

Updates a recommendation.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateRecommendationRequest**](UpdateRecommendationRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **recommendationId** | **string**| The identifier of the recommendation. | 
 **optional** | ***RecommendationsApiUpdateRecommendationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RecommendationsApiUpdateRecommendationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. account. | 

### Return type

[**Recommendation**](Recommendation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.updateRecommendationRequest.v1+json
 - **Accept**: application/vnd.recommendation.v1+json, application/vnd.recommendationApiError.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

