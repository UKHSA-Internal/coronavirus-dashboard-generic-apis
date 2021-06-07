package insight

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"

	"github.com/caarlos0/env"
)

type OperationData struct {
	OperationId       string
	ParentId          string
	TraceParent       string
	CloudRoleName     string `env:"WEBSITE_SITE_NAME" envDefault:"generic-apis"`
	CloudRoleInstance string `env:"WEBSITE_INSTANCE_ID" envDefault:""`
}

var TraceParentPattern = regexp.MustCompile(
	`^[ \t]*(?P<version>[0-9a-f]{2})-(?P<operationId>[0-9a-f]{32})-(?P<parentId>[0-9a-f]{16})-(?P<traceFlag>[0-9a-f]{2})(-.*)?[ \t]*$`)

func GetCloudRoleName() string {

	var (
		cloudRoleName  string
		roleNameExists bool
	)

	if cloudRoleName, roleNameExists = os.LookupEnv("WEBSITE_SITE_NAME"); !roleNameExists {
		cloudRoleName = "generic-apis"
	}

	return cloudRoleName

} // GetCloudRoleName

func GetCloudRoleInstance() string {

	cloudRoleInstance, _ := os.LookupEnv("WEBSITE_INSTANCE_ID")

	return cloudRoleInstance

} // GetCloudRoleInstance

func GetOperationData(traceparent string) *OperationData {

	data := &OperationData{}
	if err := env.Parse(data); err != nil {
		panic(err)
	}

	data.TraceParent = traceparent
	match := TraceParentPattern.FindStringSubmatch(data.TraceParent)
	if len(match) == 0 {
		return generateTraceParent(data)
	}

	result := make(map[string]string)
	for idx, name := range TraceParentPattern.SubexpNames() {
		if idx > 0 && name != "" {
			result[name] = match[idx]
		}
	}

	var (
		value string
		ok    bool
	)

	if value, ok = result["operationId"]; ok {
		data.OperationId = value
	}

	if value, ok = result["parentId"]; ok {
		data.ParentId = value
	}

	return data

} // GetOperationData

func generateTraceParent(data *OperationData) *OperationData {

	data.OperationId = GenerateOperationId()
	data.ParentId = GenerateParentId()
	data.TraceParent = fmt.Sprintf("00-%s-%s-01", data.OperationId, data.ParentId)

	return data

} // generateTraceParent

func GenerateOperationId() string {
	return generateRandomBits(32)
} // GenerateOperationId

func GenerateParentId() string {
	return generateRandomBits(16)
} // GenerateParentId

func generateRandomBits(n int) string {

	bytes := make([]byte, n/2)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	return hex.EncodeToString(bytes)

} // generateRandomBits
