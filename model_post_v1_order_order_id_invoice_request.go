/*
Commerce Admin REST endpoints - All inclusive

The schemas documented here are autogenerated from an instance of Adobe Commerce with B2B. Each schema represents a specific user role (Admin, Customer, and Guest) and determines which endpoints are accessible. Use the version switcher to select an Adobe Commerce version and corresponding API.  You can also <a href=\"https://developer.adobe.com/commerce/webapi/rest/quick-reference/generate-local\" target=\"_blank\">generate a local API reference</a> based on your own Adobe Commerce configuration, which allows you to see API documentation for your specific Adobe Commerce modules, third-party modules, and extension attributes.

API version: 2.4.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package magento

import (
	"encoding/json"
)

// checks if the PostV1OrderOrderIdInvoiceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostV1OrderOrderIdInvoiceRequest{}

// PostV1OrderOrderIdInvoiceRequest struct for PostV1OrderOrderIdInvoiceRequest
type PostV1OrderOrderIdInvoiceRequest struct {
	Capture              *bool                                       `json:"capture,omitempty"`
	Items                []SalesDataInvoiceItemCreationInterface     `json:"items,omitempty"`
	Notify               *bool                                       `json:"notify,omitempty"`
	AppendComment        *bool                                       `json:"appendComment,omitempty"`
	Comment              *SalesDataInvoiceCommentCreationInterface   `json:"comment,omitempty"`
	Arguments            *SalesDataInvoiceCreationArgumentsInterface `json:"arguments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PostV1OrderOrderIdInvoiceRequest PostV1OrderOrderIdInvoiceRequest

// NewPostV1OrderOrderIdInvoiceRequest instantiates a new PostV1OrderOrderIdInvoiceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostV1OrderOrderIdInvoiceRequest() *PostV1OrderOrderIdInvoiceRequest {
	this := PostV1OrderOrderIdInvoiceRequest{}
	return &this
}

// NewPostV1OrderOrderIdInvoiceRequestWithDefaults instantiates a new PostV1OrderOrderIdInvoiceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostV1OrderOrderIdInvoiceRequestWithDefaults() *PostV1OrderOrderIdInvoiceRequest {
	this := PostV1OrderOrderIdInvoiceRequest{}
	return &this
}

// GetCapture returns the Capture field value if set, zero value otherwise.
func (o *PostV1OrderOrderIdInvoiceRequest) GetCapture() bool {
	if o == nil || IsNil(o.Capture) {
		var ret bool
		return ret
	}
	return *o.Capture
}

// GetCaptureOk returns a tuple with the Capture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostV1OrderOrderIdInvoiceRequest) GetCaptureOk() (*bool, bool) {
	if o == nil || IsNil(o.Capture) {
		return nil, false
	}
	return o.Capture, true
}

// HasCapture returns a boolean if a field has been set.
func (o *PostV1OrderOrderIdInvoiceRequest) HasCapture() bool {
	if o != nil && !IsNil(o.Capture) {
		return true
	}

	return false
}

// SetCapture gets a reference to the given bool and assigns it to the Capture field.
func (o *PostV1OrderOrderIdInvoiceRequest) SetCapture(v bool) {
	o.Capture = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *PostV1OrderOrderIdInvoiceRequest) GetItems() []SalesDataInvoiceItemCreationInterface {
	if o == nil || IsNil(o.Items) {
		var ret []SalesDataInvoiceItemCreationInterface
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostV1OrderOrderIdInvoiceRequest) GetItemsOk() ([]SalesDataInvoiceItemCreationInterface, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *PostV1OrderOrderIdInvoiceRequest) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []SalesDataInvoiceItemCreationInterface and assigns it to the Items field.
func (o *PostV1OrderOrderIdInvoiceRequest) SetItems(v []SalesDataInvoiceItemCreationInterface) {
	o.Items = v
}

// GetNotify returns the Notify field value if set, zero value otherwise.
func (o *PostV1OrderOrderIdInvoiceRequest) GetNotify() bool {
	if o == nil || IsNil(o.Notify) {
		var ret bool
		return ret
	}
	return *o.Notify
}

// GetNotifyOk returns a tuple with the Notify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostV1OrderOrderIdInvoiceRequest) GetNotifyOk() (*bool, bool) {
	if o == nil || IsNil(o.Notify) {
		return nil, false
	}
	return o.Notify, true
}

// HasNotify returns a boolean if a field has been set.
func (o *PostV1OrderOrderIdInvoiceRequest) HasNotify() bool {
	if o != nil && !IsNil(o.Notify) {
		return true
	}

	return false
}

// SetNotify gets a reference to the given bool and assigns it to the Notify field.
func (o *PostV1OrderOrderIdInvoiceRequest) SetNotify(v bool) {
	o.Notify = &v
}

// GetAppendComment returns the AppendComment field value if set, zero value otherwise.
func (o *PostV1OrderOrderIdInvoiceRequest) GetAppendComment() bool {
	if o == nil || IsNil(o.AppendComment) {
		var ret bool
		return ret
	}
	return *o.AppendComment
}

// GetAppendCommentOk returns a tuple with the AppendComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostV1OrderOrderIdInvoiceRequest) GetAppendCommentOk() (*bool, bool) {
	if o == nil || IsNil(o.AppendComment) {
		return nil, false
	}
	return o.AppendComment, true
}

// HasAppendComment returns a boolean if a field has been set.
func (o *PostV1OrderOrderIdInvoiceRequest) HasAppendComment() bool {
	if o != nil && !IsNil(o.AppendComment) {
		return true
	}

	return false
}

// SetAppendComment gets a reference to the given bool and assigns it to the AppendComment field.
func (o *PostV1OrderOrderIdInvoiceRequest) SetAppendComment(v bool) {
	o.AppendComment = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *PostV1OrderOrderIdInvoiceRequest) GetComment() SalesDataInvoiceCommentCreationInterface {
	if o == nil || IsNil(o.Comment) {
		var ret SalesDataInvoiceCommentCreationInterface
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostV1OrderOrderIdInvoiceRequest) GetCommentOk() (*SalesDataInvoiceCommentCreationInterface, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *PostV1OrderOrderIdInvoiceRequest) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given SalesDataInvoiceCommentCreationInterface and assigns it to the Comment field.
func (o *PostV1OrderOrderIdInvoiceRequest) SetComment(v SalesDataInvoiceCommentCreationInterface) {
	o.Comment = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *PostV1OrderOrderIdInvoiceRequest) GetArguments() SalesDataInvoiceCreationArgumentsInterface {
	if o == nil || IsNil(o.Arguments) {
		var ret SalesDataInvoiceCreationArgumentsInterface
		return ret
	}
	return *o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostV1OrderOrderIdInvoiceRequest) GetArgumentsOk() (*SalesDataInvoiceCreationArgumentsInterface, bool) {
	if o == nil || IsNil(o.Arguments) {
		return nil, false
	}
	return o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *PostV1OrderOrderIdInvoiceRequest) HasArguments() bool {
	if o != nil && !IsNil(o.Arguments) {
		return true
	}

	return false
}

// SetArguments gets a reference to the given SalesDataInvoiceCreationArgumentsInterface and assigns it to the Arguments field.
func (o *PostV1OrderOrderIdInvoiceRequest) SetArguments(v SalesDataInvoiceCreationArgumentsInterface) {
	o.Arguments = &v
}

func (o PostV1OrderOrderIdInvoiceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostV1OrderOrderIdInvoiceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Capture) {
		toSerialize["capture"] = o.Capture
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Notify) {
		toSerialize["notify"] = o.Notify
	}
	if !IsNil(o.AppendComment) {
		toSerialize["appendComment"] = o.AppendComment
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	if !IsNil(o.Arguments) {
		toSerialize["arguments"] = o.Arguments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PostV1OrderOrderIdInvoiceRequest) UnmarshalJSON(data []byte) (err error) {
	varPostV1OrderOrderIdInvoiceRequest := _PostV1OrderOrderIdInvoiceRequest{}

	err = json.Unmarshal(data, &varPostV1OrderOrderIdInvoiceRequest)

	if err != nil {
		return err
	}

	*o = PostV1OrderOrderIdInvoiceRequest(varPostV1OrderOrderIdInvoiceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "capture")
		delete(additionalProperties, "items")
		delete(additionalProperties, "notify")
		delete(additionalProperties, "appendComment")
		delete(additionalProperties, "comment")
		delete(additionalProperties, "arguments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *PostV1OrderOrderIdInvoiceRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *PostV1OrderOrderIdInvoiceRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullablePostV1OrderOrderIdInvoiceRequest struct {
	value *PostV1OrderOrderIdInvoiceRequest
	isSet bool
}

func (v NullablePostV1OrderOrderIdInvoiceRequest) Get() *PostV1OrderOrderIdInvoiceRequest {
	return v.value
}

func (v *NullablePostV1OrderOrderIdInvoiceRequest) Set(val *PostV1OrderOrderIdInvoiceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostV1OrderOrderIdInvoiceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostV1OrderOrderIdInvoiceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostV1OrderOrderIdInvoiceRequest(val *PostV1OrderOrderIdInvoiceRequest) *NullablePostV1OrderOrderIdInvoiceRequest {
	return &NullablePostV1OrderOrderIdInvoiceRequest{value: val, isSet: true}
}

func (v NullablePostV1OrderOrderIdInvoiceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostV1OrderOrderIdInvoiceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}