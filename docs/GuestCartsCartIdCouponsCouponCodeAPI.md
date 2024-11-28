# \GuestCartsCartIdCouponsCouponCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1GuestcartsCartIdCouponsCouponCode**](GuestCartsCartIdCouponsCouponCodeAPI.md#PutV1GuestcartsCartIdCouponsCouponCode) | **Put** /V1/guest-carts/{cartId}/coupons/{couponCode} | guest-carts/{cartId}/coupons/{couponCode}



## PutV1GuestcartsCartIdCouponsCouponCode

> bool PutV1GuestcartsCartIdCouponsCouponCode(ctx, cartId, couponCode).Execute()

guest-carts/{cartId}/coupons/{couponCode}



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
	cartId := "cartId_example" // string | The cart ID.
	couponCode := "couponCode_example" // string | The coupon code data.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdCouponsCouponCodeAPI.PutV1GuestcartsCartIdCouponsCouponCode(context.Background(), cartId, couponCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdCouponsCouponCodeAPI.PutV1GuestcartsCartIdCouponsCouponCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1GuestcartsCartIdCouponsCouponCode`: bool
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdCouponsCouponCodeAPI.PutV1GuestcartsCartIdCouponsCouponCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 
**couponCode** | **string** | The coupon code data. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1GuestcartsCartIdCouponsCouponCodeRequest struct via the builder pattern


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

