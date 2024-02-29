package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	url                 = "http://localhost:8080/"
	noSpecifiedWaitTime = "a while"
	maxAllowedRetries   = 3
	defaultWaitTime     = 3
)

var (
	ErrNoWaitTime      = errors.New("no wait time provided: cannot get weather at the moment")
	ErrWaitTimeTooLong = errors.New("wait time too long: cannot get weather at the moment")
)

func main() {
	err := getWeather()
	if err != nil {
		logToStderr(err)
		os.Exit(1)
	}
}

func makeRequest() (*http.Response, error) {
	// create request
	logToStderr("setting up request...")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		err = fmt.Errorf("failed to create request: %v", err)
		logToStderr(err)
	}

	// make request
	logToStderr("making request...")
	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}

func getWeather() error {
	numberOfRetries := 0
	for {

		resp, err := makeRequest()
		// handle response or error
		if err != nil {
			msg := fmt.Sprintf("failed to make request: %v", err)
			logToStderr(msg)
			return err
		}
		if resp.StatusCode == 429 {
			// seconds given in the retry header
			secondsToWait := resp.Header["Retry-After"][0]
			if len(secondsToWait) == 1 {
				secondsAsInt, _ := strconv.Atoi(secondsToWait)
				if secondsAsInt < 5 {
					msg := fmt.Sprintf("Retrying request after %d seconds", secondsAsInt)
					logToStderr(msg)
					time.Sleep(time.Duration(secondsAsInt) * time.Second)
				} else {
					return ErrWaitTimeTooLong
				}
			} else if secondsToWait == noSpecifiedWaitTime {
				// Here are my thoughts about this:
				// So firstly, when no wait time is specified,
				// we delay by 3 seconds and retry the request.
				// But we only retry the request if we have not
				// exceeded the maxNumberOfRequests

				// My reasons for this are as follows:
				// 1. We have to be a bit more optimistic.
				// We have to give the server to request.
				// 2. Also, we have to be a bit defensive.
				// We can't keep trying the request after getting
				// no wait time repeatedly.
				numberOfRetries++
				if numberOfRetries > maxAllowedRetries {
					// no seconds or time stamp given in retry header
					msg := "exceeded max number of retries"
					logToStderr(msg)
					return ErrNoWaitTime
				}
				msg := fmt.Sprintf("no wait time provided: retrying after %v seconds", defaultWaitTime)
				logToStderr(msg)
				time.Sleep(defaultWaitTime * time.Second)
			} else {
				// time stamp given in retry header
				t, _ := time.Parse(time.RFC1123, secondsToWait)
				td := time.Until(t)
				if td.Seconds() < 5 {
					msg := fmt.Sprintf("Retrying request after %v seconds", td.Seconds())
					logToStderr(msg)
					time.Sleep(td)
				} else {
					return ErrWaitTimeTooLong
				}

			}

		}
		if resp.StatusCode == 200 {
			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				msg := fmt.Sprintf("failed to read response body: %v", err)
				logToStderr(msg)
			}
			logToStdout(string(respBody))
			return nil
		}

	}

}

func logToStderr(message any) {
	fmt.Fprint(os.Stderr, message, "\n")
}

func logToStdout(message any) {
	fmt.Fprint(os.Stdout, message, "\n")
}
