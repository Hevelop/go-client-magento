# \CouponsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1Coupons**](CouponsAPI.md#PostV1Coupons) | **Post** /V1/coupons | coupons



## PostV1Coupons

> SalesRuleDataCouponInterface PostV1Coupons(ctx).PostV1CouponsRequest(postV1CouponsRequest).Execute()

coupons



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
	postV1CouponsRequest := *openapiclient.NewPostV1CouponsRequest(*openapiclient.NewSalesRuleDataCouponInterface(int32(123), int32(123), false)) // PostV1CouponsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsAPI.PostV1Coupons(context.Background()).PostV1CouponsRequest(postV1CouponsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsAPI.PostV1Coupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1Coupons`: SalesRuleDataCouponInterface
	fmt.Fprintf(os.Stdout, "Response from `CouponsAPI.PostV1Coupons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CouponsRequest** | [**PostV1CouponsRequest**](PostV1CouponsRequest.md) |  | 

### Return type

[**SalesRuleDataCouponInterface**](SalesRuleDataCouponInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

