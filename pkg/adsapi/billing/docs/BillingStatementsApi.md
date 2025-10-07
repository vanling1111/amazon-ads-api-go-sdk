# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingStatement**](BillingStatementsApi.md#CreateBillingStatement) | **Post** /billingStatements | API to request billing statement generation for an Advertising account.
[**GetBillingStatement**](BillingStatementsApi.md#GetBillingStatement) | **Get** /billingStatements/{billingStatementRequestId} | API to fetch the latest status of Billing Statements creation request and billing statement download link if available.

# **CreateBillingStatement**
> CreateBillingStatementResponse CreateBillingStatement(ctx, body, amazonAdvertisingAPIClientId, optional)
API to request billing statement generation for an Advertising account.

Request to create billing statement for advertiser advertising in Sponsored Products/Brands/Display segment.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"nemo_transactions_view\",\"nemo_transactions_edit\",\"ManagerAccount_Dev\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\"]  **Authorized resource type**: Global Manager Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"nemo_transactions_view\",\"nemo_transactions_edit\",\"ManagerAccount_Dev\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\"]  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateBillingStatementRequest**](CreateBillingStatementRequest.md)| The request body accepts filters required for generation of the billing statement report. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***BillingStatementsApiCreateBillingStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingStatementsApiCreateBillingStatementOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| The identifier of a profile associated with the advertiser account. Used for authentication of Global Account. | 
 **amazonAdvertisingAPIManagerAccount** | **optional.**| The identifier of a manager account. | 

### Return type

[**CreateBillingStatementResponse**](CreateBillingStatementResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.createbillingstatementsrequest.v1+json
 - **Accept**: application/vnd.createbillingstatementsrequest.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBillingStatement**
> GetBillingStatementResponse GetBillingStatement(ctx, amazonAdvertisingAPIClientId, billingStatementRequestId, optional)
API to fetch the latest status of Billing Statements creation request and billing statement download link if available.

API to fetch the latest status of Billing Statements creation request and billing statement download link if available.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"nemo_transactions_view\",\"nemo_transactions_edit\",\"ManagerAccount_Dev\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\"]  **Authorized resource type**: Global Manager Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"nemo_transactions_view\",\"nemo_transactions_edit\",\"ManagerAccount_Dev\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\"]  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **billingStatementRequestId** | **string**| Billing Statement Request Id. | 
 **optional** | ***BillingStatementsApiGetBillingStatementOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingStatementsApiGetBillingStatementOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.String**| The identifier of a profile associated with the advertiser account. | 
 **amazonAdvertisingAPIManagerAccount** | **optional.String**| The identifier of a manager account. | 

### Return type

[**GetBillingStatementResponse**](GetBillingStatementResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.createbillingstatementsrequest.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

