# \SalesRulesRuleIdAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteV1SalesRulesRuleId**](SalesRulesRuleIdAPI.md#DeleteV1SalesRulesRuleId) | **Delete** /V1/salesRules/{ruleId} | salesRules/{ruleId}
[**GetV1SalesRulesRuleId**](SalesRulesRuleIdAPI.md#GetV1SalesRulesRuleId) | **Get** /V1/salesRules/{ruleId} | salesRules/{ruleId}
[**PutV1SalesRulesRuleId**](SalesRulesRuleIdAPI.md#PutV1SalesRulesRuleId) | **Put** /V1/salesRules/{ruleId} | salesRules/{ruleId}



## DeleteV1SalesRulesRuleId

> bool DeleteV1SalesRulesRuleId(ctx, ruleId).Execute()

salesRules/{ruleId}



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
	ruleId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesRulesRuleIdAPI.DeleteV1SalesRulesRuleId(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesRulesRuleIdAPI.DeleteV1SalesRulesRuleId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteV1SalesRulesRuleId`: bool
	fmt.Fprintf(os.Stdout, "Response from `SalesRulesRuleIdAPI.DeleteV1SalesRulesRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteV1SalesRulesRuleIdRequest struct via the builder pattern


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


## GetV1SalesRulesRuleId

> SalesRuleDataRuleInterface GetV1SalesRulesRuleId(ctx, ruleId).Execute()

salesRules/{ruleId}



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
	ruleId := int32(56) // int32 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesRulesRuleIdAPI.GetV1SalesRulesRuleId(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesRulesRuleIdAPI.GetV1SalesRulesRuleId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1SalesRulesRuleId`: SalesRuleDataRuleInterface
	fmt.Fprintf(os.Stdout, "Response from `SalesRulesRuleIdAPI.GetV1SalesRulesRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1SalesRulesRuleIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SalesRuleDataRuleInterface**](SalesRuleDataRuleInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1SalesRulesRuleId

> SalesRuleDataRuleInterface PutV1SalesRulesRuleId(ctx, ruleId).PostV1SalesRulesRequest(postV1SalesRulesRequest).Execute()

salesRules/{ruleId}



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
	ruleId := "ruleId_example" // string | 
	postV1SalesRulesRequest := *openapiclient.NewPostV1SalesRulesRequest(*openapiclient.NewSalesRuleDataRuleInterface([]int32{int32(123)}, []int32{int32(123)}, int32(123), false, false, false, int32(123), float32(123), int32(123), false, int32(123), false, "CouponType_example", false, int32(123))) // PostV1SalesRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SalesRulesRuleIdAPI.PutV1SalesRulesRuleId(context.Background(), ruleId).PostV1SalesRulesRequest(postV1SalesRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SalesRulesRuleIdAPI.PutV1SalesRulesRuleId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1SalesRulesRuleId`: SalesRuleDataRuleInterface
	fmt.Fprintf(os.Stdout, "Response from `SalesRulesRuleIdAPI.PutV1SalesRulesRuleId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutV1SalesRulesRuleIdRequest struct via the builder pattern


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

