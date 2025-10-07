# AudienceFilterV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | Field to filter by. Supported enums are &#x27;audienceName&#x27;, &#x27;category&#x27;, &#x27;categoryPath&#x27;, &#x27;audienceId&#x27; and &#x27;status&#x27;. The &#x27;audienceName&#x27; is a broad match filter but not an exact match. The &#x27;category&#x27; enum returns all audiences under a high-level category, whereas the &#x27;categoryPath&#x27; enum expects a path of nodes in the taxonomy tree and returns audiences attached directly to the node at the specified path. | [optional] [default to null]
**Operator** | **string** | Operator to apply to the specified filter. | [optional] [default to OPERATOR.EQ]
**Values** | **[]string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

