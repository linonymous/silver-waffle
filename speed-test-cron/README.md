## speed-test-cron

This uses the speedtest-go library, tracks upload/download speeds of the network & logs into the csv file. 

### How to run?
- Build the binary
- Run using following command
``
  ./speed-test-cron <folder-path-to-store-speed_result.csv-file>
``
### Reason:
Had to write this script so as to prove to my provider that the speed is not stable & suffers through every 1-2 hours. 

### Sample File
Sample data looks as follows:
```
Timestamp,Latency,Download Speed,Upload Speed,Host Name,Host Country,Host Sponsor,Host Distance,Duration,Errors
"Tue, 29 Jun 2021 12:26:50 IST",36.765528ms,15.091031,19.024872,Shahapur,India,BSL Technologies Private Limited,367.31,63.03,
"Tue, 29 Jun 2021 19:50:50 IST",37.61775ms,16.240799,17.338961,Shirpur,India,Adarsh Info Solutions Pvt Ltd,141.00,48.40,
"Tue, 29 Jun 2021 20:11:08 IST",33.836121ms,14.904945,17.171342,Shirpur,India,Adarsh Info Solutions Pvt Ltd,141.00,66.77,
"Tue, 29 Jun 2021 20:21:17 IST",35.36054ms,4.787227,14.692575,Shirpur,India,Adarsh Info Solutions Pvt Ltd,141.00,75.91,
"Tue, 29 Jun 2021 20:41:14 IST",36.680484ms,8.553467,13.214196,Shirpur,India,Adarsh Info Solutions Pvt Ltd,141.00,72.95,
"Tue, 29 Jun 2021 20:52:03 IST",37.233918ms,6.228755,14.233460,Shirpur,India,Adarsh Info Solutions Pvt Ltd,141.00,121.40,
"Tue, 29 Jun 2021 21:11:04 IST",36.25867ms,15.860232,18.228863,Shirpur,India,Adarsh Info Solutions Pvt Ltd,141.00,62.59,
"Tue, 29 Jun 2021 21:21:07 IST",34.951556ms,15.740327,17.030194,Shirpur,India,Adarsh Info Solutions Pvt Ltd,141.00,64.77,
"Tue, 29 Jun 2021 21:41:07 IST",33.262597ms,10.357670,14.080928,Shirpur,India,Adarsh Info Solutions Pvt Ltd,141.00,65.36,
"Tue, 29 Jun 2021 21:51:15 IST",32.785677ms,7.491792,14.069995,Shirpur,India,Adarsh Info Solutions Pvt Ltd,141.00,74.21,
"Tue, 29 Jun 2021 22:11:26 IST",33.674698ms,10.565868,15.280297,Shirpur,India,Adarsh Info Solutions Pvt Ltd,141.00,84.73,
"Tue, 29 Jun 2021 22:21:11 IST",36.329649ms,10.485783,12.529499,Shirpur,India,Adarsh Info Solutions Pvt Ltd,141.00,69.42,
"Tue, 29 Jun 2021 22:41:31 IST",35.831694ms,10.002221,14.228358,Shirpur,India,Adarsh Info Solutions Pvt Ltd,141.00,90.07,
"Tue, 29 Jun 2021 22:51:25 IST",34.22809ms,4.316624,13.515855,Shirpur,India,Adarsh Info Solutions Pvt Ltd,141.00,83.60,
"Wed, 30 Jun 2021 08:11:04 IST",36.110136ms,15.550087,18.767351,Phaltan,India,TEJNET GIGA FIBER PVT LTD,534.58,63.48,
"Wed, 30 Jun 2021 08:21:34 IST",36.878321ms,9.232425,0.000000,Phaltan,India,TEJNET GIGA FIBER PVT LTD,534.58,92.07,"Post ""http://speedtest.tejnetgigafiber.com:8080/speedtest/upload.php"": write tcp 192.168.1.3:42892->103.208.74.232:8080: use of closed network connection -UPLOAD"
"Wed, 30 Jun 2021 08:41:28 IST",38.740608ms,10.422532,0.000000,Phaltan,India,TEJNET GIGA FIBER PVT LTD,534.58,85.90,"Post ""http://speedtest.tejnetgigafiber.com:8080/speedtest/upload.php"": write tcp 192.168.1.3:53686->103.208.74.232:8080: use of closed network connection -UPLOAD"
"Wed, 30 Jun 2021 08:51:04 IST",47.536894ms,10.524829,14.637678,YAVATMAL,India,TANISH BROADBAND,63.08,63.33,
"Wed, 30 Jun 2021 09:11:11 IST",47.064365ms,14.552661,16.131750,YAVATMAL,India,TANISH BROADBAND,63.08,69.87,
"Wed, 30 Jun 2021 09:21:11 IST",47.256258ms,14.571728,16.800730,YAVATMAL,India,TANISH BROADBAND,63.08,69.05,
"Wed, 30 Jun 2021 09:41:10 IST",39.60426ms,14.530669,17.251418,Phaltan,India,TEJNET GIGA FIBER PVT LTD,534.58,67.97,
"Wed, 30 Jun 2021 09:51:09 IST",39.534218ms,15.616614,15.468398,Phaltan,India,TEJNET GIGA FIBER PVT LTD,534.58,68.22,
"Wed, 30 Jun 2021 10:11:07 IST",172.274765ms,15.895890,17.078123,YAVATMAL,India,TANISH BROADBAND,63.08,65.41,
"Wed, 30 Jun 2021 10:21:47 IST",47.357139ms,14.700706,7.513082,YAVATMAL,India,TANISH BROADBAND,63.08,105.58,
"Wed, 30 Jun 2021 10:40:51 IST",46.922171ms,14.681806,17.781136,YAVATMAL,India,TANISH BROADBAND,63.08,49.53,
"Wed, 30 Jun 2021 10:51:28 IST",39.039672ms,9.156156,0.000000,Phaltan,India,TEJNET GIGA FIBER PVT LTD,534.58,85.47,"Post ""http://speedtest.tejnetgigafiber.com:8080/speedtest/upload.php"": write tcp 192.168.1.3:36536->103.208.74.232:8080: use of closed network connection -UPLOAD"
"Wed, 30 Jun 2021 11:10:54 IST",44.304186ms,14.785465,17.793232,YAVATMAL,India,TANISH BROADBAND,63.08,52.96,
"Wed, 30 Jun 2021 11:21:05 IST",44.258723ms,16.207243,17.507227,YAVATMAL,India,TANISH BROADBAND,63.08,63.26,
"Wed, 30 Jun 2021 11:41:07 IST",47.763789ms,15.149960,18.277743,YAVATMAL,India,TANISH BROADBAND,63.08,64.67,
"Wed, 30 Jun 2021 11:51:05 IST",45.076073ms,15.294689,18.263006,YAVATMAL,India,TANISH BROADBAND,63.08,64.00,
```
