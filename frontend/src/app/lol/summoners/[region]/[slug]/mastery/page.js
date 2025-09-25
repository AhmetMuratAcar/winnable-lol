import { getMastery } from "@/lib/server/lol";

export async function generateMetadata({ params }) {
  const { slug } = await params;
  if (typeof slug !== "string" || !slug.includes("-")) {
    return { title: "Invalid Summoner | Winnable" };
  }
  const [rawGameName, rawTagLine] = slug.split("-");
  const gameName = decodeURIComponent(rawGameName || "");
  const tagLine = decodeURIComponent(rawTagLine || "");
  return {
    title: `${gameName} ${tagLine ? `#${tagLine}` : ""} - Mastery Stats | Winnable`,
    description: `Mastery stats for ${gameName}${tagLine ? ` #${tagLine}` : ""}.`,
  };
}

export default async function MasteryPage({ params }) {
  const resolvedParams = await params;
  const masteryData = await getMastery(resolvedParams);

  return (
    <div>
      <p>mastery page</p>
      {masteryData && (
        <div>
          status: {masteryData.status}
          message: {masteryData.message}
        </div>
      )}
    </div>
  );
}
