package vlc

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/zivkovicmilos/go-vlc/client"
)

var (
	errInvalidPlayOption      = errors.New("invalid play option")
	errInvalidSortMode        = errors.New("invalid playlist sort mode")
	errInvalidVolumeValue     = errors.New("invalid volume value")
	errInvalidSeekValue       = errors.New("invalid seek value")
	errInvalidPreampGainValue = errors.New("invalid preamp gain value")
	errInvalidPlaybackRate    = errors.New("invalid playback rate")
)

var (
	volumeRegex     = regexp.MustCompile(`^[+-]?\d+(%)?$`)
	seekNumberRegex = regexp.MustCompile(`^[+-]?\d+(%)?$`)
	seekFormatRegex = regexp.MustCompile(`^([+-])?(\d+[Hh])?(:)?(\d+[Mm'])?(:\d+([Ss"])?)?$`)
)

const (
	playNoAudio = "noaudio"
	playNoVideo = "novideo"
)

// executeStatusRequest executes a GET request and parses the response JSON
func (v *VLC) executeStatusRequest(params paramMap) (*Status, error) {
	endpoint := buildQueryEndpoint(baseStatus, params)

	statusRaw, err := v.client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to execute request, %s, %w", endpoint, err)
	}

	return client.ParseJSONResponse[Status](statusRaw)
}

// GetStatus returns the latest status information,
// including current item info and metadata
func (v *VLC) GetStatus() (*Status, error) {
	return v.executeStatusRequest(nil)
}

// EmptyPlaylist empties the current playlist
func (v *VLC) EmptyPlaylist() (*Status, error) {
	params := paramMap{
		commandKey: emptyCommand,
	}

	return v.executeStatusRequest(params)
}

// PlaySource adds the given source (URI) to the playlist and starts playing with the given option.
// Options can have the value of:
//   - noaudio
//   - novideo
func (v *VLC) PlaySource(source string, option ...string) (*Status, error) {
	params := paramMap{
		commandKey: inPlayCommand,
		inputKey:   source,
	}

	if len(option) >= 1 {
		op := option[0]

		// Make sure the option is valid
		if op != playNoAudio && op != playNoVideo {
			return nil, errInvalidPlayOption
		}

		params[optionKey] = option[0]
	}

	return v.executeStatusRequest(params)
}

// AddToPlaylist adds a source (URI) to the playlist
func (v *VLC) AddToPlaylist(source string) (*Status, error) {
	params := paramMap{
		commandKey: inEnqueueCommand,
		inputKey:   source,
	}

	return v.executeStatusRequest(params)
}

// PlayLastActivePlaylistItem plays the last active item
func (v *VLC) PlayLastActivePlaylistItem() (*Status, error) {
	params := paramMap{
		commandKey: playCommand,
	}

	return v.executeStatusRequest(params)
}

