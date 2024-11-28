# \CartsCartIdCouponsDeleteByCodesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV2CartsCartIdCouponsDeleteByCodes**](CartsCartIdCouponsDeleteByCodesAPI.md#PostV2CartsCartIdCouponsDeleteByCodes) | **Post** /V2/carts/{cartId}/coupons/deleteByCodes | carts/{cartId}/coupons/deleteByCodes



## PostV2CartsCartIdCouponsDeleteByCodes

> ErrorResponse PostV2CartsCartIdCouponsDeleteByCodes(ctx, cartId).PostV2CartsMineCouponsDeleteByCodesRequest(postV2CartsMineCouponsDeleteByCodesRequest).Execute()

carts/{cartId}/coupons/deleteByCodes



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
	postV2CartsMineCouponsDeleteByCodesRequest := *openapiclient.NewPostV2CartsMineCouponsDeleteByCodesRequest() // PostV2CartsMineCouponsDeleteByCodesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdCouponsDeleteByCodesAPI.PostV2CartsCartIdCouponsDeleteByCodes(context.Background(), cartId).PostV2CartsMineCouponsDeleteByCodesRequest(postV2CartsMineCouponsDeleteByCodesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdCouponsDeleteByCodesAPI.PostV2CartsCartIdCouponsDeleteByCodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV2CartsCartIdCouponsDeleteByCodes`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdCouponsDeleteByCodesAPI.PostV2CartsCartIdCouponsDeleteByCodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV2CartsCartIdCouponsDeleteByCodesRequest struct via the builder pattern


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

