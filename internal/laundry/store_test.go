package laundry

import (
	"testing"
	"time"
)

func TestStartSessionChangesMachineAndSession(t *testing.T) {
	store := NewStore()

	machine, err := store.StartSession("washer-a")
	if err != nil {
		t.Fatalf("StartSession returned error: %v", err)
	}

	if machine.Status != StatusRunning {
		t.Fatalf("expected running status, got %s", machine.Status)
	}

	session := store.CurrentSession()
	if session == nil {
		t.Fatal("expected current session")
	}

	if session.MachineID != "washer-a" {
		t.Fatalf("expected machine washer-a, got %s", session.MachineID)
	}
}

func TestClaimSessionClearsCurrentSession(t *testing.T) {
	store := NewStore()

	session := store.CurrentSession()
	if session == nil {
		t.Fatal("expected seeded session")
	}

	if err := store.ClaimSession(session.ID); err != nil {
		t.Fatalf("ClaimSession returned error: %v", err)
	}

	if store.CurrentSession() != nil {
		t.Fatal("expected current session to be cleared")
	}
}

func TestCountByStatus(t *testing.T) {
	store := NewStore()
	machines := store.ListMachines()

	if got := CountByStatus(machines, StatusAvailable); got != 1 {
		t.Fatalf("expected 1 available machine, got %d", got)
	}

	if got := CountByStatus(machines, StatusRunning); got != 1 {
		t.Fatalf("expected 1 running machine, got %d", got)
	}

	if got := CountByStatus(machines, StatusAwaitingPickup); got != 1 {
		t.Fatalf("expected 1 awaiting pickup machine, got %d", got)
	}

	if got := CountByStatus(machines, StatusOffline); got != 1 {
		t.Fatalf("expected 1 offline machine, got %d", got)
	}
}

func TestSessionProgressPercent(t *testing.T) {
	now := time.Date(2026, 4, 14, 10, 0, 0, 0, time.FixedZone("JST", 9*60*60))
	session := Session{
		StartedAt: now.Add(-10 * time.Minute).Format(time.RFC3339),
		EndsAt:    now.Add(10 * time.Minute).Format(time.RFC3339),
	}

	if got := SessionProgressPercent(session, now); got != 50 {
		t.Fatalf("expected 50 percent progress, got %d", got)
	}
}
