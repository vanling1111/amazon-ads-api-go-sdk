# Campaign

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetType** | [***BudgetType**](BudgetType.md) |  | [default to null]
**RuleBasedBudget** | [***RuleBasedBudget**](RuleBasedBudget.md) |  | [optional] [default to null]
**BrandEntityId** | **string** |  | [optional] [default to null]
**IsMultiAdGroupsEnabled** | **bool** |  | [optional] [default to null]
**Goal** | **string** | Goal will allow you to set goal type to help drive your campaign performance. If no goal is selected then it will default to PAGE_VISIT. The goal type of the campaign. - BRAND_IMPRESSION_SHARE - This goal will allow you grown your brand impression share on top of search placements - PAGE_VISIT [DEFAULT] - This goal drives traffic to your landing and detail pages through all placements - ACQUIRE_NEW_CUSTOMERS - This property is a PREVIEW ONLY and cannot be used as part of a request or response. This goal drives new customer acquisition for your brands. - AD_VIEWS - This property is a PREVIEW ONLY and cannot be used as part of a request or response. This goal maximizes view for your ads. | [optional] [default to null]
**Bidding** | [***Bidding**](Bidding.md) |  | [optional] [default to null]
**EndDate** | **string** | The format of the date is YYYY-MM-DD. | [optional] [default to null]
**CampaignId** | **string** | The identifier of the campaign. | [default to null]
**ProductLocation** | [***ProductLocation**](ProductLocation.md) |  | [optional] [default to null]
**Tags** | [***map[string]string**](map.md) |  | [optional] [default to null]
**PortfolioId** | **string** | The identifier of an existing portfolio to which the campaign is associated. | [optional] [default to null]
**CostType** | **string** | The costType can be set to determines how the campaign will bid and charge. To view the bid maximums and minimums by geography and costType, see https://advertising.amazon.com/API/docs/en-us/concepts/limits#bid-constraints-by-marketplace - CPC [Default] - Cost per click. The performance of this campaign is measured by the clicks triggered by the ad. - VCPM - Cost per 1000 viewable impressions. The performance of this campaign is measured by the viewable impressions triggered by the ad. | [optional] [default to null]
**SmartDefault** | **[]string** | The smartDefault specifies a list of the smart default options for the campaign.  &#x60;smartDefault&#x60; is optional for create campaign requests. &#x60;smartDefault&#x60; are applicable to all applicable child entities of the campaign and are not editable once the campaign is created. When using [\&quot;TARGETING\&quot;], targets will be automatically added based on the goal selected.  When [\&quot;MANUAL\&quot;] is selected, you will still be required to manually add targets.  If you don&#x27;t specify &#x60;smartDefault&#x60;, default value will be applied based on &#x60;goal&#x60; . If campaign&#x27;s &#x60;goal&#x60; is selected, &#x60;smartDefault&#x60; will be set to [\&quot;TARGETING\&quot;].  Otherwise, a campaign&#x27;s &#x60;smartDefault&#x60; will be set to [\&quot;MANUAL\&quot;].  Each element in smartDefault can be set to determines which default strategy to be used - MANUAL - Manual settings, no smart default be applied to the campaign, if MANUAL is added in the list, no other items are allowed in the list (the list must contains only one item) - TARGETING - Smart Default Targeting creation, will automatically creating targetings when create ad group  Example: [\&quot;TARGETING\&quot;] | [optional] [default to null]
**SiteRestrictions** | [**[]SiteRestriction**](SiteRestriction.md) | Restrict the ad to a particular site. siteRestrictions is an optional field.  If this field is not set, ads from the campaign will appear on Amazon - including both Amazon retail and Amazon Business.  Please note that: 1) AMAZON_BUSINESS option is only available for Amazon Business operated marketplaces (US, CA, MX, UK, DE, FR, IT, ES, IN, JP, AU); 2) siteRestrictions cannot be changed post campaign creation; 3) siteRestrictions doesn&#x27;t support shopperCohortBidding setting. | [optional] [default to null]
**Name** | **string** | The name of the campaign. | [default to null]
**State** | [***EntityState**](EntityState.md) |  | [default to null]
**StartDate** | **string** | The format of the date is YYYY-MM-DD. | [optional] [default to null]
**Budget** | **float64** |  | [default to null]
**ExtendedData** | [***CampaignExtendedData**](CampaignExtendedData.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

