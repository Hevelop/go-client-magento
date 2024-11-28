# \InventoryLowQuantityNotificationSourceCodeSkuAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventoryLowquantitynotificationSourceCodeSku**](InventoryLowQuantityNotificationSourceCodeSkuAPI.md#GetV1InventoryLowquantitynotificationSourceCodeSku) | **Get** /V1/inventory/low-quantity-notification/{sourceCode}/{sku} | inventory/low-quantity-notification/{sourceCode}/{sku}



## GetV1InventoryLowquantitynotificationSourceCodeSku

> InventoryLowQuantityNotificationApiDataSourceItemConfigurationInterface GetV1InventoryLowquantitynotificationSourceCodeSku(ctx, sourceCode, sku).Execute()

inventory/low-quantity-notification/{sourceCode}/{sku}



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
	sourceCode := "sourceCode_example" // string | 
	sku := "sku_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryLowQuantityNotificationSourceCodeSkuAPI.GetV1InventoryLowquantitynotificationSourceCodeSku(context.Background(), sourceCode, sku).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryLowQuantityNotificationSourceCodeSkuAPI.GetV1InventoryLowquantitynotificationSourceCodeSku``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventoryLowquantitynotificationSourceCodeSku`: InventoryLowQuantityNotificationApiDataSourceItemConfigurationInterface
	fmt.Fprintf(os.Stdout, "Response from `InventoryLowQuantityNotificationSourceCodeSkuAPI.GetV1InventoryLowquantitynotificationSourceCodeSku`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceCode** | **string** |  | 
**sku** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventoryLowquantitynotificationSourceCodeSkuRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**InventoryLowQuantityNotificationApiDataSourceItemConfigurationInterface**](InventoryLowQuantityNotificationApiDataSourceItemConfigurationInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

