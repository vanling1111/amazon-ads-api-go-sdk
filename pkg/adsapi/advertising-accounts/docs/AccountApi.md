# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccount**](AccountApi.md#GetAccount) | **Get** /adsAccounts/{advertisingAccountId} | Request attributes of a given advertising account.
[**ListAdsAccounts**](AccountApi.md#ListAdsAccounts) | **Post** /adsAccounts/list | List all advertising accounts for the user associated with the access token.
[**RegisterAdsAccount**](AccountApi.md#RegisterAdsAccount) | **Post** /adsAccounts | Create a new advertising account tied to a specific Amazon vendor, seller or author, or to a business who does not sell on Amazon.

# **GetAccount**
> GetAccountResponseContent GetAccount(ctx, advertisingAccountId, optional)
Request attributes of a given advertising account.

Request attributes of a given advertising account.  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **advertisingAccountId** | **string**| This is the global advertising account Id from the client. | 
 **optional** | ***AccountApiGetAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiGetAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amazonAdvertisingAPIClientId** | **optional.String**| The identifier of a client associated with a Login with Amazon account. | 

### Return type

[**GetAccountResponseContent**](GetAccountResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.accountresource.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAdsAccounts**
> ListAdsAccountsResponseContent ListAdsAccounts(ctx, optional)
List all advertising accounts for the user associated with the access token.

List all advertising accounts for the user associated with the access token.  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiListAdsAccountsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiListAdsAccountsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of ListAdsAccountsRequestContent**](ListAdsAccountsRequestContent.md)|  | 
 **amazonAdvertisingAPIClientId** | **optional.**| The identifier of a client associated with a Login with Amazon account. | 

### Return type

[**ListAdsAccountsResponseContent**](ListAdsAccountsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.listaccountsresource.v1+json
 - **Accept**: application/vnd.listaccountsresource.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RegisterAdsAccount**
> RegisterAdsAccountResponseContent RegisterAdsAccount(ctx, optional)
Create a new advertising account tied to a specific Amazon vendor, seller or author, or to a business who does not sell on Amazon.

Create a new advertising account tied to a specific Amazon vendor, seller or author, or to a business who does not sell on Amazon.  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AccountApiRegisterAdsAccountOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AccountApiRegisterAdsAccountOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RegisterAdsAccountRequestContent**](RegisterAdsAccountRequestContent.md)|  | 
 **amazonAdvertisingAPIClientId** | **optional.**| The identifier of a client associated with a Login with Amazon account. | 

### Return type

[**RegisterAdsAccountResponseContent**](RegisterAdsAccountResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.registeradsaccountresource.v1+json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

