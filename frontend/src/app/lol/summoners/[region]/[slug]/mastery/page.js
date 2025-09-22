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
  return (
    <div>
      <p>mastery page</p>
    </div>
  );
}
