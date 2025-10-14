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

const QUEUE_INFO = {
  0: {
    map: "Custom games",
    description: null,
    notes: null,
    name: "Custom",
  },
  72: {
    map: "Howling Abyss",
    description: "1v1 Snowdown Showdown games",
    notes: null,
    name: "ARAM",
  },
  73: {
    map: "Howling Abyss",
    description: "2v2 Snowdown Showdown games",
    notes: null,
    name: "ARAM",
  },
  75: {
    map: "Summoner's Rift",
    description: "6v6 Hexakill games",
    notes: null,
    name: "Hexakill",
  },
  76: {
    map: "Summoner's Rift",
    description: "Ultra Rapid Fire games",
    notes: null,
    name: "URF",
  },
  78: {
    map: "Howling Abyss",
    description: "One For All: Mirror Mode games",
    notes: null,
    name: "One for All",
  },
  83: {
    map: "Summoner's Rift",
    description: "Co-op vs AI Ultra Rapid Fire games",
    notes: null,
    name: "URF",
  },
  98: {
    map: "Twisted Treeline",
    description: "6v6 Hexakill games",
    notes: null,
    name: "Hexakill",
  },
  100: {
    map: "Butcher's Bridge",
    description: "5v5 ARAM games",
    notes: null,
    name: "ARAM",
  },
  310: {
    map: "Summoner's Rift",
    description: "Nemesis games",
    notes: null,
    name: "Nemesis",
  },
  313: {
    map: "Summoner's Rift",
    description: "Black Market Brawlers games",
    notes: null,
    name: "Brawlers",
  },
  317: {
    map: "Crystal Scar",
    description: "Definitely Not Dominion games",
    notes: null,
    name: "Dominion",
  },
  325: {
    map: "Summoner's Rift",
    description: "All Random games",
    notes: null,
    name: "Normal",
  },
  400: {
    map: "Summoner's Rift",
    description: "5v5 Draft Pick games",
    notes: null,
    name: "Normal",
  },
  420: {
    map: "Summoner's Rift",
    description: "5v5 Ranked Solo games",
    notes: null,
    name: "Ranked Solo",
  },
  430: {
    map: "Summoner's Rift",
    description: "5v5 Blind Pick games",
    notes: null,
    name: "Normal",
  },
  440: {
    map: "Summoner's Rift",
    description: "5v5 Ranked Flex games",
    notes: null,
    name: "Ranked Flex",
  },
  450: {
    map: "Howling Abyss",
    description: "5v5 ARAM games",
    notes: null,
    name: "ARAM",
  },
  490: {
    map: "Summoner's Rift",
    description: "Normal (Quickplay)",
    notes: null,
    name: "Quickplay",
  },
  600: {
    map: "Summoner's Rift",
    description: "Blood Hunt Assassin games",
    notes: null,
    name: "Dark Star",
  },
  610: {
    map: "Cosmic Ruins",
    description: "Dark Star: Singularity games",
    notes: null,
    name: "Dark Star",
  },
  700: {
    map: "Summoner's Rift",
    description: "Summoner's Rift Clash games",
    notes: null,
    name: "Clash",
  },
  720: {
    map: "Howling Abyss",
    description: "ARAM Clash games",
    notes: null,
    name: "ARAM Clash",
  },
  820: {
    map: "Twisted Treeline",
    description: "Co-op vs. AI Beginner Bot games",
    notes: null,
    name: "Bot",
  },
  870: {
    map: "Summoner's Rift",
    description: "Co-op vs. AI Intro Bot games",
    notes: null,
    name: "Bot",
  },
  880: {
    map: "Summoner's Rift",
    description: "Co-op vs. AI Beginner Bot games",
    notes: null,
    name: "Bot",
  },
  890: {
    map: "Summoner's Rift",
    description: "Co-op vs. AI Intermediate Bot games",
    notes: null,
    name: "Bot",
  },
  900: {
    map: "Summoner's Rift",
    description: "ARURF games",
    notes: null,
    name: "ARURF",
  },
  910: {
    map: "Crystal Scar",
    description: "Ascension games",
    notes: null,
    name: "Ascension",
  },
  920: {
    map: "Howling Abyss",
    description: "Legend of the Poro King games",
    notes: null,
    name: "Poro King",
  },
  940: {
    map: "Summoner's Rift",
    description: "Nexus Siege games",
    notes: null,
    name: "Nexus Siege",
  },
  950: {
    map: "Summoner's Rift",
    description: "Doom Bots Voting games",
    notes: null,
    name: "Doom Bots",
  },
  960: {
    map: "Summoner's Rift",
    description: "Doom Bots Standard games",
    notes: null,
    name: "Doom Bots",
  },
  980: {
    map: "Valoran City Park",
    description: "Star Guardian Invasion: Normal games",
    notes: null,
    name: "Star Guardian",
  },
  990: {
    map: "Valoran City Park",
    description: "Star Guardian Invasion: Onslaught games",
    notes: null,
    name: "Star Guardian",
  },
  1000: {
    map: "Overcharge",
    description: "PROJECT: Hunters games",
    notes: null,
    name: "Project",
  },
  1010: {
    map: "Summoner's Rift",
    description: "Snow ARURF games",
    notes: null,
    name: "ARURF",
  },
  1020: {
    map: "Summoner's Rift",
    description: "One for All games",
    notes: null,
    name: "One for All",
  },
  1030: {
    map: "Crash Site",
    description: "Odyssey Extraction: Intro games",
    notes: null,
    name: "Odyssey",
  },
  1040: {
    map: "Crash Site",
    description: "Odyssey Extraction: Cadet games",
    notes: null,
    name: "Odyssey",
  },
  1050: {
    map: "Crash Site",
    description: "Odyssey Extraction: Crewmember games",
    notes: null,
    name: "Odyssey",
  },
  1060: {
    map: "Crash Site",
    description: "Odyssey Extraction: Captain games",
    notes: null,
    name: "Odyssey",
  },
  1070: {
    map: "Crash Site",
    description: "Odyssey Extraction: Onslaught games",
    notes: null,
    name: "Odyssey",
  },
  1090: {
    map: "Convergence",
    description: "Teamfight Tactics games",
    notes: null,
    name: "TFT-Normal",
  },
  1100: {
    map: "Convergence",
    description: "Ranked Teamfight Tactics games",
    notes: null,
    name: "TFT-Ranked",
  },
  1110: {
    map: "Convergence",
    description: "Teamfight Tactics Tutorial games",
    notes: null,
    name: "TFT-Tutorial",
  },
  1111: {
    map: "Convergence",
    description: "Teamfight Tactics test games",
    notes: null,
    name: "TFT-Test",
  },
  1210: {
    map: "Convergence",
    description: "Teamfight Tactics Choncc's Treasure Mode",
    notes: null,
    name: "Chonc's Treasure",
  },
  1300: {
    map: "Nexus Blitz",
    description: "Nexus Blitz games",
    notes: null,
    name: "Nexus Blitz",
  },
  1400: {
    map: "Summoner's Rift",
    description: "Ultimate Spellbook games",
    notes: null,
    name: "Ultimate Spellbook",
  },
  1700: {
    map: "Rings of Wrath",
    description: "Arena",
    notes: null,
    name: "Arena",
  },
  1710: {
    map: "Rings of Wrath",
    description: "Arena",
    notes: "16 player lobby",
    name: "Arena",
  },
  1810: {
    map: "Swarm",
    description: "Swarm Mode Games",
    notes: "Swarm Mode 1 player",
    name: "Swarm",
  },
  1820: {
    map: "Swarm Mode Games",
    description: "Swarm",
    notes: "Swarm Mode 2 players",
    name: "Swarm",
  },
  1830: {
    map: "Swarm Mode Games",
    description: "Swarm",
    notes: "Swarm Mode 3 players",
    name: "Swarm",
  },
  1840: {
    map: "Swarm Mode Games",
    description: "Swarm",
    notes: "Swarm Mode 4 players",
    name: "Swarm",
  },
  1900: {
    map: "Summoner's Rift",
    description: "Pick URF games",
    notes: null,
    name: "URF",
  },
  2000: {
    map: "Summoner's Rift",
    description: "Tutorial 1",
    notes: null,
    name: "Tutorial",
  },
  2010: {
    map: "Summoner's Rift",
    description: "Tutorial 2",
    notes: null,
    name: "Tutorial",
  },
  2020: {
    map: "Summoner's Rift",
    description: "Tutorial 3",
    notes: null,
    name: "Tutorial",
  },
};

export function queueIdToName(queueID) {
  const queueName = QUEUE_INFO[queueID].name || "unknown";
  return queueName;
}

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
