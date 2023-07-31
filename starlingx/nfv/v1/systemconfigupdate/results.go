/* SPDX-License-Identifier: Apache-2.0 */
/* Copyright(c) 2023 Wind River Systems, Inc. */

package systemconfigupdate

import (
	"encoding/json"

	"github.com/gophercloud/gophercloud"
)

// DeleteResult represents the result of an delete operation.
type DeleteResult struct {
	gophercloud.ErrResult
}

type SystemConfigStep struct {
	StartDateTime string   `json:"start-date-time,omitempty"`
	EndDateTime   string   `json:"end-date-time,omitempty"`
	Timeout       int      `json:"timeout,omitempty"`
	EntityType    string   `json:"entity-type,omitempty"`
	StepId        int      `json:"step-id,omitempty"`
	EntityUuids   []string `json:"entity-uuids,omitempty"`
	StepName      string   `json:"step-name,omitempty"`
	Result        string   `json:"result,omitempty"`
	EntityNames   []string `json:"entity-names,omitempty"`
	Reason        string   `json:"reason,omitempty"`
}

type SystemConfigStage struct {
	StartDateTime string             `json:"start-date-time,omitempty"`
	EndDateTime   string             `json:"end-date-time,omitempty"`
	StageId       int                `json:"stage-id,omitempty"`
	Reason        string             `json:"reason,omitempty"`
	CurrentStep   int                `json:"current-step,omitempty"`
	Steps         []SystemConfigStep `json:"steps,omitempty"`
	Result        string             `json:"result,omitempty"`
	Timeout       int                `json:"timeout,omitempty"`
	TotalSteps    int                `json:"total-steps,omitempty"`
	Inprogress    bool               `json:"inprogress,omitempty"`
	StageName     string             `json:"stage-name,omitempty"`
}

type SystemConfigPhase struct {
	StartDateTime        string              `json:"start-date-time,omitempty"`
	EndDateTime          string              `json:"end-date-time,omitempty"`
	PhaseName            string              `json:"phase-name,omitempty"`
	CompletionPercentage int                 `json:"completion-percentage,omitempty"`
	TotalStages          int                 `json:"total-stages,omitempty"`
	StopAtStage          int                 `json:"stop-at-stage,omitempty"`
	Result               string              `json:"sresult,omitempty"`
	Timeout              int                 `json:"timeout,omitempty"`
	Reason               string              `json:"reason,omitempty"`
	Inprogress           bool                `json:"inprogress,omitempty"`
	Stages               []SystemConfigStage `json:"stages,omitempty"`
	CurrentStage         int                 `json:"current-stage,omitempty"`
}

// SystemConfigUpdateStrategy defines the data associated to a single System Config Update Strategy instance.
type SystemConfigUpdate struct {
	// ID is the generated unique UUID for the System Config Update Strategy
	ID string `json:"uuid"`

	// StrategyName is the name of the strategy.
	StrategyName string `json:"name"`

	// ControllerApplyType is the apply type for controller hosts.
	ControllerApplyType string `json:"controller-apply-type"`

	// StorageApplyType is the apply type for storage hosts.
	StorageApplyType string `json:"storage-apply-type"`

	// WorkerApplyType is the apply type for worker hosts.
	WorkerApplyType string `json:"worker-apply-type"`

	// The maximum number of worker hosts to update in parallel; only applicable
	// if ``worker-apply-type = parallel``.
	MaxParallerWorkers int `json:"max-parallel-worker-hosts,omitempty"`

	// The default instance action.
	DefaultInstanceAction string `json:"default-instance-action"`

	// The strictness of alarm checks.
	AlarmRestrictions string `json:"alarm-restrictions,omitempty"`

	// The strictness of alarm checks.
	State string `json:"state"`

	// Build Phase
	BuildPhase SystemConfigPhase `json:"build-phase,omitempty"`

	// Apply Phase
	ApplyPhase SystemConfigPhase `json:"apply-phase,omitempty"`

	// Abort Phase
	AbortPhase SystemConfigPhase `json:"abort-phase,omitempty"`
}

// GenerateSCUpdate takes a JSON string and converts it into a SystemConfigUpdate structure.
// It returns a pointer to SystemConfigUpdate and an error if the JSON parsing fails.
func GenerateSCUpdate(jsonData string) (*SystemConfigUpdate, error) {
	var data map[string]SystemConfigUpdate
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return nil, err
	}

	scUpdate := data["strategy"]
	return &scUpdate, nil
}
