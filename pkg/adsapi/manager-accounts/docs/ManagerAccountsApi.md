# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManagerAccount**](ManagerAccountsApi.md#CreateManagerAccount) | **Post** /managerAccounts | Creates a new Amazon Advertising Manager account.
[**GetManagerAccountsForUser**](ManagerAccountsApi.md#GetManagerAccountsForUser) | **Get** /managerAccounts | Returns all manager accounts that a given Amazon Ads user has access to.
[**LinkAdvertisingAccountsToManagerAccountPublicAPI**](ManagerAccountsApi.md#LinkAdvertisingAccountsToManagerAccountPublicAPI) | **Post** /managerAccounts/{managerAccountId}/associate | Link Amazon Advertising accounts or advertisers with a Manager Account.
[**UnlinkAdvertisingAccountsToManagerAccountPublicAPI**](ManagerAccountsApi.md#UnlinkAdvertisingAccountsToManagerAccountPublicAPI) | **Post** /managerAccounts/{managerAccountId}/disassociate | Unlink Amazon Advertising accounts or advertisers with a Manager Account.

# **CreateManagerAccount**
> ManagerAccount CreateManagerAccount(ctx, body, amazonAdvertisingAPIClientId)
Creates a new Amazon Advertising Manager account.

Creates a new Amazon Advertising [Manager account](https://advertising.amazon.com/help?ref_=a20m_us_blog_whtsnewfb2020_040120#GU3YDB26FR7XT3C8).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateManagerAccountRequest**](CreateManagerAccountRequest.md)| Request object required to create a new Manager account. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 

### Return type

[**ManagerAccount**](ManagerAccount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.createmanageraccountrequest.v1+json
 - **Accept**: application/vnd.manageraccount.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetManagerAccountsForUser**
> GetManagerAccountsResponse GetManagerAccountsForUser(ctx, amazonAdvertisingAPIClientId)
Returns all manager accounts that a given Amazon Ads user has access to.

Returns all [manager accounts](https://advertising.amazon.com/help?ref_=a20m_us_blog_whtsnewfb2020_040120#GU3YDB26FR7XT3C8) that a user has access to, along with metadata for the Amazon Ads accounts that are linked to each manager account. NOTE: A maximum of 50 linked accounts are returned for each manager account.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 

### Return type

[**GetManagerAccountsResponse**](GetManagerAccountsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.getmanageraccountsresponse.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LinkAdvertisingAccountsToManagerAccountPublicAPI**
> UpdateAdvertisingAccountsInManagerAccountResponse LinkAdvertisingAccountsToManagerAccountPublicAPI(ctx, body, amazonAdvertisingAPIClientId, managerAccountId)
Link Amazon Advertising accounts or advertisers with a Manager Account.

Link Amazon Advertising accounts or advertisers with a [Manager Account](https://advertising.amazon.com/help?ref_=a20m_us_blog_whtsnewfb2020_040120#GU3YDB26FR7XT3C8).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateAdvertisingAccountsInManagerAccountRequest**](UpdateAdvertisingAccountsInManagerAccountRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **managerAccountId** | **string**| Id of the Manager Account. | 

### Return type

[**UpdateAdvertisingAccountsInManagerAccountResponse**](UpdateAdvertisingAccountsInManagerAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.updateadvertisingaccountsinmanageraccountrequest.v1+json
 - **Accept**: application/vnd.updateadvertisingaccountsinmanageraccountresponse.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnlinkAdvertisingAccountsToManagerAccountPublicAPI**
> UpdateAdvertisingAccountsInManagerAccountResponse UnlinkAdvertisingAccountsToManagerAccountPublicAPI(ctx, body, amazonAdvertisingAPIClientId, managerAccountId)
Unlink Amazon Advertising accounts or advertisers with a Manager Account.

Unlink Amazon Advertising accounts or advertisers with a [Manager Account](https://advertising.amazon.com/help?ref_=a20m_us_blog_whtsnewfb2020_040120#GU3YDB26FR7XT3C8).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateAdvertisingAccountsInManagerAccountRequest**](UpdateAdvertisingAccountsInManagerAccountRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **managerAccountId** | **string**| Id of the Manager Account. | 

### Return type

[**UpdateAdvertisingAccountsInManagerAccountResponse**](UpdateAdvertisingAccountsInManagerAccountResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.updateadvertisingaccountsinmanageraccountrequest.v1+json
 - **Accept**: application/vnd.updateadvertisingaccountsinmanageraccountresponse.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

