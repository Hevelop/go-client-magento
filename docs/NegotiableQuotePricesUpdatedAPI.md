# \NegotiableQuotePricesUpdatedAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1NegotiableQuotePricesUpdated**](NegotiableQuotePricesUpdatedAPI.md#PostV1NegotiableQuotePricesUpdated) | **Post** /V1/negotiableQuote/pricesUpdated | negotiableQuote/pricesUpdated



## PostV1NegotiableQuotePricesUpdated

> bool PostV1NegotiableQuotePricesUpdated(ctx).PostV1NegotiableQuotePricesUpdatedRequest(postV1NegotiableQuotePricesUpdatedRequest).Execute()

negotiableQuote/pricesUpdated



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
	postV1NegotiableQuotePricesUpdatedRequest := *openapiclient.NewPostV1NegotiableQuotePricesUpdatedRequest([]int32{int32(123)}) // PostV1NegotiableQuotePricesUpdatedRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NegotiableQuotePricesUpdatedAPI.PostV1NegotiableQuotePricesUpdated(context.Background()).PostV1NegotiableQuotePricesUpdatedRequest(postV1NegotiableQuotePricesUpdatedRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NegotiableQuotePricesUpdatedAPI.PostV1NegotiableQuotePricesUpdated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1NegotiableQuotePricesUpdated`: bool
	fmt.Fprintf(os.Stdout, "Response from `NegotiableQuotePricesUpdatedAPI.PostV1NegotiableQuotePricesUpdated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1NegotiableQuotePricesUpdatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1NegotiableQuotePricesUpdatedRequest** | [**PostV1NegotiableQuotePricesUpdatedRequest**](PostV1NegotiableQuotePricesUpdatedRequest.md) |  | 

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

