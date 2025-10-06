import { createStore } from "../core/createStore";
export const createKycStore = () =>
  createStore<{ status: "todo" | "in_progress" | "done"; set: (s: "todo"|"in_progress"|"done") => void }>((set) => ({
    status: "todo",
    set: (s) => set({ status: s })
  }), { name: "users.kyc", devtools: true });
