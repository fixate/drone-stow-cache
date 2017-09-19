package monitor

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/date"
	"github.com/Azure/go-autorest/autorest/to"
	"net/http"
)

// ComparisonOperationType enumerates the values for comparison operation type.
type ComparisonOperationType string

const (
	// Equals specifies the equals state for comparison operation type.
	Equals ComparisonOperationType = "Equals"
	// GreaterThan specifies the greater than state for comparison operation
	// type.
	GreaterThan ComparisonOperationType = "GreaterThan"
	// GreaterThanOrEqual specifies the greater than or equal state for
	// comparison operation type.
	GreaterThanOrEqual ComparisonOperationType = "GreaterThanOrEqual"
	// LessThan specifies the less than state for comparison operation type.
	LessThan ComparisonOperationType = "LessThan"
	// LessThanOrEqual specifies the less than or equal state for comparison
	// operation type.
	LessThanOrEqual ComparisonOperationType = "LessThanOrEqual"
	// NotEquals specifies the not equals state for comparison operation type.
	NotEquals ComparisonOperationType = "NotEquals"
)

// ConditionOperator enumerates the values for condition operator.
type ConditionOperator string

const (
	// ConditionOperatorGreaterThan specifies the condition operator greater
	// than state for condition operator.
	ConditionOperatorGreaterThan ConditionOperator = "GreaterThan"
	// ConditionOperatorGreaterThanOrEqual specifies the condition operator
	// greater than or equal state for condition operator.
	ConditionOperatorGreaterThanOrEqual ConditionOperator = "GreaterThanOrEqual"
	// ConditionOperatorLessThan specifies the condition operator less than
	// state for condition operator.
	ConditionOperatorLessThan ConditionOperator = "LessThan"
	// ConditionOperatorLessThanOrEqual specifies the condition operator less
	// than or equal state for condition operator.
	ConditionOperatorLessThanOrEqual ConditionOperator = "LessThanOrEqual"
)

// MetricStatisticType enumerates the values for metric statistic type.
type MetricStatisticType string

const (
	// Average specifies the average state for metric statistic type.
	Average MetricStatisticType = "Average"
	// Max specifies the max state for metric statistic type.
	Max MetricStatisticType = "Max"
	// Min specifies the min state for metric statistic type.
	Min MetricStatisticType = "Min"
	// Sum specifies the sum state for metric statistic type.
	Sum MetricStatisticType = "Sum"
)

// RecurrenceFrequency enumerates the values for recurrence frequency.
type RecurrenceFrequency string

const (
	// Day specifies the day state for recurrence frequency.
	Day RecurrenceFrequency = "Day"
	// Hour specifies the hour state for recurrence frequency.
	Hour RecurrenceFrequency = "Hour"
	// Minute specifies the minute state for recurrence frequency.
	Minute RecurrenceFrequency = "Minute"
	// Month specifies the month state for recurrence frequency.
	Month RecurrenceFrequency = "Month"
	// None specifies the none state for recurrence frequency.
	None RecurrenceFrequency = "None"
	// Second specifies the second state for recurrence frequency.
	Second RecurrenceFrequency = "Second"
	// Week specifies the week state for recurrence frequency.
	Week RecurrenceFrequency = "Week"
	// Year specifies the year state for recurrence frequency.
	Year RecurrenceFrequency = "Year"
)

// ScaleDirection enumerates the values for scale direction.
type ScaleDirection string

const (
	// ScaleDirectionDecrease specifies the scale direction decrease state for
	// scale direction.
	ScaleDirectionDecrease ScaleDirection = "Decrease"
	// ScaleDirectionIncrease specifies the scale direction increase state for
	// scale direction.
	ScaleDirectionIncrease ScaleDirection = "Increase"
	// ScaleDirectionNone specifies the scale direction none state for scale
	// direction.
	ScaleDirectionNone ScaleDirection = "None"
)

