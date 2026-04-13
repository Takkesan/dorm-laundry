export type MachineStatus = "available" | "running" | "awaiting_pickup" | "offline";

export type Machine = {
  id: string;
  name: string;
  location: string;
  status: MachineStatus;
  cycleMinutes: number;
  currentSessionId?: string;
  remainingMinutes?: number;
  updatedAt: string;
};
