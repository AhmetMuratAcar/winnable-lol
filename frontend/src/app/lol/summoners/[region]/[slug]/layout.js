import Header from "@/app/components/Header";
import Footer from "@/app/components/Footer";
import { ProfileHeader } from "@/app/components/LeagueProfile";
import ProfileNavbar from "@/app/components/LeagueProfile/ProfileNavbar";
import { regionTagToServerCode } from "@/lib/utils/stringUtils";
import { redirect } from "next/navigation";
import { ProfileDataProvider } from "./ProfileDataProvider";

async function fetchProfile(resolvedParams) {
  const { region, slug } = resolvedParams;
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
}

export default async function SummonerLayout({ children, params }) {
  const resolvedParams = await params;
  const profileData = await fetchProfile(resolvedParams);

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

      <section className="flex-grow flex flex-col items-center gap-6">
        <ProfileHeader headerData={headerData} />
        <div className="w-7/10 space-y-3">
          <ProfileNavbar />
          <ProfileDataProvider value={profileData}>{children}</ProfileDataProvider>
        </div>
      </section>

      <Footer />
    </main>
  );
}
