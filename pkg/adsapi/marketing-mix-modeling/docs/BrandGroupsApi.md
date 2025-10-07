# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMmmBrandGroups**](BrandGroupsApi.md#ListMmmBrandGroups) | **Post** /mmm/v1/brandGroups/list | List brand groups

# **ListMmmBrandGroups**
> InlineResponse200 ListMmmBrandGroups(ctx, optional)
List brand groups

Lists the predefined brand groups for which reports may be requested. A `brandGroupId` must be provided in each request to create a report using `POST /mmm/v1/reports`. Brand groups are configured by an MMM program manager as part of the onboarding process. Contact <mmm-support@amazon.com> with any questions about the brand groups defined for your manager account.  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\"MasterAccount_Manager\",\"MasterAccount_Viewer\",\"ManagerAccount_Dev\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BrandGroupsApiListMmmBrandGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BrandGroupsApiListMmmBrandGroupsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of BrandGroupsListBody**](BrandGroupsListBody.md)|  | 

### Return type

[**InlineResponse200**](inline_response_200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

