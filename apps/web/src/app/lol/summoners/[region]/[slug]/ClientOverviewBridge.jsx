"use client";

import { useProfileData } from "./ProfileDataProvider";
import { ProfileOverview } from "@/app/components/LeagueProfile";

export default function ClientOverviewBridge() {
  const profileData = useProfileData();
  return <ProfileOverview data={profileData} />;
}
