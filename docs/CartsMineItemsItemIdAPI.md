# \CartsMineItemsItemIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1CartsMineItemsItemId**](CartsMineItemsItemIdAPI.md#DeleteV1CartsMineItemsItemId) | **Delete** /V1/carts/mine/items/{itemId} | carts/mine/items/{itemId}
[**PutV1CartsMineItemsItemId**](CartsMineItemsItemIdAPI.md#PutV1CartsMineItemsItemId) | **Put** /V1/carts/mine/items/{itemId} | carts/mine/items/{itemId}



## DeleteV1CartsMineItemsItemId

> bool DeleteV1CartsMineItemsItemId(ctx, itemId).Execute()

carts/mine/items/{itemId}



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
	itemId := int32(56) // int32 | The item ID of the item to be removed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineItemsItemIdAPI.DeleteV1CartsMineItemsItemId(context.Background(), itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineItemsItemIdAPI.DeleteV1CartsMineItemsItemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1CartsMineItemsItemId`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsMineItemsItemIdAPI.DeleteV1CartsMineItemsItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **int32** | The item ID of the item to be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1CartsMineItemsItemIdRequest struct via the builder pattern


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


## PutV1CartsMineItemsItemId

> QuoteDataCartItemInterface PutV1CartsMineItemsItemId(ctx, itemId).PostV1CartsMineItemsRequest(postV1CartsMineItemsRequest).Execute()

carts/mine/items/{itemId}



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
	itemId := "itemId_example" // string | 
	postV1CartsMineItemsRequest := *openapiclient.NewPostV1CartsMineItemsRequest(*openapiclient.NewQuoteDataCartItemInterface(float32(123), "QuoteId_example")) // PostV1CartsMineItemsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CartsMineItemsItemIdAPI.PutV1CartsMineItemsItemId(context.Background(), itemId).PostV1CartsMineItemsRequest(postV1CartsMineItemsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineItemsItemIdAPI.PutV1CartsMineItemsItemId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1CartsMineItemsItemId`: QuoteDataCartItemInterface
	fmt.Fprintf(os.Stdout, "Response from `CartsMineItemsItemIdAPI.PutV1CartsMineItemsItemId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1CartsMineItemsItemIdRequest struct via the builder pattern


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

