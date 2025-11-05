import { CURR_PATCH, IMG_PATH } from "@/lib/constants";
import { champIdToName } from "@/lib/utils/champs";
import { primaryIdToImagePath, secondaryIdToTreeImagePath } from "@/lib/utils/runes";
import { summonerIdToImagePath } from "@/lib/utils/summonerSpells";
import Image from "next/image";

export default function ExpandedMatchContainer({ completeData, userId, region }) {
  const colorMap = {
    Victory: {
      bg: "bg-(--league-win)",
      text: "text-blue-500",
      contrast: "bg-blue-500",
      user: "bg-(--user-win-contrast)",
    },
    Defeat: {
      bg: "bg-(--league-loss)",
      text: "text-(--pastel-red)",
      contrast: "bg-(--pastel-red)",
      user: "bg-(--user-loss-contrast)",
    },
  };

  const didBlueWin = completeData.winningTeam === 0;
  const {
    bg: blueTeamBg,
    text: blueTeamText,
    contrast: blueTeamContrast,
    user: blueUserContrast,
  } = didBlueWin ? colorMap["Victory"] : colorMap["Defeat"];
  const {
    bg: redTeamBg,
    text: redTeamText,
    contrast: redTeamContrast,
    user: redUserContrast,
  } = didBlueWin ? colorMap["Defeat"] : colorMap["Victory"];

  // TODO: add most damage and most damage taken to metadata
  const matchMetadata = {
    gameDuration: completeData.gameDuration,
    region: region,
    userId: userId,
    mostDamage: completeData.mostDamageDone,
    mostDamageTaken: completeData.mostDamageTaken,
    contrast: "",
    userContrast: "",
  };

  const blueTeam = completeData.participants.slice(0, 5);
  const redTeam = completeData.participants.slice(5);

  return (
    <div className="text-sm rounded overflow-hidden">
      <table className="w-full">
        <colgroup>
          <col style={{ width: "31%" }} />
          <col style={{ width: "11%" }} />
          <col style={{ width: "10%" }} />
          <col style={{ width: "18%" }} />
          <col style={{ width: "9%" }} />
          <col style={{ width: "21%" }} />
        </colgroup>
        <thead>
          <tr className="text-xs text-gray-400">
            <th className="text-left flex gap-x-1 py-2 pl-2">
              <div className={`font-bold ${blueTeamText}`}>{didBlueWin ? "Victory" : "Defeat"}</div>
              <span>Blue Side</span>
            </th>
            <th>
              <div>KDA</div>
            </th>
            <th>
              <div>CS</div>
            </th>
            <th>
              <div>Damage</div>
            </th>
            <th>
              <div>Vision</div>
            </th>
            <th>
              <div>Items</div>
            </th>
          </tr>
        </thead>
        <tbody className={`${blueTeamBg} text-xs`}>
          {blueTeam.map((pData, index) => (
            <ParticipantRow
              key={index}
              participant={pData}
              metadata={{
                ...matchMetadata,
                contrast: blueTeamContrast,
                userContrast: blueUserContrast,
              }}
            />
          ))}
        </tbody>
      </table>

      <div className="w-full bg-(--contrast) border-y-2 border-(--background) text-center">
        <p className="m-8">match data will go here</p>
      </div>

      <table className="w-full">
        <colgroup>
          <col style={{ width: "31%" }} />
          <col style={{ width: "11%" }} />
          <col style={{ width: "10%" }} />
          <col style={{ width: "18%" }} />
          <col style={{ width: "9%" }} />
          <col style={{ width: "21%" }} />
        </colgroup>
        <thead>
          <tr className="text-xs text-gray-400">
            <th className="text-left flex gap-x-1 py-2 pl-2">
              <div className={`font-bold ${redTeamText}`}>{didBlueWin ? "Defeat" : "Victory"}</div>
              <span>Red Side</span>
            </th>
            <th>
              <div>KDA</div>
            </th>
            <th>
              <div>CS</div>
            </th>
            <th>
              <div>Damage</div>
            </th>
            <th>
              <div>Vision</div>
            </th>
            <th>
              <div>Items</div>
            </th>
          </tr>
        </thead>
        <tbody className={`${redTeamBg} text-xs`}>
          {redTeam.map((pData, index) => (
            <ParticipantRow
              key={index}
              participant={pData}
              metadata={{
                ...matchMetadata,
                contrast: redTeamContrast,
                userContrast: redUserContrast,
              }}
            />
          ))}
        </tbody>
      </table>
    </div>
  );
}

