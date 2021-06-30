package main

import (
	"encoding/csv"
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
	"log"
	"os"
	"strings"
	"time"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Please provide folder path to store results ex. /root/cmd/")
	}
	folder := args[1]
	if !strings.HasSuffix(folder, "/") {
		folder = folder + "/"
	}
	fileName := fmt.Sprintf("%sspeed_result.csv", folder)

	var row [][]string
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		// if file not exist add a header row
		row = append(row, []string{"Timestamp", "Latency", "Download Speed", "Upload Speed", "Host Name", "Host Country", "Host Sponsor", "Host Distance", "Duration", "Errors"})
	}

	// Open a file
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	// SpeedTest code
	user, _ := speedtest.FetchUserInfo()
	serverList, _ := speedtest.FetchServerList(user)
	targets, _ := serverList.FindServer([]int{})
	for _, s := range targets {
		errors := ""
		start := time.Now()
		err := s.PingTest()
		if err != nil {
			errors += err.Error() + " -PING ## "
		}
		err = s.DownloadTest(false)
		if err != nil {
			errors += err.Error() + " -DOWNLOAD ## "
		}
		err = s.UploadTest(false)
		if err != nil {
			errors += err.Error() + " -UPLOAD"
		}
		elapsed := time.Since(start)
		timestamp := time.Now().Format(time.RFC1123)
		latency := fmt.Sprintf("%s", s.Latency)
		downloadSpeed := fmt.Sprintf("%f", s.DLSpeed)
		uploadSpeed := fmt.Sprintf("%f", s.ULSpeed)

		fmt.Printf("Latency: %s, Download: %f, Upload: %f, TimeElapsed: %.2f, Distance: %.2f, Server name: %s, Server Country: %s, Errors: %s \n", s.Latency, s.DLSpeed, s.ULSpeed, elapsed.Seconds(), s.Distance, s.Name, s.Country, errors)
		data := []string{timestamp, latency, downloadSpeed, uploadSpeed, s.Name, s.Country, s.Sponsor, fmt.Sprintf("%.2f", s.Distance), fmt.Sprintf("%.2f", elapsed.Seconds()), errors}
		row = append(row, data)
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for _, r := range row {
		err := writer.Write(r)
		if err != nil {
			fmt.Println("Could not write to csv file ", err)
		}
	}
}