# ListAdsRequestContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertisementPurpose** | **string** | This parameter will limit results to those matching the requested advertisement purpose. This includes the name of the product, service, or brand within the ad or affiliate marketing content. | [optional] [default to null]
**AdvertiserName** | **string** | The person or entity on whose behalf the ad or affiliate marketing content is presented. This parameter will limit results to advertisement summaries matching the requested advertiser name. | [optional] [default to null]
**DeliveryAfterDateUtc** | [**time.Time**](time.Time.md) | When specified, limits results to those with delivery date after the specified date string. The date string is specified in ISO format (YYYY-MM-DD) in UTC timezone. For example: 2020-12-16. | [optional] [default to null]
**DeliveryBeforeDateUtc** | [**time.Time**](time.Time.md) | When specified, limits results to those with a delivery date prior to the specified date string. The date string is specified in ISO format (YYYY-MM-DD) in UTC timezone. For example: 2020-12-16. | [optional] [default to null]
**IsRestricted** | **bool** | Whether the advertisement was paused, removed or suppressed based on alleged illegality or incompatibility with terms and conditions. Limits results based on their restriction status. | [optional] [default to null]
**MaxResults** | **float64** | Defaults to 10, with supported values between 1 and 1000.     If the size of the result set is larger than the limit, callers will retrieve additional results through pagination parameters. | [optional] [default to null]
**NameMatchType** | [***NameMatchType**](NameMatchType.md) |  | [optional] [default to null]
**NextToken** | **string** | Pagination token to retrieve the next set of results.     Callers must make use of the token provided in a previous response to use this parameter. | [optional] [default to null]
**SiteName** | **string** | When specified, limits results based on their delivery to a particular Amazon EU store (e.g., amazon.de, amazon.fr). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

