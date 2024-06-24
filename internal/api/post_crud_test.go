package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"news-app/internal/api"
	"news-app/internal/config"
	"news-app/internal/service"
	"news-app/pkg/model"
	"strconv"
	"testing"
)

type MockPostTool struct {
	mock.Mock
}

func (m *MockPostTool) Create(ctx context.Context, post model.Post) error {
	args := m.Called(ctx, post)
	return args.Error(0)
}

func (m *MockPostTool) GetAll(ctx context.Context, args model.PaginationArgs) ([]model.Post, error) {
	argsCalled := m.Called(ctx, args)
	return argsCalled.Get(0).([]model.Post), argsCalled.Error(1)
}

func (m *MockPostTool) GetById(ctx context.Context, args model.PaginationByID) (*model.Post, error) {
	argsCalled := m.Called(ctx, args)
	return argsCalled.Get(0).(*model.Post), argsCalled.Error(1)
}

func (m *MockPostTool) Update(ctx context.Context, post model.Post) (model.Post, error) {
	args := m.Called(ctx, post)
	return args.Get(0).(model.Post), args.Error(1)
}

func (m *MockPostTool) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestPostGetByIdHandler(t *testing.T) {
	tests := []struct {
		name            string
		id              string
		limit           string
		offset          string
		mockGetByIdFunc func(ctx context.Context, args model.PaginationByID) (*model.Post, error)
		expectedStatus  int
		expectedBody    interface{}
	}{
		{
			name:   "successful retrieval",
			id:     "1",
			limit:  "10",
			offset: "0",
			mockGetByIdFunc: func(ctx context.Context, args model.PaginationByID) (*model.Post, error) {
				return &model.Post{ID: 1, Title: "Title 1", Content: "Content 1"}, nil
			},
			expectedStatus: http.StatusOK,
			expectedBody:   &model.Post{ID: 1, Title: "Title 1", Content: "Content 1"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPostTool := new(MockPostTool)
			mockPostTool.On("GetById", mock.Anything, model.PaginationByID{
				ID: mustAtoi(tt.id),
				PaginationArgs: model.PaginationArgs{
					Limit:  mustAtoi(tt.limit),
					Offset: mustAtoi(tt.offset),
				},
			}).Return(tt.mockGetByIdFunc(context.Background(), model.PaginationByID{
				ID: mustAtoi(tt.id),
				PaginationArgs: model.PaginationArgs{
					Limit:  mustAtoi(tt.limit),
					Offset: mustAtoi(tt.offset),
				},
			}))

			mockService := &service.Service{
				PostTool: mockPostTool,
			}

			cfg := &config.Config{
				HTTP: struct {
					Host string `envconfig:"HOST"`
					Port string `envconfig:"PORT"`
				}{
					Host: "localhost",
					Port: ":8080",
				},
			}
			a := api.New(cfg, mockService)

			req, err := http.NewRequest("GET", "/posts/"+tt.id+"?limit="+tt.limit+"&offset="+tt.offset, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			router := chi.NewRouter()
			router.Get("/posts/{id}", a.PostGetByIdHandler)
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			var responseBody interface{}
			err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
			assert.NoError(t, err)

			expectedBody, err := json.Marshal(tt.expectedBody)
			assert.NoError(t, err)

			assert.JSONEq(t, string(expectedBody), string(rr.Body.Bytes()))
		})
	}
}

func TestPostUpdateHandler(t *testing.T) {
	tests := []struct {
		name             string
		input            model.Post
		mockUpdateFunc   func(ctx context.Context, post model.Post) (model.Post, error)
		expectedStatus   int
		expectedResponse model.Post
	}{
		{
			name: "successful update",
			input: model.Post{
				ID:      1,
				Title:   "Updated Title",
				Content: "Updated Content",
			},
			mockUpdateFunc: func(ctx context.Context, post model.Post) (model.Post, error) {
				return model.Post{
					ID:      1,
					Title:   "Updated Title",
					Content: "Updated Content",
				}, nil
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: model.Post{ID: 1, Title: "Updated Title", Content: "Updated Content"},
		},
		{
			name: "service error",
			input: model.Post{
				ID:      1,
				Title:   "Updated Title",
				Content: "Updated Content",
			},
			mockUpdateFunc: func(ctx context.Context, post model.Post) (model.Post, error) {
				return model.Post{}, errors.New("service error")
			},
			expectedStatus:   http.StatusInternalServerError,
			expectedResponse: model.Post{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPostTool := new(MockPostTool)
			mockPostTool.On("Update", mock.Anything, tt.input).Return(tt.mockUpdateFunc(context.Background(), tt.input))

			mockService := &service.Service{
				PostTool: mockPostTool,
			}

			cfg := &config.Config{
				HTTP: struct {
					Host string `envconfig:"HOST"`
					Port string `envconfig:"PORT"`
				}{
					Host: "localhost",
					Port: ":8080",
				},
			}
			a := api.New(cfg, mockService)

			body, err := json.Marshal(tt.input)
			assert.NoError(t, err)

			req, err := http.NewRequest("PUT", "/posts", bytes.NewReader(body))
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(a.PostUpdateHandler)
			handler.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)

			var responseBody model.Post
			err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
			assert.NoError(t, err)

			assert.Equal(t, tt.expectedResponse, responseBody)
		})
	}
}

func TestPostDeleteHandler(t *testing.T) {
	tests := []struct {
		name           string
		id             string
		mockDeleteFunc func(ctx context.Context, id int) error
		expectedStatus int
	}{
		{
			name: "successful deletion",
			id:   "1",
			mockDeleteFunc: func(ctx context.Context, id int) error {
				return nil
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "invalid id",
			id:   "invalid",
			mockDeleteFunc: func(ctx context.Context, id int) error {
				return errors.New("invalid id")
			},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "service error",
			id:   "1",
			mockDeleteFunc: func(ctx context.Context, id int) error {
				return errors.New("service error")
			},
			expectedStatus: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockPostTool := new(MockPostTool)
			mockPostTool.On("Delete", mock.Anything, mustAtoi(tt.id)).Return(tt.mockDeleteFunc(context.Background(), mustAtoi(tt.id)))

			mockService := &service.Service{
				PostTool: mockPostTool,
			}

			cfg := &config.Config{
				HTTP: struct {
					Host string `envconfig:"HOST"`
					Port string `envconfig:"PORT"`
				}{
					Host: "localhost",
					Port: ":8080",
				},
			}
			a := api.New(cfg, mockService)

			req, err := http.NewRequest("DELETE", "/posts/"+tt.id, nil)
			assert.NoError(t, err)

			rr := httptest.NewRecorder()

			router := chi.NewRouter()
			router.Delete("/posts/{id}", a.PostDeleteHandler)
			router.ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
		})
	}
}

func mustAtoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
