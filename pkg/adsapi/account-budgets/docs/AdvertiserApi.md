# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccountBudgetFeatureFlags**](AdvertiserApi.md#GetAccountBudgetFeatureFlags) | **Get** /accountBudgets/featureFlags | Gets account budget feature flags information.
[**UpdateAccountBudgetFeatureFlags**](AdvertiserApi.md#UpdateAccountBudgetFeatureFlags) | **Post** /accountBudgets/featureFlags | Creates or Updates account budget feature flags information.

# **GetAccountBudgetFeatureFlags**
> GetAccountBudgetFeatureFlagsResponse GetAccountBudgetFeatureFlags(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Gets account budget feature flags information.

Gets account budget feature flags information.  **Requires one of these permissions**: [\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**GetAccountBudgetFeatureFlagsResponse**](GetAccountBudgetFeatureFlagsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.accountBudgetFeatureFlags.v1+json, application/vnd.accountBudgetFeatureFlagsError.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateAccountBudgetFeatureFlags**
> UpdateAccountBudgetFeatureFlagsResponse UpdateAccountBudgetFeatureFlags(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Creates or Updates account budget feature flags information.

Creates or Updates account budget feature flags information.  **Requires one of these permissions**: [\"advertiser_campaign_view\",\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateAccountBudgetFeatureFlagsRequest**](UpdateAccountBudgetFeatureFlagsRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**UpdateAccountBudgetFeatureFlagsResponse**](UpdateAccountBudgetFeatureFlagsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.accountBudgetFeatureFlags.v1+json
 - **Accept**: application/vnd.accountBudgetFeatureFlags.v1+json, application/vnd.accountBudgetFeatureFlagsError.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

