package memberships

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/zuhdi751/zd_music_catalog/internal/models/memberships"
	"go.uber.org/mock/gomock"
)

func TestHandler_SignUp(t *testing.T) {
	// delete fields and args
	/* ---xxx--- */
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	/* ---xxx--- */
	/* ---xxx--- */
	mockSvc := NewMockservice(ctrlMock)
	/* ---xxx--- */

	tests := []struct {
		name               string
		mockFn             func()
		expectedStatusCode int
	}{
		// TODO: Add test cases.
		{
			name: "success",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "testpassword",
				}).Return(nil)
			},
			expectedStatusCode: 201,
		},
		{
			name: "failed",
			mockFn: func() {
				mockSvc.EXPECT().SignUp(memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "testpassword",
				}).Return(errors.New("email or username is already exists"))
			},
			expectedStatusCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			/* ---xxx--- */
			tt.mockFn()
			/* ---xxx--- */

			/* ---xxx--- */
			api := gin.New()
			/* ---xxx--- */

			h := &Handler{
				Engine:  api,
				service: mockSvc,
			}
			// h.SignUp(tt.args.c) // to be deleted

			/* ---xxx--- */
			h.ResgisterRoute()

			w := httptest.NewRecorder()

			endpoint := `/memberships/sign_up`
			model := memberships.SignUpRequest{
				Email:    "test@gmail.com",
				Username: "testusername",
				Password: "testpassword",
			}

			val, err := json.Marshal(model)
			assert.NoError(t, err)

			body := bytes.NewReader(val)
			req, err := http.NewRequest(http.MethodPost, endpoint, body)
			assert.NoError(t, err)

			h.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			/* ---xxx--- */
		})
	}
}
