import { Machine } from "@/types/machine";
import { Session } from "@/types/session";

const now = new Date("2026-04-13T22:30:00+09:00");

export const machinesMock: Machine[] = [
  {
    id: "washer-a",
    name: "洗濯機 A",
    location: "1階 ランドリー室",
    status: "available",
    cycleMinutes: 35,
    updatedAt: now.toISOString()
  },
  {
    id: "washer-b",
    name: "洗濯機 B",
    location: "1階 ランドリー室",
    status: "running",
    cycleMinutes: 40,
    currentSessionId: "session-1",
    remainingMinutes: 18,
    updatedAt: now.toISOString()
  },
  {
    id: "washer-c",
    name: "洗濯機 C",
    location: "2階 ランドリー室",
    status: "awaiting_pickup",
    cycleMinutes: 40,
    currentSessionId: "session-2",
    remainingMinutes: 0,
    updatedAt: now.toISOString()
  },
  {
    id: "washer-d",
    name: "洗濯機 D",
    location: "2階 ランドリー室",
    status: "offline",
    cycleMinutes: 40,
    updatedAt: now.toISOString()
  }
];

export const currentSessionMock: Session = {
  id: "session-1",
  machineId: "washer-b",
  machineName: "洗濯機 B",
  status: "running",
  startedAt: new Date(now.getTime() - 22 * 60 * 1000).toISOString(),
  endsAt: new Date(now.getTime() + 18 * 60 * 1000).toISOString(),
  notificationEnabled: false
};
