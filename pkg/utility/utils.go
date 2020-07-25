package utility

import (
	"bytes"
	"comet/pkg/constants"
	"comet/pkg/dto"
	"html/template"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/dongri/phonenumber"
)

func ValidateDeclarationSymptom(declaration *dto.Declaration) bool {
	if declaration.Blue == constants.HasBlue ||
		declaration.Breathe == constants.HasBreathDifficulty ||
		declaration.Chest == constants.HasChestPain ||
		declaration.Cough == constants.CoughWorse ||
		declaration.Cough == constants.StartCoughing ||
		declaration.Drowsy == constants.HasDrowsy ||
		declaration.Fever == constants.FeverWorse ||
		declaration.Fever == constants.StartFever ||
		declaration.Throat == constants.SoreThroatWorse ||
		declaration.Throat == constants.StartSoreThroat ||
		declaration.Loss == constants.HasLoss {
		return true
	}
	return false
}

func GetPatientIDAndDateFromSwabID(id string) (string, string) {
	ls := strings.Split(id, "_")
	if len(ls) == 1 {
		return ls[0], ""
	}
	if len(ls) == 0 {
		return "", ""
	}
	return ls[0], ls[1]
}

func GenerateSwabID(patientID string, date string) string {
	return patientID + "_" + date
}

func NormalizePhoneNumber(phoneNum string, countryCode string) string {
	code := countryCode

	number := phonenumber.Parse(phoneNum, code)
	if number != "" {
		return number
	}

	// default MY
	number = phonenumber.Parse(phoneNum, "MY")
	if number != "" {
		return number
	}

	country := phonenumber.GetISO3166ByNumber(phoneNum, true)
	number = phonenumber.ParseWithLandLine(phoneNum, country.CountryName)
	if number != "" {
		return number
	}

	country = phonenumber.GetISO3166ByNumber(phoneNum, false)
	number = phonenumber.ParseWithLandLine(phoneNum, country.CountryName)
	if number != "" {
		return number
	}

	country = phonenumber.GetISO3166ByNumber("+"+phoneNum, true)
	number = phonenumber.ParseWithLandLine("+"+phoneNum, country.CountryName)
	if number != "" {
		return number
	}

	country = phonenumber.GetISO3166ByNumber("+"+phoneNum, false)
	number = phonenumber.ParseWithLandLine("+"+phoneNum, country.CountryName)
	if number != "" {
		return number
	}

	return number
}

func NormalizeID(id string) string {
	re := regexp.MustCompile(`[^0-9a-zA-Z]`)
	return strings.ToUpper(re.ReplaceAllString(id, ""))
}

func NormalizeDate(date string) (string, error) {
	re := regexp.MustCompile(`[^0-9]`)
	d := strings.ToUpper(re.ReplaceAllString(date, ""))
	if len(d) < 8 {
		return "", constants.InvalidDateError
	}
	return d, nil
}

func NormalizeRole(role string) string {
	re := regexp.MustCompile(`[^a-zA-Z]`)
	return strings.ToLower(re.ReplaceAllString(role, ""))
}

func NormalizeName(name string) string {
	return strings.ToUpper(strings.Trim(name, " "))
}

func NormalizeEmail(email string) string {
	re := regexp.MustCompile(`[ ]`)
	return strings.ToLower(re.ReplaceAllString(email, ""))
}

func DeclarationIDToPatientIDAndDate(declarationID string) (string, string) {
	ls := strings.Split(declarationID, "_")
	if len(ls) == 1 {
		return ls[0], ""
	}
	if len(ls) == 0 {
		return "", ""
	}
	return ls[0], ls[1]
}

func SafeCastInt64(val interface{}) int64 {
	switch val.(type) {
	case int:
		return int64(val.(int))
	case string:
		i, err := strconv.ParseInt(val.(string), 10, 64)
		if err != nil {
			return 0
		}
		return i
	case float64:
		return int64(val.(float64))
	case int64:
		return val.(int64)
	case int32:
		return int64(val.(int32))
	default:
		return 0
	}
}

// GetDatesByRange gets date given range
func GetDatesByRange(from, to string) []string {
	var dates []string
	currentDate := from
	for currentDate <= to {
		dates = append(dates, currentDate)
		t, err := DateStringToTime(currentDate)
		if err != nil {
			return dates
		}
		currentDate = TimeToDateString(t.AddDate(0, 0, 1))
	}
	return dates
}

// MalaysiaTime gets Malaysia time
func MalaysiaTime(t time.Time) time.Time {
	// Load required location
	location, err := time.LoadLocation("Asia/Kuala_Lumpur")
	if err != nil {
		return t
	}

	return t.In(location)
}

// DaysElapsed find number of days elapsed given two days
func DaysElapsed(from time.Time, to time.Time) int64 {
	duration := (MalaysiaTime(to).Sub(MalaysiaTime(from))).Hours() / 24
	return int64(duration)
}

// TimeToMilli converts time to millisecond
func TimeToMilli(time time.Time) int64 {
	return MalaysiaTime(time).UnixNano() / 1000000
}

// MilliToTime converts millisecond to time
func MilliToTime(milli int64) time.Time {
	return MalaysiaTime(time.Unix(0, milli*int64(time.Millisecond)))
}

// DateStringToTime converts date string to time
func DateStringToTime(date string) (time.Time, error) {
	t, err := time.Parse("20060102", date)
	if err != nil {
		return time.Now(), err
	}
	t = t.Add(-8 * time.Hour)

	return MalaysiaTime(t), nil
}

// TimeToDateString timestamp to date string (yyyyMMdd)
func TimeToDateString(t time.Time) string {
	return MalaysiaTime(t).Format("20060102")
}

// RemoveZeroWidth removes zero width characters from string
func RemoveZeroWidth(t string) string {
	rslt := strings.Map(func(r rune) rune {
		if r == '↵' || r == '\n' || unicode.IsGraphic(r) &&
			r != '\u2000' &&
			r != '\u2001' &&
			r != '\u2002' &&
			r != '\u2003' &&
			r != '\u2004' &&
			r != '\u2005' &&
			r != '\u2006' &&
			r != '\u2007' &&
			r != '\u2008' &&
			r != '\u2009' &&
			r != '\u200a' &&
			r != '\u202f' &&
			r != '\u205f' &&
			r != '⠀' &&
			r != '\u3000' {
			return r
		}
		return -1
	}, t)

	// for weird characters like zalgo
	if utf8.RuneCountInString(rslt) > 500 {
		return ""
	}

	rslt = strings.Trim(rslt, " ")

	return rslt
	//re := regexp.MustCompile(`[^ -~]`)
	//return re.ReplaceAllString(t, "")
}

func ParseHTMLTemplate(templateFilename string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFilename)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
