# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDocument**](BillingDocumentsApi.md#GetDocument) | **Get** /billing/documents/{documentId} | 

# **GetDocument**
> GetDocumentResponseContent GetDocument(ctx, documentId, docType, amazonAdvertisingAPIClientId, optional)


Gets billing document(s) with id.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"nemo_transactions_view\",\"nemo_transactions_edit\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]  **Authorized resource type**: Global Manager Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"nemo_transactions_view\",\"nemo_transactions_edit\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **documentId** | **string**|  | 
  **docType** | [**[]DocType**](DocType.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***BillingDocumentsApiGetDocumentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingDocumentsApiGetDocumentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdsAccountId** | **optional.String**| Account identifier | 
 **amazonAdvertisingAPIManagerAccount** | **optional.String**| The header used to pass the ID of a manager account. Use &#x60;GET&#x60; method on the Manager Accounts resource to list the manager accounts associated with the access token passed in the HTTP Authorization header and choose &#x60;managerAccountId&#x60; from the response to pass it as input. | 

### Return type

[**GetDocumentResponseContent**](GetDocumentResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.billingDocuments.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

