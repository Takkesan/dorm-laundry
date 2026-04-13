import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { render, screen } from "@testing-library/react";

import CurrentSessionPage from "@/app/sessions/current/page";

vi.mock("@/features/sessions/queries", () => ({
  useCurrentSessionQuery: () => ({
    data: null,
    isLoading: false,
    isFetching: false,
    isError: false,
    refetch: vi.fn()
  })
}));

const renderPage = () =>
  render(
    <QueryClientProvider client={new QueryClient()}>
      <CurrentSessionPage />
    </QueryClientProvider>
  );

describe("CurrentSessionPage", () => {
  it("空状態を表示する", () => {
    renderPage();
    expect(screen.getByText("現在の利用はありません")).toBeInTheDocument();
  });
});
