package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * timeConversion converts a 12-hour AM/PM time format (e.g., "07:05:45PM")
 * to a 24-hour military time format (e.g., "19:05:45").
 *
 * The function takes a string `s` representing the time in 12-hour format as input.
 * It then returns the corresponding time in 24-hour format as a string.
 */
func timeConversion(s string) {
	// Determine if the time is in PM
	var (
		is_pm bool = strings.ToLower(s[len(s)-2:]) == "pm"
		// Split the time string (excluding AM/PM) into its components (hours, minutes, seconds)
		tl []string = strings.Split(s[:len(s)-2], ":")
	)

	// Convert the hour component from string to integer
	h, e := strconv.ParseInt(tl[0], 10, 64)

	// Check for any errors during the conversion process
	if e != nil {
		log.Fatal(e)
	}

	// If the time is in PM and the hour is less than 12, convert it to 24-hour format by adding 12
	if is_pm && h < 12 {
		tl[0] = strconv.Itoa(int(h + 12))
	}

	// If the time is in AM and the hour is 12, convert it to "00" for midnight
	if !is_pm && h == 12 {
		tl[0] = "00"
	}

	// Join the components back into a single time string in 24-hour format and print it
	fmt.Println(strings.Join(tl, ":"))
}

func main() {
	// Prompt the user to enter a time in 12-hour format
	fmt.Println("Enter input time")
	fmt.Println("example : 07:05:45PM")

	// Create a new buffered reader to read input from the standard input
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	// Read the input time from the user
	s := readLine(reader)

	// Call the timeConversion function to convert and display the time in 24-hour format
	timeConversion(s)
}

/*
 * readLine reads a line of input from the provided buffered reader.
 * It trims any trailing newline characters and returns the line as a string.
 */
func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}
