"use client";
import LeagueMatchContainer from "../LeagueMatch/DefaultMatchContainer";

export default function GamesSection({ matchPreviews = [] }) {
  if (!matchPreviews || matchPreviews.length === 0) {
    return (
      <div className="w-full bg-(--contrast) rounded border-(--contrast-border) border-1 p-4">
        <p className="text-gray-400 italic">No recent games to display</p>
      </div>
    );
  }

  // TODO: Create custom components for the different modes (eg. Arena) and map
  // to components based on queueID
  return (
    <div className="w-full space-y-2 mb-5">
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
