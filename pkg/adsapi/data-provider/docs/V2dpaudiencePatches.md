# V2dpaudiencePatches

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | **string** | Specifies the type of operation. Valid operations are &#x60;add&#x60; or &#x60;remove&#x60;. | [optional] [default to null]
**Path** | **string** | A formatted string that specifies the URL of the record. The format of the string is &#x60;/&lt;recordIdType&gt;-&lt;recordIdValue&gt;/audiences&#x60;, where &#x27;recordIdType&#x27; specifies the record&#x27;s origin and &#x27;recordIdValue&#x27; specifies the record&#x27;s Id. Valid &#x27;recordIdType&#x27; values are &#x60;COOKIE&#x60;, a cookie Id sent from a data provider to Amazon by a cookie sync; &#x60;MAID&#x60;, a mobile advertising identifier; &#x60;EXTERNAL_USER_ID&#x60;, an external id defined by data providers. | [optional] [default to null]
**Value** | **[]int64** |  | [optional] [default to null]
**Consent** | [***V2dpaudienceConsent**](v2dpaudience_consent.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

