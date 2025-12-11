"use client";
export default function RefreshButton() {
  // REMEMBER TO REVAILDATE THE TAGS OF ALL CACHED DATA IN PROFILE PAGE
  // getProfile, getMastery atm
  return (
    <button
      onClick={() => console.log("Refreshing profile...")}
      className="mt-4 px-5 py-2 rounded-md bg-(--pastel-red) text-white hover:opacity-80 hover:cursor-pointer"
    >
      Update
    </button>
  );
}
