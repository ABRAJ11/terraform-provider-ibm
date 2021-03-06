// Code generated by go-swagger; DO NOT EDIT.

package l_baas

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDReader is a Reader for the PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleID structure.
type PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDOK creates a PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDOK with default headers values
func NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDOK() *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDOK {
	return &PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDOK{}
}

/*PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDOK handles this case with default header values.

The rule update request was accepted.
*/
type PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDOK struct {
	Payload *models.ListenerPolicyRule
}

func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDOK) Error() string {
	return fmt.Sprintf("[PATCH /load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}/rules/{rule_id}][%d] patchLoadBalancersIdListenersListenerIdPoliciesPolicyIdRulesRuleIdOK  %+v", 200, o.Payload)
}

func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ListenerPolicyRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDBadRequest creates a PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDBadRequest with default headers values
func NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDBadRequest() *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDBadRequest {
	return &PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDBadRequest{}
}

/*PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDBadRequest handles this case with default header values.

An invalid rule template was provided.
*/
type PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}/rules/{rule_id}][%d] patchLoadBalancersIdListenersListenerIdPoliciesPolicyIdRulesRuleIdBadRequest  %+v", 400, o.Payload)
}

func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDNotFound creates a PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDNotFound with default headers values
func NewPatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDNotFound() *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDNotFound {
	return &PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDNotFound{}
}

/*PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDNotFound handles this case with default header values.

A load balancer, listener or policy with the specified identifier could not be found.
*/
type PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDNotFound) Error() string {
	return fmt.Sprintf("[PATCH /load_balancers/{id}/listeners/{listener_id}/policies/{policy_id}/rules/{rule_id}][%d] patchLoadBalancersIdListenersListenerIdPoliciesPolicyIdRulesRuleIdNotFound  %+v", 404, o.Payload)
}

func (o *PatchLoadBalancersIDListenersListenerIDPoliciesPolicyIDRulesRuleIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
