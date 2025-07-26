import Header from "@/app/components/Header";
import Footer from "@/app/components/Footer";

export default function NotFoundPage() {
    return (
        <main className="min-h-svh flex flex-col">
            <Header />
            <section
                id="ErrorSection"
                className="flex flex-col flex-grow items-center justify-start px-4 py-6"
            >
                <h1 className="text-3xl font-bold">404</h1>
            </section>
            <Footer />
        </main>
    );
}
