"use client";

import { useState, useTransition } from "react";

import { claimSession } from "@/lib/api/client";
import { formatDateTime } from "@/lib/utils/format";
import { Session } from "@/types/session";

type CurrentSessionCardProps = {
  session: Session;
};

export const CurrentSessionCard = ({ session }: CurrentSessionCardProps) => {
  const [message, setMessage] = useState<string | null>(null);
  const [isPending, startTransition] = useTransition();

  const handleClaim = () => {
    startTransition(async () => {
      try {
        await claimSession(session.id);
        setMessage("回収完了を受け付けました。");
      } catch {
        setMessage("回収完了を登録できませんでした。もう一度お試しください。");
      }
    });
  };

  return (
    <article className="rounded-3xl border border-line bg-white p-5 shadow-card">
      <p className="text-lg font-semibold text-ink">{session.machineName}</p>
      <dl className="mt-4 space-y-3 text-sm text-slate-700">
        <div className="flex items-center justify-between gap-3">
          <dt>状態</dt>
          <dd>{session.status === "running" ? "使用中" : "回収待ち"}</dd>
        </div>
        <div className="flex items-center justify-between gap-3">
          <dt>利用開始</dt>
          <dd>{formatDateTime(session.startedAt)}</dd>
        </div>
        <div className="flex items-center justify-between gap-3">
          <dt>終了予定</dt>
          <dd>{formatDateTime(session.endsAt)}</dd>
        </div>
      </dl>

      <div className="mt-5 flex flex-wrap gap-3">
        <button
          className="rounded-full bg-ink px-4 py-2 text-sm font-medium text-white disabled:cursor-not-allowed disabled:bg-slate-300"
          disabled={isPending}
          onClick={handleClaim}
          type="button"
        >
          {isPending ? "登録中..." : "回収完了"}
        </button>
      </div>
      {message ? <p className="mt-4 text-sm text-slate-700">{message}</p> : null}
    </article>
  );
};
