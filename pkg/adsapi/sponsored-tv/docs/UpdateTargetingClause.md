# UpdateTargetingClause

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bid** | **float64** | The bid for Ads sourced using the Targeting Clause. The default bid would be $15 CPM for the US marketplace.  Functionality varies by date and advertiser type:  After March 4th 2025 (2025-03-04): * Endemic Advertisers: Fully functional * Non-Endemic Advertisers: This field will no longer be processed - any provided values will be ignored. Bid at the adGroup level will be honored  After June 30th 2025 (2025-06-30): * For all Advertisers: This field will no longer be processed - any provided values will be ignored. Bid at the adGroup level will be honored  Note: The field remains in the API for backwards compatibility but has no effect after the dates specified above. If you are currently using this field, you should remove the implementation as the values are not being processed. | [optional] [default to null]
**State** | [***CreateOrUpdateEntityState**](CreateOrUpdateEntityState.md) |  | [default to null]
**TargetId** | **string** | The Targeting Clause ID. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

