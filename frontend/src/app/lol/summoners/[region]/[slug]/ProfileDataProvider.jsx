"use client";

import { createContext, useContext } from "react";

export const ProfileDataContext = createContext(null);

export function ProfileDataProvider({ value, children }) {
  return <ProfileDataContext.Provider value={value}>{children}</ProfileDataContext.Provider>;
}

export function useProfileData() {
  return useContext(ProfileDataContext);
}
