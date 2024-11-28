# \TaxRulesAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1TaxRules**](TaxRulesAPI.md#PostV1TaxRules) | **Post** /V1/taxRules | taxRules
[**PutV1TaxRules**](TaxRulesAPI.md#PutV1TaxRules) | **Put** /V1/taxRules | taxRules



## PostV1TaxRules

> TaxDataTaxRuleInterface PostV1TaxRules(ctx).PutV1TaxRulesRequest(putV1TaxRulesRequest).Execute()

taxRules



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
	putV1TaxRulesRequest := *openapiclient.NewPutV1TaxRulesRequest(*openapiclient.NewTaxDataTaxRuleInterface("Code_example", int32(123), int32(123), []int32{int32(123)}, []int32{int32(123)}, []int32{int32(123)})) // PutV1TaxRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxRulesAPI.PostV1TaxRules(context.Background()).PutV1TaxRulesRequest(putV1TaxRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxRulesAPI.PostV1TaxRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1TaxRules`: TaxDataTaxRuleInterface
	fmt.Fprintf(os.Stdout, "Response from `TaxRulesAPI.PostV1TaxRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1TaxRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putV1TaxRulesRequest** | [**PutV1TaxRulesRequest**](PutV1TaxRulesRequest.md) |  | 

### Return type

[**TaxDataTaxRuleInterface**](TaxDataTaxRuleInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutV1TaxRules

> TaxDataTaxRuleInterface PutV1TaxRules(ctx).PutV1TaxRulesRequest(putV1TaxRulesRequest).Execute()

taxRules



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
	putV1TaxRulesRequest := *openapiclient.NewPutV1TaxRulesRequest(*openapiclient.NewTaxDataTaxRuleInterface("Code_example", int32(123), int32(123), []int32{int32(123)}, []int32{int32(123)}, []int32{int32(123)})) // PutV1TaxRulesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TaxRulesAPI.PutV1TaxRules(context.Background()).PutV1TaxRulesRequest(putV1TaxRulesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TaxRulesAPI.PutV1TaxRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutV1TaxRules`: TaxDataTaxRuleInterface
	fmt.Fprintf(os.Stdout, "Response from `TaxRulesAPI.PutV1TaxRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPutV1TaxRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **putV1TaxRulesRequest** | [**PutV1TaxRulesRequest**](PutV1TaxRulesRequest.md) |  | 

### Return type

[**TaxDataTaxRuleInterface**](TaxDataTaxRuleInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

