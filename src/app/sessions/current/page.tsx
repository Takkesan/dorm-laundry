"use client";

import { NotificationPromptCard } from "@/components/notification/notification-prompt-card";
import { CurrentSessionCard } from "@/components/session/current-session-card";
import { EmptyState } from "@/components/ui/empty-state";
import { ErrorState } from "@/components/ui/error-state";
import { LoadingState } from "@/components/ui/loading-state";
import { SectionCard } from "@/components/ui/section-card";
import { useCurrentSessionQuery } from "@/features/sessions/queries";

export default function CurrentSessionPage() {
  const { data, isLoading, isError, refetch, isFetching } = useCurrentSessionQuery();

  return (
    <div className="space-y-4">
      <SectionCard
        title="現在の利用状況"
        description="自分の利用中セッションだけを表示します。"
      >
        {isLoading || isFetching ? <LoadingState label="利用状況を取得しています" /> : null}
        {isError ? (
          <ErrorState
            title="利用状況を取得できませんでした"
            description="時間をおいて再度お試しください。"
            onRetry={() => void refetch()}
          />
        ) : null}
        {!isLoading && !isError && !data ? (
          <EmptyState
            title="現在の利用はありません"
            description="洗濯機一覧から空きの洗濯機を選んで利用開始できます。"
          />
        ) : null}
        {data ? <CurrentSessionCard session={data} /> : null}
      </SectionCard>
      <NotificationPromptCard />
    </div>
  );
}
