# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DownloadSnapshot**](SnapshotsApi.md#DownloadSnapshot) | **Get** /v2/sp/snapshots/{snapshotId}/download | Download requested snapshot
[**GetSnapshotStatus**](SnapshotsApi.md#GetSnapshotStatus) | **Get** /v2/sp/snapshots/{snapshotId} | Get the status of a requested snapshot
[**RequestSnapshot**](SnapshotsApi.md#RequestSnapshot) | **Post** /v2/sp/{recordType}/snapshot | Request a snapshot

# **DownloadSnapshot**
> DownloadSnapshot(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, snapshotId)
Download requested snapshot

**Note: Snapshots APIs are deprecated and will be shut off on October 15, 2024. For replacement functionality, see the [exports](guides/exports/overview) API. To learn more, view the [migration guide](reference/migration-guides/snapshots-exports).**   To understand the call flow for asynchronous snapshots, see [Getting started with sponsored ads snapshots](/API/docs/en-us/guides/snapshots/get-started).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; developer account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **snapshotId** | **float64**| The snapshot identifier. | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshotStatus**
> SnapshotResponse GetSnapshotStatus(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, snapshotId)
Get the status of a requested snapshot

**Note: Snapshots APIs are deprecated and will be shut off on October 15, 2024. For replacement functionality, see the [exports](guides/exports/overview) API. To learn more, view the [migration guide](reference/migration-guides/snapshots-exports).**   To understand the call flow for asynchronous snapshots, see [Getting started with sponsored ads snapshots](/API/docs/en-us/guides/snapshots/overview).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; developer account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **snapshotId** | **float64**| The snapshot identifier. | 

### Return type

[**SnapshotResponse**](SnapshotResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RequestSnapshot**
> SnapshotResponse RequestSnapshot(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, recordType)
Request a snapshot

**Note: Snapshots APIs are deprecated and will be shut off on October 15, 2024. For replacement functionality, see the [exports](guides/exports/overview) API. To learn more, view the [migration guide](reference/migration-guides/snapshots-exports).**   Request a file-based snapshot of all entities of the specified type in the account satisfying the filtering criteria.  To understand the call flow for asynchronous snapshots, see [Getting started with sponsored ads snapshots](/API/docs/en-us/guides/snapshots/get-started).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SnapshotRequest**](SnapshotRequest.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; developer account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **recordType** | **string**| The type of entity for which the snapshot is generated. | 

### Return type

[**SnapshotResponse**](SnapshotResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

