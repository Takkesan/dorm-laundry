type EmptyStateProps = {
  title: string;
  description: string;
};

export const EmptyState = ({ title, description }: EmptyStateProps) => (
  <div className="rounded-3xl border border-dashed border-line bg-white/70 px-4 py-8 text-center">
    <p className="font-semibold text-ink">{title}</p>
    <p className="mt-2 text-sm text-slate-600">{description}</p>
  </div>
);
