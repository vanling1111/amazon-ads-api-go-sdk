# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DspAudienceDelete**](AdsApi.md#DspAudienceDelete) | **Post** /dsp/audiences/delete | 
[**DspAudienceEdit**](AdsApi.md#DspAudienceEdit) | **Put** /dsp/audiences/edit | 

# **DspAudienceDelete**
> DspAudienceDeleteResponseContent DspAudienceDelete(ctx, body, advertiserId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)


Deletes an existing targeting audience based on audience ID. Only available for the audiences of the type: *PRODUCT_PURCHASES*, *PRODUCT_VIEWS*, *PRODUCT_SIMS*, *PRODUCT_SEARCH* and *COMBINED_AUDIENCE*  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"audiences_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DspAudienceDeleteRequestContent**](DspAudienceDeleteRequestContent.md)|  | 
  **advertiserId** | **string**| ID that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). The owner of the audience. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**DspAudienceDeleteResponseContent**](DspAudienceDeleteResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dspaudiences.v1+json
 - **Accept**: application/vnd.dspaudiences.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DspAudienceEdit**
> DspAudienceEditResponseContent DspAudienceEdit(ctx, body, advertiserId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)


Updates an existing targeting audience based on an audience definition and audience ID.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"audiences_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DspAudienceEditRequestContent**](DspAudienceEditRequestContent.md)|  | 
  **advertiserId** | **string**| ID that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). The owner of the audience. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**DspAudienceEditResponseContent**](DspAudienceEditResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dspaudiences.v1+json
 - **Accept**: application/vnd.dspaudiences.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

