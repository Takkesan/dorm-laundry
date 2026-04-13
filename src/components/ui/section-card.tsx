import { ReactNode } from "react";

type SectionCardProps = {
  title?: string;
  description?: string;
  children: ReactNode;
};

export const SectionCard = ({ title, description, children }: SectionCardProps) => (
  <section className="rounded-3xl border border-line bg-white p-5 shadow-card">
    {title ? <h2 className="text-lg font-semibold text-ink">{title}</h2> : null}
    {description ? <p className="mt-1 text-sm text-slate-600">{description}</p> : null}
    <div className={title || description ? "mt-4" : ""}>{children}</div>
  </section>
);
