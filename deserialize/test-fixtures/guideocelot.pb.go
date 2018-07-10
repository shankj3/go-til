package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf2 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CredType int32

const (
	CredType_VCS  CredType = 0
	CredType_REPO CredType = 1
	CredType_K8S  CredType = 2
)

var CredType_name = map[int32]string{
	0: "VCS",
	1: "REPO",
	2: "K8S",
}
var CredType_value = map[string]int32{
	"VCS":  0,
	"REPO": 1,
	"K8S":  2,
}

func (x CredType) String() string {
	return proto.EnumName(CredType_name, int32(x))
}
func (CredType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type SubCredType int32

const (
	SubCredType_BITBUCKET SubCredType = 0
	SubCredType_GITHUB    SubCredType = 1
	SubCredType_NEXUS     SubCredType = 2
	SubCredType_MAVEN     SubCredType = 3
	SubCredType_DOCKER    SubCredType = 4
	SubCredType_KUBECONF  SubCredType = 5
)

var SubCredType_name = map[int32]string{
	0: "BITBUCKET",
	1: "GITHUB",
	2: "NEXUS",
	3: "MAVEN",
	4: "DOCKER",
	5: "KUBECONF",
}
var SubCredType_value = map[string]int32{
	"BITBUCKET": 0,
	"GITHUB":    1,
	"NEXUS":     2,
	"MAVEN":     3,
	"DOCKER":    4,
	"KUBECONF":  5,
}

func (x SubCredType) String() string {
	return proto.EnumName(SubCredType_name, int32(x))
}
func (SubCredType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BuildReq struct {
	AcctRepo string `protobuf:"bytes,1,opt,name=acctRepo" json:"acctRepo,omitempty"`
	Hash     string `protobuf:"bytes,2,opt,name=hash" json:"hash,omitempty"`
	Branch   string `protobuf:"bytes,3,opt,name=branch" json:"branch,omitempty"`
}

func (m *BuildReq) Reset()                    { *m = BuildReq{} }
func (m *BuildReq) String() string            { return proto.CompactTextString(m) }
func (*BuildReq) ProtoMessage()               {}
func (*BuildReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BuildReq) GetAcctRepo() string {
	if m != nil {
		return m.AcctRepo
	}
	return ""
}

func (m *BuildReq) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *BuildReq) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

type AllCredsWrapper struct {
	RepoCreds *RepoCredWrapper `protobuf:"bytes,1,opt,name=repoCreds" json:"repoCreds,omitempty"`
	VcsCreds  *CredWrapper     `protobuf:"bytes,3,opt,name=vcsCreds" json:"vcsCreds,omitempty"`
}

func (m *AllCredsWrapper) Reset()                    { *m = AllCredsWrapper{} }
func (m *AllCredsWrapper) String() string            { return proto.CompactTextString(m) }
func (*AllCredsWrapper) ProtoMessage()               {}
func (*AllCredsWrapper) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AllCredsWrapper) GetRepoCreds() *RepoCredWrapper {
	if m != nil {
		return m.RepoCreds
	}
	return nil
}

func (m *AllCredsWrapper) GetVcsCreds() *CredWrapper {
	if m != nil {
		return m.VcsCreds
	}
	return nil
}

type CredWrapper struct {
	Vcs []*VCSCreds `protobuf:"bytes,2,rep,name=vcs" json:"vcs,omitempty"`
}

func (m *CredWrapper) Reset()                    { *m = CredWrapper{} }
func (m *CredWrapper) String() string            { return proto.CompactTextString(m) }
func (*CredWrapper) ProtoMessage()               {}
func (*CredWrapper) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CredWrapper) GetVcs() []*VCSCreds {
	if m != nil {
		return m.Vcs
	}
	return nil
}

