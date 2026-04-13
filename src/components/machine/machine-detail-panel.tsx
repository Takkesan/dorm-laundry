"use client";

import Link from "next/link";
import { useState, useTransition } from "react";

import { StatusBadge } from "@/components/ui/status-badge";
import { startSession } from "@/lib/api/client";
import { formatDateTime, formatRelativeMinutes } from "@/lib/utils/format";
import { Machine } from "@/types/machine";

type MachineDetailPanelProps = {
  machine: Machine;
};

export const MachineDetailPanel = ({ machine }: MachineDetailPanelProps) => {
  const [message, setMessage] = useState<string | null>(null);
  const [isPending, startTransition] = useTransition();

  const handleStartSession = () => {
    startTransition(async () => {
      try {
        await startSession(machine.id);
        setMessage("利用開始を受け付けました。利用中画面から通知設定も行えます。");
      } catch {
        setMessage("利用開始に失敗しました。時間をおいてやり直してください。");
      }
    });
  };

  const canStart = machine.status === "available";

  return (
    <div className="space-y-4">
      <div className="rounded-3xl border border-line bg-white p-5 shadow-card">
        <div className="flex items-start justify-between gap-3">
          <div>
            <h1 className="text-2xl font-semibold text-ink">{machine.name}</h1>
            <p className="mt-1 text-sm text-slate-600">{machine.location}</p>
          </div>
          <StatusBadge status={machine.status} />
        </div>
        <dl className="mt-5 space-y-3 text-sm text-slate-700">
          <div className="flex items-center justify-between gap-3">
            <dt>標準運転時間</dt>
            <dd>{machine.cycleMinutes}分</dd>
          </div>
          <div className="flex items-center justify-between gap-3">
            <dt>現在の状況</dt>
            <dd>{formatRelativeMinutes(machine.remainingMinutes)}</dd>
          </div>
          <div className="flex items-center justify-between gap-3">
            <dt>最終更新</dt>
            <dd>{formatDateTime(machine.updatedAt)}</dd>
          </div>
        </dl>
      </div>

      <div className="rounded-3xl border border-line bg-white p-5 shadow-card">
        <h2 className="text-lg font-semibold text-ink">操作</h2>
        <p className="mt-1 text-sm text-slate-600">
          利用開始は空きのときだけ行えます。回収待ちや利用不可の状態では受付しません。
        </p>
        <div className="mt-4 flex flex-wrap gap-3">
          <button
            className="rounded-full bg-ink px-4 py-2 text-sm font-medium text-white disabled:cursor-not-allowed disabled:bg-slate-300"
            disabled={!canStart || isPending}
            onClick={handleStartSession}
            type="button"
          >
            {isPending ? "登録中..." : "利用開始"}
          </button>
          <Link
            className="rounded-full border border-ink px-4 py-2 text-sm font-medium text-ink"
            href="/sessions/current"
          >
            自分の利用状況を見る
          </Link>
        </div>
        {message ? <p className="mt-4 text-sm text-slate-700">{message}</p> : null}
      </div>
    </div>
  );
};
