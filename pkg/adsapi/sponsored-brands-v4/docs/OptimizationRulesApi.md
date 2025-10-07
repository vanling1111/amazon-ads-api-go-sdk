# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateSponsoredBrandsOptimizationRules**](OptimizationRulesApi.md#AssociateSponsoredBrandsOptimizationRules) | **Post** /sb/rules/optimization/associate | Associate optimization rules
[**CreateSponsoredBrandsOptimizationRules**](OptimizationRulesApi.md#CreateSponsoredBrandsOptimizationRules) | **Post** /sb/rules/optimization | Create optimization rules
[**DisassociateSponsoredBrandsOptimizationRules**](OptimizationRulesApi.md#DisassociateSponsoredBrandsOptimizationRules) | **Post** /sb/rules/optimization/disassociate | Disassociate optimization rules
[**ListSponsoredBrandsOptimizationRules**](OptimizationRulesApi.md#ListSponsoredBrandsOptimizationRules) | **Post** /sb/rules/optimization/list | List optimization rules
[**UpdateSponsoredBrandsOptimizationRules**](OptimizationRulesApi.md#UpdateSponsoredBrandsOptimizationRules) | **Put** /sb/rules/optimization | Update optimization rules

# **AssociateSponsoredBrandsOptimizationRules**
> AssociateSponsoredBrandsOptimizationRulesResponseContent AssociateSponsoredBrandsOptimizationRules(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Associate optimization rules

Currently available in beta. Associate one or more optimization rules by providing combinations of entityId-ruleId that require association.  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AssociateSponsoredBrandsOptimizationRulesRequestContent**](AssociateSponsoredBrandsOptimizationRulesRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**AssociateSponsoredBrandsOptimizationRulesResponseContent**](AssociateSponsoredBrandsOptimizationRulesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbruleoptimization.v4+json
 - **Accept**: application/vnd.sbruleoptimization.v4+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSponsoredBrandsOptimizationRules**
> CreateSponsoredBrandsOptimizationRulesResponseContent CreateSponsoredBrandsOptimizationRules(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Create optimization rules

Currently available in beta.   **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSponsoredBrandsOptimizationRulesRequestContent**](CreateSponsoredBrandsOptimizationRulesRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**CreateSponsoredBrandsOptimizationRulesResponseContent**](CreateSponsoredBrandsOptimizationRulesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbruleoptimization.v4+json
 - **Accept**: application/vnd.sbruleoptimization.v4+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisassociateSponsoredBrandsOptimizationRules**
> DisassociateSponsoredBrandsOptimizationRulesResponseContent DisassociateSponsoredBrandsOptimizationRules(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Disassociate optimization rules

Currently available in beta. Disassociate one or more optimization rules by providing combinations of entityId-ruleId that require disassociation  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DisassociateSponsoredBrandsOptimizationRulesRequestContent**](DisassociateSponsoredBrandsOptimizationRulesRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**DisassociateSponsoredBrandsOptimizationRulesResponseContent**](DisassociateSponsoredBrandsOptimizationRulesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbruleoptimization.v4+json
 - **Accept**: application/vnd.sbruleoptimization.v4+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListSponsoredBrandsOptimizationRules**
> ListSponsoredBrandsOptimizationRulesResponseContent ListSponsoredBrandsOptimizationRules(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
List optimization rules

Currently available in beta.   **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***OptimizationRulesApiListSponsoredBrandsOptimizationRulesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a OptimizationRulesApiListSponsoredBrandsOptimizationRulesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ListSponsoredBrandsOptimizationRulesRequestContent**](ListSponsoredBrandsOptimizationRulesRequestContent.md)|  | 

### Return type

[**ListSponsoredBrandsOptimizationRulesResponseContent**](ListSponsoredBrandsOptimizationRulesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbruleoptimization.v4+json
 - **Accept**: application/vnd.sbruleoptimization.v4+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSponsoredBrandsOptimizationRules**
> UpdateSponsoredBrandsOptimizationRulesResponseContent UpdateSponsoredBrandsOptimizationRules(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Update optimization rules

Currently available in beta.   **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSponsoredBrandsOptimizationRulesRequestContent**](UpdateSponsoredBrandsOptimizationRulesRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**UpdateSponsoredBrandsOptimizationRulesResponseContent**](UpdateSponsoredBrandsOptimizationRulesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbruleoptimization.v4+json
 - **Accept**: application/vnd.sbruleoptimization.v4+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

