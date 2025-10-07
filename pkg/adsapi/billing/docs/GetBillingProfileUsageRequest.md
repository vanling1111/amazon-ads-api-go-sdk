# GetBillingProfileUsageRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpandBillingProfile** | **bool** | By default only the billingProfileId linked to the country will be returned. Choose &#x60;true&#x60; if you would like to see the content of the billing profiles linked to a country | [optional] [default to false]
**ExpandFallbackBillingProfile** | **bool** | Choose &#x60;true&#x60; if you would like to see the information currently being used for billing for this marketplace. Useful when the billingProfileUsage status is not &#x60;OK&#x60; and the billing information coundn&#x27;t be propagated for that country. If set to &#x60;true&#x60;, &#x60;fallbackBillingProfiles&#x60; attribute will be returned. | [optional] [default to false]
**Filters** | [***GetBillingProfileUsageRequestFilters**](GetBillingProfileUsageRequest_filters.md) |  | [optional] [default to null]
**MaxResults** | **int64** | Max results / billing profile usage(s) to be shown in a single page. | [optional] [default to null]
**NextToken** | **string** | Offset to fetch next page with list of billing profile usage(s). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

