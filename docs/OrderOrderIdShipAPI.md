# \OrderOrderIdShipAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1OrderOrderIdShip**](OrderOrderIdShipAPI.md#PostV1OrderOrderIdShip) | **Post** /V1/order/{orderId}/ship | order/{orderId}/ship



## PostV1OrderOrderIdShip

> int32 PostV1OrderOrderIdShip(ctx, orderId).PostV1OrderOrderIdShipRequest(postV1OrderOrderIdShipRequest).Execute()

order/{orderId}/ship



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
	postV1OrderOrderIdShipRequest := *openapiclient.NewPostV1OrderOrderIdShipRequest() // PostV1OrderOrderIdShipRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrderOrderIdShipAPI.PostV1OrderOrderIdShip(context.Background(), orderId).PostV1OrderOrderIdShipRequest(postV1OrderOrderIdShipRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrderOrderIdShipAPI.PostV1OrderOrderIdShip``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1OrderOrderIdShip`: int32
	fmt.Fprintf(os.Stdout, "Response from `OrderOrderIdShipAPI.PostV1OrderOrderIdShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**orderId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1OrderOrderIdShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1OrderOrderIdShipRequest** | [**PostV1OrderOrderIdShipRequest**](PostV1OrderOrderIdShipRequest.md) |  | 

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

