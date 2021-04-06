package service

import (
	"testing"

	"github.com/cloudandbuild/blog-article-resrouces/content/blog/article-5/interfaces"
	"github.com/cloudandbuild/blog-article-resrouces/content/blog/article-5/interfaces/interfacesmocks"
	"github.com/cloudandbuild/blog-article-resrouces/content/blog/article-5/testutil"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestAService_DoSomething(t *testing.T) {
	type fields struct {
		ARepository func(ctrl *gomock.Controller) interfaces.IRepository
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				ARepository: func(ctrl *gomock.Controller) interfaces.IRepository {
					ret := interfacesmocks.NewMockIRepository(ctrl)
					ret.EXPECT().Update(
						testutil.NewMatcher(
							&interfaces.Data{
								ID: 1,
							},
							func(a, b interface{}) bool {
								if diff := cmp.Diff(a, b, cmpopts.IgnoreFields(interfaces.Data{}, "UpdatedAt")); diff != "" {
									t.Errorf("MockSomeRepository.Update mismatch (-want +got): %s", diff)
									return false
								}
								return true
							},
						),
					).Return(nil).Times(1)
					return ret
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			r := &AService{
				ARepository: tt.fields.ARepository(mockCtrl),
			}
			if err := r.DoSomething(); (err != nil) != tt.wantErr {
				t.Errorf("AService.DoSomething() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
