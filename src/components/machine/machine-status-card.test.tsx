import { render, screen } from "@testing-library/react";

import { MachineStatusCard } from "@/components/machine/machine-status-card";
import { Machine } from "@/types/machine";

const baseMachine: Machine = {
  id: "washer-a",
  name: "洗濯機 A",
  location: "1階 ランドリー室",
  status: "available",
  cycleMinutes: 35,
  updatedAt: "2026-04-13T22:30:00+09:00"
};

describe("MachineStatusCard", () => {
  it("状態ラベルを表示する", () => {
    render(<MachineStatusCard machine={{ ...baseMachine, status: "awaiting_pickup" }} />);
    expect(screen.getByText("回収待ち")).toBeInTheDocument();
  });
});
