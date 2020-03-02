package constant

type XinWenLianBoVideoInfoManifest struct {
	HlsAudioUrl string `json:"hls_audio_url"`
	AudioMp3    string `json:"audio_mp3"`
}

type XinWenLianBoVideoInfo struct {
	Manifest XinWenLianBoVideoInfoManifest `json:"manifest"`
}
