import type { Metadata, Viewport } from "next";
import Link from "next/link";
import { ReactNode } from "react";

import { PwaProvider } from "@/components/providers/pwa-provider";
import { QueryProvider } from "@/components/providers/query-provider";

import "./globals.css";

export const metadata: Metadata = {
  title: "寮ランドリー",
  description: "共同洗濯機の空き状況確認と利用登録を行うPWA"
};

export const viewport: Viewport = {
  themeColor: "#f9fbff",
  width: "device-width",
  initialScale: 1
};

const navLinks = [
  { href: "/", label: "一覧" },
  { href: "/sessions/current", label: "利用中" },
  { href: "/settings/notifications", label: "通知" }
];

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="ja">
      <body>
        <QueryProvider>
          <PwaProvider />
          <div className="mx-auto flex min-h-screen w-full max-w-md flex-col px-4 pb-8 pt-6">
            <header className="mb-6 rounded-3xl border border-white/70 bg-white/80 p-5 shadow-card backdrop-blur">
              <p className="text-sm font-medium uppercase tracking-[0.18em] text-accent">
                Dorm Laundry
              </p>
              <h1 className="mt-2 text-2xl font-semibold text-ink">共同洗濯機の状況</h1>
              <p className="mt-2 text-sm text-slate-600">
                空き状況の確認、利用開始、回収完了までをスマホで手早く行えます。
              </p>
              <nav className="mt-4 flex gap-2 overflow-x-auto pb-1">
                {navLinks.map((link) => (
                  <Link
                    key={link.href}
                    className="rounded-full border border-line bg-white px-4 py-2 text-sm text-slate-700"
                    href={link.href}
                  >
                    {link.label}
                  </Link>
                ))}
              </nav>
            </header>
            <main className="flex-1">{children}</main>
          </div>
        </QueryProvider>
      </body>
    </html>
  );
}
