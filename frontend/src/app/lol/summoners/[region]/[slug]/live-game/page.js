export async function generateMetadata({ params }) {
  const { slug } = await params;
  if (typeof slug !== "string" || !slug.includes("-")) {
    return { title: "Invalid Summoner | Winnable" };
  }
  const [rawGameName, rawTagLine] = slug.split("-");
  const gameName = decodeURIComponent(rawGameName || "");
  const tagLine = decodeURIComponent(rawTagLine || "");
  return {
    title: `${gameName} ${tagLine ? `#${tagLine}` : ""} - Live Game | Winnable`,
    description: `Live game stats for ${gameName}${tagLine ? ` #${tagLine}` : ""}.`,
  };
}

export default async function LivePage({ params }) {
  return (
    <div>
      <p>live game page</p>
    </div>
  );
}
