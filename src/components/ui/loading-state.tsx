type LoadingStateProps = {
  label?: string;
};

export const LoadingState = ({ label = "読み込み中です" }: LoadingStateProps) => (
  <div
    className="rounded-3xl border border-dashed border-line bg-white/70 px-4 py-8 text-center text-sm text-slate-600"
    role="status"
  >
    {label}
  </div>
);
