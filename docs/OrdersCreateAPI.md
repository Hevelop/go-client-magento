# \OrdersCreateAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1OrdersCreate**](OrdersCreateAPI.md#PutV1OrdersCreate) | **Put** /V1/orders/create | orders/create



## PutV1OrdersCreate

> SalesDataOrderInterface PutV1OrdersCreate(ctx).PostV1OrdersRequest(postV1OrdersRequest).Execute()

orders/create



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
	postV1OrdersRequest := *openapiclient.NewPostV1OrdersRequest(*openapiclient.NewSalesDataOrderInterface(float32(123), "CustomerEmail_example", float32(123), []openapiclient.SalesDataOrderItemInterface{*openapiclient.NewSalesDataOrderItemInterface("Sku_example")})) // PostV1OrdersRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersCreateAPI.PutV1OrdersCreate(context.Background()).PostV1OrdersRequest(postV1OrdersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersCreateAPI.PutV1OrdersCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1OrdersCreate`: SalesDataOrderInterface
	fmt.Fprintf(os.Stdout, "Response from `OrdersCreateAPI.PutV1OrdersCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1OrdersCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1OrdersRequest** | [**PostV1OrdersRequest**](PostV1OrdersRequest.md) |  | 

### Return type

[**SalesDataOrderInterface**](SalesDataOrderInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

