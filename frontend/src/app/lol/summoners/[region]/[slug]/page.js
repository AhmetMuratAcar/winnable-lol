import ClientOverviewBridge from "./ClientOverviewBridge";

export async function generateMetadata({ params }) {
  const { slug } = await params;
  if (typeof slug !== "string" || !slug.includes("-")) {
    return { title: "Invalid Summoner | Winnable" };
  }
  const [rawGameName, rawTagLine] = slug.split("-");
  const gameName = decodeURIComponent(rawGameName || "");
  const tagLine = decodeURIComponent(rawTagLine || "");
  return {
    title: `${gameName} ${tagLine ? `#${tagLine}` : ""} - Summoner Stats | Winnable`,
    description: `Summoner stats for ${gameName}${tagLine ? ` #${tagLine}` : ""}.`,
  };
}

export default function OverviewPage() {
  return <ClientOverviewBridge />;
}
