# ProductVideoSettings

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | **string** | The name of the brand being advertised. | [optional] [default to null]
**BrandLogos** | [**[]Image**](Image.md) | The brand logo image assets to be used in the ad. | [optional] [default to null]
**EnableCreativeAutoTranslation** | **bool** | If set to true and the headline and/or video are not in the marketplace&#x27;s default language, Amazon will attempt to translate them to the marketplace&#x27;s default language. If Amazon is unable to translate them, the ad will be rejected by moderation. | [optional] [default to null]
**Headlines** | **[]string** | The headline submitted as part of the ad creative. During your campaign, Amazon will optimize amongst the headlines you provide to match customer intent. | [optional] [default to null]
**LandingPage** | [***VideoLandingPage**](VideoLandingPage.md) |  | [optional] [default to null]
**ModerationStatus** | [***CreativeStatus**](CreativeStatus.md) |  | [optional] [default to null]
**Products** | [**[]AdvertisedProducts**](AdvertisedProducts.md) | The products featured in the video ad. | [optional] [default to null]
**UntranslatedHeadlines** | **[]string** | The headline entered by the advertiser. | [optional] [default to null]
**UntranslatedVideos** | [**[]Video**](Video.md) | The original video assets submitted as part of the creative. | [optional] [default to null]
**Videos** | [**[]Video**](Video.md) | The video assets used in the ad. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

