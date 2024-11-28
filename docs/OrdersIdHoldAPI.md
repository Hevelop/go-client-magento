# \OrdersIdHoldAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1OrdersIdHold**](OrdersIdHoldAPI.md#PostV1OrdersIdHold) | **Post** /V1/orders/{id}/hold | orders/{id}/hold



## PostV1OrdersIdHold

> bool PostV1OrdersIdHold(ctx, id).Execute()

orders/{id}/hold



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
	id := int32(56) // int32 | The order ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersIdHoldAPI.PostV1OrdersIdHold(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersIdHoldAPI.PostV1OrdersIdHold``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1OrdersIdHold`: bool
	fmt.Fprintf(os.Stdout, "Response from `OrdersIdHoldAPI.PostV1OrdersIdHold`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The order ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1OrdersIdHoldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

