# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SBTargetingGetRefinementsForCategory**](ProductTargetingCategoriesApi.md#SBTargetingGetRefinementsForCategory) | **Get** /sb/targets/categories/{categoryRefinementId}/refinements | Get refinements for category
[**SBTargetingGetTargetableASINCounts**](ProductTargetingCategoriesApi.md#SBTargetingGetTargetableASINCounts) | **Post** /sb/targets/products/count | Get number of products in a category
[**SBTargetingGetTargetableCategories**](ProductTargetingCategoriesApi.md#SBTargetingGetTargetableCategories) | **Get** /sb/targets/categories | Get targetable categories

# **SBTargetingGetRefinementsForCategory**
> SbTargetingGetRefinementsForCategoryResponseContent SBTargetingGetRefinementsForCategory(ctx, categoryRefinementId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Get refinements for category

Returns refinements according to category input.   Only available in the following marketplaces: US, CA, MX, UK, DE, FR, ES, IT, IN, JP  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **categoryRefinementId** | **string**| The category refinement id. Please use /sb/targets/categories or /sb/recommendations/targets/category to retrieve category IDs. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***ProductTargetingCategoriesApiSBTargetingGetRefinementsForCategoryOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductTargetingCategoriesApiSBTargetingGetRefinementsForCategoryOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **locale** | [**optional.Interface of SbTargetingLocale**](.md)| The locale to which the caller wishes to translate the targetable categories or refinements to. For example, if the caller wishes to receive the targetable categories in Simplified Chinese, the locale parameter should be set to zh_CN. If no locale is provided, the returned tagetable categories will be in the default language of the marketplace. | 
 **nextToken** | **optional.String**| Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;NextToken&#x60; field is empty, there are no further results. | 

### Return type

[**SbTargetingGetRefinementsForCategoryResponseContent**](SBTargetingGetRefinementsForCategoryResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sbtargeting.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SBTargetingGetTargetableASINCounts**
> SbTargetingGetTargetableAsinCountsResponseContent SBTargetingGetTargetableASINCounts(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Get number of products in a category

Get number of targetable asins based on refinements provided by the user.  Use `/sb/targets/categories` or `/sb/recommendations/targets/category` to retrieve the category ID. Use `/sb/targets/categories/{categoryRefinementId}/refinements` to retrieve refinements data for a category.  Only available in the following marketplaces: US, CA, MX, UK, DE, FR, ES, IT, IN, JP  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SbTargetingGetTargetableAsinCountsRequestContent**](SbTargetingGetTargetableAsinCountsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**SbTargetingGetTargetableAsinCountsResponseContent**](SBTargetingGetTargetableASINCountsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.sbtargeting.v4+json
 - **Accept**: application/vnd.sbtargeting.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SBTargetingGetTargetableCategories**
> SbTargetingGetTargetableCategoriesResponseContent SBTargetingGetTargetableCategories(ctx, supplySource, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Get targetable categories

Returns all targetable categories by default in a list. List of categories can be used to build and traverse category tree. Set query parameter `includeOnlyRootCategories=true` to return only the root categories, or set `parentCategoryRefinementId` to return children of a specific parent category. Each category node has the fields - category name, category refinement id, parent category refinement id, isTargetable flag, and ASIN count range.   Only available in the following marketplaces: US, CA, MX, UK, DE, FR, ES, IT, IN, JP  **Requires one of these permissions**: [\"advertiser_campaign_edit\",\"advertiser_campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **supplySource** | [**SbTargetingSupplySource**](.md)| [UPDATE: As of 05/28/2024, &#x60;STREAMING_VIDEO&#x60; is deprecated].   The supply source where the target will be used. Use &#x60;AMAZON&#x60; for placements on Amazon website. Use &#x60;STREAMING_VIDEO&#x60; for off-site video placements such as IMDb TV. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a &#x60;Login with Amazon&#x60; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***ProductTargetingCategoriesApiSBTargetingGetTargetableCategoriesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductTargetingCategoriesApiSBTargetingGetTargetableCategoriesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **locale** | [**optional.Interface of SbTargetingLocale**](.md)| The locale to which the caller wishes to translate the targetable categories or refinements to. For example, if the caller wishes to receive the targetable categories in Simplified Chinese, the locale parameter should be set to zh_CN. If no locale is provided, the returned tagetable categories will be in the default language of the marketplace. | 
 **includeOnlyRootCategories** | **optional.Bool**| Indicates whether to only retun root categories or not. | 
 **parentCategoryRefinementId** | **optional.String**| Returns child categories of category. | 
 **nextToken** | **optional.String**| Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the &#x60;NextToken&#x60; field is empty, there are no further results. | 

### Return type

[**SbTargetingGetTargetableCategoriesResponseContent**](SBTargetingGetTargetableCategoriesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sbtargeting.v4+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

