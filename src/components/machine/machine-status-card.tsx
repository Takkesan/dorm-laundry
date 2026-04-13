import Link from "next/link";

import { StatusBadge } from "@/components/ui/status-badge";
import { machineStatusToneMap } from "@/lib/constants/machine-status";
import { formatRelativeMinutes } from "@/lib/utils/format";
import { Machine } from "@/types/machine";

type MachineStatusCardProps = {
  machine: Machine;
};

export const MachineStatusCard = ({ machine }: MachineStatusCardProps) => (
  <article
    className={`rounded-3xl border p-5 shadow-card ${machineStatusToneMap[machine.status].panel}`}
  >
    <div className="flex items-start justify-between gap-3">
      <div>
        <p className="text-base font-semibold text-ink">{machine.name}</p>
        <p className="mt-1 text-sm text-slate-600">{machine.location}</p>
      </div>
      <StatusBadge status={machine.status} />
    </div>
    <dl className="mt-4 space-y-2 text-sm text-slate-700">
      <div className="flex items-center justify-between gap-3">
        <dt>運転時間</dt>
        <dd>{machine.cycleMinutes}分</dd>
      </div>
      <div className="flex items-center justify-between gap-3">
        <dt>状況</dt>
        <dd>{formatRelativeMinutes(machine.remainingMinutes)}</dd>
      </div>
    </dl>
    <div className="mt-5 flex gap-3">
      <Link
        className="rounded-full bg-ink px-4 py-2 text-sm font-medium text-white"
        href={`/machines/${machine.id}`}
      >
        詳細を見る
      </Link>
      <Link
        className="rounded-full border border-ink px-4 py-2 text-sm font-medium text-ink"
        href="/sessions/current"
      >
        利用中を確認
      </Link>
    </div>
  </article>
);
