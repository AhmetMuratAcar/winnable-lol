"use client";
import { ProfileHeader } from ".";
import ProfileNavbar from "./ProfileNavbar";
import RanksSection from "./RanksSection";
import GamesSection from "./GamesSection";
import RecentGames from "./RecentGames";
import PlayedWith from "./PlayedWith";

export default function ProfileSection({ data }) {
  if (!data) {
    return <p>No profile data available.</p>;
  }

  // const { gameName, tagLine, summonerLevel, ranks, masteryData } = data;
  const headerData = {
    gameName: data.gameName,
    tagLine: data.tagLine,
    region: data.region,
    summonerLevel: data.summonerLevel,
    profileIconId: data.profileIconId,
    lastUpdated: data.lastUpdated,
  };

  return (
    <section id="ProfileSection" className="flex-grow flex flex-col items-center gap-6">
      <ProfileHeader headerData={headerData} />
      <div className="w-7/10 space-y-3">
        <ProfileNavbar />
        <div className="flex space-x-4">
          <div id="leftProfile" className="flex flex-col w-3/10 space-y-3">
            <RanksSection rankData={data.ranks} />
            <PlayedWith
              playedWith={data.recentlyPlayedWith}
              playedAgainst={data.recentlyPlayedAgainst}
              gameCount={data.matchData.length}
            />
          </div>

          <div id="rightProfile" className="flex flex-col w-7/10 space-y-3">
            <RecentGames recentGames={data.recentGames} totalGameCount={data.matchData.length}/>
            <GamesSection />
          </div>
        </div>
      </div>
    </section>
  );
}
