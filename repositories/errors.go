package repositories

import (
	"fmt"
	"github.com/lib/pq"
	"regexp"
)

type ConflictError struct {
	Message string
	Err     error
}

func (e *ConflictError) Error() string {
	return e.Message
}

const PQUniqueViolation = "23505"

// parsePsqlError takes a pq.Error and returns a matching custom error
// or the error itself if no matching custom error exists
func parsePsqlError(e *pq.Error) error {
	switch e.Code {
	case PQUniqueViolation:
		column, value := extractColumnValue(e.Detail)
		msg := fmt.Sprintf("[%s] already exists with this value (%s)", column, value)

		return &ConflictError{
			Message: msg,
			Err:     e,
		}
	default:
		return e
	}
}

// parseError take an error and passes it through the respective database error parser
// or returns the error itself if there is no matching parser
func parseError(e error) error {
	pqErr, ok := e.(*pq.Error)
	if ok {
		return parsePsqlError(pqErr)
	}

	return e
}

// extractColumnValue takes a string in the form of a sql error detail
// and returns the contained column and value
func extractColumnValue(detail string) (string, string) {
	var columnFinder = regexp.MustCompile(`Key \((.+)\)=`)
	var valueFinder = regexp.MustCompile(`Key \(.+\)=\((.+)\)`)

	column := extractStringSubmatch(columnFinder, detail)
	value := extractStringSubmatch(valueFinder, detail)

	return column, value
}

// extractString takes a regex and a string and returns
// the matched string or an empty string if no match exists
func extractStringSubmatch(regex *regexp.Regexp, str string) string {
	results := regex.FindStringSubmatch(str)
	if len(results) < 2 {
		return ""
	}
	return results[1]
}
