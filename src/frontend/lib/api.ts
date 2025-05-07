type TraversalMode = "bfs" | "dfs" | "bd";

export async function getRecipeTree(
  element: string,
  mode: TraversalMode,
  amount: number
) {
  const response = await fetch(
    `http://localhost:8080/api/tree?element=${element}&mode=${mode}&amount=${amount}`
  );
  return await response.json();
}
