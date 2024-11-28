# \NegotiableCartsCartIdCouponsCouponCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1NegotiablecartsCartIdCouponsCouponCode**](NegotiableCartsCartIdCouponsCouponCodeAPI.md#PutV1NegotiablecartsCartIdCouponsCouponCode) | **Put** /V1/negotiable-carts/{cartId}/coupons/{couponCode} | negotiable-carts/{cartId}/coupons/{couponCode}



## PutV1NegotiablecartsCartIdCouponsCouponCode

> bool PutV1NegotiablecartsCartIdCouponsCouponCode(ctx, cartId, couponCode).Execute()

negotiable-carts/{cartId}/coupons/{couponCode}



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
	couponCode := "couponCode_example" // string | The coupon code data.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableCartsCartIdCouponsCouponCodeAPI.PutV1NegotiablecartsCartIdCouponsCouponCode(context.Background(), cartId, couponCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableCartsCartIdCouponsCouponCodeAPI.PutV1NegotiablecartsCartIdCouponsCouponCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1NegotiablecartsCartIdCouponsCouponCode`: bool
	fmt.Fprintf(os.Stdout, "Response from `NegotiableCartsCartIdCouponsCouponCodeAPI.PutV1NegotiablecartsCartIdCouponsCouponCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 
**couponCode** | **string** | The coupon code data. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1NegotiablecartsCartIdCouponsCouponCodeRequest struct via the builder pattern


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

