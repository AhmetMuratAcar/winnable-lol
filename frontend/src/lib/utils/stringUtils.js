const REGION_TAGLINE_MAP = {
  "North America": "NA1",
  "Europe West": "EUW",
  "Europe Nordic & East": "EUNE",
  "Middle East": "ME1",
  Oceania: "OC",
  LAS: "LAS",
  LAN: "LAN",
  "Southeast Asia": "SG2",
  Korea: "KR1",
  Japan: "JP1",
  Brazil: "BR1",
  Russia: "RU1",
  Türkiye: "TR1",
  Taiwan: "TW2",
  Vietnam: "VN2",
};

const SERVER_CODE_MAP = {
  "North America": "NA1",
  Brazil: "BR1",
  "Europe West": "EUW1",
  "Europe Nordic & East": "EUN1",
  Japan: "JP1",
  Korea: "KR",
  Oceania: "OC1",
  LAS: "LA1",
  LAN: "LA2",
  "Middle East": "ME1",
  "Southeast Asia": "SG2",
  Russia: "RU",
  Türkiye: "TR1",
  Taiwan: "TW2",
  Vietnam: "VN2",
};

const REGION_SERVER_TO_NAME = {
  NA1: "North America",
  EUW: "Europe West",
  EUNE: "Europe Nordic & East",
  ME1: "Middle East",
  OC: "Oceania",
  LAS: "LAS",
  LAN: "LAN",
  SG2: "Southeast Asia",
  KR1: "Korea",
  JP1: "Japan",
  BR1: "Brazil",
  RU1: "Russia",
  TR1: "Türkiye",
  TW2: "Taiwan",
  VN2: "Vietnam",
};

const TAGLINE_TO_SERVER = {
  NA1: "NA1",
  EUW: "EUW1",
  EUNE: "EUN1",
  ME1: "ME1",
  OC: "OC1",
  LAS: "LA1",
  LAN: "LA2",
  SG2: "SG2",
  KR1: "KR",
  JP1: "JP1",
  BR1: "BR1",
  RU1: "RU",
  TR1: "TR1",
  TW2: "TW2",
  VN2: "VN2",
};

const REGION_SERVER_TO_ABBRV = {
  BR1: "BR",
  EUN1: "EUNE",
  EUW1: "EUW",
  JP1: "JP",
  KR: "KR",
  LA1: "LAN",
  LA2: "LAS",
  ME1: "ME",
  NA1: "NA",
  OC1: "OC",
  RU: "RU",
  SG2: "SEA",
  TR1: "TR",
  TW2: "TW",
  VN2: "VN",
};

export function regionServerToAbbreviation(regionTag) {
  return REGION_SERVER_TO_ABBRV[regionTag];
}

export function regionTagToServerName(regionTag) {
  return REGION_SERVER_TO_NAME[regionTag];
}

export function regionTagToServerCode(regionTag) {
  return TAGLINE_TO_SERVER[regionTag];
}

export function idValidation({ region, riotID }) {
  const regionTag = REGION_TAGLINE_MAP[region] || null;
  const result = {
    gameName: null,
    tagLine: null,
    region: regionTag,
    isValid: false,
  };

  if (typeof riotID !== "string" || !regionTag) {
    return result;
  }

  const parts = riotID.trim().split("#");
  if (parts.length > 2) {
    return result;
  }

  const namePart = parts[0].trim();
  const tagPart = parts[1] ? parts[1].trim() : regionTag;

  result.gameName = namePart;
  result.tagLine = tagPart;

  // Validation rules
  const nameRe = /^[\p{L}0-9 _\.]{3,16}$/u;
  const tagRe = /^[\p{L}0-9\s]{2,5}$/u;

  if (nameRe.test(namePart) && tagRe.test(tagPart)) {
    result.isValid = true;
  }

  console.log(result);
  return result;
}

export function rankNumeralToNum(numeral) {
  const numerals = {
    I: 1,
    II: 2,
    III: 3,
    IV: 4,
  };

  return numerals[numeral];
}
