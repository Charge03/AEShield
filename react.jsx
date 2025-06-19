import React, { useState } from "react";
import 'bootstrap/dist/css/bootstrap.min.css';

export default function AEShieldApp() {
  const [file, setFile] = useState(null);
  const [key, setKey] = useState("");
  const [status, setStatus] = useState("");
  const [endpoint, setEndpoint] = useState("encrypt");
  const [processing, setProcessing] = useState(false);

  const isValidHexKey = (k) => /^[0-9a-fA-F]{64}$/.test(k);

  const handleSubmit = async (e) => {
    e.preventDefault();

    if (!file) {
      setStatus("Please select a file.");
      return;
    }
    if (!isValidHexKey(key)) {
      setStatus("Key must be exactly 64 hex characters (32 bytes).");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);
    formData.append("key", key);

    try {
      setProcessing(true);
      setStatus(`${endpoint === "encrypt" ? "Encrypting" : "Decrypting"}...`);

      const res = await fetch(`http://localhost:8080/${endpoint}`, {
        method: "POST",
        body: formData,
      });

      if (!res.ok) {
        const errorText = await res.text();
        throw new Error(errorText || "Server error");
      }

      const blob = await res.blob();
      const filename =
        file.name + (endpoint === "encrypt" ? ".enc" : ".dec");

      const downloadUrl = window.URL.createObjectURL(blob);
      const a = document.createElement("a");
      a.href = downloadUrl;
      a.download = filename;
      a.click();
      window.URL.revokeObjectURL(downloadUrl);

      setStatus("Done!");
    } catch (err) {
      setStatus(`Error: ${err.message}`);
    } finally {
      setProcessing(false);
    }
  };

  return (
    <div className="container">
      <h1 className="title">AEShield</h1>
      <form onSubmit={handleSubmit}>
        <label htmlFor="file">Select file to process:</label>
        <input
          id="file"
          type="file"
          onChange={(e) => setFile(e.target.files[0])}
          disabled={processing}
        />

        <label htmlFor="key">64-character Hex Key:</label>
        <input
          id="key"
          type="text"
          placeholder="Enter 64-character hex key"
          value={key}
          onChange={(e) => setKey(e.target.value)}
          disabled={processing}
        />

        <label htmlFor="operation">Operation:</label>
        <select
          id="operation"
          value={endpoint}
          onChange={(e) => setEndpoint(e.target.value)}
          disabled={processing}
        >
          <option value="encrypt">Encrypt</option>
          <option value="decrypt">Decrypt</option>
        </select>

        <button type="submit" disabled={processing}>
          {processing ? "Processing..." : "Submit"}
        </button>
      </form>

      {status && <p className="status">{status}</p>}
    </div>
  );
}
