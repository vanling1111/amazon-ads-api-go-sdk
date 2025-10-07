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
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"golang.org/x/oauth2"
)

var (
	jsonCheck = regexp.MustCompile("(?i:[application|text]/json)")
	xmlCheck  = regexp.MustCompile("(?i:[application|text]/xml)")
)

// APIClient manages communication with the Partner Opportunities API v3.0
// In most cases there should be only one, shared, APIClient.
type APIClient struct {
	cfg    *Configuration
	common service // Reuse a single struct instead of allocating one for each service on the heap.

	// API Services

	PartnerOpportunitiesApi *PartnerOpportunitiesApiService
}

type service struct {
	client *APIClient
}

// NewAPIClient creates a new API client. Requires a userAgent string describing your application.
// optionally a custom http.Client to allow for advanced features such as caching.
func NewAPIClient(cfg *Configuration) *APIClient {
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = http.DefaultClient
	}

	c := &APIClient{}
	c.cfg = cfg
	c.common.client = c

	// API Services
	c.PartnerOpportunitiesApi = (*PartnerOpportunitiesApiService)(&c.common)

	return c
}

func atoi(in string) (int, error) {
	return strconv.Atoi(in)
}

// selectHeaderContentType select a content type from the available list.
func selectHeaderContentType(contentTypes []string) string {
	if len(contentTypes) == 0 {
		return ""
	}
	if contains(contentTypes, "application/json") {
		return "application/json"
	}
	return contentTypes[0] // use the first content type specified in 'consumes'
}

// selectHeaderAccept join all accept types and return
func selectHeaderAccept(accepts []string) string {
	if len(accepts) == 0 {
		return ""
	}

	if contains(accepts, "application/json") {
		return "application/json"
	}

	return strings.Join(accepts, ",")
}

// contains is a case insenstive match, finding needle in a haystack
func contains(haystack []string, needle string) bool {
	for _, a := range haystack {
		if strings.ToLower(a) == strings.ToLower(needle) {
			return true
		}
	}
	return false
}

// Verify optional parameters are of the correct type.
func typeCheckParameter(obj interface{}, expected string, name string) error {
	// Make sure there is an object.
	if obj == nil {
		return nil
	}

	// Check the type is as expected.
	if reflect.TypeOf(obj).String() != expected {
		return fmt.Errorf("Expected %s to be of type %s but received %s.", name, expected, reflect.TypeOf(obj).String())
	}
	return nil
}

// parameterToString convert interface{} parameters to string, using a delimiter if format is provided.
func parameterToString(obj interface{}, collectionFormat string) string {
	var delimiter string

	switch collectionFormat {
	case "pipes":
		delimiter = "|"
	case "ssv":
		delimiter = " "
	case "tsv":
		delimiter = "\t"
	case "csv":
		delimiter = ","
	}

	if reflect.TypeOf(obj).Kind() == reflect.Slice {
		return strings.Trim(strings.Replace(fmt.Sprint(obj), " ", delimiter, -1), "[]")
	}

	return fmt.Sprintf("%v", obj)
}

// callAPI do the request.
func (c *APIClient) callAPI(request *http.Request) (*http.Response, error) {
	return c.cfg.HTTPClient.Do(request)
}

// Change base path to allow switching to mocks
func (c *APIClient) ChangeBasePath(path string) {
	c.cfg.BasePath = path
}

