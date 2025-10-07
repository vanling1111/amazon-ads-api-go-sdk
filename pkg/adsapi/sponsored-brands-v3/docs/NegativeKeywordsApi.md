# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveNegativeKeyword**](NegativeKeywordsApi.md#ArchiveNegativeKeyword) | **Delete** /sb/negativeKeywords/{keywordId} | Archives a negative keyword specified by identifier.
[**CreateNegativeKeywords**](NegativeKeywordsApi.md#CreateNegativeKeywords) | **Post** /sb/negativeKeywords | Creates one or more negative keywords.
[**GetNegativeKeyword**](NegativeKeywordsApi.md#GetNegativeKeyword) | **Get** /sb/negativeKeywords/{keywordId} | Gets a negative keyword specified by identifier.
[**ListNegativeKeywords**](NegativeKeywordsApi.md#ListNegativeKeywords) | **Get** /sb/negativeKeywords | Gets an array of negative keywords, filtered by optional criteria.
[**UpdateNegativeKeywords**](NegativeKeywordsApi.md#UpdateNegativeKeywords) | **Put** /sb/negativeKeywords | Updates one or more negative keywords.

# **ArchiveNegativeKeyword**
> SbKeywordResponse ArchiveNegativeKeyword(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, keywordId)
Archives a negative keyword specified by identifier.

This operation is equivalent to an update operation that sets the status field to 'archived'. Note that setting the status field to 'archived' is permanent and can't be undone. See [Developer Notes](https://advertising.amazon.com/API/docs/v2/guides/developer_notes) for more information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **keywordId** | **int64**| The identifier of an existing campaign. | 

### Return type

[**SbKeywordResponse**](SBKeywordResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sbkeywordresponse.v3+json, application/vnd.error.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNegativeKeywords**
> []SbKeywordResponse CreateNegativeKeywords(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Creates one or more negative keywords.

Creates one or more negative targeting clauses. Operation supports a maximum list size of 100 negative keywords. The `bid` and `state` can't be set at negative keyword creation. <br/>Note that negative keywords can be created on campaigns where campaign's serving status is not one of `archived`, `terminated`, `rejected`, or `ended`. <br>**Note** that negative keywords *can not* be recreated for a campaign if the negative keyword has previously been associated with a campaign and subsequently archived. 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***NegativeKeywordsApiCreateNegativeKeywordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegativeKeywordsApiCreateNegativeKeywordsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []SbNegativeKeywordsBody1**](sb_negativeKeywords_body_1.md)| An array of negative keywords. | 

### Return type

[**[]SbKeywordResponse**](SBKeywordResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sbkeywordresponse.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNegativeKeyword**
> InlineResponse2001 GetNegativeKeyword(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, keywordId)
Gets a negative keyword specified by identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **keywordId** | **int64**| The identifier of an existing negative keyword. | 

### Return type

[**InlineResponse2001**](inline_response_200_1.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sbnegativekeyword.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNegativeKeywords**
> []SbNegativeKeyword ListNegativeKeywords(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets an array of negative keywords, filtered by optional criteria.

**Note**: Negative keywords associated with BrandVideo ad groups are only available in v3.2 version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***NegativeKeywordsApiListNegativeKeywordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegativeKeywordsApiListNegativeKeywordsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Sets a zero-based offset into the requested set of negative keywords. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. | [default to 0]
 **count** | **optional.Int32**| Sets the number of negative keywords in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten negative keywords set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten negative keywords, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. | 
 **matchTypeFilter** | [**optional.Interface of NegativeMatchType**](.md)| The returned array is filtered to include only negative keywords with &#x60;matchType&#x60; set to one of the values in the specified comma-delimited list. | 
 **keywordText** | **optional.String**| The returned array includes only negative keywords with the specified text. | 
 **stateFilter** | [**optional.Interface of NegativeState**](.md)| The returned array includes only negative keywords with &#x60;state&#x60; set to the specified value. | 
 **campaignIdFilter** | **optional.String**| The returned array includes only negative keywords associated with campaigns matching those specified by identifier in the comma-delimited string. | 
 **adGroupIdFilter** | **optional.String**| The returned array includes only negative keywords associated with ad groups matching those specified by identifier in the comma-delimited string. | 
 **keywordIdFilter** | **optional.String**| The returned array includes only negative keywords with identifiers matching those specified in the comma-delimited string. | 
 **creativeType** | [**optional.Interface of CreativeType**](.md)| Filter by the type of creative the campaign is associated with. To get negative keywords associated with non-video campaigns specify &#x27;productCollection&#x27;. To get negative keywords associated with video campaigns, this must be set to &#x27;video&#x27;. Returns all negative keywords if not specified. | 

### Return type

[**[]SbNegativeKeyword**](SBNegativeKeyword.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sbnegativekeyword.v3.2+json, application/vnd.sbnegativekeyword.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNegativeKeywords**
> []SbKeywordResponse UpdateNegativeKeywords(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Updates one or more negative keywords.

Updates one or more targeting clauses. Operation supports a maximum list size of 100 negative keywords. <bold>Note</bold> that negative keywords can be updated on campaigns where campaign's serving status is not one of `archived`, `terminated`, `rejected`, or `ended`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***NegativeKeywordsApiUpdateNegativeKeywordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegativeKeywordsApiUpdateNegativeKeywordsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []SbNegativeKeywordsBody**](sb_negativeKeywords_body.md)| An array of negative keywords. | 

### Return type

[**[]SbKeywordResponse**](SBKeywordResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sbkeywordresponse.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

