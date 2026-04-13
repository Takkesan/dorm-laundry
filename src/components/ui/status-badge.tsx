import { machineStatusLabelMap, machineStatusToneMap } from "@/lib/constants/machine-status";
import { MachineStatus } from "@/types/machine";

type StatusBadgeProps = {
  status: MachineStatus;
};

export const StatusBadge = ({ status }: StatusBadgeProps) => (
  <span
    className={`inline-flex items-center rounded-full border px-3 py-1 text-sm font-medium ${machineStatusToneMap[status].badge}`}
  >
    {machineStatusLabelMap[status]}
  </span>
);
