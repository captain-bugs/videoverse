package av

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"videoverse/pkg/config"
	"videoverse/pkg/logbox"
	"videoverse/pkg/utils"
)

type TOption func(*Options)

type Options struct {
	path string
	byts []byte
}

type Resolution struct {
	Height float64 `json:"height"`
	Width  float64 `json:"width"`
}

type AudioTrack struct {
	Bitrate   int     `json:"bitrate"`
	Duration  float64 `json:"duration"`
	StartTime float64 `json:"start_time"`
	StartPts  int64   `json:"start_pts"`
}

type VideoTrack struct {
	Bitrate   int     `json:"bitrate"`
	Duration  float64 `json:"duration"`
	StartTime float64 `json:"start_time"`
	StartPts  int64   `json:"start_pts"`
	PIX_FMT   string  `json:"pix_fmt"`
}

type StreamMeta struct {
	Width     float64 `json:"width"`
	Height    float64 `json:"height"`
	FrameRate string  `json:"r_frame_rate"`
	StartTime string  `json:"start_time"`
	Duration  string  `json:"duration"`
	Bitrate   string  `json:"bit_rate"`
	CodecType string  `json:"codec_type"`
	StartPts  int64   `json:"start_pts"`
}

type AVFile struct {
	Name         string     `json:"name"`
	Path         string     `json:"path"`
	FPS          float64    `json:"fps"`
	Duration     float64    `json:"duration"`
	Resolution   Resolution `json:"resolution"`
	StartTime    float64    `json:"start_time"`
	InBytes      []byte     `json:"in_bytes"`
	OutBytes     []byte     `json:"out_bytes"`
	AudioPresent bool       `json:"is_audio_present"`
	VideoPresent bool       `json:"is_video_present"`
	Audio        AudioTrack `json:"audio_track"`
	Video        VideoTrack `json:"video_track"`
	Bitrate      int        `json:"bitrate"`
}

func execute(cmd *exec.Cmd) (*string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		err := fmt.Errorf("stdout=%v \nstderr=%v \ncommand=%v", stdout.String(), stderr.String(), cmd.String())
		return nil, err
	}
	x := stdout.String()
	return &x, nil
}

func WithBytes(byts []byte) TOption {
	return func(o *Options) { o.byts = byts }
}

func WithFile(path string) TOption {
	return func(o *Options) { o.path = path }
}

func NewAVFile(opts ...TOption) *AVFile {
	for _, o := range opts {
		opt := &Options{}
		o(opt)
		if opt.path != "" {
			return &AVFile{Path: opt.path}
		}
		if len(opt.byts) > 0 {
			return &AVFile{InBytes: opt.byts}
		}
	}
	return nil
}

