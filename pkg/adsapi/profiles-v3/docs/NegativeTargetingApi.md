# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveNegativeTargetingClause**](NegativeTargetingApi.md#ArchiveNegativeTargetingClause) | **Delete** /sd/negativeTargets/{negativeTargetId} | Sets the &#x60;state&#x60; of a negative targeting clause to &#x60;archived&#x60;.
[**CreateNegativeTargetingClauses**](NegativeTargetingApi.md#CreateNegativeTargetingClauses) | **Post** /sd/negativeTargets | Creates one or more negative targeting clauses.
[**GetNegativeTargets**](NegativeTargetingApi.md#GetNegativeTargets) | **Get** /sd/negativeTargets/{negativeTargetId} | Gets a negative targeting clause specified by identifier.
[**GetNegativeTargetsEx**](NegativeTargetingApi.md#GetNegativeTargetsEx) | **Get** /sd/negativeTargets/extended/{negativeTargetId} | Gets extended information for a negative targeting clause.
[**ListNegativeTargetingClauses**](NegativeTargetingApi.md#ListNegativeTargetingClauses) | **Get** /sd/negativeTargets | Gets a list of negative targeting clauses.
[**ListNegativeTargetingClausesEx**](NegativeTargetingApi.md#ListNegativeTargetingClausesEx) | **Get** /sd/negativeTargets/extended | Gets a list of negative targeting clause objects with extended fields.
[**UpdateNegativeTargetingClauses**](NegativeTargetingApi.md#UpdateNegativeTargetingClauses) | **Put** /sd/negativeTargets | Updates one or more negative targeting clauses.

# **ArchiveNegativeTargetingClause**
> TargetResponse ArchiveNegativeTargetingClause(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, negativeTargetId)
Sets the `state` of a negative targeting clause to `archived`.

Equivalent to using the updateNegativeTargetingClauses operation to set the `state` property of a targeting clause to `archived`. See [Developer Notes](http://advertising.amazon.com/API/docs/guides/developer_notes#Archiving) for more information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **negativeTargetId** | **int64**| The identifier of a negative targeting clause. | 

### Return type

[**TargetResponse**](TargetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNegativeTargetingClauses**
> []TargetResponse CreateNegativeTargetingClauses(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Creates one or more negative targeting clauses.

Successfully created negative targeting clauses associated with an ad group are assigned a unique target identifier. Product negative targeting clause examples: | Negative targeting clause | Description | |---------------------------|-------------| | asinSameAs=B0123456789 | Negatively target this product.| | asinBrandSameAs=12345 | Negatively target products in the brand.|

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***NegativeTargetingApiCreateNegativeTargetingClausesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegativeTargetingApiCreateNegativeTargetingClausesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []CreateNegativeTargetingClause**](CreateNegativeTargetingClause.md)| A list of up to 100 negative targeting clauses for creation. | 

### Return type

[**[]TargetResponse**](TargetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNegativeTargets**
> NegativeTargetingClause GetNegativeTargets(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, negativeTargetId)
Gets a negative targeting clause specified by identifier.

This call returns the minimal set of negative targeting clause fields, but is more efficient than getNegativeTargetsEx.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **negativeTargetId** | **int64**| The negative targeting clause identifier. | 

### Return type

[**NegativeTargetingClause**](NegativeTargetingClause.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNegativeTargetsEx**
> NegativeTargetingClauseEx GetNegativeTargetsEx(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, negativeTargetId)
Gets extended information for a negative targeting clause.

Gets a negative targeting clause with extended fields. Note that this call returns the full set of negative targeting clause extended fields, but is less efficient than getNegativeTarget.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **negativeTargetId** | **int64**| The negative targeting clause identifier. | 

### Return type

[**NegativeTargetingClauseEx**](NegativeTargetingClauseEx.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNegativeTargetingClauses**
> []NegativeTargetingClause ListNegativeTargetingClauses(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of negative targeting clauses.

Gets a list of negative targeting clauses objects for a requested set of Sponsored Display negative targets. Note that the Negative Targeting Clause object is designed for performance, and includes a small set of commonly used fields to reduce size. If the extended set of fields is required, use the negative target operations that return the NegativeTargetingClauseEx object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***NegativeTargetingApiListNegativeTargetingClausesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegativeTargetingApiListNegativeTargetingClausesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Optional. 0-indexed record offset for the result set. Defaults to 0. | 
 **count** | **optional.Int32**| Optional. Number of records to include in the paged response. Defaults to max page size. | 
 **stateFilter** | **optional.String**| Optional. Restricts results to those with state within the specified comma-separated list. Must be one of: &#x60;enabled&#x60;, &#x60;paused&#x60;, or &#x60;archived&#x60;. Default behavior is to include enabled, paused, and archived. | [default to enabled, paused, archived]
 **adGroupIdFilter** | **optional.String**| Optional list of comma separated adGroupIds. Restricts results to negative targeting clauses with the specified &#x60;adGroupId&#x60;. | 
 **campaignIdFilter** | **optional.String**| Optional. Restricts results to targeting clauses within campaigns specified in comma-separated list. | 

### Return type

[**[]NegativeTargetingClause**](NegativeTargetingClause.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNegativeTargetingClausesEx**
> []NegativeTargetingClauseEx ListNegativeTargetingClausesEx(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of negative targeting clause objects with extended fields.

Gets an array of NegativeTargetingClauseEx objects for a set of requested negative targets. Note that this call returns the full set of negative targeting clause extended fields, but is less efficient than getNegativeTargets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***NegativeTargetingApiListNegativeTargetingClausesExOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegativeTargetingApiListNegativeTargetingClausesExOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Optional. 0-indexed record offset for the result set. Defaults to 0. | 
 **count** | **optional.Int32**| Optional. Number of records to include in the paged response. Defaults to max page size. | 
 **stateFilter** | **optional.String**| Optional. Restricts results to keywords with state within the specified comma-separated list. Must be one of: &#x60;enabled&#x60;, &#x60;paused&#x60;, or &#x60;archived&#x60;. Default behavior is to include &#x60;enabled&#x60;, &#x60;paused&#x60;, and &#x60;archived&#x60;. | [default to enabled, paused, archived]
 **targetIdFilter** | **optional.String**| Optional. Restricts results to ads with the specified &#x60;tagetId&#x60; specified in comma-separated list | 
 **adGroupIdFilter** | **optional.String**| Optional list of comma separated adGroupIds. Restricts results to negative targeting clauses with the specified &#x60;adGroupId&#x60;. | 
 **campaignIdFilter** | **optional.String**| Optional. Restricts results to ads within campaigns specified in the comma-separated list. | 

### Return type

[**[]NegativeTargetingClauseEx**](NegativeTargetingClauseEx.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNegativeTargetingClauses**
> []TargetResponse UpdateNegativeTargetingClauses(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Updates one or more negative targeting clauses.

Updates one or more negative targeting clauses. Negative targeting clauses are identified using their targetId. The mutable field is `state`. Maximum length of the array is 100 objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***NegativeTargetingApiUpdateNegativeTargetingClausesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegativeTargetingApiUpdateNegativeTargetingClausesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []UpdateNegativeTargetingClause**](UpdateNegativeTargetingClause.md)| A list of up to 100 negative targeting clauses. Note that the only mutable field is &#x60;state&#x60;. | 

### Return type

[**[]TargetResponse**](TargetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

