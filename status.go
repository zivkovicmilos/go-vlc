package vlc

import (
	"fmt"
	"strconv"

	"github.com/zivkovicmilos/go-vlc/client"
)

func (v *VLC) executeStatusRequest(params paramMap) (*Status, error) {
	endpoint := buildQueryEndpoint(baseStatus, params)

	statusRaw, err := v.client.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("unable to execute request, %s, %w", endpoint, err)
	}

	return client.ParseJSONResponse[Status](statusRaw)
}

func (v *VLC) GetStatus() (*Status, error) {
	return v.executeStatusRequest(nil)
}

func (v *VLC) EmptyPlaylist() (*Status, error) {
	params := paramMap{
		commandKey: emptyCommand,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) PlaySource(source string, option ...string) (*Status, error) {
	params := paramMap{
		commandKey: inPlayCommand,
		inputKey:   source,
	}

	if len(option) > 1 {
		params[optionKey] = option[0]
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) AddToPlaylist(source string) (*Status, error) {
	params := paramMap{
		commandKey: inEnqueueCommand,
		inputKey:   source,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) PlayLastActivePlaylistItem() (*Status, error) {
	params := paramMap{
		commandKey: playCommand,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) PlayPlaylistItem(id int) (*Status, error) {
	params := paramMap{
		commandKey: playCommand,
		idKey:      strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) PauseWithLastActivePlaylistItem() (*Status, error) {
	params := paramMap{
		commandKey: pauseCommand,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) PausePlaylist(id int) (*Status, error) {
	params := paramMap{
		commandKey: pauseCommand,
		idKey:      strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) StopPlaylist() (*Status, error) {
	params := paramMap{
		commandKey: stopCommand,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) PlayNextInPlaylist() (*Status, error) {
	params := paramMap{
		commandKey: nextCommand,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) PlayPreviousInPlaylist() (*Status, error) {
	params := paramMap{
		commandKey: previousCommand,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) DeleteFromPlaylist(id int) (*Status, error) {
	params := paramMap{
		commandKey: deleteCommand,
		idKey:      strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SortPlaylist(id, val int) (*Status, error) {
	params := paramMap{
		commandKey: sortCommand,
		idKey:      strconv.Itoa(id),
		valKey:     strconv.Itoa(val),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) TogglePlaylistRandom() (*Status, error) {
	params := paramMap{
		commandKey: randomCommand,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) TogglePlaylistLoop() (*Status, error) {
	params := paramMap{
		commandKey: loopCommand,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) TogglePlaylistRepeat() (*Status, error) {
	params := paramMap{
		commandKey: repeatCommand,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) EnableServiceDiscoveryModule(module string) (*Status, error) {
	params := paramMap{
		commandKey: serviceDiscoveryCommand,
		valKey:     module,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) ToggleFullscreen() (*Status, error) {
	params := paramMap{
		commandKey: fullscreenCommand,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SetVolumeValue(volume int) (*Status, error) {
	params := paramMap{
		commandKey: volumeCommand,
		valKey:     strconv.Itoa(volume),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SetVolumePercentage(volume int64) (*Status, error) {
	params := paramMap{
		commandKey: volumeCommand,
		valKey:     fmt.Sprintf("%d%%", volume),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SeekToValue(value string) (*Status, error) {
	params := paramMap{
		commandKey: seekCommand,
		valKey:     value,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) AddSubtitle(subtitleURI string) (*Status, error) {
	params := paramMap{
		commandKey: addSubtitleCommand,
		valKey:     subtitleURI,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) ForceResumePlaylist() (*Status, error) {
	params := paramMap{
		commandKey: forceResumeCommand,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) ForcePausePlaylist() (*Status, error) {
	params := paramMap{
		commandKey: forcePauseCommand,
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SetPreamp(gain int) (*Status, error) {
	params := paramMap{
		commandKey: preampCommand,
		valKey:     strconv.Itoa(gain),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SetEQ(band, gain int) (*Status, error) {
	params := paramMap{
		commandKey: equalizerCommand,
		bandKey:    strconv.Itoa(band),
		valKey:     strconv.Itoa(gain),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) ToggleEQ(value bool) (*Status, error) {
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

func (v *VLC) SetEQPreset(id int) (*Status, error) {
	params := paramMap{
		commandKey: setpresetCommand,
		idKey:      strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SetTitle(id int) (*Status, error) {
	params := paramMap{
		commandKey: titleCommand,
		valKey:     strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SetChapter(id int) (*Status, error) {
	params := paramMap{
		commandKey: chapterCommand,
		valKey:     strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SetAudioTrack(id int) (*Status, error) {
	params := paramMap{
		commandKey: audioTrackCommand,
		valKey:     strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SetVideoTrack(id int) (*Status, error) {
	params := paramMap{
		commandKey: videoTrackCommand,
		valKey:     strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SetSubtitleTrack(id int) (*Status, error) {
	params := paramMap{
		commandKey: subtitleTrackCommand,
		valKey:     strconv.Itoa(id),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SetAudioDelay(delay float64) (*Status, error) {
	params := paramMap{
		commandKey: audioDelayCommand,
		valKey:     strconv.FormatFloat(delay, 'f', -1, 64),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SetSubtitleDelay(delay float64) (*Status, error) {
	params := paramMap{
		commandKey: subtitleDelayCommand,
		valKey:     strconv.FormatFloat(delay, 'f', -1, 64),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SetPlaybackRate(rate float64) (*Status, error) {
	params := paramMap{
		commandKey: rateCommand,
		valKey:     strconv.FormatFloat(rate, 'f', -1, 64),
	}

	return v.executeStatusRequest(params)
}

func (v *VLC) SetAspectRatio(ratio string) (*Status, error) {
	params := paramMap{
		commandKey: aspectRatioCommand,
		valKey:     ratio,
	}

	return v.executeStatusRequest(params)
}
