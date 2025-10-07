# Account

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Id of the Amazon Advertising account. | [optional] [default to null]
**AccountName** | **string** | The name given to the Amazon Advertising account. | [optional] [default to null]
**AccountType** | [***AccountType**](AccountType.md) |  | [optional] [default to null]
**DspAdvertiserId** | **string** | The identifier of a DSP advertiser. Note that this value is only populated for accounts with type &#x60;DSP_ADVERTISING_ACCOUNT&#x60;. It will be &#x60;null&#x60; for accounts of other types. | [optional] [default to null]
**MarketplaceId** | **string** | The identifier of the marketplace to which the account is associated. See [this table](https://docs.developer.amazonservices.com/en_US/dev_guide/DG_Endpoints.html) for &#x60;marketplaceId&#x60; mappings. | [optional] [default to null]
**ProfileId** | **string** | The identifier of a profile associated with the advertiser account. Note that this value is only populated for a subset of account types: &#x60;[ SELLER, VENDOR, MARKETING_CLOUD ]&#x60;. It will be &#x60;null&#x60; for accounts of other types. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

