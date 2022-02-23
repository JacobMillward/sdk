/*
 * ORY Oathkeeper
 *
 * ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.
 *
 * API version: v0.38.22-beta.1
 * Contact: hi@ory.am
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type ApiApi interface {

	/*
	 * Decisions Access Control Decision API
	 * > This endpoint works with all HTTP Methods (GET, POST, PUT, ...) and matches every path prefixed with /decision.

This endpoint mirrors the proxy capability of ORY Oathkeeper's proxy functionality but instead of forwarding the
request to the upstream server, returns 200 (request should be allowed), 401 (unauthorized), or 403 (forbidden)
status codes. This endpoint can be used to integrate with other API Proxies like Ambassador, Kong, Envoy, and many more.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiApiApiDecisionsRequest
	 */
	Decisions(ctx context.Context) ApiApiApiDecisionsRequest

	/*
	 * DecisionsExecute executes the request
	 */
	DecisionsExecute(r ApiApiApiDecisionsRequest) (*http.Response, error)

	/*
	 * GetRule Retrieve a rule
	 * Use this method to retrieve a rule from the storage. If it does not exist you will receive a 404 error.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @param id
	 * @return ApiApiApiGetRuleRequest
	 */
	GetRule(ctx context.Context, id string) ApiApiApiGetRuleRequest

	/*
	 * GetRuleExecute executes the request
	 * @return Rule
	 */
	GetRuleExecute(r ApiApiApiGetRuleRequest) (*Rule, *http.Response, error)

	/*
	 * GetVersion Get service version
	 * This endpoint returns the service version typically notated using semantic versioning.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of this service, the health status will never
refer to the cluster state, only to a single instance.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiApiApiGetVersionRequest
	 */
	GetVersion(ctx context.Context) ApiApiApiGetVersionRequest

	/*
	 * GetVersionExecute executes the request
	 * @return Version
	 */
	GetVersionExecute(r ApiApiApiGetVersionRequest) (*Version, *http.Response, error)

	/*
	 * GetWellKnownJSONWebKeys Lists cryptographic keys
	 * This endpoint returns cryptographic keys that are required to, for example, verify signatures of ID Tokens.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiApiApiGetWellKnownJSONWebKeysRequest
	 */
	GetWellKnownJSONWebKeys(ctx context.Context) ApiApiApiGetWellKnownJSONWebKeysRequest

	/*
	 * GetWellKnownJSONWebKeysExecute executes the request
	 * @return JsonWebKeySet
	 */
	GetWellKnownJSONWebKeysExecute(r ApiApiApiGetWellKnownJSONWebKeysRequest) (*JsonWebKeySet, *http.Response, error)

	/*
	 * IsInstanceAlive Check alive status
	 * This endpoint returns a 200 status code when the HTTP server is up running.
This status does currently not include checks whether the database connection is working.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of this service, the health status will never
refer to the cluster state, only to a single instance.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiApiApiIsInstanceAliveRequest
	 */
	IsInstanceAlive(ctx context.Context) ApiApiApiIsInstanceAliveRequest

	/*
	 * IsInstanceAliveExecute executes the request
	 * @return HealthStatus
	 */
	IsInstanceAliveExecute(r ApiApiApiIsInstanceAliveRequest) (*HealthStatus, *http.Response, error)

	/*
	 * IsInstanceReady Check readiness status
	 * This endpoint returns a 200 status code when the HTTP server is up running and the environment dependencies (e.g.
the database) are responsive as well.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of this service, the health status will never
refer to the cluster state, only to a single instance.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiApiApiIsInstanceReadyRequest
	 */
	IsInstanceReady(ctx context.Context) ApiApiApiIsInstanceReadyRequest

	/*
	 * IsInstanceReadyExecute executes the request
	 * @return HealthStatus
	 */
	IsInstanceReadyExecute(r ApiApiApiIsInstanceReadyRequest) (*HealthStatus, *http.Response, error)

	/*
	 * ListRules List all rules
	 * This method returns an array of all rules that are stored in the backend. This is useful if you want to get a full
view of what rules you have currently in place.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return ApiApiApiListRulesRequest
	 */
	ListRules(ctx context.Context) ApiApiApiListRulesRequest

	/*
	 * ListRulesExecute executes the request
	 * @return []Rule
	 */
	ListRulesExecute(r ApiApiApiListRulesRequest) ([]Rule, *http.Response, error)
}

