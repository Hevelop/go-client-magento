# \CartsCartIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsCartId**](CartsCartIdAPI.md#GetV1CartsCartId) | **Get** /V1/carts/{cartId} | carts/{cartId}
[**PutV1CartsCartId**](CartsCartIdAPI.md#PutV1CartsCartId) | **Put** /V1/carts/{cartId} | carts/{cartId}



## GetV1CartsCartId

> QuoteDataCartInterface GetV1CartsCartId(ctx, cartId).Execute()

carts/{cartId}



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdAPI.GetV1CartsCartId(context.Background(), cartId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdAPI.GetV1CartsCartId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsCartId`: QuoteDataCartInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdAPI.GetV1CartsCartId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsCartIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**QuoteDataCartInterface**](QuoteDataCartInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1CartsCartId

> bool PutV1CartsCartId(ctx, cartId).PutV1CartsCartIdRequest(putV1CartsCartIdRequest).Execute()

carts/{cartId}



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
	putV1CartsCartIdRequest := *openapiclient.NewPutV1CartsCartIdRequest(int32(123), int32(123)) // PutV1CartsCartIdRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsCartIdAPI.PutV1CartsCartId(context.Background(), cartId).PutV1CartsCartIdRequest(putV1CartsCartIdRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsCartIdAPI.PutV1CartsCartId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CartsCartId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsCartIdAPI.PutV1CartsCartId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cartId** | **int32** | The cart ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CartsCartIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **putV1CartsCartIdRequest** | [**PutV1CartsCartIdRequest**](PutV1CartsCartIdRequest.md) |  | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

