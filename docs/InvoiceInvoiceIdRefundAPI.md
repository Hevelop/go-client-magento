# \InvoiceInvoiceIdRefundAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1InvoiceInvoiceIdRefund**](InvoiceInvoiceIdRefundAPI.md#PostV1InvoiceInvoiceIdRefund) | **Post** /V1/invoice/{invoiceId}/refund | invoice/{invoiceId}/refund



## PostV1InvoiceInvoiceIdRefund

> int32 PostV1InvoiceInvoiceIdRefund(ctx, invoiceId).PostV1InvoiceInvoiceIdRefundRequest(postV1InvoiceInvoiceIdRefundRequest).Execute()

invoice/{invoiceId}/refund



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
	invoiceId := int32(56) // int32 | 
	postV1InvoiceInvoiceIdRefundRequest := *openapiclient.NewPostV1InvoiceInvoiceIdRefundRequest() // PostV1InvoiceInvoiceIdRefundRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceInvoiceIdRefundAPI.PostV1InvoiceInvoiceIdRefund(context.Background(), invoiceId).PostV1InvoiceInvoiceIdRefundRequest(postV1InvoiceInvoiceIdRefundRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceInvoiceIdRefundAPI.PostV1InvoiceInvoiceIdRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1InvoiceInvoiceIdRefund`: int32
	fmt.Fprintf(os.Stdout, "Response from `InvoiceInvoiceIdRefundAPI.PostV1InvoiceInvoiceIdRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**invoiceId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1InvoiceInvoiceIdRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postV1InvoiceInvoiceIdRefundRequest** | [**PostV1InvoiceInvoiceIdRefundRequest**](PostV1InvoiceInvoiceIdRefundRequest.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

