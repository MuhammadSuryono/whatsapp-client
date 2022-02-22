package models

type ParamSendMessage struct {
	Msisdn  string `form:"msisdn" json:"msisdn"`
	Message string `form:"message" json:"message"`
}

type ParamSendDocument struct {
	ParamSendMessage
	DocumentLink string `form:"document_link" json:"document_link"`
}

// ffmpeg -i file_example_MP4_1920_18MG.mp4 -profile:v baseline -level 3.0 -s 1280x720 -start_number 0 -hls_time 10 -hls_list_size 0 -f hls 1920.m3u8