func (ts *AVFile) FetchMetaInfo() {
	type Stream struct {
		Width     float64 `json:"width"`
		Height    float64 `json:"height"`
		FrameRate string  `json:"r_frame_rate"`
		StartTime string  `json:"start_time"`
		Duration  string  `json:"duration"`
		Bitrate   string  `json:"bit_rate"`
		CodecType string  `json:"codec_type"`
		StartPts  int64   `json:"start_pts"`
		PIX_FMT   string  `json:"pix_fmt"`
	}

	type Meta struct {
		Programs []interface{} `json:"programs"`
		Stream   []Stream      `json:"streams"`
	}

	var cmd *exec.Cmd
	if ts.InBytes != nil {
		cmd = exec.Command("ffprobe", "-hide_banner", "-v", "error",
			"-show_entries", "stream=width,height,r_frame_rate,duration,start_time,bit_rate,codec_type,start_pts,pix_fmt",
			"-of", "json",
			"-i", "pipe:0",
		)
		cmd.Stdin = io.NopCloser(bytes.NewReader(ts.InBytes))
	} else {
		cmd = exec.Command("ffprobe", "-hide_banner", "-v", "error",
			"-show_entries", "stream=width,height,r_frame_rate,duration,start_time,bit_rate,codec_type,start_pts,pix_fmt",
			"-of", "json",
			ts.Path,
		)
	}
	logbox.NewLogBox().Debug().Str("cmd", cmd.String()).Msg("")
	stdout, err := execute(cmd)
	if err != nil {
		logbox.NewLogBox().Error().Str("event", "FAILED_TO_FETCH_META").
			Str("file", ts.Path).Str("error", err.Error()).
			Msg("")
		return
	}

	// stdout to json
	var data Meta
	_ = json.Unmarshal([]byte(*stdout), &data)

	if len(data.Stream) == 0 {
		logbox.NewLogBox().Error().Str("event", "EMPTY_STREAM_META").Str("file", ts.Path).Str("error", "streams array empty from meta pull").Msg("")
		panic(err)
	}

	var videoTrack, audioTrack Stream
	if len(data.Stream) == 1 {
		codec := data.Stream[0].CodecType
		ts.AudioPresent = false
		ts.VideoPresent = false
		if codec == "audio" {
			ts.AudioPresent = true
			audioTrack = data.Stream[0]
		} else {
			ts.VideoPresent = true
			videoTrack = data.Stream[0]
		}
	} else {
		ts.AudioPresent = true
		ts.VideoPresent = true
		video := utils.DotFind[Stream](data.Stream, func(val Stream) bool {
			return val.CodecType == "video"
		})
		audio := utils.DotFind[Stream](data.Stream, func(val Stream) bool {
			return val.CodecType == "audio"
		})
		videoTrack = *video
		audioTrack = *audio
	}

	if ts.VideoPresent {
		fpsA := strings.Split(videoTrack.FrameRate, "/")[0]
		fpsB := strings.Split(videoTrack.FrameRate, "/")[1]
		fpsPartA, _ := strconv.Atoi(fpsA)
		fpsPartB, _ := strconv.Atoi(fpsB)
		bitrate, _ := strconv.Atoi(videoTrack.Bitrate)
		duration, _ := strconv.ParseFloat(videoTrack.Duration, 64)
		startTime, _ := strconv.ParseFloat(videoTrack.StartTime, 64)

		ts.FPS = float64(fpsPartA) / float64(fpsPartB)
		ts.StartTime = startTime
		ts.Duration = duration
		ts.Bitrate = bitrate
		ts.Resolution = Resolution{
			Height: videoTrack.Height,
			Width:  videoTrack.Width,
		}
		ts.Video = VideoTrack{
			StartTime: startTime,
			Duration:  duration,
			Bitrate:   bitrate,
			StartPts:  videoTrack.StartPts,
			PIX_FMT:   videoTrack.PIX_FMT,
		}
	}

	if ts.AudioPresent {
		duration, _ := strconv.ParseFloat(audioTrack.Duration, 64)
		bitrate, _ := strconv.Atoi(audioTrack.Bitrate)
		startTime, _ := strconv.ParseFloat(audioTrack.StartTime, 64)

		ts.Audio = AudioTrack{
			StartTime: startTime,
			Duration:  duration,
			Bitrate:   bitrate,
			StartPts:  audioTrack.StartPts,
		}
	}
}

func (ts *AVFile) FetchDuration() error {
	var cmd *exec.Cmd
	if ts.InBytes != nil {
		cmd = exec.Command(
			"ffprobe", "-hide_banner", "-v", "error",
			"-show_entries", "format=duration",
			"-v", "quiet",
			"-of", "csv=p=0",
			"-i", "pipe:0",
		)
		cmd.Stdin = io.NopCloser(bytes.NewReader(ts.InBytes))
	} else {
		cmd = exec.Command(
			"ffprobe", "-hide_banner", "-v", "error",
			"-i", ts.Path,
			"-show_entries", "format=duration",
			"-v", "quiet",
			"-of", "csv=p=0",
		)
	}

	stdout, err := execute(cmd)
	if err != nil {
		logbox.NewLogBox().Error().Err(err).Str("event", "FAILED_TO_FETCH_DURATION").Str("file", ts.Path).Msg("")
		return err
	}
	s, err := strconv.ParseFloat(strings.Replace(*stdout, "\n", "", -1), 64)
	if err != nil {
		logbox.NewLogBox().Error().Str("event", "FAILED_TO_PARSE_DURATION").Str("file", ts.Path).Str("error", err.Error()).Msg("")
		return err
	}
	ts.Duration = s
	return nil
}

func (ts *AVFile) IsValidDuration() bool {
	if ts.Duration >= config.MIN_VIDEO_DURATION && ts.Duration <= config.MAX_VIDEO_DURATION {
		return true
	}
	return false
}

