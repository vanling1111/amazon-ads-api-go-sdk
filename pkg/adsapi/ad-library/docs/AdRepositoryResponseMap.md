# AdRepositoryResponseMap

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertisementPurpose** | **string** | Displays the intent of the advertisement. | [default to null]
**AdvertiserName** | **string** | Name of the Advertiser. | [default to null]
**ContentUrls** | **[]string** | List of strings containing link to (1) the store product detail page of products shown in an ad, (2) the store brand page of the brand shown in an ad, or (3) the image of the ad. | [optional] [default to null]
**DeliveryAfterDateUtc** | [**time.Time**](time.Time.md) | The specified start date of the delivered ads.The date string is specified in ISO format (YYYY-MM-DD) in UTC timezone. For example: 2020-12-16. | [default to null]
**DeliveryBeforeDateUtc** | [**time.Time**](time.Time.md) | The specified end date of the delivered ads.The date string is specified in ISO format (YYYY-MM-DD) in UTC timezone. For example: 2020-12-16. | [default to null]
**Id** | **string** | Globally unique identifier for the advertisement, to support lookups in the repository.  It represents the combination of an ad-creative identifier, its marketplace and source of record. | [default to null]
**IllegalContentReport** | **bool** | Describes whether the ad was removed as a result of an illegal content report. | [optional] [default to null]
**IsRestricted** | **bool** | Whether the ad was paused, removed or suppressed based on alleged illegality or incompatibility with terms and conditions. | [default to null]
**PayerName** | **string** | Name of the entity who paid for the Ad. | [optional] [default to null]
**RestrictionCategory** | **string** | Reason(s) why an ad was paused, removed, or suppressed. | [optional] [default to null]
**RestrictionDetectionAutomated** | **bool** | Describes whether automation was used to identify the restricted ad. | [optional] [default to null]
**SubjectMatterUrl** | **string** | Subject matter of the advertisement (link to image or rendering). | [default to null]
**TargetingMethods** | [**[]TargetMethodRecipientGroup**](TargetMethodRecipientGroup.md) | This field lists the targeted advertising methods with the descriptions.     The targeted advertising methods includes:         1: &#x27;LOCATION&#x27;: &#x27;Their estimated location&#x27;,         2: &#x27;PAST_ACTIVITY&#x27;: &#x27;Their past activity on Amazon websites or services&#x27;,         3: &#x27;ADS_PREFERENCES&#x27;: &#x27;Their data sharing or personalized ads preferences&#x27;,         4: &#x27;SEARCH_TERMS&#x27;: &#x27;The search terms they used or the content or products they were viewing&#x27;, | [default to null]
**TotalRecipientsRange** | [***TotalRecipientsRange**](TotalRecipientsRange.md) |  | [default to null]
**TotalRecipientsRangeBySite** | [***RecipientsBySite**](RecipientsBySite.md) |  | [default to null]
**Type_** | [***ModelType**](Type.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

