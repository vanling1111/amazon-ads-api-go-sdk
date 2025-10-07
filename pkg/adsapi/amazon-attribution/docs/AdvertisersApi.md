# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdvertisersByProfile**](AdvertisersApi.md#GetAdvertisersByProfile) | **Get** /attribution/advertisers | Gets a list of advertisers associated with an Amazon Attribution account.

# **GetAdvertisersByProfile**
> AdvertiserResponse GetAdvertisersByProfile(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Gets a list of advertisers associated with an Amazon Attribution account.

For sellers, an attribution profile has one associated advertiser. For vendors, an attribution profile may have more than one associated advertiser.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**AdvertiserResponse**](AdvertiserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

