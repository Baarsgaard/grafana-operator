package v1beta1

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestGrafanaFolder_GetTitle(t *testing.T) {
	tests := []struct {
		name string
		cr   GrafanaFolder
		want string
	}{
		{
			name: "No custom title",
			cr: GrafanaFolder{
				ObjectMeta: metav1.ObjectMeta{Name: "cr-name"},
			},
			want: "cr-name",
		},
		{
			name: "Custom title",
			cr: GrafanaFolder{
				ObjectMeta: metav1.ObjectMeta{Name: "cr-name"},
				Spec: GrafanaFolderSpec{
					Title: "custom-title",
				},
			},
			want: "custom-title",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cr.GetTitle()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGrafanaFolder_GetUID(t *testing.T) {
	tests := []struct {
		name string
		cr   GrafanaFolder
		want string
	}{
		{
			name: "No custom UID",
			cr: GrafanaFolder{
				ObjectMeta: metav1.ObjectMeta{UID: "92fd2e0a-ad63-4fcf-9890-68a527cbd674"},
			},
			want: "92fd2e0a-ad63-4fcf-9890-68a527cbd674",
		},
		{
			name: "Custom UID",
			cr: GrafanaFolder{
				ObjectMeta: metav1.ObjectMeta{UID: "92fd2e0a-ad63-4fcf-9890-68a527cbd674"},
				Spec: GrafanaFolderSpec{
					CustomUID: "custom-uid",
				},
			},
			want: "custom-uid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cr.CustomUIDOrUID()
			assert.Equal(t, tt.want, got)
		})
	}
}
