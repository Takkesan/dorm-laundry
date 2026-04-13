"use client";

import { MachineStatusCard } from "@/components/machine/machine-status-card";
import { EmptyState } from "@/components/ui/empty-state";
import { ErrorState } from "@/components/ui/error-state";
import { LoadingState } from "@/components/ui/loading-state";
import { SectionCard } from "@/components/ui/section-card";
import { useMachinesQuery } from "@/features/machines/queries";

export const MachineListSection = () => {
  const { data, isLoading, isError, refetch, isFetching } = useMachinesQuery();

  return (
    <SectionCard
      title="洗濯機一覧"
      description="空き状況を見て、そのまま詳細と利用開始に進めます。"
    >
      {isLoading || isFetching ? <LoadingState label="洗濯機の状態を取得しています" /> : null}
      {isError ? <ErrorState onRetry={() => void refetch()} /> : null}
      {!isLoading && !isError && data?.length === 0 ? (
        <EmptyState
          title="表示できる洗濯機がありません"
          description="時間をおいて再度確認してください。"
        />
      ) : null}
      {!isLoading && !isError && data?.length ? (
        <div className="grid gap-4">
          {data.map((machine) => (
            <MachineStatusCard key={machine.id} machine={machine} />
          ))}
        </div>
      ) : null}
    </SectionCard>
  );
};
