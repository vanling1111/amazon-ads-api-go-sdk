# MobileMeasurementPartnerAppRegistrationV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppName** | **string** | The name of the application. | [default to null]
**BundleId** | **string** | The ID of the application with the app store it is registered with. The bundleId + platform + mmpName must be unique within your advertiser account. | [default to null]
**ConversionsCreated** | **float64** | The number of conversions associated with this mobile application. | [optional] [default to null]
**LastEventReceived** | [**time.Time**](time.Time.md) | The latest timestamp of when a conversion event for the mobile application was imported in ISO format (YYYY-MM-DDThh:mm:ssTZD). | [optional] [default to null]
**MmpAppId** | **string** | The id of the mobile measurement partner app registration. | [optional] [default to null]
**MmpName** | [***MobileMeasurementPartnerNameV1**](MobileMeasurementPartnerNameV1.md) |  | [default to null]
**Platform** | [***MobileMeasurementPartnerPlatformV1**](MobileMeasurementPartnerPlatformV1.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

