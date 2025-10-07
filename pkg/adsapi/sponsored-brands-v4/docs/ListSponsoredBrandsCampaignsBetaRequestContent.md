# ListSponsoredBrandsCampaignsBetaRequestContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignIdFilter** | [***ObjectIdFilter**](ObjectIdFilter.md) |  | [optional] [default to null]
**PortfolioIdFilter** | [***ObjectIdFilter**](ObjectIdFilter.md) |  | [optional] [default to null]
**StateFilter** | [***EntityStateFilter**](EntityStateFilter.md) |  | [optional] [default to null]
**MaxResults** | **float64** | Number of records to include in the paginated response. Defaults to max page size for given API. | [optional] [default to null]
**NextToken** | **string** | Token value allowing to navigate to the next response page. | [optional] [default to null]
**GoalTypeFilter** | [***GoalTypeFilter**](GoalTypeFilter.md) |  | [optional] [default to null]
**IncludeExtendedDataFields** | **bool** | Setting to true will slow down performance because the API needs to retrieve extra information for each campaign. | [optional] [default to null]
**NameFilter** | [***NameFilter**](NameFilter.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

