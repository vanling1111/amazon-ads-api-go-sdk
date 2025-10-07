# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkGetBillingNotifications**](BillingNotificationsApi.md#BulkGetBillingNotifications) | **Post** /billing/notifications | Get the billing notifications for a list advertising accounts.

# **BulkGetBillingNotifications**
> BulkGetBillingNotificationsResponse BulkGetBillingNotifications(ctx, body, amazonAdvertisingAPIClientId, optional)
Get the billing notifications for a list advertising accounts.

Gets an array of all currently valid billing notifications associated for each advertising account.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"adv_billing_view\",\"adv_billing_edit\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]  **Authorized resource type**: Global Manager Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"adv_billing_view\",\"adv_billing_edit\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BulkGetBillingNotificationsRequestBody**](BulkGetBillingNotificationsRequestBody.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***BillingNotificationsApiBulkGetBillingNotificationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingNotificationsApiBulkGetBillingNotificationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| The header used to pass global account associated with the advertiser account. Use &#x60;GET&#x60; method on the Global Ads Account resource to list the global ads account associated with the access token passed in the HTTP Authorization header and choose &#x60;AdvertisingAccountIdentifier&#x60; id from the response to pass it as input. | 
 **amazonAdvertisingAPIManagerAccount** | **optional.**| The header used to pass the ID of a manager account.  Use &#x60;GET&#x60; method on the Manager Accounts resource to list the manager accounts associated with the access token passed in the HTTP Authorization header and choose &#x60;managerAccountId&#x60; from the response to pass it as input. | 

### Return type

[**BulkGetBillingNotificationsResponse**](bulkGetBillingNotificationsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.billingnotifications.v1+json
 - **Accept**: application/vnd.bulkgetbillingnotificationsresponse.v1+json, application/vnd.bulkgetbillingnotificationserrorresponse.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

