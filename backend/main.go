package main

import (
    "fmt"
    "io"
    "net/http"
    "net/url"
    "strings"

    "gofr.dev/pkg/gofr"
)

func main() {
    app := gofr.New()

    // ------------------------------
    // ✅ CORS Middleware
    // ------------------------------
    app.Server.Router.Use(func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // frontend
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

            if r.Method == "OPTIONS" {
                w.WriteHeader(http.StatusOK)
                return
            }
            next.ServeHTTP(w, r)
        })
    })

    // ------------------------------
    // Test route
    // ------------------------------
    app.GET("/hello", func(ctx *gofr.Context) (interface{}, error) {
        return map[string]interface{}{"message": "Hello from Go backend"}, nil
    })

    // ------------------------------
    // NFT Mint route
    // ------------------------------
    app.POST("/api/mint-nft", func(ctx *gofr.Context) (interface{}, error) {
        endpoint := "https://api.verbwire.com/v1/nft/mint/quickMintFromMetadataUrl"
        apiKey := "YOUR_VERBWIRE_API_KEY" // replace with your key

        // Form values
        data := url.Values{}
        data.Set("allowPlatformToOperateToken", "true")
        data.Set("chain", "sepolia") // or "mumbai"
        data.Set("recipientAddress", "0x33145a6258e89b6E0796d237A3048A3852cCaeQ7") // your wallet
        data.Set("metadataUrl", "https://ipfs.io/ipfs/bafkreigjkcafrutdcbicyr3new6aoowgbscf6wgqyty45ckd3xur7ymldm")

        req, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
        if err != nil {
            return nil, err
        }

        req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
        req.Header.Set("X-API-Key", apiKey)

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            return nil, err
        }
        defer resp.Body.Close()

        body, err := io.ReadAll(resp.Body)
        if err != nil {
            return nil, err
        }

        fmt.Printf("✅ Verbwire Response Status: %d\n", resp.StatusCode)
        fmt.Printf("🔹 Raw Verbwire Response: %s\n", string(body))

        // ✅ Return raw response to frontend
        return map[string]interface{}{
            "success": true,
            "message": "NFT mint request sent!",
            "data":    string(body),
        }, nil
    })

    fmt.Println("⚡ Backend running on port 8000")
    app.Run()
}
