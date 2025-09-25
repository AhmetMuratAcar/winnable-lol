import "server-only";
import { cache } from "react";
import { redirect } from "next/navigation";
import { regionTagToServerCode } from "@/lib/utils/stringUtils";

export const getProfile = cache(async (params: { region: string; slug: string }) => {
  const { region, slug } = params;
  if (typeof slug !== "string" || !slug.includes("-")) {
    redirect("/summoner-not-found");
  }

  const [rawGameName, rawTagLine] = slug.split("-");
  const gameName = decodeURIComponent(rawGameName || "");
  const tagLine = decodeURIComponent(rawTagLine || "");
  const regionServerCode = regionTagToServerCode(region);

  let res;
  try {
    res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/lol/profile`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ region: regionServerCode, gameName, tagLine }),
    });
  } catch {
    redirect("/server-error?status=500");
  }

  if (res.status === 404) {
    redirect(`/summoner-not-found/${region}/${encodeURIComponent(slug)}`);
  }
  if (!res.ok) {
    redirect(`/server-error?status=${res.status}`);
  }

  return res.json();
});
