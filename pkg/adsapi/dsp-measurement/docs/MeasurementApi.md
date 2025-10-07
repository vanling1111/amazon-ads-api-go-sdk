# {{classname}}

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelMeasurementStudies**](MeasurementApi.md#CancelMeasurementStudies) | **Delete** /measurement/studies | Cancel existing studies. Once a study is cancelled it can not be resumed again.
[**CheckDSPAudienceResearchEligibility**](MeasurementApi.md#CheckDSPAudienceResearchEligibility) | **Post** /dsp/measurement/eligibility/audienceResearch | Checks the DSP AUDIENCE_RESEARCH study type eligibility against vendor products.
[**CheckDSPBrandLiftEligibility**](MeasurementApi.md#CheckDSPBrandLiftEligibility) | **Post** /dsp/measurement/eligibility/brandLift | Checks the DSP BRAND_LIFT study type eligibility against vendor products.
[**CheckDSPCreativeTestingEligibility**](MeasurementApi.md#CheckDSPCreativeTestingEligibility) | **Post** /dsp/measurement/eligibility/creativeTesting | Checks the DSP CREATIVE_TESTING study type eligibility against vendor products.
[**CheckDSPOmnichannelMetricsEligibility**](MeasurementApi.md#CheckDSPOmnichannelMetricsEligibility) | **Post** /dsp/measurement/eligibility/omnichannelMetrics | Checks the DSP OMNICHANNEL_METRICS study type eligibility against vendor products.
[**CheckPlanningEligibility**](MeasurementApi.md#CheckPlanningEligibility) | **Post** /measurement/planning/eligibility | Checks eligibility against all vendor products.
[**CreateDSPAudienceResearchStudy**](MeasurementApi.md#CreateDSPAudienceResearchStudy) | **Post** /dsp/measurement/studies/audienceResearch | Create new DSP AUDIENCE_RESEARCH study.
[**CreateDSPBrandLiftStudies**](MeasurementApi.md#CreateDSPBrandLiftStudies) | **Post** /dsp/measurement/studies/brandLift | Create new DSP BRAND_LIFT studies.
[**CreateDSPCreativeTestingStudy**](MeasurementApi.md#CreateDSPCreativeTestingStudy) | **Post** /dsp/measurement/studies/creativeTesting | Create new DSP CREATIVE_TESTING study.
[**CreateDSPOmnichannelMetricsStudies**](MeasurementApi.md#CreateDSPOmnichannelMetricsStudies) | **Post** /dsp/measurement/studies/omnichannelMetrics | Create new DSP OMNICHANNEL_METRICS studies.
[**CreateSurveys**](MeasurementApi.md#CreateSurveys) | **Post** /measurement/studies/surveys | Create new study surveys.
[**GetDSPAudienceResearchStudies**](MeasurementApi.md#GetDSPAudienceResearchStudies) | **Get** /dsp/measurement/studies/audienceResearch | Gets one or more DSP AUDIENCE_RESEARCH studies with requested study identifiers or an advertiser identifier.
[**GetDSPAudienceResearchStudyResult**](MeasurementApi.md#GetDSPAudienceResearchStudyResult) | **Get** /dsp/measurement/studies/audienceResearch/{studyId}/result | Get result of a DSP AUDIENCE_RESEARCH study.
[**GetDSPBrandLiftStudies**](MeasurementApi.md#GetDSPBrandLiftStudies) | **Get** /dsp/measurement/studies/brandLift | Gets one or more DSP BRAND_LIFT studies with requested study identifiers or an advertiser identifier.
[**GetDSPBrandLiftStudyResult**](MeasurementApi.md#GetDSPBrandLiftStudyResult) | **Get** /measurement/studies/brandLift/{studyId}/result | Get result of a DSP BRAND_LIFT study.
[**GetDSPCreativeTestingStudies**](MeasurementApi.md#GetDSPCreativeTestingStudies) | **Get** /dsp/measurement/studies/creativeTesting | Gets one or more DSP CREATIVE_TESTING studies with requested study identifiers or an advertiser identifier.
[**GetDSPCreativeTestingStudyResult**](MeasurementApi.md#GetDSPCreativeTestingStudyResult) | **Get** /dsp/measurement/studies/creativeTesting/{studyId}/result | Get result of a DSP CREATIVE_TESTING study.
[**GetDSPOmnichannelMetricsStudies**](MeasurementApi.md#GetDSPOmnichannelMetricsStudies) | **Get** /dsp/measurement/studies/omnichannelMetrics | Gets one or more DSP OMNICHANNEL_METRICS studies with requested study identifiers or an advertiser identifier.
[**GetDSPOmnichannelMetricsStudyResult**](MeasurementApi.md#GetDSPOmnichannelMetricsStudyResult) | **Get** /dsp/measurement/studies/omnichannelMetrics/{studyId}/result | Get result of a DSP OMNICHANNEL_METRICS study.
[**GetStudies**](MeasurementApi.md#GetStudies) | **Get** /measurement/studies | Gets base study objects given a list of studyIds or a list of advertiserIds.
[**GetSurveys**](MeasurementApi.md#GetSurveys) | **Get** /measurement/studies/surveys | Gets one or more study surveys with requested survey identifiers or a study identifier.
[**OmnichannelMetricsBrandSearch**](MeasurementApi.md#OmnichannelMetricsBrandSearch) | **Post** /measurement/vendorProducts/omnichannelMetrics/brands/list | Search for brands to be used in the OMNICHANNEL_METRICS vendor product.
[**UpdateDSPAudienceResearchStudy**](MeasurementApi.md#UpdateDSPAudienceResearchStudy) | **Put** /dsp/measurement/studies/audienceResearch/{studyId} | Update DSP AUDIENCE_RESEARCH study. This will be a full update.
[**UpdateDSPBrandLiftStudies**](MeasurementApi.md#UpdateDSPBrandLiftStudies) | **Put** /dsp/measurement/studies/brandLift | Update DSP BRAND_LIFT studies. This will be a full update.
[**UpdateDSPCreativeTestingStudy**](MeasurementApi.md#UpdateDSPCreativeTestingStudy) | **Put** /dsp/measurement/studies/creativeTesting/{studyId} | Update DSP CREATIVE_TESTING study. This will be a full update.
[**UpdateDSPOmnichannelMetricsStudies**](MeasurementApi.md#UpdateDSPOmnichannelMetricsStudies) | **Put** /dsp/measurement/studies/omnichannelMetrics | Update DSP OMNICHANNEL_METRICS studies. This will be a full update.
[**UpdateSurveys**](MeasurementApi.md#UpdateSurveys) | **Put** /measurement/studies/surveys | Update measurement surveys. This will be a full update.
[**VendorProduct**](MeasurementApi.md#VendorProduct) | **Post** /measurement/vendorProducts/list | Lists the supported measurement vendor products.
[**VendorProductPolicy**](MeasurementApi.md#VendorProductPolicy) | **Get** /measurement/vendorProducts/policies | Gets the policies for the specific vendor product(s).
[**VendorProductSurveyQuestionTemplates**](MeasurementApi.md#VendorProductSurveyQuestionTemplates) | **Get** /measurement/vendorProducts/surveyQuestionTemplates | Gets the survey question templates for the specific vendor product(s).

# **CancelMeasurementStudies**
> StudyResponsesV1 CancelMeasurementStudies(ctx, amazonAdvertisingAPIClientId, optional)
Cancel existing studies. Once a study is cancelled it can not be resumed again.

Cancel existing studies.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiCancelMeasurementStudiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiCancelMeasurementStudiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amazonAdsAccountId** | **optional.String**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **studyIds** | [**optional.Interface of []string**](string.md)| Study canonical identifiers to cancel. | 

### Return type

[**StudyResponsesV1**](StudyResponsesV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.studymanagement.v1+json, application/vnd.studymanagement.v1.1+json, application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckDSPAudienceResearchEligibility**
> EligibilityResponseV1M2 CheckDSPAudienceResearchEligibility(ctx, amazonAdvertisingAPIClientId, optional)
Checks the DSP AUDIENCE_RESEARCH study type eligibility against vendor products.

Checks the DSP AUDIENCE_RESEARCH study type eligibility status against vendor products.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiCheckDSPAudienceResearchEligibilityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiCheckDSPAudienceResearchEligibilityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DspAudienceResearchEligibilityRequestV1M2**](DspAudienceResearchEligibilityRequestV1M2.md)| The DSP audience research study eligibility request object. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **nextToken** | **optional.**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.**| Sets the maximum number of studies in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**EligibilityResponseV1M2**](EligibilityResponseV1M2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.measurementeligibility.v1.2+json
 - **Accept**: application/vnd.measurementeligibility.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckDSPBrandLiftEligibility**
> EligibilityResponseV1 CheckDSPBrandLiftEligibility(ctx, amazonAdvertisingAPIClientId, optional)
Checks the DSP BRAND_LIFT study type eligibility against vendor products.

Checks the DSP BRAND_LIFT study type eligibility status against vendor products.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiCheckDSPBrandLiftEligibilityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiCheckDSPBrandLiftEligibilityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DspBrandLiftEligibilityRequestV1**](DspBrandLiftEligibilityRequestV1.md)| The DSP brand lift eligibility request object. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **nextToken** | **optional.**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.**| Sets the maximum number of studies in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**EligibilityResponseV1**](EligibilityResponseV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.measurementeligibility.v1+json, application/vnd.measurementeligibility.v1.1+json
 - **Accept**: application/vnd.measurementeligibility.v1+json, application/vnd.measurementeligibility.v1.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckDSPCreativeTestingEligibility**
> EligibilityResponseV1M2 CheckDSPCreativeTestingEligibility(ctx, amazonAdvertisingAPIClientId, optional)
Checks the DSP CREATIVE_TESTING study type eligibility against vendor products.

Checks the DSP CREATIVE_TESTING study type eligibility status against vendor products.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiCheckDSPCreativeTestingEligibilityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiCheckDSPCreativeTestingEligibilityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DspCreativeTestingEligibilityRequestV1M2**](DspCreativeTestingEligibilityRequestV1M2.md)| The DSP creative testing study eligibility request object. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **nextToken** | **optional.**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.**| Sets the maximum number of studies in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**EligibilityResponseV1M2**](EligibilityResponseV1M2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.measurementeligibility.v1.2+json
 - **Accept**: application/vnd.measurementeligibility.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckDSPOmnichannelMetricsEligibility**
> EligibilityResponseV1M2 CheckDSPOmnichannelMetricsEligibility(ctx, amazonAdvertisingAPIClientId, optional)
Checks the DSP OMNICHANNEL_METRICS study type eligibility against vendor products.

Checks the DSP OMNICHANNEL_METRICS study type eligibility status against vendor products.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiCheckDSPOmnichannelMetricsEligibilityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiCheckDSPOmnichannelMetricsEligibilityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DspOmnichannelMetricsEligibilityRequestV1M2**](DspOmnichannelMetricsEligibilityRequestV1M2.md)| The DSP omnichannel metrics eligibility request object. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **nextToken** | **optional.**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.**| Sets the maximum number of studies in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**EligibilityResponseV1M2**](EligibilityResponseV1M2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.measurementeligibility.v1.2+json, application/vnd.measurementeligibility.v1.3+json
 - **Accept**: application/vnd.measurementeligibility.v1.2+json, application/vnd.measurementeligibility.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CheckPlanningEligibility**
> PlanningEligibilityResponseV1M3 CheckPlanningEligibility(ctx, amazonAdvertisingAPIClientId, optional)
Checks eligibility against all vendor products.

Checks eligibility against all vendor products.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiCheckPlanningEligibilityOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiCheckPlanningEligibilityOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PlanningEligibilityRequestV1M3**](PlanningEligibilityRequestV1M3.md)| Fetch measurement vendor products request object. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **nextToken** | **optional.**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.**| Sets the maximum number of studies in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**PlanningEligibilityResponseV1M3**](PlanningEligibilityResponseV1M3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.measurementeligibility.v1.1+json, application/vnd.measurementeligibility.v1.3+json
 - **Accept**: application/vnd.measurementeligibility.v1.1+json, application/vnd.measurementeligibility.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDSPAudienceResearchStudy**
> DspAudienceResearchStudyV1M2 CreateDSPAudienceResearchStudy(ctx, amazonAdvertisingAPIClientId, optional)
Create new DSP AUDIENCE_RESEARCH study.

Create new DSP AUDIENCE_RESEARCH study.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiCreateDSPAudienceResearchStudyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiCreateDSPAudienceResearchStudyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateDspAudienceResearchStudyV1M2**](CreateDspAudienceResearchStudyV1M2.md)| Create object for DSP AUDIENCE_RESEARCH study. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 

### Return type

[**DspAudienceResearchStudyV1M2**](DSPAudienceResearchStudyV1M2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.studymanagement.v1.2+json
 - **Accept**: application/vnd.studymanagement.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDSPBrandLiftStudies**
> StudyResponsesV1 CreateDSPBrandLiftStudies(ctx, amazonAdvertisingAPIClientId, optional)
Create new DSP BRAND_LIFT studies.

Create new DSP BRAND_LIFT studies.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiCreateDSPBrandLiftStudiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiCreateDSPBrandLiftStudiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []DspBrandLiftStudyV1**](DSPBrandLiftStudyV1.md)| An array of study objects. For each object, specify required fields and their values. Maximum length of the array is 1. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**StudyResponsesV1**](StudyResponsesV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.studymanagement.v1+json, application/vnd.studymanagement.v1.1+json, application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json
 - **Accept**: application/vnd.studymanagement.v1+json, application/vnd.studymanagement.v1.1+json, application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDSPCreativeTestingStudy**
> DspCreativeTestingStudyV1M2 CreateDSPCreativeTestingStudy(ctx, amazonAdvertisingAPIClientId, optional)
Create new DSP CREATIVE_TESTING study.

Create new DSP CREATIVE_TESTING study.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiCreateDSPCreativeTestingStudyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiCreateDSPCreativeTestingStudyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of CreateDspCreativeTestingStudyV1M2**](CreateDspCreativeTestingStudyV1M2.md)| Create object for DSP CREATIVE_TESTING study. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 

### Return type

[**DspCreativeTestingStudyV1M2**](DSPCreativeTestingStudyV1M2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.studymanagement.v1.2+json
 - **Accept**: application/vnd.studymanagement.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateDSPOmnichannelMetricsStudies**
> StudyResponsesV1 CreateDSPOmnichannelMetricsStudies(ctx, amazonAdvertisingAPIClientId, optional)
Create new DSP OMNICHANNEL_METRICS studies.

Create new DSP OMNICHANNEL_METRICS studies.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiCreateDSPOmnichannelMetricsStudiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiCreateDSPOmnichannelMetricsStudiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []DspOmnichannelMetricsStudyV1M2**](DSPOmnichannelMetricsStudyV1M2.md)| An array of study objects. For each object, specify required fields and their values. Maximum length of the array is 1. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 

### Return type

[**StudyResponsesV1**](StudyResponsesV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json
 - **Accept**: application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateSurveys**
> SurveyResponsesV1 CreateSurveys(ctx, amazonAdvertisingAPIClientId, optional)
Create new study surveys.

Create new study surveys.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiCreateSurveysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiCreateSurveysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []SurveyV1**](SurveyV1.md)| An array of measurement objects. For each object, specify required fields and their values. Maximum length of the array is 1. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**SurveyResponsesV1**](SurveyResponsesV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.studymanagement.v1+json, application/vnd.studymanagement.v1.1+json, application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json
 - **Accept**: application/vnd.studymanagement.v1+json, application/vnd.studymanagement.v1.1+json, application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDSPAudienceResearchStudies**
> PaginatedDspAudienceResearchStudiesV1M2 GetDSPAudienceResearchStudies(ctx, amazonAdvertisingAPIClientId, optional)
Gets one or more DSP AUDIENCE_RESEARCH studies with requested study identifiers or an advertiser identifier.

Gets one or more DSP AUDIENCE_RESEARCH studies with requested study identifiers or an advertiser identifier.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiGetDSPAudienceResearchStudiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiGetDSPAudienceResearchStudiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.String**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **studyIds** | [**optional.Interface of []string**](string.md)| Study canonical identifier to filter with. Either one of studyIds or advertiserId should be provided. | 
 **advertiserId** | **optional.String**| The advertiser canonical identifier. Either one of studyIds or advertiserId should be provided. | 
 **nextToken** | **optional.String**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.Int32**| Sets the maximum number of studies in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**PaginatedDspAudienceResearchStudiesV1M2**](PaginatedDSPAudienceResearchStudiesV1M2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.studymanagement.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDSPAudienceResearchStudyResult**
> AudienceResearchStudyResultV1M2 GetDSPAudienceResearchStudyResult(ctx, amazonAdvertisingAPIClientId, accept, studyId, optional)
Get result of a DSP AUDIENCE_RESEARCH study.

Get result of a DSP AUDIENCE_RESEARCH study. Returns 200 successful response if json resource is requested in Accept header. Returns a 307 Temporary Redirect response if any of the file types is requested and response includes a location header with the value set to an AWS S3 path where the result is located. The path expires after 60 seconds.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **accept** | **string**| The version(s) of the requested resource. Available version(s) - &#x60;application/vnd.measurementresult.v1.2+json&#x60;, &#x60;text/vnd.measurementresult.v1.2+csv&#x60;. | 
  **studyId** | **string**| The canonical identifier that represents a unique study. | 
 **optional** | ***MeasurementApiGetDSPAudienceResearchStudyResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiGetDSPAudienceResearchStudyResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.String**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 

### Return type

[**AudienceResearchStudyResultV1M2**](AudienceResearchStudyResultV1M2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.measurementresult.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDSPBrandLiftStudies**
> PaginatedDspBrandLiftStudiesV1 GetDSPBrandLiftStudies(ctx, amazonAdvertisingAPIClientId, optional)
Gets one or more DSP BRAND_LIFT studies with requested study identifiers or an advertiser identifier.

Gets one or more DSP BRAND_LIFT studies with requested study identifiers or an advertiser identifier.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiGetDSPBrandLiftStudiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiGetDSPBrandLiftStudiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amazonAdsAccountId** | **optional.String**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **studyIdFilters** | [**optional.Interface of []string**](string.md)| Study canonical identifier to filter with. Either one of studyIdFilters or advertiserId should be provided. | 
 **advertiserId** | **optional.String**| The advertiser canonical identifier. Either one of studyIdFilters or advertiserId should be provided. | 
 **nextToken** | **optional.String**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.Int32**| Sets the maximum number of studies in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**PaginatedDspBrandLiftStudiesV1**](PaginatedDSPBrandLiftStudiesV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.studymanagement.v1+json, application/vnd.studymanagement.v1.1+json, application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDSPBrandLiftStudyResult**
> BrandLiftStudyResultV1 GetDSPBrandLiftStudyResult(ctx, amazonAdvertisingAPIClientId, accept, studyId, optional)
Get result of a DSP BRAND_LIFT study.

Get result of a DSP BRAND_LIFT study. Returns 200 successful response if json resource is requested in Accept header. Returns a 307 Temporary Redirect response if any of the file types is requested and response includes a location header with the value set to an AWS S3 path where the result is located. The path expires after 60 seconds.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **accept** | **string**| The version(s) of the requested resource. Available version(s) - &#x60;application/vnd.measurementresult.v1+json&#x60;, &#x60;application/vnd.measurementresult.v1.1+json&#x60;, &#x60;text/vnd.measurementresult.v1+csv&#x60; | 
  **studyId** | **string**| The canonical identifier that represents a unique study. | 
 **optional** | ***MeasurementApiGetDSPBrandLiftStudyResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiGetDSPBrandLiftStudyResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdsAccountId** | **optional.String**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**BrandLiftStudyResultV1**](BrandLiftStudyResultV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.measurementresult.v1+json, application/vnd.measurementresult.v1.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDSPCreativeTestingStudies**
> PaginatedDspCreativeTestingStudiesV1M2 GetDSPCreativeTestingStudies(ctx, amazonAdvertisingAPIClientId, optional)
Gets one or more DSP CREATIVE_TESTING studies with requested study identifiers or an advertiser identifier.

Gets one or more DSP CREATIVE_TESTING studies with requested study identifiers or an advertiser identifier.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiGetDSPCreativeTestingStudiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiGetDSPCreativeTestingStudiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.String**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **studyIds** | [**optional.Interface of []string**](string.md)| Study canonical identifier to filter with. Either one of studyIds or advertiserId should be provided. | 
 **advertiserId** | **optional.String**| The advertiser canonical identifier. Either one of studyIds or advertiserId should be provided. | 
 **nextToken** | **optional.String**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.Int32**| Sets the maximum number of studies in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**PaginatedDspCreativeTestingStudiesV1M2**](PaginatedDSPCreativeTestingStudiesV1M2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.studymanagement.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDSPCreativeTestingStudyResult**
> CreativeTestingStudyResultV1M2 GetDSPCreativeTestingStudyResult(ctx, amazonAdvertisingAPIClientId, accept, studyId, optional)
Get result of a DSP CREATIVE_TESTING study.

Get result of a DSP CREATIVE_TESTING study. Returns 200 successful response if json resource is requested in Accept header. Returns a 307 Temporary Redirect response if any of the file types is requested and response includes a location header with the value set to an AWS S3 path where the result is located. The path expires after 60 seconds.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **accept** | **string**| The version(s) of the requested resource. Available version(s) - &#x60;application/vnd.measurementresult.v1.2+json&#x60;, &#x60;text/vnd.measurementresult.v1.2+csv&#x60;. | 
  **studyId** | **string**| The canonical identifier that represents a unique study. | 
 **optional** | ***MeasurementApiGetDSPCreativeTestingStudyResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiGetDSPCreativeTestingStudyResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.String**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 

### Return type

[**CreativeTestingStudyResultV1M2**](CreativeTestingStudyResultV1M2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.measurementresult.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDSPOmnichannelMetricsStudies**
> PaginatedDspOmnichannelMetricsStudiesV1M2 GetDSPOmnichannelMetricsStudies(ctx, amazonAdvertisingAPIClientId, optional)
Gets one or more DSP OMNICHANNEL_METRICS studies with requested study identifiers or an advertiser identifier.

Gets one or more DSP OMNICHANNEL_METRICS studies with requested study identifiers or an advertiser identifier.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiGetDSPOmnichannelMetricsStudiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiGetDSPOmnichannelMetricsStudiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.String**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **studyIds** | [**optional.Interface of []string**](string.md)| Study canonical identifier to filter with. Either one of studyIds or advertiserId should be provided. | 
 **advertiserId** | **optional.String**| The advertiser canonical identifier. Either one of studyIdFilters or advertiserId should be provided. | 
 **nextToken** | **optional.String**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.Int32**| Sets the maximum number of studies in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**PaginatedDspOmnichannelMetricsStudiesV1M2**](PaginatedDSPOmnichannelMetricsStudiesV1M2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDSPOmnichannelMetricsStudyResult**
> GetDSPOmnichannelMetricsStudyResult(ctx, amazonAdvertisingAPIClientId, accept, studyId, optional)
Get result of a DSP OMNICHANNEL_METRICS study.

Get result of a DSP OMNICHANNEL_METRICS study. Returns a 307 Temporary Redirect response if any of the file types is requested and response includes a location header with the value set to an AWS S3 path where the result is located. The path expires after 60 seconds. Accept header does not support json for OMNICHANNEL_METRICS study type.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **accept** | **string**| The version(s) of the requested resource. Available version(s) - &#x60;text/vnd.measurementresult.v1+xlsx&#x60; | 
  **studyId** | **string**| The canonical identifier that represents a unique study. | 
 **optional** | ***MeasurementApiGetDSPOmnichannelMetricsStudyResultOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiGetDSPOmnichannelMetricsStudyResultOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.String**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.measurementresult.v1.2+json, application/vnd.measurementresult.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStudies**
> PaginatedBaseStudiesV1 GetStudies(ctx, amazonAdvertisingAPIClientId, optional)
Gets base study objects given a list of studyIds or a list of advertiserIds.

Gets base study objects given a list of studyIds or a list of advertiserIds.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiGetStudiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiGetStudiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amazonAdsAccountId** | **optional.String**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **studyIds** | [**optional.Interface of []string**](string.md)| Study canonical identifier to filter with. Either one of studyIds or advertiserId should be provided. | 
 **advertiserId** | **optional.String**| The advertiser canonical identifier. Either one of studyIds or advertiserId should be provided. | 
 **nextToken** | **optional.String**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.Int32**| Sets the maximum number of studies in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**PaginatedBaseStudiesV1**](PaginatedBaseStudiesV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.studymanagement.v1+json, application/vnd.studymanagement.v1.1+json, application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSurveys**
> PaginatedSurveysV1 GetSurveys(ctx, amazonAdvertisingAPIClientId, optional)
Gets one or more study surveys with requested survey identifiers or a study identifier.

Gets one or more study surveys with requested survey identifiers or a study identifier.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiGetSurveysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiGetSurveysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amazonAdsAccountId** | **optional.String**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **surveyIds** | [**optional.Interface of []string**](string.md)| Survey canonical identifier to filter with. Either one of surveyIds or studyId should be provided. | 
 **studyId** | **optional.String**| A study canonical identifier. Either one of surveyIds or studyId should be provided. | 
 **nextToken** | **optional.String**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.Int32**| Sets the maximum number of studies in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**PaginatedSurveysV1**](PaginatedSurveysV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.studymanagement.v1+json, application/vnd.studymanagement.v1.1+json, application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OmnichannelMetricsBrandSearch**
> PaginatedOmnichannelMetricsBrandsV1M2 OmnichannelMetricsBrandSearch(ctx, amazonAdvertisingAPIClientId, optional)
Search for brands to be used in the OMNICHANNEL_METRICS vendor product.

Search for brands to be used in the OMNICHANNEL_METRICS vendor product.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiOmnichannelMetricsBrandSearchOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiOmnichannelMetricsBrandSearchOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of OmnichannelMetricsBrandSearchRequestV1M2**](OmnichannelMetricsBrandSearchRequestV1M2.md)| Fetch measurement vendor products request object. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **nextToken** | **optional.**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.**| Sets the maximum number of brands in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**PaginatedOmnichannelMetricsBrandsV1M2**](PaginatedOmnichannelMetricsBrandsV1M2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.ocmbrands.v1.2+json, application/vnd.ocmbrands.v1.3+json
 - **Accept**: application/vnd.ocmbrands.v1.2+json, application/vnd.ocmbrands.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDSPAudienceResearchStudy**
> DspAudienceResearchStudyV1M2 UpdateDSPAudienceResearchStudy(ctx, amazonAdvertisingAPIClientId, studyId, optional)
Update DSP AUDIENCE_RESEARCH study. This will be a full update.

Update DSP AUDIENCE_RESEARCH study. This will be a full update.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **studyId** | **string**| The canonical identifier that represents a unique study. | 
 **optional** | ***MeasurementApiUpdateDSPAudienceResearchStudyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiUpdateDSPAudienceResearchStudyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateDspAudienceResearchStudyV1M2**](UpdateDspAudienceResearchStudyV1M2.md)| Update object for DSP AUDIENCE_RESEARCH study. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 

### Return type

[**DspAudienceResearchStudyV1M2**](DSPAudienceResearchStudyV1M2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.studymanagement.v1.2+json
 - **Accept**: application/vnd.studymanagement.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDSPBrandLiftStudies**
> StudyResponsesV1 UpdateDSPBrandLiftStudies(ctx, amazonAdvertisingAPIClientId, optional)
Update DSP BRAND_LIFT studies. This will be a full update.

Update DSP BRAND_LIFT studies. This will be a full update.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiUpdateDSPBrandLiftStudiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiUpdateDSPBrandLiftStudiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []DspBrandLiftStudyV1**](DSPBrandLiftStudyV1.md)| An array of measurement objects. For each object, specify required fields and their values. Maximum length of the array is 1. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**StudyResponsesV1**](StudyResponsesV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.studymanagement.v1+json, application/vnd.studymanagement.v1.1+json, application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json
 - **Accept**: application/vnd.studymanagement.v1+json, application/vnd.studymanagement.v1.1+json, application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDSPCreativeTestingStudy**
> DspCreativeTestingStudyV1M2 UpdateDSPCreativeTestingStudy(ctx, amazonAdvertisingAPIClientId, studyId, optional)
Update DSP CREATIVE_TESTING study. This will be a full update.

Update DSP CREATIVE_TESTING study. This will be a full update.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
  **studyId** | **string**| The canonical identifier that represents a unique study. | 
 **optional** | ***MeasurementApiUpdateDSPCreativeTestingStudyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiUpdateDSPCreativeTestingStudyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of UpdateDspCreativeTestingStudyV1M2**](UpdateDspCreativeTestingStudyV1M2.md)| Update object for DSP CREATIVE_TESTING study. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 

### Return type

[**DspCreativeTestingStudyV1M2**](DSPCreativeTestingStudyV1M2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.studymanagement.v1.2+json
 - **Accept**: application/vnd.studymanagement.v1.2+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateDSPOmnichannelMetricsStudies**
> StudyResponsesV1 UpdateDSPOmnichannelMetricsStudies(ctx, amazonAdvertisingAPIClientId, optional)
Update DSP OMNICHANNEL_METRICS studies. This will be a full update.

Update DSP OMNICHANNEL_METRICS studies. This will be a full update.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_edit\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiUpdateDSPOmnichannelMetricsStudiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiUpdateDSPOmnichannelMetricsStudiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []DspOmnichannelMetricsStudyV1M2**](DSPOmnichannelMetricsStudyV1M2.md)| An array of measurement objects. For each object, specify required fields and their values. Maximum length of the array is 1. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 

### Return type

[**StudyResponsesV1**](StudyResponsesV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json
 - **Accept**: application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSurveys**
> SurveyResponsesV1 UpdateSurveys(ctx, amazonAdvertisingAPIClientId, optional)
Update measurement surveys. This will be a full update.

Update measurement surveys. This will be a full update.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiUpdateSurveysOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiUpdateSurveysOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of []SurveyV1**](SurveyV1.md)| An array of survey objects. For each object, specify required fields and their values. Maximum length of the array is 1. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 

### Return type

[**SurveyResponsesV1**](SurveyResponsesV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.studymanagement.v1+json, application/vnd.studymanagement.v1.1+json, application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json
 - **Accept**: application/vnd.studymanagement.v1+json, application/vnd.studymanagement.v1.1+json, application/vnd.studymanagement.v1.2+json, application/vnd.studymanagement.v1.3+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VendorProduct**
> PaginatedVendorProductsV1 VendorProduct(ctx, amazonAdvertisingAPIClientId, optional)
Lists the supported measurement vendor products.

Lists the supported measurement vendors products.  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiVendorProductOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiVendorProductOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of VendorProductRequestV1**](VendorProductRequestV1.md)| Fetch measurement vendor products request object. | 
 **amazonAdvertisingAPIScope** | **optional.**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **amazonAdsAccountId** | **optional.**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **nextToken** | **optional.**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.**| Sets the maximum number of studies in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**PaginatedVendorProductsV1**](PaginatedVendorProductsV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/vnd.measurementvendor.v1+json, application/vnd.measurementvendor.v1.1+json
 - **Accept**: application/vnd.measurementvendor.v1+json, application/vnd.measurementvendor.v1.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VendorProductPolicy**
> PaginatedVendorProductPoliciesV1 VendorProductPolicy(ctx, amazonAdvertisingAPIClientId, optional)
Gets the policies for the specific vendor product(s).

Gets the policies for the specific vendor product(s).  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiVendorProductPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiVendorProductPolicyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amazonAdsAccountId** | **optional.String**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **vendorProductIds** | [**optional.Interface of []string**](string.md)| Vendor product canonical identifier to filter with. | 
 **nextToken** | **optional.String**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.Int32**| Sets the maximum number of studies in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**PaginatedVendorProductPoliciesV1**](PaginatedVendorProductPoliciesV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.measurementvendor.v1+json, application/vnd.measurementvendor.v1.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VendorProductSurveyQuestionTemplates**
> PaginatedSurveyQuestionTemplatesV1 VendorProductSurveyQuestionTemplates(ctx, amazonAdvertisingAPIClientId, optional)
Gets the survey question templates for the specific vendor product(s).

Gets the survey question templates for the specific vendor product(s).  **Authorized resource type**: DSP Rodeo Entity ID, DSP Advertiser Account ID  **Parameter name**: Amazon-Ads-AccountId  **Parameter in**: header  **Requires one of these permissions**: [\"campaign_view\"]

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **amazonAdvertisingAPIClientId** | **string**| The identifier of a client associated with a \&quot;Login with Amazon\&quot; account. | 
 **optional** | ***MeasurementApiVendorProductSurveyQuestionTemplatesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a MeasurementApiVendorProductSurveyQuestionTemplatesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **amazonAdsAccountId** | **optional.String**| The identifier that describe DSP advertiser level accounts that exists under a manager account (previously under a DSP entity). Exposed in the /dsp/advertisers API. | 
 **amazonAdvertisingAPIScope** | **optional.String**| The identifier of a profile associated with the advertiser account. Use &#x60;GET&#x60; method on Profiles resource to list profiles associated with the access token passed in the HTTP Authorization header and choose profile id &#x60;profileId&#x60; from the response to pass it as input. | 
 **vendorProductIds** | [**optional.Interface of []string**](string.md)| Vendor product canonical identifier to filter with. | 
 **surveyQuestionTemplateIds** | [**optional.Interface of []string**](string.md)| Vendor product survey question template identifier to filter with. | 
 **nextToken** | **optional.String**| Token from a previous request. Use in conjunction with the &#x60;maxResults&#x60; parameter to control pagination of the returned array. | 
 **maxResults** | **optional.Int32**| Sets the maximum number of studies in the returned array. Use in conjunction with the &#x60;nextToken&#x60; parameter to control pagination. The range for maxResults is [1,100] with default as 10. For example, supplying maxResults&#x3D;20 with a previously returned token will fetch up to the next 20 items. In some cases, fewer items may be returned. | [default to 10]

### Return type

[**PaginatedSurveyQuestionTemplatesV1**](PaginatedSurveyQuestionTemplatesV1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/vnd.measurementvendor.v1+json, application/vnd.measurementvendor.v1.1+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

