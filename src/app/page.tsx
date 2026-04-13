import Link from "next/link";

import { MachineListSection } from "@/components/machine/machine-list-section";
import { SectionCard } from "@/components/ui/section-card";

export default function HomePage() {
  return (
    <div className="space-y-4">
      <SectionCard
        title="いま見たい情報"
        description="まず空きを見て、必要なら利用開始まで進めます。"
      >
        <div className="grid gap-3">
          <Link className="rounded-2xl border border-line px-4 py-3 text-sm" href="/sessions/current">
            自分の利用中セッションを確認する
          </Link>
          <Link
            className="rounded-2xl border border-line px-4 py-3 text-sm"
            href="/settings/notifications"
          >
            通知設定を確認する
          </Link>
        </div>
      </SectionCard>
      <MachineListSection />
    </div>
  );
}
