package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var AIClient = &http.Client{
	Timeout: 15*time.Second,
}

func AIRequest(ctx context.Context, url string, apiKey string, payload any , result any) error {
	// convert payload into json
	body, err := json.Marshal(payload)
	if err != nil {
		return  err
	}

	// create new request
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	// set headers 
	req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+apiKey)

	// send request
	res, err := AIClient.Do(req)
	if err != nil {
		return  err
	}
	defer res.Body.Close()

	// read resp
	respBody, err := io.ReadAll(res.Body)
    if err != nil {
		return  err
	}

	// check status code 
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("err %s", string(respBody))
	}

	return json.Unmarshal(respBody, result)
}