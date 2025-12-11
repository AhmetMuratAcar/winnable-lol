import { CURR_PATCH, IMG_PATH } from "../constants";

const SUMMONER_SPELLS = {
  1: "SummonerBoost",
  3: "SummonerExhaust",
  4: "SummonerFlash",
  6: "SummonerHaste",
  7: "SummonerHeal",
  11: "SummonerSmite",
  12: "SummonerTeleport",
  13: "SummonerMana",
  14: "SummonerDot",
  21: "SummonerBarrier",
  30: "SummonerPoroRecall",
  31: "SummonerPoroThrow",
  32: "SummonerSnowball",
  39: "SummonerSnowURFSnowball_Mark",
  54: "Summoner_UltBookPlaceholder",
  55: "Summoner_UltBookSmitePlaceholder",
  2201: "SummonerCherryHold",
  2202: "SummonerCherryFlash",
};

export function summonerIdToImagePath(summonerSpellId) {
  const summonerSpell = SUMMONER_SPELLS[summonerSpellId] || "unknown";

  if (summonerSpell === "unknown") {
    return `${IMG_PATH}/${CURR_PATCH}/img/spell/Summoner_UltBookPlaceholder.png`;
  }

  return `${IMG_PATH}/${CURR_PATCH}/img/spell/${summonerSpell}.png`;
}
