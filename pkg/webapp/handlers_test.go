package webapp

import (
	"gosearch/pkg/crawler"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

var testMux *mux.Router

func TestMain(m *testing.M) {

	testMux = mux.NewRouter()
	wa := WebApp{
		addr: "127.0.0.1:8080",
		Pages: []crawler.Document{
			{
				ID:    123,
				Title: "Test title",
				URL:   "https://testurl.com/123",
			},
		},
		Words: map[string][]int{
			"testKey": {1, 2, 3},
		},
	}
	wa.endpoints(testMux)
	m.Run()
}

func Test_pagesHandler(t *testing.T) {
	// Создаём HTTP=запрос.
	req := httptest.NewRequest(http.MethodGet, "/pages", nil)

	// Объект для записи ответа HTTP-сервера.
	rr := httptest.NewRecorder()

	// Вызов маршрутизатора и обслуживание запроса.
	testMux.ServeHTTP(rr, req)

	// Анализ ответа сервера (неверный метод HTTP).
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	// t.Log("Response: ", rr.Body)

	//=========================================================

	req = httptest.NewRequest(http.MethodGet, "/pages", nil)

	rr = httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	body := rr.Body.String()
	if !strings.Contains(body, "======[1]=======\nDocumentID: 123\nTitle: Test title\nURL: https://testurl.com/123\n") {
		t.Fatal(body)
	}
}

func Test_wordsHandler(t *testing.T) {
	// Создаём HTTP=запрос.
	req := httptest.NewRequest(http.MethodGet, "/words", nil)

	// Объект для записи ответа HTTP-сервера.
	rr := httptest.NewRecorder()

	// Вызов маршрутизатора и обслуживание запроса.
	testMux.ServeHTTP(rr, req)

	// Анализ ответа сервера (неверный метод HTTP).
	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}

	// t.Log("Response: ", rr.Body)

	//=========================================================

	req = httptest.NewRequest(http.MethodGet, "/words", nil)

	rr = httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("код неверен: получили %d, а хотели %d", rr.Code, http.StatusOK)
	}
	body := rr.Body.String()
	want := "1: testKey[3]\n    ID: 1\n    ID: 2\n    ID: 3\n======[END]======="
	if !strings.Contains(body, want) {
		t.Log(want)
		t.Fatal(body)

	}
	// fmt.Fprintf(w, "%d: %s[%d]\n", i, s, len(document))
	// 	for _, id := range document {
	// 		fmt.Fprintf(w, "    ID: %v\n", id)
	// 	}
}