function ParticipantRow({ participant, metadata }) {
  const targetUser = participant.riotIdGameName + participant.riotIdTagLine === metadata.userId;

  const KDA =
    participant.deaths === 0
      ? participant.kills + participant.assists
      : (participant.kills + participant.assists) / participant.deaths;
  const CSPM = participant.totalMinionsKilled / (metadata.gameDuration / 60);
  const champName = champIdToName(participant.championId);
  const damagePct = Math.round(
    (participant.totalDamageDealtToChampions / metadata.mostDamage) * 100,
  );
  const takenPct = Math.round((participant.totalDamageTaken / metadata.mostDamageTaken) * 100);

  return (
    <tr className={`text-center ${targetUser ? metadata.userContrast : ""}`}>
      <td className="py-1 pl-2">
        <div className="flex items-center gap-1">
          <div className="relative inline-block">
            <Image
              src={`${IMG_PATH}/${CURR_PATCH}/img/champion/${champName}.png`}
              width={32}
              height={32}
              className="rounded"
              alt={`${champName} image`}
            />
            <span className="absolute bg-(--contrast) bottom-0 right-0 text-center">
              {participant.champLevel}
            </span>
          </div>

          <ul className="flex flex-col items-center">
            <li>
              <Image
                src={summonerIdToImagePath(participant.summoner1Id)}
                width={16}
                height={16}
                className="rounded mb-0.5"
                alt="summoner spell image"
              />
            </li>
            <li>
              <Image
                src={summonerIdToImagePath(participant.summoner2Id)}
                width={16}
                height={16}
                className="rounded"
                alt="summoner spell image"
              />
            </li>
          </ul>

          <ul className="flex flex-col items-center">
            <li>
              <Image
                src={primaryIdToImagePath(participant.runes.mainTree.keystone)}
                width={16}
                height={16}
                className="mb-0.5"
                alt="rune image"
              />
            </li>
            <li>
              <Image
                src={secondaryIdToTreeImagePath(participant.runes.secondaryTree.rune1)}
                width={16}
                height={16}
                alt="rune image"
              />
            </li>
          </ul>

          <div className="flex items-center min-w-0">
            <a
              href={`/lol/summoners/${metadata.region}/${participant.riotIdGameName}-${participant.riotIdTagline}`}
              title={`${participant.riotIdGameName} #${participant.riotIdTagline}`}
              target="_blank"
              className="block overflow-hidden text-ellipsis whitespace-nowrap hover:underline decoration-white"
            >
              <span className="text-white">{participant.riotIdGameName}</span>
              <span className="text-gray-400"> #{participant.riotIdTagline}</span>
            </a>
          </div>
        </div>
      </td>
      <td>
        <p>
          {participant.kills} / {participant.deaths} / {participant.assists}
        </p>
        <div className="text-gray-400">{KDA.toFixed(2)}</div>
      </td>
      <td>
        <p>{participant.totalMinionsKilled}</p>
        <div className="text-gray-400">{CSPM.toFixed(1)}/m</div>
      </td>
      <td>
        <div className="flex justify-between gap-4">
          <div id="damageDone" className="w-1/2">
            <div>{participant.totalDamageDealtToChampions.toLocaleString("en-US")}</div>
            <div className="w-full h-1.5 mt-1 flex rounded overflow-hidden">
              <div className="bg-(--pastel-red)" style={{ width: `${damagePct}%` }}></div>
              <div className="bg-(--contrast)" style={{ width: `${100 - damagePct}%` }}></div>
            </div>
          </div>

          <div id="damageTaken" className="w-1/2">
            <div>{participant.totalDamageTaken.toLocaleString("en-US")}</div>
            <div className="w-full h-1.5 mt-1 flex rounded overflow-hidden">
              <div className="bg-gray-500" style={{ width: `${damagePct}%` }}></div>
              <div className="bg-(--contrast)" style={{ width: `${100 - damagePct}%` }}></div>
            </div>
          </div>
        </div>
      </td>
      <td>
        <div>{participant.visionScore}</div>
      </td>
      <td className="pr-2">
        <div className="flex gap-1">
          {participant.items.map((itemId, index) => (
            <div
              className={`h-6 w-6 rounded flex items-center justify-center ${metadata.userContrast}`}
              key={`${itemId} - ${index}`}
            >
              {itemId !== 0 && (
                <Image
                  height={24}
                  width={24}
                  src={`${IMG_PATH}/${CURR_PATCH}/img/item/${itemId}.png`}
                  alt={`Item ${itemId} image`}
                  className="rounded"
                />
              )}
            </div>
          ))}
        </div>
      </td>
    </tr>
  );
}
