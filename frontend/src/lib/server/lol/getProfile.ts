import "server-only";
import { unstable_cache } from "next/cache";
import { redirect } from "next/navigation";
import { regionTagToServerCode } from "@/lib/utils/stringUtils";

type Params = { region: string; slug: string };
// TODO: add ProfileData type

export async function getProfile(params: Params) {
  const { region, slug } = params;

  const cached = unstable_cache(
    async () => {
      if (typeof slug !== "string" || !slug.includes("-")) {
        redirect("/summoner-not-found");
      }

      const [rawGameName, rawTagLine] = slug.split("-");
      const gameName = decodeURIComponent(rawGameName || "");
      const tagLine = decodeURIComponent(rawTagLine || "");
      const regionServerCode = regionTagToServerCode(region);

      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/lol/profile`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ region: regionServerCode, gameName, tagLine }),
      });

      if (res.status === 404) {
        redirect(`/summoner-not-found/${region}/${encodeURIComponent(slug)}`);
      }
      if (!res.ok) {
        redirect(`/server-error?status=${res.status}`);
      }

      return res.json();
    },
    ["profile", region, slug],
    { revalidate: 300, tags: [`profile:${region}:${slug}`] }, // 5 minutes
  );

  return cached();
}
