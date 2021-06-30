package main

import (
	http2 "cowin/http"
	"cowin/model"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func RetrieveBeneficiaryIds(client *http2.Client) map[string]string {
	beneficiary := client.GetBeneficiary()
	if beneficiary == nil {
		log.Println("Could not get the beneficiary list!")
		return nil
	}
	var result = make(map[string]string)
	for _, b := range beneficiary.Beneficiaries {
		result[b.BeneficiaryReferenceID] = b.Name
	}
	return result
}

func ChooseBeneficiaries(available map[string]string) []string {
	fmt.Printf("%50s", "following are the available beneficiaries:\n")
	count := 1
	var response = make(map[int]string)
	for key, value := range available {
		response[count] = key
		fmt.Printf("SID: %d %5s Name: %30s ID: %15s\n", count, " ", value, key)
		count++
	}
	var input = make(map[string]string)
	var inp int
	for {
		fmt.Println("Input the SIDs of beneficiaries to book slots for(input -1 to confirm SIDs input): ")
		_, err := fmt.Scanf("%d", &inp)
		if _, ok := response[inp]; err != nil || !ok {
			if inp == -1 {
				fmt.Println("Entries noted! Moving ahead...")
				break
			}
			fmt.Println("Invalid SID! Input again!")
		}
		input[response[inp]] = ""
	}
	var result = make([]string, 0)
	for k, _ := range input {
		result = append(result, k)
	}
	return result
}

func InputPinCodes() []string {
	// Add the pin-code of the place to refer the centers around
	pin := InputString("Pin Code (colon(:) separated for multiple pin codes): ")
	var pinCodes []string
	pins := strings.Split(pin, ":")
	for _, pin := range pins {
		if pin == "" || len(pin) != 6 {
			continue
		}
		pinCodes = append(pinCodes, pin)
	}
	return pinCodes
}

func InputCaptcha(h *http2.Client) string {
	// Get captcha & resolve it
	c := h.GetCaptcha()
	err := ioutil.WriteFile("captcha.svg", []byte(c.Captcha), 0777)
	if err != nil {
		fmt.Println("Could not save captcha! Rerun! ")
		return ""
	}
	fmt.Print("Open up the image and enter captcha:")
	var captcha string
	_, err = fmt.Scan(&captcha)
	if err != nil {
		fmt.Println("Invalid/unsupported captcha! Retry again!")
		return ""
	}
	captcha = strings.TrimSpace(captcha)
	return captcha
}

func InputString(variable string) string {
	var value string
	for {
		fmt.Printf("Provide %s: ", variable)
		_, err := fmt.Scan(&value)
		if err != nil || value == "" {
			fmt.Println("Invalid input! Please try again!")
			continue
		}
		break
	}
	return value
}

func main()  {
	// To book slots for older people change this ageLimit or change the code at LINE 41
	ageLimit := 44
	// beneficiary is the list of users' reference ids to book the slots for
	var beneficiary []string = []string{"56711904877400", "36341989442300"}
	// After OTP validation on covid, copy the auth token from `authorization` header inside the api call /beneficiary in developer console
	auth := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX25hbWUiOiIwZjgwNzhjNC0yN2JlLTQ1ZTEtODZlYy0xMzNjNGM3MjY3ZDciLCJ1c2VyX2lkIjoiMGY4MDc4YzQtMjdiZS00NWUxLTg2ZWMtMTMzYzRjNzI2N2Q3IiwidXNlcl90eXBlIjoiQkVORUZJQ0lBUlkiLCJtb2JpbGVfbnVtYmVyIjo4NDIxNTYzNTM5LCJiZW5lZmljaWFyeV9yZWZlcmVuY2VfaWQiOjU2NzExOTA0ODc3NDAwLCJzZWNyZXRfa2V5IjoiYjVjYWIxNjctNzk3Ny00ZGYxLTgwMjctYTYzYWExNDRmMDRlIiwidWEiOiJNb3ppbGxhLzUuMCAoWDExOyBMaW51eCB4ODZfNjQpIEFwcGxlV2ViS2l0LzUzNy4zNiAoS0hUTUwsIGxpa2UgR2Vja28pIENocm9tZS85MC4wLjQ0MzAuOTMgU2FmYXJpLzUzNy4zNiIsImRhdGVfbW9kaWZpZWQiOiIyMDIxLTA1LTExVDE0OjI5OjI0LjM0NFoiLCJpYXQiOjE2MjA3NDMzNjQsImV4cCI6MTYyMDc0NDI2NH0.PFDM7_6hrPJK6K3zSe8lfaABsBoI7js7qS0Dhza6tgM"

	// Get PinCodes
	pinCodes := []string{"412202", "411027", "411026"}
	if len(pinCodes) == 0 {
		fmt.Println("Invalid pin codes! Rerun!")
		return
	}

	// Initialize http2 Client
	h := http2.InitializeClient(http.DefaultClient, "https://cdn-api.co-vin.in/api/v2", auth)

	// Get Beneficiaries
	//b := RetrieveBeneficiaryIds(h)
	//beneficiary = ChooseBeneficiaries(b)
	if len(beneficiary) == 0 {
		fmt.Println("Could not get the beneficiaries! Retry again!")
		os.Exit(100)
	}

	// Get Captcha
	captcha := InputCaptcha(h)
	if captcha == "" {
		fmt.Println("Could not get the captcha!")
		os.Exit(200)
	}
	//captcha := ""
	count := 0
	start := time.Now()
	for {
		pin := pinCodes[count%len(pinCodes)]
		fmt.Println("Getting schedules for ", pin, "....")
		schedules := h.GetSchedules(pin)
		if schedules == nil {
			fmt.Println("Did not get any schedules! Sleeping for 2 seconds")
			time.Sleep(time.Second)
			continue
		}
		if len(schedules.Centers) == 0 {
			fmt.Printf("For %s pin, there are no centers available!\n", pin)
		}

		for _, center := range schedules.Centers {
			for _, session := range center.Sessions {
				if session.MinAgeLimit > ageLimit {
					continue
				}
				if session.AvailableCapacity > 0 {
					fmt.Println("Available session found!")
					fmt.Println(session.MinAgeLimit)
					fmt.Println(session.Slots)
					fmt.Println(session.Date)
					fmt.Println(session.AvailableCapacity)
					if len(session.Slots) == 0 {
						continue
					}
					body := model.ScheduleRequest{
						CenterID:      center.CenterID,
						SessionID:     session.SessionID,
						Beneficiaries: beneficiary,
						Slot:          session.Slots[0],
						Dose:          1,
						Captcha:       captcha,
					}
					fmt.Println("Making the booking...")
					response := h.BookSlot(&body)
					if response != nil && response.AppointmentConfirmationNo != "" {
						fmt.Println(count)
						fmt.Println(time.Since(start))
						fmt.Println("SLOTS BOOKED FOR ", session.Slots[0], " ON ", session.Date)
						fmt.Println("APPOINTMENT NO ", response.AppointmentConfirmationNo)
						os.Exit(100)
					} else {
						fmt.Println("could not book the slots! Retrying...")
					}
				}
			}
		}

		time.Sleep(800 * time.Millisecond)
		count++
		if count % 70 == 0 {
			fmt.Println("Time since start: ", time.Since(start))
		}
	}
}