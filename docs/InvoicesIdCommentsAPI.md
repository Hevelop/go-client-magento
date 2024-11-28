# \InvoicesIdCommentsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InvoicesIdComments**](InvoicesIdCommentsAPI.md#GetV1InvoicesIdComments) | **Get** /V1/invoices/{id}/comments | invoices/{id}/comments



## GetV1InvoicesIdComments

> SalesDataInvoiceCommentSearchResultInterface GetV1InvoicesIdComments(ctx, id).Execute()

invoices/{id}/comments



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
	id := int32(56) // int32 | The invoice ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoicesIdCommentsAPI.GetV1InvoicesIdComments(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesIdCommentsAPI.GetV1InvoicesIdComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InvoicesIdComments`: SalesDataInvoiceCommentSearchResultInterface
	fmt.Fprintf(os.Stdout, "Response from `InvoicesIdCommentsAPI.GetV1InvoicesIdComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The invoice ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InvoicesIdCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalesDataInvoiceCommentSearchResultInterface**](SalesDataInvoiceCommentSearchResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

