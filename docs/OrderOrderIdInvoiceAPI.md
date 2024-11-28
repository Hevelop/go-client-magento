# \OrderOrderIdInvoiceAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1OrderOrderIdInvoice**](OrderOrderIdInvoiceAPI.md#PostV1OrderOrderIdInvoice) | **Post** /V1/order/{orderId}/invoice | order/{orderId}/invoice



## PostV1OrderOrderIdInvoice

> int32 PostV1OrderOrderIdInvoice(ctx, orderId).PostV1OrderOrderIdInvoiceRequest(postV1OrderOrderIdInvoiceRequest).Execute()

order/{orderId}/invoice



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
	orderId := int32(56) // int32 | 
	postV1OrderOrderIdInvoiceRequest := *openapiclient.NewPostV1OrderOrderIdInvoiceRequest() // PostV1OrderOrderIdInvoiceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderOrderIdInvoiceAPI.PostV1OrderOrderIdInvoice(context.Background(), orderId).PostV1OrderOrderIdInvoiceRequest(postV1OrderOrderIdInvoiceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderOrderIdInvoiceAPI.PostV1OrderOrderIdInvoice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1OrderOrderIdInvoice`: int32
	fmt.Fprintf(os.Stdout, "Response from `OrderOrderIdInvoiceAPI.PostV1OrderOrderIdInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1OrderOrderIdInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1OrderOrderIdInvoiceRequest** | [**PostV1OrderOrderIdInvoiceRequest**](PostV1OrderOrderIdInvoiceRequest.md) |  | 

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

