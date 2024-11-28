# \InventorySourceSelectionAlgorithmListAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventorySourceselectionalgorithmlist**](InventorySourceSelectionAlgorithmListAPI.md#GetV1InventorySourceselectionalgorithmlist) | **Get** /V1/inventory/source-selection-algorithm-list | inventory/source-selection-algorithm-list



## GetV1InventorySourceselectionalgorithmlist

> []InventorySourceSelectionApiDataSourceSelectionAlgorithmInterface GetV1InventorySourceselectionalgorithmlist(ctx).Execute()

inventory/source-selection-algorithm-list



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventorySourceSelectionAlgorithmListAPI.GetV1InventorySourceselectionalgorithmlist(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventorySourceSelectionAlgorithmListAPI.GetV1InventorySourceselectionalgorithmlist``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventorySourceselectionalgorithmlist`: []InventorySourceSelectionApiDataSourceSelectionAlgorithmInterface
	fmt.Fprintf(os.Stdout, "Response from `InventorySourceSelectionAlgorithmListAPI.GetV1InventorySourceselectionalgorithmlist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventorySourceselectionalgorithmlistRequest struct via the builder pattern


### Return type

[**[]InventorySourceSelectionApiDataSourceSelectionAlgorithmInterface**](InventorySourceSelectionApiDataSourceSelectionAlgorithmInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

