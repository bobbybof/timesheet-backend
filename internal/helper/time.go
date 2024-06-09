package helper

import "time"

func TimeFormat(timeP string) (time.Time, error) {
	layout := "15:04"
	t, err := time.Parse(layout, timeP)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func DateFormat(date string) (time.Time, error) {
	layout := "2006-01-02 15:04"
	t, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}
