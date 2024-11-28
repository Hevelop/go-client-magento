# \SalesRulesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1SalesRules**](SalesRulesAPI.md#PostV1SalesRules) | **Post** /V1/salesRules | salesRules



## PostV1SalesRules

> SalesRuleDataRuleInterface PostV1SalesRules(ctx).PostV1SalesRulesRequest(postV1SalesRulesRequest).Execute()

salesRules



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
	postV1SalesRulesRequest := *openapiclient.NewPostV1SalesRulesRequest(*openapiclient.NewSalesRuleDataRuleInterface([]int32{int32(123)}, []int32{int32(123)}, int32(123), false, false, false, int32(123), float32(123), int32(123), false, int32(123), false, "CouponType_example", false, int32(123))) // PostV1SalesRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesRulesAPI.PostV1SalesRules(context.Background()).PostV1SalesRulesRequest(postV1SalesRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesRulesAPI.PostV1SalesRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1SalesRules`: SalesRuleDataRuleInterface
	fmt.Fprintf(os.Stdout, "Response from `SalesRulesAPI.PostV1SalesRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1SalesRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1SalesRulesRequest** | [**PostV1SalesRulesRequest**](PostV1SalesRulesRequest.md) |  | 

### Return type

[**SalesRuleDataRuleInterface**](SalesRuleDataRuleInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

