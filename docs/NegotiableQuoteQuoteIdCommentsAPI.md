# \NegotiableQuoteQuoteIdCommentsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1NegotiableQuoteQuoteIdComments**](NegotiableQuoteQuoteIdCommentsAPI.md#GetV1NegotiableQuoteQuoteIdComments) | **Get** /V1/negotiableQuote/{quoteId}/comments | negotiableQuote/{quoteId}/comments



## GetV1NegotiableQuoteQuoteIdComments

> []NegotiableQuoteDataCommentInterface GetV1NegotiableQuoteQuoteIdComments(ctx, quoteId).Execute()

negotiableQuote/{quoteId}/comments



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/Hevelop/go-client-magento"
)

func main() {
	quoteId := int32(56) // int32 | Negotiable Quote ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableQuoteQuoteIdCommentsAPI.GetV1NegotiableQuoteQuoteIdComments(context.Background(), quoteId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableQuoteQuoteIdCommentsAPI.GetV1NegotiableQuoteQuoteIdComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1NegotiableQuoteQuoteIdComments`: []NegotiableQuoteDataCommentInterface
	fmt.Fprintf(os.Stdout, "Response from `NegotiableQuoteQuoteIdCommentsAPI.GetV1NegotiableQuoteQuoteIdComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **int32** | Negotiable Quote ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1NegotiableQuoteQuoteIdCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]NegotiableQuoteDataCommentInterface**](NegotiableQuoteDataCommentInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

