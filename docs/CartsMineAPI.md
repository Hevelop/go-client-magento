# \CartsMineAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsMine**](CartsMineAPI.md#GetV1CartsMine) | **Get** /V1/carts/mine | carts/mine
[**PostV1CartsMine**](CartsMineAPI.md#PostV1CartsMine) | **Post** /V1/carts/mine | carts/mine
[**PutV1CartsMine**](CartsMineAPI.md#PutV1CartsMine) | **Put** /V1/carts/mine | carts/mine



## GetV1CartsMine

> QuoteDataCartInterface GetV1CartsMine(ctx).Execute()

carts/mine



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
	resp, r, err := apiClient.CartsMineAPI.GetV1CartsMine(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineAPI.GetV1CartsMine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsMine`: QuoteDataCartInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMineAPI.GetV1CartsMine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsMineRequest struct via the builder pattern


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


## PostV1CartsMine

> int32 PostV1CartsMine(ctx).Execute()

carts/mine



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
	resp, r, err := apiClient.CartsMineAPI.PostV1CartsMine(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineAPI.PostV1CartsMine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMine`: int32
	fmt.Fprintf(os.Stdout, "Response from `CartsMineAPI.PostV1CartsMine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMineRequest struct via the builder pattern


### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1CartsMine

> ErrorResponse PutV1CartsMine(ctx).PutV1CartsMineRequest(putV1CartsMineRequest).Execute()

carts/mine



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
	putV1CartsMineRequest := *openapiclient.NewPutV1CartsMineRequest(*openapiclient.NewQuoteDataCartInterface(int32(123), *openapiclient.NewCustomerDataCustomerInterface("Email_example", "Firstname_example", "Lastname_example"), int32(123))) // PutV1CartsMineRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineAPI.PutV1CartsMine(context.Background()).PutV1CartsMineRequest(putV1CartsMineRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineAPI.PutV1CartsMine``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CartsMine`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `CartsMineAPI.PutV1CartsMine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CartsMineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putV1CartsMineRequest** | [**PutV1CartsMineRequest**](PutV1CartsMineRequest.md) |  | 

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

