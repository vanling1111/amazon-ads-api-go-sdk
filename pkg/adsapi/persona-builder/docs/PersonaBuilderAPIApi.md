# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BandedSize**](PersonaBuilderAPIApi.md#BandedSize) | **Post** /insights/bandedSize | 
[**Demographics**](PersonaBuilderAPIApi.md#Demographics) | **Post** /insights/demographics | 
[**PrimeVideo**](PersonaBuilderAPIApi.md#PrimeVideo) | **Post** /insights/primeVideo | 
[**TopCategoriesPurchased**](PersonaBuilderAPIApi.md#TopCategoriesPurchased) | **Post** /insights/topCategoriesPurchased | 
[**TopOverlappingAudiences**](PersonaBuilderAPIApi.md#TopOverlappingAudiences) | **Post** /insights/topOverlappingAudiences | 

# **BandedSize**
> BandedSizeResponseContent BandedSize(ctx, advertiserId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Get banded size of number of unique customers that are in the input expression.  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **advertiserId** | **string**| The identifier of the advertiser, retrieved from /dsp/advertisers, that you&#x27;d like to retrieve insights for. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***PersonaBuilderAPIApiBandedSizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonaBuilderAPIApiBandedSizeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of InputExpression**](InputExpression.md)|  | 

### Return type

[**BandedSizeResponseContent**](BandedSizeResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.bandedsizeinputexpression.v1+json
 - **Accept**: application/vnd.bandedsizeinsights.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Demographics**
> DemographicsResponseContent Demographics(ctx, advertiserId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Get demographic insights for the input expression.  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **advertiserId** | **string**| The identifier of the advertiser, retrieved from /dsp/advertisers, that you&#x27;d like to retrieve insights for. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***PersonaBuilderAPIApiDemographicsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonaBuilderAPIApiDemographicsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of InputExpression**](InputExpression.md)|  | 

### Return type

[**DemographicsResponseContent**](DemographicsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.demographicinputexpressions.v1+json
 - **Accept**: application/vnd.demographicinsights.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PrimeVideo**
> PrimeVideoResponseContent PrimeVideo(ctx, advertiserId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Get Prime Video insights for the input expression.  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **advertiserId** | **string**| The identifier of the advertiser, retrieved from /dsp/advertisers, that you&#x27;d like to retrieve insights for. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***PersonaBuilderAPIApiPrimeVideoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonaBuilderAPIApiPrimeVideoOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of PrimeVideoInputExpression**](PrimeVideoInputExpression.md)|  | 

### Return type

[**PrimeVideoResponseContent**](PrimeVideoResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.primevideoinputexpressions.v1+json
 - **Accept**: application/vnd.primevideoinsights.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TopCategoriesPurchased**
> TopCategoriesPurchasedResponseContent TopCategoriesPurchased(ctx, advertiserId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Get insights on top retail categories purchased by customers in the input expression.  **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **advertiserId** | **string**| The identifier of the advertiser, retrieved from /dsp/advertisers, that you&#x27;d like to retrieve insights for. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***PersonaBuilderAPIApiTopCategoriesPurchasedOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonaBuilderAPIApiTopCategoriesPurchasedOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of InputExpression**](InputExpression.md)|  | 
 **maxResults** | **optional.**| Sets the maximum number of objects in the returned array. Use in conjunction with the nextToken parameter to control pagination. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned.&lt;br/&gt;Default: 30; &lt;br/&gt;Minimum: 1; &lt;br/&gt;Maximum:250. | 
 **nextToken** | **optional.**| Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the nextToken field is empty, there are no further results. | 

### Return type

[**TopCategoriesPurchasedResponseContent**](TopCategoriesPurchasedResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.topcategoriespurchasedinputexpression.v1+json
 - **Accept**: application/vnd.topcategoriespurchasedinsights.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TopOverlappingAudiences**
> TopOverlappingAudiencesResponseContent TopOverlappingAudiences(ctx, advertiserId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Get top audiences overlapping with the input expression.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Advertising-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **advertiserId** | **string**| The identifier of the advertiser, retrieved from /dsp/advertisers, that you&#x27;d like to retrieve insights for. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***PersonaBuilderAPIApiTopOverlappingAudiencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PersonaBuilderAPIApiTopOverlappingAudiencesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of TopOverlappingAudiencesInputExpression**](TopOverlappingAudiencesInputExpression.md)|  | 
 **maxResults** | **optional.**| Sets the maximum number of objects in the returned array. Use in conjunction with the nextToken parameter to control pagination. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned.&lt;br/&gt;Default: 30; &lt;br/&gt;Minimum: 1; &lt;br/&gt;Maximum:250. | 
 **nextToken** | **optional.**| Operations that return paginated results include a pagination token in this field. To retrieve the next page of results, call the same operation and specify this token in the request. If the nextToken field is empty, there are no further results. | 

### Return type

[**TopOverlappingAudiencesResponseContent**](TopOverlappingAudiencesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.topoverlappingaudiencesinputexpression.v1+json
 - **Accept**: application/vnd.topoverlappingaudiencesinsights.v1+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

