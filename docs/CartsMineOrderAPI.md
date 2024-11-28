# \CartsMineOrderAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CartsMineOrder**](CartsMineOrderAPI.md#PutV1CartsMineOrder) | **Put** /V1/carts/mine/order | carts/mine/order



## PutV1CartsMineOrder

> int32 PutV1CartsMineOrder(ctx).PutV1CartsMineOrderRequest(putV1CartsMineOrderRequest).Execute()

carts/mine/order



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
	putV1CartsMineOrderRequest := *openapiclient.NewPutV1CartsMineOrderRequest() // PutV1CartsMineOrderRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineOrderAPI.PutV1CartsMineOrder(context.Background()).PutV1CartsMineOrderRequest(putV1CartsMineOrderRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineOrderAPI.PutV1CartsMineOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CartsMineOrder`: int32
	fmt.Fprintf(os.Stdout, "Response from `CartsMineOrderAPI.PutV1CartsMineOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CartsMineOrderRequest struct via the builder pattern


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

