// TODO: add a proper return type instead of unknown
export async function getMatch(matchId: string): Promise<unknown> {
  const url = new URL(`${process.env.NEXT_PUBLIC_BACKEND_URL}/lol/match`);
  url.searchParams.set("matchID", matchId);

  const res = await fetch(url.toString(), { cache: "no-store" });

  if (!res.ok) {
    throw new Error(`Failed to fetch match ${matchId}: ${res.status} ${res.statusText}`);
  }

  return res.json();
}
