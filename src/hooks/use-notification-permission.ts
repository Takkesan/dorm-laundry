"use client";

import { useEffect, useState } from "react";

import { NotificationPreference } from "@/types/notification";

const defaultPreference: NotificationPreference = {
  permission: "default",
  supported: false
};

export const useNotificationPermission = () => {
  const [preference, setPreference] = useState<NotificationPreference>(defaultPreference);

  useEffect(() => {
    if (typeof window === "undefined" || !("Notification" in window)) {
      setPreference(defaultPreference);
      return;
    }

    setPreference({
      permission: Notification.permission,
      supported: true
    });
  }, []);

  const requestPermission = async () => {
    if (typeof window === "undefined" || !("Notification" in window)) {
      return defaultPreference;
    }

    const permission = await Notification.requestPermission();
    const nextPreference = {
      permission,
      supported: true
    } satisfies NotificationPreference;

    setPreference(nextPreference);
    return nextPreference;
  };

  return {
    preference,
    requestPermission
  };
};
