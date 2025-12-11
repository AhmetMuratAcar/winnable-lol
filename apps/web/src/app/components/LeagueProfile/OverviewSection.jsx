"use client";
import RanksSection from "./RanksSection";
import GamesSection from "./GamesSection";
import RecentGames from "./RecentGames";
import PlayedWith from "./PlayedWith";

export default function ProfileOverview({ data }) {
  if (!data) {
    return <p>No profile data available.</p>;
  }

  return (
    <div id="profileOverview" className="flex space-x-4">
      <div id="leftProfile" className="flex flex-col w-3/10 space-y-3">
        <RanksSection rankData={data.ranks} />
        <PlayedWith
          playedWith={data.recentlyPlayedWith}
          playedAgainst={data.recentlyPlayedAgainst}
          gameCount={data.matchPreviews.length}
        />
      </div>

      <div id="rightProfile" className="flex flex-col w-7/10 space-y-3">
        <RecentGames recentGames={data.recentGames} totalGameCount={data.matchPreviews.length} />
        <GamesSection matchPreviews={data.matchPreviews} />
      </div>
    </div>
  );
}
