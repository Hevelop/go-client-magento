# # SalesDataShipmentTrackInterface
Shipment track interface. A shipment is a delivery package that contains products. A shipment document accompanies the shipment. This document lists the products and their quantities in the delivery package. Merchants and customers can track shipments.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId**| **int32** | The order_id for the shipment package.  |
**CreatedAt**| **string** | Created-at timestamp.  | [optional]
**EntityId**| **int32** | Shipment package ID.  | [optional]
**ParentId**| **int32** | Parent ID.  |
**UpdatedAt**| **string** | Updated-at timestamp.  | [optional]
**Weight**| **float32** | Weight.  |
**Qty**| **float32** | Quantity.  |
**Description**| **string** | Description.  |
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Sales\\Api\\Data\\ShipmentTrackInterface  | [optional]
**TrackNumber**| **string** | Track number.  |
**Title**| **string** | Title.  |
**CarrierCode**| **string** | Carrier code.  |


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

