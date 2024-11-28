# \InventoryGetDistanceProviderCodeAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1InventoryGetdistanceprovidercode**](InventoryGetDistanceProviderCodeAPI.md#GetV1InventoryGetdistanceprovidercode) | **Get** /V1/inventory/get-distance-provider-code | inventory/get-distance-provider-code



## GetV1InventoryGetdistanceprovidercode

> string GetV1InventoryGetdistanceprovidercode(ctx).Execute()

inventory/get-distance-provider-code



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
	resp, r, err := apiClient.InventoryGetDistanceProviderCodeAPI.GetV1InventoryGetdistanceprovidercode(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InventoryGetDistanceProviderCodeAPI.GetV1InventoryGetdistanceprovidercode``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1InventoryGetdistanceprovidercode`: string
	fmt.Fprintf(os.Stdout, "Response from `InventoryGetDistanceProviderCodeAPI.GetV1InventoryGetdistanceprovidercode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetV1InventoryGetdistanceprovidercodeRequest struct via the builder pattern


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

