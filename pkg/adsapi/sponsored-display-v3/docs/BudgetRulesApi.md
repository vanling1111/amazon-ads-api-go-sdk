# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssociatedBudgetRulesForSDCampaigns**](BudgetRulesApi.md#CreateAssociatedBudgetRulesForSDCampaigns) | **Post** /sd/campaigns/{campaignId}/budgetRules | Associates one or more budget rules to a campaign specified by identifer.
[**CreateBudgetRulesForSDCampaigns**](BudgetRulesApi.md#CreateBudgetRulesForSDCampaigns) | **Post** /sd/budgetRules | Creates one or more budget rules.
[**DisassociateAssociatedBudgetRuleForSDCampaigns**](BudgetRulesApi.md#DisassociateAssociatedBudgetRuleForSDCampaigns) | **Delete** /sd/campaigns/{campaignId}/budgetRules/{budgetRuleId} | Disassociates a budget rule specified by identifier from a campaign specified by identifier.
[**GetBudgetRuleByRuleIdForSDCampaigns**](BudgetRulesApi.md#GetBudgetRuleByRuleIdForSDCampaigns) | **Get** /sd/budgetRules/{budgetRuleId} | Gets a budget rule specified by identifier.
[**GetCampaignsAssociatedWithSDBudgetRule**](BudgetRulesApi.md#GetCampaignsAssociatedWithSDBudgetRule) | **Get** /sd/budgetRules/{budgetRuleId}/campaigns | Gets all the campaigns associated with a budget rule
[**GetSDBudgetRulesForAdvertiser**](BudgetRulesApi.md#GetSDBudgetRulesForAdvertiser) | **Get** /sd/budgetRules | Get all budget rules created by an advertiser
[**ListAssociatedBudgetRulesForSDCampaigns**](BudgetRulesApi.md#ListAssociatedBudgetRulesForSDCampaigns) | **Get** /sd/campaigns/{campaignId}/budgetRules | Gets a list of budget rules associated to a campaign specified by identifier.
[**UpdateBudgetRulesForSDCampaigns**](BudgetRulesApi.md#UpdateBudgetRulesForSDCampaigns) | **Put** /sd/budgetRules | Update one or more budget rules.

# **CreateAssociatedBudgetRulesForSDCampaigns**
> CreateAssociatedBudgetRulesResponse CreateAssociatedBudgetRulesForSDCampaigns(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, campaignId)
Associates one or more budget rules to a campaign specified by identifer.

A maximum of 250 rules can be associated to a campaign. Note that the name of each rule associated to a campaign is required to be unique.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateAssociatedBudgetRulesRequest**](CreateAssociatedBudgetRulesRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 
  **campaignId** | **float64**| The campaign identifier. | 

### Return type

[**CreateAssociatedBudgetRulesResponse**](CreateAssociatedBudgetRulesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBudgetRulesForSDCampaigns**
> CreateBudgetRulesResponse CreateBudgetRulesForSDCampaigns(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Creates one or more budget rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSdBudgetRulesRequest**](CreateSdBudgetRulesRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 

### Return type

[**CreateBudgetRulesResponse**](CreateBudgetRulesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisassociateAssociatedBudgetRuleForSDCampaigns**
> DisassociateAssociatedBudgetRuleResponse DisassociateAssociatedBudgetRuleForSDCampaigns(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, campaignId, budgetRuleId)
Disassociates a budget rule specified by identifier from a campaign specified by identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 
  **campaignId** | **float64**| The campaign identifier. | 
  **budgetRuleId** | **string**| The budget rule identifier. | 

### Return type

[**DisassociateAssociatedBudgetRuleResponse**](DisassociateAssociatedBudgetRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBudgetRuleByRuleIdForSDCampaigns**
> GetSdBudgetRuleResponse GetBudgetRuleByRuleIdForSDCampaigns(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, budgetRuleId)
Gets a budget rule specified by identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 
  **budgetRuleId** | **string**| The budget rule identifier. | 

### Return type

[**GetSdBudgetRuleResponse**](GetSDBudgetRuleResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignsAssociatedWithSDBudgetRule**
> SdGetAssociatedCampaignsResponse GetCampaignsAssociatedWithSDBudgetRule(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, budgetRuleId, pageSize, optional)
Gets all the campaigns associated with a budget rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 
  **budgetRuleId** | **string**| The budget rule identifier. | 
  **pageSize** | **float64**| Sets a limit on the number of results returned. Maximum limit of &#x60;pageSize&#x60; is 30. | 
 **optional** | ***BudgetRulesApiGetCampaignsAssociatedWithSDBudgetRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BudgetRulesApiGetCampaignsAssociatedWithSDBudgetRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **nextToken** | **optional.String**| To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;nextToken&#x60; field is empty, there are no further results. | 

### Return type

[**SdGetAssociatedCampaignsResponse**](SDGetAssociatedCampaignsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSDBudgetRulesForAdvertiser**
> GetSdBudgetRulesForAdvertiserResponse GetSDBudgetRulesForAdvertiser(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, pageSize, optional)
Get all budget rules created by an advertiser

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 
  **pageSize** | **float64**| Sets a limit on the number of results returned. Maximum limit of &#x60;pageSize&#x60; is 30. | 
 **optional** | ***BudgetRulesApiGetSDBudgetRulesForAdvertiserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BudgetRulesApiGetSDBudgetRulesForAdvertiserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nextToken** | **optional.String**| To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;nextToken&#x60; field is empty, there are no further results. | 

### Return type

[**GetSdBudgetRulesForAdvertiserResponse**](GetSDBudgetRulesForAdvertiserResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAssociatedBudgetRulesForSDCampaigns**
> SdListAssociatedBudgetRulesResponse ListAssociatedBudgetRulesForSDCampaigns(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, campaignId)
Gets a list of budget rules associated to a campaign specified by identifier.

**Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 
  **campaignId** | **float64**| The campaign identifier. | 

### Return type

[**SdListAssociatedBudgetRulesResponse**](SDListAssociatedBudgetRulesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBudgetRulesForSDCampaigns**
> UpdateBudgetRulesResponse UpdateBudgetRulesForSDCampaigns(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Update one or more budget rules.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSdBudgetRulesRequest**](UpdateSdBudgetRulesRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 

### Return type

[**UpdateBudgetRulesResponse**](UpdateBudgetRulesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

