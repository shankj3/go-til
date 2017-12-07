package deserialize

import (
	"testing"
	"io/ioutil"
	pb "github.com/shankj3/ocelot/protos/out"
	"bytes"
	"bitbucket.org/level11consulting/go-til/test"
)

const TestOcelot = "test-fixtures/ocelot.yml"
const TestRepos = "test-fixtures/repo.json"

func TestDeserializer_YAMLToStruct(t *testing.T) {
	testOcelot, _ := ioutil.ReadFile(TestOcelot)
	d := New()
	ocelot := &pb.BuildConfig{}
	d.YAMLToStruct(testOcelot, ocelot)

	if ocelot.Image != "test" {
		t.Error(test.StrFormatErrors("ocelot image", "test", ocelot.Image))
	}
	if len(ocelot.Packages) != 2 {
		t.Error(test.IntFormatErrors("docker package list size", 2, len(ocelot.Packages)))
	}
	if ocelot.Env["BUILD_DEBUG"] != "1" {
		t.Error(test.StrFormatErrors("build debug value in global env", "1", ocelot.Env["BUILD_DEBUG"]))
	}
	if ocelot.Env["SEARCH_URL"] != "https://google.com" {
		t.Error(test.StrFormatErrors("search url value in global env", "https://google.com", ocelot.Env["SEARCH_URL"]))
	}
	if ocelot.Before.Env != nil {
		t.Error(test.GenericStrFormatErrors("before stages environment", nil, ocelot.Before.Env))
	}
	if ocelot.Before.Script[0] != "sh echo \"hello\"" {
		t.Error(test.StrFormatErrors("before stages first script", "sh echo \"hello\"", ocelot.Before.Script[0]))
	}
	if ocelot.After.Env["BUILD_DEBUG"] != "0" {
		t.Error(test.StrFormatErrors("after stages BUILD_DEBUG val", "0", ocelot.After.Env["BUILD_DEBUG"]))
	}
	//can we assume parsing looks good if the above values have been set or do I have to write it for all the fields
}

func TestDeserializer_JSONToProto(t *testing.T) {
	repositories := &pb.PaginatedRepository{}
	testRepo, _ := ioutil.ReadFile(TestRepos)
	d := New()
	d.JSONToProto(ioutil.NopCloser(bytes.NewReader(testRepo)), repositories)

	if repositories.Pagelen != 10 {
		t.Error(test.IntFormatErrors("repository page len", 10, int(repositories.Pagelen)))
	}
	if repositories.Page != 1 {
		t.Error(test.IntFormatErrors("repository current page", 1, int(repositories.Page)))
	}
	if repositories.Size != 1 {
		t.Error(test.IntFormatErrors("repository response size", 1, int(repositories.Size)))
	}
	if len(repositories.Values) != 1 {
		t.Error(test.IntFormatErrors("repository values length", 1, len(repositories.Values)))
	}
	if repositories.Values[0].Name != "test-ocelot" {
		t.Error(test.StrFormatErrors("repository name", "test-ocelot", repositories.Values[0].Name))
	}
	if repositories.Values[0].FullName != "mariannefeng/test-ocelot" {
		t.Error(test.StrFormatErrors("repository full name", "mariannefeng/test-ocelot", repositories.Values[0].FullName))
	}
	if repositories.Values[0].Type != "repository" {
		t.Error(test.StrFormatErrors("repository type", "repository", repositories.Values[0].Type))
	}
	if repositories.Values[0].Links.Hooks.Href != "https://api.bitbucket.org/2.0/repositories/mariannefeng/test-ocelot/hooks" {
		t.Error(test.StrFormatErrors("webhook", "https://api.bitbucket.org/2.0/repositories/mariannefeng/test-ocelot/hooks", repositories.Values[0].Links.Hooks.Href))
	}
}