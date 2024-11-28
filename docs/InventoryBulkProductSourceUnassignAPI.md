# \InventoryBulkProductSourceUnassignAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1InventoryBulkproductsourceunassign**](InventoryBulkProductSourceUnassignAPI.md#PostV1InventoryBulkproductsourceunassign) | **Post** /V1/inventory/bulk-product-source-unassign | inventory/bulk-product-source-unassign



## PostV1InventoryBulkproductsourceunassign

> int32 PostV1InventoryBulkproductsourceunassign(ctx).PostV1InventoryBulkproductsourceassignRequest(postV1InventoryBulkproductsourceassignRequest).Execute()

inventory/bulk-product-source-unassign



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
	resp, r, err := apiClient.InventoryBulkProductSourceUnassignAPI.PostV1InventoryBulkproductsourceunassign(context.Background()).PostV1InventoryBulkproductsourceassignRequest(postV1InventoryBulkproductsourceassignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryBulkProductSourceUnassignAPI.PostV1InventoryBulkproductsourceunassign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1InventoryBulkproductsourceunassign`: int32
	fmt.Fprintf(os.Stdout, "Response from `InventoryBulkProductSourceUnassignAPI.PostV1InventoryBulkproductsourceunassign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1InventoryBulkproductsourceunassignRequest struct via the builder pattern


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

