package vlc

import "encoding/xml"

type Status struct {
	AudioFilters  map[string]string `json:"audiofilters"`
	Information   *Information      `json:"information,omitempty"`
	Stats         *Stats            `json:"stats,omitempty"`
	AspectRatio   string            `json:"aspectratio,omitempty"`
	Version       string            `json:"version"`
	State         string            `json:"state"`
	Equalizer     []Equalizer       `json:"equalizer"`
	VideoEffects  VideoEffects      `json:"videoeffects"`
	FullScreen    uint64            `json:"fullscreen"`
	Length        uint64            `json:"length"`
	APIVersion    uint64            `json:"apiversion"`
	Rate          float64           `json:"rate"`
	Volume        uint64            `json:"volume"`
	Time          uint64            `json:"time"`
	SeekSec       uint64            `json:"seek_sec"`
	CurrentPLID   int64             `json:"currentplid"`
	Position      float64           `json:"position"`
	AudioDelay    float64           `json:"audiodelay"`
	SubtitleDelay float64           `json:"subtitledelay"`
	Repeat        bool              `json:"repeat"`
	Loop          bool              `json:"loop"`
	Random        bool              `json:"random"`
}

type VideoEffects struct {
	Hue        float64 `json:"hue"`
	Saturation float64 `json:"saturation"`
	Contrast   float64 `json:"contrast"`
	Brightness float64 `json:"brightness"`
	Gamma      float64 `json:"gamma"`
}

type Equalizer struct {
	Presets map[string]string `json:"presets"`
	Bands   map[string]string `json:"bands"`
	Preamp  float64           `json:"preamp"`
}

type Information struct {
	Chapters []any                  `json:"chapters"` // TODO define
	Category map[string]StreamTable `json:"category"`
	Titles   []any                  `json:"titles"` // TODO define
	Chapter  int64                  `json:"chapter"`
	Title    int64                  `json:"title"`
}

type StreamTable struct {
	FileName string `json:"filename"` // Only present for the "meta" key

	DecodedFormat         string `json:"Decoded_format,omitempty"`
	ColorTransferFunction string `json:"Color_transfer_function,omitempty"`
	ChromaLocation        string `json:"Chroma_location,omitempty"`
	VideoResolution       string `json:"Video_resolution,omitempty"`
	FrameRate             string `json:"Frame_rate,omitempty"`
	Codec                 string `json:"Codec,omitempty"`
	Orientation           string `json:"Orientation,omitempty"`
	ColorSpace            string `json:"Color_space,omitempty"`
	Type                  string `json:"Type,omitempty"`
	ColorPrimaries        string `json:"Color_primaries,omitempty"`
	BufferDimensions      string `json:"Buffer_dimensions,omitempty"`
	Channels              string `json:"Channels,omitempty"`
	BitsPerSample         string `json:"Bits_per_sample,omitempty"`
	SampleRate            string `json:"Sample_rate,omitempty"`
}

type Stats struct {
	InputBitRate        float64 `json:"inputbitrate"`
	AveragedEMuxBitrate float64 `json:"averagedemuxbitrate"`
	DemuxBitRate        float64 `json:"demuxbitrate"`
	AverageInputBitRate float64 `json:"averageinputbitrate"`
	SendBitRate         float64 `json:"sendbitrate"`
	LosABuffers         uint64  `json:"lostabuffers"`
	ReadPackets         uint64  `json:"readpackets"`
	SentBytes           uint64  `json:"sentbytes"`
	DisplayedPictures   uint64  `json:"displayedpictures"`
	DemuxReadPackets    uint64  `json:"demuxreadpackets"`
	SentPackets         uint64  `json:"sentpackets"`
	DemuxReadBytes      uint64  `json:"demuxreadbytes"`
	DecodeAudio         uint64  `json:"decodedaudio"`
	PlayedABuffers      uint64  `json:"playedabuffers"`
	DemuxDiscontinuity  uint64  `json:"demuxdiscontinuity"`
	LostPictures        uint64  `json:"lostpictures"`
	DecodedVideo        uint64  `json:"decodedvideo"`
	ReadBytes           uint64  `json:"readbytes"`
	DemuxCorrupted      uint64  `json:"demuxcorrupted"`
}

type Playlist struct {
	Ro       string     `json:"ro"`
	Type     string     `json:"type"`
	Name     string     `json:"name"`
	ID       string     `json:"id"`
	URI      string     `json:"uri,omitempty"`
	Current  string     `json:"current,omitempty"`
	Children []Playlist `json:"children,omitempty"`
	Duration int64      `json:"duration,omitempty"`
}

type VLM struct {
	XMLName xml.Name `xml:"vlm"`
	Error   string   `xml:"error"`
}
