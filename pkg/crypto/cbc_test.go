package crypto

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	type args struct {
		orig string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"A", args{"123456"}, false},
		{"B", args{"z-(!@#$0?1%^&*)+A"}, false},
		{"C", args{"Test*AES@Encrypt&"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Encrypt() input: %s", tt.args.orig)
			got, err := Encrypt(tt.args.orig)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Encrypt() output = %s", got)

			orig, err := Decrypt(got)
			if err != nil {
				t.Errorf("Decrypt() error = %v", err)
			}
			t.Logf("Decrypt() output = %s", orig)
			if orig != tt.args.orig {
				t.Errorf("Decrypt() result not equal to origin")
			}
		})
	}
}

func TestDecrypt(t *testing.T) {
	type args struct {
		cryted string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"B", args{"HdbAgl/+HeWUq+jZx35Wq6Tc3fC97PK2+MdQ4w+MgLc="}, false},
		{"C", args{"v2a4/sfENbcY2YJHMYr31FsyIvRM1Nba2/AQqi/0yZc="}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("Decrypt() input: %s", tt.args.cryted)
			got, err := Decrypt(tt.args.cryted)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("Decrypt() output = %s", got)

			chiper, err := Encrypt(got)
			if err != nil {
				t.Errorf("Encrypt() error = %v", err)
			}
			t.Logf("Encrypt() output = %s", chiper)
			if chiper != tt.args.cryted {
				t.Errorf("Encrypt() result (%s) not equal to cryted (%s)", chiper, tt.args.cryted)
			}
		})
	}
}
