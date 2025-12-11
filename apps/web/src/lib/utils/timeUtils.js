export function calcLastUpdated(lastUpdated) {
  const now = new Date();
  const updated = new Date(lastUpdated);
  const timeDiffSec = (now.getTime() - updated.getTime()) / 1000;
  // console.log(`NOW: ${now} UPDATED: ${updated} DIFF: ${timeDiffSec}`);

  if (timeDiffSec < 60) {
    return "just now";
  }

  if (timeDiffSec < 3600) {
    const minutes = Math.floor(timeDiffSec / 60);
    return plural(minutes, "minute");
  }

  if (timeDiffSec < 86400) {
    const hours = Math.floor(timeDiffSec / 3600);
    return plural(hours, "hour");
  }

  if (timeDiffSec < 604800) {
    const days = Math.floor(timeDiffSec / 86400);
    return `over ${plural(days, "day")}`;
  }

  if (timeDiffSec < 2592000) {
    const weeks = Math.floor(timeDiffSec / 604800);
    return `over ${plural(weeks, "week")}`;
  }

  const months = Math.floor(timeDiffSec / 2592000);
  return `over ${plural(months, "month")}`;
}

function plural(num, unit) {
  return `${num} ${unit}${num === 1 ? "" : "s"} ago`;
}

export function totalToReadable(gameDuration) {
  const minutes = Math.floor(gameDuration / 60);
  const seconds = gameDuration % 60;

  if (seconds.toString().length === 1) {
    return `${minutes}:0${seconds}`;
  }

  return `${minutes}:${seconds}`;
}

export function calcWhenPlayed(gameEndTimestamp) {
  const now = Date.now();
  const diffMs = now - gameEndTimestamp;
  const diffSec = diffMs / 1000;
  const diffMin = diffSec / 60;
  const diffHr = diffMin / 60;
  const diffDay = diffHr / 24;
  const diffWeek = diffDay / 7;
  const diffMonth = diffDay / 30;
  const diffYear = diffDay / 365;

  if (diffMin < 1) return "Just now";
  if (diffHr < 1) return `${Math.round(diffMin)} minute${Math.round(diffMin) !== 1 ? "s" : ""} ago`;
  if (diffDay < 1) return `${Math.round(diffHr)} hour${Math.round(diffHr) !== 1 ? "s" : ""} ago`;
  if (diffWeek < 1) return `${Math.round(diffDay)} day${Math.round(diffDay) !== 1 ? "s" : ""} ago`;
  if (diffMonth < 1)
    return `${Math.round(diffWeek)} week${Math.round(diffWeek) !== 1 ? "s" : ""} ago`;
  if (diffYear < 1)
    return `${Math.round(diffMonth)} month${Math.round(diffMonth) !== 1 ? "s" : ""} ago`;
  return `${Math.round(diffYear)} year${Math.round(diffYear) !== 1 ? "s" : ""} ago`;
}
