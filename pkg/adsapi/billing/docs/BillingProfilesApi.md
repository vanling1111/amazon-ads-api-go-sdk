# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyBillingProfile**](BillingProfilesApi.md#ApplyBillingProfile) | **Post** /billingProfileUsages | API to link one or more countries with a billing profile. This association is known as &#x27;applying&#x27; a billing profile.
[**CreateBillingProfiles**](BillingProfilesApi.md#CreateBillingProfiles) | **Post** /billingProfiles | API to create one or more billing profile(s).
[**GetBillingProfileAgreementContent**](BillingProfilesApi.md#GetBillingProfileAgreementContent) | **Get** /billingProfileAgreementContents/{billingProfileAgreementContentId} | API to fetch agreement contents related to billing profiles.
[**GetBillingProfileUsages**](BillingProfilesApi.md#GetBillingProfileUsages) | **Post** /billingProfileUsages/list | Lists the billing profiles linked to each country of global ads account.
[**GetBillingProfiles**](BillingProfilesApi.md#GetBillingProfiles) | **Post** /billingProfiles/list | Fetches billing profiles present under the global account. 
[**UpdateBillingProfiles**](BillingProfilesApi.md#UpdateBillingProfiles) | **Put** /billingProfiles | API to update one or more billing profile(s).

# **ApplyBillingProfile**
> ApplyBillingProfileResponse ApplyBillingProfile(ctx, body, amazonAdvertisingAPIClientId, optional)
API to link one or more countries with a billing profile. This association is known as 'applying' a billing profile.

API to link one or more countries with a billing profile. The linked BillingProfile's billing informationwill be used for invoicing, taxes and other billing workflows. A single billing profile can be linked to multiple countries.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_view\",\"adv_billing_edit\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]  **Authorized resource type**: Global Manager Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_view\",\"adv_billing_edit\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ApplyBillingProfileRequest**](ApplyBillingProfileRequest.md)| The request body accepts a list of billing profile identifiers and advertisers for association. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***BillingProfilesApiApplyBillingProfileOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingProfilesApiApplyBillingProfileOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| The identifier of an account. Account can be a global advertising account. | 
 **amazonAdvertisingAPIManagerAccount** | **optional.**| The identifier of a manager account. | 

### Return type

[**ApplyBillingProfileResponse**](ApplyBillingProfileResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.billingProfileUsage.v1+json
 - **Accept**: application/vnd.billingProfileUsage.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBillingProfiles**
> CreateOrUpdateBillingProfilesResponse CreateBillingProfiles(ctx, body, amazonAdvertisingAPIClientId, optional)
API to create one or more billing profile(s).

Creates one or more billing profiles. A Billing Profile contains billing information (address, taxes, agreements etc..) that will be used for invoicing and other billing workflows. Currently, Billing Profiles can only be created for a Global Sponsored Ads vendor account. Please note that you need to link (`POST /billingProfileUsages`) a billingProfile to a country for that billingInformation to be used for that country.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_view\",\"adv_billing_edit\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]  **Authorized resource type**: Global Manager Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_view\",\"adv_billing_edit\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateBillingProfilesRequest**](CreateBillingProfilesRequest.md)| The request body accepts a list of billing profiles to be created in batch. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***BillingProfilesApiCreateBillingProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingProfilesApiCreateBillingProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| The identifier of an account. Account can be a global advertising account. | 
 **amazonAdvertisingAPIManagerAccount** | **optional.**| The identifier of a manager account. | 

### Return type

[**CreateOrUpdateBillingProfilesResponse**](CreateOrUpdateBillingProfilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.billingProfile.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBillingProfileAgreementContent**
> BillingProfileAgreementContentResponse GetBillingProfileAgreementContent(ctx, amazonAdvertisingAPIClientId, billingProfileAgreementContentId, optional)
API to fetch agreement contents related to billing profiles.

User needs to provide consent to certain agreements before creating a billing profile. This API provides a way for users to go through the agreement content.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_view\",\"adv_billing_edit\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]  **Authorized resource type**: Global Manager Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_view\",\"adv_billing_edit\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **billingProfileAgreementContentId** | **string**| agreementId for which the content needs to be fetched. This is same as the &#x60;documentName&#x60; in create/update billing profiles request | 
 **optional** | ***BillingProfilesApiGetBillingProfileAgreementContentOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingProfilesApiGetBillingProfileAgreementContentOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.String**| The identifier of an account. Account can be a global advertising account. | 
 **amazonAdvertisingAPIManagerAccount** | **optional.String**| The identifier of a manager account. | 
 **languageOfPreference** | **optional.String**| This selection controls the language of preference (or &#x60;locale&#x60;) and the returned values of agreement content to match that language. Preferred locale can be chosen among the list of valid language codes. Check the table below for supported language code. &lt;br/&gt;&lt;br/&gt;&lt;table border&#x3D;1&gt;&lt;caption&gt; **Supported Locales Table** &lt;/caption&gt;&lt;tr&gt;    &lt;th&gt;Locale code&lt;/th&gt;    &lt;th&gt;Language&lt;/th&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;en_US&lt;/td&gt;    &lt;td&gt;English&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;ja_JP&lt;/td&gt;    &lt;td&gt;Japanese&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;ar_SA&lt;/td&gt;    &lt;td&gt;Arabic&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;de_DE&lt;/td&gt;    &lt;td&gt;German&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;fr_FR&lt;/td&gt;    &lt;td&gt;French&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;it_IT&lt;/td&gt;    &lt;td&gt;Italian&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;es_MX&lt;/td&gt;    &lt;td&gt;Spanish&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;nl_NL&lt;/td&gt;    &lt;td&gt;Dutch&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;sv_SE&lt;/td&gt;    &lt;td&gt;Swedish&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;pl_PL&lt;/td&gt;    &lt;td&gt;Polish&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;    &lt;td&gt;tr_TR&lt;/td&gt;    &lt;td&gt;Turkish&lt;/td&gt;&lt;/tr&gt; &lt;tr&gt;    &lt;td&gt;pt_BR&lt;/td&gt;    &lt;td&gt;Portuguese&lt;/td&gt;&lt;/tr&gt; &lt;/table&gt; | [default to en_US]

### Return type

[**BillingProfileAgreementContentResponse**](BillingProfileAgreementContentResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.billingProfileAgreement.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBillingProfileUsages**
> GetBillingProfileUsageResponse GetBillingProfileUsages(ctx, body, amazonAdvertisingAPIClientId, optional)
Lists the billing profiles linked to each country of global ads account.

Lists the billing profiles linked to each country of global ads account. You can further narrow down the search by providing the countries you want the billing profiles for.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_view\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]  **Authorized resource type**: Global Manager Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_view\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetBillingProfileUsageRequest**](GetBillingProfileUsageRequest.md)| The request body to fetch billing profiles linked to each country of global ads account. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***BillingProfilesApiGetBillingProfileUsagesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingProfilesApiGetBillingProfileUsagesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| The identifier of an account. Account can be a global advertising account. | 
 **amazonAdvertisingAPIManagerAccount** | **optional.**| The identifier of a manager account. | 

### Return type

[**GetBillingProfileUsageResponse**](GetBillingProfileUsageResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.billingProfileUsage.v1+json
 - **Accept**: application/vnd.billingProfileUsage.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBillingProfiles**
> GetBillingProfilesResponse GetBillingProfiles(ctx, body, amazonAdvertisingAPIClientId, optional)
Fetches billing profiles present under the global account. 

Fetches billing profiles present under the global account. You can limit the search results by providing filters like `defaultBillingProfileFilter` and `billingProfileIdFilter` in which case only results matching the filters will be returned. This is a Paginated API request and returns a paginationToken (`nextToken`) if more that `50` results match. You can customize the pageSize to be less than `50` by providing `maxResults` key in request.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_view\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]  **Authorized resource type**: Global Manager Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_view\",\"MasterAccount_Viewer\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**GetBillingProfilesRequest**](GetBillingProfilesRequest.md)| The request body to fetch one or more billing profile(s). | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***BillingProfilesApiGetBillingProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingProfilesApiGetBillingProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| The identifier of an account. Account can be a global advertising account. | 
 **amazonAdvertisingAPIManagerAccount** | **optional.**| The identifier of a manager account. | 

### Return type

[**GetBillingProfilesResponse**](GetBillingProfilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.billingProfile.v1+json
 - **Accept**: application/vnd.billingProfile.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBillingProfiles**
> CreateOrUpdateBillingProfilesResponse UpdateBillingProfiles(ctx, body, amazonAdvertisingAPIClientId, optional)
API to update one or more billing profile(s).

Updates one or more billing profiles under a global account Please note that `isBillTo` and `type` are immutable attributes and cannot be updated -- in this case, user can always create a new billing profile with same details but with different `isBillTo` and `type` attribute values.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_view\",\"adv_billing_edit\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]  **Authorized resource type**: Global Manager Account ID, Profile ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"adv_billing_view\",\"adv_billing_edit\",\"MasterAccount_Manager\",\"ManagerAccount_Dev\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateBillingProfilesRequest**](UpdateBillingProfilesRequest.md)| The request body accepts a list of billing profiles to be updated in batch. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***BillingProfilesApiUpdateBillingProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BillingProfilesApiUpdateBillingProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **amazonAdsAccountId** | **optional.**| The identifier of an account. Account can be a global advertising account. | 
 **amazonAdvertisingAPIManagerAccount** | **optional.**| The identifier of a manager account. | 

### Return type

[**CreateOrUpdateBillingProfilesResponse**](CreateOrUpdateBillingProfilesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.billingProfile.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