// ScaleType enumerates the values for scale type.
type ScaleType string

const (
	// ChangeCount specifies the change count state for scale type.
	ChangeCount ScaleType = "ChangeCount"
	// ExactCount specifies the exact count state for scale type.
	ExactCount ScaleType = "ExactCount"
	// PercentChangeCount specifies the percent change count state for scale
	// type.
	PercentChangeCount ScaleType = "PercentChangeCount"
)

// TimeAggregationOperator enumerates the values for time aggregation operator.
type TimeAggregationOperator string

const (
	// TimeAggregationOperatorAverage specifies the time aggregation operator
	// average state for time aggregation operator.
	TimeAggregationOperatorAverage TimeAggregationOperator = "Average"
	// TimeAggregationOperatorLast specifies the time aggregation operator last
	// state for time aggregation operator.
	TimeAggregationOperatorLast TimeAggregationOperator = "Last"
	// TimeAggregationOperatorMaximum specifies the time aggregation operator
	// maximum state for time aggregation operator.
	TimeAggregationOperatorMaximum TimeAggregationOperator = "Maximum"
	// TimeAggregationOperatorMinimum specifies the time aggregation operator
	// minimum state for time aggregation operator.
	TimeAggregationOperatorMinimum TimeAggregationOperator = "Minimum"
	// TimeAggregationOperatorTotal specifies the time aggregation operator
	// total state for time aggregation operator.
	TimeAggregationOperatorTotal TimeAggregationOperator = "Total"
)

// TimeAggregationType enumerates the values for time aggregation type.
type TimeAggregationType string

const (
	// TimeAggregationTypeAverage specifies the time aggregation type average
	// state for time aggregation type.
	TimeAggregationTypeAverage TimeAggregationType = "Average"
	// TimeAggregationTypeCount specifies the time aggregation type count state
	// for time aggregation type.
	TimeAggregationTypeCount TimeAggregationType = "Count"
	// TimeAggregationTypeMaximum specifies the time aggregation type maximum
	// state for time aggregation type.
	TimeAggregationTypeMaximum TimeAggregationType = "Maximum"
	// TimeAggregationTypeMinimum specifies the time aggregation type minimum
	// state for time aggregation type.
	TimeAggregationTypeMinimum TimeAggregationType = "Minimum"
	// TimeAggregationTypeTotal specifies the time aggregation type total state
	// for time aggregation type.
	TimeAggregationTypeTotal TimeAggregationType = "Total"
)

// ActivityLogAlert is an Azure activity log alert.
type ActivityLogAlert struct {
	Scopes      *[]string                       `json:"scopes,omitempty"`
	Enabled     *bool                           `json:"enabled,omitempty"`
	Condition   *ActivityLogAlertAllOfCondition `json:"condition,omitempty"`
	Actions     *ActivityLogAlertActionList     `json:"actions,omitempty"`
	Description *string                         `json:"description,omitempty"`
}

// ActivityLogAlertActionGroup is a pointer to an Azure Action Group.
type ActivityLogAlertActionGroup struct {
	ActionGroupID     *string             `json:"actionGroupId,omitempty"`
	WebhookProperties *map[string]*string `json:"webhookProperties,omitempty"`
}

// ActivityLogAlertActionList is a list of activity log alert actions.
type ActivityLogAlertActionList struct {
	ActionGroups *[]ActivityLogAlertActionGroup `json:"actionGroups,omitempty"`
}

// ActivityLogAlertAllOfCondition is an Activity Log alert condition that is
// met when all its member conditions are met.
type ActivityLogAlertAllOfCondition struct {
	AllOf *[]ActivityLogAlertLeafCondition `json:"allOf,omitempty"`
}

// ActivityLogAlertLeafCondition is an Activity Log alert condition that is met
// by comparing an activity log field and value.
type ActivityLogAlertLeafCondition struct {
	Field  *string `json:"field,omitempty"`
	Equals *string `json:"equals,omitempty"`
}

