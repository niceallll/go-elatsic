package aliyunopsfile

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

type AliyunBody struct {
	Summary   string `json:"summary"`
	Priority  string `json:"priority"`
	Name      string `json:"name"`
	Timestamp string `json:"timestamp"`
	Source    string `json:"source"`
}

func Openfile(sa, aliname string) {
	aliyunSummary := sa
	aliyunName := aliname
	aliyunToken := "https://public-alert.aliyuncs.com/event/standard/xxxxxxxxxx"
	aliyunPriority := "P3"

	aliyunbody := createAliyunbody(aliyunSummary, aliyunName, aliyunPriority)
	sendAliyunAlert(aliyunbody, aliyunToken)

}

func createAliyunbody(summary, name, priority string) AliyunBody {
	hostName, _ := os.Hostname()
	hostIP := getLocalIP()
	aliyunSource := "主机：" + hostName + "  IP:" + hostIP
	aliyunTimestamp := time.Now().Format("2006-01-02 15:04:05")
	return AliyunBody{
		Summary:   summary,
		Priority:  priority,
		Name:      name,
		Timestamp: aliyunTimestamp,
		Source:    aliyunSource,
	}
}

func sendAliyunAlert(aliyunBody AliyunBody, token string) {
	aliyunURL := token
	aliyunBodyJSON, _ := json.Marshal(aliyunBody)

	fmt.Println(string(aliyunBodyJSON))

	req, _ := http.NewRequest("POST", aliyunURL, bytes.NewBuffer(aliyunBodyJSON))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Request sent successfully.")
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}

	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}

	return ""
}
