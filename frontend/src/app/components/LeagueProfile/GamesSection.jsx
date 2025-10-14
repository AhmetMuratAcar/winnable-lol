"use client";
import { queueIdToName } from "@/lib/utils/stringUtils";
import { useState } from "react";

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

  return (
    <div
      className={`
        flex w-full rounded ${bg} overflow-hidden
      `}
    >
      <div id="previewData" className="flex-box w-full p-3">
        <p>{queueIdToName(matchData.queueId)}</p>
        <p>
          {matchData.matchId} - {gameResult}
        </p>
      </div>
      <div
        id="dataButtonContainer"
        className={`
          ${button} w-1/20
          transition-colors duration-300
        `}
      >
        <button onClick={handleFilterOpening} className="w-full h-full hover:cursor-pointer">
          {!isOpen ? "V" : "A"}
        </button>
      </div>

      <div id="completeData">{isOpen && <p>Opened</p>}</div>
    </div>
  );
}
