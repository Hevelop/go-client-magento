# \CartsMineBalanceApplyAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1CartsMineBalanceApply**](CartsMineBalanceApplyAPI.md#PostV1CartsMineBalanceApply) | **Post** /V1/carts/mine/balance/apply | carts/mine/balance/apply



## PostV1CartsMineBalanceApply

> bool PostV1CartsMineBalanceApply(ctx).Execute()

carts/mine/balance/apply



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
	resp, r, err := apiClient.CartsMineBalanceApplyAPI.PostV1CartsMineBalanceApply(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CartsMineBalanceApplyAPI.PostV1CartsMineBalanceApply``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1CartsMineBalanceApply`: bool
	fmt.Fprintf(os.Stdout, "Response from `CartsMineBalanceApplyAPI.PostV1CartsMineBalanceApply`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1CartsMineBalanceApplyRequest struct via the builder pattern


### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

