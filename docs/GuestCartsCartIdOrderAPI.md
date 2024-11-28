# \GuestCartsCartIdOrderAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1GuestcartsCartIdOrder**](GuestCartsCartIdOrderAPI.md#PutV1GuestcartsCartIdOrder) | **Put** /V1/guest-carts/{cartId}/order | guest-carts/{cartId}/order



## PutV1GuestcartsCartIdOrder

> int32 PutV1GuestcartsCartIdOrder(ctx, cartId).PutV1CartsMineOrderRequest(putV1CartsMineOrderRequest).Execute()

guest-carts/{cartId}/order



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
	putV1CartsMineOrderRequest := *openapiclient.NewPutV1CartsMineOrderRequest() // PutV1CartsMineOrderRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdOrderAPI.PutV1GuestcartsCartIdOrder(context.Background(), cartId).PutV1CartsMineOrderRequest(putV1CartsMineOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdOrderAPI.PutV1GuestcartsCartIdOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1GuestcartsCartIdOrder`: int32
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdOrderAPI.PutV1GuestcartsCartIdOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1GuestcartsCartIdOrderRequest struct via the builder pattern


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

