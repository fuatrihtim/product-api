package product

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"io"
	"net/http/httptest"
	"testing"
)

func TestCreate(t *testing.T) {
	type wanted struct {
		statusCode   int
		expectedKeys []string
	}

	// Create fiber app for testing purposes
	app := fiber.New()
	app.Post("/", Create)

	// Tests struct
	tests := []struct {
		name        string
		description string
		endpoint    string
		payload     any
		want        wanted
	}{
		{
			name:        "Create a fake product",
			description: "This test should return 200 status code and created product",
			endpoint:    "/",
			payload: map[string]any{
				"code":         "testCode",
				"name":         "testname",
				"stokquantity": 10,
			},
			want: wanted{
				statusCode: 200,
				expectedKeys: []string{
					"Id",
					"Code",
					"Name",
					"StockQuantity",
				},
			},
		},
		{
			name:        "Create a product with blank code",
			description: "This test should return 400 status code and error message",
			endpoint:    "/",
			payload: map[string]any{
				"code":         "",
				"name":         "testname2",
				"stokquantity": 20,
			},
			want: wanted{
				statusCode: 400,
				expectedKeys: []string{
					"message",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.payload)
			if err != nil {
				t.Errorf("Cannot parse body: %v", err)
				t.Fail()
			}

			req := httptest.NewRequest("POST", tt.endpoint, bytes.NewReader(body))
			req.Header.Add("Content-Type", "application/json")

			// Create request
			res, err := app.Test(req)
			if err != nil {
				t.Errorf("Cannot test Fiber handler: %v", err)
				t.Fail()
			}

			// Assertions
			if res.StatusCode != tt.want.statusCode {
				t.Errorf("Expected status code %d, got %d", tt.want.statusCode, res.StatusCode)
			}

			answer, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Cannot parse body: %v", err)
			}

			var message map[string]any
			err = json.Unmarshal([]byte(answer), &message)
			if err != nil {
				t.Errorf("Cannot unmarshal response: %v", err)
			}

			for _, s := range tt.want.expectedKeys {
				if _, ok := message[s]; !ok {
					t.Errorf("Expected response body to contain key %s but it can not be found", s)
				}
			}
		})
	}
}
