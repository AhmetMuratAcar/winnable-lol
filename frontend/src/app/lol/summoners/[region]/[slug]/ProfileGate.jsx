// Server component
import { getProfile } from "@/lib/server/lol";
import { ProfileHeader } from "@/app/components/LeagueProfile";
import ProfileNavbar from "@/app/components/LeagueProfile/ProfileNavbar";

export default async function ProfileGate({ params, children }) {
  // This await suspends the rendering of all children so
  // the mastery and live-game tabs dont make their calls to
  // the server if it redirects to an error page.
  const profileData = await getProfile(params);

  const headerData = {
    gameName: profileData.gameName,
    tagLine: profileData.tagLine,
    region: profileData.region,
    summonerLevel: profileData.summonerLevel,
    profileIconId: profileData.profileIconId,
    lastUpdated: profileData.lastUpdated,
  };

  return (
    <>
      <ProfileHeader headerData={headerData} />
      <div className="w-7/10 space-y-3">
        <ProfileNavbar />
        <ProfileDataProvider value={profileData}>{children}</ProfileDataProvider>
      </div>
    </>
  );
}
