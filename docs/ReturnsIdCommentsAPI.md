# \ReturnsIdCommentsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1ReturnsIdComments**](ReturnsIdCommentsAPI.md#GetV1ReturnsIdComments) | **Get** /V1/returns/{id}/comments | returns/{id}/comments
[**PostV1ReturnsIdComments**](ReturnsIdCommentsAPI.md#PostV1ReturnsIdComments) | **Post** /V1/returns/{id}/comments | returns/{id}/comments



## GetV1ReturnsIdComments

> RmaDataCommentSearchResultInterface GetV1ReturnsIdComments(ctx, id).Execute()

returns/{id}/comments



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
	id := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsIdCommentsAPI.GetV1ReturnsIdComments(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsIdCommentsAPI.GetV1ReturnsIdComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1ReturnsIdComments`: RmaDataCommentSearchResultInterface
	fmt.Fprintf(os.Stdout, "Response from `ReturnsIdCommentsAPI.GetV1ReturnsIdComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1ReturnsIdCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RmaDataCommentSearchResultInterface**](RmaDataCommentSearchResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostV1ReturnsIdComments

> bool PostV1ReturnsIdComments(ctx, id).PostV1ReturnsIdCommentsRequest(postV1ReturnsIdCommentsRequest).Execute()

returns/{id}/comments



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
	postV1ReturnsIdCommentsRequest := *openapiclient.NewPostV1ReturnsIdCommentsRequest(*openapiclient.NewRmaDataCommentInterface("Comment_example", int32(123), "CreatedAt_example", int32(123), false, false, "Status_example", false)) // PostV1ReturnsIdCommentsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReturnsIdCommentsAPI.PostV1ReturnsIdComments(context.Background(), id).PostV1ReturnsIdCommentsRequest(postV1ReturnsIdCommentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReturnsIdCommentsAPI.PostV1ReturnsIdComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1ReturnsIdComments`: bool
	fmt.Fprintf(os.Stdout, "Response from `ReturnsIdCommentsAPI.PostV1ReturnsIdComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1ReturnsIdCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1ReturnsIdCommentsRequest** | [**PostV1ReturnsIdCommentsRequest**](PostV1ReturnsIdCommentsRequest.md) |  | 

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

