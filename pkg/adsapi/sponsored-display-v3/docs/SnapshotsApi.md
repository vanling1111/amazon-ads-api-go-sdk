# {{classname}}

All URIs are relative to *https://advertising-api.amazon.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshot**](SnapshotsApi.md#CreateSnapshot) | **Post** /sd/{recordType}/snapshot | Request a file-based snapshot of all entities of the specified type in the account satisfying the filtering criteria
[**DownloadSnapshot**](SnapshotsApi.md#DownloadSnapshot) | **Get** /sd/snapshots/{snapshotId}/download | Download previously requested snapshot
[**GetSnapshot**](SnapshotsApi.md#GetSnapshot) | **Get** /sd/snapshots/{snapshotId} | Retrieve status, metadata, and location of previously requested snapshot

# **CreateSnapshot**
> SnapshotResponse CreateSnapshot(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, recordType, optional)
Request a file-based snapshot of all entities of the specified type in the account satisfying the filtering criteria

**Note: Snapshots APIs are deprecated and will be shut off on October 15, 2024. For replacement functionality, see the [exports](guides/exports/overview) API. To learn more, view the [migration guide](reference/migration-guides/snapshots-exports).**  To understand the call flow for asynchronous snapshots, see [Getting started with sponsored ads snapshots](/API/docs/en-us/guides/snapshots/get-started).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **recordType** | **string**| **Note: Snapshots APIs are deprecated and will be shut off on October 15, 2024. For replacement functionality, see the [exports](guides/exports/overview) API. To learn more, view the [migration guide](reference/migration-guides/snapshots-exports).**  The type of entity for which the snapshot should be generated. Must be one of: &#x60;campaigns&#x60;, &#x60;adgroups&#x60;, &#x60;productAds&#x60;, or &#x60;targets&#x60;. | 
 **optional** | ***SnapshotsApiCreateSnapshotOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SnapshotsApiCreateSnapshotOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of SnapshotRequest**](SnapshotRequest.md)| Request a snapshot file for all entities of a single record type. | 

### Return type

[**SnapshotResponse**](SnapshotResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DownloadSnapshot**
> DownloadSnapshot(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, snapshotId)
Download previously requested snapshot

**Note: Snapshots APIs are deprecated and will be shut off on October 15, 2024. For replacement functionality, see the [exports](guides/exports/overview) API. To learn more, view the [migration guide](reference/migration-guides/snapshots-exports).**   To understand the call flow for asynchronous snapshots, see [Getting started with sponsored ads snapshots](/API/docs/en-us/guides/snapshots/overview).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **snapshotId** | **string**| The Snapshot identifier. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSnapshot**
> SnapshotResponse GetSnapshot(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, snapshotId)
Retrieve status, metadata, and location of previously requested snapshot

**Note: Snapshots APIs are deprecated and will be shut off on October 15, 2024. For replacement functionality, see the [exports](guides/exports/overview) API. To learn more, view the [migration guide](reference/migration-guides/snapshots-exports).**   To understand the call flow for asynchronous snapshots, see [Getting started with sponsored ads snapshots](/API/docs/en-us/guides/snapshots/get-started).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header. | 
  **snapshotId** | **string**| The Snapshot identifier. | 

### Return type

[**SnapshotResponse**](SnapshotResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth), [oauth2AuthorizationCode](../README.md#oauth2AuthorizationCode)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