type SSHKeyWrapper struct {
	AcctName   string `protobuf:"bytes,1,opt,name=acctName" json:"acctName,omitempty"`
	PrivateKey []byte `protobuf:"bytes,2,opt,name=privateKey,proto3" json:"privateKey,omitempty"`
	Type       string `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
}

func (m *SSHKeyWrapper) Reset()                    { *m = SSHKeyWrapper{} }
func (m *SSHKeyWrapper) String() string            { return proto.CompactTextString(m) }
func (*SSHKeyWrapper) ProtoMessage()               {}
func (*SSHKeyWrapper) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SSHKeyWrapper) GetAcctName() string {
	if m != nil {
		return m.AcctName
	}
	return ""
}

func (m *SSHKeyWrapper) GetPrivateKey() []byte {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *SSHKeyWrapper) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type VCSCreds struct {
	ClientId     string      `protobuf:"bytes,1,opt,name=clientId" json:"clientId,omitempty"`
	ClientSecret string      `protobuf:"bytes,2,opt,name=clientSecret" json:"clientSecret,omitempty"`
	Identifier   string      `protobuf:"bytes,8,opt,name=identifier" json:"identifier,omitempty"`
	TokenURL     string      `protobuf:"bytes,3,opt,name=tokenURL" json:"tokenURL,omitempty"`
	AcctName     string      `protobuf:"bytes,4,opt,name=acctName" json:"acctName,omitempty"`
	SshFileLoc   string      `protobuf:"bytes,6,opt,name=sshFileLoc" json:"sshFileLoc,omitempty"`
	Type         CredType    `protobuf:"varint,9,opt,name=type,enum=models.CredType" json:"type,omitempty"`
	SubType      SubCredType `protobuf:"varint,10,opt,name=subType,enum=models.SubCredType" json:"subType,omitempty"`
}

func (m *VCSCreds) Reset()                    { *m = VCSCreds{} }
func (m *VCSCreds) String() string            { return proto.CompactTextString(m) }
func (*VCSCreds) ProtoMessage()               {}
func (*VCSCreds) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *VCSCreds) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *VCSCreds) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *VCSCreds) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *VCSCreds) GetTokenURL() string {
	if m != nil {
		return m.TokenURL
	}
	return ""
}

func (m *VCSCreds) GetAcctName() string {
	if m != nil {
		return m.AcctName
	}
	return ""
}

func (m *VCSCreds) GetSshFileLoc() string {
	if m != nil {
		return m.SshFileLoc
	}
	return ""
}

func (m *VCSCreds) GetType() CredType {
	if m != nil {
		return m.Type
	}
	return CredType_VCS
}

func (m *VCSCreds) GetSubType() SubCredType {
	if m != nil {
		return m.SubType
	}
	return SubCredType_BITBUCKET
}

type RepoCredWrapper struct {
	Repo []*RepoCreds `protobuf:"bytes,3,rep,name=repo" json:"repo,omitempty"`
}

func (m *RepoCredWrapper) Reset()                    { *m = RepoCredWrapper{} }
func (m *RepoCredWrapper) String() string            { return proto.CompactTextString(m) }
func (*RepoCredWrapper) ProtoMessage()               {}
func (*RepoCredWrapper) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RepoCredWrapper) GetRepo() []*RepoCreds {
	if m != nil {
		return m.Repo
	}
	return nil
}

type RepoCreds struct {
	Username   string      `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password   string      `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	RepoUrl    string      `protobuf:"bytes,6,opt,name=repoUrl" json:"repoUrl,omitempty"`
	Identifier string      `protobuf:"bytes,8,opt,name=identifier" json:"identifier,omitempty"`
	AcctName   string      `protobuf:"bytes,4,opt,name=acctName" json:"acctName,omitempty"`
	Type       CredType    `protobuf:"varint,9,opt,name=type,enum=models.CredType" json:"type,omitempty"`
	SubType    SubCredType `protobuf:"varint,10,opt,name=subType,enum=models.SubCredType" json:"subType,omitempty"`
}

func (m *RepoCreds) Reset()                    { *m = RepoCreds{} }
func (m *RepoCreds) String() string            { return proto.CompactTextString(m) }
func (*RepoCreds) ProtoMessage()               {}
func (*RepoCreds) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RepoCreds) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RepoCreds) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *RepoCreds) GetRepoUrl() string {
	if m != nil {
		return m.RepoUrl
	}
	return ""
}

func (m *RepoCreds) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *RepoCreds) GetAcctName() string {
	if m != nil {
		return m.AcctName
	}
	return ""
}

func (m *RepoCreds) GetType() CredType {
	if m != nil {
		return m.Type
	}
	return CredType_VCS
}

func (m *RepoCreds) GetSubType() SubCredType {
	if m != nil {
		return m.SubType
	}
	return SubCredType_BITBUCKET
}

type K8SCreds struct {
	AcctName    string      `protobuf:"bytes,1,opt,name=acctName" json:"acctName,omitempty"`
	K8SContents string      `protobuf:"bytes,2,opt,name=k8sContents" json:"k8sContents,omitempty"`
	Identifier  string      `protobuf:"bytes,3,opt,name=identifier" json:"identifier,omitempty"`
	Type        CredType    `protobuf:"varint,4,opt,name=type,enum=models.CredType" json:"type,omitempty"`
	SubType     SubCredType `protobuf:"varint,5,opt,name=subType,enum=models.SubCredType" json:"subType,omitempty"`
}

func (m *K8SCreds) Reset()                    { *m = K8SCreds{} }
func (m *K8SCreds) String() string            { return proto.CompactTextString(m) }
func (*K8SCreds) ProtoMessage()               {}
func (*K8SCreds) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *K8SCreds) GetAcctName() string {
	if m != nil {
		return m.AcctName
	}
	return ""
}

func (m *K8SCreds) GetK8SContents() string {
	if m != nil {
		return m.K8SContents
	}
	return ""
}

func (m *K8SCreds) GetIdentifier() string {
	if m != nil {
		return m.Identifier
	}
	return ""
}

func (m *K8SCreds) GetType() CredType {
	if m != nil {
		return m.Type
	}
	return CredType_VCS
}

func (m *K8SCreds) GetSubType() SubCredType {
	if m != nil {
		return m.SubType
	}
	return SubCredType_BITBUCKET
}

type K8SCredsWrapper struct {
	K8SCreds []*K8SCreds `protobuf:"bytes,2,rep,name=k8sCreds" json:"k8sCreds,omitempty"`
}

func (m *K8SCredsWrapper) Reset()                    { *m = K8SCredsWrapper{} }
func (m *K8SCredsWrapper) String() string            { return proto.CompactTextString(m) }
func (*K8SCredsWrapper) ProtoMessage()               {}
func (*K8SCredsWrapper) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *K8SCredsWrapper) GetK8SCreds() []*K8SCreds {
	if m != nil {
		return m.K8SCreds
	}
	return nil
}

type StatusQuery struct {
	Hash        string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	AcctName    string `protobuf:"bytes,2,opt,name=acctName" json:"acctName,omitempty"`
	RepoName    string `protobuf:"bytes,3,opt,name=repoName" json:"repoName,omitempty"`
	PartialRepo string `protobuf:"bytes,4,opt,name=partialRepo" json:"partialRepo,omitempty"`
}

func (m *StatusQuery) Reset()                    { *m = StatusQuery{} }
func (m *StatusQuery) String() string            { return proto.CompactTextString(m) }
func (*StatusQuery) ProtoMessage()               {}
func (*StatusQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *StatusQuery) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *StatusQuery) GetAcctName() string {
	if m != nil {
		return m.AcctName
	}
	return ""
}

func (m *StatusQuery) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

func (m *StatusQuery) GetPartialRepo() string {
	if m != nil {
		return m.PartialRepo
	}
	return ""
}

type BuildQuery struct {
	Hash    string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	BuildId int64  `protobuf:"varint,2,opt,name=buildId" json:"buildId,omitempty"`
}

func (m *BuildQuery) Reset()                    { *m = BuildQuery{} }
func (m *BuildQuery) String() string            { return proto.CompactTextString(m) }
func (*BuildQuery) ProtoMessage()               {}
func (*BuildQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *BuildQuery) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *BuildQuery) GetBuildId() int64 {
	if m != nil {
		return m.BuildId
	}
	return 0
}

type Builds struct {
	Builds map[string]*BuildRuntimeInfo `protobuf:"bytes,1,rep,name=builds" json:"builds,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Builds) Reset()                    { *m = Builds{} }
func (m *Builds) String() string            { return proto.CompactTextString(m) }
func (*Builds) ProtoMessage()               {}
func (*Builds) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *Builds) GetBuilds() map[string]*BuildRuntimeInfo {
	if m != nil {
		return m.Builds
	}
	return nil
}

type BuildRuntimeInfo struct {
	Done     bool   `protobuf:"varint,1,opt,name=done" json:"done,omitempty"`
	Ip       string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
	GrpcPort string `protobuf:"bytes,3,opt,name=grpcPort" json:"grpcPort,omitempty"`
	Hash     string `protobuf:"bytes,4,opt,name=hash" json:"hash,omitempty"`
	AcctName string `protobuf:"bytes,5,opt,name=acctName" json:"acctName,omitempty"`
	RepoName string `protobuf:"bytes,6,opt,name=repoName" json:"repoName,omitempty"`
}

func (m *BuildRuntimeInfo) Reset()                    { *m = BuildRuntimeInfo{} }
func (m *BuildRuntimeInfo) String() string            { return proto.CompactTextString(m) }
func (*BuildRuntimeInfo) ProtoMessage()               {}
func (*BuildRuntimeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *BuildRuntimeInfo) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *BuildRuntimeInfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *BuildRuntimeInfo) GetGrpcPort() string {
	if m != nil {
		return m.GrpcPort
	}
	return ""
}

func (m *BuildRuntimeInfo) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *BuildRuntimeInfo) GetAcctName() string {
	if m != nil {
		return m.AcctName
	}
	return ""
}

func (m *BuildRuntimeInfo) GetRepoName() string {
	if m != nil {
		return m.RepoName
	}
	return ""
}

type LineResponse struct {
	OutputLine string `protobuf:"bytes,1,opt,name=outputLine" json:"outputLine,omitempty"`
}

