import { RefreshButton } from ".";
import { regionServerToAbbreviation } from "@/app/utils/idValidation";

export default function ProfileHeader({ headerData }) {
  if (!headerData) {
    return <p>No profileHeader data available</p>;
  }

  const { gameName, tagLine, region, summonerLevel, profileIconId } =
    headerData;
  return (
    <section id="profileHeader">
      <div>
        <div>
          <p>
            {profileIconId}: {summonerLevel}
          </p>
        </div>

        <div>
          <p>
            {gameName}#{tagLine}
          </p>
          <p>{regionServerToAbbreviation(region)}</p>
        </div>

        <RefreshButton />
      </div>
    </section>
  );
}
