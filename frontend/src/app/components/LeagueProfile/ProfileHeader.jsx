"use client";

import { RefreshButton } from ".";
import { regionServerToAbbreviation } from "@/lib/utils/stringUtils";
import Image from "next/image";
import { IMG_PATH, CURR_PATCH } from "@/lib/constants";
import { calcLastUpdated } from "@/lib/utils/timeUtils";

export default function ProfileHeader({ headerData }) {
  if (!headerData) {
    return <p>No profileHeader data available</p>;
  }

  const { gameName, tagLine, region, summonerLevel, profileIconId, lastUpdated } = headerData;
  const profileIconSrc = `${IMG_PATH}/${CURR_PATCH}/img/profileicon/${profileIconId}.png`;
  return (
    <section
      id="profileHeader"
      className="bg-(--contrast) mt-10 p-10 border-1 border-(--contrast-border) rounded w-full max-w-2xl mx-auto"
    >
      <div className="flex gap-6 items-start">
        <div id="iconAndLevel" className="relative inline-block self-start">
          <Image
            src={profileIconSrc}
            alt="summoner profile icon"
            width={100}
            height={100}
            className="rounded-lg border border-(--contrast-border)"
          />
          <p
            className="
              absolute left-1/2 -translate-x-1/2
              bottom-0 translate-y-1/2
              min-w-[3ch] text-center
              bg-gray-800 text-white text-sm font-semibold
              px-2 py-0.5 rounded-md shadow-md border border-(--contrast-border)
            "
          >
            {summonerLevel}
          </p>
        </div>

        <div className="flex flex-col">
          <p className="text-2xl font-bold">
            {gameName}
            <span className="text-gray-400"> #{tagLine}</span>
          </p>
          <p className="text-gray-300">{regionServerToAbbreviation(region)}</p>

          <div id="update" className="mt-2">
            <RefreshButton />
            <div>
              <span className="text-gray-400 text-xs">
                {" "}
                Last updated: {calcLastUpdated(lastUpdated)}
              </span>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
