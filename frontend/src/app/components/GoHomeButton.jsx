"use client"
import { useRouter } from "next/navigation";

export default function GoHomeButton() {
    const router = useRouter();

    return (
        <button
            onClick={() => router.push("/")}
            className="mt-4 px-6 py-2 rounded-md bg-[var(--pastel-red)] text-white font-bold hover:opacity-80 transition"
        >
            Home
        </button>
    );
}