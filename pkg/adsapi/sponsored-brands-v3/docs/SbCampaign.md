# SbCampaign

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int32** | The campaign identifier. | [optional] [default to null]
**Name** | **string** | The name of the campaign. This name must be unique to the Amazon Ads account to which the campaign is associated. Maximum length of the string is 128 characters. | [optional] [default to null]
**Tags** | [***map[string]string**](map.md) |  | [optional] [default to null]
**Budget** | **float64** | The budget amount associated with the campaign. | [optional] [default to null]
**BudgetType** | [***BudgetType**](BudgetType.md) |  | [optional] [default to null]
**StartDate** | **string** | The YYYYMMDD start date for the campaign. If this field is not set to a value, the current date is used. | [optional] [default to null]
**EndDate** | **string** | The YYYYMMDD end date for the campaign. Must be greater than the value for &#x60;startDate&#x60;. If not specified, the campaign has no end date and runs continuously. | [optional] [default to null]
**State** | [***State**](State.md) |  | [optional] [default to null]
**ServingStatus** | **string** | |Value|Description| |-----|-----------| |asinNotBuyable| The associated ASIN cannot be purchased due to eligibility or availability.| |billingError| The billing information associated with the account requires correction.| |ended| The value specified in the &#x60;endDate&#x60; field is in the past.| |landingPageNotAvailable| The specified landing page is not available. This may be caused by an incorrect address or a landing page with less than three ASINs.| |outOfBudget| The campaign has run out of budget.| |paused|The campaign state is set to &#x60;paused&#x60;.| |pendingReview| A newly created campaign that has not passed moderation review. Note that moderation review may take up to 72 hours. |ready| The campaign is scheduled for a future date.| |rejected| The campaign failed moderation review.| |running| The campaign is enabled and serving.| |scheduled| A transitive state between &#x60;ready&#x60; and &#x60;running&#x60;, as child entities associated with the campaign move to a running state.| |terminated| The state of the campaign is set to &#x60;archived&#x60;.| &lt;br/&gt; | [optional] [default to null]
**BrandEntityId** | **string** | The brand entity identifier. Note that this field is required for sellers. For more information, see the [Stores reference](https://advertising.amazon.com/API/docs/v2/reference/stores) or [Brands reference](https://advertising.amazon.com/API/docs/v3/reference/SponsoredBrands/Brands). | [optional] [default to null]
**BidOptimization** | **bool** | Set to &#x60;true&#x60; to allow Amazon to automatically optimize bids for placements below top of search. | [optional] [default to true]
**BidMultiplier** | **float64** | A bid multiplier. Note that this field can only be set when &#x27;bidOptimization&#x27; is set to false. Value is a percentage to two decimal places. Example: If set to -40.00 for a $5.00 bid, the resulting bid is $3.00. | [optional] [default to null]
**BidAdjustments** | [**[]BidAdjustment**](BidAdjustment.md) | List of bid adjustment for each placement group. BidMultiplier cannot be specified when bidAdjustments presents. &#x60;Not supported for video campaigns&#x60; | [optional] [default to null]
**PortfolioId** | **int64** | The identifier of the portfolio to which the campaign is associated. | [optional] [default to null]
**Creative** | [***SbCreative**](SBCreative.md) |  | [optional] [default to null]
**LandingPage** | [***SbLandingPage**](SBLandingPage.md) |  | [optional] [default to null]
**Keywords** | [***[]SbCreateCampaignPositiveKeywordInner**](array.md) |  | [optional] [default to null]
**NegativeKeywords** | [***[]SbCreateCampaignNegativeKeywordInner**](array.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

