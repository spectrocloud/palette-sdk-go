// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// V1VMVirtualMachineInstanceSpec VirtualMachineInstanceSpec is a description of a VirtualMachineInstance.
//
// swagger:model v1VmVirtualMachineInstanceSpec
type V1VMVirtualMachineInstanceSpec struct {

	// Specifies a set of public keys to inject into the vm guest
	AccessCredentials []*V1VMAccessCredential `json:"accessCredentials"`

	// affinity
	Affinity *V1VMAffinity `json:"affinity,omitempty"`

	// dns config
	DNSConfig *V1VMPodDNSConfig `json:"dnsConfig,omitempty"`

	// Set DNS policy for the pod. Defaults to "ClusterFirst". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.
	DNSPolicy string `json:"dnsPolicy,omitempty"`

	// domain
	// Required: true
	Domain *V1VMDomainSpec `json:"domain"`

	// EvictionStrategy can be set to "LiveMigrate" if the VirtualMachineInstance should be migrated instead of shut-off in case of a node drain.
	EvictionStrategy string `json:"evictionStrategy,omitempty"`

	// Specifies the hostname of the vmi If not specified, the hostname will be set to the name of the vmi, if dhcp or cloud-init is configured properly.
	Hostname string `json:"hostname,omitempty"`

	// liveness probe
	LivenessProbe *V1VMProbe `json:"livenessProbe,omitempty"`

	// List of networks that can be attached to a vm's virtual interface.
	Networks []*V1VMNetwork `json:"networks"`

	// NodeSelector is a selector which must be true for the vmi to fit on a node. Selector which must match a node's labels for the vmi to be scheduled on that node. More info: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/
	NodeSelector map[string]string `json:"nodeSelector,omitempty"`

	// If specified, indicates the pod's priority. If not specified, the pod priority will be default or zero if there is no default.
	PriorityClassName string `json:"priorityClassName,omitempty"`

	// readiness probe
	ReadinessProbe *V1VMProbe `json:"readinessProbe,omitempty"`

	// If specified, the VMI will be dispatched by specified scheduler. If not specified, the VMI will be dispatched by default scheduler.
	SchedulerName string `json:"schedulerName,omitempty"`

	// StartStrategy can be set to "Paused" if Virtual Machine should be started in paused state.
	StartStrategy string `json:"startStrategy,omitempty"`

	// If specified, the fully qualified vmi hostname will be "<hostname>.<subdomain>.<pod namespace>.svc.<cluster domain>". If not specified, the vmi will not have a domainname at all. The DNS entry will resolve to the vmi, no matter if the vmi itself can pick up a hostname.
	Subdomain string `json:"subdomain,omitempty"`

	// Grace period observed after signalling a VirtualMachineInstance to stop after which the VirtualMachineInstance is force terminated.
	TerminationGracePeriodSeconds int64 `json:"terminationGracePeriodSeconds,omitempty"`

	// If toleration is specified, obey all the toleration rules.
	Tolerations []*V1VMToleration `json:"tolerations"`

	// TopologySpreadConstraints describes how a group of VMIs will be spread across a given topology domains. K8s scheduler will schedule VMI pods in a way which abides by the constraints.
	TopologySpreadConstraints []*V1VMTopologySpreadConstraint `json:"topologySpreadConstraints"`

	// List of volumes that can be mounted by disks belonging to the vmi.
	Volumes []*V1VMVolume `json:"volumes"`
}

