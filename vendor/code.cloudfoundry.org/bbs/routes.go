package bbs

import "github.com/tedsuo/rata"

const (
	// Ping
	PingRoute_r0 = "Ping"

	// Domains
	DomainsRoute_r0      = "Domains"
	UpsertDomainRoute_r0 = "UpsertDomain"

	// Actual LRPs
	ActualLRPsRoute_r0                          = "ActualLRPs"
	ActualLRPGroupsRoute_r0                     = "ActualLRPGroups"
	ActualLRPGroupsByProcessGuidRoute_r0        = "ActualLRPGroupsByProcessGuid"
	ActualLRPGroupByProcessGuidAndIndexRoute_r0 = "ActualLRPGroupsByProcessGuidAndIndex"

	// Actual LRP Lifecycle
	ClaimActualLRPRoute_r0  = "ClaimActualLRP"
	StartActualLRPRoute_r0  = "StartActualLRP"
	CrashActualLRPRoute_r0  = "CrashActualLRP"
	FailActualLRPRoute_r0   = "FailActualLRP"
	RemoveActualLRPRoute_r0 = "RemoveActualLRP"
	RetireActualLRPRoute_r0 = "RetireActualLRP"

	// Evacuation
	RemoveEvacuatingActualLRPRoute_r0 = "RemoveEvacuatingActualLRP"
	EvacuateClaimedActualLRPRoute_r0  = "EvacuateClaimedActualLRP"
	EvacuateCrashedActualLRPRoute_r0  = "EvacuateCrashedActualLRP"
	EvacuateStoppedActualLRPRoute_r0  = "EvacuateStoppedActualLRP"
	EvacuateRunningActualLRPRoute_r0  = "EvacuateRunningActualLRP"

	// Desired LRPs
	DesiredLRPsRoute_r2               = "DesiredLRPs"
	DesiredLRPSchedulingInfosRoute_r0 = "DesiredLRPSchedulingInfos"
	DesiredLRPByProcessGuidRoute_r2   = "DesiredLRPByProcessGuid"

	// Desire LRP Lifecycle
	DesireDesiredLRPRoute_r2 = "DesireDesiredLRP"
	UpdateDesiredLRPRoute_r0 = "UpdateDesireLRP"
	RemoveDesiredLRPRoute    = "RemoveDesiredLRP"

	// Tasks
	TasksRoute_r2         = "Tasks"
	TaskByGuidRoute_r2    = "TaskByGuid"
	DesireTaskRoute_r2    = "DesireTask"
	StartTaskRoute_r0     = "StartTask"
	CancelTaskRoute_r0    = "CancelTask"
	FailTaskRoute_r0      = "FailTask"
	RejectTaskRoute_r0    = "RejectTask"
	CompleteTaskRoute_r0  = "CompleteTask"
	ResolvingTaskRoute_r0 = "ResolvingTask"
	DeleteTaskRoute_r0    = "DeleteTask"

	// Event Streaming
	EventStreamRoute_r0     = "EventStream_r0"
	TaskEventStreamRoute_r0 = "TaskEventStream_r0"

	// Cell Presence
	CellsRoute_r0 = "Cells"
)

