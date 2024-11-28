# # RmaDataRmaInterface
Interface RmaInterface

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncrementId**| **string** | Entity_id  |
**EntityId**| **int32** | Entity_id  |
**OrderId**| **int32** | Order_id  |
**OrderIncrementId**| **string** | Order_increment_id  |
**StoreId**| **int32** | Store_id  |
**CustomerId**| **int32** | Customer_id  |
**DateRequested**| **string** | Date_requested  |
**CustomerCustomEmail**| **string** | Customer_custom_email  |
**Items**| [**[]RmaDataItemInterface**](RmaDataItemInterface.md) | Items  |
**Status**| **string** | Status  |
**Comments**| [**[]RmaDataCommentInterface**](RmaDataCommentInterface.md) | Comments list  |
**Tracks**| [**[]RmaDataTrackInterface**](RmaDataTrackInterface.md) | Tracks list  |
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Rma\\Api\\Data\\RmaInterface  | [optional]
**CustomAttributes**| [**[]FrameworkAttributeInterface**](FrameworkAttributeInterface.md) | Custom attributes values.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

