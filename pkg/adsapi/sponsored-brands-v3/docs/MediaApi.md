# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CompleteUpload**](MediaApi.md#CompleteUpload) | **Put** /media/complete | The API is used to notify that the upload is completed.
[**DescribeMedia**](MediaApi.md#DescribeMedia) | **Get** /media/describe | API to poll for media status

# **CompleteUpload**
> InlineResponse2003 CompleteUpload(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
The API is used to notify that the upload is completed.

The API should be called once the media is uploaded to the location provided by the /media/upload API endpoint. The API creates a Media resource for the uploaded media. Media resource is comprised of Media Identifier. The Media Identifier can be used to attach media to Ad Program (Sponsored Brands).  The API internally kicks off the asynchronous validation and processing workflow of the uploaded media. As a result, Media may not be immediately available for usage (to create Sponsored Brands Video Campaign) as soon as the response is received. See /media/describe API doc for instructions on when media is ready for campaign creation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MediaCompleteBody**](MediaCompleteBody.md)| The  upload location | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DescribeMedia**
> InlineResponse2004 DescribeMedia(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, mediaId)
API to poll for media status

API to poll for media status. In order to attach media to campaign, media should be in either `PendingDeepValidation` or `Available` status.  `Available` status guarantees that media has completed processing and published for usage.  Though media can be attached to campaign once the status of the media transitions to `PendingDeepValidation`, media could still fail additional validation and transition to `Failed` status. For example in the context of SBV, SBV campaign can be created when status transitions to `PendingDeepValidation`, it could result in SBV campaign to be rejected later if media transitions to `Failed` status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **mediaId** | **string**| Media Identifier | 

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

