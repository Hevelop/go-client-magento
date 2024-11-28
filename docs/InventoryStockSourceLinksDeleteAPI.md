# \InventoryStockSourceLinksDeleteAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1InventoryStocksourcelinksdelete**](InventoryStockSourceLinksDeleteAPI.md#PostV1InventoryStocksourcelinksdelete) | **Post** /V1/inventory/stock-source-links-delete | inventory/stock-source-links-delete



## PostV1InventoryStocksourcelinksdelete

> ErrorResponse PostV1InventoryStocksourcelinksdelete(ctx).PostV1InventoryStocksourcelinksRequest(postV1InventoryStocksourcelinksRequest).Execute()

inventory/stock-source-links-delete



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
	postV1InventoryStocksourcelinksRequest := *openapiclient.NewPostV1InventoryStocksourcelinksRequest([]openapiclient.InventoryApiDataStockSourceLinkInterface{*openapiclient.NewInventoryApiDataStockSourceLinkInterface()}) // PostV1InventoryStocksourcelinksRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InventoryStockSourceLinksDeleteAPI.PostV1InventoryStocksourcelinksdelete(context.Background()).PostV1InventoryStocksourcelinksRequest(postV1InventoryStocksourcelinksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryStockSourceLinksDeleteAPI.PostV1InventoryStocksourcelinksdelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1InventoryStocksourcelinksdelete`: ErrorResponse
	fmt.Fprintf(os.Stdout, "Response from `InventoryStockSourceLinksDeleteAPI.PostV1InventoryStocksourcelinksdelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostV1InventoryStocksourcelinksdeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postV1InventoryStocksourcelinksRequest** | [**PostV1InventoryStocksourcelinksRequest**](PostV1InventoryStocksourcelinksRequest.md) |  | 

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

