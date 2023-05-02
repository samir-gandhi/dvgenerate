//go:generate go run ./cmd/generate/generate.go

package dvgenerate

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	dv "github.com/samir-gandhi/davinci-client-go/davinci"

	// "strings"
	"text/template"
)

func Generate() {

	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		panic(err)
	}
	fmt.Println("dir:", dir)
	var file *os.File
	packageDir := "./"
	currentDir := "./"
	fileName := currentDir + "docs/resources/connection.md"

	if !strings.Contains(dir, "samir-gandhi/dvgenerate") {
		file, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			panic(err)
		}
		packageDir = "./vendor/github.com/samir-gandhi/dvgenerate/"
	} else {
		file, err = os.Create(fileName)
		if err != nil {
			panic(err)
		}
	}

	t, err := template.ParseFiles(packageDir + "cmd/generate/connector.tmpl")
	if err != nil {
		panic(err)
	}

	conns, err := readConnectors()
	if err != nil {
		panic(err)
	}

	err = t.ExecuteTemplate(file, "allconnections", conns)
	if err != nil {
		panic(err)
	}
}

func readConnectors() ([]dv.Connector, error) {
	c, err := testClient()
	if err != nil {
		return nil, err
	}
	connectors, err := c.ReadConnectors(&c.CompanyID, nil)
	if err != nil {
		return nil, err
	}
	sort.Sort(ConnByName(connectors))
	return connectors, nil
}

func testClient() (*dv.APIClient, error) {
	e := ""
	username := os.Getenv("PINGONE_USERNAME")
	if username == "" {
		e = e + "PINGONE_USERNAME "
	}
	password := os.Getenv("PINGONE_PASSWORD")
	if password == "" {
		e = e + "PINGONE_PASSWORD "
	}
	region := os.Getenv("PINGONE_REGION")
	if region == "" {
		e = e + "PINGONE_REGION "
	}
	environment_id := os.Getenv("PINGONE_ENVIRONMENT_ID")
	if environment_id == "" {
		e = e + "PINGONE_ENVIRONMENT_ID "
	}
	if e != "" {
		return nil, fmt.Errorf("missing environment variables: %s", e)
	}

	cInput := dv.ClientInput{
		Username:        username,
		Password:        password,
		PingOneRegion:   region,
		PingOneSSOEnvId: environment_id,
	}
	c, err := dv.NewClient(&cInput)
	if err != nil {
		return nil, err
	}
	if environment_id != "" {
		c.CompanyID = environment_id
	}

	return c, nil
}

type ConnByName []dv.Connector

func (a ConnByName) Len() int           { return len(a) }
func (a ConnByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a ConnByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