// ActivityLogAlertList is a list of activity log alerts.
type ActivityLogAlertList struct {
	autorest.Response `json:"-"`
	Value             *[]ActivityLogAlertResource `json:"value,omitempty"`
}

// ActivityLogAlertPatch is an Azure activity log alert for patch operations.
type ActivityLogAlertPatch struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// ActivityLogAlertResource is an activity log alert resource.
type ActivityLogAlertResource struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	*ActivityLogAlert `json:"properties,omitempty"`
}

// ActivityLogAlertResourcePatch is an activity log alert resource for patch
// operations.
type ActivityLogAlertResourcePatch struct {
	ID                     *string             `json:"id,omitempty"`
	Name                   *string             `json:"name,omitempty"`
	Type                   *string             `json:"type,omitempty"`
	Location               *string             `json:"location,omitempty"`
	Tags                   *map[string]*string `json:"tags,omitempty"`
	*ActivityLogAlertPatch `json:"properties,omitempty"`
}

// AlertRule is an alert rule.
type AlertRule struct {
	Name            *string        `json:"name,omitempty"`
	Description     *string        `json:"description,omitempty"`
	IsEnabled       *bool          `json:"isEnabled,omitempty"`
	Condition       *RuleCondition `json:"condition,omitempty"`
	Actions         *[]RuleAction  `json:"actions,omitempty"`
	LastUpdatedTime *date.Time     `json:"lastUpdatedTime,omitempty"`
}

// AlertRuleResource is the alert rule resource.
type AlertRuleResource struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	*AlertRule        `json:"properties,omitempty"`
}

// AlertRuleResourceCollection is represents a collection of alert rule
// resources.
type AlertRuleResourceCollection struct {
	autorest.Response `json:"-"`
	Value             *[]AlertRuleResource `json:"value,omitempty"`
}

// AutoscaleNotification is autoscale notification.
type AutoscaleNotification struct {
	Operation *string                `json:"operation,omitempty"`
	Email     *EmailNotification     `json:"email,omitempty"`
	Webhooks  *[]WebhookNotification `json:"webhooks,omitempty"`
}

// AutoscaleProfile is autoscale profile.
type AutoscaleProfile struct {
	Name       *string        `json:"name,omitempty"`
	Capacity   *ScaleCapacity `json:"capacity,omitempty"`
	Rules      *[]ScaleRule   `json:"rules,omitempty"`
	FixedDate  *TimeWindow    `json:"fixedDate,omitempty"`
	Recurrence *Recurrence    `json:"recurrence,omitempty"`
}

// AutoscaleSetting is a setting that contains all of the configuration for the
// automatic scaling of a resource.
type AutoscaleSetting struct {
	Profiles          *[]AutoscaleProfile      `json:"profiles,omitempty"`
	Notifications     *[]AutoscaleNotification `json:"notifications,omitempty"`
	Enabled           *bool                    `json:"enabled,omitempty"`
	Name              *string                  `json:"name,omitempty"`
	TargetResourceURI *string                  `json:"targetResourceUri,omitempty"`
}

// AutoscaleSettingResource is the autoscale setting resource.
type AutoscaleSettingResource struct {
	autorest.Response `json:"-"`
	ID                *string             `json:"id,omitempty"`
	Name              *string             `json:"name,omitempty"`
	Type              *string             `json:"type,omitempty"`
	Location          *string             `json:"location,omitempty"`
	Tags              *map[string]*string `json:"tags,omitempty"`
	*AutoscaleSetting `json:"properties,omitempty"`
}

// AutoscaleSettingResourceCollection is represents a collection of autoscale
// setting resources.
type AutoscaleSettingResourceCollection struct {
	autorest.Response `json:"-"`
	Value             *[]AutoscaleSettingResource `json:"value,omitempty"`
	NextLink          *string                     `json:"nextLink,omitempty"`
}

