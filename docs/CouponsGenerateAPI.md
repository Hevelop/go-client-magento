# \CouponsGenerateAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CouponsGenerate**](CouponsGenerateAPI.md#PostV1CouponsGenerate) | **Post** /V1/coupons/generate | coupons/generate



## PostV1CouponsGenerate

> []string PostV1CouponsGenerate(ctx).PostV1CouponsGenerateRequest(postV1CouponsGenerateRequest).Execute()

coupons/generate



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
	postV1CouponsGenerateRequest := *openapiclient.NewPostV1CouponsGenerateRequest(*openapiclient.NewSalesRuleDataCouponGenerationSpecInterface(int32(123), "Format_example", int32(123), int32(123))) // PostV1CouponsGenerateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CouponsGenerateAPI.PostV1CouponsGenerate(context.Background()).PostV1CouponsGenerateRequest(postV1CouponsGenerateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CouponsGenerateAPI.PostV1CouponsGenerate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CouponsGenerate`: []string
	fmt.Fprintf(os.Stdout, "Response from `CouponsGenerateAPI.PostV1CouponsGenerate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CouponsGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1CouponsGenerateRequest** | [**PostV1CouponsGenerateRequest**](PostV1CouponsGenerateRequest.md) |  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

