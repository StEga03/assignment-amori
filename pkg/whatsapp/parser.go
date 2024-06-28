package whatsapp

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/assignment-amori/pkg/utils"
)

func ParseString(str string) ([]Message, error) {
	trimmedStr := strings.TrimRight(str, "\r\n")
	lines := utils.RegexNewlines.Split(trimmedStr, -1)

	if len(lines) == 0 {
		return nil, nil
	}

	parsedRawMsgs, err := parseRawMessages(newRawMessages(lines))
	if err != nil {
		return nil, err
	}

	return newMessages(parsedRawMsgs)
}

func newRawMessages(lines []string) []rawMessage {
	var msgs []rawMessage

	for _, line := range lines {
		// If the line doesn't match the regex it's probably part of the previous
		// message or a "WhatsApp event"
		if !utils.RegexParserRegular.MatchString(line) {
			// If it doesn't match the first regex but still matches the system regex
			// it should be considered a "WhatsApp event" so it gets labeled "system"
			if utils.RegexParserSystem.MatchString(line) {
				msgs = append(msgs, rawMessage{system: true, body: line})
			} else if len(msgs) > 0 {
				// Else, it's part of the previous message and should be concatenated
				lastMessage := &msgs[len(msgs)-1]
				lastMessage.body += "\n" + line
			}
		} else {
			msgs = append(msgs, rawMessage{system: false, body: line})
		}
	}

	return msgs
}

func parseSystemMessageBody(body string) (parsedRawMessage, error) {
	matches := utils.RegexParserSystem.FindStringSubmatch(body)
	if len(matches) < 5 {
		return parsedRawMessage{}, errors.New("invalid system message body format")
	}

	var parsed parsedRawMessage
	parsed.date = matches[1]
	parsed.time = matches[2]
	parsed.body = matches[4]

	// Handle optional ampm field
	if matches[3] != "" {
		ampm := matches[3]
		parsed.ampm = &ampm
	}

	// Note: author is nil for system messages
	return parsed, nil
}

func parseRegularMessageBody(body string) (parsedRawMessage, error) {
	matches := utils.RegexParserRegular.FindStringSubmatch(body)
	if len(matches) < 6 {
		return parsedRawMessage{}, errors.New("invalid regular message body format")
	}

	var parsed parsedRawMessage
	parsed.date = matches[1]
	parsed.time = matches[2]
	parsed.author = &matches[4]
	parsed.body = matches[5]

	// Handle optional ampm field
	if matches[3] != "" {
		ampm := matches[3]
		parsed.ampm = &ampm
	}

	return parsed, nil
}

func parseRawMessages(msgs []rawMessage) ([]parsedRawMessage, error) {
	if len(msgs) == 0 {
		return nil, nil
	}

	parsedRawMessages := make([]parsedRawMessage, 0, len(msgs))

	for _, msg := range msgs {
		// If it's a system message another regex should be used to parse it
		if msg.system {
			parsed, err := parseSystemMessageBody(msg.body)
			if err != nil {
				return nil, err
			}

			parsedRawMessages = append(parsedRawMessages, parsed)
		} else {
			parsed, err := parseRegularMessageBody(msg.body)
			if err != nil {
				return nil, err
			}

			parsedRawMessages = append(parsedRawMessages, parsed)
		}
	}

	return parsedRawMessages, nil
}

func newMessages(msgs []parsedRawMessage) ([]Message, error) {
	dateMap := make(map[string]bool)
	for _, msg := range msgs {
		dateMap[msg.date] = true
	}

	numDates, err := extractNumericDates(msgs)
	if err != nil {
		return nil, err
	}

	dateOrder, err := utils.InferDateOrder(numDates)
	if err != nil {
		return nil, err
	}

	if dateOrder == utils.DateOrderUnknown {
		return nil, errors.New("unable to infer date format")
	}

	outputMsgs := make([]Message, 0, len(msgs))

	for _, msg := range msgs {
		var daysStr, monthsStr, yearsStr string

		dc, err := utils.OrderDateComponents(msg.date)
		if err != nil {
			return nil, err
		}

		if dateOrder == utils.DateOrderDaysFirst {
			daysStr, monthsStr, yearsStr = dc[0], dc[1], dc[2]
		} else {
			monthsStr, daysStr, yearsStr = dc[0], dc[1], dc[2]
		}

		yearsStr, monthsStr, daysStr = utils.NormalizeDate(yearsStr, monthsStr, daysStr)

		timeStr := msg.time

		if msg.ampm != nil {
			timeStr, err = utils.ConvertTime12To24(msg.time, utils.NormalizeAMPM(*msg.ampm))
			if err != nil {
				return nil, err
			}
		}

		timeStr = utils.NormalizeTime(timeStr)
		timeParts := utils.RegexSplitTime.Split(timeStr, -1)

		hoursStr, minutesStr, secondsStr := timeParts[0], timeParts[1], timeParts[2]

		year, _ := strconv.Atoi(yearsStr)
		month, _ := strconv.Atoi(monthsStr)
		day, _ := strconv.Atoi(daysStr)
		hour, _ := strconv.Atoi(hoursStr)
		minute, _ := strconv.Atoi(minutesStr)
		second, _ := strconv.Atoi(secondsStr)

		date := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.UTC)

		outputMsg := Message{
			Author:   msg.author,
			Body:     msg.body,
			Date:     date,
			BodyType: "text",
		}

		outputMsgs = append(outputMsgs, outputMsg)
	}

	return outputMsgs, nil
}

func extractNumericDates(msgs []parsedRawMessage) ([]utils.NumericDate, error) {
	dateMap := make(map[string]bool)
	for _, msg := range msgs {
		dateMap[msg.date] = true
	}

	var numDates []utils.NumericDate

	for date := range dateMap {
		dc, err := utils.OrderDateComponents(date)
		if err != nil {
			return nil, err
		}

		numDate, err := utils.DateComponentsToNumericDate(dc)
		if err != nil {
			return nil, err
		}

		numDates = append(numDates, numDate)
	}

	return numDates, nil
}
