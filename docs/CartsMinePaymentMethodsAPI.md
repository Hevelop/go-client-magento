# \CartsMinePaymentMethodsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsMinePaymentmethods**](CartsMinePaymentMethodsAPI.md#GetV1CartsMinePaymentmethods) | **Get** /V1/carts/mine/payment-methods | carts/mine/payment-methods



## GetV1CartsMinePaymentmethods

> []QuoteDataPaymentMethodInterface GetV1CartsMinePaymentmethods(ctx).Execute()

carts/mine/payment-methods



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMinePaymentMethodsAPI.GetV1CartsMinePaymentmethods(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMinePaymentMethodsAPI.GetV1CartsMinePaymentmethods``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsMinePaymentmethods`: []QuoteDataPaymentMethodInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMinePaymentMethodsAPI.GetV1CartsMinePaymentmethods`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsMinePaymentmethodsRequest struct via the builder pattern


### Return type

[**[]QuoteDataPaymentMethodInterface**](QuoteDataPaymentMethodInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