var Routes = rata.Routes{
	// Ping
	{Path: "/v1/ping", Method: "POST", Name: PingRoute_r0},

	// Domains
	{Path: "/v1/domains/list", Method: "POST", Name: DomainsRoute_r0},
	{Path: "/v1/domains/upsert", Method: "POST", Name: UpsertDomainRoute_r0},

	// Actual LRPs
	{Path: "/v1/actual_lrps/list", Method: "POST", Name: ActualLRPsRoute_r0},
	{Path: "/v1/actual_lrp_groups/list", Method: "POST", Name: ActualLRPGroupsRoute_r0},
	{Path: "/v1/actual_lrp_groups/list_by_process_guid", Method: "POST", Name: ActualLRPGroupsByProcessGuidRoute_r0},
	{Path: "/v1/actual_lrp_groups/get_by_process_guid_and_index", Method: "POST", Name: ActualLRPGroupByProcessGuidAndIndexRoute_r0},

	// Actual LRP Lifecycle
	{Path: "/v1/actual_lrps/claim", Method: "POST", Name: ClaimActualLRPRoute_r0},
	{Path: "/v1/actual_lrps/start", Method: "POST", Name: StartActualLRPRoute_r0},
	{Path: "/v1/actual_lrps/crash", Method: "POST", Name: CrashActualLRPRoute_r0},
	{Path: "/v1/actual_lrps/fail", Method: "POST", Name: FailActualLRPRoute_r0},
	{Path: "/v1/actual_lrps/remove", Method: "POST", Name: RemoveActualLRPRoute_r0},
	{Path: "/v1/actual_lrps/retire", Method: "POST", Name: RetireActualLRPRoute_r0},

	// Evacuation
	{Path: "/v1/actual_lrps/remove_evacuating", Method: "POST", Name: RemoveEvacuatingActualLRPRoute_r0},
	{Path: "/v1/actual_lrps/evacuate_claimed", Method: "POST", Name: EvacuateClaimedActualLRPRoute_r0},
	{Path: "/v1/actual_lrps/evacuate_crashed", Method: "POST", Name: EvacuateCrashedActualLRPRoute_r0},
	{Path: "/v1/actual_lrps/evacuate_stopped", Method: "POST", Name: EvacuateStoppedActualLRPRoute_r0},
	{Path: "/v1/actual_lrps/evacuate_running", Method: "POST", Name: EvacuateRunningActualLRPRoute_r0},

	// Desired LRPs
	{Path: "/v1/desired_lrp_scheduling_infos/list", Method: "POST", Name: DesiredLRPSchedulingInfosRoute_r0},

	{Path: "/v1/desired_lrps/list.r2", Method: "POST", Name: DesiredLRPsRoute_r2},
	{Path: "/v1/desired_lrps/get_by_process_guid.r2", Method: "POST", Name: DesiredLRPByProcessGuidRoute_r2},

	// Desire LPR Lifecycle
	{Path: "/v1/desired_lrp/desire.r2", Method: "POST", Name: DesireDesiredLRPRoute_r2},
	{Path: "/v1/desired_lrp/update", Method: "POST", Name: UpdateDesiredLRPRoute_r0},
	{Path: "/v1/desired_lrp/remove", Method: "POST", Name: RemoveDesiredLRPRoute},

	// Tasks
	{Path: "/v1/tasks/list.r2", Method: "POST", Name: TasksRoute_r2},
	{Path: "/v1/tasks/get_by_task_guid.r2", Method: "POST", Name: TaskByGuidRoute_r2},

	// Task Lifecycle
	{Path: "/v1/tasks/desire.r2", Method: "POST", Name: DesireTaskRoute_r2},
	{Path: "/v1/tasks/start", Method: "POST", Name: StartTaskRoute_r0},
	{Path: "/v1/tasks/cancel", Method: "POST", Name: CancelTaskRoute_r0},
	{Path: "/v1/tasks/fail", Method: "POST", Name: FailTaskRoute_r0},
	{Path: "/v1/tasks/reject", Method: "POST", Name: RejectTaskRoute_r0},
	{Path: "/v1/tasks/complete", Method: "POST", Name: CompleteTaskRoute_r0},
	{Path: "/v1/tasks/resolving", Method: "POST", Name: ResolvingTaskRoute_r0},
	{Path: "/v1/tasks/delete", Method: "POST", Name: DeleteTaskRoute_r0},

	// Event Streaming
	{Path: "/v1/events", Method: "GET", Name: EventStreamRoute_r0},
	{Path: "/v1/events/tasks", Method: "POST", Name: TaskEventStreamRoute_r0},

	// Cells
	{Path: "/v1/cells/list.r1", Method: "POST", Name: CellsRoute_r0},
}
