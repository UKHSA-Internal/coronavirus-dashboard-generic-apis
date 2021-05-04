package insight

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"regexp"
)

type OperationData struct {
	OperationId string
	ParentId    string
	TraceParent string
}

var TraceParentPattern = regexp.MustCompile(
	`^[ \t]*(?P<version>[0-9a-f]{2})-(?P<operationId>[0-9a-f]{32})-(?P<parentId>[0-9a-f]{16})-(?P<traceFlag>[0-9a-f]{2})(-.*)?[ \t]*$`)

func getCloudRoleName() string {

	var (
		cloudRoleName  string
		roleNameExists bool
	)

	if cloudRoleName, roleNameExists = os.LookupEnv("WEBSITE_SITE_NAME"); !roleNameExists {
		cloudRoleName = "generic-apis"
	}

	return cloudRoleName

} // getCloudRoleName

func GetOperationData(traceparent string) *OperationData {

	match := TraceParentPattern.FindStringSubmatch(traceparent)

	if len(match) == 0 {
		return generateTraceParent()
	}

	response := &OperationData{TraceParent: traceparent}

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
		response.OperationId = value
	}

	if value, ok = result["parentId"]; ok {
		response.ParentId = value
	}

	return response

} // GetOperationData

func generateTraceParent() *OperationData {

	response := &OperationData{
		OperationId: GenerateOperationId(),
		ParentId:    GenerateParentId(),
	}

	response.TraceParent = fmt.Sprintf("00-%s-%s-00", response.OperationId, response.ParentId)

	return response

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
