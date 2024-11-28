# # InventoryInStorePickupApiDataSearchRequestInterface
Endpoint used to search Pickup Locations by different parameters: - by attribute filters fields @see \\Magento\\InventoryInStorePickupApi\\Api\\Data\\SearchRequest\\FiltersInterface - by distance to the address @see \\Magento\\InventoryInStorePickupApi\\Api\\Data\\SearchRequest\\AreaInterface Also, endpoint supports paging and sort orders.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Area**| [**InventoryInStorePickupApiDataSearchRequestAreaInterface**](InventoryInStorePickupApiDataSearchRequestAreaInterface.md) |   | [optional]
**Filters**| [**InventoryInStorePickupApiDataSearchRequestFiltersInterface**](InventoryInStorePickupApiDataSearchRequestFiltersInterface.md) |   | [optional]
**PageSize**| **int32** | Page size.  | [optional]
**CurrentPage**| **int32** | Current page.  |
**ScopeType**| **string** | Sales Channel Type.  |
**ScopeCode**| **string** | Sales Channel code.  |
**Sort**| [**[]FrameworkSortOrder**](FrameworkSortOrder.md) | Sort Order.  | [optional]
**ExtensionAttributes**| [**InventoryInStorePickupApiDataSearchRequestExtensionInterface**](InventoryInStorePickupApiDataSearchRequestExtensionInterface.md) |   | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

