# Fee

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cost** | [***CurrencyAmount**](currencyAmount.md) |  | [default to null]
**FeeIdentifiers** | [***FeeIdentifiers**](Fee Identifiers.md) |  | [optional] [default to null]
**FeeType** | **string** | * &#x60;PLATFORM_FEE&#x60;: Billable fee set at the Rodeo Entity level by internal users which reflects the cost of using the Amazon DSP   * Supply Cost * Platform Fee % * &#x60;AGENCY_FEE&#x60;: Non-billable fee set at the Rodeo Order level by external users which reflects the fee that the agency is charging the end customer   * Total Cost * Agency Fee % * &#x60;AUDIENCE_FEE&#x60;: Billable fee automatically calculated at the Rodeo Line Item level when external users choose Amazon 1P data segments for campaign targeting   * Impressions with Audience Fees * Audience Fee (CPM)/1000 * &#x60;3P_[AUTO_]NON_ABSORBED_FEE&#x60;: Billable fee automatically calculated at the Rodeo Line Item level when external users choose Automotive data segments and/or DMP data segments for campaign targeting   * Impressions * Billable 3p Fee / 1000 * &#x60;REGULATORY_ADVERTISING_FEE&#x60;: Fees derive from ads serving in specific countries and/or for ads purchased from advertisers in specific countries during the period in which you are billed. * &#x60;OMNICHANNEL_METRICS_FEE&#x60;: Billable fee set at DSP order level by internal users, which reflects the cost of using Omnichannel metrics measurement   * Supply Cost * Omnichannel Metrics Fee % * &#x60;3P_PREBID_FEE&#x60;: Billable fee automatically calculated when external users choose third party prebid targeting products for supply quality filtering.   * Impressions with 3P Prebid Fees * 3P Prebid Fee (CPM)/1000 % * &#x60;REWARDED_ADS_COST&#x60;: Billable, non taxable cost calculated for the rewards distributed when external users choose Rewarded Ads creative template for campaign targeting  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

