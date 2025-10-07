# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SbCreateThemes**](ThemeTargetingApi.md#SbCreateThemes) | **Post** /sb/themes | Create one or more theme targets.
[**SbListThemes**](ThemeTargetingApi.md#SbListThemes) | **Post** /sb/themes/list | 
[**SbUpdateThemes**](ThemeTargetingApi.md#SbUpdateThemes) | **Put** /sb/themes | Updates one or more theme targets.

# **SbCreateThemes**
> SbCreateThemesResponse SbCreateThemes(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Create one or more theme targets.

Note that this endpoint does not support for **Author profiles**.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***ThemeTargetingApiSbCreateThemesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ThemeTargetingApiSbCreateThemesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SbThemesBody1**](SbThemesBody1.md)| A list of theme targets for creation. &lt;br/&gt;Note that theme targets can be created on multi-adGroups campaigns and where campaign serving status is not one of &#x60;archived&#x60;, &#x60;terminated&#x60;, &#x60;rejected&#x60;, or &#x60;ended&#x60; and adgroup state is not &#x27;archived&#x27;. &lt;br/&gt;Note that this operation supports a maximum list size of 100 theme targets and only one target can be created for each themeType per adGroup. &lt;br/&gt;Note the bid is only mutable when the keyword&#x27;s corresponding campaign does not have any enabled optimization rule. | 

### Return type

[**SbCreateThemesResponse**](SBCreateThemesResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sbthemescreateresponse.v3+json, application/vnd.sberror.v3.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SbListThemes**
> InlineResponse20012 SbListThemes(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)


Note that this endpoint does not support for **Author profiles**. Gets a list of theme targets associated with the client identifier passed in the authorization header, filtered by specified criteria.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***ThemeTargetingApiSbListThemesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ThemeTargetingApiSbListThemesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of ThemesListBody**](ThemesListBody.md)| A set of filters. | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sbthemeslistresponse.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SbUpdateThemes**
> InlineResponse20013 SbUpdateThemes(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Updates one or more theme targets.

Note that this endpoint does not support for **Author profiles**.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
 **optional** | ***ThemeTargetingApiSbUpdateThemesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ThemeTargetingApiSbUpdateThemesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of SbThemesBody**](SbThemesBody.md)| A list of theme targets with updated values. &lt;br/&gt;Note that theme targets can be updated on multi-adGroups campaigns and where campaign serving status is not one of &#x60;archived&#x60;, &#x60;terminated&#x60;, &#x60;rejected&#x60;, or &#x60;ended&#x60; and adgroup state is not &#x27;archived&#x27;. &lt;br/&gt;Note that this operation supports a maximum list size of 100 theme targets. Also theme target can not be archived. &lt;br/&gt;Note the bid is only mutable when the keyword&#x27;s corresponding campaign does not have any enabled optimization rule. | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sbthemesupdateresponse.v3+json, application/vnd.sberror.v3.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

