export async function getRecipeTree(element: string) {
  const response = await fetch(
    `http://localhost:8080/api/tree?element=${element}`
  );
  return await response.json();
}
