"use client";

import { useMachinesQuery } from "@/features/machines/queries";
import { machineStatusLabelMap } from "@/lib/constants/machine-status";

export const AdminMachineSummary = () => {
  const { data = [] } = useMachinesQuery();

  return (
    <div className="rounded-3xl border border-line bg-white p-5 shadow-card">
      <h1 className="text-xl font-semibold text-ink">管理画面</h1>
      <p className="mt-1 text-sm text-slate-600">状態確認に必要な最小情報のみ表示します。</p>
      <ul className="mt-4 space-y-3">
        {data.map((machine) => (
          <li
            key={machine.id}
            className="flex items-center justify-between gap-3 rounded-2xl border border-line px-4 py-3 text-sm"
          >
            <span>{machine.name}</span>
            <span className="font-medium text-slate-700">
              {machineStatusLabelMap[machine.status]}
            </span>
          </li>
        ))}
      </ul>
    </div>
  );
};
