# PartnerOpportunitiesOpportunityV1

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | **string** | The intended audience of the opportunity. For example, it might be targeted towards optimizing partner metrics or the metrics of advertisers that the partner manages. | [default to null]
**CallToAction** | **string** | An explanation of why it&#x27;s recommended to take the actions detailed in the opportunity&#x27;s data file. | [default to null]
**CreatedDate** | [**time.Time**](time.Time.md) | When the opportunity was created, in ISO 8601 format. This should never change. | [default to null]
**DataMetadata** | [***AllOfPartnerOpportunitiesOpportunityV1DataMetadata**](AllOfPartnerOpportunitiesOpportunityV1DataMetadata.md) | Contains the most recent data file information for the opportunity.  Can be used to track the availability of a partner opportunity data file:    if dataMetadata.rowCount &gt; 0. | [default to null]
**DataUrl** | **string** | The URL through which an opportunity&#x27;s data file (in CSV format) can be downloaded.  A simple GET request is all that is necessary, which will automatically redirect to a presigned, short-lived URL.  URLs expire in 15 minutes. | [default to null]
**Description** | **string** | A detailed description of the opportunity and how it is pertinent to partners. May provide a summary of the underlying data provided in the opportunity data file. | [default to null]
**Objective** | **string** | The objective of the opportunity. For example, an objective might be to drive sales, raise brand awareness, etc.  Deprecated. | [optional] [default to null]
**ObjectiveType** | **string** | The objective type of the opportunity. For example, an objective type might be around providing the unlaunched ASINs you can optimize or deals you can action on. | [default to null]
**PartnerOpportunityId** | **string** | The unique ID for the opportunity. | [default to null]
**Product** | **string** | The Amazon Advertising product to which the opportunity corresponds, like Amazon DSP, Video Ads, etc. | [default to null]
**Title** | **string** | The title of the opportunity. | [default to null]
**UpdatedDate** | [**time.Time**](time.Time.md) | When the opportunity was last updated, in ISO 8601 format. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

