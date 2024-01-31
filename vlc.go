package vlc

import (
	"github.com/zivkovicmilos/go-vlc/client"
)

// VLC http server commands extracted from
// https://github.com/videolan/vlc/blob/f7bb59d9f51cc10b25ff86d34a3eff744e60c46e/share/lua/http/requests/README.txt

const (
	baseStatus   = "requests/status.json"
	basePlaylist = "requests/playlist.json"
	baseBrowse   = "requests/browse.json"
)

const (
	commandKey = "command"
	inputKey   = "input"
	idKey      = "id"
	valKey     = "val"
	optionKey  = "option"
	bandKey    = "band"
	dirKey     = "dir"
	uriKey     = "uri"
)

const (
	// Playlist commands //
	stopCommand             = "pl_stop"
	emptyCommand            = "pl_empty"
	playCommand             = "pl_play"
	pauseCommand            = "pl_pause"
	nextCommand             = "pl_next"
	previousCommand         = "pl_previous"
	deleteCommand           = "pl_delete"
	sortCommand             = "pl_sort"
	randomCommand           = "pl_random"
	loopCommand             = "pl_loop"
	repeatCommand           = "pl_repeat"
	serviceDiscoveryCommand = "pl_sd"
	forceResumeCommand      = "pl_forceresume"
	forcePauseCommand       = "pl_forcepause"

	// Input commands //
	inPlayCommand    = "in_play"
	inEnqueueCommand = "in_enqueue"

	// General commands //
	fullscreenCommand    = "fullscreen"
	volumeCommand        = "volume"
	seekCommand          = "seek"
	addSubtitleCommand   = "addsubtitle"
	preampCommand        = "preamp"
	equalizerCommand     = "equalizer"
	enableeqCommand      = "enableeq"
	setpresetCommand     = "setpreset"
	titleCommand         = "title"
	chapterCommand       = "title"
	audioTrackCommand    = "audio_track"
	videoTrackCommand    = "video_track"
	subtitleTrackCommand = "subtitle_track"
	audioDelayCommand    = "audiodelay"
	subtitleDelayCommand = "subdelay"
	rateCommand          = "rate"
	aspectRatioCommand   = "aspectratio"
)

// VLC is an instance of the VLC HTTP client
type VLC struct {
	client client.Client
}

// NewVLC creates a new VLC HTTP client instance
func NewVLC(client client.Client) *VLC {
	return &VLC{
		client: client,
	}
}
