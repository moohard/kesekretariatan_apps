package testutil

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

// FiberTestApp creates a Fiber app for testing
type FiberTestApp struct {
	App *fiber.App
	t   *testing.T
}

// NewFiberTestApp creates a new Fiber test app with default config
func NewFiberTestApp(t *testing.T) *FiberTestApp {
	app := fiber.New(fiber.Config{
		AppName: "SIKERMA Test",
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    "TEST_ERROR",
					"message": err.Error(),
				},
			})
		},
	})

	return &FiberTestApp{App: app, t: t}
}

// Request makes a test request to the app
func (ft *FiberTestApp) Request(method, path string, body any, headers map[string]string) *httptest.ResponseRecorder {
	var bodyReader bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&bodyReader).Encode(body); err != nil {
			ft.t.Fatalf("failed to encode body: %v", err)
		}
	}

	req := httptest.NewRequest(method, path, &bodyReader)
	req.Header.Set("Content-Type", "application/json")
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// Fiber v3 Test uses fiber.TestConfig
	resp, err := ft.App.Test(req, fiber.TestConfig{
		Timeout: 0, // No timeout
	})
	if err != nil {
		ft.t.Fatalf("request failed: %v", err)
	}

	// Convert to ResponseRecorder
	recorder := httptest.NewRecorder()
	recorder.Code = resp.StatusCode
	recorder.Body = &bytes.Buffer{}
	if _, err := recorder.Body.ReadFrom(resp.Body); err != nil {
		ft.t.Fatalf("failed to read response body: %v", err)
	}

	return recorder
}

// Get makes a GET request
func (ft *FiberTestApp) Get(path string, headers map[string]string) *httptest.ResponseRecorder {
	return ft.Request(fiber.MethodGet, path, nil, headers)
}

// Post makes a POST request
func (ft *FiberTestApp) Post(path string, body any, headers map[string]string) *httptest.ResponseRecorder {
	return ft.Request(fiber.MethodPost, path, body, headers)
}

// Put makes a PUT request
func (ft *FiberTestApp) Put(path string, body any, headers map[string]string) *httptest.ResponseRecorder {
	return ft.Request(fiber.MethodPut, path, body, headers)
}

// Delete makes a DELETE request
func (ft *FiberTestApp) Delete(path string, headers map[string]string) *httptest.ResponseRecorder {
	return ft.Request(fiber.MethodDelete, path, nil, headers)
}

// AssertStatus asserts the response status code
func AssertStatus(t *testing.T, recorder *httptest.ResponseRecorder, expected int) {
	if recorder.Code != expected {
		t.Errorf("expected status %d, got %d", expected, recorder.Code)
	}
}

// AssertJSON asserts the JSON response body
func AssertJSON[T any](t *testing.T, recorder *httptest.ResponseRecorder, expected T) {
	var actual T
	if err := json.NewDecoder(recorder.Body).Decode(&actual); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	expectedJSON, _ := json.Marshal(expected)
	actualJSON, _ := json.Marshal(actual)

	if string(expectedJSON) != string(actualJSON) {
		t.Errorf("expected JSON %s, got %s", expectedJSON, actualJSON)
	}
}

// ParseResponse parses the response body into the given type
func ParseResponse[T any](t *testing.T, recorder *httptest.ResponseRecorder) T {
	var result T
	if err := json.NewDecoder(recorder.Body).Decode(&result); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	return result
}
