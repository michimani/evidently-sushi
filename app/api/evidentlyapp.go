package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sushi/libs"
	"time"
)

var (
	entityIDParamKey string = "entity_id"
	evidentlyProject string = "FoodProject"
	evidentlyFeature string = "SushiFeature"
)

type EvidentlyApp struct{}

var _ ServerInterface = (*EvidentlyApp)(nil)

func NewEvidentlyApp() *EvidentlyApp {
	return &EvidentlyApp{}
}

func (t *EvidentlyApp) EvaluateFeature(w http.ResponseWriter, r *http.Request, params EvaluateFeatureParams) {
	qp, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeErrorResponse(&w, err)
		return
	}

	entityID := fmt.Sprintf("%x", time.Now().UnixNano())
	if e, ok := qp[entityIDParamKey]; ok && len(e) > 0 {
		entityID = qp[entityIDParamKey][0]
	}
	in := &libs.EvaluateFeatureInput{
		Project:  evidentlyProject,
		Feature:  evidentlyFeature,
		EntityID: entityID,
	}

	c, err := newEvidentlyClient()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeErrorResponse(&w, err)
	}

	res, err := c.EvaluateFeature(context.Background(), in)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeErrorResponse(&w, err)
	}

	result := Feature{
		Value:    &res.StringValue,
		Reason:   &res.Reason,
		Name:     &evidentlyFeature,
		EntityId: &entityID,
	}

	w.WriteHeader(http.StatusOK)
	jerr := json.NewEncoder(w).Encode(result)
	if jerr != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func newEvidentlyClient() (*libs.EvidentlyClient, error) {
	cin := &libs.NewClientInput{
		Context: context.Background(),
		Region:  os.Getenv("AWS_DEFAULT_REGION"),
	}
	return libs.NewEvidentlyClient(cin)
}

func writeErrorResponse(w *http.ResponseWriter, err error) {
	_ = json.NewEncoder(*w).Encode(Error{
		Message: err.Error(),
	})
}
