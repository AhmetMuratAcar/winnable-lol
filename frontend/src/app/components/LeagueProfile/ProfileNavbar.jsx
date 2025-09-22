"use client";

import Link from "next/link";
import { usePathname, useParams } from "next/navigation";

export default function ProfileNavbar() {
  const pathname = usePathname();
  const params = useParams();

  const region = Array.isArray(params.region) ? params.region[0] : params.region;
  const slug = Array.isArray(params.slug) ? params.slug[0] : params.slug;

  const basePath = `/lol/summoners/${region}/${slug}`;

  const isMastery =
    pathname === `${basePath}/mastery` || pathname.startsWith(`${basePath}/mastery/`);
  const isLiveGame =
    pathname === `${basePath}/live-game` || pathname.startsWith(`${basePath}/live-game/`);
  const isOverview = !isMastery && !isLiveGame;

  const baseBtn = "px-12 py-1 rounded-md font-bold text-white";
  const active = "bg-(--pastel-red)";
  const inactive = "hover:bg-(--background)";

  return (
    <div className="w-full bg-(--contrast) rounded border-(--contrast-border) border-1 space-x-5 p-2">
      <Link href={basePath} className={`${baseBtn} ${isOverview ? active : inactive}`}>
        Overview
      </Link>

      <Link href={`${basePath}/mastery`} className={`${baseBtn} ${isMastery ? active : inactive}`}>
        Masteries
      </Link>

      <Link
        href={`${basePath}/live-game`}
        className={`${baseBtn} ${isLiveGame ? active : inactive}`}
      >
        Live Game
      </Link>
    </div>
  );
}