func (m *LineResponse) Reset()                    { *m = LineResponse{} }
func (m *LineResponse) String() string            { return proto.CompactTextString(m) }
func (*LineResponse) ProtoMessage()               {}
func (*LineResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *LineResponse) GetOutputLine() string {
	if m != nil {
		return m.OutputLine
	}
	return ""
}

type RepoAccount struct {
	Repo    string `protobuf:"bytes,1,opt,name=repo" json:"repo,omitempty"`
	Account string `protobuf:"bytes,2,opt,name=account" json:"account,omitempty"`
	Limit   int32  `protobuf:"varint,3,opt,name=limit" json:"limit,omitempty"`
}

func (m *RepoAccount) Reset()                    { *m = RepoAccount{} }
func (m *RepoAccount) String() string            { return proto.CompactTextString(m) }
func (*RepoAccount) ProtoMessage()               {}
func (*RepoAccount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *RepoAccount) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *RepoAccount) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *RepoAccount) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type Status struct {
	BuildSum   *BuildSummary `protobuf:"bytes,1,opt,name=buildSum" json:"buildSum,omitempty"`
	Stages     []*Stage      `protobuf:"bytes,2,rep,name=stages" json:"stages,omitempty"`
	IsInConsul bool          `protobuf:"varint,3,opt,name=isInConsul" json:"isInConsul,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *Status) GetBuildSum() *BuildSummary {
	if m != nil {
		return m.BuildSum
	}
	return nil
}

func (m *Status) GetStages() []*Stage {
	if m != nil {
		return m.Stages
	}
	return nil
}

func (m *Status) GetIsInConsul() bool {
	if m != nil {
		return m.IsInConsul
	}
	return false
}

type Stage struct {
	Stage         string                      `protobuf:"bytes,1,opt,name=stage" json:"stage,omitempty"`
	Error         string                      `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
	Status        int32                       `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	Messages      []string                    `protobuf:"bytes,4,rep,name=messages" json:"messages,omitempty"`
	StartTime     *google_protobuf2.Timestamp `protobuf:"bytes,5,opt,name=startTime" json:"startTime,omitempty"`
	StageDuration float64                     `protobuf:"fixed64,6,opt,name=stageDuration" json:"stageDuration,omitempty"`
}

func (m *Stage) Reset()                    { *m = Stage{} }
func (m *Stage) String() string            { return proto.CompactTextString(m) }
func (*Stage) ProtoMessage()               {}
func (*Stage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *Stage) GetStage() string {
	if m != nil {
		return m.Stage
	}
	return ""
}

func (m *Stage) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Stage) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Stage) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *Stage) GetStartTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Stage) GetStageDuration() float64 {
	if m != nil {
		return m.StageDuration
	}
	return 0
}

type BuildSummary struct {
	Hash          string                      `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Failed        bool                        `protobuf:"varint,2,opt,name=failed" json:"failed,omitempty"`
	BuildTime     *google_protobuf2.Timestamp `protobuf:"bytes,3,opt,name=buildTime" json:"buildTime,omitempty"`
	Account       string                      `protobuf:"bytes,4,opt,name=account" json:"account,omitempty"`
	BuildDuration float64                     `protobuf:"fixed64,5,opt,name=buildDuration" json:"buildDuration,omitempty"`
	Repo          string                      `protobuf:"bytes,6,opt,name=repo" json:"repo,omitempty"`
	Branch        string                      `protobuf:"bytes,7,opt,name=branch" json:"branch,omitempty"`
	BuildId       int64                       `protobuf:"varint,8,opt,name=buildId" json:"buildId,omitempty"`
	QueueTime     *google_protobuf2.Timestamp `protobuf:"bytes,9,opt,name=queueTime" json:"queueTime,omitempty"`
}

func (m *BuildSummary) Reset()                    { *m = BuildSummary{} }
func (m *BuildSummary) String() string            { return proto.CompactTextString(m) }
func (*BuildSummary) ProtoMessage()               {}
func (*BuildSummary) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *BuildSummary) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *BuildSummary) GetFailed() bool {
	if m != nil {
		return m.Failed
	}
	return false
}

func (m *BuildSummary) GetBuildTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.BuildTime
	}
	return nil
}

func (m *BuildSummary) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *BuildSummary) GetBuildDuration() float64 {
	if m != nil {
		return m.BuildDuration
	}
	return 0
}

func (m *BuildSummary) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *BuildSummary) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *BuildSummary) GetBuildId() int64 {
	if m != nil {
		return m.BuildId
	}
	return 0
}

func (m *BuildSummary) GetQueueTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.QueueTime
	}
	return nil
}

type Summaries struct {
	Sums []*BuildSummary `protobuf:"bytes,1,rep,name=sums" json:"sums,omitempty"`
}

func (m *Summaries) Reset()                    { *m = Summaries{} }
func (m *Summaries) String() string            { return proto.CompactTextString(m) }
func (*Summaries) ProtoMessage()               {}
func (*Summaries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *Summaries) GetSums() []*BuildSummary {
	if m != nil {
		return m.Sums
	}
	return nil
}

type PollRequest struct {
	Account      string                      `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Repo         string                      `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Cron         string                      `protobuf:"bytes,4,opt,name=cron" json:"cron,omitempty"`
	Branches     string                      `protobuf:"bytes,5,opt,name=branches" json:"branches,omitempty"`
	LastCronTime *google_protobuf2.Timestamp `protobuf:"bytes,6,opt,name=lastCronTime" json:"lastCronTime,omitempty"`
}

func (m *PollRequest) Reset()                    { *m = PollRequest{} }
func (m *PollRequest) String() string            { return proto.CompactTextString(m) }
func (*PollRequest) ProtoMessage()               {}
func (*PollRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *PollRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *PollRequest) GetRepo() string {
	if m != nil {
		return m.Repo
	}
	return ""
}

func (m *PollRequest) GetCron() string {
	if m != nil {
		return m.Cron
	}
	return ""
}

func (m *PollRequest) GetBranches() string {
	if m != nil {
		return m.Branches
	}
	return ""
}

func (m *PollRequest) GetLastCronTime() *google_protobuf2.Timestamp {
	if m != nil {
		return m.LastCronTime
	}
	return nil
}

type Polls struct {
	Polls []*PollRequest `protobuf:"bytes,1,rep,name=polls" json:"polls,omitempty"`
}

func (m *Polls) Reset()                    { *m = Polls{} }
func (m *Polls) String() string            { return proto.CompactTextString(m) }
func (*Polls) ProtoMessage()               {}
func (*Polls) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *Polls) GetPolls() []*PollRequest {
	if m != nil {
		return m.Polls
	}
	return nil
}

