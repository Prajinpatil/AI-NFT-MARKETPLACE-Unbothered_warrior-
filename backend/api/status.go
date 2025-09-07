package api

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "gofr.dev/pkg/gofr"
)

func CheckNFTStatus(ctx *gofr.Context) (interface{}, error) {
    transactionID := ctx.PathParam("id") // /api/check-status/:id
    apiKey := "sk_live_aa613753-3653-4d8c-b2ff-3c33492f9ca8"

    url := fmt.Sprintf("https://api.verbwire.com/v1/nft/data/transaction?transactionId=%s", transactionID)

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("accept", "application/json")
    req.Header.Set("X-API-Key", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)

    var result map[string]interface{}
    json.Unmarshal(body, &result)

    return result, nil
}
