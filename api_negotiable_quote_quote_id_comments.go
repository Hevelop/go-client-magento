/*
Commerce Admin REST endpoints - All inclusive

The schemas documented here are autogenerated from an instance of Adobe Commerce with B2B. Each schema represents a specific user role (Admin, Customer, and Guest) and determines which endpoints are accessible. Use the version switcher to select an Adobe Commerce version and corresponding API.  You can also <a href=\"https://developer.adobe.com/commerce/webapi/rest/quick-reference/generate-local\" target=\"_blank\">generate a local API reference</a> based on your own Adobe Commerce configuration, which allows you to see API documentation for your specific Adobe Commerce modules, third-party modules, and extension attributes.

API version: 2.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package magento

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type NegotiableQuoteQuoteIdCommentsAPI interface {

	/*
		GetV1NegotiableQuoteQuoteIdComments negotiableQuote/{quoteId}/comments

		Returns comments for a specified negotiable quote.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param quoteId Negotiable Quote ID.
		@return ApiGetV1NegotiableQuoteQuoteIdCommentsRequest
	*/
	GetV1NegotiableQuoteQuoteIdComments(ctx context.Context, quoteId int32) ApiGetV1NegotiableQuoteQuoteIdCommentsRequest

	// GetV1NegotiableQuoteQuoteIdCommentsExecute executes the request
	//  @return []NegotiableQuoteDataCommentInterface
	GetV1NegotiableQuoteQuoteIdCommentsExecute(r ApiGetV1NegotiableQuoteQuoteIdCommentsRequest) ([]NegotiableQuoteDataCommentInterface, *http.Response, error)
}

// NegotiableQuoteQuoteIdCommentsAPIService NegotiableQuoteQuoteIdCommentsAPI service
type NegotiableQuoteQuoteIdCommentsAPIService service

type ApiGetV1NegotiableQuoteQuoteIdCommentsRequest struct {
	ctx        context.Context
	ApiService NegotiableQuoteQuoteIdCommentsAPI
	quoteId    int32
}

func (r ApiGetV1NegotiableQuoteQuoteIdCommentsRequest) Execute() ([]NegotiableQuoteDataCommentInterface, *http.Response, error) {
	return r.ApiService.GetV1NegotiableQuoteQuoteIdCommentsExecute(r)
}

/*
GetV1NegotiableQuoteQuoteIdComments negotiableQuote/{quoteId}/comments

Returns comments for a specified negotiable quote.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param quoteId Negotiable Quote ID.
	@return ApiGetV1NegotiableQuoteQuoteIdCommentsRequest
*/
func (a *NegotiableQuoteQuoteIdCommentsAPIService) GetV1NegotiableQuoteQuoteIdComments(ctx context.Context, quoteId int32) ApiGetV1NegotiableQuoteQuoteIdCommentsRequest {
	return ApiGetV1NegotiableQuoteQuoteIdCommentsRequest{
		ApiService: a,
		ctx:        ctx,
		quoteId:    quoteId,
	}
}

// Execute executes the request
//
//	@return []NegotiableQuoteDataCommentInterface
func (a *NegotiableQuoteQuoteIdCommentsAPIService) GetV1NegotiableQuoteQuoteIdCommentsExecute(r ApiGetV1NegotiableQuoteQuoteIdCommentsRequest) ([]NegotiableQuoteDataCommentInterface, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []NegotiableQuoteDataCommentInterface
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NegotiableQuoteQuoteIdCommentsAPIService.GetV1NegotiableQuoteQuoteIdComments")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/V1/negotiableQuote/{quoteId}/comments"
	localVarPath = strings.Replace(localVarPath, "{"+"quoteId"+"}", url.PathEscape(parameterValueToString(r.quoteId, "quoteId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/xml"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		var v ErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
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
