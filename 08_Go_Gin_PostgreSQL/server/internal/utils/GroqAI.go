// package utils

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"time"
// )

// var AIClient = &http.Client{
// 	Timeout: 15 * time.Second,
// }

// func AIRequest(ctx context.Context, url string, apiKey string, payload any, result any) error {
// 	// convert payload to JSON
// 	body, err := json.Marshal(payload)
// 	if err != nil {
// 		return err
// 	}

// 	// create request
// 	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
// 	if err != nil {
// 		return err
// 	}

// 	// headers
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+apiKey)

// 	// send request
// 	resp, err := AIClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// check status
// 	if resp.StatusCode != http.StatusOK {
// 		b, _ := io.ReadAll(resp.Body)
// 		return fmt.Errorf("ai api error: %s", string(b))
// 	}

//		// decode response into struct
//		return json.NewDecoder(resp.Body).Decode(result)
//	}

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

func AIRequest(ctx context.Context, url string, apikey string, payload any, result any) error {
 
 
	// convert payload to json
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// make new req
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return  err
	}

	// set headers 
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apikey)

	// send req
	res, err := AIClient.Do(req)
	if err != nil {
		return  err
	}
	defer res.Body.Close()

	// check status code 
	if res.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(res.Body)
		return fmt.Errorf("err %s", string(b))
	}

	// convert resp to json and store in struct
	return json.NewDecoder(res.Body).Decode(result)


}