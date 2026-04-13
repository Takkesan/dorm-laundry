import Link from "next/link";

import { MachineDetailPanel } from "@/components/machine/machine-detail-panel";
import { ErrorState } from "@/components/ui/error-state";
import { getMachine } from "@/lib/api/client";

type MachineDetailPageProps = {
  params: Promise<{
    machineId: string;
  }>;
};

export default async function MachineDetailPage({ params }: MachineDetailPageProps) {
  const { machineId } = await params;

  try {
    const machine = await getMachine(machineId);

    return (
      <div className="space-y-4">
        <Link className="text-sm text-slate-600" href="/">
          一覧に戻る
        </Link>
        <MachineDetailPanel machine={machine} />
      </div>
    );
  } catch {
    return (
      <div className="space-y-4">
        <Link className="text-sm text-slate-600" href="/">
          一覧に戻る
        </Link>
        <ErrorState
          title="洗濯機情報が見つかりません"
          description="URLを確認して、一覧画面から再度選択してください。"
        />
      </div>
    );
  }
}
