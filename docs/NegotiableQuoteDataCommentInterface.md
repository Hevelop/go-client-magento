# # NegotiableQuoteDataCommentInterface
Interface CommentInterface

## Properties 


Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId**| **int32** | Comment ID.  |
**ParentId**| **int32** | Negotiable quote ID, that this comment belongs to.  |
**CreatorType**| **int32** | The comment creator type.  |
**IsDecline**| **int32** | Is quote was declined by seller.  |
**IsDraft**| **int32** | Is quote draft flag.  |
**CreatorId**| **int32** | Comment creator ID.  |
**Comment**| **string** | Comment.  |
**CreatedAt**| **string** | Comment created at.  |
**ExtensionAttributes**| **map[string]interface{}** | ExtensionInterface class for @see \\Magento\\NegotiableQuote\\Api\\Data\\CommentInterface  | [optional]
**Attachments**| [**[]NegotiableQuoteDataCommentAttachmentInterface**](NegotiableQuoteDataCommentAttachmentInterface.md) | Existing attachments.  |


[[Back to Model list]](../../README.md#models) [[Back to API list]](../../README.md#endpoints) [[Back to README]](../../README.md)

