# \InventoryBulkPartialSourceTransferAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1InventoryBulkpartialsourcetransfer**](InventoryBulkPartialSourceTransferAPI.md#PostV1InventoryBulkpartialsourcetransfer) | **Post** /V1/inventory/bulk-partial-source-transfer | inventory/bulk-partial-source-transfer



## PostV1InventoryBulkpartialsourcetransfer

> ErrorResponse PostV1InventoryBulkpartialsourcetransfer(ctx).PostV1InventoryBulkpartialsourcetransferRequest(postV1InventoryBulkpartialsourcetransferRequest).Execute()

inventory/bulk-partial-source-transfer



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
	postV1InventoryBulkpartialsourcetransferRequest := *openapiclient.NewPostV1InventoryBulkpartialsourcetransferRequest("OriginSourceCode_example", "DestinationSourceCode_example", []openapiclient.InventoryCatalogApiDataPartialInventoryTransferItemInterface{*openapiclient.NewInventoryCatalogApiDataPartialInventoryTransferItemInterface("Sku_example", float32(123))}) // PostV1InventoryBulkpartialsourcetransferRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryBulkPartialSourceTransferAPI.PostV1InventoryBulkpartialsourcetransfer(context.Background()).PostV1InventoryBulkpartialsourcetransferRequest(postV1InventoryBulkpartialsourcetransferRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryBulkPartialSourceTransferAPI.PostV1InventoryBulkpartialsourcetransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1InventoryBulkpartialsourcetransfer`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `InventoryBulkPartialSourceTransferAPI.PostV1InventoryBulkpartialsourcetransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1InventoryBulkpartialsourcetransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1InventoryBulkpartialsourcetransferRequest** | [**PostV1InventoryBulkpartialsourcetransferRequest**](PostV1InventoryBulkpartialsourcetransferRequest.md) |  | 

### Return type

[**ErrorResponse**](ErrorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

