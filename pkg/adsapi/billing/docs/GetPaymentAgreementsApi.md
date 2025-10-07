# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPaymentAgreements**](GetPaymentAgreementsApi.md#GetPaymentAgreements) | **Post** /billing/paymentAgreements/list | 

# **GetPaymentAgreements**
> AdPaymentsGetPaymentAgreementOutput GetPaymentAgreements(ctx, amazonAdvertisingAPIClientId, agreementType, optional)


Gets current payment agreement for a customer.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_view\",\"adv_billing_edit\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]  **Authorized resource type**: Global Manager Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_view\",\"adv_billing_edit\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **agreementType** | [**AdPaymentsPaymentAgreementType**](.md)| The type of payment agreement to query for. | 
 **optional** | ***GetPaymentAgreementsApiGetPaymentAgreementsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a GetPaymentAgreementsApiGetPaymentAgreementsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.String**| The header used to pass global account associated with the advertiser account. Use &#x60;GET&#x60; method on the Global Ads Account resource to list the global ads account associated with the access token passed in the HTTP Authorization header and choose &#x60;AdvertisingAccountIdentifier&#x60; id from the response to pass it as input. | 
 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdvertisingAPIManagerAccount** | **optional.String**| The header used to pass the ID of a manager account.  Use &#x60;GET&#x60; method on the Manager Accounts resource to list the manager accounts associated with the access token passed in the HTTP Authorization header and choose &#x60;managerAccountId&#x60; from the response to pass it as input. | 
 **nextToken** | **optional.String**| To retrieve the next page of results, call the same operation and specify this token in the request. If the nextToken field is empty, there are no further results. | 

### Return type

[**AdPaymentsGetPaymentAgreementOutput**](AdPaymentsGetPaymentAgreementOutput.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.paymentagreements.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

