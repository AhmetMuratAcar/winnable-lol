"use client";
import { useState } from "react";
import { usePathname } from "next/navigation";
import Image from "next/image";
import { IMG_PATH, CURR_PATCH } from "@/lib/constants";

function PlayerRow({ item, region, baseProfileIconPath }) {
  return (
    <div className="border-b border-(--contrast-border)">
      <a
        href={`/lol/summoners/${region}/${item.gameName}-${item.tagLine}`}
        target="_blank"
        rel="noopener noreferrer"
        className="grid grid-cols-[50%_15%_20%_15%] items-center gap-x-0 py-1 pl-3 text-left hover:bg-(--background)"
      >
        <div className="flex items-center gap-2 min-w-0">
          <Image
            src={`${baseProfileIconPath}/${item.profileIconID}.png`}
            alt="summoner profile icon"
            width={20}
            height={20}
            className="rounded-lg"
          />
          <span className="truncate">
            {item.gameName}#{item.tagLine}
          </span>
        </div>
        <div className="text-center whitespace-nowrap">{item.gamesPlayed}</div>
        <div className="text-center whitespace-nowrap">
          {item.wins} - {item.losses}
        </div>
        <div className="text-center whitespace-nowrap">
          {Math.round((item.wins / item.gamesPlayed) * 100)}%
        </div>
      </a>
    </div>
  );
}

export default function PlayedWith({ playedWith, playedAgainst, gameCount }) {
  const [activeTab, setActiveTab] = useState("with");
  const region = usePathname().split("/")[3];
  const baseProfileIconPath = `${IMG_PATH}/${CURR_PATCH}/img/profileicon`;
  const currData = activeTab === "width" ? playedWith : playedAgainst;

  return (
    <div className="pt-4 bg-(--contrast) rounded border-(--contrast-border) border-t border-l border-r mb-5">
      <div className="justify-items-center text-sm">
        <p>
          Recently Played {activeTab === "with" ? "With " : "Against "}
          (Last {gameCount} Game{gameCount > 1 ? "s" : ""})
        </p>
      </div>

      <div className="flex justify-center space-x-2 my-4 px-2">
        <button
          onClick={() => setActiveTab("with")}
          className={`border-1 border-(--contrast-border) rounded-md text-white hover:cursor-pointer w-1/2
            ${activeTab === "with" ? "bg-(--pastel-red)" : "bg-(--contrast) hover:bg-(--background)"}`}
        >
          With
        </button>
        <button
          onClick={() => setActiveTab("against")}
          className={`border-1 border-(--contrast-border) rounded-md text-white hover:cursor-pointer w-1/2
            ${activeTab === "against" ? "bg-(--pastel-red)" : "bg-(--contrast) hover:bg-(--background)"}`}
        >
          Against
        </button>
      </div>

      <div>
        <div className="py-1 text-gray-400 grid grid-cols-[50%_15%_20%_15%] text-sm border-t border-b border-(--contrast-border) text-left [&_div:nth-child(n+2)]:text-center pl-3">
          <div>Summoner</div>
          <div>Played</div>
          <div>W-L</div>
          <div>WR</div>
        </div>
        {activeTab === "with" && (
          <div className="text-sm">
            {playedWith.map((item, index) => (
              <PlayerRow
                key={index}
                item={item}
                region={region}
                baseProfileIconPath={baseProfileIconPath}
              />
            ))}
          </div>
        )}

        {activeTab === "against" && (
          <div className="text-sm">
            {playedAgainst.map((item, index) => (
              <PlayerRow
                key={index}
                item={item}
                region={region}
                baseProfileIconPath={baseProfileIconPath}
              />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
