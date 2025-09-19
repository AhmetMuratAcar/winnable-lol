"use client";
import Image from "next/image";
import dynamic from "next/dynamic";
import { useState, useMemo } from "react";

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
    if (!recentGames) {
      return { wins: 0, losses: 0, kda: 0.0 };
    }

    const summaries = recentGames.matchSummaries ?? [];
    const apiRole = roleKeyToApi(active);
    const qKey = String(activeQueue);

    // : { kills; deaths; assists; wins; losses; games }
    const kdaFromTotals = (t) => {
      if (!t || !t.games) return { wins: 0, losses: 0, kda: 0.0 };
      const { kills, deaths, assists, wins, losses } = t;
      const kda = deaths === 0 ? kills + assists : (kills + assists) / deaths;
      return { wins, losses, kda };
    };

    if (active === "all" && activeQueue === "all" && recentGames.totalsAll) {
      return kdaFromTotals(recentGames.totalsAll);
    }

    if (active !== "all" && activeQueue === "all") {
      const t = recentGames.totalsByRole?.[apiRole];
      if (t) return kdaFromTotals(t);
      return { wins: 0, losses: 0, kda: null };
    }

    if (active === "all" && activeQueue !== "all") {
      const t = recentGames.totalsByQueue?.[qKey];
      if (t) return kdaFromTotals(t);
      return { wins: 0, losses: 0, kda: null };
    }

    const base =
      activeQueue !== "all" && Array.isArray(recentGames.filteredByQueue)
        ? recentGames.filteredByQueue // if you keep a pre-filtered array around
        : summaries;

    const filtered = base.filter((m) => {
      const okRole = active === "all" ? true : m.role === apiRole;
      const okQueue = activeQueue === "all" ? true : String(m.queueID) === qKey;
      return okRole && okQueue;
    });

    if (!filtered.length) return { wins: 0, losses: 0, kda: null };

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

    const kda = agg.d === 0 ? agg.k + agg.a : (agg.k + agg.a) / agg.d;
    return { wins: agg.w, losses: agg.l, kda };
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
                  style={{ width: "20px", height: "auto" }}
                />
              </button>
            );
          })}
        </div>
      </div>

      {/* Displaying raw data for now */}
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
                ? filteredByQueue // queue-only selection (fast path you already maintain)
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
                Showing: <span className="font-semibold">{roleKeyToName[active]}</span> •{" "}
                {shownCount} game{shownCount !== 1 ? "s" : ""} — {stats.wins}W-{stats.losses}L • KDA{" "}
                {kdaText}
              </div>

              {shownCount === 0 ? (
                <div className="text-gray-400 italic">No games for this selection.</div>
              ) : (
                <div>
                  {/* NEW: win/loss donut for current role/queue tab */}
                  <div className="mb-3 flex justify-center">
                    <WinLossDonut wins={stats.wins} losses={stats.losses} />
                  </div>

                  <ul className="space-y-1">
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
                  </ul>
                </div>
              )}
            </>
          );
        })()}
      </div>
    </div>
  );
}
