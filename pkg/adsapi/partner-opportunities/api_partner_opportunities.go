
/*
 * Partner Opportunities
 *
 * <br><br>**Purpose of the API**<br><br>The API is for partners to retrieve **'opportunities'** relevant to optimizing their book of business. Opportunities are actionable insights for partners that drive value to advertisers.<br><br>For a step by step guide to using our API, we recommend our official [developer guide](https://advertising.amazon.com/API/docs/en-us/guides/recommendations/partner-opportunities/overview).<br><br>**How should partners use this set of APIs?**<br><br>Partners should use this set of APIs by first calling the **GET** `/partnerOpportunities` route to obtain a list of opportunities that are pertinent (and exclusive) to their book of business (the advertisers they manage).<br><br>It should be noted that the opportunities returned in this list call may not have data for that day (opportunities are updated every 24 hours). To check if an opportunity has data, introspect on the value of the `opportunities[].dataMetadata.rowCount` field. If the value is greater than zero, data is available and a data file can be downloaded.<br><br>Once an opportunity of interest has been identified, the given opportunity data can be downloaded using the **GET** `/partnerOpportunities/{partnerOpportunityId}/file` route, which will redirect to a pre-signed file download URL containing all opportunity data in a CSV format. This operation will return a 404 if no data is available for the given opportunity.<br><br>Partners can get aggregated information about opportunities available to them using the **GET** `/partnerOpportunities/summary` route. This information includes total opportunity count, count of opportunities with data available to be downloaded, and a count of unique advertisers across all opportunities. <br><br>**Required Headers**<br><br>Currently, there are two headers that are required for all API calls. If these headers are not correctly provided and properly formatted, the API call requests will fail.<br><br>1. `Accept:` Must be set to a supported API version (see below), using the format described on the [Advertising API portal](https://advertising.amazon.com/API/docs/en-us/concepts/compatibility-versioning-policy). The request and response payloads are identical between versions v1, v1.1, and v1.2.<br>    - Version 1 (Recommended): `'application/vnd.partneropportunity.v1+json'`<br>    - Version 1.1: `'application/vnd.partneropportunity.v1.1+json'`<br>    - Version 1.2: `'application/vnd.partneropportunity.v1.2+json'`<br>  <br>2. `Amazon-Advertising-API-Manager-Account:` 'Partner Network Account ID' which is accessible from Partner Network under the ['User settings'](https://advertising.amazon.com/partner-network/settings) link in the upper right corner.<br>    - Example: `'amzn1.ads1.ma1.abcd1234...'`<br><br>**Applying Opportunities**<br>Partners can take action against opportunities based on their business needs. For example, partners may choose to use existing Advertising API resources to launch new ASINs or modify campaign settings. Some entries in your opportunity data file will have `recommendationId`s and `applyApiEndPoint`s. These opportunities will have objective type `AMAZON_ACCOUNT_TEAM_RECOMMENDATIONS` and can be supplied back to the `/apply` endpoint to programmatically take action on each opportunity.<br><br>See [How to use the Partner Opportunities API](https://advertising.amazon.com/API/docs/en-us/guides/recommendations/partner-opportunities/how-to) for additional details. <br><br>**Advanced Usage**<br><br>**Pagination**<br><br>The `/partnerOpportunities` route supports pagination of results via query parameters. <br><br>**GET** `/partnerOpportunities` calls support using *optional* query parameters `maxResults` and `nextToken` for paginated requests and responses. <br>**GET** `/partnerOpportunities` responses include tokens which can be used to navigate to the first, previous, next, and last pages of results. To navigate to the desired page, pass one of these provided tokens as a query parameter for the next call to **GET** `/partnerOpportunities`.<br><br>Examples:<br>- **GET** `/partnerOpportunities?maxResults=10` will return 10 opportunities and a valid series of pagination tokens.<br>- **GET** `/partnerOpportunities?maxResults=10&nextToken=[next-page-token]` will get the next 10 opportunities. <br><br>**Filtering**<br><br>The `/partnerOpportunities` and `/partnerOpportunities/summary` routes support filtering results via query parameters.<br><br>**GET** `/partnerOpportunities` calls support using *optional* query parameters `audience`, `objectiveType`, `product`, `advertiserId`, and `profileId` for filtering responses.<br>**GET** `/partnerOpportunities` responses will include only opportunities which have values matching the requested filters. If no filters are specified, all opportunities are returned.<br><br>Examples:<br>- **GET** `/partnerOpportunities?objectiveType=UNLAUNCHED_ASINS&product=AMAZON_DSP,SPONSORED_PRODUCTS` will return all opportunities which have a `objectiveType` of `UNLAUNCHED_ASINS` and a `product` of `AMAZON_DSP` OR `SPONSORED_PRODUCTS`.<br>- **GET** `/partnerOpportunities?advertiserId=A12345,B67890` will return all opportunities that contain data for `advertiserId` `A12345` OR `B67890`.<br>- **GET** `/partnerOpportunities?profileId=12345,67890` will return all opportunities that contain data for `profileId` `12345` OR `67890`.<br>- **GET** `/partnerOpportunities?advertiserId=A12345,B67890&profileId=12345,67890` will return all opportunities that contain data for `advertiserId` `A12345` OR `B67890` OR `profileId` `12345` OR `67890`.<br><br>**GET** `/partnerOpportunities/summary` calls support using *optional* query parameters `audience`, `objectiveType`, and `product` for filtering responses.<br>**GET** `/partnerOpportunities/summary` responses will include metadata for the filter values for opportunities under the `availableAudiences`, `availableObjectiveTypes`, etc. fields. This metadata will include the available filter values along with the number of opportunities for those filter values.<br><br>Examples:<br>- **GET** `/partnerOpportunities/summary?product=SPONSORED_PRODUCTS` will provide a summary of all opportunities that have a `product` value of `SPONSORED_PRODUCTS`.<br><br>**Localization**<br><br>The `/partnerOpportunities` route supports a `locale` query parameter for retrieving opportunities in a localized manner.<br><br>**GET** `/partnerOpportunities` calls support the *optional* query parameter `locale` for localizing responses.<br>**GET** `/partnerOpportunities` responses will be localized to the requested `locale` if possible and filtered out if not.<br><br>Example:<br>- **GET** `/partnerOpportunities?locale=zh_CN` to request opportunity responses localized in Chinese.<br><br>**Additional Resources**<br><br>For more information on **CURL** command formatting, please see [the Amazon Ads Website](https://advertising.amazon.com/API/docs/en-us/getting-started/first-call).<br>
 *
 * API version: 3.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package partneropportunities

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type PartnerOpportunitiesApiService service
/*
PartnerOpportunitiesApiService
Retrieves the current status of applied recommendations.  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\&quot;MasterAccount_Manager\&quot;,\&quot;MasterAccount_Viewer\&quot;,\&quot;ManagerAccount_Dev\&quot;,\&quot;MA_ReadOnly\&quot;]
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param amazonAdvertisingAPIClientId The identifier of a client associated with a &#x27;Login with Amazon&#x27; account.
 * @param amazonAdvertisingAPIManagerAccount &#x27;Partner Network Account ID&#x27; which is accessible from Partner Network under the [&#x27;User settings&#x27;](https://advertising.amazon.com/partner-network/settings) link in the upper right corner.
 * @param partnerOpportunityId
@return PartnerOpportunitiesApplicationStatusResponseDtoV1
*/
func (a *PartnerOpportunitiesApiService) PartnerOpportunitiesApplicationStatus(ctx context.Context, body PartnerOpportunitiesApplicationStatusRequestDtoV1, amazonAdvertisingAPIClientId string, amazonAdvertisingAPIManagerAccount string, partnerOpportunityId string) (PartnerOpportunitiesApplicationStatusResponseDtoV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PartnerOpportunitiesApplicationStatusResponseDtoV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/partnerOpportunities/{partnerOpportunityId}/applicationStatus"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerOpportunityId"+"}", fmt.Sprintf("%v", partnerOpportunityId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/vnd.partneropportunity.v1+json", "application/vnd.partneropportunity.v1.1+json", "application/vnd.partneropportunity.v1.2+json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/vnd.partneropportunity.v1+json", "application/vnd.partneropportunity.v1.1+json", "application/vnd.partneropportunity.v1.2+json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Amazon-Advertising-API-ClientId"] = parameterToString(amazonAdvertisingAPIClientId, "")
	localVarHeaderParams["Amazon-Advertising-API-Manager-Account"] = parameterToString(amazonAdvertisingAPIManagerAccount, "")
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v PartnerOpportunitiesApplicationStatusResponseDtoV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v PartnerOpportunitiesApplicationStatusErrorDtoV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v PartnerOpportunitiesApplicationStatusErrorDtoV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
PartnerOpportunitiesApiService
Applies a given set of recommendations. Application may be asynchronous. Application status may be checked using the applicationStatus operation. Note that all required parameters are retrieved from opportunity data.  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\&quot;MasterAccount_Manager\&quot;,\&quot;MasterAccount_Viewer\&quot;,\&quot;ManagerAccount_Dev\&quot;,\&quot;MA_ReadOnly\&quot;]
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param body
 * @param amazonAdvertisingAPIClientId The identifier of a client associated with a &#x27;Login with Amazon&#x27; account.
 * @param amazonAdvertisingAPIManagerAccount &#x27;Partner Network Account ID&#x27; which is accessible from Partner Network under the [&#x27;User settings&#x27;](https://advertising.amazon.com/partner-network/settings) link in the upper right corner.
 * @param partnerOpportunityId

*/
func (a *PartnerOpportunitiesApiService) PartnerOpportunitiesApply(ctx context.Context, body PartnerOpportunitiesApplyRequestDtoV1, amazonAdvertisingAPIClientId string, amazonAdvertisingAPIManagerAccount string, partnerOpportunityId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/partnerOpportunities/{partnerOpportunityId}/apply"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerOpportunityId"+"}", fmt.Sprintf("%v", partnerOpportunityId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/vnd.partneropportunity.v1+json", "application/vnd.partneropportunity.v1.1+json", "application/vnd.partneropportunity.v1.2+json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/vnd.partneropportunity.v1+json", "application/vnd.partneropportunity.v1.1+json", "application/vnd.partneropportunity.v1.2+json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Amazon-Advertising-API-ClientId"] = parameterToString(amazonAdvertisingAPIClientId, "")
	localVarHeaderParams["Amazon-Advertising-API-Manager-Account"] = parameterToString(amazonAdvertisingAPIManagerAccount, "")
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v PartnerOpportunitiesApplyErrorDtoV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 422 {
			var v PartnerOpportunitiesApplyErrorDtoV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 500 {
			var v PartnerOpportunitiesApplyErrorDtoV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarHttpResponse, newErr
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
PartnerOpportunitiesApiService
Gets a 307 - TEMPORARY_REDIRECT to an opportunity data file.  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\&quot;MasterAccount_Manager\&quot;,\&quot;MasterAccount_Viewer\&quot;,\&quot;ManagerAccount_Dev\&quot;,\&quot;MA_ReadOnly\&quot;]
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param partnerOpportunityId The opportunity ID for which the file URL is desired.
 * @param amazonAdvertisingAPIClientId The identifier of a client associated with a &#x27;Login with Amazon&#x27; account.
 * @param amazonAdvertisingAPIManagerAccount &#x27;Partner Network Account ID&#x27; which is accessible from Partner Network under the [&#x27;User settings&#x27;](https://advertising.amazon.com/partner-network/settings) link in the upper right corner.
 * @param optional nil or *PartnerOpportunitiesApiPartnerOpportunitiesGetOpportunityFileOpts - Optional Parameters:
     * @param "FileFormat" (optional.String) -  Specify the desired file format for the opportunity data. Default is CSV.

*/

type PartnerOpportunitiesApiPartnerOpportunitiesGetOpportunityFileOpts struct {
    FileFormat optional.String
}

func (a *PartnerOpportunitiesApiService) PartnerOpportunitiesGetOpportunityFile(ctx context.Context, partnerOpportunityId string, amazonAdvertisingAPIClientId string, amazonAdvertisingAPIManagerAccount string, localVarOptionals *PartnerOpportunitiesApiPartnerOpportunitiesGetOpportunityFileOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/partnerOpportunities/{partnerOpportunityId}/file"
	localVarPath = strings.Replace(localVarPath, "{"+"partnerOpportunityId"+"}", fmt.Sprintf("%v", partnerOpportunityId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.FileFormat.IsSet() {
		localVarQueryParams.Add("fileFormat", parameterToString(localVarOptionals.FileFormat.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Amazon-Advertising-API-ClientId"] = parameterToString(amazonAdvertisingAPIClientId, "")
	localVarHeaderParams["Amazon-Advertising-API-Manager-Account"] = parameterToString(amazonAdvertisingAPIManagerAccount, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}


	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
/*
PartnerOpportunitiesApiService
Gets a list of opportunities specific to the partner making the request.  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\&quot;MasterAccount_Manager\&quot;,\&quot;MasterAccount_Viewer\&quot;,\&quot;ManagerAccount_Dev\&quot;,\&quot;MA_ReadOnly\&quot;]
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param amazonAdvertisingAPIClientId The identifier of a client associated with a &#x27;Login with Amazon&#x27; account.
 * @param amazonAdvertisingAPIManagerAccount &#x27;Partner Network Account ID&#x27; which is accessible from Partner Network under the [&#x27;User settings&#x27;](https://advertising.amazon.com/partner-network/settings) link in the upper right corner.
 * @param optional nil or *PartnerOpportunitiesApiPartnerOpportunitiesListOpportunitiesOpts - Optional Parameters:
     * @param "MaxResults" (optional.Float64) -  The maximum number of results to return in a single page.
     * @param "NextToken" (optional.String) -  An obfuscated cursor value that indicates which &#x27;page&#x27; of results should be returned next.
     * @param "Locale" (optional.Interface of Locale) -  The desired locale for opportunity data to be presented in. The &#x60;title&#x60;, &#x60;description&#x60;, and &#x60;callToAction&#x60; fields of the response items support localization.
     * @param "RetrieveTranslationKeys" (optional.Bool) - 
     * @param "AdvertiserId" (optional.Interface of []string) -  Filter for opportunities using advertiserId. Each advertiserId must be fewer than 50 characters.
     * @param "ProfileId" (optional.Interface of []float64) -  Filter for opportunities using profileId. Each profileId must be fewer than 50 characters.
     * @param "Audience" (optional.Interface of []string) -  Filter for opportunities with these audience values. * PARTNER_MANAGED_ADVERTISERS - Recommendation relates to advertisers the partner manages. * PARTNER_MANAGED_AD_BUSINESS - Recommendation relates to other partners the partner interacts with. * PARTNER - Recommendation relates to you, the partner.
     * @param "ObjectiveType" (optional.Interface of []string) -  Filter for opportunities with these objectiveType values. * AD_API_ENDPOINT_ADOPTION - Recommendation relates to adopting a new API endpoint. * ADVERTISER_INSIGHTS - Recommendation relates to advertiser insights. * AMAZON_ACCOUNT_TEAM_RECOMMENDATIONS - Recommendation is provided by the Amazon Ads Account Team. * BENCHMARKING_INSIGHTS - Recommendation relates to performance insights comparing against similar partners, brands, campaigns, or ASINs. * CAMPAIGN_OPTIMIZATION - Recommendation relates to optimizing campaigns. * CATEGORY_INSIGHTS - Recommendation relates to advertising insights across product categories.. * CLICK_CREDITS - Recommendation relates to available click credits. * DEALS - Recommendation relates to deals. * MARKETPLACE_EXPANSION - Recommendation relates to expanding to new marketplaces. * NEW_TO_BRAND_INSIGHTS - Recommendation relates to new to brand advertising insights. * PARTNER_GROWTH - Recommendation relates to growing your business as a partner. * PATH_TO_PURCHASE_INSIGHTS - Recommendation relates to path to purchase insights. * READY_TO_LAUNCH_CAMPAIGNS - Recommendation relates to ready to launch campaigns. * RETAIL_INSIGHTS - Recommendation related to retail insights about products you manage. * SHARE_OF_VOICE_INSIGHTS - Recommendation relates to share of voice for a particular audience. * UNLAUNCHED_ASINS - Recommendation relates to ASINs you manage that are not enrolled in advertising campaigns. 
     * @param "Product" (optional.Interface of []string) -  Filter for opportunities with these product values. * AMAZON_DSP - Recommendation relates to the Amazon DSP. * AMAZON_LIVE - Recommendation relates to Amazon&#x27;s Live Show and Tell program. * CROSS_PRODUCT - Recommendation relates to multiple Amazon Ads products including leveraging insights between products * POSTS - Deprecated * SPONSORED_BRANDS - Recommendation relates to Sponsored Brands. * SPONSORED_BRANDS_VIDEO - Recommendation relates to Sponsored Brands Video. * SPONSORED_DISPLAY - Recommendation relates to Sponsored Display. * SPONSORED_DISPLAY_VIDEO - Recommendation relates to Sponsored Display Video. * SPONSORED_PRODUCTS - Recommendation relates to Sponsored Products. * SPONSORED_TV - Recommendation relates to Sponsored TV. * STORES - Recommendation relates to building a storefront page on Amazon. * VIDEO_ADS - Deprecated value, replaced by SPONSORED_BRANDS_VIDEO and SPONSORED_DISPLAY_VIDEO values.
@return PartnerOpportunitiesOpportunitiesPageV1
*/

type PartnerOpportunitiesApiPartnerOpportunitiesListOpportunitiesOpts struct {
    MaxResults optional.Float64
    NextToken optional.String
    Locale optional.Interface
    RetrieveTranslationKeys optional.Bool
    AdvertiserId optional.Interface
    ProfileId optional.Interface
    Audience optional.Interface
    ObjectiveType optional.Interface
    Product optional.Interface
}

func (a *PartnerOpportunitiesApiService) PartnerOpportunitiesListOpportunities(ctx context.Context, amazonAdvertisingAPIClientId string, amazonAdvertisingAPIManagerAccount string, localVarOptionals *PartnerOpportunitiesApiPartnerOpportunitiesListOpportunitiesOpts) (PartnerOpportunitiesOpportunitiesPageV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PartnerOpportunitiesOpportunitiesPageV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/partnerOpportunities"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.MaxResults.IsSet() {
		localVarQueryParams.Add("maxResults", parameterToString(localVarOptionals.MaxResults.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.NextToken.IsSet() {
		localVarQueryParams.Add("nextToken", parameterToString(localVarOptionals.NextToken.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Locale.IsSet() {
		localVarQueryParams.Add("locale", parameterToString(localVarOptionals.Locale.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RetrieveTranslationKeys.IsSet() {
		localVarQueryParams.Add("retrieveTranslationKeys", parameterToString(localVarOptionals.RetrieveTranslationKeys.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AdvertiserId.IsSet() {
		localVarQueryParams.Add("advertiserId", parameterToString(localVarOptionals.AdvertiserId.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.ProfileId.IsSet() {
		localVarQueryParams.Add("profileId", parameterToString(localVarOptionals.ProfileId.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Audience.IsSet() {
		localVarQueryParams.Add("audience", parameterToString(localVarOptionals.Audience.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.ObjectiveType.IsSet() {
		localVarQueryParams.Add("objectiveType", parameterToString(localVarOptionals.ObjectiveType.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Product.IsSet() {
		localVarQueryParams.Add("product", parameterToString(localVarOptionals.Product.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/vnd.partneropportunity.v1+json", "application/vnd.partneropportunity.v1.1+json", "application/vnd.partneropportunity.v1.2+json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Amazon-Advertising-API-ClientId"] = parameterToString(amazonAdvertisingAPIClientId, "")
	localVarHeaderParams["Amazon-Advertising-API-Manager-Account"] = parameterToString(amazonAdvertisingAPIManagerAccount, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v PartnerOpportunitiesOpportunitiesPageV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
/*
PartnerOpportunitiesApiService
Gets aggregated information about all opportunities specific to the partner making the request. Supported since V1.1.  **Authorized resource type**: Global Manager Account ID  **Parameter name**: Amazon-Advertising-API-Manager-Account  **Parameter in**: header  **Requires one of these permissions**: [\&quot;MasterAccount_Manager\&quot;,\&quot;MasterAccount_Viewer\&quot;,\&quot;ManagerAccount_Dev\&quot;,\&quot;MA_ReadOnly\&quot;]
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param amazonAdvertisingAPIClientId The identifier of a client associated with a &#x27;Login with Amazon&#x27; account.
 * @param amazonAdvertisingAPIManagerAccount &#x27;Partner Network Account ID&#x27; which is accessible from Partner Network under the [&#x27;User settings&#x27;](https://advertising.amazon.com/partner-network/settings) link in the upper right corner.
 * @param optional nil or *PartnerOpportunitiesApiPartnerOpportunitiesSummarizeOpportunitiesOpts - Optional Parameters:
     * @param "Audience" (optional.Interface of []string) -  Filter for opportunities with these audience values. * PARTNER_MANAGED_ADVERTISERS - Recommendation relates to advertisers the partner manages. * PARTNER_MANAGED_AD_BUSINESS - Recommendation relates to other partners the partner interacts with. * PARTNER - Recommendation relates to you, the partner.
     * @param "ObjectiveType" (optional.Interface of []string) -  Filter for opportunities with these objectiveType values. * AD_API_ENDPOINT_ADOPTION - Recommendation relates to adopting a new API endpoint. * ADVERTISER_INSIGHTS - Recommendation relates to advertiser insights. * AMAZON_ACCOUNT_TEAM_RECOMMENDATIONS - Recommendation is provided by the Amazon Ads Account Team. * BENCHMARKING_INSIGHTS - Recommendation relates to performance insights comparing against similar partners, brands, campaigns, or ASINs. * CAMPAIGN_OPTIMIZATION - Recommendation relates to optimizing campaigns. * CATEGORY_INSIGHTS - Recommendation relates to advertising insights across product categories.. * CLICK_CREDITS - Recommendation relates to available click credits. * DEALS - Recommendation relates to deals. * MARKETPLACE_EXPANSION - Recommendation relates to expanding to new marketplaces. * NEW_TO_BRAND_INSIGHTS - Recommendation relates to new to brand advertising insights. * PARTNER_GROWTH - Recommendation relates to growing your business as a partner. * PATH_TO_PURCHASE_INSIGHTS - Recommendation relates to path to purchase insights. * READY_TO_LAUNCH_CAMPAIGNS - Recommendation relates to ready to launch campaigns. * RETAIL_INSIGHTS - Recommendation related to retail insights about products you manage. * SHARE_OF_VOICE_INSIGHTS - Recommendation relates to share of voice for a particular audience. * UNLAUNCHED_ASINS - Recommendation relates to ASINs you manage that are not enrolled in advertising campaigns. 
     * @param "Product" (optional.Interface of []string) -  Filter for opportunities with these product values. * AMAZON_DSP - Recommendation relates to the Amazon DSP. * AMAZON_LIVE - Recommendation relates to Amazon&#x27;s Live Show and Tell program. * CROSS_PRODUCT - Recommendation relates to multiple Amazon Ads products including leveraging insights between products * POSTS - Deprecated * SPONSORED_BRANDS - Recommendation relates to Sponsored Brands. * SPONSORED_BRANDS_VIDEO - Recommendation relates to Sponsored Brands Video. * SPONSORED_DISPLAY - Recommendation relates to Sponsored Display. * SPONSORED_DISPLAY_VIDEO - Recommendation relates to Sponsored Display Video. * SPONSORED_PRODUCTS - Recommendation relates to Sponsored Products. * SPONSORED_TV - Recommendation relates to Sponsored TV. * STORES - Recommendation relates to building a storefront page on Amazon. * VIDEO_ADS - Deprecated value, replaced by SPONSORED_BRANDS_VIDEO and SPONSORED_DISPLAY_VIDEO values.
@return PartnerOpportunitiesOpportunitiesSummaryV1
*/

type PartnerOpportunitiesApiPartnerOpportunitiesSummarizeOpportunitiesOpts struct {
    Audience optional.Interface
    ObjectiveType optional.Interface
    Product optional.Interface
}

func (a *PartnerOpportunitiesApiService) PartnerOpportunitiesSummarizeOpportunities(ctx context.Context, amazonAdvertisingAPIClientId string, amazonAdvertisingAPIManagerAccount string, localVarOptionals *PartnerOpportunitiesApiPartnerOpportunitiesSummarizeOpportunitiesOpts) (PartnerOpportunitiesOpportunitiesSummaryV1, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		localVarReturnValue PartnerOpportunitiesOpportunitiesSummaryV1
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/partnerOpportunities/summary"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Audience.IsSet() {
		localVarQueryParams.Add("audience", parameterToString(localVarOptionals.Audience.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.ObjectiveType.IsSet() {
		localVarQueryParams.Add("objectiveType", parameterToString(localVarOptionals.ObjectiveType.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Product.IsSet() {
		localVarQueryParams.Add("product", parameterToString(localVarOptionals.Product.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/vnd.partneropportunity.v1+json", "application/vnd.partneropportunity.v1.1+json", "application/vnd.partneropportunity.v1.2+json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Amazon-Advertising-API-ClientId"] = parameterToString(amazonAdvertisingAPIClientId, "")
	localVarHeaderParams["Amazon-Advertising-API-Manager-Account"] = parameterToString(amazonAdvertisingAPIManagerAccount, "")
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil { 
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v PartnerOpportunitiesOpportunitiesSummaryV1
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
