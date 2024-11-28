# \RewardMineUseRewardAPI

All URIs are relative to *http://example.com/rest/default*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PostV1RewardMineUsereward**](RewardMineUseRewardAPI.md#PostV1RewardMineUsereward) | **Post** /V1/reward/mine/use-reward | reward/mine/use-reward



## PostV1RewardMineUsereward

> bool PostV1RewardMineUsereward(ctx).Execute()

reward/mine/use-reward



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
	resp, r, err := apiClient.RewardMineUseRewardAPI.PostV1RewardMineUsereward(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RewardMineUseRewardAPI.PostV1RewardMineUsereward``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostV1RewardMineUsereward`: bool
	fmt.Fprintf(os.Stdout, "Response from `RewardMineUseRewardAPI.PostV1RewardMineUsereward`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPostV1RewardMineUserewardRequest struct via the builder pattern


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

