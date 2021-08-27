/*
Cisco Intersight

Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-08-10T21:48:06Z.

API version: 1.0.9-4430
Contact: intersight@cisco.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
)

// Linger please
var (
	_ _context.Context
)

// TelemetryApiService TelemetryApi service
type TelemetryApiService service

type ApiQueryTelemetryDatasourceMetadataRequest struct {
	ctx                                     _context.Context
	ApiService                              *TelemetryApiService
	telemetryDruidDataSourceMetadataRequest *TelemetryDruidDataSourceMetadataRequest
}

// The Druid request schema for time series queries.
func (r ApiQueryTelemetryDatasourceMetadataRequest) TelemetryDruidDataSourceMetadataRequest(telemetryDruidDataSourceMetadataRequest TelemetryDruidDataSourceMetadataRequest) ApiQueryTelemetryDatasourceMetadataRequest {
	r.telemetryDruidDataSourceMetadataRequest = &telemetryDruidDataSourceMetadataRequest
	return r
}

func (r ApiQueryTelemetryDatasourceMetadataRequest) Execute() ([]TelemetryDruidDataSourceMetadataResult, *_nethttp.Response, error) {
	return r.ApiService.QueryTelemetryDatasourceMetadataExecute(r)
}

/*
QueryTelemetryDatasourceMetadata Perform a Druid DatasourceMetadata request.

Endpoint that exposes Druid DatasourceMetadata requests for time series data.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryTelemetryDatasourceMetadataRequest
*/
func (a *TelemetryApiService) QueryTelemetryDatasourceMetadata(ctx _context.Context) ApiQueryTelemetryDatasourceMetadataRequest {
	return ApiQueryTelemetryDatasourceMetadataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []TelemetryDruidDataSourceMetadataResult
func (a *TelemetryApiService) QueryTelemetryDatasourceMetadataExecute(r ApiQueryTelemetryDatasourceMetadataRequest) ([]TelemetryDruidDataSourceMetadataResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TelemetryDruidDataSourceMetadataResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelemetryApiService.QueryTelemetryDatasourceMetadata")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telemetry/DataSourceMetadata"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.telemetryDruidDataSourceMetadataRequest == nil {
		return localVarReturnValue, nil, reportError("telemetryDruidDataSourceMetadataRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.telemetryDruidDataSourceMetadataRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v TelemetryDruidError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiQueryTelemetryGroupByRequest struct {
	ctx                          _context.Context
	ApiService                   *TelemetryApiService
	telemetryDruidGroupByRequest *TelemetryDruidGroupByRequest
}

// The Druid request schema for time series queries.
func (r ApiQueryTelemetryGroupByRequest) TelemetryDruidGroupByRequest(telemetryDruidGroupByRequest TelemetryDruidGroupByRequest) ApiQueryTelemetryGroupByRequest {
	r.telemetryDruidGroupByRequest = &telemetryDruidGroupByRequest
	return r
}

func (r ApiQueryTelemetryGroupByRequest) Execute() ([]TelemetryDruidGroupByResult, *_nethttp.Response, error) {
	return r.ApiService.QueryTelemetryGroupByExecute(r)
}

/*
QueryTelemetryGroupBy Perform a Druid GroupBy request.

Endpoint that exposes Druid GroupBy requests for time series data.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryTelemetryGroupByRequest
*/
func (a *TelemetryApiService) QueryTelemetryGroupBy(ctx _context.Context) ApiQueryTelemetryGroupByRequest {
	return ApiQueryTelemetryGroupByRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []TelemetryDruidGroupByResult
func (a *TelemetryApiService) QueryTelemetryGroupByExecute(r ApiQueryTelemetryGroupByRequest) ([]TelemetryDruidGroupByResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TelemetryDruidGroupByResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelemetryApiService.QueryTelemetryGroupBy")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telemetry/GroupBys"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.telemetryDruidGroupByRequest == nil {
		return localVarReturnValue, nil, reportError("telemetryDruidGroupByRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.telemetryDruidGroupByRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v TelemetryDruidError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiQueryTelemetryScanRequest struct {
	ctx                       _context.Context
	ApiService                *TelemetryApiService
	telemetryDruidScanRequest *TelemetryDruidScanRequest
}

// The Druid request schema for time series queries.
func (r ApiQueryTelemetryScanRequest) TelemetryDruidScanRequest(telemetryDruidScanRequest TelemetryDruidScanRequest) ApiQueryTelemetryScanRequest {
	r.telemetryDruidScanRequest = &telemetryDruidScanRequest
	return r
}

func (r ApiQueryTelemetryScanRequest) Execute() ([]TelemetryDruidScanResult, *_nethttp.Response, error) {
	return r.ApiService.QueryTelemetryScanExecute(r)
}

/*
QueryTelemetryScan Perform a Druid Scan request.

Endpoint that exposes Druid Scan requests for time series data.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryTelemetryScanRequest
*/
func (a *TelemetryApiService) QueryTelemetryScan(ctx _context.Context) ApiQueryTelemetryScanRequest {
	return ApiQueryTelemetryScanRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []TelemetryDruidScanResult
func (a *TelemetryApiService) QueryTelemetryScanExecute(r ApiQueryTelemetryScanRequest) ([]TelemetryDruidScanResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TelemetryDruidScanResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelemetryApiService.QueryTelemetryScan")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telemetry/Scans"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.telemetryDruidScanRequest == nil {
		return localVarReturnValue, nil, reportError("telemetryDruidScanRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.telemetryDruidScanRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v TelemetryDruidError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiQueryTelemetrySearchRequest struct {
	ctx                         _context.Context
	ApiService                  *TelemetryApiService
	telemetryDruidSearchRequest *TelemetryDruidSearchRequest
}

// The Druid request schema for time series queries.
func (r ApiQueryTelemetrySearchRequest) TelemetryDruidSearchRequest(telemetryDruidSearchRequest TelemetryDruidSearchRequest) ApiQueryTelemetrySearchRequest {
	r.telemetryDruidSearchRequest = &telemetryDruidSearchRequest
	return r
}

func (r ApiQueryTelemetrySearchRequest) Execute() ([]TelemetryDruidSearchResult, *_nethttp.Response, error) {
	return r.ApiService.QueryTelemetrySearchExecute(r)
}

/*
QueryTelemetrySearch Perform a Druid Search request.

Endpoint that exposes Druid Search requests for time series data.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryTelemetrySearchRequest
*/
func (a *TelemetryApiService) QueryTelemetrySearch(ctx _context.Context) ApiQueryTelemetrySearchRequest {
	return ApiQueryTelemetrySearchRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []TelemetryDruidSearchResult
func (a *TelemetryApiService) QueryTelemetrySearchExecute(r ApiQueryTelemetrySearchRequest) ([]TelemetryDruidSearchResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TelemetryDruidSearchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelemetryApiService.QueryTelemetrySearch")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telemetry/Searches"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.telemetryDruidSearchRequest == nil {
		return localVarReturnValue, nil, reportError("telemetryDruidSearchRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.telemetryDruidSearchRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v TelemetryDruidError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiQueryTelemetrySegmentMetadataRequest struct {
	ctx                                  _context.Context
	ApiService                           *TelemetryApiService
	telemetryDruidSegmentMetadataRequest *TelemetryDruidSegmentMetadataRequest
}

// The Druid request schema for time series queries.
func (r ApiQueryTelemetrySegmentMetadataRequest) TelemetryDruidSegmentMetadataRequest(telemetryDruidSegmentMetadataRequest TelemetryDruidSegmentMetadataRequest) ApiQueryTelemetrySegmentMetadataRequest {
	r.telemetryDruidSegmentMetadataRequest = &telemetryDruidSegmentMetadataRequest
	return r
}

func (r ApiQueryTelemetrySegmentMetadataRequest) Execute() ([]TelemetryDruidSegmentMetadataResult, *_nethttp.Response, error) {
	return r.ApiService.QueryTelemetrySegmentMetadataExecute(r)
}

/*
QueryTelemetrySegmentMetadata Perform a Druid SegmentMetadata request.

Endpoint that exposes Druid SegmentMetadata requests for time series data.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryTelemetrySegmentMetadataRequest
*/
func (a *TelemetryApiService) QueryTelemetrySegmentMetadata(ctx _context.Context) ApiQueryTelemetrySegmentMetadataRequest {
	return ApiQueryTelemetrySegmentMetadataRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []TelemetryDruidSegmentMetadataResult
func (a *TelemetryApiService) QueryTelemetrySegmentMetadataExecute(r ApiQueryTelemetrySegmentMetadataRequest) ([]TelemetryDruidSegmentMetadataResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TelemetryDruidSegmentMetadataResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelemetryApiService.QueryTelemetrySegmentMetadata")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telemetry/SegmentMetadata"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.telemetryDruidSegmentMetadataRequest == nil {
		return localVarReturnValue, nil, reportError("telemetryDruidSegmentMetadataRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.telemetryDruidSegmentMetadataRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v TelemetryDruidError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiQueryTelemetryTimeBoundaryRequest struct {
	ctx                               _context.Context
	ApiService                        *TelemetryApiService
	telemetryDruidTimeBoundaryRequest *TelemetryDruidTimeBoundaryRequest
}

// The Druid request schema for time series queries.
func (r ApiQueryTelemetryTimeBoundaryRequest) TelemetryDruidTimeBoundaryRequest(telemetryDruidTimeBoundaryRequest TelemetryDruidTimeBoundaryRequest) ApiQueryTelemetryTimeBoundaryRequest {
	r.telemetryDruidTimeBoundaryRequest = &telemetryDruidTimeBoundaryRequest
	return r
}

func (r ApiQueryTelemetryTimeBoundaryRequest) Execute() ([]TelemetryDruidTimeBoundaryResult, *_nethttp.Response, error) {
	return r.ApiService.QueryTelemetryTimeBoundaryExecute(r)
}

/*
QueryTelemetryTimeBoundary Perform a Druid TimeBoundary request.

Endpoint that exposes Druid TimeBoundary requests for time series data.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryTelemetryTimeBoundaryRequest
*/
func (a *TelemetryApiService) QueryTelemetryTimeBoundary(ctx _context.Context) ApiQueryTelemetryTimeBoundaryRequest {
	return ApiQueryTelemetryTimeBoundaryRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []TelemetryDruidTimeBoundaryResult
func (a *TelemetryApiService) QueryTelemetryTimeBoundaryExecute(r ApiQueryTelemetryTimeBoundaryRequest) ([]TelemetryDruidTimeBoundaryResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TelemetryDruidTimeBoundaryResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelemetryApiService.QueryTelemetryTimeBoundary")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telemetry/TimeBoundaries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.telemetryDruidTimeBoundaryRequest == nil {
		return localVarReturnValue, nil, reportError("telemetryDruidTimeBoundaryRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.telemetryDruidTimeBoundaryRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v TelemetryDruidError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiQueryTelemetryTimeSeriesRequest struct {
	ctx                             _context.Context
	ApiService                      *TelemetryApiService
	telemetryDruidTimeSeriesRequest *TelemetryDruidTimeSeriesRequest
}

// The Druid request schema for time series queries.
func (r ApiQueryTelemetryTimeSeriesRequest) TelemetryDruidTimeSeriesRequest(telemetryDruidTimeSeriesRequest TelemetryDruidTimeSeriesRequest) ApiQueryTelemetryTimeSeriesRequest {
	r.telemetryDruidTimeSeriesRequest = &telemetryDruidTimeSeriesRequest
	return r
}

func (r ApiQueryTelemetryTimeSeriesRequest) Execute() ([]TelemetryDruidIntervalResult, *_nethttp.Response, error) {
	return r.ApiService.QueryTelemetryTimeSeriesExecute(r)
}

/*
QueryTelemetryTimeSeries Perform a Druid TimeSeries request.

Endpoint that exposes Druid requests for time series data. This endpoint exposes specifically TimeSeries requests and broker information.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryTelemetryTimeSeriesRequest
*/
func (a *TelemetryApiService) QueryTelemetryTimeSeries(ctx _context.Context) ApiQueryTelemetryTimeSeriesRequest {
	return ApiQueryTelemetryTimeSeriesRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []TelemetryDruidIntervalResult
func (a *TelemetryApiService) QueryTelemetryTimeSeriesExecute(r ApiQueryTelemetryTimeSeriesRequest) ([]TelemetryDruidIntervalResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TelemetryDruidIntervalResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelemetryApiService.QueryTelemetryTimeSeries")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telemetry/TimeSeries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.telemetryDruidTimeSeriesRequest == nil {
		return localVarReturnValue, nil, reportError("telemetryDruidTimeSeriesRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.telemetryDruidTimeSeriesRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v TelemetryDruidError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiQueryTelemetryTopNRequest struct {
	ctx                       _context.Context
	ApiService                *TelemetryApiService
	telemetryDruidTopNRequest *TelemetryDruidTopNRequest
}

// The Druid request schema for time series queries.
func (r ApiQueryTelemetryTopNRequest) TelemetryDruidTopNRequest(telemetryDruidTopNRequest TelemetryDruidTopNRequest) ApiQueryTelemetryTopNRequest {
	r.telemetryDruidTopNRequest = &telemetryDruidTopNRequest
	return r
}

func (r ApiQueryTelemetryTopNRequest) Execute() ([]TelemetryDruidTopNResult, *_nethttp.Response, error) {
	return r.ApiService.QueryTelemetryTopNExecute(r)
}

/*
QueryTelemetryTopN Perform a Druid TopN request.

Endpoint that exposes Druid TopN requests for time series data.

 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQueryTelemetryTopNRequest
*/
func (a *TelemetryApiService) QueryTelemetryTopN(ctx _context.Context) ApiQueryTelemetryTopNRequest {
	return ApiQueryTelemetryTopNRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//  @return []TelemetryDruidTopNResult
func (a *TelemetryApiService) QueryTelemetryTopNExecute(r ApiQueryTelemetryTopNRequest) ([]TelemetryDruidTopNResult, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []TelemetryDruidTopNResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TelemetryApiService.QueryTelemetryTopN")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v1/telemetry/Topns"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.telemetryDruidTopNRequest == nil {
		return localVarReturnValue, nil, reportError("telemetryDruidTopNRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.telemetryDruidTopNRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v TelemetryDruidError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
