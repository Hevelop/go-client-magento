# # CustomerDataAttributeMetadataInterface
Customer attribute metadata interface.

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FrontendInput**| **string** | HTML for input element.  |
**InputFilter**| **string** | Template used for input (e.g. \&quot;date\&quot;)  |
**StoreLabel**| **string** | Label of the store.  |
**ValidationRules**| [**[]CustomerDataValidationRuleInterface**](CustomerDataValidationRuleInterface.md) | Validation rules.  |
**MultilineCount**| **int32** | Of lines of the attribute value.  |
**Visible**| **bool** | Attribute is visible on frontend.  |
**Required**| **bool** | Attribute is required.  |
**DataModel**| **string** | Data model for attribute.  |
**Options**| [**[]CustomerDataOptionInterface**](CustomerDataOptionInterface.md) | Options of the attribute (key &#x3D;&gt; value pairs for select)  |
**FrontendClass**| **string** | Class which is used to display the attribute on frontend.  |
**UserDefined**| **bool** | Current attribute has been defined by a user.  |
**SortOrder**| **int32** | Attributes sort order.  |
**FrontendLabel**| **string** | Label which supposed to be displayed on frontend.  |
**Note**| **string** | The note attribute for the element.  |
**System**| **bool** | This is a system attribute.  |
**BackendType**| **string** | Backend type.  |
**IsUsedInGrid**| **bool** | It is used in customer grid  | [optional]
**IsVisibleInGrid**| **bool** | It is visible in customer grid  | [optional]
**IsFilterableInGrid**| **bool** | It is filterable in customer grid  | [optional]
**IsSearchableInGrid**| **bool** | It is searchable in customer grid  | [optional]
**AttributeCode**| **string** | Code of the attribute.  |


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