// AutoscaleSettingResourceCollectionPreparer prepares a request to retrieve the next set of results. It returns
// nil if no more results exist.
func (client AutoscaleSettingResourceCollection) AutoscaleSettingResourceCollectionPreparer() (*http.Request, error) {
	if client.NextLink == nil || len(to.String(client.NextLink)) <= 0 {
		return nil, nil
	}
	return autorest.Prepare(&http.Request{},
		autorest.AsJSON(),
		autorest.AsGet(),
		autorest.WithBaseURL(to.String(client.NextLink)))
}

// EmailNotification is email notification of an autoscale event.
type EmailNotification struct {
	SendToSubscriptionAdministrator    *bool     `json:"sendToSubscriptionAdministrator,omitempty"`
	SendToSubscriptionCoAdministrators *bool     `json:"sendToSubscriptionCoAdministrators,omitempty"`
	CustomEmails                       *[]string `json:"customEmails,omitempty"`
}

// ErrorResponse is describes the format of Error response.
type ErrorResponse struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

// Incident is an alert incident indicates the activation status of an alert
// rule.
type Incident struct {
	autorest.Response `json:"-"`
	Name              *string    `json:"name,omitempty"`
	RuleName          *string    `json:"ruleName,omitempty"`
	IsActive          *bool      `json:"isActive,omitempty"`
	ActivatedTime     *date.Time `json:"activatedTime,omitempty"`
	ResolvedTime      *date.Time `json:"resolvedTime,omitempty"`
}

// IncidentListResult is the List incidents operation response.
type IncidentListResult struct {
	autorest.Response `json:"-"`
	Value             *[]Incident `json:"value,omitempty"`
}

// LocationThresholdRuleCondition is a rule condition based on a certain number
// of locations failing.
type LocationThresholdRuleCondition struct {
	DataSource          *RuleDataSource `json:"dataSource,omitempty"`
	WindowSize          *string         `json:"windowSize,omitempty"`
	FailedLocationCount *int32          `json:"failedLocationCount,omitempty"`
}

// LogProfileCollection is represents a collection of log profiles.
type LogProfileCollection struct {
	autorest.Response `json:"-"`
	Value             *[]LogProfileResource `json:"value,omitempty"`
}

// LogProfileProperties is the log profile properties.
type LogProfileProperties struct {
	StorageAccountID *string          `json:"storageAccountId,omitempty"`
	ServiceBusRuleID *string          `json:"serviceBusRuleId,omitempty"`
	Locations        *[]string        `json:"locations,omitempty"`
	Categories       *[]string        `json:"categories,omitempty"`
	RetentionPolicy  *RetentionPolicy `json:"retentionPolicy,omitempty"`
}

// LogProfileResource is the log profile resource.
type LogProfileResource struct {
	autorest.Response     `json:"-"`
	ID                    *string             `json:"id,omitempty"`
	Name                  *string             `json:"name,omitempty"`
	Type                  *string             `json:"type,omitempty"`
	Location              *string             `json:"location,omitempty"`
	Tags                  *map[string]*string `json:"tags,omitempty"`
	*LogProfileProperties `json:"properties,omitempty"`
}

// LogSettings is part of MultiTenantDiagnosticSettings. Specifies the settings
// for a particular log.
type LogSettings struct {
	Category        *string          `json:"category,omitempty"`
	Enabled         *bool            `json:"enabled,omitempty"`
	RetentionPolicy *RetentionPolicy `json:"retentionPolicy,omitempty"`
}

// ManagementEventAggregationCondition is how the data that is collected should
// be combined over time.
type ManagementEventAggregationCondition struct {
	Operator   ConditionOperator `json:"operator,omitempty"`
	Threshold  *float64          `json:"threshold,omitempty"`
	WindowSize *string           `json:"windowSize,omitempty"`
}

