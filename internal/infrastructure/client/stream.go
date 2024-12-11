package client

import (
	"bufio"
	"io"
	"net/http"
)

type StreamReader struct {
	isFinished bool

	reader   *bufio.Reader
	response *http.Response
}

func (stream *StreamReader) Recv() (*http.Response, error) {
	var err error
	if stream.isFinished {
		err = io.EOF
		return nil, err

	}
	return stream.response, nil
}

func (stream *StreamReader) ProcessLine() ([]byte, error) {
	response, err := stream.Recv()
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (stream *StreamReader) processLines() {

}

func (stream *StreamReader) Close() {
	stream.response.Body.Close()
}