// Validate validates this v1 Vm virtual machine instance spec
func (m *V1VMVirtualMachineInstanceSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAffinity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLivenessProbe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadinessProbe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTolerations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopologySpreadConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) validateAccessCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.AccessCredentials) { // not required
		return nil
	}

	for i := 0; i < len(m.AccessCredentials); i++ {
		if swag.IsZero(m.AccessCredentials[i]) { // not required
			continue
		}

		if m.AccessCredentials[i] != nil {
			if err := m.AccessCredentials[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accessCredentials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accessCredentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) validateAffinity(formats strfmt.Registry) error {
	if swag.IsZero(m.Affinity) { // not required
		return nil
	}

	if m.Affinity != nil {
		if err := m.Affinity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("affinity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("affinity")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) validateDNSConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.DNSConfig) { // not required
		return nil
	}

	if m.DNSConfig != nil {
		if err := m.DNSConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dnsConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dnsConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	if m.Domain != nil {
		if err := m.Domain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) validateLivenessProbe(formats strfmt.Registry) error {
	if swag.IsZero(m.LivenessProbe) { // not required
		return nil
	}

	if m.LivenessProbe != nil {
		if err := m.LivenessProbe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("livenessProbe")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("livenessProbe")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) validateNetworks(formats strfmt.Registry) error {
	if swag.IsZero(m.Networks) { // not required
		return nil
	}

	for i := 0; i < len(m.Networks); i++ {
		if swag.IsZero(m.Networks[i]) { // not required
			continue
		}

		if m.Networks[i] != nil {
			if err := m.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) validateReadinessProbe(formats strfmt.Registry) error {
	if swag.IsZero(m.ReadinessProbe) { // not required
		return nil
	}

	if m.ReadinessProbe != nil {
		if err := m.ReadinessProbe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("readinessProbe")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("readinessProbe")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) validateTolerations(formats strfmt.Registry) error {
	if swag.IsZero(m.Tolerations) { // not required
		return nil
	}

	for i := 0; i < len(m.Tolerations); i++ {
		if swag.IsZero(m.Tolerations[i]) { // not required
			continue
		}

		if m.Tolerations[i] != nil {
			if err := m.Tolerations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tolerations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tolerations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) validateTopologySpreadConstraints(formats strfmt.Registry) error {
	if swag.IsZero(m.TopologySpreadConstraints) { // not required
		return nil
	}

	for i := 0; i < len(m.TopologySpreadConstraints); i++ {
		if swag.IsZero(m.TopologySpreadConstraints[i]) { // not required
			continue
		}

		if m.TopologySpreadConstraints[i] != nil {
			if err := m.TopologySpreadConstraints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topologySpreadConstraints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("topologySpreadConstraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) validateVolumes(formats strfmt.Registry) error {
	if swag.IsZero(m.Volumes) { // not required
		return nil
	}

	for i := 0; i < len(m.Volumes); i++ {
		if swag.IsZero(m.Volumes[i]) { // not required
			continue
		}

		if m.Volumes[i] != nil {
			if err := m.Volumes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this v1 Vm virtual machine instance spec based on the context it is used
func (m *V1VMVirtualMachineInstanceSpec) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAccessCredentials(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAffinity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNSConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDomain(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLivenessProbe(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReadinessProbe(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTolerations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTopologySpreadConstraints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) contextValidateAccessCredentials(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AccessCredentials); i++ {

		if m.AccessCredentials[i] != nil {

			if swag.IsZero(m.AccessCredentials[i]) { // not required
				return nil
			}

			if err := m.AccessCredentials[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("accessCredentials" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("accessCredentials" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) contextValidateAffinity(ctx context.Context, formats strfmt.Registry) error {

	if m.Affinity != nil {

		if swag.IsZero(m.Affinity) { // not required
			return nil
		}

		if err := m.Affinity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("affinity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("affinity")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) contextValidateDNSConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.DNSConfig != nil {

		if swag.IsZero(m.DNSConfig) { // not required
			return nil
		}

		if err := m.DNSConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dnsConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dnsConfig")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) contextValidateDomain(ctx context.Context, formats strfmt.Registry) error {

	if m.Domain != nil {

		if err := m.Domain.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) contextValidateLivenessProbe(ctx context.Context, formats strfmt.Registry) error {

	if m.LivenessProbe != nil {

		if swag.IsZero(m.LivenessProbe) { // not required
			return nil
		}

		if err := m.LivenessProbe.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("livenessProbe")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("livenessProbe")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Networks); i++ {

		if m.Networks[i] != nil {

			if swag.IsZero(m.Networks[i]) { // not required
				return nil
			}

			if err := m.Networks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) contextValidateReadinessProbe(ctx context.Context, formats strfmt.Registry) error {

	if m.ReadinessProbe != nil {

		if swag.IsZero(m.ReadinessProbe) { // not required
			return nil
		}

		if err := m.ReadinessProbe.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("readinessProbe")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("readinessProbe")
			}
			return err
		}
	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) contextValidateTolerations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tolerations); i++ {

		if m.Tolerations[i] != nil {

			if swag.IsZero(m.Tolerations[i]) { // not required
				return nil
			}

			if err := m.Tolerations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tolerations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tolerations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) contextValidateTopologySpreadConstraints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TopologySpreadConstraints); i++ {

		if m.TopologySpreadConstraints[i] != nil {

			if swag.IsZero(m.TopologySpreadConstraints[i]) { // not required
				return nil
			}

			if err := m.TopologySpreadConstraints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topologySpreadConstraints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("topologySpreadConstraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *V1VMVirtualMachineInstanceSpec) contextValidateVolumes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Volumes); i++ {

		if m.Volumes[i] != nil {

			if swag.IsZero(m.Volumes[i]) { // not required
				return nil
			}

			if err := m.Volumes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("volumes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1VMVirtualMachineInstanceSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1VMVirtualMachineInstanceSpec) UnmarshalBinary(b []byte) error {
	var res V1VMVirtualMachineInstanceSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
