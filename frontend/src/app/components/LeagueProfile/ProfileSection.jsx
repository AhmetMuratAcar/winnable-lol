"use client";
import { ProfileHeader } from ".";

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
    </section>
  );
}
