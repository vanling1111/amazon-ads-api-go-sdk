# StartMigrationJobRequestContent

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignIds** | **[]string** | Provide list of campaign ids that needs to be migrated | [default to null]
**IsStagedMigration** | **bool** | Set this flag to true if you want generate new campaign ID based on V3 campaign ID. These campaigns will not be visible through V4 campaign list call. If set to true not all campaign entities such as ad group, targeting, ad, or creatives are created. Use this flag for staging purpose only. By default it will always be false | [optional] [default to null]
**NewCampaignState** | **string** | This is optional parameter. By default, the new migrated campaigns will have the original status of V3 campaigns. If this parameter is set, then all newly migrated campaigns will have this state.  Supported campaign states | State                                              |  Description | |----------------------------------------------------------|--------------| | ENABLED                               | Campaign entity has ENABLED state | | [optional] [default to null]
**EnableThemeTargeting** | **bool** | By default, theme targeting is set true if no value is provide. To disable theme targeting, set this flag to false. | [default to null]
**BrandEntityId** | **string** | Please note that brandEntityId is only required for sellers. You can get the brandEntityId by calling the &lt;a href &#x3D; https://advertising.amazon.com/API/docs/en-us/sponsored-brands/3-0/openapi#tag/Brands/operation/getBrands&gt;GET /brands&lt;/a&gt; endpoint. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

