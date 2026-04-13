"use client";

import { useState, useTransition } from "react";

import { useNotificationPermission } from "@/hooks/use-notification-permission";
import { savePushSubscription } from "@/lib/api/client";

export const NotificationPromptCard = () => {
  const { preference, requestPermission } = useNotificationPermission();
  const [message, setMessage] = useState<string | null>(null);
  const [isPending, startTransition] = useTransition();

  const handleEnable = () => {
    startTransition(async () => {
      const result = await requestPermission();

      if (result.permission !== "granted") {
        setMessage("通知はまだ有効ではありません。必要になったときに再度設定できます。");
        return;
      }

      try {
        await savePushSubscription();
        setMessage("通知を有効にしました。洗濯終了が近づいたらお知らせします。");
      } catch {
        setMessage("通知設定の保存に失敗しました。時間をおいて再度お試しください。");
      }
    });
  };

  if (!preference.supported) {
    return (
      <div className="rounded-3xl border border-line bg-white p-5 shadow-card">
        <h2 className="text-lg font-semibold text-ink">通知設定</h2>
        <p className="mt-2 text-sm text-slate-600">
          この端末では通知に対応していません。ブラウザを変えずにそのまま利用できます。
        </p>
      </div>
    );
  }

  return (
    <div className="rounded-3xl border border-line bg-white p-5 shadow-card">
      <h2 className="text-lg font-semibold text-ink">通知設定</h2>
      <p className="mt-2 text-sm text-slate-600">
        利用中の洗濯が終わる頃に気づきやすくなります。必要なときだけ有効にしてください。
      </p>
      <p className="mt-3 text-sm text-slate-700">
        現在の状態:{" "}
        {preference.permission === "granted"
          ? "有効"
          : preference.permission === "denied"
            ? "拒否済み"
            : "未設定"}
      </p>
      <button
        className="mt-4 rounded-full bg-ink px-4 py-2 text-sm font-medium text-white disabled:cursor-not-allowed disabled:bg-slate-300"
        disabled={isPending || preference.permission === "granted"}
        onClick={handleEnable}
        type="button"
      >
        {preference.permission === "granted"
          ? "通知は有効です"
          : isPending
            ? "設定中..."
            : "通知を有効にする"}
      </button>
      {message ? <p className="mt-4 text-sm text-slate-700">{message}</p> : null}
    </div>
  );
};
