# TopCategoriesPurchasedResponseContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdatedAt** | [**time.Time**](time.Time.md) | UTC timestamp in ISO 8601 format indicating when insight was last generated for the input audience set. | [default to null]
**NextToken** | **string** | Optional: If present, there are more insights than initially returned. Use this token to call the operation again                     and have the additional insights returned. The token is valid for 8 hours from the initial request. | [optional] [default to null]
**RetailCategories** | [**[]TopRetailCategoryInsight**](TopRetailCategoryInsight.md) | Top retail categories purchased by customers in the input expression., ordered by the affinity score.                         Affinity represents a measure of how likely customers in the input expression are to make a purchase                     from the category. An affinity of 5 indicates that customers in the input expression are 5 times as likely to                     buy from this category than the average of customers on Amazon. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

