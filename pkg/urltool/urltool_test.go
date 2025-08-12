package urltool

import "testing"

func TestGetBasePath(t *testing.T) {
	type args struct {
		targetUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "base",
			args:    args{targetUrl: "https://www.baidu.com/test"},
			want:    "test",
			wantErr: false,
		},
		{
			name:    "invalid url",
			args:    args{targetUrl: "/xxxx/231"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "empty",
			args:    args{targetUrl: ""},
			want:    "",
			wantErr: true,
		},
		{
			name:    "with query",
			args:    args{targetUrl: "https://www.baidu.com/test?query=1"},
			want:    "test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBasePath(tt.args.targetUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBasePath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBasePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
