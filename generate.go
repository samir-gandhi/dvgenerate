//go:generate go run ./cmd/generate/generate.go

package dvgenerate

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"unicode"

	dv "github.com/samir-gandhi/davinci-client-go/davinci"
	"github.com/samir-gandhi/dvgenerate/internal"
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
	Value              *string
}

func Generate() {
	GenerateReferenceTemplate()
	GenerateConnectorHCLExamples()
}

func GenerateReferenceTemplate() {

	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		panic(err)
	}
	fmt.Println("dir:", dir)
	var file *os.File
	packageDir := "./"
	currentDir := "./"
	fileName := currentDir + "templates/guides/connector-reference.md.tmpl"

	if !strings.Contains(dir, "samir-gandhi/dvgenerate") {
		file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}

		defer file.Close()

		packageDir = "./vendor/github.com/samir-gandhi/dvgenerate/"
	} else {
		file, err = os.Create(fileName)
		if err != nil {
			panic(err)
		}

		defer file.Close()
	}

	t, err := template.ParseFiles(packageDir + "cmd/generate/connector_reference.tmpl")
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

func GenerateConnectorHCLExamples() {

	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		panic(err)
	}
	fmt.Println("dir:", dir)

	conns, err := readConnectors()
	if err != nil {
		panic(err)
	}

	for _, conn := range conns {

		var file *os.File
		packageDir := "./"
		currentDir := "./"
		fileName := currentDir + fmt.Sprintf("examples/connectors/%s.tf", conn.ConnectorId)

		if !strings.Contains(dir, "samir-gandhi/dvgenerate") {
			file, err = os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0644)
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

		err = t.Execute(file, conn)
		if err != nil {
			panic(err)
		}
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

						var propertyType string
						if prop.Type != nil {
							switch *prop.Type {
							case "string", "boolean", "number", "":
								propertyType = *prop.Type
							default:
								propertyType = "json"
							}
						} else {
							propertyType = "string"
						}

						connectorProperty := connectorDocPropertyData{
							Name:               key,
							Type:               &propertyType,
							Description:        description,
							ConsoleDisplayName: prop.DisplayName,
						}

						exampleFound := false
						if v, ok := internal.ExampleValues[connectorId][key]; ok {
							connectorProperty.Value = &v.Value

							if v.OverridingType != nil {
								connectorProperty.Type = v.OverridingType
							}

							exampleFound = true
						}

						if !exampleFound {
							defaultValue := fmt.Sprintf("var.%s_property_%s", strings.ToLower(connectorId), camelToSnake(key))
							connectorProperty.Value = &defaultValue
						}

						connectorProperties = append(connectorProperties, connectorProperty)
					}
				}
			}
		}

		sort.Sort(connectorProperties)

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

func (a connectionByName) Len() int { return len(a) }
func (a connectionByName) Less(i, j int) bool {
	return fmt.Sprintf("%s%s", a[i].ConnectorName, a[i].ConnectorId) < fmt.Sprintf("%s%s", a[j].ConnectorName, a[j].ConnectorId)
}
func (a connectionByName) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func camelToSnake(camel string) string {
	// A buffer to build the output string
	var buf bytes.Buffer

	// Loop through each rune in the string
	for i, r := range camel {
		// If the rune is an uppercase letter and it's not the first character,
		// write an underscore to the buffer
		if unicode.IsUpper(r) && i > 0 {
			buf.WriteRune('_')
		}
		// Write the lowercase version of the current rune to the buffer
		buf.WriteRune(unicode.ToLower(r))
	}

	// Return the contents of the buffer as a string
	return buf.String()
}
