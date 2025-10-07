# AssetBasedCreativeSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalHtml** | **string** | Additional HTML to include with the render response for display inventory targets. | [optional] [default to null]
**Backgrounds** | [**[]Background**](Background.md) | The background which is displayed on the ad. | [optional] [default to null]
**BodyText** | **[]string** | The body text to use for the Asset Based Creative experience. | [optional] [default to null]
**Brand** | **string** | The brand of the product(s) being advertised. | [optional] [default to null]
**CallToActions** | [***AssetBasedCreativeCallToAction**](AssetBasedCreativeCallToAction.md) |  | [optional] [default to null]
**ClickTrackingUrls** | [**[]CreativeTrackingUrl**](CreativeTrackingUrl.md) | The third party urls to trigger when an click is recorded. | [optional] [default to null]
**CreativeSizes** | [**[]Size**](Size.md) | The placement sizes this creative should serve on. | [optional] [default to null]
**CustomVideos** | [***Video**](Video.md) |  | [optional] [default to null]
**Disclaimers** | **string** | The disclaimers to use for the Asset Based Creative experience. | [optional] [default to null]
**EnableCreativeAutoTranslation** | **bool** | If set to true and the headline and/or video are not in the marketplace&#x27;s default language, Amazon will attempt to translate them to the marketplace&#x27;s default language. If Amazon is unable to translate them, the ad will be rejected by moderation. | [optional] [default to null]
**HasTermsAndConditions** | **bool** | Indicates that the ad promotes a free product or service and has qualifying terms and conditions applicable to the customer. LandingPageURL must link out to a page detailing terms and conditions or contain a link to those. | [optional] [default to null]
**Headlines** | **[]string** | The headline(s) to use for the Asset Based Creative experience. | [optional] [default to null]
**Images** | [**[]Image**](Image.md) | The image(s) to use. | [optional] [default to null]
**ImpressionTrackingUrls** | [**[]CreativeTrackingUrl**](CreativeTrackingUrl.md) | The third party urls to trigger when an impression is recorded. | [optional] [default to null]
**InventoryTypes** | [**[]ComponentInventoryType**](ComponentInventoryType.md) | The inventory types this creative should serve on. | [optional] [default to null]
**LandingPage** | [***ComponentLandingPage**](ComponentLandingPage.md) |  | [optional] [default to null]
**Language** | [***LanguageLocale**](LanguageLocale.md) |  | [optional] [default to null]
**Logos** | [**[]Image**](Image.md) | The logos to use for the Asset Based Creative experience. | [optional] [default to null]
**ModerationStatus** | [***CreativeStatus**](CreativeStatus.md) |  | [optional] [default to null]
**OptimizationGoalKpi** | [***CreativeOptimizationGoalKpi**](CreativeOptimizationGoalKpi.md) |  | [optional] [default to null]
**ResponsiveSizingBehavior** | [***ResponsiveSizingBehavior**](ResponsiveSizingBehavior.md) |  | [optional] [default to null]
**SquareImages** | [**[]Image**](Image.md) | The square image(s) to use. | [optional] [default to null]
**TallImages** | [**[]Image**](Image.md) | The tall image(s) to use. | [optional] [default to null]
**UntranslatedHeadlines** | **[]string** | The headline entered by the advertiser. | [optional] [default to null]
**WideImages** | [**[]Image**](Image.md) | The wide image(s) to use. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

