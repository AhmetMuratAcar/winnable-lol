"use client";
import Image from "next/image";
import { useState } from "react";

export default function RecentGames({ recentGames, totalGameCount }) {
  // totalGameCount = 0;
  // recentGames.matchSummaries = [];
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
    { key: "all", src: "/images/roles/all.svg", rounded: "rounded-l" },
    { key: "top", src: "/images/roles/top.svg" },
    { key: "jungle", src: "/images/roles/jungle.svg" },
    { key: "middle", src: "/images/roles/middle.svg" },
    { key: "bottom", src: "/images/roles/bottom.svg" },
    { key: "utility", src: "/images/roles/utility.svg", rounded: "rounded-r" },
  ];

  // default
  return (
    <div className="w-full bg-(--contrast) rounded border-(--contrast-border) border- p-4">
      <h3 className="font-semibold border-b-1 border-(--contrast-border) mb-3">
        Recent Ranked Games
      </h3>

      <div className="flex">
        {roles.map(({ key, src, rounded }) => {
          const isActive = active === key;

          return (
            <button
              key={key}
              onClick={() => setActive(key)}
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
                alt={`${key} role icon`}
                style={{ width: "20px", height: "20px" }}
              />
            </button>
          );
        })}
      </div>

      <p>data:</p>
    </div>
  );
}
