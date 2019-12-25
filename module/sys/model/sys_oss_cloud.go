package model

type SysOssCloud struct {
	AliyunDomain          string `json:"aliyunDomain" gconv:"AliyunDomain,omitempty"`
	AliyunEndPoint        string `json:"aliyunEndPoint" gconv:"AliyunEndPoint,omitempty"`
	AliyunAccessKeyId     string `json:"aliyunAccessKeyId" gconv:"AliyunAccessKeyId,omitempty"`
	AliyunAccessKeySecret string `json:"aliyunAccessKeySecret" gconv:"AliyunAccessKeySecret,omitempty"`
	AliyunBucketName      string `json:"aliyunBucketName" gconv:"AliyunBucketName,omitempty"`
}
