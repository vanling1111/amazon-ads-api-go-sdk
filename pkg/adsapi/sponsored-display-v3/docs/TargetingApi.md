# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveTargetingClause**](TargetingApi.md#ArchiveTargetingClause) | **Delete** /sd/targets/{targetId} | Sets the &#x60;state&#x60; of a targeting clause to &#x60;archived&#x60;.
[**CreateTargetingClauses**](TargetingApi.md#CreateTargetingClauses) | **Post** /sd/targets | Creates one or more targeting clauses.
[**GetTargets**](TargetingApi.md#GetTargets) | **Get** /sd/targets/{targetId} | Gets a targeting clause specified by identifier.
[**GetTargetsEx**](TargetingApi.md#GetTargetsEx) | **Get** /sd/targets/extended/{targetId} | Gets extended information for a targeting clause.
[**ListTargetingClauses**](TargetingApi.md#ListTargetingClauses) | **Get** /sd/targets | Gets a list of targeting clauses.
[**ListTargetingClausesEx**](TargetingApi.md#ListTargetingClausesEx) | **Get** /sd/targets/extended | Gets a list of targeting clause objects with extended fields.
[**UpdateTargetingClauses**](TargetingApi.md#UpdateTargetingClauses) | **Put** /sd/targets | Updates one or more targeting clauses.

# **ArchiveTargetingClause**
> TargetResponse ArchiveTargetingClause(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, targetId)
Sets the `state` of a targeting clause to `archived`.

Equivalent to using the `updateTargetingClauses` operation to set the `state` property of a targeting clause to `archived`. See [Developer Notes](http://advertising.amazon.com/API/docs/guides/developer_notes#Archiving) for more information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **targetId** | **int64**| The identifer of a targeting clause. | 

### Return type

[**TargetResponse**](TargetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTargetingClauses**
> []TargetResponse CreateTargetingClauses(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Creates one or more targeting clauses.

Successfully created targeting clauses are assigned a unique `targetId` value.  Create new targeting clauses for **both** Audience and Contextual campaigns with tactic 'T00030' using the above and the following: | Contextual targeting clause | Description | |------------------|-------------| | similarProduct | Dynamic segment to target products that are similar to the advertised asin. We recommend using 'similarProduct' targeting for all adGroups. | | asinSameAs=B0123456789 | Target this product. | | asinCategorySameAs=12345 | Target products in the category. | | asinCategorySameAs=12345 asinBrandSameAs=45678 | Target products in the category and brand. |  **Refinements:** - asinBrandSameAs - asinPriceBetween - asinPriceGreaterThan - asinPriceLessThan - asinReviewRatingLessThan - asinReviewRatingGreaterThan - asinReviewRatingBetween - asinIsPrimeShippingEligible - asinAgeRangeSameAs - asinGenreSameAs  **Refinement Notes:** * Brand, price, and review predicates are optional and may only be specified if category is also specified. * Review predicates accept numbers between 0 and 5 and are inclusive. * When using either of the 'between' strings to construct a targeting expression the format of the string is 'double-double' where the first double must be smaller than the second double. Prices are not inclusive. * 'similarProduct' has no expression value or refinements.  | Audience targeting clause | Description | |------------------|-------------| | views(exactProduct lookback=30) | Target an audience that has viewed the advertised asins in the past 7,14,30,60, or 90 days. Note: This target should only be used for productAds with SKU or ASIN. | | views(similarProduct lookback=60) | Target an audience that has viewed similar products to the advertised asins in the past 7,14,30,60, or 90 days. Note: This target should only be used for productAds with SKU or ASIN.| | views(asinCategorySameAs=12345 lookback=90) | Target an audience that has viewed products in the given category in the past 7,14,30,60, or 90 days. | | views(asinCategorySameAs=12345 asinBrandSameAs=45678 asinPriceBetween=50-100 lookback=60) | Target an audience that has viewed products in the given category, brand, and price range in the past 7,14,30,60, or 90 days. | | purchases(relatedProduct lookback=180) | Target an audience that has purchased a related product in the past 7,14,30,60,90,180 or 365 days. Note: This target should only be used for productAds with SKU or ASIN.| | purchases(exactProduct lookback=365) | Target an audience that has purchased the advertised asins in the past 7,14,30,60,90,180 or 365 days. Note: This target should only be used for productAds with SKU or ASIN.| | purchases(asinCategorySameAs=12345 asinBrandSameAs=45678 asinPriceBetween=50-100 lookback=90) | Target an audience that has purchased products in the given category, brand, and price range in the past 7,14,30,60,90,180 or 365 days | | audiencesLikelyInterestedInAd | Dynamic segment to target audiences that are likely to consider and buy from your business. We recommend that you add this segment to all of your campaigns to help you achieve greater reach and campaign scale. Note: this is only supported when using landingPageType of OFF_AMAZON_LINK | | audience(audienceSameAs=12345) | Target the given Amazon Audience. |  | Content targeting clause | Description | |------------------|-------------| | contentCategorySameAs=amzn1.iab-content.325 | Target an audience that has engaged with or shown interest in the given entertainment category. |  Notes on content targeting: * The `contentCategorySameAs` targeting predicate is required * For all iab categories, see [Discovering entertainment target categories](https://advertising.amazon.com/API/docs/en-us/guides/sponsored-display/entertainment-targeting#discovering-entertainment-target-categories).  Note: 1. You can still create new targeting clauses for Contextual campaigns with tactic 'T00020' using the Contextual clauses above. 2. There is a limit of 200 targeting clauses per request for T00030. 3. There is a limit of 1000 targeting clauses per request for T00020. 4. There is a total limit of 1000 targeting clauses per ad group. 5. If you receive the error of \"Cannot create targeting clause: audience size is too small\", please expand or broaden your targeting clause to increase the audience size.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***TargetingApiCreateTargetingClausesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingApiCreateTargetingClausesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []CreateTargetingClause**](CreateTargetingClause.md)| A list of targeting clauses for creation. | 

### Return type

[**[]TargetResponse**](TargetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTargets**
> TargetingClause GetTargets(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, targetId)
Gets a targeting clause specified by identifier.

This call returns the minimal set of targeting clause fields.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **targetId** | **int64**| The identifier of a targeting clause. | 

### Return type

[**TargetingClause**](TargetingClause.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTargetsEx**
> TargetingClauseEx GetTargetsEx(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, targetId)
Gets extended information for a targeting clause.

Gets a targeting clause object with extended fields. Note that this call returns the full set of targeting clause extended fields, but is less efficient than getTarget.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **targetId** | **int64**| The identifier of a targeting clause. | 

### Return type

[**TargetingClauseEx**](TargetingClauseEx.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTargetingClauses**
> []TargetingClause ListTargetingClauses(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of targeting clauses.

Gets a list of targeting clauses objects for a requested set of Sponsored Display targets. Note that the Targeting Clause object is designed for performance, and includes a small set of commonly used fields to reduce size. If the extended set of fields is required, use the target operations that return the TargetingClauseEx object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***TargetingApiListTargetingClausesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingApiListTargetingClausesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Optional. 0-indexed record offset for the result set. Defaults to 0. | 
 **count** | **optional.Int32**| Optional. Number of records to include in the paged response. Defaults to max page size. | 
 **stateFilter** | **optional.String**| Optional. Restricts results to those with &#x60;state&#x60; set to values in the specified comma-separated list. | [default to enabled, paused, archived]
 **targetIdFilter** | **optional.String**| Optional. Restricts results to ads with the specified &#x60;tagetId&#x60; specified in comma-separated list | 
 **adGroupIdFilter** | **optional.String**| Optional list of comma separated adGroupIds. Restricts results to targeting clauses with the specified &#x60;adGroupId&#x60;. | 
 **campaignIdFilter** | **optional.String**| Optional. Restricts results to targeting clauses within campaigns specified in comma-separated list. | 

### Return type

[**[]TargetingClause**](TargetingClause.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTargetingClausesEx**
> []TargetingClauseEx ListTargetingClausesEx(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of targeting clause objects with extended fields.

Gets an array of TargetingClauseEx objects for a set of requested targets. Note that this call returns the full set of targeting clause extended fields, but is less efficient than getTargets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***TargetingApiListTargetingClausesExOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingApiListTargetingClausesExOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Optional. 0-indexed record offset for the result set. Defaults to 0. | 
 **count** | **optional.Int32**| Optional. Number of records to include in the paged response. Defaults to max page size. | 
 **stateFilter** | **optional.String**| Optional. Restricts results to keywords with state within the specified comma-separated list. Must be one of: &#x60;enabled&#x60;, &#x60;paused&#x60;, or &#x60;archived&#x60;. Default behavior is to include enabled, paused, and archived. | [default to enabled, paused, archived]
 **targetIdFilter** | **optional.String**| Optional. Restricts results to ads with the specified &#x60;tagetId&#x60; specified in comma-separated list | 
 **adGroupIdFilter** | **optional.String**| Optional list of comma separated adGroupIds. Restricts results to targeting clauses with the specified &#x60;adGroupId&#x60;. | 
 **campaignIdFilter** | **optional.String**| Optional. Restricts results to ads within campaigns specified in comma-separated list. | 

### Return type

[**[]TargetingClauseEx**](TargetingClauseEx.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTargetingClauses**
> []TargetResponse UpdateTargetingClauses(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Updates one or more targeting clauses.

Updates one or more targeting clauses. Targeting clauses are identified using their targetId. The mutable fields are `bid` and `state`. Maximum length of the array is 100 objects.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***TargetingApiUpdateTargetingClausesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a TargetingApiUpdateTargetingClausesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []UpdateTargetingClause**](UpdateTargetingClause.md)| A list of up to 1000 targeting clauses. Mutable fields:
* &#x60;state&#x60;
* &#x60;bid&#x60; (only mutable when the targeting clause&#x27;s adGroup does not have any enabled optimization rule) | 

### Return type

[**[]TargetResponse**](TargetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

