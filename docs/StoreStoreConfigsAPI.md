# \StoreStoreConfigsAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetV1StoreStoreConfigs**](StoreStoreConfigsAPI.md#GetV1StoreStoreConfigs) | **Get** /V1/store/storeConfigs | store/storeConfigs



## GetV1StoreStoreConfigs

> []StoreDataStoreConfigInterface GetV1StoreStoreConfigs(ctx).StoreCodes(storeCodes).Execute()

store/storeConfigs



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
	storeCodes := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StoreStoreConfigsAPI.GetV1StoreStoreConfigs(context.Background()).StoreCodes(storeCodes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StoreStoreConfigsAPI.GetV1StoreStoreConfigs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetV1StoreStoreConfigs`: []StoreDataStoreConfigInterface
	fmt.Fprintf(os.Stdout, "Response from `StoreStoreConfigsAPI.GetV1StoreStoreConfigs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetV1StoreStoreConfigsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **storeCodes** | **[]string** |  | 

### Return type

[**[]StoreDataStoreConfigInterface**](StoreDataStoreConfigInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