// prepareRequest build the request
func (c *APIClient) prepareRequest(
	ctx context.Context,
	path string, method string,
	postBody interface{},
	headerParams map[string]string,
	queryParams url.Values,
	formParams url.Values,
	fileName string,
	fileBytes []byte) (localVarRequest *http.Request, err error) {

	var body *bytes.Buffer

	// Detect postBody type and post.
	if postBody != nil {
		contentType := headerParams["Content-Type"]
		if contentType == "" {
			contentType = detectContentType(postBody)
			headerParams["Content-Type"] = contentType
		}

		body, err = setBody(postBody, contentType)
		if err != nil {
			return nil, err
		}
	}

	// add form parameters and file if available.
	if strings.HasPrefix(headerParams["Content-Type"], "multipart/form-data") && len(formParams) > 0 || (len(fileBytes) > 0 && fileName != "") {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and multipart form at the same time.")
		}
		body = &bytes.Buffer{}
		w := multipart.NewWriter(body)

		for k, v := range formParams {
			for _, iv := range v {
				if strings.HasPrefix(k, "@") { // file
					err = addFile(w, k[1:], iv)
					if err != nil {
						return nil, err
					}
				} else { // form value
					w.WriteField(k, iv)
				}
			}
		}
		if len(fileBytes) > 0 && fileName != "" {
			w.Boundary()
			//_, fileNm := filepath.Split(fileName)
			part, err := w.CreateFormFile("file", filepath.Base(fileName))
			if err != nil {
				return nil, err
			}
			_, err = part.Write(fileBytes)
			if err != nil {
				return nil, err
			}
			// Set the Boundary in the Content-Type
			headerParams["Content-Type"] = w.FormDataContentType()
		}

		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
		w.Close()
	}

	if strings.HasPrefix(headerParams["Content-Type"], "application/x-www-form-urlencoded") && len(formParams) > 0 {
		if body != nil {
			return nil, errors.New("Cannot specify postBody and x-www-form-urlencoded form at the same time.")
		}
		body = &bytes.Buffer{}
		body.WriteString(formParams.Encode())
		// Set Content-Length
		headerParams["Content-Length"] = fmt.Sprintf("%d", body.Len())
	}

	// Setup path and query parameters
	url, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	// Adding Query Param
	query := url.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	// Encode the parameters.
	url.RawQuery = query.Encode()

	// Generate a new request
	if body != nil {
		localVarRequest, err = http.NewRequest(method, url.String(), body)
	} else {
		localVarRequest, err = http.NewRequest(method, url.String(), nil)
	}
	if err != nil {
		return nil, err
	}

	// add header parameters, if any
	if len(headerParams) > 0 {
		headers := http.Header{}
		for h, v := range headerParams {
			headers.Set(h, v)
		}
		localVarRequest.Header = headers
	}

	// Override request host, if applicable
	if c.cfg.Host != "" {
		localVarRequest.Host = c.cfg.Host
	}

	// Add the user agent to the request.
	localVarRequest.Header.Add("User-Agent", c.cfg.UserAgent)

	if ctx != nil {
		// add context to the request
		localVarRequest = localVarRequest.WithContext(ctx)

		// Walk through any authentication.

		// OAuth2 authentication
		if tok, ok := ctx.Value(ContextOAuth2).(oauth2.TokenSource); ok {
			// We were able to grab an oauth2 token from the context
			var latestToken *oauth2.Token
			if latestToken, err = tok.Token(); err != nil {
				return nil, err
			}

			latestToken.SetAuthHeader(localVarRequest)
		}

		// Basic HTTP Authentication
		if auth, ok := ctx.Value(ContextBasicAuth).(BasicAuth); ok {
			localVarRequest.SetBasicAuth(auth.UserName, auth.Password)
		}

		// AccessToken Authentication
		if auth, ok := ctx.Value(ContextAccessToken).(string); ok {
			localVarRequest.Header.Add("Authorization", "Bearer "+auth)
		}
	}

	for header, value := range c.cfg.DefaultHeader {
		localVarRequest.Header.Add(header, value)
	}

	return localVarRequest, nil
}

func (c *APIClient) decode(v interface{}, b []byte, contentType string) (err error) {
		if strings.Contains(contentType, "application/xml") {
			if err = xml.Unmarshal(b, v); err != nil {
				return err
			}
			return nil
		} else if strings.Contains(contentType, "application/json") {
			if err = json.Unmarshal(b, v); err != nil {
				return err
			}
			return nil
		}
	return errors.New("undefined response type")
}

