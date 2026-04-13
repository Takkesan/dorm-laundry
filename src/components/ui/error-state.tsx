type ErrorStateProps = {
  title?: string;
  description?: string;
  actionLabel?: string;
  onRetry?: () => void;
};

export const ErrorState = ({
  title = "読み込みに失敗しました",
  description = "通信状況を確認して、もう一度お試しください。",
  actionLabel = "再読み込み",
  onRetry
}: ErrorStateProps) => (
  <div className="rounded-3xl border border-rose-200 bg-rose-50 px-4 py-6">
    <p className="font-semibold text-rose-900">{title}</p>
    <p className="mt-2 text-sm text-rose-800">{description}</p>
    {onRetry ? (
      <button
        className="mt-4 rounded-full bg-rose-700 px-4 py-2 text-sm font-medium text-white"
        onClick={onRetry}
        type="button"
      >
        {actionLabel}
      </button>
    ) : null}
  </div>
);
