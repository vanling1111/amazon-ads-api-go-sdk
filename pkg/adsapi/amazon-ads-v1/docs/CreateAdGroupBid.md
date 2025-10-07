# CreateAdGroupBid

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseBid** | **float64** | The lower bound bid used for the ads in the ad group. | [optional] [default to null]
**DefaultBid** | **float64** | The default maximum bid for ads and targets in the ad group. This is used in sponsored ads as the maximum bid during the auction. | [optional] [default to null]
**MarketplaceSettings** | [**[]CreateAdGroupBidMarketplaceSetting**](CreateAdGroupBidMarketplaceSetting.md) | The bid associated with the ad group at specified marketplace level. Either one of bid or marketplaceSettings should always be specified | [optional] [default to null]
**MaxAverageBid** | **float64** | The max average bid that will be targeted on the ad group across all of the bids (a single bid could be lower or higher that this number). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

