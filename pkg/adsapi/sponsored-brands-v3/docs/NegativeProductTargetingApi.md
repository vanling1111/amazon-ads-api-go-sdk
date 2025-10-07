# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveNegativeTarget**](NegativeProductTargetingApi.md#ArchiveNegativeTarget) | **Delete** /sb/negativeTargets/{negativeTargetId} | Archives a negative target specified by identifier. Note that archiving is permanent, and once a negative target has been archived it can&#x27;t be made active again.
[**CreateNegativeTargets**](NegativeProductTargetingApi.md#CreateNegativeTargets) | **Post** /sb/negativeTargets | Create one or more negative targets.
[**GetNegativeTarget**](NegativeProductTargetingApi.md#GetNegativeTarget) | **Get** /sb/negativeTargets/{negativeTargetId} | Gets a negative target specified by identifier.
[**ListNegativeTargets**](NegativeProductTargetingApi.md#ListNegativeTargets) | **Post** /sb/negativeTargets/list | Gets a list of product negative targets associated with the client identifier passed in the authorization header, filtered by specified criteria.
[**UpdateNegativeTargets**](NegativeProductTargetingApi.md#UpdateNegativeTargets) | **Put** /sb/negativeTargets | Updates one or more negative targets.

# **ArchiveNegativeTarget**
> SbTargetingClauseResponse ArchiveNegativeTarget(ctx, negativeTargetId)
Archives a negative target specified by identifier. Note that archiving is permanent, and once a negative target has been archived it can't be made active again.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **negativeTargetId** | [**int64**](.md)|  | 

### Return type

[**SbTargetingClauseResponse**](SBTargetingClauseResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sbnegativetarget.v3+json, application/vnd.sberror.v3.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateNegativeTargets**
> SbCreateTargetsResponse CreateNegativeTargets(ctx, optional)
Create one or more negative targets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NegativeProductTargetingApiCreateNegativeTargetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegativeProductTargetingApiCreateNegativeTargetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SbNegativeTargetsBody1**](SbNegativeTargetsBody1.md)| A list of negative targeting clauses for creation. &lt;br/&gt;Note that negative targeting clauses can be created on campaigns where serving status is not one of &#x60;archived&#x60;, &#x60;terminated&#x60;, &#x60;rejected&#x60;, or &#x60;ended&#x60;. &lt;br/&gt;Note that this operation supports a maximum list size of 100 negative targets. | 

### Return type

[**SbCreateTargetsResponse**](SBCreateTargetsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sbcreatenegativetargetsrequest.v3+json, application/vnd.sberror.v3.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetNegativeTarget**
> SbNegativeTargetingClause GetNegativeTarget(ctx, negativeTargetId)
Gets a negative target specified by identifier.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **negativeTargetId** | [**int64**](.md)| The identifier of an existing negative target. | 

### Return type

[**SbNegativeTargetingClause**](SBNegativeTargetingClause.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.sbnegativetarget.v3+json, application/vnd.sberror.v3.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListNegativeTargets**
> InlineResponse20010 ListNegativeTargets(ctx, optional)
Gets a list of product negative targets associated with the client identifier passed in the authorization header, filtered by specified criteria.

**Note**: Negative targets associated with BrandVideo ad groups are only available in v3.2 version.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NegativeProductTargetingApiListNegativeTargetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegativeProductTargetingApiListNegativeTargetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NegativeTargetsListBody**](NegativeTargetsListBody.md)| A set of filters. | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.sblistnegativetargetsresponse.v3.2+json, application/vnd.sblistnegativetargetsresponse.v3+json, application/vnd.sberror.v3.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateNegativeTargets**
> InlineResponse20011 UpdateNegativeTargets(ctx, optional)
Updates one or more negative targets.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NegativeProductTargetingApiUpdateNegativeTargetsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NegativeProductTargetingApiUpdateNegativeTargetsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SbNegativeTargetsBody**](SbNegativeTargetsBody.md)| A list of negative targets with updated values. &lt;br/&gt;Note that negative targeting clauses can be created on campaigns where serving status is not one of &#x60;archived&#x60;, &#x60;terminated&#x60;, &#x60;rejected&#x60;, or &#x60;ended&#x60;. &lt;br/&gt;Note that this operation supports a maximum list size of 100 negative targets. | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/vnd.updatenegativetargetsresponse.v3+json, application/vnd.sberror.v3.0+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

