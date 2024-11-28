# \OrdersItemsIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1OrdersItemsId**](OrdersItemsIdAPI.md#GetV1OrdersItemsId) | **Get** /V1/orders/items/{id} | orders/items/{id}



## GetV1OrdersItemsId

> SalesDataOrderItemInterface GetV1OrdersItemsId(ctx, id).Execute()

orders/items/{id}



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
	id := int32(56) // int32 | The order item ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersItemsIdAPI.GetV1OrdersItemsId(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersItemsIdAPI.GetV1OrdersItemsId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1OrdersItemsId`: SalesDataOrderItemInterface
	fmt.Fprintf(os.Stdout, "Response from `OrdersItemsIdAPI.GetV1OrdersItemsId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The order item ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1OrdersItemsIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalesDataOrderItemInterface**](SalesDataOrderItemInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

