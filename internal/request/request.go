package request

import (
	"fmt"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

// Add a vallidation as only HTTP/1.1 is allowed
// Add a validation to check only method is capital letters
func (r *RequestLine) ValidHeader() bool {
	if r.HttpVersion != "1.1" {
		return false
	}

	for _, c := range r.Method {
		if c < 'A' || c > 'Z' {
			return false
		}
	}

	return true
}

var ERROR_BAD_HEADER = fmt.Errorf("bad http Header")
var ERROR_BAD_START_LINE = fmt.Errorf("bad start line")
var LINE_SEPARATOR = "\r\n" // The registered nurse (CRLF)

func RequestFromReader(reader io.Reader) (Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return Request{}, fmt.Errorf("bad reading data from reader: %w", err)
	}

	requestLine, _, err := parseRequestLine(string(data))
	if err != nil {
		return Request{}, fmt.Errorf("bad parsing request line: %w", err)
	}

	return Request{*requestLine}, nil
}

/*
Parse the request line from a string.
Return the request line, the rest of the string, and an error if any.
*/
func parseRequestLine(line string) (*RequestLine, string, error) {
	index := strings.Index(line, LINE_SEPARATOR)

	//No line separator found, just return nil, line and nil. No complications, already life is complicated
	if index == -1 {
		return nil, line, nil
	}
	startLine := line[:index]
	restOfMessage := line[index+len(LINE_SEPARATOR):] // Just return that No need to parse this

	//fetch the 3 parts of the start line.
	startLineParts := strings.Split(startLine, " ")
	if len(startLineParts) != 3 { //Error out if not 3 parts.
		return nil, restOfMessage, ERROR_BAD_START_LINE
	}

	httpVersionParts := strings.Split(startLineParts[2], "/")
	if len(httpVersionParts) != 2 {
		return nil, restOfMessage, ERROR_BAD_START_LINE
	}

	requestLine := RequestLine{
		Method:        startLineParts[0],
		RequestTarget: startLineParts[1],
		HttpVersion:   httpVersionParts[1],
	}

	if !requestLine.ValidHeader() {
		return nil, restOfMessage, ERROR_BAD_HEADER
	}

	return &requestLine, restOfMessage, nil
}
