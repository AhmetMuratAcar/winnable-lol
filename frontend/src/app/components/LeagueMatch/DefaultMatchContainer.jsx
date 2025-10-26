"use client";
import { matchRegionToServer, queueIdToName } from "@/lib/utils/stringUtils";
import { calcWhenPlayed, totalToReadable } from "@/lib/utils/timeUtils";
import { useState } from "react";
import { CURR_PATCH, IMG_PATH } from "@/lib/constants";
import Image from "next/image";
import { champIdToName } from "@/lib/utils/champs";
import { summonerIdToImagePath } from "@/lib/utils/summonerSpells";
import { primaryIdToImagePath, secondaryIdToTreeImagePath } from "@/lib/utils/runes";

export default function LeagueMatchContainer({ matchData }) {
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
  const { kills, deaths, assists } = matchData.userPreview;
  const KDA = deaths === 0 ? kills + assists : (kills + assists) / deaths;
  const CSPM = matchData.userPreview.totalMinionsKilled / (matchData.gameDuration / 60);
  const gameEndTimestamp = matchData.gameStartTimeStamp + matchData.gameDuration * 1000;

  const colorMap = {
    Default: {
      bg: "bg-(--league-remake)",
      contrast: "bg-(--contrast)",
      button: "bg-(--contrast) hover:bg-(--league-remake)",
    },
    Victory: {
      bg: "bg-(--league-win)",
      contrast: "bg-blue-500",
      button: "bg-blue-500 hover:bg-(--league-win)",
    },
    Defeat: {
      bg: "bg-(--league-loss)",
      contrast: "bg-red-500",
      button: "bg-(--pastel-red) hover:bg-(--league-loss)",
    },
  };
  const { bg, contrast, button } = colorMap[gameResult] || colorMap.Default;

  const [isOpen, setIsOpen] = useState(false);
  const [showDiv, setShowDiv] = useState(false);
  const [details, setDetails] = useState(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const handleToggle = () => {
    // first click → create and show the div
    if (!showDiv) {
      setShowDiv(true);
      setIsOpen(true);
      return;
    }
    // next clicks → toggle visibility
    setIsOpen((prev) => !prev);
  };

  const mid = Math.ceil(matchData.participants.length / 2);
  const columns = [matchData.participants.slice(0, mid), matchData.participants.slice(mid)];

  return (
    <div className="w-full">
      <div
        className={`
        flex w-full rounded ${bg} overflow-hidden
      `}
      >
        <div
          id="previewData"
          className="text-xs grid w-full py-1 px-3 gap-6 items-center grid-cols-[100px_1fr_max-content]"
        >
          <div id="matchDetails" className="flex flex-col justify-center">
            <p className="font-bold text-sm">{queueIdToName(matchData.queueId)}</p>
            <p className="text-gray-400">{calcWhenPlayed(gameEndTimestamp)}</p>
            <div className={`block ${contrast} h-px w-1/2 my-1`}></div>
            <p>{gameResult}</p>
            <p className="text-gray-400">{totalToReadable(matchData.gameDuration)}</p>
          </div>

          <div id="previewMain" className="flex items-center justify-center space-x-5">
            <div id="previewImages" className="flex flex-row items-center gap-1">
              <div className="relative inline-block">
                <Image
                  src={`${IMG_PATH}/${CURR_PATCH}/img/champion/${userChamp}.png`}
                  width={52}
                  height={52}
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
                    width={25}
                    height={25}
                    className="mb-0.5"
                    alt="summoner spell image"
                  />
                </li>
                <li>
                  <Image
                    src={summonerIdToImagePath(matchData.userPreview.summoner2Id)}
                    width={25}
                    height={25}
                    alt="summoner spell image"
                  />
                </li>
              </ul>
              <ul id="Runes">
                <li>
                  <Image
                    src={primaryIdToImagePath(matchData.userPreview.primaryRune)}
                    width={25}
                    height={25}
                    className="mb-0.5"
                    alt="rune image"
                  />
                </li>
                <li>
                  <Image
                    src={secondaryIdToTreeImagePath(matchData.userPreview.secondaryRune)}
                    width={25}
                    height={25}
                    alt="rune image"
                  />
                </li>
              </ul>
            </div>

            <div id="previewStats" className="flex flex-col items-center justify-center">
              <div className="text-base">
                <span>{matchData.userPreview.kills} / </span>
                <span className="text-red-500">{matchData.userPreview.deaths} </span>
                <span>/ {matchData.userPreview.assists}</span>
              </div>
              <p className="text-gray-400">{KDA.toFixed(2)} KDA</p>
              <p className="text-gray-400">
                {matchData.userPreview.totalMinionsKilled} CS ({CSPM.toFixed(1)})
              </p>
            </div>

            <div id="previewItems" className="flex items-center gap-2">
              {/* 2 rows of 3 */}
              <div className="grid grid-cols-3 grid-rows-2 gap-1">
                {matchData.userPreview.items.slice(0, 6).map((itemId, index) => (
                  <div className={`h-6 w-6 ${contrast} rounded`} key={`${itemId} - ${index}`}>
                    {itemId !== 0 && (
                      <Image
                        height={25}
                        width={25}
                        src={`${IMG_PATH}/${CURR_PATCH}/img/item/${itemId}.png`}
                        alt={`Item ${itemId} image`}
                        className="rounded"
                      />
                    )}
                  </div>
                ))}
              </div>

              {/* Ward */}
              <div className="flex items-center">
                <div className="w-10 h-10flex items-center justify-center rounded">
                  <Image
                    height={25}
                    width={25}
                    src={`${IMG_PATH}/${CURR_PATCH}/img/item/${matchData.userPreview.items[6]}.png`}
                    alt="item image"
                    className="rounded"
                  />
                </div>
              </div>
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
          <button
            onClick={handleToggle}
            className="w-full h-full pb-1 flex items-end justify-center cursor-pointer"
            aria-expanded={isOpen}
            aria-controls={`match-extra-${matchData.matchId}`}
          >
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="24"
              height="24"
              fill="currentColor"
              viewBox="0 0 24 24"
              className={`transition-transform duration-300 ${isOpen ? "rotate-180" : ""}`}
            >
              <path d="M7 10l5 5 5-5H7z" />
            </svg>
          </button>
        </div>
      </div>

      {showDiv && (
        <div id={`match-extra-${matchData.matchId}`} className="w-full rounded mt-2">
          {isOpen && (
            <div className="w-full bg-(--contrast)  rounded p-3">
              <p className="text-sm">Opened up: {matchData.matchId}</p>
            </div>
          )}
        </div>
      )}
    </div>
  );
}
