# UpdateTargetingClause

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetId** | **int64** |  | [default to null]
**State** | **string** |  | [optional] [default to null]
**Bid** | **float32** | The bid will override the adGroup bid if specified. This field is not used for negative targeting clauses. The bid must be less than the maximum allowable bid for the campaign&#x27;s marketplace; for a list of maximum allowable bids, find the [\&quot;Bid constraints by marketplace\&quot; table in our documentation overview](https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace). You cannot manually set a bid when the targeting clause&#x27;s adGroup has an enabled optimization rule. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

