import { useEffect, useState } from "react";

function App() {
  const [backendMessage, setBackendMessage] = useState("");

  
  useEffect(() => {
  fetch("http://localhost:8000/hello")
    .then((res) => res.json())
    .then((data) => {
      console.log("Data received:", data); // 🔍 Log full response
      console.log("Message to set:", data.data.message); // 🔍 Log the message
      setBackendMessage(data.data.message); // ✅ Update state
      console.log("State updated to:", data.data.message); // 🔥 Confirm
    })
    .catch((err) => {
      console.error("Fetch error:", err); // ✅ Log any error
    });
}, []);   

  return (
    <div className="flex flex-col justify-center items-center h-screen bg-gray-900 text-white">
      <h1 className="text-4xl font-bold">React + GoFr + Tailwind 🚀</h1>
      <p className="mt-4 text-lg">{backendMessage}</p>
    </div>
  );
}

export default App;
