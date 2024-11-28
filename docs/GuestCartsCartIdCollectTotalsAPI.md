# \GuestCartsCartIdCollectTotalsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1GuestcartsCartIdCollecttotals**](GuestCartsCartIdCollectTotalsAPI.md#PutV1GuestcartsCartIdCollecttotals) | **Put** /V1/guest-carts/{cartId}/collect-totals | guest-carts/{cartId}/collect-totals



## PutV1GuestcartsCartIdCollecttotals

> QuoteDataTotalsInterface PutV1GuestcartsCartIdCollecttotals(ctx, cartId).PutV1CartsMineCollecttotalsRequest(putV1CartsMineCollecttotalsRequest).Execute()

guest-carts/{cartId}/collect-totals



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
	cartId := "cartId_example" // string | The cart ID.
	putV1CartsMineCollecttotalsRequest := *openapiclient.NewPutV1CartsMineCollecttotalsRequest(*openapiclient.NewQuoteDataPaymentInterface("Method_example")) // PutV1CartsMineCollecttotalsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdCollectTotalsAPI.PutV1GuestcartsCartIdCollecttotals(context.Background(), cartId).PutV1CartsMineCollecttotalsRequest(putV1CartsMineCollecttotalsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdCollectTotalsAPI.PutV1GuestcartsCartIdCollecttotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1GuestcartsCartIdCollecttotals`: QuoteDataTotalsInterface
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdCollectTotalsAPI.PutV1GuestcartsCartIdCollecttotals`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1GuestcartsCartIdCollecttotalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1CartsMineCollecttotalsRequest** | [**PutV1CartsMineCollecttotalsRequest**](PutV1CartsMineCollecttotalsRequest.md) |  | 

### Return type

[**QuoteDataTotalsInterface**](QuoteDataTotalsInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

