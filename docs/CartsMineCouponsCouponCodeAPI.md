# \CartsMineCouponsCouponCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PutV1CartsMineCouponsCouponCode**](CartsMineCouponsCouponCodeAPI.md#PutV1CartsMineCouponsCouponCode) | **Put** /V1/carts/mine/coupons/{couponCode} | carts/mine/coupons/{couponCode}



## PutV1CartsMineCouponsCouponCode

> bool PutV1CartsMineCouponsCouponCode(ctx, couponCode).Execute()

carts/mine/coupons/{couponCode}



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
	couponCode := "couponCode_example" // string | The coupon code data.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineCouponsCouponCodeAPI.PutV1CartsMineCouponsCouponCode(context.Background(), couponCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineCouponsCouponCodeAPI.PutV1CartsMineCouponsCouponCode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CartsMineCouponsCouponCode`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsMineCouponsCouponCodeAPI.PutV1CartsMineCouponsCouponCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponCode** | **string** | The coupon code data. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CartsMineCouponsCouponCodeRequest struct via the builder pattern


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

