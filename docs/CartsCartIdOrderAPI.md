# \CartsCartIdOrderAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CartsCartIdOrder**](CartsCartIdOrderAPI.md#PutV1CartsCartIdOrder) | **Put** /V1/carts/{cartId}/order | carts/{cartId}/order



## PutV1CartsCartIdOrder

> int32 PutV1CartsCartIdOrder(ctx, cartId).PutV1CartsMineOrderRequest(putV1CartsMineOrderRequest).Execute()

carts/{cartId}/order



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
	cartId := int32(56) // int32 | The cart ID.
	putV1CartsMineOrderRequest := *openapiclient.NewPutV1CartsMineOrderRequest() // PutV1CartsMineOrderRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdOrderAPI.PutV1CartsCartIdOrder(context.Background(), cartId).PutV1CartsMineOrderRequest(putV1CartsMineOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdOrderAPI.PutV1CartsCartIdOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CartsCartIdOrder`: int32
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdOrderAPI.PutV1CartsCartIdOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CartsCartIdOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1CartsMineOrderRequest** | [**PutV1CartsMineOrderRequest**](PutV1CartsMineOrderRequest.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

