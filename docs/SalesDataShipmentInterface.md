# # SalesDataShipmentInterface
Shipment interface. A shipment is a delivery package that contains products. A shipment document accompanies the shipment. This document lists the products and their quantities in the delivery package.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingAddressId**| **int32** | Billing address ID.  | [optional]
**CreatedAt**| **string** | Created-at timestamp.  | [optional]
**CustomerId**| **int32** | Customer ID.  | [optional]
**EmailSent**| **int32** | Email-sent flag value.  | [optional]
**EntityId**| **int32** | Shipment ID.  | [optional]
**IncrementId**| **string** | Increment ID.  | [optional]
**OrderId**| **int32** | Order ID.  |
**Packages**| [**[]SalesDataShipmentPackageInterface**](SalesDataShipmentPackageInterface.md) | Array of packages, if any. Otherwise, null.  | [optional]
**ShipmentStatus**| **int32** | Shipment status.  | [optional]
**ShippingAddressId**| **int32** | Shipping address ID.  | [optional]
**ShippingLabel**| **string** | Shipping label.  | [optional]
**StoreId**| **int32** | Store ID.  | [optional]
**TotalQty**| **float32** | Total quantity.  | [optional]
**TotalWeight**| **float32** | Total weight.  | [optional]
**UpdatedAt**| **string** | Updated-at timestamp.  | [optional]
**Items**| [**[]SalesDataShipmentItemInterface**](SalesDataShipmentItemInterface.md) | Array of items.  |
**Tracks**| [**[]SalesDataShipmentTrackInterface**](SalesDataShipmentTrackInterface.md) | Array of tracks.  |
**Comments**| [**[]SalesDataShipmentCommentInterface**](SalesDataShipmentCommentInterface.md) | Array of comments.  |
**ExtensionAttributes**| [**SalesDataShipmentExtensionInterface**](SalesDataShipmentExtensionInterface.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

