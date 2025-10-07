# InlineResponse20017

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **int64** | The campaign identifier. | [optional] [default to null]
**ModerationStatus** | **string** | The moderation status of the campaign. |Status|Description| |------|-----------| |APPROVED|Moderation for the campaign is complete.| |IN_PROGRESS|Moderation for the campaign is in progress. The expected date and time for completion are specfied in the &#x60;etaForModeration&#x60; field.| |REJECTED|The campaign has failed moderation. Specific information about the content that violated policy is available in &#x60;policyViolations&#x60;.| | [optional] [default to null]
**EtaForModeration** | [**time.Time**](time.Time.md) | Expected date and time by which moderation will be complete. Note that this field is present in the response only when &#x60;moderationStatus&#x60; is set to &#x60;IN_PROGRESS&#x60;. | [optional] [default to null]
**PolicyViolations** | [**[]InlineResponse20017PolicyViolations**](inline_response_200_17_policyViolations.md) | A list of policy violations for a campaign that has failed moderation. Note that this field is present in the response only when &#x60;moderationStatus&#x60; is set to &#x60;REJECTED&#x60;. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

