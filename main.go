package homebridge

import (
    "github.com/go-resty/resty/v2"
    "log"
	"encoding/json"
)

type APIClient struct {
    httpClient *resty.Client
    BaseURL    string
    Token      string
}

// NewAPIClient creates a new API client instance
func NewAPIClient(baseURL string) *APIClient {
    return &APIClient{
        httpClient: resty.New(),
        BaseURL:    baseURL,
    }
}

type LoginCredentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
    OTP      string `json:"otp,omitempty"`
}

type LoginResponse struct {
    Token string `json:"token"`
}

func (c *APIClient) Login(username, password, otp string) error {
    credentials := LoginCredentials{
        Username: username,
        Password: password,
        OTP:      otp,
    }

    resp, err := c.httpClient.R().
        SetHeader("Content-Type", "application/json").
        SetBody(credentials).
        Post(c.BaseURL + "/api/auth/login")

    if err != nil {
        return err
    }

    var loginResp LoginResponse
    if err := json.Unmarshal(resp.Body(), &loginResp); err != nil {
        return err
    }

    c.Token = loginResp.Token
    return nil
}

func (c *APIClient) GetServerPairing() {
    // Example: GET request using the token
    resp, err := c.httpClient.R().
        SetAuthToken(c.Token).
        Get(c.BaseURL + "/api/server/pairing")

    if err != nil {
        log.Fatalf("Error on request: %v", err)
    }

    log.Printf("Response: %v", resp)
}
