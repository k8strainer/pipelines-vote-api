package main

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

// TestVoteEndpoint prüft das GET/POST-Verhalten auf /vote
func TestVoteEndpoint(t *testing.T) {
    // 1) Router vorbereiten
    router := setupRouter()

    // 2) Test: GET /vote -> sollte 200 liefern und anfänglich {"a":0,"b":0}
    reqGET, _ := http.NewRequest("GET", "/vote", nil)
    wGET := httptest.NewRecorder()
    router.ServeHTTP(wGET, reqGET)

    if wGET.Code != http.StatusOK {
        t.Errorf("GET /vote: erwartet 200, bekam %d", wGET.Code)
    }

    // 3) Test: POST /vote -> Ein Vote für "a"
    voteBody := `{"voter_id":"Alice","vote":"a"}`
    reqPOST, _ := http.NewRequest("POST", "/vote", strings.NewReader(voteBody))
    wPOST := httptest.NewRecorder()
    router.ServeHTTP(wPOST, reqPOST)

    if wPOST.Code != http.StatusOK {
        t.Errorf("POST /vote: erwartet 200, bekam %d", wPOST.Code)
    }

    // 4) Nochmals GET /vote -> a sollte 1 sein, b = 0
    reqGET2, _ := http.NewRequest("GET", "/vote", nil)
    wGET2 := httptest.NewRecorder()
    router.ServeHTTP(wGET2, reqGET2)

    if wGET2.Code != http.StatusOK {
        t.Errorf("GET /vote (nach POST): erwartet 200, bekam %d", wGET2.Code)
    }

    // Hier könnte man z. B. das JSON auswerten, um sicherzustellen,
    // dass "a" = 1 und "b" = 0 darin enthaltne ist.
    // (Bei Go 1.4 ist noch kein "t.Logf(...)" vorhanden,
    //  aber prinzipiell kannst es mit fmt.Println(wGET2.Body.String()) gecheckt werden.)
}
