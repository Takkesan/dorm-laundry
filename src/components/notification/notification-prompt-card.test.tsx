import { render, screen } from "@testing-library/react";
import userEvent from "@testing-library/user-event";

import { NotificationPromptCard } from "@/components/notification/notification-prompt-card";

describe("NotificationPromptCard", () => {
  it("未対応端末では案内を表示する", () => {
    // JSDOM で Notification を未対応にして分岐を固定する
    // eslint-disable-next-line @typescript-eslint/no-explicit-any
    delete (window as any).Notification;
    render(<NotificationPromptCard />);
    expect(
      screen.getByText("この端末では通知に対応していません。ブラウザを変えずにそのまま利用できます。")
    ).toBeInTheDocument();
  });

  it("許可ボタンを表示する", async () => {
    const user = userEvent.setup();
    Object.defineProperty(window, "Notification", {
      writable: true,
      value: {
        permission: "default",
        requestPermission: vi.fn().mockResolvedValue("denied")
      }
    });

    render(<NotificationPromptCard />);
    await user.click(screen.getByRole("button", { name: "通知を有効にする" }));
    expect(
      screen.getByText("通知はまだ有効ではありません。必要になったときに再度設定できます。")
    ).toBeInTheDocument();
  });
});
