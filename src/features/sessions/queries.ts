import { useQuery } from "@tanstack/react-query";

import { getCurrentSession } from "@/lib/api/client";

export const useCurrentSessionQuery = () =>
  useQuery({
    queryKey: ["sessions", "current"],
    queryFn: getCurrentSession
  });
