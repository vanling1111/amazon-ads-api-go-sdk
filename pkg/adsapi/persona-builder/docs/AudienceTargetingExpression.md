# AudienceTargetingExpression

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audiences** | [**[]AudienceV1**](AudienceV1.md) | Specify groups of audiences to include or exclude when targeting.&lt;ul&gt;&lt;li&gt;Included groups are joined by an intersection. Users must be in all of the groups specified to be included.&lt;/li&gt;&lt;li&gt;Audiences within a group are joined by a union. Users can be in any of the audiences specified to be included/excluded.&lt;/li&gt;&lt;li&gt;Users within the same group must either be included or excluded.&lt;/li&gt;&lt;li&gt;You can specify up to 10 groups to be included, but only 1 group to be excluded.&lt;/li&gt;&lt;li&gt;Up to 500 audienceIds can be specified per group.&lt;/li&gt;&lt;ul&gt; | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

