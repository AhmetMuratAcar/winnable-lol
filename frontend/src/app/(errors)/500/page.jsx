"use client"
import Header from "@/app/components/Header";
import Footer from "@/app/components/Footer";
import GoHomeButton from "@/app/components/GoHomeButton";
import { useSearchParams } from "next/navigation";

export default function customErrorPage() {
    const searchParams = useSearchParams()
    const status = searchParams.get('status') || '500'

    return (
        <main className="min-h-svh flex flex-col">
            <Header />
            <section
                id="ErrorSection"
                className="flex flex-col flex-grow items-center justify-start px-4 py-6"
            >
                <h1 className="text-3xl font-bold">Server Error</h1>
                <p className="mt-4 text-gray-600">Something went wrong on our end.</p>
                <p className="mt-2 text-sm text-gray-500">Error Code: {status}</p>
                <GoHomeButton />
            </section>
            <Footer />
        </main>
    );
}