// PlayPlaylistItem plays the playlist item with the given ID
func (v *VLC) PlayPlaylistItem(id int) (*Status, error) {
	params := paramMap{
		commandKey: playCommand,
		idKey:      strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

// PauseWithLastActivePlaylistItem pauses playback. If the current state was 'stop',
// it plays the current item. If there is no current item, it plays the first item in the playlist
func (v *VLC) PauseWithLastActivePlaylistItem() (*Status, error) {
	params := paramMap{
		commandKey: pauseCommand,
	}

	return v.executeStatusRequest(params)
}

// PausePlaylist pauses playback. If the current state was 'stop',
// it plays the item with the given ID
func (v *VLC) PausePlaylist(id int) (*Status, error) {
	params := paramMap{
		commandKey: pauseCommand,
		idKey:      strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

// ForceResumePlaylist resumes playback if paused, otherwise does nothing
func (v *VLC) ForceResumePlaylist() (*Status, error) {
	params := paramMap{
		commandKey: forceResumeCommand,
	}

	return v.executeStatusRequest(params)
}

// ForcePausePlaylist pauses playback if not paused, otherwise does nothing
func (v *VLC) ForcePausePlaylist() (*Status, error) {
	params := paramMap{
		commandKey: forcePauseCommand,
	}

	return v.executeStatusRequest(params)
}

// StopPlaylist stops the playback
func (v *VLC) StopPlaylist() (*Status, error) {
	params := paramMap{
		commandKey: stopCommand,
	}

	return v.executeStatusRequest(params)
}

// PlayNextInPlaylist plays the next item in the playlist
func (v *VLC) PlayNextInPlaylist() (*Status, error) {
	params := paramMap{
		commandKey: nextCommand,
	}

	return v.executeStatusRequest(params)
}

// PlayPreviousInPlaylist plays the previous item in the playlist
func (v *VLC) PlayPreviousInPlaylist() (*Status, error) {
	params := paramMap{
		commandKey: previousCommand,
	}

	return v.executeStatusRequest(params)
}

// DeleteFromPlaylist deletes the item with the given ID from the playlist
func (v *VLC) DeleteFromPlaylist(id int) (*Status, error) {
	params := paramMap{
		commandKey: deleteCommand,
		idKey:      strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

// SortPlaylist sorts the playlist using the given sort mode and order.
// If id=0 items will be sorted in the normal order, otherwise if id=1 they will be
// sorted in the reverse order.
//
// A non-exhaustive list of sort modes:
//   - 0 ID
//   - 1 Name
//   - 3 Author
//   - 5 Random
//   - 7 Track number
func (v *VLC) SortPlaylist(id, mode int) (*Status, error) {
	// Make sure the sort ID is valid
	if id != 0 && id != 1 {
		return nil, errInvalidSortMode
	}

	params := paramMap{
		commandKey: sortCommand,
		idKey:      strconv.Itoa(id),
		valKey:     strconv.Itoa(mode),
	}

	return v.executeStatusRequest(params)
}

// TogglePlaylistRandom toggles random playlist playback
func (v *VLC) TogglePlaylistRandom() (*Status, error) {
	params := paramMap{
		commandKey: randomCommand,
	}

	return v.executeStatusRequest(params)
}

// TogglePlaylistLoop toggles a playlist playback loop
func (v *VLC) TogglePlaylistLoop() (*Status, error) {
	params := paramMap{
		commandKey: loopCommand,
	}

	return v.executeStatusRequest(params)
}

// TogglePlaylistRepeat toggles a playlist playback repeat
func (v *VLC) TogglePlaylistRepeat() (*Status, error) {
	params := paramMap{
		commandKey: repeatCommand,
	}

	return v.executeStatusRequest(params)
}

// EnableServiceDiscoveryModule enables the given service discovery module.
// Typical values are:
//   - sap
//   - shoutcast
//   - podcast
//   - hal
func (v *VLC) EnableServiceDiscoveryModule(module string) (*Status, error) {
	params := paramMap{
		commandKey: serviceDiscoveryCommand,
		valKey:     module,
	}

	return v.executeStatusRequest(params)
}

// ToggleFullscreen toggles fullscreen playback
func (v *VLC) ToggleFullscreen() (*Status, error) {
	params := paramMap{
		commandKey: fullscreenCommand,
	}

	return v.executeStatusRequest(params)
}

// SetVolume sets the playback volume.
//
// Allowed values are of the form:
// +<int>, -<int>, <int> or <int>%
func (v *VLC) SetVolume(volume string) (*Status, error) {
	// Make sure the volume value is valid
	if !volumeRegex.MatchString(volume) {
		return nil, errInvalidVolumeValue
	}

	params := paramMap{
		commandKey: volumeCommand,
		valKey:     volume,
	}

	return v.executeStatusRequest(params)
}

// SeekToValue seeks the playback to the given value.
//
// Allowed values are of the form:
//
// [+ or -][<int><H or h>:][<int><M or m or '>:][<int><nothing or S or s or ">]
//
// or [+ or -]<int>%
//
// (value between [ ] are optional, value between < > are mandatory)
//
// examples:
// 1000 -> seek to the 1000th second
// +1H:2M -> seek 1 hour and 2 minutes forward
// -10% -> seek 10% back
func (v *VLC) SeekToValue(value string) (*Status, error) {
	// Make sure the seek value is valid
	if !seekNumberRegex.MatchString(value) &&
		!seekFormatRegex.MatchString(value) {
		return nil, errInvalidSeekValue
	}

	params := paramMap{
		commandKey: seekCommand,
		valKey:     value,
	}

	return v.executeStatusRequest(params)
}

// AddSubtitle adds the given subtitle to the currently playing file
func (v *VLC) AddSubtitle(subtitleURI string) (*Status, error) {
	params := paramMap{
		commandKey: addSubtitleCommand,
		valKey:     subtitleURI,
	}

	return v.executeStatusRequest(params)
}

// SetPreamp sets the preamp value.
//
// Must be >=-20 and <=20
func (v *VLC) SetPreamp(gain int) (*Status, error) {
	// Make sure the gain value is valid
	if gain < -20 || gain > 20 {
		return nil, errInvalidPreampGainValue
	}

	params := paramMap{
		commandKey: preampCommand,
		valKey:     strconv.Itoa(gain),
	}

	return v.executeStatusRequest(params)
}

// SetEQ sets the gain for a specific band.
//
// Gain must be in Db and >=-20 and <=20
func (v *VLC) SetEQ(band, gain int) (*Status, error) {
	// Make sure the gain value is valid
	if gain < -20 || gain > 20 {
		return nil, errInvalidPreampGainValue
	}

	params := paramMap{
		commandKey: equalizerCommand,
		bandKey:    strconv.Itoa(band),
		valKey:     strconv.Itoa(gain),
	}

	return v.executeStatusRequest(params)
}

// EnableEQ enables or disables the equalizer
func (v *VLC) EnableEQ(value bool) (*Status, error) {
	enableValue := "0"

	if value {
		enableValue = "1"
	}

	params := paramMap{
		commandKey: enableeqCommand,
		valKey:     enableValue,
	}

	return v.executeStatusRequest(params)
}

// SetEQPreset sets the equalizer preset as per the ID specified.
//
// Equalizer gains:
// Band 0: 60 Hz, 1: 170 Hz, 2: 310 Hz, 3: 600 Hz, 4: 1 kHz,
// 5: 3 kHz, 6: 6 kHz, 7: 12 kHz , 8: 14 kHz , 9: 16 kHz
func (v *VLC) SetEQPreset(id int) (*Status, error) {
	params := paramMap{
		commandKey: setpresetCommand,
		idKey:      strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

// SelectTitle selects the title with the given title ID
func (v *VLC) SelectTitle(id int) (*Status, error) {
	params := paramMap{
		commandKey: titleCommand,
		valKey:     strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

// SelectChapter selects the chapter with the given chapter ID
func (v *VLC) SelectChapter(id int) (*Status, error) {
	params := paramMap{
		commandKey: chapterCommand,
		valKey:     strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

// SelectAudioTrack selects the audio track with the given audio track ID
// (use the number from the stream)
func (v *VLC) SelectAudioTrack(id int) (*Status, error) {
	params := paramMap{
		commandKey: audioTrackCommand,
		valKey:     strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

// SelectVideoTrack selects the video track with the given video track ID
// (use the number from the stream)
func (v *VLC) SelectVideoTrack(id int) (*Status, error) {
	params := paramMap{
		commandKey: videoTrackCommand,
		valKey:     strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

// SelectSubtitleTrack selects the subtitle track with the given subtitle track ID
// (use the number from the stream)
func (v *VLC) SelectSubtitleTrack(id int) (*Status, error) {
	params := paramMap{
		commandKey: subtitleTrackCommand,
		valKey:     strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

// SetAudioDelay sets the audio delay in seconds
func (v *VLC) SetAudioDelay(delay float64) (*Status, error) {
	params := paramMap{
		commandKey: audioDelayCommand,
		valKey:     fmt.Sprintf("%f", delay),
	}

	return v.executeStatusRequest(params)
}

// SetSubtitleDelay sets the subtitle delay in seconds
func (v *VLC) SetSubtitleDelay(delay float64) (*Status, error) {
	params := paramMap{
		commandKey: subtitleDelayCommand,
		valKey:     fmt.Sprintf("%f", delay),
	}

	return v.executeStatusRequest(params)
}

// SetPlaybackRate sets the playback rate.
//
// Must be > 0
func (v *VLC) SetPlaybackRate(rate float64) (*Status, error) {
	// Make sure the playback rate is valid
	if rate <= 0 {
		return nil, errInvalidPlaybackRate
	}

	params := paramMap{
		commandKey: rateCommand,
		valKey:     fmt.Sprintf("%f", rate),
	}

	return v.executeStatusRequest(params)
}

// SetAspectRatio sets the aspect ratio.
//
// Must be one of the following values. Any other value will reset aspect ratio to default
//   - 1:1
//   - 4:3
//   - 5:4
//   - 16:9
//   - 16:10
//   - 221:100
//   - 235:100
//   - 239:100
func (v *VLC) SetAspectRatio(ratio string) (*Status, error) {
	params := paramMap{
		commandKey: aspectRatioCommand,
		valKey:     ratio,
	}

	return v.executeStatusRequest(params)
}