func (ts *AVFile) Validate() map[string]any {
	var errs = make(map[string]any)
	if !ts.IsValidDuration() {
		msg := fmt.Sprintf("file duration is %vs, should be between %vs and %vs", ts.Duration, config.MIN_VIDEO_DURATION, config.MAX_VIDEO_DURATION)
		errs["duration"] = msg
	}
	if !ts.VideoPresent {
		errs["track"] = "video track not found"
	}
	return errs
}

func (ts *AVFile) Trim(start, end float64) (*AVFile, error) {
	if ts.Duration < end {
		logbox.NewLogBox().Error().Str("event", "TRIM_END_GREATER_THAN_DURATION").Str("file", ts.Name).Msg("")
		return nil, fmt.Errorf("end time is greater than the duration of the video")
	}

	var cmd *exec.Cmd
	if ts.InBytes != nil {
		cmd = exec.Command(
			"ffmpeg", "-hide_banner", "-v", "error",
			"-i", "pipe:0",
			"-ss", utils.ConvertSecondsToDuration(start),
			"-to", utils.ConvertSecondsToDuration(end),
			"-c", "copy",
			"-f", "flv",
			"pipe:1",
		)
		cmd.Stdin = io.NopCloser(bytes.NewReader(ts.InBytes))
	} else {
		cmd = exec.Command(
			"ffmpeg", "-hide_banner", "-v", "error",
			"-i", ts.Path,
			"-ss", utils.ConvertSecondsToDuration(start),
			"-to", utils.ConvertSecondsToDuration(end),
			"-c", "copy",
			"-f", "flv",
			"pipe:1",
		)
	}
	logbox.NewLogBox().Debug().Str("cmd", cmd.String()).Msg("")

	stdout, err := execute(cmd)
	if err != nil {
		logbox.NewLogBox().Error().Str("event", "FAILED_TO_TRIM").Str("file", ts.Path).Str("error", err.Error()).Msg("")
		return nil, err
	}

	var trimmed AVFile
	trimmed.InBytes = []byte(*stdout)
	trimmed.Name = fmt.Sprintf("T_%v_%s", utils.GenerateUUID(), ts.Name)
	trimmed.Path = config.FILE_UPLOAD_PATH + "/" + trimmed.Name
	if err := trimmed.SaveToDisk(); err != nil {
		return nil, err
	}
	trimmed.FetchMetaInfo()
	_ = trimmed.FetchDuration()

	return &trimmed, nil
}

func (ts *AVFile) Read() {
	file, err := os.Open(ts.Path)
	if err != nil {
		return
	}
	defer func(file *os.File) { _ = file.Close() }(file)
	fileInfo, err := file.Stat()
	if err != nil {
		return
	}

	var size = fileInfo.Size()
	buffer := make([]byte, size)
	_, _ = file.Read(buffer)
	ts.InBytes = buffer
}

func (ts *AVFile) SaveToDisk() error {
	var file *os.File
	var err error

	file, err = os.Create(ts.Path)
	if err != nil {
		return fmt.Errorf("event=%s file=%s error=%v", "FAILED_TO_CREATE_ON_DISK", ts.Path, err.Error())
	}
	defer file.Close()
	_, err = file.Write(ts.InBytes)
	if err != nil {
		return fmt.Errorf("event=%s file=%s error=%v", "FAILED_TO_WRITE_ON_DISK", ts.Path, err.Error())
	}
	file.Sync()
	return nil
}

func Merge(filename, filepath string, files []*AVFile) *AVFile {

	//"concat:input1|input2"
	var args []string
	args = append(args, "-hide_banner", "-v", "error")
	for _, file := range files {
		args = append(args, "-i", file.Path)
	}
	args = append(args, "-filter_complex", fmt.Sprintf("concat=n=%d:v=1:a=1", len(files)), filepath)

	cmd := exec.Command("ffmpeg", args...)
	logbox.NewLogBox().Debug().Str("cmd", cmd.String()).Msg("")
	_, err := execute(cmd)
	if err != nil {
		logbox.NewLogBox().Error().Str("event", "FAILED_TO_MERGE_VIDEOS").Str("error", err.Error()).Msg("")
		return nil
	}
	merged := &AVFile{
		Name: fmt.Sprintf("M_%v_%s", utils.GenerateUUID(), filename),
		Path: filepath,
	}
	merged.FetchMetaInfo()
	merged.Read()
	return merged
}
