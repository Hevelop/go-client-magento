# \NegotiableCartsCartIdCouponsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1NegotiablecartsCartIdCoupons**](NegotiableCartsCartIdCouponsAPI.md#DeleteV1NegotiablecartsCartIdCoupons) | **Delete** /V1/negotiable-carts/{cartId}/coupons | negotiable-carts/{cartId}/coupons



## DeleteV1NegotiablecartsCartIdCoupons

> bool DeleteV1NegotiablecartsCartIdCoupons(ctx, cartId).Execute()

negotiable-carts/{cartId}/coupons



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
	cartId := int32(56) // int32 | The cart ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableCartsCartIdCouponsAPI.DeleteV1NegotiablecartsCartIdCoupons(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableCartsCartIdCouponsAPI.DeleteV1NegotiablecartsCartIdCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1NegotiablecartsCartIdCoupons`: bool
	fmt.Fprintf(os.Stdout, "Response from `NegotiableCartsCartIdCouponsAPI.DeleteV1NegotiablecartsCartIdCoupons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1NegotiablecartsCartIdCouponsRequest struct via the builder pattern


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

