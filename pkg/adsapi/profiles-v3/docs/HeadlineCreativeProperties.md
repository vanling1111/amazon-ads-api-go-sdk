# HeadlineCreativeProperties

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headline** | **string** | A marketing phrase to display on the ad. This field is optional and mutable. Maximum number of characters allowed is 50. | [optional] [default to null]
**HasTermsAndConditions** | **bool** | Indicates that the ad promotes a free product or service (e.g., &#x27;buy one get one free&#x27; or &#x27;free one-month trial&#x27;) and has qualifying terms and conditions applicable to your customer. Only supported for productAds with landingPageType of OFF_AMAZON_LINK. LandingPageURL must link out to a page detailing terms and conditions or contain a link to those. | [optional] [default to null]
**OriginalHeadline** | **string** | The original headline submitted by the advertiser. If &#x27;consentToTranslate&#x27; is set to true and translation is SUCCESSFUL then &#x60;headline&#x60; will return the translated headline whereas &#x60;originalHeadline&#x60; will return the original headline. In all other cases, &#x27;originalHeadline&#x27; and &#x60;headline&#x60; both will return the original headline. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

