package pet

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/endpoint"
	kithttp "github.com/go-kit/kit/transport/http"
)

type Set struct {
	GetById endpoint.Endpoint
}

func NewHTTPClient(baseUrl string) (Set, error) {
	base, err := url.Parse(baseUrl)
	if err != nil {
		return Set{}, err
	}

	var getById endpoint.Endpoint
	{
		getById = kithttp.NewClient("GET", base.JoinPath("/api/v1/pets/{id}"), encodeGetByIdRequest, decodeGetByIdResponse).Endpoint()
	}

	return Set{
		GetById: getById,
	}, nil
}

func encodeGetByIdRequest(_ context.Context, req *http.Request, request interface{}) error {
	r := request.(GetByIdRequest)

	path, err := url.JoinPath(req.URL.Path, r.Id)
	if err != nil {
		return err
	}

	req.URL.Path = path
	return nil
}

func decodeGetByIdResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("error receiving response")
	}

	var out GetByIdResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	return out, nil
}

type GetByIdRequest struct {
	Id string `json:"id"`
}

type GetByIdResponse struct {
	Profile Pet
	Err     string
}
