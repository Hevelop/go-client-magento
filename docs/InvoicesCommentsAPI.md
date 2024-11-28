# \InvoicesCommentsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1InvoicesComments**](InvoicesCommentsAPI.md#PostV1InvoicesComments) | **Post** /V1/invoices/comments | invoices/comments



## PostV1InvoicesComments

> SalesDataInvoiceCommentInterface PostV1InvoicesComments(ctx).PostV1InvoicesCommentsRequest(postV1InvoicesCommentsRequest).Execute()

invoices/comments



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
	postV1InvoicesCommentsRequest := *openapiclient.NewPostV1InvoicesCommentsRequest(*openapiclient.NewSalesDataInvoiceCommentInterface(int32(123), int32(123), "Comment_example", int32(123))) // PostV1InvoicesCommentsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesCommentsAPI.PostV1InvoicesComments(context.Background()).PostV1InvoicesCommentsRequest(postV1InvoicesCommentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesCommentsAPI.PostV1InvoicesComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1InvoicesComments`: SalesDataInvoiceCommentInterface
	fmt.Fprintf(os.Stdout, "Response from `InvoicesCommentsAPI.PostV1InvoicesComments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1InvoicesCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1InvoicesCommentsRequest** | [**PostV1InvoicesCommentsRequest**](PostV1InvoicesCommentsRequest.md) |  | 

### Return type

[**SalesDataInvoiceCommentInterface**](SalesDataInvoiceCommentInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

