import { IMG_PATH } from "../constants";

const KEYSTONE_RUNES = {
  // Domination
  8112: { key: "Electrocute", tree: "Domination" },
  8128: { key: "DarkHarvest", tree: "Domination" },
  9923: { key: "HailOfBlades", tree: "Domination" },

  // Inspiration
  8351: { key: "GlacialAugment", tree: "Inspiration" },
  8360: { key: "UnsealedSpellbook", tree: "Inspiration" },
  8369: { key: "FirstStrike", tree: "Inspiration" },

  // Precision
  8005: { key: "PressTheAttack", tree: "Precision" },
  8008: { key: "LethalTempo", tree: "Precision" },
  8021: { key: "FleetFootwork", tree: "Precision" },
  8010: { key: "Conqueror", tree: "Precision" },

  // Resolve
  8437: { key: "GraspOfTheUndying", tree: "Resolve" },
  8439: { key: "Aftershock", tree: "Resolve" },
  8465: { key: "Guardian", tree: "Resolve" },

  // Sorcery
  8214: { key: "SummonAery", tree: "Sorcery" },
  8229: { key: "ArcaneComet", tree: "Sorcery" },
  8230: { key: "PhaseRush", tree: "Sorcery" },
};

const RUNE_TREE_MAP = {
  // DOMINATION
  9923: "Domination",
  8126: "Domination",
  8139: "Domination",
  8143: "Domination",
  8137: "Domination",
  8140: "Domination",
  8141: "Domination",
  8135: "Domination",
  8105: "Domination",
  8106: "Domination",

  // WHIMSY (Inspiration)
  8306: "Whimsy",
  8304: "Whimsy",
  8321: "Whimsy",
  8313: "Whimsy",
  8352: "Whimsy",
  8345: "Whimsy",
  8347: "Whimsy",
  8410: "Whimsy",
  8316: "Whimsy",

  // PRECISION
  9101: "Precision",
  9111: "Precision",
  8009: "Precision",
  9104: "Precision",
  9105: "Precision",
  9103: "Precision",
  8014: "Precision",
  8017: "Precision",
  8299: "Precision",

  // RESOLVE
  8446: "Resolve",
  8463: "Resolve",
  8401: "Resolve",
  8429: "Resolve",
  8444: "Resolve",
  8473: "Resolve",
  8451: "Resolve",
  8453: "Resolve",
  8242: "Resolve",

  // SORCERY
  8224: "Sorcery",
  8226: "Sorcery",
  8275: "Sorcery",
  8210: "Sorcery",
  8234: "Sorcery",
  8233: "Sorcery",
  8237: "Sorcery",
  8232: "Sorcery",
  8236: "Sorcery",
};

const TREE_EXTENSIONS = {
  Domination: "7200",
  Precision: "7201",
  Sorcery: "7202",
  Whimsy: "7203",
  Resolve: "7204",
};

export function primaryIdToImagePath(primaryRuneId) {
  const runeInfo = KEYSTONE_RUNES[primaryRuneId] || "unknown";

  if (runeInfo === "unknown") {
    return `${IMG_PATH}/img/perk-images/Styles/RunesIcon.png`;
  }

  let runePath;
  if (runeInfo.key === "LethalTempo") {
    runePath = `${IMG_PATH}/img/perk-images/Styles/${runeInfo.tree}/${runeInfo.key}/${runeInfo.key}Temp.png`;
  } else if (runeInfo.key === "Aftershock") {
    runePath = `${IMG_PATH}/img/perk-images/Styles/${runeInfo.tree}/Veteran${runeInfo.key}/Veteran${runeInfo.key}.png`;
  } else {
    runePath = `${IMG_PATH}/img/perk-images/Styles/${runeInfo.tree}/${runeInfo.key}/${runeInfo.key}.png`;
  }
  return runePath;
}

export function secondaryIdToTreeImagePath(secondaryRuneId) {
  const runeTree = RUNE_TREE_MAP[secondaryRuneId] || "unknown";

  if (runeTree === "unknown") {
    return `${IMG_PATH}/img/perk-images/Styles/RunesIcon.png`;
  }

  const extension = TREE_EXTENSIONS[runeTree];
  const runePath = `${IMG_PATH}/img/perk-images/Styles/${extension}_${runeTree}.png`;
  return runePath;
}
