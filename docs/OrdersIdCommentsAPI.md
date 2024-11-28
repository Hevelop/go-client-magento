# \OrdersIdCommentsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1OrdersIdComments**](OrdersIdCommentsAPI.md#GetV1OrdersIdComments) | **Get** /V1/orders/{id}/comments | orders/{id}/comments
[**PostV1OrdersIdComments**](OrdersIdCommentsAPI.md#PostV1OrdersIdComments) | **Post** /V1/orders/{id}/comments | orders/{id}/comments



## GetV1OrdersIdComments

> SalesDataOrderStatusHistorySearchResultInterface GetV1OrdersIdComments(ctx, id).Execute()

orders/{id}/comments



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
	resp, r, err := apiClient.OrdersIdCommentsAPI.GetV1OrdersIdComments(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersIdCommentsAPI.GetV1OrdersIdComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1OrdersIdComments`: SalesDataOrderStatusHistorySearchResultInterface
	fmt.Fprintf(os.Stdout, "Response from `OrdersIdCommentsAPI.GetV1OrdersIdComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The order ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1OrdersIdCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalesDataOrderStatusHistorySearchResultInterface**](SalesDataOrderStatusHistorySearchResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1OrdersIdComments

> bool PostV1OrdersIdComments(ctx, id).PostV1OrdersIdCommentsRequest(postV1OrdersIdCommentsRequest).Execute()

orders/{id}/comments



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
	postV1OrdersIdCommentsRequest := *openapiclient.NewPostV1OrdersIdCommentsRequest(*openapiclient.NewSalesDataOrderStatusHistoryInterface("Comment_example", int32(123), int32(123), int32(123))) // PostV1OrdersIdCommentsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersIdCommentsAPI.PostV1OrdersIdComments(context.Background(), id).PostV1OrdersIdCommentsRequest(postV1OrdersIdCommentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersIdCommentsAPI.PostV1OrdersIdComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1OrdersIdComments`: bool
	fmt.Fprintf(os.Stdout, "Response from `OrdersIdCommentsAPI.PostV1OrdersIdComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The order ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1OrdersIdCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1OrdersIdCommentsRequest** | [**PostV1OrdersIdCommentsRequest**](PostV1OrdersIdCommentsRequest.md) |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

