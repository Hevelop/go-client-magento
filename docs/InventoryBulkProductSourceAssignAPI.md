# \InventoryBulkProductSourceAssignAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1InventoryBulkproductsourceassign**](InventoryBulkProductSourceAssignAPI.md#PostV1InventoryBulkproductsourceassign) | **Post** /V1/inventory/bulk-product-source-assign | inventory/bulk-product-source-assign



## PostV1InventoryBulkproductsourceassign

> int32 PostV1InventoryBulkproductsourceassign(ctx).PostV1InventoryBulkproductsourceassignRequest(postV1InventoryBulkproductsourceassignRequest).Execute()

inventory/bulk-product-source-assign



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
	postV1InventoryBulkproductsourceassignRequest := *openapiclient.NewPostV1InventoryBulkproductsourceassignRequest([]string{"Skus_example"}, []string{"SourceCodes_example"}) // PostV1InventoryBulkproductsourceassignRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryBulkProductSourceAssignAPI.PostV1InventoryBulkproductsourceassign(context.Background()).PostV1InventoryBulkproductsourceassignRequest(postV1InventoryBulkproductsourceassignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryBulkProductSourceAssignAPI.PostV1InventoryBulkproductsourceassign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1InventoryBulkproductsourceassign`: int32
	fmt.Fprintf(os.Stdout, "Response from `InventoryBulkProductSourceAssignAPI.PostV1InventoryBulkproductsourceassign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1InventoryBulkproductsourceassignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1InventoryBulkproductsourceassignRequest** | [**PostV1InventoryBulkproductsourceassignRequest**](PostV1InventoryBulkproductsourceassignRequest.md) |  | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

