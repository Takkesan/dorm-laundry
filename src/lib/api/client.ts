import { z } from "zod";

import { currentSessionMock, machinesMock } from "@/lib/api/mock-data";
import { Machine } from "@/types/machine";
import { Session } from "@/types/session";

const machineSchema = z.object({
  id: z.string(),
  name: z.string(),
  location: z.string(),
  status: z.enum(["available", "running", "awaiting_pickup", "offline"]),
  cycleMinutes: z.number(),
  currentSessionId: z.string().optional(),
  remainingMinutes: z.number().optional(),
  updatedAt: z.string()
});

const sessionSchema = z.object({
  id: z.string(),
  machineId: z.string(),
  machineName: z.string(),
  status: z.enum(["running", "awaiting_pickup"]),
  startedAt: z.string(),
  endsAt: z.string(),
  notificationEnabled: z.boolean()
});

const wait = async (ms = 250) => {
  await new Promise((resolve) => setTimeout(resolve, ms));
};

export const listMachines = async (): Promise<Machine[]> => {
  await wait();
  return z.array(machineSchema).parse(machinesMock);
};

export const getMachine = async (machineId: string): Promise<Machine> => {
  await wait();
  const machine = machinesMock.find((item) => item.id === machineId);

  if (!machine) {
    throw new Error("machine_not_found");
  }

  return machineSchema.parse(machine);
};

export const getCurrentSession = async (): Promise<Session | null> => {
  await wait();
  return sessionSchema.parse(currentSessionMock);
};

export const startSession = async (machineId: string): Promise<{ ok: true; machineId: string }> => {
  await wait();
  return { ok: true, machineId };
};

export const claimSession = async (sessionId: string): Promise<{ ok: true; sessionId: string }> => {
  await wait();
  return { ok: true, sessionId };
};

export const savePushSubscription = async (): Promise<{ ok: true }> => {
  await wait();
  return { ok: true };
};
