package laundry

type MachineStatus string

const (
	StatusAvailable      MachineStatus = "available"
	StatusRunning        MachineStatus = "running"
	StatusAwaitingPickup MachineStatus = "awaiting_pickup"
	StatusOffline        MachineStatus = "offline"
)

type Machine struct {
	ID               string
	Name             string
	Location         string
	Status           MachineStatus
	CycleMinutes     int
	CurrentSessionID string
	RemainingMinutes int
	UpdatedAt        string
}

type SessionStatus string

const (
	SessionRunning        SessionStatus = "running"
	SessionAwaitingPickup SessionStatus = "awaiting_pickup"
)

type Session struct {
	ID                  string
	MachineID           string
	MachineName         string
	Status              SessionStatus
	StartedAt           string
	EndsAt              string
	NotificationEnabled bool
}

type NotificationPreference struct {
	Enabled   bool
	Supported bool
}

type StatusSummary struct {
	Status MachineStatus
	Label  string
	Count  int
}
