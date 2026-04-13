import { NotificationPromptCard } from "@/components/notification/notification-prompt-card";
import { SectionCard } from "@/components/ui/section-card";

export default function NotificationSettingsPage() {
  return (
    <div className="space-y-4">
      <SectionCard
        title="通知について"
        description="初回表示では自動で許可を求めず、必要なときだけ設定できるようにしています。"
      >
        <p className="text-sm text-slate-700">
          通知がなくても一覧確認と利用登録は行えます。回収忘れを減らしたい場合だけ有効にしてください。
        </p>
      </SectionCard>
      <NotificationPromptCard />
    </div>
  );
}
