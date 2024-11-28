# \InventorySourceSelectionAlgorithmResultAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1InventorySourceselectionalgorithmresult**](InventorySourceSelectionAlgorithmResultAPI.md#PostV1InventorySourceselectionalgorithmresult) | **Post** /V1/inventory/source-selection-algorithm-result | inventory/source-selection-algorithm-result



## PostV1InventorySourceselectionalgorithmresult

> InventorySourceSelectionApiDataSourceSelectionResultInterface PostV1InventorySourceselectionalgorithmresult(ctx).PostV1InventorySourceselectionalgorithmresultRequest(postV1InventorySourceselectionalgorithmresultRequest).Execute()

inventory/source-selection-algorithm-result



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
	postV1InventorySourceselectionalgorithmresultRequest := *openapiclient.NewPostV1InventorySourceselectionalgorithmresultRequest(*openapiclient.NewInventorySourceSelectionApiDataInventoryRequestInterface(int32(123), []openapiclient.InventorySourceSelectionApiDataItemRequestInterface{*openapiclient.NewInventorySourceSelectionApiDataItemRequestInterface("Sku_example", float32(123))}), "AlgorithmCode_example") // PostV1InventorySourceselectionalgorithmresultRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventorySourceSelectionAlgorithmResultAPI.PostV1InventorySourceselectionalgorithmresult(context.Background()).PostV1InventorySourceselectionalgorithmresultRequest(postV1InventorySourceselectionalgorithmresultRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventorySourceSelectionAlgorithmResultAPI.PostV1InventorySourceselectionalgorithmresult``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1InventorySourceselectionalgorithmresult`: InventorySourceSelectionApiDataSourceSelectionResultInterface
	fmt.Fprintf(os.Stdout, "Response from `InventorySourceSelectionAlgorithmResultAPI.PostV1InventorySourceselectionalgorithmresult`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1InventorySourceselectionalgorithmresultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1InventorySourceselectionalgorithmresultRequest** | [**PostV1InventorySourceselectionalgorithmresultRequest**](PostV1InventorySourceselectionalgorithmresultRequest.md) |  | 

### Return type

[**InventorySourceSelectionApiDataSourceSelectionResultInterface**](InventorySourceSelectionApiDataSourceSelectionResultInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, application/xml
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

