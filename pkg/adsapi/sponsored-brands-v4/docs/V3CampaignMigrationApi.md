# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MigrationJobResults**](V3CampaignMigrationApi.md#MigrationJobResults) | **Post** /sb/v4/legacyCampaigns/migrationJob/results | List Migration Results of all Campaign.
[**MigrationJobStatus**](V3CampaignMigrationApi.md#MigrationJobStatus) | **Post** /sb/v4/legacyCampaigns/migrationJob/status | List Migration Job Status.
[**MigrationResults**](V3CampaignMigrationApi.md#MigrationResults) | **Post** /sb/v4/legacyCampaigns/overallMigrationResults | Lists all Campaign Migration results for an advertiser.
[**StartMigrationJob**](V3CampaignMigrationApi.md#StartMigrationJob) | **Post** /sb/v4/legacyCampaigns/migrationJob | Creates Migration Job for V3 campaigns.

# **MigrationJobResults**
> MigrationJobResultsResponseContent MigrationJobResults(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
List Migration Results of all Campaign.

List Migration Results of all Campaign.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MigrationJobResultsRequestContent**](MigrationJobResultsRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**MigrationJobResultsResponseContent**](MigrationJobResultsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.SponsoredBrands.SponsoredBrandsMigrationApi.v4+json
 - **Accept**: application/vnd.SponsoredBrands.SponsoredBrandsMigrationApi.v4+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrationJobStatus**
> MigrationJobStatusResponseContent MigrationJobStatus(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
List Migration Job Status.

List Migration Job Status.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MigrationJobStatusRequestContent**](MigrationJobStatusRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**MigrationJobStatusResponseContent**](MigrationJobStatusResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.SponsoredBrands.SponsoredBrandsMigrationApi.v4+json
 - **Accept**: application/vnd.SponsoredBrands.SponsoredBrandsMigrationApi.v4+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrationResults**
> MigrationResultsResponseContent MigrationResults(ctx, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope, optional)
Lists all Campaign Migration results for an advertiser.

Lists all Campaign Migration results for an advertiser

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **optional** | ***V3CampaignMigrationApiMigrationResultsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V3CampaignMigrationApiMigrationResultsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of MigrationResultsRequestContent**](MigrationResultsRequestContent.md)|  | 

### Return type

[**MigrationResultsResponseContent**](MigrationResultsResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.SponsoredBrands.SponsoredBrandsMigrationApi.v4+json
 - **Accept**: application/vnd.SponsoredBrands.SponsoredBrandsMigrationApi.v4+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartMigrationJob**
> StartMigrationJobResponseContent StartMigrationJob(ctx, body, amazonAdvertisingAPIClientId, amazonAdvertisingAPIScope)
Creates Migration Job for V3 campaigns.

Creates Migration Job for V3 campaigns.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**StartMigrationJobRequestContent**](StartMigrationJobRequestContent.md)|  | 
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **amazonAdvertisingAPIScope** | **string**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**StartMigrationJobResponseContent**](StartMigrationJobResponseContent.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.SponsoredBrands.SponsoredBrandsMigrationApi.v4+json
 - **Accept**: application/vnd.SponsoredBrands.SponsoredBrandsMigrationApi.v4+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

