# \CouponsDeleteByCodesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CouponsDeleteByCodes**](CouponsDeleteByCodesAPI.md#PostV1CouponsDeleteByCodes) | **Post** /V1/coupons/deleteByCodes | coupons/deleteByCodes



## PostV1CouponsDeleteByCodes

> SalesRuleDataCouponMassDeleteResultInterface PostV1CouponsDeleteByCodes(ctx).PostV1CouponsDeleteByCodesRequest(postV1CouponsDeleteByCodesRequest).Execute()

coupons/deleteByCodes



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
	postV1CouponsDeleteByCodesRequest := *openapiclient.NewPostV1CouponsDeleteByCodesRequest([]string{"Codes_example"}) // PostV1CouponsDeleteByCodesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsDeleteByCodesAPI.PostV1CouponsDeleteByCodes(context.Background()).PostV1CouponsDeleteByCodesRequest(postV1CouponsDeleteByCodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsDeleteByCodesAPI.PostV1CouponsDeleteByCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CouponsDeleteByCodes`: SalesRuleDataCouponMassDeleteResultInterface
	fmt.Fprintf(os.Stdout, "Response from `CouponsDeleteByCodesAPI.PostV1CouponsDeleteByCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CouponsDeleteByCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CouponsDeleteByCodesRequest** | [**PostV1CouponsDeleteByCodesRequest**](PostV1CouponsDeleteByCodesRequest.md) |  | 

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

