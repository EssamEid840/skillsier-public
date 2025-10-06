import { usersFetch } from "../clients/rest";

export type PortfolioItem = {
  id: string;
  title: string;
  imageUri?: string;
  url?: string;
};

export async function listPortfolio() {
  return usersFetch<PortfolioItem[]>("rest/portfolio", { method: "GET" });
}

export async function createPortfolio(item: Omit<PortfolioItem, "id">) {
  return usersFetch<PortfolioItem>("rest/portfolio", {
    method: "POST",
    body: JSON.stringify(item)
  });
}

export async function deletePortfolio(id: string) {
  return usersFetch<{ ok: true }>(`rest/portfolio/${id}`, { method: "DELETE" });
}
