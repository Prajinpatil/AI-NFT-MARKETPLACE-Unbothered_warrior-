package api

import (
    "bytes"
    "encoding/json"
    "io"
    "net/http"

    "gofr.dev/pkg/gofr"
)

func MintNFT(ctx *gofr.Context) (interface{}, error) {
    url := "https://api.verbwire.com/v1/nft/mint/quickMint"
    apiKey := "<YOUR_VERBWIRE_API_KEY>"

    payload := map[string]string{
        "chain":            "goerli",
        "name":             "HackOdisha NFT",
        "description":      "Demo NFT minted from Verbwire + GoFr",
        "recipientAddress": "<YOUR_WALLET_ADDRESS>",
        "imageUrl":         "https://verbwire.io/sample-nft.png",
    }

    // Marshal payload
    body, err := json.Marshal(payload)
    if err != nil {
        return nil, err // ✅ Check error
    }

    // Create request
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
    if err != nil {
        return nil, err // ✅ Check error
    }

    // Set headers
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-API-Key", apiKey)

    // Execute request
    resp, err := http.DefaultClient.Do(req)   
    if err != nil {
        return nil, err // ✅ Check error before using resp
    }
    defer resp.Body.Close() // ✅ Safe to use now

    // Read response body
    responseBody, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err // ✅ Check error
    }

    // Return response as string
    return string(responseBody), nil
}   