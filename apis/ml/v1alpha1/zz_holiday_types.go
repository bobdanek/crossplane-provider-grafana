/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type CustomPeriodsInitParameters struct {
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The name of the custom period.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type CustomPeriodsObservation struct {
	EndTime *string `json:"endTime,omitempty" tf:"end_time,omitempty"`

	// The name of the custom period.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	StartTime *string `json:"startTime,omitempty" tf:"start_time,omitempty"`
}

type CustomPeriodsParameters struct {

	// +kubebuilder:validation:Optional
	EndTime *string `json:"endTime" tf:"end_time,omitempty"`

	// The name of the custom period.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	StartTime *string `json:"startTime" tf:"start_time,omitempty"`
}

type HolidayInitParameters struct {

	// A list of custom periods for the holiday.
	CustomPeriods []CustomPeriodsInitParameters `json:"customPeriods,omitempty" tf:"custom_periods,omitempty"`

	// A description of the holiday.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The timezone to use for events in the iCal file pointed to by ical_url.
	IcalTimezone *string `json:"icalTimezone,omitempty" tf:"ical_timezone,omitempty"`

	// A URL to an iCal file containing all occurrences of the holiday.
	IcalURL *string `json:"icalUrl,omitempty" tf:"ical_url,omitempty"`

	// The name of the holiday.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type HolidayObservation struct {

	// A list of custom periods for the holiday.
	CustomPeriods []CustomPeriodsObservation `json:"customPeriods,omitempty" tf:"custom_periods,omitempty"`

	// A description of the holiday.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The timezone to use for events in the iCal file pointed to by ical_url.
	IcalTimezone *string `json:"icalTimezone,omitempty" tf:"ical_timezone,omitempty"`

	// A URL to an iCal file containing all occurrences of the holiday.
	IcalURL *string `json:"icalUrl,omitempty" tf:"ical_url,omitempty"`

	// The name of the holiday.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type HolidayParameters struct {

	// A list of custom periods for the holiday.
	// +kubebuilder:validation:Optional
	CustomPeriods []CustomPeriodsParameters `json:"customPeriods,omitempty" tf:"custom_periods,omitempty"`

	// A description of the holiday.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The timezone to use for events in the iCal file pointed to by ical_url.
	// +kubebuilder:validation:Optional
	IcalTimezone *string `json:"icalTimezone,omitempty" tf:"ical_timezone,omitempty"`

	// A URL to an iCal file containing all occurrences of the holiday.
	// +kubebuilder:validation:Optional
	IcalURL *string `json:"icalUrl,omitempty" tf:"ical_url,omitempty"`

	// The name of the holiday.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// HolidaySpec defines the desired state of Holiday
type HolidaySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HolidayParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider HolidayInitParameters `json:"initProvider,omitempty"`
}

// HolidayStatus defines the observed state of Holiday.
type HolidayStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HolidayObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Holiday is the Schema for the Holidays API. A holiday describes time periods where a time series is expected to behave differently to normal. To use a holiday in a job, use its id in the holidays attribute of a grafana_machine_learning_job:
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,grafana}
type Holiday struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   HolidaySpec   `json:"spec"`
	Status HolidayStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HolidayList contains a list of Holidays
type HolidayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Holiday `json:"items"`
}

// Repository type metadata.
var (
	Holiday_Kind             = "Holiday"
	Holiday_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Holiday_Kind}.String()
	Holiday_KindAPIVersion   = Holiday_Kind + "." + CRDGroupVersion.String()
	Holiday_GroupVersionKind = CRDGroupVersion.WithKind(Holiday_Kind)
)

func init() {
	SchemeBuilder.Register(&Holiday{}, &HolidayList{})
}
