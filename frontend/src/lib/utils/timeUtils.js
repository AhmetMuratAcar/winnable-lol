export function calcLastUpdated(lastUpdated) {
  const now = new Date();
  const updated = new Date(lastUpdated);
  const timeDiffSec = (now.getTime() - updated.getTime()) / 1000;
  console.log(`NOW: ${now} UPDATED: ${updated} DIFF: ${timeDiffSec}`);

  if (timeDiffSec < 60) {
    return "just now";
  }

  if (timeDiffSec < 3600) {
    const minutes = Math.floor(timeDiffSec / 60);
    return plural(minutes, "minute");
  }

  if (timeDiffSec < 86400) {
    const hours = Math.floor(timeDiffSec / 24);
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
