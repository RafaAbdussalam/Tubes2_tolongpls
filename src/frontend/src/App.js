import HomePage from "./pages/HomePage";
import { useEffect } from "react";

function App() {
  useEffect(() => {
    document.title = "Recipe Finder | Little Alchemy 2";
  }, []);
  return <HomePage />;
}

export default App;
