package memberships

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zuhdi751/zd_music_catalog/internal/configs"
	"github.com/zuhdi751/zd_music_catalog/internal/models/memberships"
	"go.uber.org/mock/gomock"
	"gorm.io/gorm"
)

func Test_service_SignUp(t *testing.T) {

	// hapus filed bawaan dan ganti dengan ini
	/* ---xxx--- */
	ctrlMock := gomock.NewController(t)
	defer ctrlMock.Finish()
	/* ---xxx--- */
	/* ---xxx--- */
	// mock repository level
	mockRepo := NewMockrepository(ctrlMock)
	/* ---xxx--- */

	type args struct {
		request memberships.SignUpRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				request: memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(nil)
			},
		},
		{
			name: "failed when GetUser",
			args: args{
				request: memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, assert.AnError)
			},
		},
		{
			name: "failed when CreateUser",
			args: args{
				request: memberships.SignUpRequest{
					Email:    "test@gmail.com",
					Username: "testusername",
					Password: "password",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mockRepo.EXPECT().GetUser(args.request.Email, args.request.Username, uint(0)).Return(nil, gorm.ErrRecordNotFound)
				mockRepo.EXPECT().CreateUser(gomock.Any()).Return(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			/* ---xxx--- */
			tt.mockFn(tt.args)
			/* ---xxx--- */

			s := &service{
				cfg:        &configs.Config{},
				repository: mockRepo,
			}
			if err := s.SignUp(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("service.SignUp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
