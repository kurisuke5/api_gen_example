package health_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"

	domain "github.com/kurisuke5/api_gen_example/backend/domain/health"
	"github.com/kurisuke5/api_gen_example/backend/domain/health/mock_health"
	"github.com/kurisuke5/api_gen_example/backend/interfaces/api/health"
	"github.com/kurisuke5/api_gen_example/backend/interfaces/props"
	"github.com/kurisuke5/api_gen_example/backend/interfaces/wrapper"
)

func TestGetHealthCheckController_GetHealthCheck(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		setup   func(*mock_health.MockRepository)
		req     *health.GetHealthCheckRequest
		wantRes *health.GetHealthCheckResponse
		wantErr *wrapper.APIError
	}{
		{
			name: "success",
			setup: func(ar *mock_health.MockRepository) {
				ar.EXPECT().Check(gomock.Any()).Return(nil)
			},
			wantRes: &health.GetHealthCheckResponse{
				Status: http.StatusOK,
			},
			wantErr: nil,
		},
		{
			name: "error",
			setup: func(ar *mock_health.MockRepository) {
				ar.EXPECT().Check(gomock.Any()).Return(xerrors.New("unhealthy database"))
			},
			wantErr: &wrapper.APIError{
				Status: http.StatusInternalServerError,
				Body:   nil,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			hr := mock_health.NewMockRepository(ctrl)
			u := domain.NewHealthService(hr)
			g := &health.GetHealthCheckController{
				ControllerProps: &props.ControllerProps{
					HealthService: u,
				},
			}

			e := echo.New()
			c := e.NewContext(new(http.Request), httptest.NewRecorder())

			tt.setup(hr)

			got, err := g.GetHealthCheck(c, tt.req)
			if tt.wantErr != nil {
				if err == nil {
					t.Error("expected error, but got nil")
					return
				}
				if diff := cmp.Diff(tt.wantErr, err.(*wrapper.APIError),
					cmpopts.IgnoreFields(wrapper.APIError{}, "err")); diff != "" {
					t.Errorf("unexpected error:(-want +got):\n%s", diff)
				}
				return
			}

			if err != nil {
				t.Errorf("err should be nil, but got %q", err)
				return
			}
			if diff := cmp.Diff(tt.wantRes, got); diff != "" {
				t.Errorf("GetHealthCheck() mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
