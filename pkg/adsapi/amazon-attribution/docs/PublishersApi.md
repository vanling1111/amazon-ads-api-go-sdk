# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPublishers**](PublishersApi.md#GetPublishers) | **Get** /attribution/publishers | Gets a list of all available publishers.

# **GetPublishers**
> PublishersResponse GetPublishers(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Gets a list of all available publishers.

Use the response to determine whether to use either the macroTags or nonMacroTemplateTags resource to get tags for a certain publisher.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**PublishersResponse**](PublishersResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

