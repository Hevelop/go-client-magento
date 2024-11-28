# \CartsCartIdCouponsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CartsCartIdCoupons**](CartsCartIdCouponsAPI.md#DeleteV1CartsCartIdCoupons) | **Delete** /V1/carts/{cartId}/coupons | carts/{cartId}/coupons
[**GetV1CartsCartIdCoupons**](CartsCartIdCouponsAPI.md#GetV1CartsCartIdCoupons) | **Get** /V1/carts/{cartId}/coupons | carts/{cartId}/coupons
[**GetV2CartsCartIdCoupons**](CartsCartIdCouponsAPI.md#GetV2CartsCartIdCoupons) | **Get** /V2/carts/{cartId}/coupons | carts/{cartId}/coupons
[**PostV2CartsCartIdCoupons**](CartsCartIdCouponsAPI.md#PostV2CartsCartIdCoupons) | **Post** /V2/carts/{cartId}/coupons | carts/{cartId}/coupons
[**PutV2CartsCartIdCoupons**](CartsCartIdCouponsAPI.md#PutV2CartsCartIdCoupons) | **Put** /V2/carts/{cartId}/coupons | carts/{cartId}/coupons



## DeleteV1CartsCartIdCoupons

> bool DeleteV1CartsCartIdCoupons(ctx, cartId).Execute()

carts/{cartId}/coupons



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
	resp, r, err := apiClient.CartsCartIdCouponsAPI.DeleteV1CartsCartIdCoupons(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdCouponsAPI.DeleteV1CartsCartIdCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CartsCartIdCoupons`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdCouponsAPI.DeleteV1CartsCartIdCoupons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CartsCartIdCouponsRequest struct via the builder pattern


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


## GetV1CartsCartIdCoupons

> string GetV1CartsCartIdCoupons(ctx, cartId).Execute()

carts/{cartId}/coupons



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
	resp, r, err := apiClient.CartsCartIdCouponsAPI.GetV1CartsCartIdCoupons(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdCouponsAPI.GetV1CartsCartIdCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsCartIdCoupons`: string
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdCouponsAPI.GetV1CartsCartIdCoupons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsCartIdCouponsRequest struct via the builder pattern


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


## GetV2CartsCartIdCoupons

> []string GetV2CartsCartIdCoupons(ctx, cartId).Execute()

carts/{cartId}/coupons



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
	resp, r, err := apiClient.CartsCartIdCouponsAPI.GetV2CartsCartIdCoupons(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdCouponsAPI.GetV2CartsCartIdCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV2CartsCartIdCoupons`: []string
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdCouponsAPI.GetV2CartsCartIdCoupons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV2CartsCartIdCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV2CartsCartIdCoupons

> ErrorResponse PostV2CartsCartIdCoupons(ctx, cartId).PutV2CartsMineCouponsRequest(putV2CartsMineCouponsRequest).Execute()

carts/{cartId}/coupons



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
	cartId := int32(56) // int32 | 
	putV2CartsMineCouponsRequest := *openapiclient.NewPutV2CartsMineCouponsRequest([]string{"CouponCodes_example"}) // PutV2CartsMineCouponsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdCouponsAPI.PostV2CartsCartIdCoupons(context.Background(), cartId).PutV2CartsMineCouponsRequest(putV2CartsMineCouponsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdCouponsAPI.PostV2CartsCartIdCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV2CartsCartIdCoupons`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdCouponsAPI.PostV2CartsCartIdCoupons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV2CartsCartIdCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV2CartsMineCouponsRequest** | [**PutV2CartsMineCouponsRequest**](PutV2CartsMineCouponsRequest.md) |  | 

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


## PutV2CartsCartIdCoupons

> ErrorResponse PutV2CartsCartIdCoupons(ctx, cartId).PutV2CartsMineCouponsRequest(putV2CartsMineCouponsRequest).Execute()

carts/{cartId}/coupons



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
	cartId := int32(56) // int32 | 
	putV2CartsMineCouponsRequest := *openapiclient.NewPutV2CartsMineCouponsRequest([]string{"CouponCodes_example"}) // PutV2CartsMineCouponsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdCouponsAPI.PutV2CartsCartIdCoupons(context.Background(), cartId).PutV2CartsMineCouponsRequest(putV2CartsMineCouponsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdCouponsAPI.PutV2CartsCartIdCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV2CartsCartIdCoupons`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdCouponsAPI.PutV2CartsCartIdCoupons`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV2CartsCartIdCouponsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV2CartsMineCouponsRequest** | [**PutV2CartsMineCouponsRequest**](PutV2CartsMineCouponsRequest.md) |  | 

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

