import { useEffect, useState } from "react";

function App() {
  const [backendMessage, setBackendMessage] = useState("");

  useEffect(() => {
    fetch("http://localhost:8000/hello")
      .then((res) => res.json())
      .then((data) => setBackendMessage(data.message))
      .catch((err) => console.error(err));
  }, []);

  return (
    <div className="flex flex-col justify-center items-center h-screen bg-gray-900 text-white">
      <h1 className="text-4xl font-bold">React + GoFr + Tailwind 🚀</h1>
      <p className="mt-4 text-lg">{backendMessage}</p>
    </div>
  );
}

export default App;
