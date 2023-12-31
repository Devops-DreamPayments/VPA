/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// ProvisioningRequestSpecApplyConfiguration represents an declarative configuration of the ProvisioningRequestSpec type for use
// with apply.
type ProvisioningRequestSpecApplyConfiguration struct {
	PodSets              []PodSetApplyConfiguration `json:"podSets,omitempty"`
	ProvisioningClass    *string                    `json:"provisioningClass,omitempty"`
	AdditionalParameters map[string]string          `json:"additionalParameters,omitempty"`
}

// ProvisioningRequestSpecApplyConfiguration constructs an declarative configuration of the ProvisioningRequestSpec type for use with
// apply.
func ProvisioningRequestSpec() *ProvisioningRequestSpecApplyConfiguration {
	return &ProvisioningRequestSpecApplyConfiguration{}
}

// WithPodSets adds the given value to the PodSets field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the PodSets field.
func (b *ProvisioningRequestSpecApplyConfiguration) WithPodSets(values ...*PodSetApplyConfiguration) *ProvisioningRequestSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPodSets")
		}
		b.PodSets = append(b.PodSets, *values[i])
	}
	return b
}

// WithProvisioningClass sets the ProvisioningClass field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProvisioningClass field is set to the value of the last call.
func (b *ProvisioningRequestSpecApplyConfiguration) WithProvisioningClass(value string) *ProvisioningRequestSpecApplyConfiguration {
	b.ProvisioningClass = &value
	return b
}

// WithAdditionalParameters puts the entries into the AdditionalParameters field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the AdditionalParameters field,
// overwriting an existing map entries in AdditionalParameters field with the same key.
func (b *ProvisioningRequestSpecApplyConfiguration) WithAdditionalParameters(entries map[string]string) *ProvisioningRequestSpecApplyConfiguration {
	if b.AdditionalParameters == nil && len(entries) > 0 {
		b.AdditionalParameters = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.AdditionalParameters[k] = v
	}
	return b
}
