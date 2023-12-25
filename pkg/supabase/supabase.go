package supabase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type LoginWithEmailAndPasswordResponse struct {
	AccessToken string `json:"access_token"`
}

func LoginWithEmailAndPassword(email, password string) (string, error) {
	supabaseURL := os.Getenv("SUPABASE_URI") + "/auth/v1/token?grant_type=password"
	supabaseKey := os.Getenv("SUPABASE_ANON2")

	supabaseRequest := map[string]string{
		"email":    email,
		"password": password,
	}

	jsonData, err := json.Marshal(supabaseRequest)
	if err != nil {
		return "", err
	}

	client := &http.Client{}

	loginRequest, err := http.NewRequest("POST", supabaseURL, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err
	}

	loginRequest.Header.Add("Authorization", "Bearer "+supabaseKey)
	loginRequest.Header.Add("apikey", supabaseKey)

	response, err := client.Do(loginRequest)

	if err != nil {
		fmt.Println("Error posting: ", err)
		return "", err
	}

	defer loginRequest.Body.Close()

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	var jsonResponse LoginWithEmailAndPasswordResponse

	err = json.Unmarshal(responseBody, &jsonResponse)

	if err != nil {
		fmt.Println("Error unmarshalling JSON jsonResponse2:", err)
		return "", err
	}

	return jsonResponse.AccessToken, nil
}
