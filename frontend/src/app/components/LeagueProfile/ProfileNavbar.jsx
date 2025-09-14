"use client";
export default function ProfileNavbar() {
  return (
    <div className="w-full bg-(--contrast) rounded border-(--contrast-border) border-1 space-x-5 p-2">
      <button
        onClick={() => console.log("Overview section")}
        className="bg-(--pastel-red) px-12 py-1 rounded-md text-white hover:opacity-80 hover:cursor-pointer font-bold"
      >
        Overview
      </button>

      <button
        onClick={() => console.log("Masteries section")}
        className="px-12 py-1 rounded-md text-white hover:bg-(--background) hover:cursor-pointer font-bold"
      >
        Masteries
      </button>

      <button
        onClick={() => console.log("Live Game section")}
        className="px-12 py-1 rounded-md text-white hover:bg-(--background) hover:cursor-pointer font-bold"
      >
        Live Game
      </button>
    </div>
  );
}
