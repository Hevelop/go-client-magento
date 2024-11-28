# \CouponsCouponIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CouponsCouponId**](CouponsCouponIdAPI.md#DeleteV1CouponsCouponId) | **Delete** /V1/coupons/{couponId} | coupons/{couponId}
[**GetV1CouponsCouponId**](CouponsCouponIdAPI.md#GetV1CouponsCouponId) | **Get** /V1/coupons/{couponId} | coupons/{couponId}
[**PutV1CouponsCouponId**](CouponsCouponIdAPI.md#PutV1CouponsCouponId) | **Put** /V1/coupons/{couponId} | coupons/{couponId}



## DeleteV1CouponsCouponId

> bool DeleteV1CouponsCouponId(ctx, couponId).Execute()

coupons/{couponId}



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
	couponId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsCouponIdAPI.DeleteV1CouponsCouponId(context.Background(), couponId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsCouponIdAPI.DeleteV1CouponsCouponId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CouponsCouponId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CouponsCouponIdAPI.DeleteV1CouponsCouponId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CouponsCouponIdRequest struct via the builder pattern


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


## GetV1CouponsCouponId

> SalesRuleDataCouponInterface GetV1CouponsCouponId(ctx, couponId).Execute()

coupons/{couponId}



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
	couponId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsCouponIdAPI.GetV1CouponsCouponId(context.Background(), couponId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsCouponIdAPI.GetV1CouponsCouponId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CouponsCouponId`: SalesRuleDataCouponInterface
	fmt.Fprintf(os.Stdout, "Response from `CouponsCouponIdAPI.GetV1CouponsCouponId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CouponsCouponIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalesRuleDataCouponInterface**](SalesRuleDataCouponInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1CouponsCouponId

> SalesRuleDataCouponInterface PutV1CouponsCouponId(ctx, couponId).PostV1CouponsRequest(postV1CouponsRequest).Execute()

coupons/{couponId}



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
	couponId := "couponId_example" // string | 
	postV1CouponsRequest := *openapiclient.NewPostV1CouponsRequest(*openapiclient.NewSalesRuleDataCouponInterface(int32(123), int32(123), false)) // PostV1CouponsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsCouponIdAPI.PutV1CouponsCouponId(context.Background(), couponId).PostV1CouponsRequest(postV1CouponsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsCouponIdAPI.PutV1CouponsCouponId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CouponsCouponId`: SalesRuleDataCouponInterface
	fmt.Fprintf(os.Stdout, "Response from `CouponsCouponIdAPI.PutV1CouponsCouponId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**couponId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CouponsCouponIdRequest struct via the builder pattern


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

