# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentProfiles**](CreatePaymentProfilesApi.md#CreatePaymentProfiles) | **Post** /billing/paymentProfiles | 

# **CreatePaymentProfiles**
> AdPaymentsCreatePaymentProfilesOutput CreatePaymentProfiles(ctx, body, amazonAdsAccountId, amazonAdvertisingAPIClientId, optional)


Creates or updates payment profiles.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_edit\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]  **Authorized resource type**: Global Manager Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_edit\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AdPaymentsCreatePaymentProfileInput**](AdPaymentsCreatePaymentProfileInput.md)| This API creates or updates payment profiles. | 
  **amazonAdsAccountId** | **string**| The identifier of a profile associated with the advertiser account. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***CreatePaymentProfilesApiCreatePaymentProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a CreatePaymentProfilesApiCreatePaymentProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIManagerAccount** | **optional.**| The header used to pass the ID of a manager account.  Use &#x60;GET&#x60; method on the Manager Accounts resource to list the manager accounts associated with the access token passed in the HTTP Authorization header and choose &#x60;managerAccountId&#x60; from the response to pass it as input. | 

### Return type

[**AdPaymentsCreatePaymentProfilesOutput**](AdPaymentsCreatePaymentProfilesOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.paymentprofiles.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

