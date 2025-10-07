# ListPostsResponseContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsNumPostsOverLimit** | **bool** | No more than 10,000 posts can be sorted by a field other than CREATED_DATE. This value represents if that limit has been reached. If this is the case, you should instruct your users to add more filters to the call. | [optional] [default to null]
**NextToken** | **string** | Pagination token to get next page of results from previous call. | [optional] [default to null]
**Posts** | [**[]Post**](Post.md) | A list of posts. | [optional] [default to null]
**TotalPosts** | **float64** | Total number of posts that exist within the conditions of the request. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

