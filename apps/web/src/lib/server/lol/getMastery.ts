import "server-only";
import { redirect } from "next/navigation";
import { regionTagToServerCode } from "@/lib/utils/stringUtils";

type Params = { region: string; slug: string };

// TODO: replace unknown with mastery response type.
export async function getMastery(params: Params): Promise<unknown> {
  const { region, slug } = params;

  if (typeof slug !== "string" || !slug.includes("-")) {
    redirect("/summoner-not-found");
  }

  const [rawGameName, rawTagLine] = slug.split("-");
  const gameName = decodeURIComponent(rawGameName || "");
  const tagLine = decodeURIComponent(rawTagLine || "");
  const regionServerCode = regionTagToServerCode(region);

  const url = new URL(`${process.env.NEXT_PUBLIC_BACKEND_URL}/lol/mastery`);
  url.searchParams.set("region", regionServerCode);
  url.searchParams.set("gameName", gameName);
  url.searchParams.set("tagLine", tagLine);

  const res = await fetch(url.toString(), {
    next: {
      revalidate: 300, // 5 minutes
      tags: [`mastery:${region}:${slug}`],
    },
  });

  if (res.status === 404) {
    redirect(`/summoner-not-found/${region}/${gameName}-${tagLine}`);
  }
  if (!res.ok) {
    redirect(`/server-error?status=${res.status}`);
  }

  return res.json();
}
