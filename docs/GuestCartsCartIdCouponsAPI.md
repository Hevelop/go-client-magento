# \GuestCartsCartIdCouponsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1GuestcartsCartIdCoupons**](GuestCartsCartIdCouponsAPI.md#DeleteV1GuestcartsCartIdCoupons) | **Delete** /V1/guest-carts/{cartId}/coupons | guest-carts/{cartId}/coupons
[**GetV1GuestcartsCartIdCoupons**](GuestCartsCartIdCouponsAPI.md#GetV1GuestcartsCartIdCoupons) | **Get** /V1/guest-carts/{cartId}/coupons | guest-carts/{cartId}/coupons



## DeleteV1GuestcartsCartIdCoupons

> bool DeleteV1GuestcartsCartIdCoupons(ctx, cartId).Execute()

guest-carts/{cartId}/coupons



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdCouponsAPI.DeleteV1GuestcartsCartIdCoupons(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdCouponsAPI.DeleteV1GuestcartsCartIdCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1GuestcartsCartIdCoupons`: bool
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdCouponsAPI.DeleteV1GuestcartsCartIdCoupons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1GuestcartsCartIdCouponsRequest struct via the builder pattern


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


## GetV1GuestcartsCartIdCoupons

> string GetV1GuestcartsCartIdCoupons(ctx, cartId).Execute()

guest-carts/{cartId}/coupons



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GuestCartsCartIdCouponsAPI.GetV1GuestcartsCartIdCoupons(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GuestCartsCartIdCouponsAPI.GetV1GuestcartsCartIdCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1GuestcartsCartIdCoupons`: string
	fmt.Fprintf(os.Stdout, "Response from `GuestCartsCartIdCouponsAPI.GetV1GuestcartsCartIdCoupons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **string** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1GuestcartsCartIdCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

