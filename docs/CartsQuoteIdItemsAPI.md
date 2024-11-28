# \CartsQuoteIdItemsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CartsQuoteIdItems**](CartsQuoteIdItemsAPI.md#PostV1CartsQuoteIdItems) | **Post** /V1/carts/{quoteId}/items | carts/{quoteId}/items



## PostV1CartsQuoteIdItems

> QuoteDataCartItemInterface PostV1CartsQuoteIdItems(ctx, quoteId).PostV1CartsMineItemsRequest(postV1CartsMineItemsRequest).Execute()

carts/{quoteId}/items



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
	postV1CartsMineItemsRequest := *openapiclient.NewPostV1CartsMineItemsRequest(*openapiclient.NewQuoteDataCartItemInterface(float32(123), "QuoteId_example")) // PostV1CartsMineItemsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsQuoteIdItemsAPI.PostV1CartsQuoteIdItems(context.Background(), quoteId).PostV1CartsMineItemsRequest(postV1CartsMineItemsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsQuoteIdItemsAPI.PostV1CartsQuoteIdItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsQuoteIdItems`: QuoteDataCartItemInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsQuoteIdItemsAPI.PostV1CartsQuoteIdItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**quoteId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsQuoteIdItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CartsMineItemsRequest** | [**PostV1CartsMineItemsRequest**](PostV1CartsMineItemsRequest.md) |  | 

### Return type

[**QuoteDataCartItemInterface**](QuoteDataCartItemInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

