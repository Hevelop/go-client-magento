# \CartsMineCouponsDeleteByCodesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV2CartsMineCouponsDeleteByCodes**](CartsMineCouponsDeleteByCodesAPI.md#PostV2CartsMineCouponsDeleteByCodes) | **Post** /V2/carts/mine/coupons/deleteByCodes | carts/mine/coupons/deleteByCodes



## PostV2CartsMineCouponsDeleteByCodes

> ErrorResponse PostV2CartsMineCouponsDeleteByCodes(ctx).PostV2CartsMineCouponsDeleteByCodesRequest(postV2CartsMineCouponsDeleteByCodesRequest).Execute()

carts/mine/coupons/deleteByCodes



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
	postV2CartsMineCouponsDeleteByCodesRequest := *openapiclient.NewPostV2CartsMineCouponsDeleteByCodesRequest() // PostV2CartsMineCouponsDeleteByCodesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineCouponsDeleteByCodesAPI.PostV2CartsMineCouponsDeleteByCodes(context.Background()).PostV2CartsMineCouponsDeleteByCodesRequest(postV2CartsMineCouponsDeleteByCodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineCouponsDeleteByCodesAPI.PostV2CartsMineCouponsDeleteByCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV2CartsMineCouponsDeleteByCodes`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `CartsMineCouponsDeleteByCodesAPI.PostV2CartsMineCouponsDeleteByCodes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV2CartsMineCouponsDeleteByCodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV2CartsMineCouponsDeleteByCodesRequest** | [**PostV2CartsMineCouponsDeleteByCodesRequest**](PostV2CartsMineCouponsDeleteByCodesRequest.md) |  | 

### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

