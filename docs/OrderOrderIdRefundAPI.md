# \OrderOrderIdRefundAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1OrderOrderIdRefund**](OrderOrderIdRefundAPI.md#PostV1OrderOrderIdRefund) | **Post** /V1/order/{orderId}/refund | order/{orderId}/refund



## PostV1OrderOrderIdRefund

> int32 PostV1OrderOrderIdRefund(ctx, orderId).PostV1OrderOrderIdRefundRequest(postV1OrderOrderIdRefundRequest).Execute()

order/{orderId}/refund



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
	postV1OrderOrderIdRefundRequest := *openapiclient.NewPostV1OrderOrderIdRefundRequest() // PostV1OrderOrderIdRefundRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderOrderIdRefundAPI.PostV1OrderOrderIdRefund(context.Background(), orderId).PostV1OrderOrderIdRefundRequest(postV1OrderOrderIdRefundRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderOrderIdRefundAPI.PostV1OrderOrderIdRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1OrderOrderIdRefund`: int32
	fmt.Fprintf(os.Stdout, "Response from `OrderOrderIdRefundAPI.PostV1OrderOrderIdRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1OrderOrderIdRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1OrderOrderIdRefundRequest** | [**PostV1OrderOrderIdRefundRequest**](PostV1OrderOrderIdRefundRequest.md) |  | 

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

