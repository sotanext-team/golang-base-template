package schema

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	config "app-api/configs"

	"github.com/sirupsen/logrus"
	"github.com/sony/sonyflake"
)

var sf *sonyflake.Sonyflake

func init() {
	var st sonyflake.Settings
	if !config.IsDev() {
		st.MachineID = ECSMachineID
	}
	sf = sonyflake.NewSonyflake(st)
	if sf == nil {
		if st.MachineID != nil {
			_, err := st.MachineID()
			logrus.Error(err)
			panic("sonyflake unabled to get machine ID")
		}
		panic("sonyflake not created")
	}
}

func NextID() uint64 {
	id, err := sf.NextID()
	if err != nil {
		panic("NextID() returned error: " + err.Error())
	}
	return id
}

// ECSMachineID retrieves the private IP address of the AWS ECS container
// and returns its lower 16 bits.
func ECSMachineID() (uint16, error) {
	ip, err := awsECSPrivateIPv4()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}

func awsECSPrivateIPv4() (net.IP, error) {
	// ECS_CONTAINER_METADATA_URI_V4 is available in ECS container. Doc:
	// https://docs.aws.amazon.com/AmazonECS/latest/userguide/task-metadata-endpoint-v4-fargate.html
	res, err := http.Get(os.Getenv("ECS_CONTAINER_METADATA_URI_V4"))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	containerInfo := struct {
		Networks []struct {
			IPv4Addresses []string
		}
	}{}
	err = json.Unmarshal(body, &containerInfo)
	if err != nil {
		return nil, err
	}

	ip := net.ParseIP(containerInfo.Networks[0].IPv4Addresses[0])
	if ip == nil {
		return nil, errors.New("invalid ip address")
	}
	return ip.To4(), nil
}
