# \CouponsDeleteByIdsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CouponsDeleteByIds**](CouponsDeleteByIdsAPI.md#PostV1CouponsDeleteByIds) | **Post** /V1/coupons/deleteByIds | coupons/deleteByIds



## PostV1CouponsDeleteByIds

> SalesRuleDataCouponMassDeleteResultInterface PostV1CouponsDeleteByIds(ctx).PostV1CouponsDeleteByIdsRequest(postV1CouponsDeleteByIdsRequest).Execute()

coupons/deleteByIds



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
	postV1CouponsDeleteByIdsRequest := *openapiclient.NewPostV1CouponsDeleteByIdsRequest([]int32{int32(123)}) // PostV1CouponsDeleteByIdsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsDeleteByIdsAPI.PostV1CouponsDeleteByIds(context.Background()).PostV1CouponsDeleteByIdsRequest(postV1CouponsDeleteByIdsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsDeleteByIdsAPI.PostV1CouponsDeleteByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CouponsDeleteByIds`: SalesRuleDataCouponMassDeleteResultInterface
	fmt.Fprintf(os.Stdout, "Response from `CouponsDeleteByIdsAPI.PostV1CouponsDeleteByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CouponsDeleteByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CouponsDeleteByIdsRequest** | [**PostV1CouponsDeleteByIdsRequest**](PostV1CouponsDeleteByIdsRequest.md) |  | 

### Return type

[**SalesRuleDataCouponMassDeleteResultInterface**](SalesRuleDataCouponMassDeleteResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

