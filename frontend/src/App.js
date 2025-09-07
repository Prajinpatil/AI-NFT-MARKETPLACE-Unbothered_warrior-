import { useState } from "react";

function App() {
  const [status, setStatus] = useState("");
  const [txID, setTxID] = useState("");
  const [txHash, setTxHash] = useState("");
  const [tokenId, setTokenId] = useState("");
  const [minting, setMinting] = useState(false);

  const generateAndMint = async () => {
    setMinting(true);
    setStatus("Generating metadata...");

    // 1) generate metadata -> returns metadataUrl
    const genRes = await fetch("http://localhost:8000/api/generate", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ prompt: "futuristic library badge, neon, vector" }),
    });
    const genData = await genRes.json();
    if (!genData.metadataUrl) {
      setStatus("Metadata generation failed.");
      setMinting(false);
      return;
    }

    setStatus("Metadata ready. Minting NFT...");
    const mintRes = await fetch("http://localhost:8000/api/mint-nft", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ metadataUrl: genData.metadataUrl, chain: "sepolia" }),
    });
    const mintData = await mintRes.json();
    console.log("mintData", mintData);
    if (!mintData.transactionID) {
      setStatus("Failed to start minting.");
      setMinting(false);
      return;
    }
    setTxID(mintData.transactionID);
    setStatus("Mint started. Polling for confirmation...");

    // Poll for final txHash/tokenId
    const poll = setInterval(async () => {
      const statusRes = await fetch(`http://localhost:8000/api/mint-status/${mintData.transactionID}`);
      const statusData = await statusRes.json();
      console.log("statusData", statusData);

      if (statusData.txHash) {
        setTxHash(statusData.txHash);
        setTokenId(statusData.tokenId || "");
        setStatus("NFT minted! 🎉");
        clearInterval(poll);
        setMinting(false);
      }
    }, 5000);
  };

  return (
    <div className="flex flex-col items-center justify-center h-screen bg-gray-900 text-white">
      <h1 className="text-4xl font-bold">AI → IPFS → Verbwire Mint</h1>
      <p className="my-4">{status}</p>

      {txID && <p>TransactionID: {txID}</p>}
      {txHash && <p>TxHash: {txHash}</p>}
      {tokenId && <p>TokenId: {tokenId}</p>}

      <button onClick={generateAndMint} className="px-6 py-3 bg-indigo-600 rounded">
        {minting ? "Working..." : "Generate & Mint NFT"}
      </button>
    </div>
  );
}

export default App;
