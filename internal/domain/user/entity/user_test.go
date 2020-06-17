package entity

import "testing"

func TestUser_CheckPwd(t *testing.T) {
	type fields struct {
		ID       int
		UserName string
		PassWord string
	}
	type args struct {
		pwd string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"check right", fields{1, "test1", "123456"},
			args{"123456"}, true},
		{"check error", fields{1, "test1", "123456"},
			args{"111111"}, false},
		{"empty pwd", fields{1, "test1", "123456"},
			args{""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				ID:       tt.fields.ID,
				UserName: tt.fields.UserName,
				PassWord: tt.fields.PassWord,
			}
			if got := u.CheckPwd(tt.args.pwd); got != tt.want {
				t.Errorf("User.CheckPwd() = %v, want %v", got, tt.want)
			}
		})
	}
}
