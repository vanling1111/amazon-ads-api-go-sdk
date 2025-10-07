# SbVideoListCampaignItem

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LandingPage** | [***SbDetailPageLandingPage**](SBDetailPageLandingPage.md) |  | [optional] [default to null]
**Creative** | [***SbVideoCreative**](SBVideoCreative.md) |  | [optional] [default to null]
**SupplySource** | [***SupplySource**](SupplySource.md) |  | [optional] [default to null]
**CampaignId** | **int64** | The campaign identifier. | [optional] [default to null]
**Name** | **string** |  | [optional] [default to null]
**Tags** | [***map[string]string**](map.md) |  | [optional] [default to null]
**Budget** | **float64** |  | [optional] [default to null]
**BudgetType** | [***BudgetType**](BudgetType.md) |  | [optional] [default to null]
**StartDate** | **string** |  | [optional] [default to null]
**EndDate** | **string** |  | [optional] [default to null]
**State** | [***Object**](.md) |  | [optional] [default to null]
**ServingStatus** | **string** | |Status|Description| |------|-----------| |ASIN_NOT_BUYABLE| The ASIN can&#x27;t be purchased due to eligibility or availability.| |BILLING_ERROR| Billing information requires correction.| |ENDED| THe &#x60;endDate&#x60; specified in the campaign object occurs in the past.| |LANDING_PAGE_NOT_AVAILABLE| The specified landing page is not available. This may be caused by an incorrect address or a landing page with less than three ASINs.| |OUT_OF_BUDGET| The campaign has run out of budget.| |PAUSED| The campaign state set to &#x60;paused&#x60;.| |PENDING_REVIEW|: A newly created campaign that has not passed moderation review. Note that moderation review may take up to 72 hours.| |READY| The campaign is scheduled for a future date.| |REJECTED| The campaign failed moderation review.| |RUNNING| The campaign is enabled and serving.| |SCHEDULED| A transitive state between &#x60;ready&#x60; and &#x60;running&#x60;, as child entities associated with the campaign move to a running state.| |TERMINATED|The state of the campaign is set to &#x60;archived&#x60;.| | [optional] [default to null]
**BrandEntityId** | **string** | The brand entity identifier. Note that this field is required for sellers. For more information, see the [Stores reference](https://advertising.amazon.com/API/docs/v2/reference/stores) or [Brands reference](https://advertising.amazon.com/API/docs/v3/reference/SponsoredBrands/Brands). | [optional] [default to null]
**PortfolioId** | **int64** | The identifier of the portfolio to which the campaign is associated. | [optional] [default to null]
**AdFormat** | [***AdFormat**](AdFormat.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

