import Header from "@/app/components/Header";
import Footer from "@/app/components/Footer";
import ProfileNavbar from "@/app/components/LeagueProfile/ProfileNavbar";
import { ProfileHeader, ProfileOverview } from "@/app/components/LeagueProfile";
import { regionTagToServerCode } from "@/lib/utils/stringUtils";
import { redirect } from "next/navigation";

export async function generateMetadata({ params }) {
  const { slug } = await params;

  if (typeof slug !== "string" || !slug.includes("-")) {
    return { title: "Invalid Summoner | Winnable" };
  }

  const [rawGameName, rawTagLine] = slug.split("-");
  const gameName = decodeURIComponent(rawGameName || "");
  const tagLine = decodeURIComponent(rawTagLine || "");

  return {
    title: `${gameName} ${tagLine ? `#${tagLine}` : ""} - Summoner Stats | Winnable`,
    description: `Summoner stats for ${gameName}${tagLine ? ` #${tagLine}` : ""}.`,
  };
}

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
  const headerData = {
    gameName: profileData.gameName,
    tagLine: profileData.tagLine,
    region: profileData.region,
    summonerLevel: profileData.summonerLevel,
    profileIconId: profileData.profileIconId,
    lastUpdated: profileData.lastUpdated,
  };

  return (
    <main className="min-h-svh flex flex-col bg-(--background)">
      <Header />

      <section id="ProfileOverview" className="flex-grow flex flex-col items-center gap-6">
        <ProfileHeader headerData={headerData} />

        <div className="w-7/10 space-y-3">
          <ProfileNavbar />
          <ProfileOverview data={profileData} />
        </div>
      </section>

      <Footer />
    </main>
  );
}
