# \InvoicesIdVoidAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1InvoicesIdVoid**](InvoicesIdVoidAPI.md#PostV1InvoicesIdVoid) | **Post** /V1/invoices/{id}/void | invoices/{id}/void



## PostV1InvoicesIdVoid

> bool PostV1InvoicesIdVoid(ctx, id).Execute()

invoices/{id}/void



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
	resp, r, err := apiClient.InvoicesIdVoidAPI.PostV1InvoicesIdVoid(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoicesIdVoidAPI.PostV1InvoicesIdVoid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1InvoicesIdVoid`: bool
	fmt.Fprintf(os.Stdout, "Response from `InvoicesIdVoidAPI.PostV1InvoicesIdVoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The invoice ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1InvoicesIdVoidRequest struct via the builder pattern


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

