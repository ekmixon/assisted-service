// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// UpdateClusterReader is a Reader for the UpdateCluster structure.
type UpdateClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateClusterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewUpdateClusterMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateClusterConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateClusterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateClusterCreated creates a UpdateClusterCreated with default headers values
func NewUpdateClusterCreated() *UpdateClusterCreated {
	return &UpdateClusterCreated{}
}

/* UpdateClusterCreated describes a response with status code 201, with default header values.

Success.
*/
type UpdateClusterCreated struct {
	Payload *models.Cluster
}

func (o *UpdateClusterCreated) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}][%d] updateClusterCreated  %+v", 201, o.Payload)
}
func (o *UpdateClusterCreated) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *UpdateClusterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterBadRequest creates a UpdateClusterBadRequest with default headers values
func NewUpdateClusterBadRequest() *UpdateClusterBadRequest {
	return &UpdateClusterBadRequest{}
}

/* UpdateClusterBadRequest describes a response with status code 400, with default header values.

Error.
*/
type UpdateClusterBadRequest struct {
	Payload *models.Error
}

func (o *UpdateClusterBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}][%d] updateClusterBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateClusterBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterUnauthorized creates a UpdateClusterUnauthorized with default headers values
func NewUpdateClusterUnauthorized() *UpdateClusterUnauthorized {
	return &UpdateClusterUnauthorized{}
}

/* UpdateClusterUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type UpdateClusterUnauthorized struct {
	Payload *models.InfraError
}

func (o *UpdateClusterUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}][%d] updateClusterUnauthorized  %+v", 401, o.Payload)
}
func (o *UpdateClusterUnauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UpdateClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterForbidden creates a UpdateClusterForbidden with default headers values
func NewUpdateClusterForbidden() *UpdateClusterForbidden {
	return &UpdateClusterForbidden{}
}

/* UpdateClusterForbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type UpdateClusterForbidden struct {
	Payload *models.InfraError
}

func (o *UpdateClusterForbidden) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}][%d] updateClusterForbidden  %+v", 403, o.Payload)
}
func (o *UpdateClusterForbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *UpdateClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterNotFound creates a UpdateClusterNotFound with default headers values
func NewUpdateClusterNotFound() *UpdateClusterNotFound {
	return &UpdateClusterNotFound{}
}

/* UpdateClusterNotFound describes a response with status code 404, with default header values.

Error.
*/
type UpdateClusterNotFound struct {
	Payload *models.Error
}

func (o *UpdateClusterNotFound) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}][%d] updateClusterNotFound  %+v", 404, o.Payload)
}
func (o *UpdateClusterNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterMethodNotAllowed creates a UpdateClusterMethodNotAllowed with default headers values
func NewUpdateClusterMethodNotAllowed() *UpdateClusterMethodNotAllowed {
	return &UpdateClusterMethodNotAllowed{}
}

/* UpdateClusterMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type UpdateClusterMethodNotAllowed struct {
	Payload *models.Error
}

func (o *UpdateClusterMethodNotAllowed) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}][%d] updateClusterMethodNotAllowed  %+v", 405, o.Payload)
}
func (o *UpdateClusterMethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateClusterMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterConflict creates a UpdateClusterConflict with default headers values
func NewUpdateClusterConflict() *UpdateClusterConflict {
	return &UpdateClusterConflict{}
}

/* UpdateClusterConflict describes a response with status code 409, with default header values.

Error.
*/
type UpdateClusterConflict struct {
	Payload *models.Error
}

func (o *UpdateClusterConflict) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}][%d] updateClusterConflict  %+v", 409, o.Payload)
}
func (o *UpdateClusterConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateClusterConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateClusterInternalServerError creates a UpdateClusterInternalServerError with default headers values
func NewUpdateClusterInternalServerError() *UpdateClusterInternalServerError {
	return &UpdateClusterInternalServerError{}
}

/* UpdateClusterInternalServerError describes a response with status code 500, with default header values.

Error.
*/
type UpdateClusterInternalServerError struct {
	Payload *models.Error
}

func (o *UpdateClusterInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /v1/clusters/{cluster_id}][%d] updateClusterInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateClusterInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *UpdateClusterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
