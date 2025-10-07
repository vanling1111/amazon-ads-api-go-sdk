# CreativeModeration

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreativeId** | **float64** | Unique identifier of the creative. | [default to null]
**CreativeType** | [***CreativeTypeInCreativeResponse**](CreativeTypeInCreativeResponse.md) |  | [default to null]
**ModerationStatus** | **string** | The moderation status of the creative. |Status|Description| |------|-----------| |APPROVED|Moderation for the creative is complete.| |IN_PROGRESS|Moderation for the creative is in progress. The expected date and time for completion are specfied in the &#x60;etaForModeration&#x60; field.| |REJECTED|The creative has failed moderation. Specific information about the content that violated policy is available in &#x60;policyViolations&#x60;.| | [default to null]
**EtaForModeration** | [**time.Time**](time.Time.md) | Expected date and time by which moderation will be complete. | [default to null]
**PolicyViolations** | [**[]CreativeModerationPolicyViolations**](CreativeModeration_policyViolations.md) | A list of policy violations for a creative that has failed moderation. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

