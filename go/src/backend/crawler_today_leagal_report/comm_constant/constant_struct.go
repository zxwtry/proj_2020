package comm_constant

type XinWenLianBoVideoInfoManifest struct {
	HlsAudioUrl string `json:"hls_audio_url"`
	AudioMp3    string `json:"audio_mp3"`
}

type XinWenLianBoVideoInfo struct {
	Manifest XinWenLianBoVideoInfoManifest `json:"manifest"`
}

type ShenDuGuoJiVideoSetInfo struct {
	Guid      string `json:"guid"`
	Id        string `json:"id"`
	Time      string `json:"time"`
	Title     string `json:"title"`
	Image     string `json:"image"`
	FocusDate uint64 `json:"focus_date"`
	Brief     string `json:"brief"`
	Url       string `json:"url"`
}

type ShenDuGuoJiVideoSetData struct {
	Total int32                     `json:"total"`
	List  []ShenDuGuoJiVideoSetInfo `json:"list"`
}
type ShenDuGuoJiVideoSet struct {
	Data ShenDuGuoJiVideoSetData `json:"data"`
}
