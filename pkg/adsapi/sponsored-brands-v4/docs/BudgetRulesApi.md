# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssociatedBudgetRulesForSBCampaigns**](BudgetRulesApi.md#CreateAssociatedBudgetRulesForSBCampaigns) | **Post** /sb/campaigns/{campaignId}/budgetRules | Associate budget rules to campaign
[**CreateBudgetRulesForSBCampaigns**](BudgetRulesApi.md#CreateBudgetRulesForSBCampaigns) | **Post** /sb/budgetRules | Create budget rules
[**DisassociateAssociatedBudgetRuleForSBCampaigns**](BudgetRulesApi.md#DisassociateAssociatedBudgetRuleForSBCampaigns) | **Delete** /sb/campaigns/{campaignId}/budgetRules/{budgetRuleId} | Disassociate budget rule from campaign
[**GetBudgetRuleByRuleIdForSBCampaigns**](BudgetRulesApi.md#GetBudgetRuleByRuleIdForSBCampaigns) | **Get** /sb/budgetRules/{budgetRuleId} | Get budget rule by ID
[**GetCampaignsAssociatedWithSBBudgetRule**](BudgetRulesApi.md#GetCampaignsAssociatedWithSBBudgetRule) | **Get** /sb/budgetRules/{budgetRuleId}/campaigns | Get campaigns associated with budget rule
[**GetSBBudgetRulesForAdvertiser**](BudgetRulesApi.md#GetSBBudgetRulesForAdvertiser) | **Get** /sb/budgetRules | Get budget rules
[**ListAssociatedBudgetRulesForSBCampaigns**](BudgetRulesApi.md#ListAssociatedBudgetRulesForSBCampaigns) | **Get** /sb/campaigns/{campaignId}/budgetRules | Get budget rules associated to campaign
[**UpdateBudgetRulesForSBCampaigns**](BudgetRulesApi.md#UpdateBudgetRulesForSBCampaigns) | **Put** /sb/budgetRules | Update budget rules

# **CreateAssociatedBudgetRulesForSBCampaigns**
> CreateAssociatedBudgetRulesResponse CreateAssociatedBudgetRulesForSBCampaigns(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, campaignId)
Associate budget rules to campaign

A maximum of 250 rules can be associated to a campaign. Note that the name of each rule associated to a campaign is required to be unique.  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

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

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateBudgetRulesForSBCampaigns**
> CreateBudgetRulesResponse CreateBudgetRulesForSBCampaigns(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Create budget rules

  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreateSbBudgetRulesRequest**](CreateSbBudgetRulesRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 

### Return type

[**CreateBudgetRulesResponse**](CreateBudgetRulesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DisassociateAssociatedBudgetRuleForSBCampaigns**
> DisassociateAssociatedBudgetRuleResponse DisassociateAssociatedBudgetRuleForSBCampaigns(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, campaignId, budgetRuleId)
Disassociate budget rule from campaign

  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

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

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetBudgetRuleByRuleIdForSBCampaigns**
> GetSbBudgetRuleResponse GetBudgetRuleByRuleIdForSBCampaigns(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, budgetRuleId)
Get budget rule by ID

  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 
  **budgetRuleId** | **string**| The budget rule identifier. | 

### Return type

[**GetSbBudgetRuleResponse**](GetSBBudgetRuleResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetCampaignsAssociatedWithSBBudgetRule**
> SbGetAssociatedCampaignsResponse GetCampaignsAssociatedWithSBBudgetRule(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, budgetRuleId, pageSize, optional)
Get campaigns associated with budget rule

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 
  **budgetRuleId** | **string**| The budget rule identifier. | 
  **pageSize** | **float64**| Sets a limit on the number of results returned. Maximum limit of &#x60;pageSize&#x60; is 30. | 
 **optional** | ***BudgetRulesApiGetCampaignsAssociatedWithSBBudgetRuleOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BudgetRulesApiGetCampaignsAssociatedWithSBBudgetRuleOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **nextToken** | **optional.String**| To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;nextToken&#x60; field is empty, there are no further results. | 

### Return type

[**SbGetAssociatedCampaignsResponse**](SBGetAssociatedCampaignsResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSBBudgetRulesForAdvertiser**
> GetSbBudgetRulesForAdvertiserResponse GetSBBudgetRulesForAdvertiser(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, pageSize, optional)
Get budget rules

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 
  **pageSize** | **float64**| Sets a limit on the number of results returned. Maximum limit of &#x60;pageSize&#x60; is 30. | 
 **optional** | ***BudgetRulesApiGetSBBudgetRulesForAdvertiserOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BudgetRulesApiGetSBBudgetRulesForAdvertiserOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **nextToken** | **optional.String**| To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;nextToken&#x60; field is empty, there are no further results. | 

### Return type

[**GetSbBudgetRulesForAdvertiserResponse**](GetSBBudgetRulesForAdvertiserResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListAssociatedBudgetRulesForSBCampaigns**
> SbListAssociatedBudgetRulesResponse ListAssociatedBudgetRulesForSBCampaigns(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, campaignId)
Get budget rules associated to campaign

  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 
  **campaignId** | **float64**| The campaign identifier. | 

### Return type

[**SbListAssociatedBudgetRulesResponse**](SBListAssociatedBudgetRulesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateBudgetRulesForSBCampaigns**
> UpdateBudgetRulesResponse UpdateBudgetRulesForSBCampaigns(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Update budget rules

  **Requires one of these permissions**: [\"advertiser_campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateSbBudgetRulesRequest**](UpdateSbBudgetRulesRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. This is a required header for advertisers and integrators using the Advertising API. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. This is a required header for advertisers and integrators using the Advertising API. | 

### Return type

[**UpdateBudgetRulesResponse**](UpdateBudgetRulesResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

