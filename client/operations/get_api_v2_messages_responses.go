// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetAPIV2MessagesReader is a Reader for the GetAPIV2Messages structure.
type GetAPIV2MessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV2MessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAPIV2MessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAPIV2MessagesOK creates a GetAPIV2MessagesOK with default headers values
func NewGetAPIV2MessagesOK() *GetAPIV2MessagesOK {
	return &GetAPIV2MessagesOK{}
}

/* GetAPIV2MessagesOK describes a response with status code 200, with default header values.

Successful response
*/
type GetAPIV2MessagesOK struct {
	Payload *GetAPIV2MessagesOKBody
}

func (o *GetAPIV2MessagesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/messages][%d] getApiV2MessagesOK  %+v", 200, o.Payload)
}
func (o *GetAPIV2MessagesOK) GetPayload() *GetAPIV2MessagesOKBody {
	return o.Payload
}

func (o *GetAPIV2MessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetAPIV2MessagesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetAPIV2MessagesOKBody Messages
swagger:model GetAPIV2MessagesOKBody
*/
type GetAPIV2MessagesOKBody struct {

	// Number of returned messages
	Count int64 `json:"count,omitempty"`

	// messages
	Messages []*GetAPIV2MessagesOKBodyMessagesItems0 `json:"messages"`

	// Start index of first returned message
	Start int64 `json:"start,omitempty"`

	// Total number of stored messages
	Total int64 `json:"total,omitempty"`
}

// Validate validates this get API v2 messages o k body
func (o *GetAPIV2MessagesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAPIV2MessagesOKBody) validateMessages(formats strfmt.Registry) error {
	if swag.IsZero(o.Messages) { // not required
		return nil
	}

	for i := 0; i < len(o.Messages); i++ {
		if swag.IsZero(o.Messages[i]) { // not required
			continue
		}

		if o.Messages[i] != nil {
			if err := o.Messages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getApiV2MessagesOK" + "." + "messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get API v2 messages o k body based on the context it is used
func (o *GetAPIV2MessagesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMessages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAPIV2MessagesOKBody) contextValidateMessages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Messages); i++ {

		if o.Messages[i] != nil {
			if err := o.Messages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getApiV2MessagesOK" + "." + "messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAPIV2MessagesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAPIV2MessagesOKBody) UnmarshalBinary(b []byte) error {
	var res GetAPIV2MessagesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAPIV2MessagesOKBodyMessagesItems0 Message
swagger:model GetAPIV2MessagesOKBodyMessagesItems0
*/
type GetAPIV2MessagesOKBodyMessagesItems0 struct {

	// created
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// from
	From *GetAPIV2MessagesOKBodyMessagesItems0From `json:"from,omitempty"`

	// headers
	Headers interface{} `json:"headers,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// size
	Size int64 `json:"size,omitempty"`

	// to
	To []*GetAPIV2MessagesOKBodyMessagesItems0ToItems0 `json:"to"`
}

// Validate validates this get API v2 messages o k body messages items0
func (o *GetAPIV2MessagesOKBodyMessagesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAPIV2MessagesOKBodyMessagesItems0) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(o.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", o.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetAPIV2MessagesOKBodyMessagesItems0) validateFrom(formats strfmt.Registry) error {
	if swag.IsZero(o.From) { // not required
		return nil
	}

	if o.From != nil {
		if err := o.From.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (o *GetAPIV2MessagesOKBodyMessagesItems0) validateTo(formats strfmt.Registry) error {
	if swag.IsZero(o.To) { // not required
		return nil
	}

	for i := 0; i < len(o.To); i++ {
		if swag.IsZero(o.To[i]) { // not required
			continue
		}

		if o.To[i] != nil {
			if err := o.To[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("to" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get API v2 messages o k body messages items0 based on the context it is used
func (o *GetAPIV2MessagesOKBodyMessagesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFrom(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAPIV2MessagesOKBodyMessagesItems0) contextValidateFrom(ctx context.Context, formats strfmt.Registry) error {

	if o.From != nil {
		if err := o.From.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("from")
			}
			return err
		}
	}

	return nil
}

func (o *GetAPIV2MessagesOKBodyMessagesItems0) contextValidateTo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.To); i++ {

		if o.To[i] != nil {
			if err := o.To[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("to" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetAPIV2MessagesOKBodyMessagesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAPIV2MessagesOKBodyMessagesItems0) UnmarshalBinary(b []byte) error {
	var res GetAPIV2MessagesOKBodyMessagesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAPIV2MessagesOKBodyMessagesItems0From Path
swagger:model GetAPIV2MessagesOKBodyMessagesItems0From
*/
type GetAPIV2MessagesOKBodyMessagesItems0From struct {

	// domain
	Domain string `json:"domain,omitempty"`

	// mailbox
	Mailbox string `json:"mailbox,omitempty"`

	// params
	Params string `json:"params,omitempty"`

	// relays
	Relays []string `json:"relays"`
}

// Validate validates this get API v2 messages o k body messages items0 from
func (o *GetAPIV2MessagesOKBodyMessagesItems0From) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get API v2 messages o k body messages items0 from based on context it is used
func (o *GetAPIV2MessagesOKBodyMessagesItems0From) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAPIV2MessagesOKBodyMessagesItems0From) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAPIV2MessagesOKBodyMessagesItems0From) UnmarshalBinary(b []byte) error {
	var res GetAPIV2MessagesOKBodyMessagesItems0From
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetAPIV2MessagesOKBodyMessagesItems0ToItems0 Path
swagger:model GetAPIV2MessagesOKBodyMessagesItems0ToItems0
*/
type GetAPIV2MessagesOKBodyMessagesItems0ToItems0 struct {

	// domain
	Domain string `json:"domain,omitempty"`

	// mailbox
	Mailbox string `json:"mailbox,omitempty"`

	// params
	Params string `json:"params,omitempty"`

	// relays
	Relays []string `json:"relays"`
}

// Validate validates this get API v2 messages o k body messages items0 to items0
func (o *GetAPIV2MessagesOKBodyMessagesItems0ToItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get API v2 messages o k body messages items0 to items0 based on context it is used
func (o *GetAPIV2MessagesOKBodyMessagesItems0ToItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetAPIV2MessagesOKBodyMessagesItems0ToItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetAPIV2MessagesOKBodyMessagesItems0ToItems0) UnmarshalBinary(b []byte) error {
	var res GetAPIV2MessagesOKBodyMessagesItems0ToItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
