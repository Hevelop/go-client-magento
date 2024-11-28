# # CatalogDataProductAttributeInterface


## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExtensionAttributes**| [**EavDataAttributeExtensionInterface**](EavDataAttributeExtensionInterface.md) |   | [optional]
**IsWysiwygEnabled**| **bool** | WYSIWYG flag  | [optional]
**IsHtmlAllowedOnFront**| **bool** | The HTML tags are allowed on the frontend  | [optional]
**UsedForSortBy**| **bool** | It is used for sorting in product listing  | [optional]
**IsFilterable**| **bool** | It used in layered navigation  | [optional]
**IsFilterableInSearch**| **bool** | It is used in search results layered navigation  | [optional]
**IsUsedInGrid**| **bool** | It is used in catalog product grid  | [optional]
**IsVisibleInGrid**| **bool** | It is visible in catalog product grid  | [optional]
**IsFilterableInGrid**| **bool** | It is filterable in catalog product grid  | [optional]
**Position**| **int32** | Position  | [optional]
**ApplyTo**| **[]string** | Apply to value for the element  | [optional]
**IsSearchable**| **string** | The attribute can be used in Quick Search  | [optional]
**IsVisibleInAdvancedSearch**| **string** | The attribute can be used in Advanced Search  | [optional]
**IsComparable**| **string** | The attribute can be compared on the frontend  | [optional]
**IsUsedForPromoRules**| **string** | The attribute can be used for promo rules  | [optional]
**IsVisibleOnFront**| **string** | The attribute is visible on the frontend  | [optional]
**UsedInProductListing**| **string** | The attribute can be used in product listing  | [optional]
**IsVisible**| **bool** | Attribute is visible on frontend.  | [optional]
**Scope**| **string** | Attribute scope  | [optional]
**AttributeId**| **int32** | Id of the attribute.  | [optional]
**AttributeCode**| **string** | Code of the attribute.  |
**FrontendInput**| **string** | HTML for input element.  |
**EntityTypeId**| **string** | Entity type id  |
**IsRequired**| **bool** | Attribute is required.  |
**Options**| [**[]EavDataAttributeOptionInterface**](EavDataAttributeOptionInterface.md) | Options of the attribute (key &#x3D;&gt; value pairs for select)  | [optional]
**IsUserDefined**| **bool** | Current attribute has been defined by a user.  | [optional]
**DefaultFrontendLabel**| **string** | Frontend label for default store  | [optional]
**FrontendLabels**| [**[]EavDataAttributeFrontendLabelInterface**](EavDataAttributeFrontendLabelInterface.md) | Frontend label for each store  |
**Note**| **string** | The note attribute for the element.  | [optional]
**BackendType**| **string** | Backend type.  | [optional]
**BackendModel**| **string** | Backend model  | [optional]
**SourceModel**| **string** | Source model  | [optional]
**DefaultValue**| **string** | Default value for the element.  | [optional]
**IsUnique**| **string** | This is a unique attribute  | [optional]
**FrontendClass**| **string** | Frontend class of attribute  | [optional]
**ValidationRules**| [**[]EavDataAttributeValidationRuleInterface**](EavDataAttributeValidationRuleInterface.md) | Validation rules.  | [optional]
**CustomAttributes**| [**[]FrameworkAttributeInterface**](FrameworkAttributeInterface.md) | Custom attributes values.  | [optional]


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