func init() {
	proto.RegisterType((*BuildReq)(nil), "models.BuildReq")
	proto.RegisterType((*AllCredsWrapper)(nil), "models.AllCredsWrapper")
	proto.RegisterType((*CredWrapper)(nil), "models.CredWrapper")
	proto.RegisterType((*SSHKeyWrapper)(nil), "models.SSHKeyWrapper")
	proto.RegisterType((*VCSCreds)(nil), "models.VCSCreds")
	proto.RegisterType((*RepoCredWrapper)(nil), "models.RepoCredWrapper")
	proto.RegisterType((*RepoCreds)(nil), "models.RepoCreds")
	proto.RegisterType((*K8SCreds)(nil), "models.K8sCreds")
	proto.RegisterType((*K8SCredsWrapper)(nil), "models.K8sCredsWrapper")
	proto.RegisterType((*StatusQuery)(nil), "models.StatusQuery")
	proto.RegisterType((*BuildQuery)(nil), "models.BuildQuery")
	proto.RegisterType((*Builds)(nil), "models.Builds")
	proto.RegisterType((*BuildRuntimeInfo)(nil), "models.BuildRuntimeInfo")
	proto.RegisterType((*LineResponse)(nil), "models.LineResponse")
	proto.RegisterType((*RepoAccount)(nil), "models.RepoAccount")
	proto.RegisterType((*Status)(nil), "models.Status")
	proto.RegisterType((*Stage)(nil), "models.Stage")
	proto.RegisterType((*BuildSummary)(nil), "models.BuildSummary")
	proto.RegisterType((*Summaries)(nil), "models.Summaries")
	proto.RegisterType((*PollRequest)(nil), "models.PollRequest")
	proto.RegisterType((*Polls)(nil), "models.Polls")
	proto.RegisterEnum("models.CredType", CredType_name, CredType_value)
	proto.RegisterEnum("models.SubCredType", SubCredType_name, SubCredType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for GuideOcelot service

type GuideOcelotClient interface {
	GetVCSCreds(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*CredWrapper, error)
	SetVCSCreds(ctx context.Context, in *VCSCreds, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	CheckConn(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	SetVCSPrivateKey(ctx context.Context, in *SSHKeyWrapper, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	GetRepoCreds(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*RepoCredWrapper, error)
	SetRepoCreds(ctx context.Context, in *RepoCreds, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	SetK8SCreds(ctx context.Context, in *K8SCreds, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	GetK8SCreds(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*K8SCredsWrapper, error)
	GetAllCreds(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*AllCredsWrapper, error)
	BuildRuntime(ctx context.Context, in *BuildQuery, opts ...grpc.CallOption) (*Builds, error)
	Logs(ctx context.Context, in *BuildQuery, opts ...grpc.CallOption) (GuideOcelot_LogsClient, error)
	LastFewSummaries(ctx context.Context, in *RepoAccount, opts ...grpc.CallOption) (*Summaries, error)
	GetStatus(ctx context.Context, in *StatusQuery, opts ...grpc.CallOption) (*Status, error)
	WatchRepo(ctx context.Context, in *RepoAccount, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	BuildRepoAndHash(ctx context.Context, in *BuildReq, opts ...grpc.CallOption) (GuideOcelot_BuildRepoAndHashClient, error)
	PollRepo(ctx context.Context, in *PollRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	DeletePollRepo(ctx context.Context, in *PollRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	FindWerker(ctx context.Context, in *BuildReq, opts ...grpc.CallOption) (*BuildRuntimeInfo, error)
	ListPolledRepos(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*Polls, error)
}

type guideOcelotClient struct {
	cc *grpc.ClientConn
}

func NewGuideOcelotClient(cc *grpc.ClientConn) GuideOcelotClient {
	return &guideOcelotClient{cc}
}

func (c *guideOcelotClient) GetVCSCreds(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*CredWrapper, error) {
	out := new(CredWrapper)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/GetVCSCreds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) SetVCSCreds(ctx context.Context, in *VCSCreds, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/SetVCSCreds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) CheckConn(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/CheckConn", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) SetVCSPrivateKey(ctx context.Context, in *SSHKeyWrapper, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/SetVCSPrivateKey", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) GetRepoCreds(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*RepoCredWrapper, error) {
	out := new(RepoCredWrapper)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/GetRepoCreds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) SetRepoCreds(ctx context.Context, in *RepoCreds, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/SetRepoCreds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) SetK8SCreds(ctx context.Context, in *K8SCreds, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/SetK8sCreds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) GetK8SCreds(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*K8SCredsWrapper, error) {
	out := new(K8SCredsWrapper)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/GetK8sCreds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) GetAllCreds(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*AllCredsWrapper, error) {
	out := new(AllCredsWrapper)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/GetAllCreds", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) BuildRuntime(ctx context.Context, in *BuildQuery, opts ...grpc.CallOption) (*Builds, error) {
	out := new(Builds)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/BuildRuntime", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) Logs(ctx context.Context, in *BuildQuery, opts ...grpc.CallOption) (GuideOcelot_LogsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GuideOcelot_serviceDesc.Streams[0], c.cc, "/models.GuideOcelot/Logs", opts...)
	if err != nil {
		return nil, err
	}
	x := &guideOcelotLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GuideOcelot_LogsClient interface {
	Recv() (*LineResponse, error)
	grpc.ClientStream
}

type guideOcelotLogsClient struct {
	grpc.ClientStream
}

func (x *guideOcelotLogsClient) Recv() (*LineResponse, error) {
	m := new(LineResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *guideOcelotClient) LastFewSummaries(ctx context.Context, in *RepoAccount, opts ...grpc.CallOption) (*Summaries, error) {
	out := new(Summaries)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/LastFewSummaries", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) GetStatus(ctx context.Context, in *StatusQuery, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/GetStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) WatchRepo(ctx context.Context, in *RepoAccount, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/WatchRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) BuildRepoAndHash(ctx context.Context, in *BuildReq, opts ...grpc.CallOption) (GuideOcelot_BuildRepoAndHashClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_GuideOcelot_serviceDesc.Streams[1], c.cc, "/models.GuideOcelot/BuildRepoAndHash", opts...)
	if err != nil {
		return nil, err
	}
	x := &guideOcelotBuildRepoAndHashClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GuideOcelot_BuildRepoAndHashClient interface {
	Recv() (*LineResponse, error)
	grpc.ClientStream
}

type guideOcelotBuildRepoAndHashClient struct {
	grpc.ClientStream
}

func (x *guideOcelotBuildRepoAndHashClient) Recv() (*LineResponse, error) {
	m := new(LineResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *guideOcelotClient) PollRepo(ctx context.Context, in *PollRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/PollRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) DeletePollRepo(ctx context.Context, in *PollRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/DeletePollRepo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) FindWerker(ctx context.Context, in *BuildReq, opts ...grpc.CallOption) (*BuildRuntimeInfo, error) {
	out := new(BuildRuntimeInfo)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/FindWerker", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guideOcelotClient) ListPolledRepos(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*Polls, error) {
	out := new(Polls)
	err := grpc.Invoke(ctx, "/models.GuideOcelot/ListPolledRepos", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GuideOcelot service

type GuideOcelotServer interface {
	GetVCSCreds(context.Context, *google_protobuf.Empty) (*CredWrapper, error)
	SetVCSCreds(context.Context, *VCSCreds) (*google_protobuf.Empty, error)
	CheckConn(context.Context, *google_protobuf.Empty) (*google_protobuf.Empty, error)
	SetVCSPrivateKey(context.Context, *SSHKeyWrapper) (*google_protobuf.Empty, error)
	GetRepoCreds(context.Context, *google_protobuf.Empty) (*RepoCredWrapper, error)
	SetRepoCreds(context.Context, *RepoCreds) (*google_protobuf.Empty, error)
	SetK8SCreds(context.Context, *K8SCreds) (*google_protobuf.Empty, error)
	GetK8SCreds(context.Context, *google_protobuf.Empty) (*K8SCredsWrapper, error)
	GetAllCreds(context.Context, *google_protobuf.Empty) (*AllCredsWrapper, error)
	BuildRuntime(context.Context, *BuildQuery) (*Builds, error)
	Logs(*BuildQuery, GuideOcelot_LogsServer) error
	LastFewSummaries(context.Context, *RepoAccount) (*Summaries, error)
	GetStatus(context.Context, *StatusQuery) (*Status, error)
	WatchRepo(context.Context, *RepoAccount) (*google_protobuf.Empty, error)
	BuildRepoAndHash(*BuildReq, GuideOcelot_BuildRepoAndHashServer) error
	PollRepo(context.Context, *PollRequest) (*google_protobuf.Empty, error)
	DeletePollRepo(context.Context, *PollRequest) (*google_protobuf.Empty, error)
	FindWerker(context.Context, *BuildReq) (*BuildRuntimeInfo, error)
	ListPolledRepos(context.Context, *google_protobuf.Empty) (*Polls, error)
}

func RegisterGuideOcelotServer(s *grpc.Server, srv GuideOcelotServer) {
	s.RegisterService(&_GuideOcelot_serviceDesc, srv)
}

func _GuideOcelot_GetVCSCreds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).GetVCSCreds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/GetVCSCreds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).GetVCSCreds(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_SetVCSCreds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VCSCreds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).SetVCSCreds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/SetVCSCreds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).SetVCSCreds(ctx, req.(*VCSCreds))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_CheckConn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).CheckConn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/CheckConn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).CheckConn(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_SetVCSPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SSHKeyWrapper)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).SetVCSPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/SetVCSPrivateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).SetVCSPrivateKey(ctx, req.(*SSHKeyWrapper))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_GetRepoCreds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).GetRepoCreds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/GetRepoCreds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).GetRepoCreds(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_SetRepoCreds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoCreds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).SetRepoCreds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/SetRepoCreds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).SetRepoCreds(ctx, req.(*RepoCreds))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_SetK8SCreds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(K8SCreds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).SetK8SCreds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/SetK8SCreds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).SetK8SCreds(ctx, req.(*K8SCreds))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_GetK8SCreds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).GetK8SCreds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/GetK8SCreds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).GetK8SCreds(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_GetAllCreds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).GetAllCreds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/GetAllCreds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).GetAllCreds(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_BuildRuntime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).BuildRuntime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/BuildRuntime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).BuildRuntime(ctx, req.(*BuildQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_Logs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BuildQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GuideOcelotServer).Logs(m, &guideOcelotLogsServer{stream})
}

type GuideOcelot_LogsServer interface {
	Send(*LineResponse) error
	grpc.ServerStream
}

type guideOcelotLogsServer struct {
	grpc.ServerStream
}

func (x *guideOcelotLogsServer) Send(m *LineResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GuideOcelot_LastFewSummaries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).LastFewSummaries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/LastFewSummaries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).LastFewSummaries(ctx, req.(*RepoAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).GetStatus(ctx, req.(*StatusQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_WatchRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepoAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).WatchRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/WatchRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).WatchRepo(ctx, req.(*RepoAccount))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_BuildRepoAndHash_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BuildReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GuideOcelotServer).BuildRepoAndHash(m, &guideOcelotBuildRepoAndHashServer{stream})
}

