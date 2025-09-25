import Header from "@/app/components/Header";
import Footer from "@/app/components/Footer";
import ProfileGate from "./ProfileGate";

export default async function SummonerLayout({ children, params }) {
  const resolvedParams = await params;

  return (
    <main className="min-h-svh flex flex-col bg-(--background)">
      <Header />

      <section className="flex-grow flex flex-col items-center gap-6">
        <ProfileGate params={resolvedParams}>{children}</ProfileGate>
      </section>

      <Footer />
    </main>
  );
}
