# \ShipmentIdCommentsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ShipmentIdComments**](ShipmentIdCommentsAPI.md#GetV1ShipmentIdComments) | **Get** /V1/shipment/{id}/comments | shipment/{id}/comments
[**PostV1ShipmentIdComments**](ShipmentIdCommentsAPI.md#PostV1ShipmentIdComments) | **Post** /V1/shipment/{id}/comments | shipment/{id}/comments



## GetV1ShipmentIdComments

> SalesDataShipmentCommentSearchResultInterface GetV1ShipmentIdComments(ctx, id).Execute()

shipment/{id}/comments



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
	id := int32(56) // int32 | The shipment ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentIdCommentsAPI.GetV1ShipmentIdComments(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentIdCommentsAPI.GetV1ShipmentIdComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ShipmentIdComments`: SalesDataShipmentCommentSearchResultInterface
	fmt.Fprintf(os.Stdout, "Response from `ShipmentIdCommentsAPI.GetV1ShipmentIdComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The shipment ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ShipmentIdCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalesDataShipmentCommentSearchResultInterface**](SalesDataShipmentCommentSearchResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1ShipmentIdComments

> SalesDataShipmentCommentInterface PostV1ShipmentIdComments(ctx, id).PostV1ShipmentIdCommentsRequest(postV1ShipmentIdCommentsRequest).Execute()

shipment/{id}/comments



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
	id := "id_example" // string | 
	postV1ShipmentIdCommentsRequest := *openapiclient.NewPostV1ShipmentIdCommentsRequest(*openapiclient.NewSalesDataShipmentCommentInterface(int32(123), int32(123), "Comment_example", int32(123))) // PostV1ShipmentIdCommentsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentIdCommentsAPI.PostV1ShipmentIdComments(context.Background(), id).PostV1ShipmentIdCommentsRequest(postV1ShipmentIdCommentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentIdCommentsAPI.PostV1ShipmentIdComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ShipmentIdComments`: SalesDataShipmentCommentInterface
	fmt.Fprintf(os.Stdout, "Response from `ShipmentIdCommentsAPI.PostV1ShipmentIdComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ShipmentIdCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ShipmentIdCommentsRequest** | [**PostV1ShipmentIdCommentsRequest**](PostV1ShipmentIdCommentsRequest.md) |  | 

### Return type

[**SalesDataShipmentCommentInterface**](SalesDataShipmentCommentInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

