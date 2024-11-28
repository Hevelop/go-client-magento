# # CatalogDataCategoryInterface
Category data interface.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id**| **int32** | Category id.  | [optional]
**ParentId**| **int32** | Parent category ID  | [optional]
**Name**| **string** | Category name  | [optional]
**IsActive**| **bool** | Whether category is active  | [optional]
**Position**| **int32** | Category position  | [optional]
**Level**| **int32** | Category level  | [optional]
**Children**| **string** | Children ids comma separated.  | [optional]
**CreatedAt**| **string** | Category creation date and time.  | [optional]
**UpdatedAt**| **string** | Category last update date and time.  | [optional]
**Path**| **string** | Category full path.  | [optional]
**AvailableSortBy**| **[]string** | Available sort by for category.  | [optional]
**IncludeInMenu**| **bool** | Category is included in menu.  | [optional]
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\Catalog\\Api\\Data\\CategoryInterface  | [optional]
**CustomAttributes**| [**[]FrameworkAttributeInterface**](FrameworkAttributeInterface.md) | Custom attributes values.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

