# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ProgramEligibility**](ProgramEligibilityApi.md#ProgramEligibility) | **Post** /eligibility/programs | 

# **ProgramEligibility**
> ProgramEligibilityResponseContent ProgramEligibility(ctx, amazonAdvertisingAPIClientId, optional)


Checks the advertiser's eligibility to ad programs.  **Authorized resource type**: Global Ad Account ID, Profile ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***ProgramEligibilityApiProgramEligibilityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProgramEligibilityApiProgramEligibilityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ProgramEligibilityRequestContent**](ProgramEligibilityRequestContent.md)|  | 
 **acceptLanguage** | [**optional.Interface of AcceptLanguage**](.md)| Specify the language in which the response is returned. | 
 **amazonAdsAccountId** | **optional.**| The header used to pass global account associated with the advertiser account Use &#x60;GET&#x60; method on the Global Ads Account resource to list the global ads account associated with the access token passed in the HTTP Authorization header and choose AdvertisingAccountIdentifier id from the response to pass it as input. Use for v2 global calls | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **contentType** | **optional.**| The content type of the request. | 

### Return type

[**ProgramEligibilityResponseContent**](ProgramEligibilityResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/vnd.programeligibility.v2+json
 - **Accept**: application/json, application/vnd.programeligibility.v2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

