# \CreditmemoAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1Creditmemo**](CreditmemoAPI.md#PostV1Creditmemo) | **Post** /V1/creditmemo | creditmemo



## PostV1Creditmemo

> SalesDataCreditmemoInterface PostV1Creditmemo(ctx).PostV1CreditmemoRequest(postV1CreditmemoRequest).Execute()

creditmemo



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
	postV1CreditmemoRequest := *openapiclient.NewPostV1CreditmemoRequest(*openapiclient.NewSalesDataCreditmemoInterface(int32(123), []openapiclient.SalesDataCreditmemoItemInterface{*openapiclient.NewSalesDataCreditmemoItemInterface(float32(123), float32(123), int32(123), int32(123), float32(123))})) // PostV1CreditmemoRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CreditmemoAPI.PostV1Creditmemo(context.Background()).PostV1CreditmemoRequest(postV1CreditmemoRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreditmemoAPI.PostV1Creditmemo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1Creditmemo`: SalesDataCreditmemoInterface
	fmt.Fprintf(os.Stdout, "Response from `CreditmemoAPI.PostV1Creditmemo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CreditmemoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CreditmemoRequest** | [**PostV1CreditmemoRequest**](PostV1CreditmemoRequest.md) |  | 

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

