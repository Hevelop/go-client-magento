# \CartsMineCouponsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CartsMineCoupons**](CartsMineCouponsAPI.md#DeleteV1CartsMineCoupons) | **Delete** /V1/carts/mine/coupons | carts/mine/coupons
[**GetV1CartsMineCoupons**](CartsMineCouponsAPI.md#GetV1CartsMineCoupons) | **Get** /V1/carts/mine/coupons | carts/mine/coupons
[**GetV2CartsMineCoupons**](CartsMineCouponsAPI.md#GetV2CartsMineCoupons) | **Get** /V2/carts/mine/coupons | carts/mine/coupons
[**PostV2CartsMineCoupons**](CartsMineCouponsAPI.md#PostV2CartsMineCoupons) | **Post** /V2/carts/mine/coupons | carts/mine/coupons
[**PutV2CartsMineCoupons**](CartsMineCouponsAPI.md#PutV2CartsMineCoupons) | **Put** /V2/carts/mine/coupons | carts/mine/coupons



## DeleteV1CartsMineCoupons

> bool DeleteV1CartsMineCoupons(ctx).Execute()

carts/mine/coupons



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineCouponsAPI.DeleteV1CartsMineCoupons(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineCouponsAPI.DeleteV1CartsMineCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CartsMineCoupons`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsMineCouponsAPI.DeleteV1CartsMineCoupons`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CartsMineCouponsRequest struct via the builder pattern


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


## GetV1CartsMineCoupons

> string GetV1CartsMineCoupons(ctx).Execute()

carts/mine/coupons



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineCouponsAPI.GetV1CartsMineCoupons(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineCouponsAPI.GetV1CartsMineCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsMineCoupons`: string
	fmt.Fprintf(os.Stdout, "Response from `CartsMineCouponsAPI.GetV1CartsMineCoupons`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsMineCouponsRequest struct via the builder pattern


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


## GetV2CartsMineCoupons

> []string GetV2CartsMineCoupons(ctx).Execute()

carts/mine/coupons



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineCouponsAPI.GetV2CartsMineCoupons(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineCouponsAPI.GetV2CartsMineCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV2CartsMineCoupons`: []string
	fmt.Fprintf(os.Stdout, "Response from `CartsMineCouponsAPI.GetV2CartsMineCoupons`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV2CartsMineCouponsRequest struct via the builder pattern


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


## PostV2CartsMineCoupons

> ErrorResponse PostV2CartsMineCoupons(ctx).PutV2CartsMineCouponsRequest(putV2CartsMineCouponsRequest).Execute()

carts/mine/coupons



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
	putV2CartsMineCouponsRequest := *openapiclient.NewPutV2CartsMineCouponsRequest([]string{"CouponCodes_example"}) // PutV2CartsMineCouponsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineCouponsAPI.PostV2CartsMineCoupons(context.Background()).PutV2CartsMineCouponsRequest(putV2CartsMineCouponsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineCouponsAPI.PostV2CartsMineCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV2CartsMineCoupons`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `CartsMineCouponsAPI.PostV2CartsMineCoupons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV2CartsMineCouponsRequest struct via the builder pattern


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


## PutV2CartsMineCoupons

> ErrorResponse PutV2CartsMineCoupons(ctx).PutV2CartsMineCouponsRequest(putV2CartsMineCouponsRequest).Execute()

carts/mine/coupons



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
	putV2CartsMineCouponsRequest := *openapiclient.NewPutV2CartsMineCouponsRequest([]string{"CouponCodes_example"}) // PutV2CartsMineCouponsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineCouponsAPI.PutV2CartsMineCoupons(context.Background()).PutV2CartsMineCouponsRequest(putV2CartsMineCouponsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineCouponsAPI.PutV2CartsMineCoupons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV2CartsMineCoupons`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `CartsMineCouponsAPI.PutV2CartsMineCoupons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV2CartsMineCouponsRequest struct via the builder pattern


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

