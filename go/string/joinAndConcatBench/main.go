package main

import (
	"bytes"
	"fmt"
	"strings"
)

func sprintfAndStringConcat() {
	var msg string
	msg += fmt.Sprintf("userid : %d", 1)
	msg += fmt.Sprintf("location: %s", "ab")
}

func sprintfAndSliceAppendAndStringJoin() {
	var arr []string
	var msg string

	arr = append(arr, fmt.Sprintf("userid : %d", 1))
	arr = append(arr, fmt.Sprintf("location: %s", "ab"))
	msg += strings.Join(arr[:], "")
}

func sprintfAndPreallocatedArrayAndJoin() {
	var arr [2]string
	var msg string

	arr[0] = fmt.Sprintf("userid : %d", 1)
	arr[1] = fmt.Sprintf("location: %s", "ab")
	msg += strings.Join(arr[:], "")
}

func sprintfAndPreallocatedSliceAndJoin() {
	var arr = make([]string, 0, 10)
	var msg string

	arr = append(arr, fmt.Sprintf("userid : %d", 1))
	arr = append(arr, fmt.Sprintf("location: %s", "ab"))
	msg += strings.Join(arr[:], "")
}

func noSprintfAndSliceAppend() {
	var arr []string
	var msg string

	arr = append(arr, "userid :"+"1")
	arr = append(arr, "location: "+"ab")
	msg += strings.Join(arr, "")
}

func noSprintfAndPreallocatedSlice() {
	var arr = make([]string, 0, 10)
	var msg string

	arr = append(arr, "userid :"+"1")
	arr = append(arr, "location: "+"ab")
	msg += strings.Join(arr, "")
}

func bytesBufferAppend() {
	var msg bytes.Buffer
	msg.WriteString("userid : " + "1")
	msg.WriteString("location : " + "ab")
	//msg.String()
}

func bytesBufferAppendSprintf() {
	var msg bytes.Buffer
	msg.WriteString(fmt.Sprintf("userid : %d", 1))
	msg.WriteString(fmt.Sprintf("location : %s", "ab"))
	//msg.String()
}