// ManagementEventRuleCondition is a management event rule condition.
type ManagementEventRuleCondition struct {
	DataSource  *RuleDataSource                      `json:"dataSource,omitempty"`
	Aggregation *ManagementEventAggregationCondition `json:"aggregation,omitempty"`
}

// MetricSettings is part of MultiTenantDiagnosticSettings. Specifies the
// settings for a particular metric.
type MetricSettings struct {
	TimeGrain       *string          `json:"timeGrain,omitempty"`
	Enabled         *bool            `json:"enabled,omitempty"`
	RetentionPolicy *RetentionPolicy `json:"retentionPolicy,omitempty"`
}

// MetricTrigger is the trigger that results in a scaling action.
type MetricTrigger struct {
	MetricName        *string                 `json:"metricName,omitempty"`
	MetricResourceURI *string                 `json:"metricResourceUri,omitempty"`
	TimeGrain         *string                 `json:"timeGrain,omitempty"`
	Statistic         MetricStatisticType     `json:"statistic,omitempty"`
	TimeWindow        *string                 `json:"timeWindow,omitempty"`
	TimeAggregation   TimeAggregationType     `json:"timeAggregation,omitempty"`
	Operator          ComparisonOperationType `json:"operator,omitempty"`
	Threshold         *float64                `json:"threshold,omitempty"`
}

// Recurrence is the repeating times at which this profile begins. This element
// is not used if the FixedDate element is used.
type Recurrence struct {
	Frequency RecurrenceFrequency `json:"frequency,omitempty"`
	Schedule  *RecurrentSchedule  `json:"schedule,omitempty"`
}

// RecurrentSchedule is the scheduling constraints for when the profile begins.
type RecurrentSchedule struct {
	TimeZone *string   `json:"timeZone,omitempty"`
	Days     *[]string `json:"days,omitempty"`
	Hours    *[]int32  `json:"hours,omitempty"`
	Minutes  *[]int32  `json:"minutes,omitempty"`
}

// Resource is an azure resource object
type Resource struct {
	ID       *string             `json:"id,omitempty"`
	Name     *string             `json:"name,omitempty"`
	Type     *string             `json:"type,omitempty"`
	Location *string             `json:"location,omitempty"`
	Tags     *map[string]*string `json:"tags,omitempty"`
}

// RetentionPolicy is specifies the retention policy for the log.
type RetentionPolicy struct {
	Enabled *bool  `json:"enabled,omitempty"`
	Days    *int32 `json:"days,omitempty"`
}

// RuleAction is the action that is performed when the alert rule becomes
// active, and when an alert condition is resolved.
type RuleAction struct {
}

// RuleCondition is the condition that results in the alert rule being
// activated.
type RuleCondition struct {
	DataSource *RuleDataSource `json:"dataSource,omitempty"`
}

// RuleDataSource is the resource from which the rule collects its data.
type RuleDataSource struct {
	ResourceURI *string `json:"resourceUri,omitempty"`
}

// RuleEmailAction is specifies the action to send email when the rule
// condition is evaluated. The discriminator is always RuleEmailAction in this
// case.
type RuleEmailAction struct {
	SendToServiceOwners *bool     `json:"sendToServiceOwners,omitempty"`
	CustomEmails        *[]string `json:"customEmails,omitempty"`
}

// RuleManagementEventClaimsDataSource is the claims for a rule management
// event data source.
type RuleManagementEventClaimsDataSource struct {
	EmailAddress *string `json:"emailAddress,omitempty"`
}

// RuleManagementEventDataSource is a rule management event data source. The
// discriminator fields is always RuleManagementEventDataSource in this case.
type RuleManagementEventDataSource struct {
	ResourceURI          *string                              `json:"resourceUri,omitempty"`
	EventName            *string                              `json:"eventName,omitempty"`
	EventSource          *string                              `json:"eventSource,omitempty"`
	Level                *string                              `json:"level,omitempty"`
	OperationName        *string                              `json:"operationName,omitempty"`
	ResourceGroupName    *string                              `json:"resourceGroupName,omitempty"`
	ResourceProviderName *string                              `json:"resourceProviderName,omitempty"`
	Status               *string                              `json:"status,omitempty"`
	SubStatus            *string                              `json:"subStatus,omitempty"`
	Claims               *RuleManagementEventClaimsDataSource `json:"claims,omitempty"`
}

