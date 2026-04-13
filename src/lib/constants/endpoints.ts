export const endpoints = {
  machines: "/machines",
  machine: (machineId: string) => `/machines/${machineId}`,
  sessions: "/sessions",
  currentSession: "/sessions/current",
  claimSession: (sessionId: string) => `/sessions/${sessionId}/claim`,
  pushSubscriptions: "/push-subscriptions"
} as const;
