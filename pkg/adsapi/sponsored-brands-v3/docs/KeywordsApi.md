# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveKeyword**](KeywordsApi.md#ArchiveKeyword) | **Delete** /sb/keywords/{keywordId} | Archives a keyword specified by identifier.
[**CreateKeywords**](KeywordsApi.md#CreateKeywords) | **Post** /sb/keywords | Creates one or more keywords.
[**GetKeyword**](KeywordsApi.md#GetKeyword) | **Get** /sb/keywords/{keywordId} | Gets a keyword specified by identifier.
[**ListKeywords**](KeywordsApi.md#ListKeywords) | **Get** /sb/keywords | Gets an array of keywords, filtered by optional criteria.
[**UpdateKeywords**](KeywordsApi.md#UpdateKeywords) | **Put** /sb/keywords | Updates one or more keywords.

# **ArchiveKeyword**
> SbKeywordResponse ArchiveKeyword(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, keywordId)
Archives a keyword specified by identifier.

This operation is equivalent to an update operation that sets the status field to 'archived'. Note that setting the status field to 'archived' is permanent and can't be undone. See [Developer Notes](https://advertising.amazon.com/API/docs/v2/guides/developer_notes) for more information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **keywordId** | **int64**|  | 

### Return type

[**SbKeywordResponse**](SBKeywordResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sbkeywordresponse.v3+json, application/vnd.error.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateKeywords**
> SbKeywordResponse CreateKeywords(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Creates one or more keywords.

Note that `state` can't be set at keyword creation. <br/>Note that keywords can be created on campaigns where serving status is not one of `archived`, `terminated`, `rejected`, or `ended`. <br/>Note that this operation supports a maximum list size of 100 keywords.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***KeywordsApiCreateKeywordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KeywordsApiCreateKeywordsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []SbKeywordsBody1**](sb_keywords_body_1.md)| An array of keywords. &lt;br/&gt;Note the bid is only mutable when the keyword&#x27;s corresponding campaign does not have any enabled optimization rule. | 

### Return type

[**SbKeywordResponse**](SBKeywordResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sbkeywordresponse.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKeyword**
> SbKeyword GetKeyword(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, keywordId, optional)
Gets a keyword specified by identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **keywordId** | **int64**| The identifier of an existing keyword. | 
 **optional** | ***KeywordsApiGetKeywordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KeywordsApiGetKeywordOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **locale** | **optional.String**| The returned array includes only keywords associated with locale matching those specified by identifier. | 

### Return type

[**SbKeyword**](SBKeyword.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sbkeyword.v3+json, application/vnd.error.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListKeywords**
> []SbKeyword ListKeywords(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets an array of keywords, filtered by optional criteria.

**Note**: Keywords associated with BrandVideo ad groups are only available in v3.2 version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***KeywordsApiListKeywordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KeywordsApiListKeywordsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Sets a zero-based offset into the requested set of keywords. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. | [default to 0]
 **count** | **optional.Int32**| Sets the number of keywords in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten keywords set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten keywords, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. | 
 **matchTypeFilter** | [**optional.Interface of MatchType**](.md)| The returned array is filtered to include only keywords with &#x60;matchType&#x60; set to one of the values in the specified comma-delimited list. | 
 **keywordText** | **optional.String**| The returned array includes only keywords with the specified text. | 
 **stateFilter** | **optional.String**| The returned array is filtered to include only keywords with &#x27;state&#x27; set to one of the values in the specified comma-delimited list. | [default to enabled,paused]
 **campaignIdFilter** | **optional.String**| The returned array includes only keywords associated with campaigns matching those specified by identifier in the comma-delimited string. | 
 **adGroupIdFilter** | **optional.String**| The returned array includes only keywords associated with ad groups matching those specified by identifier in the comma-delimited string. | 
 **keywordIdFilter** | **optional.String**| The returned array includes only keywords with identifiers matching those specified in the comma-delimited string. | 
 **creativeType** | [**optional.Interface of CreativeType**](.md)| Filter by the type of creative the campaign is associated with. To get keywords associated with non-video campaigns specify &#x27;productCollection&#x27;. To get keywords associated with video campaigns, this must be set to &#x27;video&#x27;. Returns all keywords if not specified. | 
 **locale** | **optional.String**| The returned array includes only keywords with locale matching those specified string. | 

### Return type

[**[]SbKeyword**](SBKeyword.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sbkeyword.v3.2+json, application/vnd.sbkeyword.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateKeywords**
> SbKeywordResponse UpdateKeywords(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Updates one or more keywords.

Updates one or more targeting clauses. Operation supports a maximum list size of 100 keywords. <bold>Note</bold> that negative keywords can be updated on campaigns where campaign's serving status is not one of `archived`, `terminated`, `rejected`, or `ended`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***KeywordsApiUpdateKeywordsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a KeywordsApiUpdateKeywordsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []SbKeywordsBody**](sb_keywords_body.md)| An array of keywords. &lt;br/&gt;Note the bid is only mutable when the keyword&#x27;s corresponding campaign does not have any enabled optimization rule. | 

### Return type

[**SbKeywordResponse**](SBKeywordResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sbkeywordresponse.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

