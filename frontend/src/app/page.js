import Footer from "./components/Footer";
import Header from "./components/Header";
import MainSection from "./components/MainSection";

export default function Home() {
  return (
    <main className="min-h-svh flex flex-col">
      <Header />
      <MainSection />
      <Footer />
    </main>
  );
}