package extra

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"go.uber.org/mock/gomock"

	hClient "github.com/ory/hydra-client-go/v2"
	kClient "github.com/ory/kratos-client-go"
)

//go:generate mockgen -build_flags=--mod=mod -package extra -destination ./mock_logger.go -source=../../internal/logging/interfaces.go
//go:generate mockgen -build_flags=--mod=mod -package extra -destination ./mock_extra.go -source=./interfaces.go

func TestHandleConsentSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := NewMockLoggerInterface(ctrl)
	mockService := NewMockServiceInterface(ctrl)

	session := kClient.NewSession("test", *kClient.NewIdentity("test", "test.json", "https://test.com/test.json", map[string]string{"name": "name"}))
	consent := hClient.NewOAuth2ConsentRequest("challenge")
	accept := hClient.NewOAuth2RedirectTo("test")

	req := httptest.NewRequest(http.MethodGet, "/api/consent", nil)

	values := req.URL.Query()
	values.Add("consent_challenge", "7bb518c4eec2454dbb289f5fdb4c0ee2")
	req.URL.RawQuery = values.Encode()

	w := httptest.NewRecorder()

	mockService.EXPECT().CheckSession(gomock.Any(), req.Cookies()).Return(session, nil)
	mockService.EXPECT().GetConsent(gomock.Any(), "7bb518c4eec2454dbb289f5fdb4c0ee2").Return(consent, nil)
	mockService.EXPECT().AcceptConsent(gomock.Any(), session.Identity, consent).Return(accept, nil)

	mux := chi.NewMux()
	NewAPI(mockService, mockLogger).RegisterEndpoints(mux)

	mux.ServeHTTP(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected HTTP status code 200 got %v", res.StatusCode)
	}

	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		t.Fatalf("expected error to be nil got %v", err)
	}

	redirect := hClient.NewOAuth2RedirectToWithDefaults()
	if err := json.Unmarshal(data, redirect); err != nil {
		t.Fatalf("expected error to be nil got %v", err)
	}

	if redirect.RedirectTo != accept.RedirectTo {
		t.Fatalf("expected %s, got %s.", accept.RedirectTo, redirect.RedirectTo)
	}
}

func TestHandleConsentFailOnAcceptConsent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := NewMockLoggerInterface(ctrl)
	mockService := NewMockServiceInterface(ctrl)

	session := kClient.NewSession("test", *kClient.NewIdentity("test", "test.json", "https://test.com/test.json", map[string]string{"name": "name"}))
	consent := hClient.NewOAuth2ConsentRequest("challenge")

	req := httptest.NewRequest(http.MethodGet, "/api/consent", nil)

	values := req.URL.Query()
	values.Add("consent_challenge", "7bb518c4eec2454dbb289f5fdb4c0ee2")
	req.URL.RawQuery = values.Encode()

	w := httptest.NewRecorder()

	mockService.EXPECT().CheckSession(gomock.Any(), req.Cookies()).Return(session, nil)
	mockService.EXPECT().GetConsent(gomock.Any(), "7bb518c4eec2454dbb289f5fdb4c0ee2").Return(consent, nil)
	mockService.EXPECT().AcceptConsent(gomock.Any(), session.Identity, consent).Return(nil, fmt.Errorf("error"))
	mockLogger.EXPECT().Errorf(gomock.Any(), gomock.Any()).Times(1)

	mux := chi.NewMux()
	NewAPI(mockService, mockLogger).RegisterEndpoints(mux)

	mux.ServeHTTP(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("expected HTTP status code 403 got %v", res.StatusCode)
	}
}

func TestHandleConsentFailOnGetConsent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := NewMockLoggerInterface(ctrl)
	mockService := NewMockServiceInterface(ctrl)

	session := kClient.NewSession("test", *kClient.NewIdentity("test", "test.json", "https://test.com/test.json", map[string]string{"name": "name"}))

	req := httptest.NewRequest(http.MethodGet, "/api/consent", nil)

	values := req.URL.Query()
	values.Add("consent_challenge", "7bb518c4eec2454dbb289f5fdb4c0ee2")
	req.URL.RawQuery = values.Encode()

	w := httptest.NewRecorder()

	mockService.EXPECT().CheckSession(gomock.Any(), req.Cookies()).Return(session, nil)
	mockService.EXPECT().GetConsent(gomock.Any(), "7bb518c4eec2454dbb289f5fdb4c0ee2").Return(nil, fmt.Errorf("error"))
	mockLogger.EXPECT().Errorf(gomock.Any(), gomock.Any()).Times(1)

	mux := chi.NewMux()
	NewAPI(mockService, mockLogger).RegisterEndpoints(mux)

	mux.ServeHTTP(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("expected HTTP status code 403 got %v", res.StatusCode)
	}
}

func TestHandleConsentFailOnCheckSession(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLogger := NewMockLoggerInterface(ctrl)
	mockService := NewMockServiceInterface(ctrl)

	req := httptest.NewRequest(http.MethodGet, "/api/consent", nil)

	values := req.URL.Query()
	values.Add("consent_challenge", "7bb518c4eec2454dbb289f5fdb4c0ee2")
	req.URL.RawQuery = values.Encode()

	w := httptest.NewRecorder()

	mockService.EXPECT().CheckSession(gomock.Any(), req.Cookies()).Return(nil, fmt.Errorf("error"))
	mockLogger.EXPECT().Errorf(gomock.Any(), gomock.Any()).Times(1)

	mux := chi.NewMux()
	NewAPI(mockService, mockLogger).RegisterEndpoints(mux)

	mux.ServeHTTP(w, req)

	res := w.Result()

	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("expected HTTP status code 403 got %v", res.StatusCode)
	}
}
