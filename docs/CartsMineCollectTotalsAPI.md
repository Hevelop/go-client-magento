# \CartsMineCollectTotalsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CartsMineCollecttotals**](CartsMineCollectTotalsAPI.md#PutV1CartsMineCollecttotals) | **Put** /V1/carts/mine/collect-totals | carts/mine/collect-totals



## PutV1CartsMineCollecttotals

> QuoteDataTotalsInterface PutV1CartsMineCollecttotals(ctx).PutV1CartsMineCollecttotalsRequest(putV1CartsMineCollecttotalsRequest).Execute()

carts/mine/collect-totals



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
	putV1CartsMineCollecttotalsRequest := *openapiclient.NewPutV1CartsMineCollecttotalsRequest(*openapiclient.NewQuoteDataPaymentInterface("Method_example")) // PutV1CartsMineCollecttotalsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineCollectTotalsAPI.PutV1CartsMineCollecttotals(context.Background()).PutV1CartsMineCollecttotalsRequest(putV1CartsMineCollecttotalsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineCollectTotalsAPI.PutV1CartsMineCollecttotals``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CartsMineCollecttotals`: QuoteDataTotalsInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMineCollectTotalsAPI.PutV1CartsMineCollecttotals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CartsMineCollecttotalsRequest struct via the builder pattern


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

