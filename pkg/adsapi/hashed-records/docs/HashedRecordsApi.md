# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UploadHashedRecords**](HashedRecordsApi.md#UploadHashedRecords) | **Post** /dp/records/hashed | Upload a batch of hashed records for matching

# **UploadHashedRecords**
> InlineResponse200 UploadHashedRecords(ctx, authorization, amazonAdvertisingAPIClientId, contentType, optional)
Upload a batch of hashed records for matching

Saves a batch of hashed personally-identifiable information (PII) records to be matched with Amazon identities for future use.  All inputs must be properly normalized and SHA-256 hashed as defined [in the documentation](https://advertising.amazon.com/help/GCCXMZYCK4RXWS6C).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **authorization** | **string**| A valid access token. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **contentType** | **string**| The Content-Type is application/json. | 
 **optional** | ***HashedRecordsApiUploadHashedRecordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a HashedRecordsApiUploadHashedRecordsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of RecordsHashedBody**](RecordsHashedBody.md)| Hashed PII records to be matched with Amazon identities for future use.  All inputs must be properly normalized and SHA-256 hashed as defined [in the documentation](https://advertising.amazon.com/help/GCCXMZYCK4RXWS6C). Max input size 5MB. | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.dpuploadhashedrecordsrequest.v3+json
 - **Accept**: application/vnd.dpuploadhashedrecordsresponse.v3+json, application/vnd.dpuploadhashedrecordsresponse.v3.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

