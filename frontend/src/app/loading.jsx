import Header from "@/app/components/Header";
import Footer from "@/app/components/Footer";

export default function Loading() {
  return (
    <main className="min-h-svh flex flex-col">
      <Header />
      <section
        id="ProfileSectionLoading"
        className="flex-grow px-4 py-6 flex flex-col items-center"
      >
        <div className="w-full max-w-3xl animate-pulse space-y-6">
          <section className="flex-grow grid place-items-center">
            <p className="py-4">Building...</p>
            <div className="size-10 animate-spin rounded-full border-4 border-white/30 border-t-white" />
          </section>
        </div>
      </section>
      <Footer />
    </main>
  );
}