type GuideOcelot_BuildRepoAndHashServer interface {
	Send(*LineResponse) error
	grpc.ServerStream
}

type guideOcelotBuildRepoAndHashServer struct {
	grpc.ServerStream
}

func (x *guideOcelotBuildRepoAndHashServer) Send(m *LineResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _GuideOcelot_PollRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).PollRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/PollRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).PollRepo(ctx, req.(*PollRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_DeletePollRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).DeletePollRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/DeletePollRepo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).DeletePollRepo(ctx, req.(*PollRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_FindWerker_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).FindWerker(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/FindWerker",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).FindWerker(ctx, req.(*BuildReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuideOcelot_ListPolledRepos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuideOcelotServer).ListPolledRepos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/models.GuideOcelot/ListPolledRepos",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuideOcelotServer).ListPolledRepos(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _GuideOcelot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "models.GuideOcelot",
	HandlerType: (*GuideOcelotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVCSCreds",
			Handler:    _GuideOcelot_GetVCSCreds_Handler,
		},
		{
			MethodName: "SetVCSCreds",
			Handler:    _GuideOcelot_SetVCSCreds_Handler,
		},
		{
			MethodName: "CheckConn",
			Handler:    _GuideOcelot_CheckConn_Handler,
		},
		{
			MethodName: "SetVCSPrivateKey",
			Handler:    _GuideOcelot_SetVCSPrivateKey_Handler,
		},
		{
			MethodName: "GetRepoCreds",
			Handler:    _GuideOcelot_GetRepoCreds_Handler,
		},
		{
			MethodName: "SetRepoCreds",
			Handler:    _GuideOcelot_SetRepoCreds_Handler,
		},
		{
			MethodName: "SetK8sCreds",
			Handler:    _GuideOcelot_SetK8SCreds_Handler,
		},
		{
			MethodName: "GetK8sCreds",
			Handler:    _GuideOcelot_GetK8SCreds_Handler,
		},
		{
			MethodName: "GetAllCreds",
			Handler:    _GuideOcelot_GetAllCreds_Handler,
		},
		{
			MethodName: "BuildRuntime",
			Handler:    _GuideOcelot_BuildRuntime_Handler,
		},
		{
			MethodName: "LastFewSummaries",
			Handler:    _GuideOcelot_LastFewSummaries_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _GuideOcelot_GetStatus_Handler,
		},
		{
			MethodName: "WatchRepo",
			Handler:    _GuideOcelot_WatchRepo_Handler,
		},
		{
			MethodName: "PollRepo",
			Handler:    _GuideOcelot_PollRepo_Handler,
		},
		{
			MethodName: "DeletePollRepo",
			Handler:    _GuideOcelot_DeletePollRepo_Handler,
		},
		{
			MethodName: "FindWerker",
			Handler:    _GuideOcelot_FindWerker_Handler,
		},
		{
			MethodName: "ListPolledRepos",
			Handler:    _GuideOcelot_ListPolledRepos_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Logs",
			Handler:       _GuideOcelot_Logs_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BuildRepoAndHash",
			Handler:       _GuideOcelot_BuildRepoAndHash_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "guideocelot.proto",
}

func init() { proto.RegisterFile("guideocelot.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xcd, 0x52, 0x1b, 0xc7,
	0x13, 0xd7, 0xea, 0x8b, 0x55, 0x4b, 0xc0, 0x32, 0xf6, 0xdf, 0x7f, 0x95, 0x92, 0x4a, 0xa8, 0x29,
	0x3b, 0x45, 0x5c, 0x09, 0xd8, 0x24, 0xae, 0xc2, 0x24, 0xb1, 0x03, 0xb2, 0xc0, 0x14, 0x18, 0xf0,
	0x0a, 0x6c, 0xdf, 0x52, 0xcb, 0x6a, 0x80, 0x2d, 0x56, 0xbb, 0xcb, 0xce, 0x2c, 0x2e, 0x55, 0x4e,
	0xc9, 0x1b, 0xa4, 0x72, 0xc8, 0x2d, 0x2f, 0x90, 0x6b, 0x8e, 0xc9, 0x39, 0x0f, 0x90, 0x57, 0xc8,
	0x39, 0xcf, 0x90, 0x9a, 0xaf, 0xd5, 0x48, 0x20, 0x70, 0xa5, 0x2a, 0x27, 0x4d, 0xf7, 0x74, 0xf7,
	0xf4, 0xef, 0xd7, 0xb3, 0x3d, 0x2d, 0x98, 0x3b, 0xc9, 0x82, 0x1e, 0x89, 0x7d, 0x12, 0xc6, 0x6c,
	0x31, 0x49, 0x63, 0x16, 0xa3, 0x6a, 0x3f, 0xee, 0x91, 0x90, 0xb6, 0xde, 0x3b, 0x89, 0xe3, 0x93,
	0x90, 0x2c, 0x09, 0xed, 0x51, 0x76, 0xbc, 0x44, 0xfa, 0x09, 0x1b, 0x48, 0xa3, 0xd6, 0xfb, 0x6a,
	0xd3, 0x4b, 0x82, 0x25, 0x2f, 0x8a, 0x62, 0xe6, 0xb1, 0x20, 0x8e, 0xa8, 0xda, 0xfd, 0x70, 0xdc,
	0x95, 0x05, 0x7d, 0x42, 0x99, 0xd7, 0x4f, 0xa4, 0x01, 0x76, 0xc1, 0x5e, 0xcf, 0x82, 0xb0, 0xe7,
	0x92, 0x73, 0xd4, 0x02, 0xdb, 0xf3, 0x7d, 0xe6, 0x92, 0x24, 0x6e, 0x5a, 0xf3, 0xd6, 0x42, 0xcd,
	0xcd, 0x65, 0x84, 0xa0, 0x7c, 0xea, 0xd1, 0xd3, 0x66, 0x51, 0xe8, 0xc5, 0x1a, 0xdd, 0x81, 0xea,
	0x51, 0xea, 0x45, 0xfe, 0x69, 0xb3, 0x24, 0xb4, 0x4a, 0xc2, 0x03, 0x98, 0x5d, 0x0b, 0xc3, 0x76,
	0x4a, 0x7a, 0xf4, 0x75, 0xea, 0x25, 0x09, 0x49, 0xd1, 0x23, 0xa8, 0xa5, 0x24, 0x89, 0x85, 0x4e,
	0xc4, 0xae, 0x2f, 0xff, 0x7f, 0x51, 0xc2, 0x5b, 0x74, 0xd5, 0x86, 0xb2, 0x75, 0x87, 0x96, 0x68,
	0x09, 0xec, 0x0b, 0x9f, 0x4a, 0xaf, 0x92, 0xf0, 0xba, 0xa5, 0xbd, 0x4c, 0x8f, 0xdc, 0x08, 0x3f,
	0x84, 0xba, 0xb1, 0x81, 0x30, 0x94, 0x2e, 0x7c, 0xda, 0x2c, 0xce, 0x97, 0x16, 0xea, 0xcb, 0x8e,
	0x76, 0x7d, 0xd5, 0xee, 0x0a, 0x6b, 0x97, 0x6f, 0xe2, 0x6f, 0x60, 0xba, 0xdb, 0x7d, 0xbe, 0x4d,
	0x06, 0xda, 0x49, 0xd1, 0xb0, 0xeb, 0xf5, 0x89, 0x49, 0x03, 0x97, 0xd1, 0x07, 0x00, 0x49, 0x1a,
	0x5c, 0x78, 0x8c, 0x6c, 0x93, 0x81, 0x20, 0xa3, 0xe1, 0x1a, 0x1a, 0x4e, 0x13, 0x1b, 0x24, 0x44,
	0x11, 0x22, 0xd6, 0xf8, 0xa7, 0x22, 0xd8, 0xfa, 0x48, 0x1e, 0xdc, 0x0f, 0x03, 0x12, 0xb1, 0xad,
	0x9e, 0x0e, 0xae, 0x65, 0x84, 0xa1, 0x21, 0xd7, 0x5d, 0xe2, 0xa7, 0x84, 0x29, 0xae, 0x47, 0x74,
	0x3c, 0x81, 0xa0, 0x47, 0x22, 0x16, 0x1c, 0x07, 0x24, 0x6d, 0xda, 0xc2, 0xc2, 0xd0, 0xf0, 0xf8,
	0x2c, 0x3e, 0x23, 0xd1, 0xa1, 0xbb, 0xa3, 0x92, 0xc8, 0xe5, 0x11, 0x60, 0xe5, 0xcb, 0xc0, 0x28,
	0x3d, 0xdd, 0x08, 0x42, 0xb2, 0x13, 0xfb, 0xcd, 0xaa, 0x8c, 0x3b, 0xd4, 0xa0, 0xbb, 0x0a, 0x58,
	0x6d, 0xde, 0x5a, 0x98, 0x19, 0x52, 0xc9, 0x41, 0x1d, 0x0c, 0x12, 0x22, 0xa1, 0xa2, 0x4f, 0x61,
	0x8a, 0x66, 0x47, 0x5c, 0xd1, 0x04, 0x61, 0x98, 0x97, 0xab, 0x9b, 0x1d, 0xe5, 0xb6, 0xda, 0x06,
	0xaf, 0xc0, 0xec, 0x58, 0xf1, 0xd1, 0x3d, 0x28, 0xf3, 0xf2, 0x37, 0x4b, 0xa2, 0x64, 0x73, 0xe3,
	0x77, 0x84, 0xba, 0x62, 0x1b, 0xff, 0x6d, 0x41, 0x2d, 0xd7, 0x71, 0x60, 0x19, 0x25, 0x69, 0x64,
	0x54, 0x4c, 0xcb, 0x7c, 0x2f, 0xf1, 0x28, 0x7d, 0x1b, 0xa7, 0x3d, 0x45, 0x68, 0x2e, 0xa3, 0x26,
	0x4c, 0xf1, 0x68, 0x87, 0x69, 0xa8, 0x10, 0x6b, 0xf1, 0x5d, 0x68, 0x9e, 0x48, 0xe5, 0x7f, 0x42,
	0xd5, 0xef, 0x16, 0xd8, 0xdb, 0x2b, 0x34, 0xc7, 0x3b, 0xf1, 0x86, 0xce, 0x43, 0xfd, 0x6c, 0x85,
	0xb6, 0xe3, 0x88, 0x91, 0x88, 0x51, 0x05, 0xd9, 0x54, 0x8d, 0x61, 0x2b, 0x5d, 0xc2, 0xa6, 0xf3,
	0x2f, 0xbf, 0x6b, 0xfe, 0x95, 0x77, 0xc8, 0xff, 0x29, 0xcc, 0xea, 0xf4, 0x75, 0xa9, 0x3f, 0x01,
	0xfb, 0x4c, 0xa9, 0xc6, 0xbf, 0x50, 0x6d, 0xea, 0xe6, 0x16, 0xf8, 0x5b, 0xa8, 0x77, 0x99, 0xc7,
	0x32, 0xfa, 0x32, 0x23, 0xe9, 0x20, 0xef, 0x47, 0x96, 0xd1, 0x8f, 0x4c, 0x5a, 0x8a, 0x63, 0xb4,
	0xb4, 0xc0, 0xe6, 0xb5, 0x15, 0x7b, 0xea, 0xbb, 0xd0, 0x32, 0xa7, 0x2c, 0xf1, 0x52, 0x16, 0x78,
	0xa1, 0x68, 0x7d, 0xb2, 0x9e, 0xa6, 0x0a, 0xaf, 0x02, 0x88, 0x2e, 0x39, 0xf9, 0xec, 0x26, 0x4c,
	0x1d, 0x71, 0x8b, 0x2d, 0x79, 0xcb, 0x4a, 0xae, 0x16, 0xf1, 0x0f, 0x16, 0x54, 0x85, 0x33, 0x45,
	0xcb, 0x50, 0x15, 0x5a, 0xde, 0x02, 0x39, 0xde, 0x96, 0xc6, 0x2b, 0xf7, 0xd5, 0x4f, 0x27, 0x62,
	0xe9, 0xc0, 0x55, 0x96, 0xad, 0x2e, 0xd4, 0x0d, 0x35, 0x72, 0xa0, 0x74, 0x46, 0x06, 0xea, 0x68,
	0xbe, 0x44, 0x8b, 0x50, 0xb9, 0xf0, 0xc2, 0x4c, 0x42, 0xae, 0x2f, 0x37, 0x47, 0x62, 0xba, 0x59,
	0xc4, 0x3b, 0xfe, 0x56, 0x74, 0x1c, 0xbb, 0xd2, 0x6c, 0xb5, 0xb8, 0x62, 0xe1, 0x9f, 0x2d, 0x70,
	0xc6, 0xf7, 0x39, 0xac, 0x5e, 0x1c, 0xc9, 0x1b, 0x65, 0xbb, 0x62, 0x8d, 0x66, 0xa0, 0x18, 0x24,
	0x8a, 0xcc, 0x62, 0x90, 0x70, 0x1a, 0x4f, 0xd2, 0xc4, 0xdf, 0x8f, 0x53, 0xa6, 0x69, 0xd4, 0x72,
	0x4e, 0x4b, 0x79, 0x42, 0x49, 0x2a, 0xd7, 0x94, 0xa4, 0x3a, 0x5a, 0x12, 0xbc, 0x08, 0x8d, 0x9d,
	0x20, 0x22, 0x2e, 0xa1, 0x49, 0x1c, 0x51, 0xd1, 0x9e, 0xe2, 0x8c, 0x25, 0x19, 0xe3, 0x5a, 0x85,
	0xde, 0xd0, 0xe0, 0x97, 0x50, 0xe7, 0x85, 0x5a, 0xf3, 0xfd, 0x38, 0x8b, 0x44, 0x2a, 0xe9, 0xf0,
	0x15, 0x13, 0x6b, 0x5e, 0x21, 0x4f, 0x6e, 0x2b, 0x3c, 0x5a, 0x44, 0xb7, 0xa1, 0x12, 0x06, 0xfd,
	0x40, 0x22, 0xaa, 0xb8, 0x52, 0xc0, 0xdf, 0x59, 0x50, 0x95, 0x37, 0x0e, 0x3d, 0x00, 0x5b, 0x54,
	0xa3, 0x9b, 0xf5, 0xd5, 0xe3, 0x75, 0x7b, 0x84, 0xe5, 0x6e, 0xd6, 0xef, 0x7b, 0xe9, 0xc0, 0xcd,
	0xad, 0xd0, 0x3d, 0xa8, 0x52, 0xe6, 0x9d, 0x10, 0x7d, 0xb3, 0xa7, 0xf3, 0x8f, 0x83, 0x6b, 0x5d,
	0xb5, 0x29, 0x3e, 0x45, 0xba, 0x15, 0xb5, 0xe3, 0x88, 0x66, 0xa1, 0x38, 0xde, 0x76, 0x0d, 0x0d,
	0xfe, 0xc3, 0x82, 0x8a, 0xf0, 0xe0, 0x39, 0x0a, 0x1f, 0x05, 0x49, 0x0a, 0x5c, 0x4b, 0xd2, 0x34,
	0x4e, 0x15, 0x22, 0x29, 0xf0, 0x77, 0x99, 0x8a, 0xc4, 0x15, 0x20, 0x25, 0x71, 0xc2, 0xfb, 0x84,
	0x52, 0x91, 0x56, 0x79, 0xbe, 0xc4, 0x09, 0xd7, 0x32, 0x5a, 0x81, 0x1a, 0x65, 0x5e, 0xca, 0x0e,
	0x02, 0x55, 0x29, 0x7e, 0x3b, 0xe5, 0xf0, 0xb0, 0xa8, 0x87, 0x87, 0xc5, 0x03, 0x3d, 0x3c, 0xb8,
	0x43, 0x63, 0x74, 0x17, 0xa6, 0x45, 0x32, 0xcf, 0xb2, 0x54, 0x8c, 0x1e, 0xa2, 0x96, 0x96, 0x3b,
	0xaa, 0xc4, 0xbf, 0x16, 0xa1, 0x61, 0x72, 0x75, 0xe5, 0x47, 0x74, 0x07, 0xaa, 0xc7, 0x5e, 0x10,
	0x12, 0xf9, 0x0d, 0xd9, 0xae, 0x92, 0x78, 0x72, 0x82, 0x59, 0x91, 0x5c, 0xe9, 0xe6, 0xe4, 0x72,
	0x63, 0xb3, 0xe8, 0xe5, 0xd1, 0xa2, 0xdf, 0x85, 0x69, 0x61, 0x96, 0xa7, 0x5d, 0x91, 0x69, 0x8f,
	0x28, 0xf3, 0x8b, 0x54, 0x35, 0x2e, 0xd2, 0x70, 0xec, 0x99, 0x32, 0xc7, 0x1e, 0xb3, 0x05, 0xd8,
	0x23, 0x2d, 0x80, 0xe7, 0x7f, 0x9e, 0x91, 0x8c, 0x88, 0xfc, 0x6b, 0x37, 0xe7, 0x9f, 0x1b, 0xe3,
	0x47, 0x50, 0x93, 0x84, 0x05, 0x84, 0xa2, 0x05, 0x28, 0xd3, 0xac, 0xaf, 0x9b, 0xc7, 0xd5, 0x57,
	0x50, 0x58, 0xe0, 0x5f, 0x2c, 0xa8, 0xef, 0xc7, 0x61, 0xe8, 0x92, 0xf3, 0x8c, 0x50, 0x66, 0xd2,
	0x60, 0x8d, 0xd2, 0xa0, 0x01, 0x16, 0x0d, 0x80, 0x08, 0xca, 0x7e, 0x1a, 0x47, 0xfa, 0x43, 0xe6,
	0x6b, 0x7e, 0x77, 0x24, 0x4c, 0x42, 0xf5, 0x87, 0xac, 0x65, 0xf4, 0x04, 0x1a, 0xa1, 0x47, 0x59,
	0x3b, 0x8d, 0x23, 0x81, 0xb0, 0x7a, 0x23, 0xc2, 0x11, 0x7b, 0xbc, 0x0c, 0x15, 0x9e, 0x2c, 0x45,
	0x1f, 0x43, 0x25, 0xe1, 0x0b, 0x85, 0x30, 0x7f, 0x51, 0x0c, 0x28, 0xae, 0xb4, 0xb8, 0xff, 0x11,
	0xd8, 0xfa, 0x91, 0x41, 0x53, 0x50, 0x7a, 0xd5, 0xee, 0x3a, 0x05, 0x64, 0x43, 0xd9, 0xed, 0xec,
	0xef, 0x39, 0x16, 0x57, 0x6d, 0xaf, 0x74, 0x9d, 0xe2, 0xfd, 0x37, 0x50, 0x37, 0xde, 0x23, 0x34,
	0x0d, 0xb5, 0xf5, 0xad, 0x83, 0xf5, 0xc3, 0xf6, 0x76, 0xe7, 0xc0, 0x29, 0x20, 0x80, 0xea, 0xe6,
	0xd6, 0xc1, 0xf3, 0xc3, 0x75, 0xc7, 0x42, 0x35, 0xa8, 0xec, 0x76, 0xde, 0x1c, 0x76, 0x9d, 0x22,
	0x5f, 0xbe, 0x58, 0x7b, 0xd5, 0xd9, 0x75, 0x4a, 0xdc, 0xe2, 0xd9, 0x5e, 0x7b, 0xbb, 0xe3, 0x3a,
	0x65, 0xd4, 0x00, 0x7b, 0xfb, 0x70, 0xbd, 0xd3, 0xde, 0xdb, 0xdd, 0x70, 0x2a, 0xcb, 0xbf, 0xd5,
	0xa0, 0xbe, 0xc9, 0x67, 0xf6, 0x3d, 0x31, 0xb3, 0xa3, 0x17, 0x50, 0xdf, 0x24, 0x2c, 0x1f, 0xf4,
	0xee, 0x5c, 0x82, 0xdf, 0xe1, 0x53, 0x7b, 0xeb, 0xaa, 0x01, 0x16, 0xcf, 0x7d, 0xff, 0xe7, 0x5f,
	0x3f, 0x16, 0xeb, 0xa8, 0xb6, 0x74, 0xf1, 0x70, 0xc9, 0x17, 0xfe, 0x2f, 0xa0, 0xde, 0x35, 0xc2,
	0x5d, 0x1a, 0x5e, 0x5b, 0x13, 0x0e, 0xc0, 0xb7, 0x45, 0xac, 0x19, 0x3c, 0x8c, 0xb5, 0x6a, 0xdd,
	0x47, 0x5f, 0x41, 0xad, 0x7d, 0x4a, 0xfc, 0xb3, 0x76, 0x1c, 0x45, 0x13, 0x73, 0x9b, 0x14, 0xb2,
	0x80, 0xda, 0xe0, 0xc8, 0x6c, 0xf6, 0x87, 0xb3, 0xee, 0xff, 0xf2, 0x9e, 0x66, 0x8e, 0xcf, 0xd7,
	0x04, 0x59, 0x83, 0xc6, 0x26, 0x61, 0xc3, 0xb1, 0x6d, 0x52, 0x1a, 0x93, 0xfe, 0x19, 0xe0, 0x02,
	0xfa, 0x02, 0x1a, 0x5d, 0x33, 0xc4, 0xe5, 0x01, 0xf1, 0x9a, 0xf3, 0x1f, 0x0b, 0x4a, 0xf3, 0x29,
	0xea, 0xd2, 0xb4, 0x71, 0x8d, 0xeb, 0xd7, 0xa2, 0xb8, 0xb9, 0xeb, 0x8d, 0x99, 0x8f, 0xcd, 0x3a,
	0x79, 0x04, 0xfd, 0xbf, 0xe8, 0xe6, 0x08, 0x63, 0xff, 0xa0, 0x70, 0x01, 0x7d, 0xae, 0x3a, 0xa8,
	0x7a, 0xb3, 0x11, 0x1a, 0x69, 0x00, 0x62, 0x34, 0x69, 0xcd, 0x8c, 0x4e, 0x14, 0xc2, 0xab, 0xbc,
	0x13, 0x9f, 0xd0, 0x2b, 0xad, 0xf3, 0x16, 0x62, 0xbe, 0xb5, 0xb8, 0xf0, 0xc0, 0x42, 0x5f, 0x82,
	0xb3, 0xe3, 0x51, 0xb6, 0x41, 0xde, 0x0e, 0xdb, 0xcf, 0x2d, 0x93, 0x6b, 0xf5, 0xd2, 0xb6, 0xe6,
	0x86, 0x53, 0x9f, 0xb2, 0xc3, 0x05, 0xb4, 0x0c, 0xb5, 0x4d, 0xc2, 0xd4, 0xe3, 0x79, 0xcb, 0x78,
	0xfa, 0xf4, 0xf8, 0x36, 0xcc, 0x53, 0x2a, 0x71, 0x01, 0xad, 0x42, 0xed, 0xb5, 0xc7, 0xfc, 0x53,
	0xf1, 0x6f, 0xf3, 0xca, 0xa3, 0x26, 0x57, 0xe7, 0x89, 0x9e, 0x66, 0xb8, 0x75, 0xd4, 0x7b, 0xce,
	0xdf, 0x12, 0x67, 0x74, 0x0e, 0x22, 0xe7, 0xd7, 0xa0, 0x7d, 0x0c, 0xb6, 0x6c, 0x31, 0xe6, 0xd1,
	0x46, 0xd3, 0xb9, 0xe6, 0xe8, 0xa7, 0x30, 0xf3, 0x8c, 0x84, 0x84, 0x91, 0x7f, 0x1b, 0x60, 0x15,
	0x60, 0x23, 0x88, 0x7a, 0xaf, 0x49, 0x7a, 0x46, 0xd2, 0x2b, 0xb2, 0x9e, 0x38, 0xcf, 0x09, 0xdf,
	0xd9, 0x9d, 0x80, 0x32, 0x7e, 0x10, 0x11, 0xe0, 0x27, 0xdf, 0xab, 0x69, 0x33, 0x2b, 0x8a, 0x0b,
	0x47, 0x55, 0x61, 0xf0, 0xd9, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x13, 0xc3, 0xcf, 0x78,
	0x10, 0x00, 0x00,
}
