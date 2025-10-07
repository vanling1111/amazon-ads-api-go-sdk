# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DspCreateAudiencesPost**](AudiencesApi.md#DspCreateAudiencesPost) | **Post** /dsp/audiences | Creates an audience.

# **DspCreateAudiencesPost**
> DspAudienceResponse DspCreateAudiencesPost(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, advertiserId, optional)
Creates an audience.

Creates a targeting audience based on an audience definition.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\",\"audiences_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
  **advertiserId** | **string**| The advertiser to create audience for. | 
 **optional** | ***AudiencesApiDspCreateAudiencesPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AudiencesApiDspCreateAudiencesPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of []DspAudienceCreateRequestItem**](DspAudienceCreateRequestItem.md)| An array of audience objects. For each object, specify required fields and their values. Maximum length of the array is 1. | 

### Return type

[**DspAudienceResponse**](DspAudienceResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dspaudiences.v1+json
 - **Accept**: application/vnd.dspaudiencesresponse.v1+json, application/vnd.dspaudienceserror.v1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

