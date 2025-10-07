# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DspAdvertisersAdvertiserIdGet**](AdvertiserApi.md#DspAdvertisersAdvertiserIdGet) | **Get** /dsp/advertisers/{advertiserId} | 
[**DspAdvertisersGet**](AdvertiserApi.md#DspAdvertisersGet) | **Get** /dsp/advertisers | 

# **DspAdvertisersAdvertiserIdGet**
> DspAdvertiserV1 DspAdvertisersAdvertiserIdGet(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, advertiserId)


Returns advertiser information based on given advertiser id.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The client identifier of the customer making the request. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. For DSP profiles, the &#x60;type&#x60; field of the &#x60;accountInfo&#x60; object must be set to &#x60;agency&#x60; and the &#x60;subType&#x60; must not be &#x60;AMAZON_ATTRIBUTION&#x60;. | 
  **advertiserId** | **string**| Unique id to identify advertiser | 

### Return type

[**DspAdvertiserV1**](DspAdvertiserV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DspAdvertisersGet**
> DspAdvertisersV1 DspAdvertisersGet(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Returns a list of advertisers with information which satisfy the filtering criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The client identifier of the customer making the request. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use GET method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. For DSP profiles, the &#x60;type&#x60; field of the &#x60;accountInfo&#x60; object must be set to &#x60;agency&#x60; and the &#x60;subType&#x60; must not be &#x60;AMAZON_ATTRIBUTION&#x60;. | 
 **optional** | ***AdvertiserApiDspAdvertisersGetOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AdvertiserApiDspAdvertisersGetOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Sets a cursor into the requested set of advertisers. Use in conjunction with the count parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0. | [default to 0]
 **count** | **optional.Int32**| Sets the number of advertisers to be returned in a single call. Maximum of 100 advertisers per call. | [default to 100]
 **advertiserIdFilter** | **optional.String**| List of comma separated advertiser ids to filter the advertisers. If no advertiser ids provided, all advertisers in this entity will be returned. | 

### Return type

[**DspAdvertisersV1**](DspAdvertisersV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

