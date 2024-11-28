# \CreditmemoRefundAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CreditmemoRefund**](CreditmemoRefundAPI.md#PostV1CreditmemoRefund) | **Post** /V1/creditmemo/refund | creditmemo/refund



## PostV1CreditmemoRefund

> SalesDataCreditmemoInterface PostV1CreditmemoRefund(ctx).PostV1CreditmemoRefundRequest(postV1CreditmemoRefundRequest).Execute()

creditmemo/refund



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
	postV1CreditmemoRefundRequest := *openapiclient.NewPostV1CreditmemoRefundRequest(*openapiclient.NewSalesDataCreditmemoInterface(int32(123), []openapiclient.SalesDataCreditmemoItemInterface{*openapiclient.NewSalesDataCreditmemoItemInterface(float32(123), float32(123), int32(123), int32(123), float32(123))})) // PostV1CreditmemoRefundRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditmemoRefundAPI.PostV1CreditmemoRefund(context.Background()).PostV1CreditmemoRefundRequest(postV1CreditmemoRefundRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditmemoRefundAPI.PostV1CreditmemoRefund``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CreditmemoRefund`: SalesDataCreditmemoInterface
	fmt.Fprintf(os.Stdout, "Response from `CreditmemoRefundAPI.PostV1CreditmemoRefund`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CreditmemoRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CreditmemoRefundRequest** | [**PostV1CreditmemoRefundRequest**](PostV1CreditmemoRefundRequest.md) |  | 

### Return type

[**SalesDataCreditmemoInterface**](SalesDataCreditmemoInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

