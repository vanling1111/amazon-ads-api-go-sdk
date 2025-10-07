# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveTarget**](ProductTargetingApi.md#ArchiveTarget) | **Delete** /sb/targets/{targetId} | Archives a target specified by identifier. Note that archiving is permanent, and once a target has been archived it can&#x27;t be made active again.
[**CreateTargets**](ProductTargetingApi.md#CreateTargets) | **Post** /sb/targets | Create one or more targets.
[**GetTarget**](ProductTargetingApi.md#GetTarget) | **Get** /sb/targets/{targetId} | Gets a target specified by identifier.
[**ListTargets**](ProductTargetingApi.md#ListTargets) | **Post** /sb/targets/list | 
[**UpdateTargets**](ProductTargetingApi.md#UpdateTargets) | **Put** /sb/targets | Updates one or more targets.

# **ArchiveTarget**
> SbTargetingClauseResponse ArchiveTarget(ctx, targetId)
Archives a target specified by identifier. Note that archiving is permanent, and once a target has been archived it can't be made active again.

The identifier of an existing target.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **targetId** | [**int64**](.md)|  | 

### Return type

[**SbTargetingClauseResponse**](SBTargetingClauseResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sbtargetresponse.v3+json, application/vnd.sberror.v3.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateTargets**
> SbCreateTargetsResponse CreateTargets(ctx, optional)
Create one or more targets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductTargetingApiCreateTargetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductTargetingApiCreateTargetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SbTargetsBody1**](SbTargetsBody1.md)| A list of targeting clauses for creation. &lt;br/&gt;Note that targets can be created on campaigns where serving status is not one of &#x60;archived&#x60;, &#x60;terminated&#x60;, &#x60;rejected&#x60;, or &#x60;ended&#x60;. &lt;br/&gt;Note that this operation supports a maximum list size of 100 targets. &lt;br/&gt;Note the bid is only mutable when the keyword&#x27;s corresponding campaign does not have any enabled optimization rule. | 

### Return type

[**SbCreateTargetsResponse**](SBCreateTargetsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sbcreatetargetsresponse.v3+json, application/vnd.sberror.v3.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTarget**
> SbTargetingClause GetTarget(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, targetId)
Gets a target specified by identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a **Login with Amazon** account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **targetId** | [**int64**](.md)| The identifier of an existing target. | 

### Return type

[**SbTargetingClause**](SBTargetingClause.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sbtarget.v3+json, application/vnd.sberror.v3.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListTargets**
> InlineResponse2008 ListTargets(ctx, optional)


Gets a list of product targets associated with the client identifier passed in the authorization header, filtered by specified criteria.  **Note**: Product targets associated with BrandVideo ad groups are only available in v3.2 version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductTargetingApiListTargetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductTargetingApiListTargetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TargetsListBody**](TargetsListBody.md)| A set of filters. | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sblisttargetsresponse.v3.2+json, application/vnd.sblisttargetsresponse.v3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTargets**
> InlineResponse2009 UpdateTargets(ctx, optional)
Updates one or more targets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProductTargetingApiUpdateTargetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProductTargetingApiUpdateTargetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SbTargetsBody**](SbTargetsBody.md)| A list of targets with updated values. &lt;br/&gt;Note that targets can be updated on campaigns where serving status is not one of &#x60;archived&#x60;, &#x60;terminated&#x60;, &#x60;rejected&#x60;, or &#x60;ended&#x60;. &lt;br/&gt;Note that this operation supports a maximum list size of 100 targets. &lt;br/&gt;Note the bid is only mutable when the keyword&#x27;s corresponding campaign does not have any enabled optimization rule. | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.updatetargetsresponse.v3+json, application/vnd.sberror.v3.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