// RuleMetricDataSource is a rule metric data source. The discriminator value
// is always RuleMetricDataSource in this case.
type RuleMetricDataSource struct {
	ResourceURI *string `json:"resourceUri,omitempty"`
	MetricName  *string `json:"metricName,omitempty"`
}

// RuleWebhookAction is specifies the action to post to service when the rule
// condition is evaluated. The discriminator is always RuleWebhookAction in
// this case.
type RuleWebhookAction struct {
	ServiceURI *string             `json:"serviceUri,omitempty"`
	Properties *map[string]*string `json:"properties,omitempty"`
}

// ScaleAction is the parameters for the scaling action.
type ScaleAction struct {
	Direction ScaleDirection `json:"direction,omitempty"`
	Type      ScaleType      `json:"type,omitempty"`
	Value     *string        `json:"value,omitempty"`
	Cooldown  *string        `json:"cooldown,omitempty"`
}

// ScaleCapacity is the number of instances that can be used during this
// profile.
type ScaleCapacity struct {
	Minimum *string `json:"minimum,omitempty"`
	Maximum *string `json:"maximum,omitempty"`
	Default *string `json:"default,omitempty"`
}

// ScaleRule is a rule that provide the triggers and parameters for the scaling
// action.
type ScaleRule struct {
	MetricTrigger *MetricTrigger `json:"metricTrigger,omitempty"`
	ScaleAction   *ScaleAction   `json:"scaleAction,omitempty"`
}

// ServiceDiagnosticSettings is the diagnostic settings for service.
type ServiceDiagnosticSettings struct {
	StorageAccountID            *string           `json:"storageAccountId,omitempty"`
	ServiceBusRuleID            *string           `json:"serviceBusRuleId,omitempty"`
	EventHubAuthorizationRuleID *string           `json:"eventHubAuthorizationRuleId,omitempty"`
	Metrics                     *[]MetricSettings `json:"metrics,omitempty"`
	Logs                        *[]LogSettings    `json:"logs,omitempty"`
	WorkspaceID                 *string           `json:"workspaceId,omitempty"`
}

// ServiceDiagnosticSettingsResource is description of a service diagnostic
// setting
type ServiceDiagnosticSettingsResource struct {
	autorest.Response          `json:"-"`
	ID                         *string             `json:"id,omitempty"`
	Name                       *string             `json:"name,omitempty"`
	Type                       *string             `json:"type,omitempty"`
	Location                   *string             `json:"location,omitempty"`
	Tags                       *map[string]*string `json:"tags,omitempty"`
	*ServiceDiagnosticSettings `json:"properties,omitempty"`
}

// ThresholdRuleCondition is a rule condition based on a metric crossing a
// threshold.
type ThresholdRuleCondition struct {
	DataSource      *RuleDataSource         `json:"dataSource,omitempty"`
	Operator        ConditionOperator       `json:"operator,omitempty"`
	Threshold       *float64                `json:"threshold,omitempty"`
	WindowSize      *string                 `json:"windowSize,omitempty"`
	TimeAggregation TimeAggregationOperator `json:"timeAggregation,omitempty"`
}

// TimeWindow is a specific date-time for the profile.
type TimeWindow struct {
	TimeZone *string    `json:"timeZone,omitempty"`
	Start    *date.Time `json:"start,omitempty"`
	End      *date.Time `json:"end,omitempty"`
}

// WebhookNotification is webhook notification of an autoscale event.
type WebhookNotification struct {
	ServiceURI *string             `json:"serviceUri,omitempty"`
	Properties *map[string]*string `json:"properties,omitempty"`
}
