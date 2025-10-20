import "server-only";
import { unstable_cache } from "next/cache";
import { redirect } from "next/navigation";
import { regionTagToServerCode } from "@/lib/utils/stringUtils";

type Params = { region: string; slug: string };
// TODO: add ProfileData type

export async function getProfile(params: Params) {
  const { region, slug } = params;

  // TODO: this endpoint is now a GET on the server, update to use next caching instead of unstable
  const cached = unstable_cache(
    async () => {
      if (typeof slug !== "string" || !slug.includes("-")) {
        redirect("/summoner-not-found");
      }

      const [rawGameName, rawTagLine] = slug.split("-");
      const gameName = decodeURIComponent(rawGameName || "");
      const tagLine = decodeURIComponent(rawTagLine || "");
      const regionServerCode = regionTagToServerCode(region);

      const url = new URL(`${process.env.NEXT_PUBLIC_BACKEND_URL}/lol/profile`);
      url.searchParams.set("region", regionServerCode);
      url.searchParams.set("gameName", gameName);
      url.searchParams.set("tagLine", tagLine);

      const res = await fetch(url);

      if (res.status === 404) {
        redirect(`/summoner-not-found/${region}/${gameName}-${tagLine}`);
      }
      if (!res.ok) {
        redirect(`/server-error?status=${res.status}`);
      }

      return res.json();
    },
    ["profile", region, slug],
    { revalidate: 1, tags: [`profile:${region}:${slug}`] }, // 5 minutes
  );

  return cached();
}
