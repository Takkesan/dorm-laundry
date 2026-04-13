export const formatRelativeMinutes = (minutes?: number) => {
  if (minutes === undefined) {
    return "残り時間不明";
  }

  if (minutes <= 0) {
    return "まもなく終了";
  }

  return `残り約${minutes}分`;
};

export const formatDateTime = (iso: string) =>
  new Intl.DateTimeFormat("ja-JP", {
    month: "numeric",
    day: "numeric",
    hour: "2-digit",
    minute: "2-digit"
  }).format(new Date(iso));
