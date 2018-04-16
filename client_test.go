package loach

import (
	"crypto/x509/pkix"
	"testing"

	"github.com/learnergo/loach/model"
)

const (
	Path string = "static\\file.yaml"

	Admin_Key  string = "static\\admin.key"
	Admin_Cert string = "static\\admin.crt"
)

func Test_NewClient(t *testing.T) {
	t.Log("Start test")
	_, err := NewClient(Path)
	if err != nil {
		t.Error("Failed to init client") // 如果不是如预期的那么就报错
	} else {
		t.Error("Succeeded to init client")
	}
	t.Log("End test")
}

func Test_LoadAdmin(t *testing.T) {
	t.Log("Load test")
	_, err := LoadIdentity(Admin_Key, Admin_Cert)
	if err != nil {
		t.Error("Failed to load admin") // 如果不是如预期的那么就报错
	} else {
		t.Log("Succeeded to load admin")
	}
	t.Log("Load test")
}

func Test_Register(t *testing.T) {
	t.Log("Register test")
	client, err := NewClient(Path)
	if err != nil {
		t.Error("Failed to init client") // 如果不是如预期的那么就报错
	} else {
		t.Log("Succeeded to init client")
	}

	admin, err := LoadIdentity(Admin_Key, Admin_Cert)
	if err != nil {
		t.Error("Failed to load admin") // 如果不是如预期的那么就报错
	} else {
		t.Log("Succeeded to load admin")
	}

	t.Log("Build request params")

	var attrs []model.RegisterAttribute
	attrs = append(attrs, model.RegisterAttribute{Name: "hf.Registrar.Roles", Value: "peer"})
	attrs = append(attrs, model.RegisterAttribute{Name: "hf.Revoker", Value: "false"})

	request := &model.RegisterRequest{
		EnrollID:       "peer0.org1.example.com",
		Type:           "peer",
		Secret:         "adminpwd",
		MaxEnrollments: -1,
		Affiliation:    "org1.department1",
		Attrs:          attrs,
	}

	t.Log("Start Request")
	response, err := client.Register(admin, request)
	if err != nil {
		t.Errorf("Failed Register,err=%s", err)
	} else {
		if !response.Success {
			t.Errorf("Failed Register,err=%s", response.Error())
		}
		if len(response.Errors) > 0 {
			t.Errorf("Failed Register,err=%s", response.Error())
		}
		result := response.Result.Secret
		t.Logf("Register success,password=%s", result)
	}

	t.Log("Register test")
}

func Test_Enroll(t *testing.T) {
	t.Log("Enroll test")
	client, err := NewClient(Path)
	if err != nil {
		t.Error("Failed to init client") // 如果不是如预期的那么就报错
	} else {
		t.Log("Succeeded to init client")
	}

	t.Log("Build request params")

	request := &model.EnrollRequest{
		EnrollID: "peer0.org1.example.com",
		Secret:   "adminpwd",
		Profile:  "tls",
		Hosts:    []string{"hynet"},
		Name: pkix.Name{
			Country:            []string{"CN"},
			Organization:       []string{"BeiJing"},
			OrganizationalUnit: []string{"BeiJing"},
			CommonName:         "peer0.org1.example.com",
		},
	}
	response, err := client.Enroll(request)
	if err != nil {
		t.Errorf("Failed Enroll,err=%s", err)
	} else {
		strIdentity, _ := model.MarshalIdentity(response.Identity)
		certChain := model.CertToString(response.CertChain)
		t.Logf("Enroll success,strIdentity=%s\ncertChain=%s", strIdentity, certChain)
	}
	t.Log("Enroll test")
}
