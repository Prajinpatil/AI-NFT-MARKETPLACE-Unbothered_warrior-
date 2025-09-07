package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"gofr.dev/pkg/gofr"
)

func MintNFT(ctx *gofr.Context) (interface{}, error) {
	url := "https://api.verbwire.com/v1/nft/mint/quickMintFromMetadataUrl"
	apiKey := "sk_live_aa613753-3653-4d8c-b2ff-3c33492f9ca8"

	// read request body: { "metadataUrl": "https://..." }
	var reqData struct {
		MetadataUrl string `json:"metadataUrl"`
		Recipient   string `json:"recipient"` // optional, fallback
		Chain       string `json:"chain"`
	}
	if err := ctx.Bind(&reqData); err != nil {
		return nil, err
	}
	if reqData.MetadataUrl == "" {
		return nil, fmt.Errorf("metadataUrl required")
	}
	if reqData.Chain == "" {
		reqData.Chain = "sepolia"
	}
	if reqData.Recipient == "" {
		reqData.Recipient = "0x33145a6258e89b6E0796d237A3048A3852cCaeQ7" // your default
	}

	payload := fmt.Sprintf(
		"allowPlatformToOperateToken=true&chain=%s&recipientAddress=%s&metadataUrl=%s",
		reqData.Chain,
		reqData.Recipient,
		reqData.MetadataUrl,
	)

	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("X-API-Key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	fmt.Printf("✅ Verbwire Response Status: %d\n", resp.StatusCode)
	fmt.Printf("🔹 Raw Verbwire Response: %s\n", string(body))

	var parsed struct {
		QuickMint struct {
			TransactionID string `json:"transactionID"`
			Status        string `json:"status"`
			BlockExplorer string `json:"blockExplorer"`
		} `json:"quick_mint"`
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, err
	}
	return map[string]interface{}{
		"success":       true,
		"status":        parsed.QuickMint.Status,
		"transactionID": parsed.QuickMint.TransactionID,
		"blockExplorer": parsed.QuickMint.BlockExplorer,
		"raw":           string(body),
	}, nil
}
