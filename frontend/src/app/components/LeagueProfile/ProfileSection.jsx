"use client";
import { ProfileHeader } from ".";

export default function ProfileSection({ data }) {
  if (!data) {
    return <p>No profile data available.</p>;
  }

  const { gameName, tagLine, summonerLevel, ranks, masteryData } = data;
  const headerData = {
    gameName: data.gameName,
    tagLine: data.tagLine,
    region: data.region,
    summonerLevel: data.summonerLevel,
    profileIconId: data.profileIconId,
  };

  const rank = ranks && ranks.length > 0 ? ranks[0] : null;
  const topChampions = masteryData?.championMasteries?.slice(0, 3) || [];

  return (
    <section
      id="profileSection"
      className="flex-grow flex flex-col items-center gap-6 py-6"
    >
      <ProfileHeader headerData={headerData} />

      {/* Basic info */}
      <div className="text-center">
        <h2 className="text-2xl font-bold">
          {gameName}#{tagLine}
        </h2>
        <p className="text-gray-400">Level {summonerLevel}</p>
      </div>

      {/* Rank info */}
      {rank && (
        <div className="bg-[var(--contrast)] text-white px-6 py-4 rounded-xl shadow">
          <p className="font-semibold">
            Ranked {rank.queueType.replaceAll("_", " ")}
          </p>
          <p>
            {rank.tier} {rank.rank} - {rank.leaguePoints} LP
          </p>
          <p>
            {rank.wins}W / {rank.losses}L
          </p>
        </div>
      )}

      {/* Top mastery champs */}
      <div className="w-full max-w-md space-y-3">
        <h3 className="text-xl font-semibold">Top Champions</h3>
        {topChampions.map((c) => (
          <div
            key={c.championId}
            className="flex justify-between bg-[var(--contrast)] text-white px-4 py-2 rounded-lg"
          >
            <span>Champion ID: {c.championId}</span>
            <span>
              Lvl {c.championLevel} • {c.championPoints.toLocaleString()} pts
            </span>
          </div>
        ))}
      </div>
    </section>
  );
}
