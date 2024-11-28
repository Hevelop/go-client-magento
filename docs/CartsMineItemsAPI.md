# \CartsMineItemsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CartsMineItems**](CartsMineItemsAPI.md#GetV1CartsMineItems) | **Get** /V1/carts/mine/items | carts/mine/items
[**PostV1CartsMineItems**](CartsMineItemsAPI.md#PostV1CartsMineItems) | **Post** /V1/carts/mine/items | carts/mine/items



## GetV1CartsMineItems

> []QuoteDataCartItemInterface GetV1CartsMineItems(ctx).Execute()

carts/mine/items



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
	resp, r, err := apiClient.CartsMineItemsAPI.GetV1CartsMineItems(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineItemsAPI.GetV1CartsMineItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CartsMineItems`: []QuoteDataCartItemInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMineItemsAPI.GetV1CartsMineItems`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CartsMineItemsRequest struct via the builder pattern


### Return type

[**[]QuoteDataCartItemInterface**](QuoteDataCartItemInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1CartsMineItems

> QuoteDataCartItemInterface PostV1CartsMineItems(ctx).PostV1CartsMineItemsRequest(postV1CartsMineItemsRequest).Execute()

carts/mine/items



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
	postV1CartsMineItemsRequest := *openapiclient.NewPostV1CartsMineItemsRequest(*openapiclient.NewQuoteDataCartItemInterface(float32(123), "QuoteId_example")) // PostV1CartsMineItemsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineItemsAPI.PostV1CartsMineItems(context.Background()).PostV1CartsMineItemsRequest(postV1CartsMineItemsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineItemsAPI.PostV1CartsMineItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMineItems`: QuoteDataCartItemInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMineItemsAPI.PostV1CartsMineItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMineItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CartsMineItemsRequest** | [**PostV1CartsMineItemsRequest**](PostV1CartsMineItemsRequest.md) |  | 

### Return type

[**QuoteDataCartItemInterface**](QuoteDataCartItemInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

