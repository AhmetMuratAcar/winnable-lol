import Header from "@/app/components/Header";
import Footer from "@/app/components/Footer";

export default function SummonerNotFoundPage() {
    return (
        <main className="min-h-svh flex flex-col">
            <Header />
            <section
                id="ErrorSection"
                className="flex flex-col flex-grow items-center justify-start px-4 py-6"
            >
                <h1 className="text-3xl font-bold">Summoner Not Found</h1>
            </section>
            <Footer />
        </main>
    );
}
