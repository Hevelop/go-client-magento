# \CreditmemoIdCommentsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1CreditmemoIdComments**](CreditmemoIdCommentsAPI.md#GetV1CreditmemoIdComments) | **Get** /V1/creditmemo/{id}/comments | creditmemo/{id}/comments
[**PostV1CreditmemoIdComments**](CreditmemoIdCommentsAPI.md#PostV1CreditmemoIdComments) | **Post** /V1/creditmemo/{id}/comments | creditmemo/{id}/comments



## GetV1CreditmemoIdComments

> SalesDataCreditmemoCommentSearchResultInterface GetV1CreditmemoIdComments(ctx, id).Execute()

creditmemo/{id}/comments



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
	id := int32(56) // int32 | The credit memo ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditmemoIdCommentsAPI.GetV1CreditmemoIdComments(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditmemoIdCommentsAPI.GetV1CreditmemoIdComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1CreditmemoIdComments`: SalesDataCreditmemoCommentSearchResultInterface
	fmt.Fprintf(os.Stdout, "Response from `CreditmemoIdCommentsAPI.GetV1CreditmemoIdComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The credit memo ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1CreditmemoIdCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalesDataCreditmemoCommentSearchResultInterface**](SalesDataCreditmemoCommentSearchResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1CreditmemoIdComments

> SalesDataCreditmemoCommentInterface PostV1CreditmemoIdComments(ctx, id).PostV1CreditmemoIdCommentsRequest(postV1CreditmemoIdCommentsRequest).Execute()

creditmemo/{id}/comments



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
	postV1CreditmemoIdCommentsRequest := *openapiclient.NewPostV1CreditmemoIdCommentsRequest(*openapiclient.NewSalesDataCreditmemoCommentInterface("Comment_example", int32(123), int32(123), int32(123))) // PostV1CreditmemoIdCommentsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditmemoIdCommentsAPI.PostV1CreditmemoIdComments(context.Background(), id).PostV1CreditmemoIdCommentsRequest(postV1CreditmemoIdCommentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditmemoIdCommentsAPI.PostV1CreditmemoIdComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CreditmemoIdComments`: SalesDataCreditmemoCommentInterface
	fmt.Fprintf(os.Stdout, "Response from `CreditmemoIdCommentsAPI.PostV1CreditmemoIdComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CreditmemoIdCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1CreditmemoIdCommentsRequest** | [**PostV1CreditmemoIdCommentsRequest**](PostV1CreditmemoIdCommentsRequest.md) |  | 

### Return type

[**SalesDataCreditmemoCommentInterface**](SalesDataCreditmemoCommentInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

