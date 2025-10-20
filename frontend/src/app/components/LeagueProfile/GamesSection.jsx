"use client";
import { matchRegionToServer, queueIdToName } from "@/lib/utils/stringUtils";
import { totalToReadable } from "@/lib/utils/timeUtils";
import { useState } from "react";
import { CURR_PATCH, IMG_PATH } from "@/lib/constants";
import Image from "next/image";
import { champIdToName } from "@/lib/utils/champs";
import { summonerIdToImagePath } from "@/lib/utils/summonerSpells";
import { primaryIdToImagePath, secondaryIdToTreeImagePath } from "@/lib/utils/runes";

export default function GamesSection({ matchPreviews = [] }) {
  if (!matchPreviews || matchPreviews.length === 0) {
    return (
      <div className="w-full bg-(--contrast) rounded border-(--contrast-border) border-1 p-4">
        <p className="text-gray-400 italic">No recent games to display</p>
      </div>
    );
  }

  return (
    <div className="w-full space-y-3 mb-5">
      {matchPreviews.map((matchData, index) => (
        <LeagueMatchContainer key={index} matchData={matchData} />
      ))}

      <button
        onClick={() => console.log("Loading more games...")}
        className="
          w-full bg-(--contrast) rounded 
          border-1 border-(--contrast-border) 
          hover:cursor-pointer hover:opacity-60
          py-2
          font-semibold
        "
      >
        Load More Games
      </button>
    </div>
  );
}

// TODO: Create custom components for the different modes (eg. Arena) and map
// to components based on queueID
function LeagueMatchContainer({ matchData, open = false }) {
  const didWin =
    matchData.userPreview.team === matchData.winningTeam && !matchData.gameEndedInEarlySurrender;

  let gameResult;
  if (matchData.gameEndedInEarlySurrender) {
    gameResult = "Remake";
  } else if (didWin) {
    gameResult = "Victory";
  } else {
    gameResult = "Defeat";
  }
  const matchRegion = matchData.matchId.split("_")[0];
  const region = matchRegionToServer(matchRegion);
  const userChamp = champIdToName(matchData.userPreview.championId);

  const colorMap = {
    Default: {
      bg: "bg-(--league-remake)",
      border: "border-(--league-remake)",
      button: "bg-(--contrast) hover:bg-(--league-remake)",
    },
    Victory: {
      bg: "bg-(--league-win)",
      border: "border-(--league-win)",
      button: "bg-blue-500 hover:bg-(--league-win)",
    },
    Defeat: {
      bg: "bg-(--league-loss)",
      border: "border-(--league-loss)",
      button: "bg-(--pastel-red) hover:bg-(--league-loss)",
    },
  };
  const { bg, button } = colorMap[gameResult] || colorMap.Default;

  const [isOpen, setIsOpen] = useState(open);
  const handleFilterOpening = () => {
    setIsOpen((prev) => !prev);
  };

  const mid = Math.ceil(matchData.participants.length / 2);
  const columns = [matchData.participants.slice(0, mid), matchData.participants.slice(mid)];

  return (
    <div
      className={`
        flex w-full rounded ${bg} overflow-hidden
      `}
    >
      <div
        id="previewData"
        className="text-xs grid w-full p-3 gap-6 items-center grid-cols-[100px_1fr_max-content]"
      >
        <div id="matchDetails" className="flex flex-col justify-center">
          <p>{queueIdToName(matchData.queueId)}</p>
          <p>X Days Ago</p>
          <p>
            {/* {matchData.matchId} -  */}
            {gameResult}
          </p>
          <p>{totalToReadable(matchData.gameDuration)}</p>
        </div>

        <div id="matchStats" className="flex items-center justify-center">
          <div className="flex flex-row items-center gap-1">
            <div className="relative inline-block">
              <Image
                src={`${IMG_PATH}/${CURR_PATCH}/img/champion/${userChamp}.png`}
                width={48}
                height={48}
                className="rounded"
                alt={`${userChamp} image`}
              />
              <span className="absolute bg-(--contrast) bottom-0 left-0 text-center">
                {matchData.userPreview.champLevel}
              </span>
            </div>
            <ul id="Summoners">
              <li>
                <Image
                  src={summonerIdToImagePath(matchData.userPreview.summoner1Id)}
                  width={22}
                  height={22}
                  className="mb-0.5"
                  alt="summoner spell image"
                />
              </li>
              <li>
                <Image
                  src={summonerIdToImagePath(matchData.userPreview.summoner2Id)}
                  width={22}
                  height={22}
                  alt="summoner spell image"
                />
              </li>
            </ul>
            <ul id="Runes">
              <li>
                <Image
                  src={primaryIdToImagePath(matchData.userPreview.primaryRune)}
                  width={22}
                  height={22}
                  className="mb-0.5"
                  alt="rune image"
                />
              </li>
              <li>
                <Image
                  src={secondaryIdToTreeImagePath(matchData.userPreview.secondaryRune)}
                  width={22}
                  height={22}
                  alt="rune image"
                />
              </li>
            </ul>
          </div>
        </div>

        <div id="matchParticipants" className="hidden md:flex gap-1">
          {columns.map((col, idx) => (
            <div key={idx} className="flex flex-col gap-0.5">
              {col.map((p, i) => {
                const name = p.riotIdGameName;
                const tagLine = p.riotIdTagline;
                const championName = champIdToName(p.championId);

                return (
                  <a
                    href={`/lol/summoners/${region}/${name}-${tagLine}`}
                    className="flex items-center gap-1"
                    title={`${name} #${tagLine}`}
                    target="_blank"
                    key={i}
                  >
                    <Image
                      src={`${IMG_PATH}/${CURR_PATCH}/img/champion/${championName}.png`}
                      width={16}
                      height={16}
                      className="rounded shrink-0"
                      alt={`${championName} image`}
                    />
                    <div className="flex w-[60px] items-center gap-1">
                      <span className="truncate">{name.trim()}</span>
                    </div>
                  </a>
                );
              })}
            </div>
          ))}
        </div>
      </div>
      <div
        id="dataButtonContainer"
        className={`
          ${button} w-1/20
          transition-colors duration-300
        `}
      >
        <button onClick={handleFilterOpening} className="w-full h-full hover:cursor-pointer">
          {!isOpen ? "v" : "A"}
        </button>
      </div>

      <div id="completeData">{isOpen && <p>Opened</p>}</div>
    </div>
  );
}
