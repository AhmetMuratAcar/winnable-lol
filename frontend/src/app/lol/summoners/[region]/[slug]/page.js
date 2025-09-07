import Header from "@/app/components/Header";
import Footer from "@/app/components/Footer";
import { ProfileSection } from "@/app/components/LeagueProfile";
import { regionTagToServerCode } from "@/lib/utils/stringUtils";
import { redirect } from "next/navigation";

export default async function SummonerPage({ params }) {
  const { region, slug } = await params;

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
      cache: "no-store",
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

  const profileData = await res.json();
  if (process.env.NODE_ENV === "development") {
    const {
      profileIconId,
      gameName,
      tagLine,
      summonerLevel,
      ranks = [],
      masteryData = {},
      matchData = [],
    } = profileData;

    console.log("Fetched profile data:");
    console.log("  profileIconId:", profileIconId);
    console.log("  gameName:", gameName);
    console.log("  tagLine:", tagLine);
    console.log("  summonerLevel:", summonerLevel);

    console.log("  ranks:");
    ranks.forEach((rank, i) => {
      console.log(`    [${i}]`, rank);
    });

    console.log("  masteryData length:", masteryData?.championMasteries?.length ?? 0);
    console.log("  matchData length:", matchData?.length ?? 0);
  }

  return (
    <main className="min-h-svh flex flex-col bg-(--background)">
      <Header />
      <ProfileSection data={profileData} />
      <Footer />
    </main>
  );
}
