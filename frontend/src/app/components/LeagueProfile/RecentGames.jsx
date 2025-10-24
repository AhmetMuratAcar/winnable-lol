"use client";
import Image from "next/image";
import dynamic from "next/dynamic";
import { useState, useMemo } from "react";
import { champIdToName } from "@/lib/utils/champs";
import { IMG_PATH } from "@/lib/constants";
import RoleBarGraph from "../charts/RoleBarGraph";

const WinLossDonut = dynamic(() => import("@/app/components/charts/WinLossDonut"), {
  ssr: false,
});

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
        <p className="text-gray-400 italic">You have no recent games played</p>
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
        <p className="text-gray-400 italic">
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

  const roleKeyToApi = (k) => k.toUpperCase();

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
    const emptyRoleWL = { wins: 0, losses: 0 };
    const emptyRolesWL = {
      top: { ...emptyRoleWL },
      jungle: { ...emptyRoleWL },
      mid: { ...emptyRoleWL },
      bot: { ...emptyRoleWL },
      utility: { ...emptyRoleWL },
    };

    const empty = {
      wins: 0,
      losses: 0,
      kda: null,
      avgKills: null,
      avgDeaths: null,
      avgAssists: null,
      roles: emptyRolesWL,
    };

    if (!recentGames) return empty;

    const summaries = recentGames.matchSummaries ?? [];
    const apiRole = roleKeyToApi(active);
    const qKey = String(activeQueue);

    const normalizeRole = (r) => {
      const s = r?.toLowerCase();
      if (s === "middle") return "mid";
      if (s === "bottom") return "bot";
      return s;
    };

    const base =
      activeQueue !== "all" && Array.isArray(recentGames.filteredByQueue)
        ? recentGames.filteredByQueue
        : summaries;

    const filtered = base.filter((m) => {
      const okRole = active === "all" ? true : m.role === apiRole;
      const okQueue = activeQueue === "all" ? true : String(m.queueID) === qKey;
      return okRole && okQueue;
    });

    if (!filtered.length) return empty;

    // data for role bar chart
    const roleWinLoss =
      active === "all"
        ? filtered.reduce(
            (acc, m) => {
              const key = normalizeRole(m.role);
              if (!acc[key]) acc[key] = { wins: 0, losses: 0 };
              if (m.didWin) acc[key].wins += 1;
              else acc[key].losses += 1;
              return acc;
            },
            { ...emptyRolesWL },
          )
        : { ...emptyRolesWL };

    const fromTotals = (t) => {
      if (!t || !t.games) return { ...empty, roles: roleWinLoss };
      const { kills, deaths, assists, wins, losses, games } = t;
      const kda = deaths === 0 ? kills + assists : (kills + assists) / deaths;
      return {
        wins,
        losses,
        kda,
        avgKills: kills / games,
        avgDeaths: deaths / games,
        avgAssists: assists / games,
        roles: roleWinLoss,
      };
    };

    if (active === "all" && activeQueue === "all" && recentGames.totalsAll) {
      return fromTotals(recentGames.totalsAll);
    }
    if (active !== "all" && activeQueue === "all") {
      const t = recentGames.totalsByRole?.[apiRole];
      if (t) return fromTotals(t);
    }
    if (active === "all" && activeQueue !== "all") {
      const t = recentGames.totalsByQueue?.[qKey];
      if (t) return fromTotals(t);
    }

    // Fallback: compute aggregate stats from filtered
    const agg = filtered.reduce(
      (a, m) => {
        a.k += m.kills;
        a.d += m.deaths;
        a.a += m.assists;
        m.didWin ? a.w++ : a.l++;
        return a;
      },
      { k: 0, d: 0, a: 0, w: 0, l: 0 },
    );

    const games = filtered.length;
    const kda = agg.d === 0 ? agg.k + agg.a : (agg.k + agg.a) / agg.d;

    return {
      wins: agg.w,
      losses: agg.l,
      kda,
      avgKills: agg.k / games,
      avgDeaths: agg.d / games,
      avgAssists: agg.a / games,
      roles: roleWinLoss, // <-- new shape
    };
  }, [active, activeQueue, filteredByQueue, recentGames]);

  function kdaFromTotals(k, d, a) {
    return d === 0 ? k + a : (k + a) / d;
  }

  const topChamps = useMemo(() => {
    if (!recentGames) return [];

    const summaries = recentGames.matchSummaries ?? [];
    const apiRole = roleKeyToApi(active);
    const qKey = String(activeQueue);

    // prefer the pre-filtered array for the queue
    const base =
      activeQueue !== "all" && Array.isArray(recentGames.filteredByQueue)
        ? recentGames.filteredByQueue
        : summaries;

    // Apply chosen filters (role + queue)
    const filtered = base.filter((m) => {
      const okRole = active === "all" ? true : m.role === apiRole;
      const okQueue = activeQueue === "all" ? true : String(m.queueID) === qKey;
      return okRole && okQueue;
    });

    if (!filtered.length) return [];

    // Aggregate per championID from the FILTERED set
    const byChamp = new Map();
    for (const m of filtered) {
      const champName = champIdToName(m.championID);
      const c = byChamp.get(champName) ?? {
        championName: champName,
        games: 0,
        wins: 0,
        losses: 0,
        kills: 0,
        deaths: 0,
        assists: 0,
      };

      c.games += 1;
      c.kills += m.kills;
      c.deaths += m.deaths;
      c.assists += m.assists;
      m.didWin ? c.wins++ : c.losses++;

      byChamp.set(champName, c);
    }

    const enriched = [...byChamp.values()].map((c) => ({
      championName: c.championName,
      games: c.games,
      wins: c.wins,
      losses: c.losses,
      kda: kdaFromTotals(c.kills, c.deaths, c.assists),
      avgKills: c.kills / c.games,
      avgDeaths: c.deaths / c.games,
      avgAssists: c.assists / c.games,
      winRate: c.wins / c.games,
    }));

    // Sort: most games → higher winRate → higher KDA
    enriched.sort((a, b) => b.games - a.games || b.winRate - a.winRate || b.kda - a.kda);

    return enriched.slice(0, 3);
  }, [recentGames, active, activeQueue, filteredByQueue]);

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
                  style={{ width: "20px", height: "auto" }}
                />
              </button>
            );
          })}
        </div>
      </div>

      <div className="mt-4 text-sm">
        {(() => {
          const summaries = recentGames?.matchSummaries ?? [];
          const totalsAll = recentGames?.totalsAll ?? recentGames?.totalsALl; // tolerate typo
          const apiRole = roleKeyToApi(active).toUpperCase();
          const qKey = String(activeQueue);

          // Decide which games to render in the list
          const displayedGames =
            active !== "all" && activeQueue !== "all"
              ? summaries.filter((m) => m.role === apiRole && String(m.queueID) === qKey)
              : active === "all" && activeQueue !== "all"
                ? filteredByQueue
                : active !== "all" && activeQueue === "all"
                  ? summaries.filter((m) => m.role === apiRole)
                  : summaries; // all/all

          // Show count from pre-aggregates when possible
          const shownCount =
            active === "all" && activeQueue === "all"
              ? (totalsAll?.games ?? displayedGames.length)
              : active !== "all" && activeQueue === "all"
                ? (recentGames?.totalsByRole?.[apiRole]?.games ?? displayedGames.length)
                : active === "all" && activeQueue !== "all"
                  ? (recentGames?.totalsByQueue?.[qKey]?.games ?? displayedGames.length)
                  : displayedGames.length;

          const kdaText = stats.kda == null ? "—" : stats.kda.toFixed(2);

          return (
            <>
              <div className="mb-2 text-gray-300">
                Role: <span className="font-semibold">{roleKeyToName[active]}</span> • {shownCount}{" "}
                game{shownCount !== 1 ? "s" : ""}: {stats.wins}W - {stats.losses}L
              </div>

              {shownCount === 0 ? (
                <div className="text-gray-400 italic">No games for this selection.</div>
              ) : (
                <div>
                  <div
                    className={
                      active === "all"
                        ? "mb-3 flex items-center gap-6" // row w/ bar graph
                        : "mb-3 flex justify-center gap-12" // center stats + champs
                    }
                  >
                    {/* donut + stats */}
                    <div id="stats" className="flex items-center gap-4">
                      <WinLossDonut wins={stats.wins} losses={stats.losses} />
                      <div className="flex flex-col text-center">
                        <p className="font-bold">{kdaText} KDA</p>
                        <p className="text-gray-300">
                          {stats.avgKills.toFixed(2)} / {stats.avgDeaths.toFixed(2)} /{" "}
                          {stats.avgAssists.toFixed(2)}
                        </p>
                      </div>
                    </div>

                    <div id="topChamps" className="space-y-2">
                      {topChamps.map((c, i) => (
                        <div key={i} className="flex items-center text-gray-300 text-xs">
                          <Image
                            src={`${IMG_PATH}/img/champion/tiles/${c.championName}_0.jpg`}
                            width={25}
                            height={25}
                            className="rounded mr-2 border border-(--contrast-border) img-auto"
                            alt={`${c.championName} image`}
                          />
                          <div className="flex flex-col">
                            <p>
                              {(c.winRate * 100).toFixed()}% ({c.wins}W - {c.losses}L)
                            </p>
                            <p>{c.kda.toFixed(2)} KDA</p>
                          </div>
                        </div>
                      ))}
                    </div>

                    {/* bar graph only for all */}
                    {active === "all" && (
                      <div className="flex-1 min-w-0 h-28 self-center">
                        <RoleBarGraph roles={stats.roles} />
                      </div>
                    )}
                  </div>

                  {/* game list stays the same */}
                  {/* <ul className="space-y-1">
                    {displayedGames.map((g, i) => (
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
                  </ul> */}
                </div>
              )}
            </>
          );
        })()}
      </div>
    </div>
  );
}
