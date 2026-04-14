package laundry

import (
	"errors"
	"fmt"
	"slices"
	"sync"
	"time"
)

var (
	ErrMachineNotFound = errors.New("machine_not_found")
	ErrMachineBusy     = errors.New("machine_not_available")
	ErrSessionNotFound = errors.New("session_not_found")
)

type Store struct {
	mu                  sync.RWMutex
	machines            []Machine
	currentSession      *Session
	notificationEnabled bool
	now                 func() time.Time
}

func NewStore() *Store {
	now := time.Date(2026, 4, 13, 22, 30, 0, 0, time.FixedZone("JST", 9*60*60))

	return &Store{
		now: func() time.Time {
			return now
		},
		machines: []Machine{
			{
				ID:           "washer-a",
				Name:         "洗濯機 A",
				Location:     "1階 ランドリー室",
				Status:       StatusAvailable,
				CycleMinutes: 35,
				UpdatedAt:    now.Format(time.RFC3339),
			},
			{
				ID:               "washer-b",
				Name:             "洗濯機 B",
				Location:         "1階 ランドリー室",
				Status:           StatusRunning,
				CycleMinutes:     40,
				CurrentSessionID: "session-1",
				RemainingMinutes: 18,
				UpdatedAt:        now.Format(time.RFC3339),
			},
			{
				ID:               "washer-c",
				Name:             "洗濯機 C",
				Location:         "2階 ランドリー室",
				Status:           StatusAwaitingPickup,
				CycleMinutes:     40,
				CurrentSessionID: "session-2",
				RemainingMinutes: 0,
				UpdatedAt:        now.Format(time.RFC3339),
			},
			{
				ID:           "washer-d",
				Name:         "洗濯機 D",
				Location:     "2階 ランドリー室",
				Status:       StatusOffline,
				CycleMinutes: 40,
				UpdatedAt:    now.Format(time.RFC3339),
			},
		},
		currentSession: &Session{
			ID:                  "session-1",
			MachineID:           "washer-b",
			MachineName:         "洗濯機 B",
			Status:              SessionRunning,
			StartedAt:           now.Add(-22 * time.Minute).Format(time.RFC3339),
			EndsAt:              now.Add(18 * time.Minute).Format(time.RFC3339),
			NotificationEnabled: false,
		},
	}
}

func (s *Store) ListMachines() []Machine {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return slices.Clone(s.machines)
}

func (s *Store) Now() time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.now()
}

func (s *Store) GetMachine(machineID string) (Machine, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for _, machine := range s.machines {
		if machine.ID == machineID {
			return machine, nil
		}
	}

	return Machine{}, ErrMachineNotFound
}

func (s *Store) CurrentSession() *Session {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.currentSession == nil {
		return nil
	}

	session := *s.currentSession
	return &session
}

func (s *Store) NotificationPreference() NotificationPreference {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return NotificationPreference{
		Enabled:   s.notificationEnabled,
		Supported: true,
	}
}

func (s *Store) StartSession(machineID string) (Machine, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	index := -1
	for i, machine := range s.machines {
		if machine.ID == machineID {
			index = i
			break
		}
	}

	if index == -1 {
		return Machine{}, ErrMachineNotFound
	}

	machine := s.machines[index]
	if machine.Status != StatusAvailable {
		return machine, ErrMachineBusy
	}

	now := s.now()
	sessionID := fmt.Sprintf("session-%d", now.Unix())

	machine.Status = StatusRunning
	machine.CurrentSessionID = sessionID
	machine.RemainingMinutes = machine.CycleMinutes
	machine.UpdatedAt = now.Format(time.RFC3339)
	s.machines[index] = machine

	s.currentSession = &Session{
		ID:                  sessionID,
		MachineID:           machine.ID,
		MachineName:         machine.Name,
		Status:              SessionRunning,
		StartedAt:           now.Format(time.RFC3339),
		EndsAt:              now.Add(time.Duration(machine.CycleMinutes) * time.Minute).Format(time.RFC3339),
		NotificationEnabled: s.notificationEnabled,
	}

	return machine, nil
}

func (s *Store) ClaimSession(sessionID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.currentSession == nil || s.currentSession.ID != sessionID {
		return ErrSessionNotFound
	}

	for i, machine := range s.machines {
		if machine.ID != s.currentSession.MachineID {
			continue
		}

		machine.Status = StatusAvailable
		machine.CurrentSessionID = ""
		machine.RemainingMinutes = 0
		machine.UpdatedAt = s.now().Format(time.RFC3339)
		s.machines[i] = machine
		break
	}

	s.currentSession = nil
	return nil
}

