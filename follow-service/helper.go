package main

import (
	"encoding/json"
	"io"
	"net/http"
	"xws/follow-service/data"
)

// renderJSON renders 'v' as JSON and writes it as a response into w.
func renderJSON(w http.ResponseWriter, v interface{}) {
	//span := tracer.StartSpanFromContext(ctx, "renderJSON")
	//defer span.Finish()

	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		return
	}

}

func decodeBody(r io.Reader) (*data.Follow, error) {
	//span := tracer.StartSpanFromContext(ctx, "decodeBody")
	//defer span.Finish()

	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var rt data.Follow
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return &rt, nil
}
