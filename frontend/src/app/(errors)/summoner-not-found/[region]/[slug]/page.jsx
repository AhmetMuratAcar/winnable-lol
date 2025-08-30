import Header from "@/app/components/Header";
import Footer from "@/app/components/Footer";
import GoHomeButton from "@/app/components/GoHomeButton";
import { regionTagToServerName } from "@/app/utils/idValidation";

export default async function SummonerNotFoundPage({ params }) {
  const { region, slug } = await params;
  const raw = decodeURIComponent(slug);

  let displayName = raw;
  if (raw.includes("-")) {
    const [gameName = "", tagLine = ""] = raw.split("-");
    if (gameName && tagLine) {
      displayName = `${gameName}#${tagLine}`;
    }
  }

  return (
    <main className="min-h-svh flex flex-col">
      <Header />
      <section
        id="ErrorSection"
        className="flex flex-col flex-grow items-center justify-center px-4 py-6 text-center"
      >
        <h1 className="text-3xl font-bold mb-4">Summoner Not Found</h1>
        <p className="text-lg text-gray-500 mb-6">
          <span className="font-semibold">{displayName}</span> does not exist in{" "}
          <span className="font-semibold">{regionTagToServerName(region)}</span>.
        </p>
        <GoHomeButton />
      </section>
      <Footer />
    </main>
  );
}