"use client";
export default function RefreshButton() {
  return (
    <button
      onClick={console.log("Refreshing profile...")}
      className="mt-4 px-3 py-2 rounded-md bg-[var(--pastel-red)] text-white hover:opacity-80 hover:cursor-pointer"
    >
      Update
    </button>
  );
}
