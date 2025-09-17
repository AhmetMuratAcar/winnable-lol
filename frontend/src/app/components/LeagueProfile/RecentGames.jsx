"use client";
import Image from "next/image";
import { useState, useMemo } from "react";

const roleKeyToName = {
  all: "All",
  top: "Top",
  jungle: "Jungle",
  middle: "Mid",
  bottom: "Bot",
  utility: "Support",
};

export default function RecentGames({ recentGames, totalGameCount }) {
  // No games played
  if (totalGameCount === 0) {
    return (
      <div className="w-full bg-(--contrast) rounded border-(--contrast-border) border- p-4">
        <h3 className="font-semibold border-b-1 border-(--contrast-border) mb-3">
          Recent Ranked Games
        </h3>
        <p>You have no recent games played</p>
      </div>
    );
  }

  // No ranked games played
  if (recentGames.matchSummaries.length === 0) {
    return (
      <div className="w-full bg-(--contrast) rounded border-(--contrast-border) border- p-4">
        <h3 className="font-semibold border-b-1 border-(--contrast-border) mb-3">
          Recent Ranked Games
        </h3>
        <p>
          You have no ranked games in your last {totalGameCount} game
          {totalGameCount > 1 ? "s" : ""}
        </p>
      </div>
    );
  }

  const [active, setActive] = useState("all");
  const roles = [
    { key: "all", src: "/images/roles/all.svg", rounded: "rounded-l", label: "All" },
    { key: "top", src: "/images/roles/top.svg", label: "Top" },
    { key: "jungle", src: "/images/roles/jungle.svg", label: "Jungle" },
    { key: "middle", src: "/images/roles/middle.svg", label: "Middle" },
    { key: "bottom", src: "/images/roles/bottom.svg", label: "Bottom" },
    { key: "utility", src: "/images/roles/utility.svg", rounded: "rounded-r", label: "Support" },
  ];

  const [activeQueue, setActiveQueue] = useState("all");
  const queues = [
    { key: "all", label: "All", rounded: "rounded-l" },
    { key: 420, label: "Solo" },
    { key: 440, label: "Flex", rounded: "rounded-r" },
  ];

  const roleKeyToApi = (k) => (k === "all" ? "ALL" : k.toUpperCase());

  // Filtering by active role and queue
  const filtered = useMemo(() => {
    if (active === "all") return recentGames.matchSummaries ?? [];
    const apiRole = roleKeyToApi(active);
    return (recentGames.matchSummaries ?? []).filter((m) => m.role === apiRole);
  }, [active, recentGames]);

  const filteredByQueue = useMemo(() => {
    if (activeQueue === "all") return filtered;
    return filtered.filter((m) => m.queueID === activeQueue);
  }, [activeQueue, filtered]);

  // Compute quick stats for the filtered set
  const stats = useMemo(() => {
    if (active === "all" && activeQueue === "all") {
      return {
        wins: recentGames.wins,
        losses: recentGames.losses,
        kda: recentGames.totalKDA,
      };
    }

    let wins = 0;
    let losses = 0;
    for (const g of filteredByQueue) {
      g.didWin ? wins++ : losses++;
    }

    if (active !== "all" && activeQueue === "all") {
      const apiRole = roleKeyToApi(active);
      const fromServer = recentGames.KDAsByRole?.[apiRole];
      return { wins, losses, kda: fromServer !== undefined ? fromServer : 0 };
    }

    let k = 0,
      d = 0,
      a = 0;
    for (const g of filteredByQueue) {
      k += g.kills || 0;
      d += g.deaths || 0;
      a += g.assists || 0;
    }
    const kda = filteredByQueue.length ? (k + a) / Math.max(1, d) : 0;
    return { wins, losses, kda };
  }, [active, activeQueue, filteredByQueue, recentGames]);

  // default
  return (
    <div className="w-full bg-(--contrast) rounded border-(--contrast-border) border- p-4">
      <div className="border-b-1 border-(--contrast-border) mb-3 pb-4 flex items-center justify-between">
        <h3 className="font-semibold">Last {recentGames.matchSummaries.length} Ranked Games</h3>

        {/* Queue filter buttons */}
        <div className="flex justify-center flex-1">
          <div className="flex">
            {queues.map(({ key, label, rounded }) => {
              const isActive = activeQueue === key;
              return (
                <button
                  key={String(key)}
                  onClick={() => setActiveQueue(key)}
                  aria-pressed={isActive}
                  className={`w-16 text-center p-1 border-1 ${rounded ?? ""} cursor-pointer transition-colors ${
                    isActive
                      ? "bg-(--pastel-red) border-white"
                      : "bg-transparent border-(--contrast-border) hover:border-white"
                  }`}
                >
                  {label}
                </button>
              );
            })}
          </div>
        </div>

        {/* Role filter buttons */}
        <div className="flex">
          {roles.map(({ key, src, rounded, label }) => {
            const isActive = active === key;
            return (
              <button
                key={key}
                onClick={() => setActive(key)}
                aria-pressed={isActive}
                className={`p-2 border-1 ${rounded ?? ""} cursor-pointer transition-colors ${
                  isActive
                    ? "bg-(--pastel-red) border-white"
                    : "bg-transparent border-(--contrast-border) hover:border-white"
                }`}
              >
                <Image
                  src={src}
                  width={20}
                  height={20}
                  alt={`${label} role icon`}
                  style={{ width: "20px", height: "20px" }}
                />
              </button>
            );
          })}
        </div>
      </div>

      {/* Displaying raw data for now */}
      <div className="mt-4 text-sm">
        <div className="mb-2 text-gray-300">
          Showing: <span className="font-semibold">{roleKeyToName[active]}</span> •{" "}
          {filteredByQueue.length} game{filteredByQueue.length !== 1 ? "s" : ""} — {stats.wins}W-
          {stats.losses}L • KDA {stats.kda.toFixed(2)}
        </div>

        {filteredByQueue.length === 0 ? (
          <div className="text-gray-400 italic">No games for this selection.</div>
        ) : (
          <ul className="space-y-1">
            {filteredByQueue.map((g, i) => (
              <li
                key={i}
                className="flex items-center justify-between rounded border-1 border-(--contrast-border) px-2 py-1"
              >
                <span className="text-gray-300">
                  Champ #{g.championID} vs #{g.oppChampionID} • {g.role} • Q{g.queueID}
                </span>
                <span className="text-gray-400">
                  {g.kills}/{g.deaths}/{g.assists} • {g.didWin ? "Win" : "Loss"}
                </span>
              </li>
            ))}
          </ul>
        )}
      </div>
    </div>
  );
}
