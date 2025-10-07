# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateOptimizationRulesWithAdGroup**](OptimizationRulesBetaApi.md#AssociateOptimizationRulesWithAdGroup) | **Post** /sd/adGroups/{adGroupId}/optimizationRules | Associate one or more optimization rules to an ad group specified by identifier.
[**CreateOptimizationRules**](OptimizationRulesBetaApi.md#CreateOptimizationRules) | **Post** /sd/optimizationRules | Creates one or more optimization rules, also known as outcome optimizations.
[**DisassociateOptimizationRulesFromAdGroup**](OptimizationRulesBetaApi.md#DisassociateOptimizationRulesFromAdGroup) | **Post** /sd/adGroups/{adGroupId}/optimizationRules/disassociate | Disassociate one or more optimization rules from an ad group specified by identifier.
[**ListOptimizationRules**](OptimizationRulesBetaApi.md#ListOptimizationRules) | **Get** /sd/optimizationRules | Gets a list of optimization rules.
[**SdAdGroupsAdGroupIdOptimizationRulesGet**](OptimizationRulesBetaApi.md#SdAdGroupsAdGroupIdOptimizationRulesGet) | **Get** /sd/adGroups/{adGroupId}/optimizationRules | Gets a list of optimization rules associated to an adgroup specified by identifier.
[**SdOptimizationRulesOptimizationRuleIdGet**](OptimizationRulesBetaApi.md#SdOptimizationRulesOptimizationRuleIdGet) | **Get** /sd/optimizationRules/{optimizationRuleId} | Gets a requested optimization rule.
[**UpdateOptimizationRules**](OptimizationRulesBetaApi.md#UpdateOptimizationRules) | **Put** /sd/optimizationRules | Updates one or more optimization rules.

# **AssociateOptimizationRulesWithAdGroup**
> []OptimizationRuleResponse AssociateOptimizationRulesWithAdGroup(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, adGroupId, optional)
Associate one or more optimization rules to an ad group specified by identifier.

 * When an optimization rule is associated to an ad group, manual bids for individual targets will be overridden. * Only one optimization rule can be associated per adGroup. This note will be removed when multiple rules are supported per adGroup.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **adGroupId** | **int64**| The identifier of the ad group. | 
 **optional** | ***OptimizationRulesBetaApiAssociateOptimizationRulesWithAdGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OptimizationRulesBetaApiAssociateOptimizationRulesWithAdGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of CreateAssociatedOptimizationRulesRequest**](CreateAssociatedOptimizationRulesRequest.md)| A list of optimization rule identifiers. Only one optimization rule identifier is currently supported per request. This note will be removed when multiple rule identifiers are supported.

For each ad group, only one optimization rule metric name is supported, based on the ad group&#x27;s &#x60;bidOptimization&#x60; type. Refer to the following table for the metric names supported for each type.
|  AdGroup.bidOptimization |     Supported OptimizationRule.metricName       |
|------------------|--------------------|
|   reach       | COST_PER_THOUSAND_VIEWABLE_IMPRESSIONS  |
|   clicks      | COST_PER_CLICK          |
|  conversions  | COST_PER_ORDER          | | 

### Return type

[**[]OptimizationRuleResponse**](OptimizationRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOptimizationRules**
> []OptimizationRuleResponse CreateOptimizationRules(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Creates one or more optimization rules, also known as outcome optimizations.

 * When an optimization rule is associated to an ad group, manual bids for individual targets will be overridden. * Optimization rules can only be associated to ad groups that have productAds with ASIN or SKU. * We may add targets while your campaign is running to try to stay at or below a cost per order value you have specified. * Only one optimization rule can be associated per adGroup.  * If you are using optimization rules, the following campaign budget must be at least:   - 5x the value of any COST_PER_ORDER threshold.   - 10x the value of any COST_PER_THOUSAND_VIEWABLE_IMPRESSIONS threshold.   - 20x the value of any COST_PER_CLICK threshold.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***OptimizationRulesBetaApiCreateOptimizationRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OptimizationRulesBetaApiCreateOptimizationRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []CreateOptimizationRule**](CreateOptimizationRule.md)| An array of OptimizationRule objects. For each object, specify required fields and their values. Required fields are &#x60;state&#x60; and &#x60;ruleConditions&#x60;. | 

### Return type

[**[]OptimizationRuleResponse**](OptimizationRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisassociateOptimizationRulesFromAdGroup**
> OptimizationRuleAssociationResponse DisassociateOptimizationRulesFromAdGroup(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, adGroupId, optional)
Disassociate one or more optimization rules from an ad group specified by identifier.

 * Only one optimization rule can be disassociated per adGroup. This note will be removed when multiple rules are supported per adGroup. * When an optimization rule is disassociated from an ad group, you can set the manual bids for the individual targets under the adGroup.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **adGroupId** | **int64**| The identifier of the ad group. | 
 **optional** | ***OptimizationRulesBetaApiDisassociateOptimizationRulesFromAdGroupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OptimizationRulesBetaApiDisassociateOptimizationRulesFromAdGroupOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of CreateAssociatedOptimizationRulesRequest**](CreateAssociatedOptimizationRulesRequest.md)| A list of optimization rule identifiers. Only one optimization rule identifier is currently supported per request. This note will be removed when multiple rule identifiers are supported. | 

### Return type

[**OptimizationRuleAssociationResponse**](OptimizationRuleAssociationResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListOptimizationRules**
> []OptimizationRule ListOptimizationRules(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Gets a list of optimization rules.

Gets an array of OptimizationRule objects for a requested set of Sponsored Display optimization rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***OptimizationRulesBetaApiListOptimizationRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OptimizationRulesBetaApiListOptimizationRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startIndex** | **optional.Int32**| Optional. Sets a cursor into the requested set of optimization rules. Use in conjunction with the &#x60;count&#x60; parameter to control pagination of the returned array. 0-indexed record offset for the result set, defaults to 0. | 
 **count** | **optional.Int32**| Optional. Sets the number of OptimizationRule objects in the returned array. Use in conjunction with the &#x60;startIndex&#x60; parameter to control pagination. For example, to return the first ten optimization rules set &#x60;startIndex&#x3D;0&#x60; and &#x60;count&#x3D;10&#x60;. To return the next ten optimization rules, set &#x60;startIndex&#x3D;10&#x60; and &#x60;count&#x3D;10&#x60;, and so on. Defaults to max page size. | 
 **stateFilter** | **optional.String**| Optional. The returned array is filtered to include only optimization rules with state set to one of the values in the specified comma-delimited list. Available values:   - enabled   - paused [COMING LATER]   - enabled, paused [COMING LATER] | [default to enabled]
 **name** | **optional.String**| Optional. The returned array includes only optimization rules with the specified name using an exact string match. | 
 **optimizationRuleIdFilter** | **optional.String**| Optional. The returned array is filtered to include only optimization rules associated with the optimization rule identifiers in the specified comma-delimited list.  Maximum size limit 50. | 
 **adGroupIdFilter** | **optional.String**| Optional. The returned array is filtered to include only optimization rules associated with the ad group identifiers in the comma-delimited list.  Maximum size limit 50. | 

### Return type

[**[]OptimizationRule**](OptimizationRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SdAdGroupsAdGroupIdOptimizationRulesGet**
> []OptimizationRule SdAdGroupsAdGroupIdOptimizationRulesGet(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, adGroupId)
Gets a list of optimization rules associated to an adgroup specified by identifier.

Gets an OptimizationRule object for a requested Sponsored Display optimization rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **adGroupId** | **int64**| The identifier of the ad group. | 

### Return type

[**[]OptimizationRule**](OptimizationRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SdOptimizationRulesOptimizationRuleIdGet**
> OptimizationRule SdOptimizationRulesOptimizationRuleIdGet(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optimizationRuleId)
Gets a requested optimization rule.

Gets an OptimizationRule object for a requested Sponsored Display optimization rule.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **optimizationRuleId** | **string**| The identifier of the requested optimization rule. | 

### Return type

[**OptimizationRule**](OptimizationRule.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOptimizationRules**
> []OptimizationRuleResponse UpdateOptimizationRules(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Updates one or more optimization rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***OptimizationRulesBetaApiUpdateOptimizationRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OptimizationRulesBetaApiUpdateOptimizationRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of []UpdateOptimizationRule**](UpdateOptimizationRule.md)| An array of OptimizationRule objects. For each object, specify an optimization rule identifier and mutable fields with their updated values. The mutable fields are &#x60;ruleName&#x60;, &#x60;state&#x60;, and &#x60;ruleConditions&#x60;. | 

### Return type

[**[]OptimizationRuleResponse**](OptimizationRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