// Add a file to the multipart request
func addFile(w *multipart.Writer, fieldName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	part, err := w.CreateFormFile(fieldName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	return err
}

// Prevent trying to import "fmt"
func reportError(format string, a ...interface{}) error {
	return fmt.Errorf(format, a...)
}

// Set request body from an interface{}
func setBody(body interface{}, contentType string) (bodyBuf *bytes.Buffer, err error) {
	if bodyBuf == nil {
		bodyBuf = &bytes.Buffer{}
	}

	if reader, ok := body.(io.Reader); ok {
		_, err = bodyBuf.ReadFrom(reader)
	} else if b, ok := body.([]byte); ok {
		_, err = bodyBuf.Write(b)
	} else if s, ok := body.(string); ok {
		_, err = bodyBuf.WriteString(s)
	} else if s, ok := body.(*string); ok {
		_, err = bodyBuf.WriteString(*s)
	} else if jsonCheck.MatchString(contentType) {
		err = json.NewEncoder(bodyBuf).Encode(body)
	} else if xmlCheck.MatchString(contentType) {
		xml.NewEncoder(bodyBuf).Encode(body)
	}

	if err != nil {
		return nil, err
	}

	if bodyBuf.Len() == 0 {
		err = fmt.Errorf("Invalid body type %s\n", contentType)
		return nil, err
	}
	return bodyBuf, nil
}

// detectContentType method is used to figure out `Request.Body` content type for request header
func detectContentType(body interface{}) string {
	contentType := "text/plain; charset=utf-8"
	kind := reflect.TypeOf(body).Kind()

	switch kind {
	case reflect.Struct, reflect.Map, reflect.Ptr:
		contentType = "application/json; charset=utf-8"
	case reflect.String:
		contentType = "text/plain; charset=utf-8"
	default:
		if b, ok := body.([]byte); ok {
			contentType = http.DetectContentType(b)
		} else if kind == reflect.Slice {
			contentType = "application/json; charset=utf-8"
		}
	}

	return contentType
}

// Ripped from https://github.com/gregjones/httpcache/blob/master/httpcache.go
type cacheControl map[string]string

func parseCacheControl(headers http.Header) cacheControl {
	cc := cacheControl{}
	ccHeader := headers.Get("Cache-Control")
	for _, part := range strings.Split(ccHeader, ",") {
		part = strings.Trim(part, " ")
		if part == "" {
			continue
		}
		if strings.ContainsRune(part, '=') {
			keyval := strings.Split(part, "=")
			cc[strings.Trim(keyval[0], " ")] = strings.Trim(keyval[1], ",")
		} else {
			cc[part] = ""
		}
	}
	return cc
}

// CacheExpires helper function to determine remaining time before repeating a request.
func CacheExpires(r *http.Response) time.Time {
	// Figure out when the cache expires.
	var expires time.Time
	now, err := time.Parse(time.RFC1123, r.Header.Get("date"))
	if err != nil {
		return time.Now()
	}
	respCacheControl := parseCacheControl(r.Header)

	if maxAge, ok := respCacheControl["max-age"]; ok {
		lifetime, err := time.ParseDuration(maxAge + "s")
		if err != nil {
			expires = now
		}
		expires = now.Add(lifetime)
	} else {
		expiresHeader := r.Header.Get("Expires")
		if expiresHeader != "" {
			expires, err = time.Parse(time.RFC1123, expiresHeader)
			if err != nil {
				expires = now
			}
		}
	}
	return expires
}

func strlen(s string) int {
	return utf8.RuneCountInString(s)
}

// GenericSwaggerError Provides access to the body, error and model on returned errors.
type GenericSwaggerError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e GenericSwaggerError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e GenericSwaggerError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e GenericSwaggerError) Model() interface{} {
	return e.model
}
