// import Image from "next/image";
import Footer from "./components/Footer";
import Navbar from "./components/Navbar";
import MainSection from "./components/MainSection";

export default function Home() {
  return (
    <main className="min-h-svh flex flex-col">
      <Navbar />
      <MainSection />
      <Footer />
    </main>
  );
}

