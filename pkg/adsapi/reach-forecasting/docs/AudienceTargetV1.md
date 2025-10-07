# AudienceTargetV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudienceId** | **string** | The unique identifier for the audience. Use the [audiences](https://advertising.amazon.com/API/docs/en-us/audiences/#/Discovery/listAudiences) discovery resource to look up audience identifiers. &lt;li&gt;Included groups are joined by an intersection. Users must be in all of the groups specified to be included.&lt;/li&gt; &lt;li&gt;Audiences within an inclusion group are joined by OR/ANY logic. Audiences within the exclusion group are joined by AND/ALL logic. This means that users can be in any of the audiences inside a group to be included or excluded.&lt;/li&gt; &lt;li&gt;All audiences within the same group must either be included or excluded.&lt;/li&gt; &lt;li&gt;You can specify up to 10 groups to be included, but only 1 group to be excluded.&lt;/li&gt; &lt;li&gt;Up to 500 audience IDs can be specified per group.&lt;/li&gt; | [default to null]
**GroupId** | **string** | The string identifying a group of audiences. Only numbers formatted as strings are accepted (e.g., \&quot;1\&quot;). To add audiences to a new group, choose any string not currently being used on this forecast. To add audiences to an existing group, use the existing groupId from this forecast. You may specify up to 10 include groups and 1 exclude group. | [optional] [default to null]
**TargetType** | [***AudienceTargetTypeV1**](AudienceTargetTypeV1.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

