package storage

import (
	"sync"
	"testing"
	"time"
)

func TestInMemoryStorage_GetFullUrlByButty(t *testing.T) {
	type fields struct {
		Storage     map[string]Link
		FreeLinksId map[int64]bool
		mux         sync.RWMutex
	}

	var test fields

	test.Storage = map[string]Link{}
	test.Storage["test"] = Link{
		ShortUrl:    "test",
		Url:         "www.ya.ru",
		CreatedDate: time.Now(),
		Id:          0,
	}

	type args struct {
		buttyUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "Simple test",
			fields: test,
			args: args{
				"test",
			},
			want:    "www.ya.ru",
			wantErr: false,
		},
		{
			name:   "Not found test",
			fields: test,
			args: args{
				"test_not",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InMemoryStorage{
				Storage:     tt.fields.Storage,
				FreeLinksId: tt.fields.FreeLinksId,
				mux:         tt.fields.mux,
			}
			got, err := i.GetFullUrlByButty(tt.args.buttyUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFullUrlByButty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetFullUrlByButty() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryStorage_NewButtyUrl(t *testing.T) {
	type fields struct {
		Storage     map[string]Link
		FreeLinksId map[int64]bool
		mux         sync.RWMutex
	}
	var test fields

	test.Storage = map[string]Link{}
	test.FreeLinksId = map[int64]bool{}

	test.Storage["test"] = Link{
		ShortUrl:    "test",
		Url:         "www.ya.ru",
		CreatedDate: time.Now(),
		Id:          0,
	}

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name:   "Simple test",
			fields: test,
			args: args{
				"www.ya.ru",
			},
			want:    "0",
			wantErr: false,
		},
		{
			name:   "Not found test",
			fields: test,
			args: args{
				"www.ya.sru",
			},
			want:    "1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InMemoryStorage{
				Storage:     tt.fields.Storage,
				FreeLinksId: tt.fields.FreeLinksId,
				mux:         tt.fields.mux,
			}
			got, err := i.NewButtyUrl(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewButtyUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("NewButtyUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_idToUrl(t *testing.T) {
	type args struct {
		id int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "0",
			args: args{
				0,
			},
			want: "0",
		},
		{
			name: "10",
			args: args{
				10,
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := idToUrl(tt.args.id); got != tt.want {
				t.Errorf("idToUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
