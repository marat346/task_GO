package handler

func TestHandler(t *testing.T) {
	srv := http.NewServer(http.HandlerFunc(handler))

	resp, err := http.Get(srv.URL)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	textBytes, err := ioutil.ReaAll(resp.Body)
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	defer resp.Body.Close()

	text := string(textBytes)
	if text != "message from test hadler" {
		t.Log(err)
		t.Fail()
	}
}
