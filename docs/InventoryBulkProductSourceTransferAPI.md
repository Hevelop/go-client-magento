# \InventoryBulkProductSourceTransferAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1InventoryBulkproductsourcetransfer**](InventoryBulkProductSourceTransferAPI.md#PostV1InventoryBulkproductsourcetransfer) | **Post** /V1/inventory/bulk-product-source-transfer | inventory/bulk-product-source-transfer



## PostV1InventoryBulkproductsourcetransfer

> bool PostV1InventoryBulkproductsourcetransfer(ctx).PostV1InventoryBulkproductsourcetransferRequest(postV1InventoryBulkproductsourcetransferRequest).Execute()

inventory/bulk-product-source-transfer



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
	postV1InventoryBulkproductsourcetransferRequest := *openapiclient.NewPostV1InventoryBulkproductsourcetransferRequest([]string{"Skus_example"}, "OriginSource_example", "DestinationSource_example", false) // PostV1InventoryBulkproductsourcetransferRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryBulkProductSourceTransferAPI.PostV1InventoryBulkproductsourcetransfer(context.Background()).PostV1InventoryBulkproductsourcetransferRequest(postV1InventoryBulkproductsourcetransferRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryBulkProductSourceTransferAPI.PostV1InventoryBulkproductsourcetransfer``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1InventoryBulkproductsourcetransfer`: bool
	fmt.Fprintf(os.Stdout, "Response from `InventoryBulkProductSourceTransferAPI.PostV1InventoryBulkproductsourcetransfer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1InventoryBulkproductsourcetransferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1InventoryBulkproductsourcetransferRequest** | [**PostV1InventoryBulkproductsourcetransferRequest**](PostV1InventoryBulkproductsourcetransferRequest.md) |  | 

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

