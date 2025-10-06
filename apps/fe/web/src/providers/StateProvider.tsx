"use client";
import React from "react";
import { createSessionStore, createPreferencesStore } from "@skillsier/state";

export const sessionStore = createSessionStore();
export const preferencesStore = createPreferencesStore();

export const StateContext = React.createContext({ session: sessionStore, prefs: preferencesStore });
export function StateProvider({ children }: { children: React.ReactNode }) {
  return <StateContext.Provider value={{ session: sessionStore, prefs: preferencesStore }}>{children}</StateContext.Provider>;
}
