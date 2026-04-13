import { MachineStatus } from "@/types/machine";

export const machineStatusLabelMap: Record<MachineStatus, string> = {
  available: "空き",
  running: "使用中",
  awaiting_pickup: "回収待ち",
  offline: "利用不可"
};

export const machineStatusToneMap: Record<
  MachineStatus,
  { badge: string; panel: string }
> = {
  available: {
    badge: "bg-emerald-100 text-emerald-800 border-emerald-200",
    panel: "border-emerald-200 bg-emerald-50/80"
  },
  running: {
    badge: "bg-sky-100 text-sky-800 border-sky-200",
    panel: "border-sky-200 bg-sky-50/80"
  },
  awaiting_pickup: {
    badge: "bg-amber-100 text-amber-900 border-amber-200",
    panel: "border-amber-200 bg-amber-50/90"
  },
  offline: {
    badge: "bg-slate-200 text-slate-700 border-slate-300",
    panel: "border-slate-300 bg-slate-100/90"
  }
};
