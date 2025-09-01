import Header from "@/app/components/Header";
import Footer from "@/app/components/Footer";

export default async function SlowPage() {
  await new Promise((resolve) => setTimeout(resolve, 8000));

  return (
    <main className="min-h-svh flex flex-col">
      <Header />
      <div className="flex flex-col flex-grow items-center justify-center px-4 py-6 text-center">
        <h1 className="text-3xl font-bold">Slow Page Loaded</h1>
        <p className="text-gray-500 mt-4">
          This page delayed on purpose so you could see the global loading
          spinner.
        </p>
      </div>
      <Footer />
    </main>
  );
}
