package config

const Sensor_Type_Unknown = 0
const Sensor_Type_Face = 1
const Sensor_Type_Capture = 2
const Sensor_Type_Ipc = 3
const Sensor_Type_Video = 4
const Sensor_Type_Picture = 5
const Sensor_Type_Gate = 6

var TYPE_REVERSE_MAPPING = map[int64]string{
	Sensor_Type_Face: "video",
	Sensor_Type_Ipc:  "image",
}

var TYPE_MAPPING = map[string]int64{
	"video": Sensor_Type_Face,
	"image": Sensor_Type_Ipc,
}

const (
	DATE_FORMAT     = "2006-01-02"
	TIME_FORMAT     = "15:04:05"
	DATETIME_FORMAT = "2006-01-02 15:04:05"
)
