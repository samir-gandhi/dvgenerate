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

type connectorDocData struct {
	ConnectorName string
	ConnectorId   string
	Properties    []connectorDocPropertyData
}

type connectorDocPropertyData struct {
	Name               string
	Type               *string
	Description        *string
	ConsoleDisplayName *string
}

func Generate() {

	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		panic(err)
	}
	fmt.Println("dir:", dir)

	// nameReg := regexp.MustCompile(`name=([A-Za-z0-9_]+),`)
	t, err := template.ParseFiles(dir + "/internal/templates/connector.tmpl")
	if err != nil {
		panic(err)
	}

	fileName := t.Name() + "_generated.md"
	// fileName := strings.ToLower(t.Name()) + "_generated.md"
	file, err := os.Create(fmt.Sprintf("generated/%s", fileName))
	if err != nil {
		panic(err)
	}

	conns, err := readConnectors()
	if err != nil {
		panic(err)
	}

	err = t.Execute(file, conns)
	if err != nil {
		panic(err)
	}
}

func readConnectors() (connectionByName, error) {
	c, err := testClient()
	if err != nil {
		return nil, err
	}
	environment_id := os.Getenv("PINGONE_ENVIRONMENT_ID")
	connectors, err := c.ReadConnectors(&environment_id, nil)
	if err != nil {
		return nil, err
	}

	connectorList := make(connectionByName, 0)

	for _, conn := range connectors {

		connectorProperties := make(connectionPropertyByName, 0)
		if acv := conn.AccountConfigView; acv != nil {
			for key, prop := range conn.Properties {

				for _, acv := range acv.Items {

					if acvProperty := acv.PropertyName; acvProperty != nil && *acvProperty == key {

						description := prop.Info

						if description != nil && strings.TrimSpace(*description) != "" && !strings.HasSuffix(strings.TrimSpace(*description), ".") {
							descriptionTemp := fmt.Sprintf("%s.", *description)
							description = &descriptionTemp
						}

						connectorProperties = append(connectorProperties, connectorDocPropertyData{
							Name:               key,
							Type:               prop.Type,
							Description:        description,
							ConsoleDisplayName: prop.DisplayName,
						})
					}
				}
			}
		}

		sort.Sort(connectorProperties)

		var connectorId, connectorName string

		if v := conn.ConnectorID; v != nil {
			connectorId = *v
		} else {
			connectorId = "No value"
		}

		if v := conn.Name; v != nil {
			connectorName = *v
		} else {
			connectorName = "No name"
		}

		connectorList = append(connectorList, connectorDocData{
			ConnectorName: connectorName,
			ConnectorId:   connectorId,
			Properties:    connectorProperties,
		})
	}

	sort.Sort(connectorList)

	return connectorList, nil
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

	return c, nil
}

type connectionPropertyByName []connectorDocPropertyData

func (a connectionPropertyByName) Len() int           { return len(a) }
func (a connectionPropertyByName) Less(i, j int) bool { return a[i].Name < a[j].Name }
func (a connectionPropertyByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type connectionByName []connectorDocData

func (a connectionByName) Len() int           { return len(a) }
func (a connectionByName) Less(i, j int) bool { return a[i].ConnectorName < a[j].ConnectorName }
func (a connectionByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
