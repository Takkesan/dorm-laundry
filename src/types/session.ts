export type Session = {
  id: string;
  machineId: string;
  machineName: string;
  status: "running" | "awaiting_pickup";
  startedAt: string;
  endsAt: string;
  notificationEnabled: boolean;
};
