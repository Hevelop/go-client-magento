# \NegotiableQuoteQuoteIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1NegotiableQuoteQuoteId**](NegotiableQuoteQuoteIdAPI.md#PutV1NegotiableQuoteQuoteId) | **Put** /V1/negotiableQuote/{quoteId} | negotiableQuote/{quoteId}



## PutV1NegotiableQuoteQuoteId

> ErrorResponse PutV1NegotiableQuoteQuoteId(ctx, quoteId).PutV1CartsMineRequest(putV1CartsMineRequest).Execute()

negotiableQuote/{quoteId}



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
	quoteId := "quoteId_example" // string | 
	putV1CartsMineRequest := *openapiclient.NewPutV1CartsMineRequest(*openapiclient.NewQuoteDataCartInterface(int32(123), *openapiclient.NewCustomerDataCustomerInterface("Email_example", "Firstname_example", "Lastname_example"), int32(123))) // PutV1CartsMineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableQuoteQuoteIdAPI.PutV1NegotiableQuoteQuoteId(context.Background(), quoteId).PutV1CartsMineRequest(putV1CartsMineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableQuoteQuoteIdAPI.PutV1NegotiableQuoteQuoteId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1NegotiableQuoteQuoteId`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `NegotiableQuoteQuoteIdAPI.PutV1NegotiableQuoteQuoteId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1NegotiableQuoteQuoteIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1CartsMineRequest** | [**PutV1CartsMineRequest**](PutV1CartsMineRequest.md) |  | 

### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

