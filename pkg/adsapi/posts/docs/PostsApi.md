# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePost**](PostsApi.md#CreatePost) | **Post** /bp/v2/posts | API to create a Post. 
[**GetPost**](PostsApi.md#GetPost) | **Get** /bp/v2/posts/{postId} | Get a post. 
[**GetPostProducts**](PostsApi.md#GetPostProducts) | **Get** /bp/v2/products/list | Get information for a list of products. 
[**GetProfile**](PostsApi.md#GetProfile) | **Get** /bp/v2/profiles/{profileId} | Get information for a Post Profile. 
[**GetProfileMetrics**](PostsApi.md#GetProfileMetrics) | **Post** /bp/v2/profiles/{profileId}/metrics | Get brand level performance metrics from Posts. 
[**GetReportDownloadLink**](PostsApi.md#GetReportDownloadLink) | **Get** /bp/v2/profiles/{profileId}/metrics/download | Get a URL to download a metrics report for the Post Profile. 
[**ListPosts**](PostsApi.md#ListPosts) | **Post** /bp/v2/posts/list | API to get posts data along with performance metrics for each post. 
[**ListProfiles**](PostsApi.md#ListProfiles) | **Get** /bp/v2/profiles | Get a list of Post Profiles that the advertiser has access to. 
[**SubmitPostForReview**](PostsApi.md#SubmitPostForReview) | **Put** /bp/v2/posts/{postId}/submitForReview | Submit a Post for review. 
[**UpdatePost**](PostsApi.md#UpdatePost) | **Put** /bp/v2/posts/{postId} | Update a Post&#x27;s data. 
[**WithdrawPost**](PostsApi.md#WithdrawPost) | **Put** /bp/v2/posts/{postId}/unpublish | Unpublishes a Post from Amazon. 

# **CreatePost**
> CreatePostResponseContent CreatePost(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
API to create a Post. 

API to create a Post.   **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CreatePostRequestContent**](CreatePostRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account.      Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with      the access token passed in the HTTP Authorization header. | 

### Return type

[**CreatePostResponseContent**](CreatePostResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.bpPost.v2+json
 - **Accept**: application/vnd.bpPost.v2+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPost**
> GetPostResponseContent GetPost(ctx, postId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Get a post. 

Get a post.   **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **postId** | **string**| Identifier for a post. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account.      Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with      the access token passed in the HTTP Authorization header. | 

### Return type

[**GetPostResponseContent**](GetPostResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.bpPost.v2+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetPostProducts**
> GetPostProductsResponseContent GetPostProducts(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Get information for a list of products. 

Get information for a list of products.   **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account.      Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with      the access token passed in the HTTP Authorization header. | 
 **optional** | ***PostsApiGetPostProductsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostsApiGetPostProductsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **asins** | [**optional.Interface of []string**](string.md)| A list of product identifiers. | 

### Return type

[**GetPostProductsResponseContent**](GetPostProductsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.bpProduct.v2+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProfile**
> GetProfileResponseContent GetProfile(ctx, profileId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Get information for a Post Profile. 

Get information for a Post Profile.   **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**| Identifier for a profile. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account.      Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with      the access token passed in the HTTP Authorization header. | 

### Return type

[**GetProfileResponseContent**](GetProfileResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.bpProfile.v2+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProfileMetrics**
> GetProfileMetricsResponseContent GetProfileMetrics(ctx, profileId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Get brand level performance metrics from Posts. 

Get brand level performance metrics from Posts.   **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account.      Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with      the access token passed in the HTTP Authorization header. | 
 **optional** | ***PostsApiGetProfileMetricsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostsApiGetProfileMetricsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of GetProfileMetricsRequestContent**](GetProfileMetricsRequestContent.md)|  | 

### Return type

[**GetProfileMetricsResponseContent**](GetProfileMetricsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.bpProfile.v2+json
 - **Accept**: application/vnd.bpProfile.v2+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetReportDownloadLink**
> GetReportDownloadLinkResponseContent GetReportDownloadLink(ctx, profileId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Get a URL to download a metrics report for the Post Profile. 

Get a URL to download a metrics report for the Post Profile.   **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **profileId** | **string**|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account.      Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with      the access token passed in the HTTP Authorization header. | 
 **optional** | ***PostsApiGetReportDownloadLinkOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostsApiGetReportDownloadLinkOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **metricStartDate** | **optional.String**| The start date to get metrics for. The value is in ISO8601 full-date format (UTC). For example, 2020-08-16. | 
 **metricEndDate** | **optional.String**| The end date to get metrics for. The value is in ISO8601 full-date format (UTC). For example, 2020-08-16. | 

### Return type

[**GetReportDownloadLinkResponseContent**](GetReportDownloadLinkResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.bpProfile.v2+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListPosts**
> ListPostsResponseContent ListPosts(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
API to get posts data along with performance metrics for each post. 

API to get posts data along with performance metrics for each post.   **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ListPostsRequestContent**](ListPostsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account.      Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with      the access token passed in the HTTP Authorization header. | 

### Return type

[**ListPostsResponseContent**](ListPostsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.bpPost.v2+json
 - **Accept**: application/vnd.bpPost.v2+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListProfiles**
> ListProfilesResponseContent ListProfiles(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Get a list of Post Profiles that the advertiser has access to. 

Get a list of Post Profiles that the advertiser has access to.   **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account.      Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with      the access token passed in the HTTP Authorization header. | 
 **optional** | ***PostsApiListProfilesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PostsApiListProfilesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nextToken** | **optional.String**| Pagination token to get next page of results from previous call. | 
 **maxResults** | **optional.Float64**| Number of items to be returned on this call. | 

### Return type

[**ListProfilesResponseContent**](ListProfilesResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.bpProfile.v2+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SubmitPostForReview**
> SubmitPostForReviewResponseContent SubmitPostForReview(ctx, body, postId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Submit a Post for review. 

Submit a Post for review.   **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SubmitPostForReviewRequestContent**](SubmitPostForReviewRequestContent.md)|  | 
  **postId** | **string**| Identifier for a post. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account.      Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with      the access token passed in the HTTP Authorization header. | 

### Return type

[**SubmitPostForReviewResponseContent**](SubmitPostForReviewResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.bpPost.v2+json
 - **Accept**: application/vnd.bpPost.v2+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePost**
> UpdatePostResponseContent UpdatePost(ctx, body, postId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Update a Post's data. 

Update a Post's data.   **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePostRequestContent**](UpdatePostRequestContent.md)|  | 
  **postId** | **string**| Identifier for a post. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account.      Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with      the access token passed in the HTTP Authorization header. | 

### Return type

[**UpdatePostResponseContent**](UpdatePostResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.bpPost.v2+json
 - **Accept**: application/vnd.bpPost.v2+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **WithdrawPost**
> WithdrawPostResponseContent WithdrawPost(ctx, body, postId, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Unpublishes a Post from Amazon. 

Unpublishes a Post from Amazon.   **Requires one of these permissions**: []

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WithdrawPostRequestContent**](WithdrawPostRequestContent.md)|  | 
  **postId** | **string**| Identifier for a post. | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account.      Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with      the access token passed in the HTTP Authorization header. | 

### Return type

[**WithdrawPostResponseContent**](WithdrawPostResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.bpPost.v2+json
 - **Accept**: application/vnd.bpPost.v2+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

