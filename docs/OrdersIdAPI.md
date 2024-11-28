# \OrdersIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1OrdersId**](OrdersIdAPI.md#GetV1OrdersId) | **Get** /V1/orders/{id} | orders/{id}



## GetV1OrdersId

> SalesDataOrderInterface GetV1OrdersId(ctx, id).Execute()

orders/{id}



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
	resp, r, err := apiClient.OrdersIdAPI.GetV1OrdersId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersIdAPI.GetV1OrdersId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1OrdersId`: SalesDataOrderInterface
	fmt.Fprintf(os.Stdout, "Response from `OrdersIdAPI.GetV1OrdersId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The order ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1OrdersIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalesDataOrderInterface**](SalesDataOrderInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