// ApiApiService ApiApi service
type ApiApiService service

type ApiApiApiDecisionsRequest struct {
	ctx context.Context
	ApiService ApiApi
}


func (r ApiApiApiDecisionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DecisionsExecute(r)
}

/*
 * Decisions Access Control Decision API
 * > This endpoint works with all HTTP Methods (GET, POST, PUT, ...) and matches every path prefixed with /decision.

This endpoint mirrors the proxy capability of ORY Oathkeeper's proxy functionality but instead of forwarding the
request to the upstream server, returns 200 (request should be allowed), 401 (unauthorized), or 403 (forbidden)
status codes. This endpoint can be used to integrate with other API Proxies like Ambassador, Kong, Envoy, and many more.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiApiApiDecisionsRequest
 */
func (a *ApiApiService) Decisions(ctx context.Context) ApiApiApiDecisionsRequest {
	return ApiApiApiDecisionsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *ApiApiService) DecisionsExecute(r ApiApiApiDecisionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiApiService.Decisions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/decisions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v InlineResponse500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v InlineResponse500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InlineResponse500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiApiApiGetRuleRequest struct {
	ctx context.Context
	ApiService ApiApi
	id string
}


func (r ApiApiApiGetRuleRequest) Execute() (*Rule, *http.Response, error) {
	return r.ApiService.GetRuleExecute(r)
}

/*
 * GetRule Retrieve a rule
 * Use this method to retrieve a rule from the storage. If it does not exist you will receive a 404 error.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param id
 * @return ApiApiApiGetRuleRequest
 */
func (a *ApiApiService) GetRule(ctx context.Context, id string) ApiApiApiGetRuleRequest {
	return ApiApiApiGetRuleRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

/*
 * Execute executes the request
 * @return Rule
 */
func (a *ApiApiService) GetRuleExecute(r ApiApiApiGetRuleRequest) (*Rule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *Rule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiApiService.GetRule")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rules/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterToString(r.id, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v InlineResponse500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InlineResponse500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiApiGetVersionRequest struct {
	ctx context.Context
	ApiService ApiApi
}


func (r ApiApiApiGetVersionRequest) Execute() (*Version, *http.Response, error) {
	return r.ApiService.GetVersionExecute(r)
}

/*
 * GetVersion Get service version
 * This endpoint returns the service version typically notated using semantic versioning.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of this service, the health status will never
refer to the cluster state, only to a single instance.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiApiApiGetVersionRequest
 */
func (a *ApiApiService) GetVersion(ctx context.Context) ApiApiApiGetVersionRequest {
	return ApiApiApiGetVersionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return Version
 */
func (a *ApiApiService) GetVersionExecute(r ApiApiApiGetVersionRequest) (*Version, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *Version
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiApiService.GetVersion")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/version"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiApiGetWellKnownJSONWebKeysRequest struct {
	ctx context.Context
	ApiService ApiApi
}


func (r ApiApiApiGetWellKnownJSONWebKeysRequest) Execute() (*JsonWebKeySet, *http.Response, error) {
	return r.ApiService.GetWellKnownJSONWebKeysExecute(r)
}

/*
 * GetWellKnownJSONWebKeys Lists cryptographic keys
 * This endpoint returns cryptographic keys that are required to, for example, verify signatures of ID Tokens.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiApiApiGetWellKnownJSONWebKeysRequest
 */
func (a *ApiApiService) GetWellKnownJSONWebKeys(ctx context.Context) ApiApiApiGetWellKnownJSONWebKeysRequest {
	return ApiApiApiGetWellKnownJSONWebKeysRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return JsonWebKeySet
 */
func (a *ApiApiService) GetWellKnownJSONWebKeysExecute(r ApiApiApiGetWellKnownJSONWebKeysRequest) (*JsonWebKeySet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *JsonWebKeySet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiApiService.GetWellKnownJSONWebKeys")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/.well-known/jwks.json"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InlineResponse500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiApiIsInstanceAliveRequest struct {
	ctx context.Context
	ApiService ApiApi
}


func (r ApiApiApiIsInstanceAliveRequest) Execute() (*HealthStatus, *http.Response, error) {
	return r.ApiService.IsInstanceAliveExecute(r)
}

/*
 * IsInstanceAlive Check alive status
 * This endpoint returns a 200 status code when the HTTP server is up running.
This status does currently not include checks whether the database connection is working.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of this service, the health status will never
refer to the cluster state, only to a single instance.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiApiApiIsInstanceAliveRequest
 */
func (a *ApiApiService) IsInstanceAlive(ctx context.Context) ApiApiApiIsInstanceAliveRequest {
	return ApiApiApiIsInstanceAliveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return HealthStatus
 */
func (a *ApiApiService) IsInstanceAliveExecute(r ApiApiApiIsInstanceAliveRequest) (*HealthStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *HealthStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiApiService.IsInstanceAlive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/health/alive"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InlineResponse500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiApiIsInstanceReadyRequest struct {
	ctx context.Context
	ApiService ApiApi
}


func (r ApiApiApiIsInstanceReadyRequest) Execute() (*HealthStatus, *http.Response, error) {
	return r.ApiService.IsInstanceReadyExecute(r)
}

/*
 * IsInstanceReady Check readiness status
 * This endpoint returns a 200 status code when the HTTP server is up running and the environment dependencies (e.g.
the database) are responsive as well.

If the service supports TLS Edge Termination, this endpoint does not require the
`X-Forwarded-Proto` header to be set.

Be aware that if you are running multiple nodes of this service, the health status will never
refer to the cluster state, only to a single instance.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiApiApiIsInstanceReadyRequest
 */
func (a *ApiApiService) IsInstanceReady(ctx context.Context) ApiApiApiIsInstanceReadyRequest {
	return ApiApiApiIsInstanceReadyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return HealthStatus
 */
func (a *ApiApiService) IsInstanceReadyExecute(r ApiApiApiIsInstanceReadyRequest) (*HealthStatus, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *HealthStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiApiService.IsInstanceReady")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/health/ready"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v HealthNotReadyStatus
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiApiApiListRulesRequest struct {
	ctx context.Context
	ApiService ApiApi
	limit *int64
	offset *int64
}

func (r ApiApiApiListRulesRequest) Limit(limit int64) ApiApiApiListRulesRequest {
	r.limit = &limit
	return r
}
func (r ApiApiApiListRulesRequest) Offset(offset int64) ApiApiApiListRulesRequest {
	r.offset = &offset
	return r
}

func (r ApiApiApiListRulesRequest) Execute() ([]Rule, *http.Response, error) {
	return r.ApiService.ListRulesExecute(r)
}

/*
 * ListRules List all rules
 * This method returns an array of all rules that are stored in the backend. This is useful if you want to get a full
view of what rules you have currently in place.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiApiApiListRulesRequest
 */
func (a *ApiApiService) ListRules(ctx context.Context) ApiApiApiListRulesRequest {
	return ApiApiApiListRulesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return []Rule
 */
func (a *ApiApiService) ListRulesExecute(r ApiApiApiListRulesRequest) ([]Rule, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []Rule
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ApiApiService.ListRules")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/rules"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.offset != nil {
		localVarQueryParams.Add("offset", parameterToString(*r.offset, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v InlineResponse500
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
