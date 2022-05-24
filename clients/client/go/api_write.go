/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * API version: v0.0.1-alpha.181
 * Contact: support@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

type WriteApi interface {

	/*
	 * CreateRelationTuple Create a Relation Tuple
	 * Use this endpoint to create a relation tuple.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return WriteApiApiCreateRelationTupleRequest
	 */
	CreateRelationTuple(ctx context.Context) WriteApiApiCreateRelationTupleRequest

	/*
	 * CreateRelationTupleExecute executes the request
	 * @return RelationQuery
	 */
	CreateRelationTupleExecute(r WriteApiApiCreateRelationTupleRequest) (*RelationQuery, *http.Response, error)

	/*
	 * DeleteRelationTuples Delete Relation Tuples
	 * Use this endpoint to delete relation tuples
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return WriteApiApiDeleteRelationTuplesRequest
	 */
	DeleteRelationTuples(ctx context.Context) WriteApiApiDeleteRelationTuplesRequest

	/*
	 * DeleteRelationTuplesExecute executes the request
	 */
	DeleteRelationTuplesExecute(r WriteApiApiDeleteRelationTuplesRequest) (*http.Response, error)

	/*
	 * PatchRelationTuples Patch Multiple Relation Tuples
	 * Use this endpoint to patch one or more relation tuples.
	 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	 * @return WriteApiApiPatchRelationTuplesRequest
	 */
	PatchRelationTuples(ctx context.Context) WriteApiApiPatchRelationTuplesRequest

	/*
	 * PatchRelationTuplesExecute executes the request
	 */
	PatchRelationTuplesExecute(r WriteApiApiPatchRelationTuplesRequest) (*http.Response, error)
}

// WriteApiService WriteApi service
type WriteApiService service

type WriteApiApiCreateRelationTupleRequest struct {
	ctx context.Context
	ApiService WriteApi
	relationQuery *RelationQuery
}

func (r WriteApiApiCreateRelationTupleRequest) RelationQuery(relationQuery RelationQuery) WriteApiApiCreateRelationTupleRequest {
	r.relationQuery = &relationQuery
	return r
}

func (r WriteApiApiCreateRelationTupleRequest) Execute() (*RelationQuery, *http.Response, error) {
	return r.ApiService.CreateRelationTupleExecute(r)
}

/*
 * CreateRelationTuple Create a Relation Tuple
 * Use this endpoint to create a relation tuple.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return WriteApiApiCreateRelationTupleRequest
 */
func (a *WriteApiService) CreateRelationTuple(ctx context.Context) WriteApiApiCreateRelationTupleRequest {
	return WriteApiApiCreateRelationTupleRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return RelationQuery
 */
func (a *WriteApiService) CreateRelationTupleExecute(r WriteApiApiCreateRelationTupleRequest) (*RelationQuery, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  *RelationQuery
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WriteApiService.CreateRelationTuple")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.relationQuery
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericError
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

type WriteApiApiDeleteRelationTuplesRequest struct {
	ctx context.Context
	ApiService WriteApi
	namespace *string
	object *string
	relation *string
	subjectId *string
	subjectSetNamespace *string
	subjectSetObject *string
	subjectSetRelation *string
}

func (r WriteApiApiDeleteRelationTuplesRequest) Namespace(namespace string) WriteApiApiDeleteRelationTuplesRequest {
	r.namespace = &namespace
	return r
}
func (r WriteApiApiDeleteRelationTuplesRequest) Object(object string) WriteApiApiDeleteRelationTuplesRequest {
	r.object = &object
	return r
}
func (r WriteApiApiDeleteRelationTuplesRequest) Relation(relation string) WriteApiApiDeleteRelationTuplesRequest {
	r.relation = &relation
	return r
}
func (r WriteApiApiDeleteRelationTuplesRequest) SubjectId(subjectId string) WriteApiApiDeleteRelationTuplesRequest {
	r.subjectId = &subjectId
	return r
}
func (r WriteApiApiDeleteRelationTuplesRequest) SubjectSetNamespace(subjectSetNamespace string) WriteApiApiDeleteRelationTuplesRequest {
	r.subjectSetNamespace = &subjectSetNamespace
	return r
}
func (r WriteApiApiDeleteRelationTuplesRequest) SubjectSetObject(subjectSetObject string) WriteApiApiDeleteRelationTuplesRequest {
	r.subjectSetObject = &subjectSetObject
	return r
}
func (r WriteApiApiDeleteRelationTuplesRequest) SubjectSetRelation(subjectSetRelation string) WriteApiApiDeleteRelationTuplesRequest {
	r.subjectSetRelation = &subjectSetRelation
	return r
}

func (r WriteApiApiDeleteRelationTuplesRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteRelationTuplesExecute(r)
}

/*
 * DeleteRelationTuples Delete Relation Tuples
 * Use this endpoint to delete relation tuples
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return WriteApiApiDeleteRelationTuplesRequest
 */
func (a *WriteApiService) DeleteRelationTuples(ctx context.Context) WriteApiApiDeleteRelationTuplesRequest {
	return WriteApiApiDeleteRelationTuplesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *WriteApiService) DeleteRelationTuplesExecute(r WriteApiApiDeleteRelationTuplesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WriteApiService.DeleteRelationTuples")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.namespace != nil {
		localVarQueryParams.Add("namespace", parameterToString(*r.namespace, ""))
	}
	if r.object != nil {
		localVarQueryParams.Add("object", parameterToString(*r.object, ""))
	}
	if r.relation != nil {
		localVarQueryParams.Add("relation", parameterToString(*r.relation, ""))
	}
	if r.subjectId != nil {
		localVarQueryParams.Add("subject_id", parameterToString(*r.subjectId, ""))
	}
	if r.subjectSetNamespace != nil {
		localVarQueryParams.Add("subject_set.namespace", parameterToString(*r.subjectSetNamespace, ""))
	}
	if r.subjectSetObject != nil {
		localVarQueryParams.Add("subject_set.object", parameterToString(*r.subjectSetObject, ""))
	}
	if r.subjectSetRelation != nil {
		localVarQueryParams.Add("subject_set.relation", parameterToString(*r.subjectSetRelation, ""))
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericError
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

type WriteApiApiPatchRelationTuplesRequest struct {
	ctx context.Context
	ApiService WriteApi
	patchDelta *[]PatchDelta
}

func (r WriteApiApiPatchRelationTuplesRequest) PatchDelta(patchDelta []PatchDelta) WriteApiApiPatchRelationTuplesRequest {
	r.patchDelta = &patchDelta
	return r
}

func (r WriteApiApiPatchRelationTuplesRequest) Execute() (*http.Response, error) {
	return r.ApiService.PatchRelationTuplesExecute(r)
}

/*
 * PatchRelationTuples Patch Multiple Relation Tuples
 * Use this endpoint to patch one or more relation tuples.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return WriteApiApiPatchRelationTuplesRequest
 */
func (a *WriteApiService) PatchRelationTuples(ctx context.Context) WriteApiApiPatchRelationTuplesRequest {
	return WriteApiApiPatchRelationTuplesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 */
func (a *WriteApiService) PatchRelationTuplesExecute(r WriteApiApiPatchRelationTuplesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WriteApiService.PatchRelationTuples")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/relation-tuples"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.patchDelta
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v GenericError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v GenericError
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