func (s *Store) EnableNotifications() *Session {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.notificationEnabled = true
	if s.currentSession != nil {
		s.currentSession.NotificationEnabled = true
		session := *s.currentSession
		return &session
	}

	return nil
}

func StatusLabel(status MachineStatus) string {
	switch status {
	case StatusAvailable:
		return "空き"
	case StatusRunning:
		return "使用中"
	case StatusAwaitingPickup:
		return "回収待ち"
	case StatusOffline:
		return "利用不可"
	default:
		return string(status)
	}
}

func SessionStatusLabel(status SessionStatus) string {
	if status == SessionAwaitingPickup {
		return "回収待ち"
	}

	return "使用中"
}

func StatusDescription(status MachineStatus) string {
	switch status {
	case StatusAvailable:
		return "いま利用開始できます。"
	case StatusRunning:
		return "ほかの利用者が使用中です。"
	case StatusAwaitingPickup:
		return "洗濯は終了していますが、まだ回収されていません。"
	case StatusOffline:
		return "故障または停止中のため利用できません。"
	default:
		return ""
	}
}

func MachineStatusSummary(machine Machine) string {
	switch machine.Status {
	case StatusAvailable:
		return "空いています。現地に向かう前に利用開始できます。"
	case StatusRunning:
		return fmt.Sprintf("%sで終了見込みです。", RelativeMinutes(machine.RemainingMinutes))
	case StatusAwaitingPickup:
		return "洗濯は終わっています。空くまで少し待ってください。"
	case StatusOffline:
		return "この洗濯機は現在停止中です。別の洗濯機を選んでください。"
	default:
		return ""
	}
}

func ActionHint(machine Machine) string {
	switch machine.Status {
	case StatusAvailable:
		return "空きのため、そのまま利用開始できます。"
	case StatusRunning:
		return "利用開始はできません。終了予定を確認してください。"
	case StatusAwaitingPickup:
		return "回収待ちのため利用開始はできません。"
	case StatusOffline:
		return "利用不可のため受付していません。"
	default:
		return ""
	}
}

func SessionSummary(session Session, now time.Time) string {
	if session.Status == SessionAwaitingPickup {
		return "洗濯は終了しています。回収完了を登録してください。"
	}

	endTime, err := time.Parse(time.RFC3339, session.EndsAt)
	if err != nil {
		return "終了予定を確認中です。"
	}

	remaining := int(endTime.Sub(now).Minutes())
	if remaining < 0 {
		remaining = 0
	}

	return fmt.Sprintf("%sで終了予定です。", RelativeMinutes(remaining))
}

func RelativeMinutes(minutes int) string {
	if minutes <= 0 {
		return "まもなく終了"
	}

	return fmt.Sprintf("残り約%d分", minutes)
}

func CountByStatus(machines []Machine, status MachineStatus) int {
	count := 0
	for _, machine := range machines {
		if machine.Status == status {
			count++
		}
	}

	return count
}

func SessionProgressPercent(session Session, now time.Time) int {
	startTime, startErr := time.Parse(time.RFC3339, session.StartedAt)
	endTime, endErr := time.Parse(time.RFC3339, session.EndsAt)
	if startErr != nil || endErr != nil {
		return 0
	}

	total := int(endTime.Sub(startTime).Minutes())
	if total <= 0 {
		return 0
	}

	elapsed := int(now.Sub(startTime).Minutes())
	if elapsed < 0 {
		elapsed = 0
	}
	if elapsed > total {
		elapsed = total
	}

	return (elapsed * 100) / total
}

func SummarizeMachines(machines []Machine) []StatusSummary {
	statuses := []MachineStatus{
		StatusAvailable,
		StatusRunning,
		StatusAwaitingPickup,
		StatusOffline,
	}

	summary := make([]StatusSummary, 0, len(statuses))
	for _, status := range statuses {
		count := 0
		for _, machine := range machines {
			if machine.Status == status {
				count++
			}
		}

		if count == 0 {
			continue
		}

		summary = append(summary, StatusSummary{
			Status: status,
			Label:  StatusLabel(status),
			Count:  count,
		})
	}

	return summary
}
