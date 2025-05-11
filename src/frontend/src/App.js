import HomePage from "./pages/HomePage";
import { useEffect } from "react";

function App() {
  useEffect(() => {
    document.title = "Recipe Calculator | Little Alchemy 2";
  }, []);
  return <HomePage />;
}

export default App;
