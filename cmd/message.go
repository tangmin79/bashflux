/**
 * Copyright (c) 2016 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package cmd

import (
	"strings"
)

// GetMsg - gets messages from the channel
func GetMsg(id string, startTime string, endTime string) string {
	url := UrlHTTP + "/channels/" + id +
	                 "/msg?start_time=" + startTime + "&end_time=" + endTime
	resp, err := netClient.Get(url)
	body := GetHttpRespBody(resp, err)

	return body
}

// SendMsg - publishes SenML message on the channel
func SendMsg(id string, msg string) string {
	url := UrlHTTP + "/channels/" + id + "/msg"
	sr := strings.NewReader(msg)
	resp, err := netClient.Post(url, "application/json", sr)
	body := GetHttpRespBody(resp, err)

	return body
}
