import { useQuery } from "@tanstack/react-query";

import { getMachine, listMachines } from "@/lib/api/client";

export const useMachinesQuery = () =>
  useQuery({
    queryKey: ["machines"],
    queryFn: listMachines
  });

export const useMachineDetailQuery = (machineId: string) =>
  useQuery({
    queryKey: ["machines", machineId],
    queryFn: () => getMachine(machineId)
  });
