import { getRecipeTree } from "@/lib/api"; // adjust the path if needed

export default async function Page() {
  const data = await getRecipeTree("Brick", "bfs", 3);

  return (
    <div>
      <h1>Tree for Brick</h1>
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </div>
  );
}
