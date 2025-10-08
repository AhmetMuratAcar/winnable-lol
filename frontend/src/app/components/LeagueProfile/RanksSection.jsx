import { rankNumeralToNum } from "@/lib/utils/stringUtils";
import Image from "next/image";

export default function RanksSection({ rankData = [] }) {
  const soloq = rankData.find((item) => item.queueType === "RANKED_SOLO_5x5");
  const flex = rankData.find((item) => item.queueType === "RANKED_FLEX_SR");

  function RankCard(rank, label) {
    const winRate = rank ? Math.round((rank.wins / (rank.wins + rank.losses)) * 100) : null;

    return (
      <div className="p-4 rounded border-1 border-(--contrast-border) bg-(--contrast) text-white">
        <div className="flex items-center justify-between ">
          <h3 className="font-semibold">{label}</h3>
          {!rank && <span className="italic text-gray-400">Unranked</span>}
        </div>

        {rank && (
          <div>
            <div className="flex items-center border-t-1 border-(--contrast-border)">
              <div className="flex items-center gap-4">
                <Image
                  src={`/images/Ranked-Emblems/Rank=${rank.tier}.png`}
                  alt={`${label} queue emblem`}
                  width={62}
                  height={62}
                />
                <div className="leading-tight">
                  <p className="text-l font-extrabold">
                    {rank.tier.charAt(0) + rank.tier.slice(1).toLowerCase()}{" "}
                    {rank.tier !== "CHALLENGER" && rank.tier !== "GRANDMASTER"
                      ? rankNumeralToNum(rank.rank)
                      : ""}
                  </p>
                  <p className="text-gray-400 text-sm">{rank.leaguePoints} LP</p>
                </div>
              </div>

              <div className="ml-auto text-right leading-tight">
                <p className="text-sm">
                  {rank.wins}W {rank.losses}L
                </p>
                <p className="text-gray-400 text-sm">Win Rate: {winRate}%</p>
              </div>
            </div>

            <div className="mt-2 h-2 w-full flex rounded overflow-hidden">
              <div className="bg-green-500" style={{ width: `${winRate}%` }} />
              <div className="bg-(--pastel-red)" style={{ width: `${100 - winRate}%` }} />
            </div>
          </div>
        )}
      </div>
    );
  }

  return (
    <section className="space-y-3">
      {RankCard(soloq, "Solo/Duo")}
      {RankCard(flex, "Flex")}
    </section>
  );
}
