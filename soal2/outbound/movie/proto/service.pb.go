// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: outbound/movie/proto/service.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OmdbId string `protobuf:"bytes,1,opt,name=omdbId,proto3" json:"omdbId,omitempty"`
}

func (x *GetMovieRequest) Reset() {
	*x = GetMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outbound_movie_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieRequest) ProtoMessage() {}

func (x *GetMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_outbound_movie_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieRequest.ProtoReflect.Descriptor instead.
func (*GetMovieRequest) Descriptor() ([]byte, []int) {
	return file_outbound_movie_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetMovieRequest) GetOmdbId() string {
	if x != nil {
		return x.OmdbId
	}
	return ""
}

type Ratings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	Value  string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Ratings) Reset() {
	*x = Ratings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outbound_movie_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ratings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ratings) ProtoMessage() {}

func (x *Ratings) ProtoReflect() protoreflect.Message {
	mi := &file_outbound_movie_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ratings.ProtoReflect.Descriptor instead.
func (*Ratings) Descriptor() ([]byte, []int) {
	return file_outbound_movie_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *Ratings) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Ratings) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title      string     `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year       string     `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	Rated      string     `protobuf:"bytes,3,opt,name=rated,proto3" json:"rated,omitempty"`
	Released   string     `protobuf:"bytes,4,opt,name=released,proto3" json:"released,omitempty"`
	Runtime    string     `protobuf:"bytes,5,opt,name=runtime,proto3" json:"runtime,omitempty"`
	Genre      string     `protobuf:"bytes,6,opt,name=genre,proto3" json:"genre,omitempty"`
	Director   string     `protobuf:"bytes,7,opt,name=director,proto3" json:"director,omitempty"`
	Writer     string     `protobuf:"bytes,8,opt,name=writer,proto3" json:"writer,omitempty"`
	Actors     string     `protobuf:"bytes,9,opt,name=actors,proto3" json:"actors,omitempty"`
	Plot       string     `protobuf:"bytes,10,opt,name=plot,proto3" json:"plot,omitempty"`
	Language   string     `protobuf:"bytes,11,opt,name=language,proto3" json:"language,omitempty"`
	Country    string     `protobuf:"bytes,12,opt,name=country,proto3" json:"country,omitempty"`
	Awards     string     `protobuf:"bytes,13,opt,name=awards,proto3" json:"awards,omitempty"`
	Poster     string     `protobuf:"bytes,14,opt,name=poster,proto3" json:"poster,omitempty"`
	Ratings    []*Ratings `protobuf:"bytes,15,rep,name=ratings,proto3" json:"ratings,omitempty"`
	MetaScore  string     `protobuf:"bytes,16,opt,name=metaScore,proto3" json:"metaScore,omitempty"`
	ImdbRating string     `protobuf:"bytes,17,opt,name=imdbRating,proto3" json:"imdbRating,omitempty"`
	ImdbVotes  string     `protobuf:"bytes,18,opt,name=imdbVotes,proto3" json:"imdbVotes,omitempty"`
	ImdbID     string     `protobuf:"bytes,19,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type       string     `protobuf:"bytes,20,opt,name=type,proto3" json:"type,omitempty"`
	Dvd        string     `protobuf:"bytes,21,opt,name=dvd,proto3" json:"dvd,omitempty"`
	BoxOffice  string     `protobuf:"bytes,22,opt,name=boxOffice,proto3" json:"boxOffice,omitempty"`
	Production string     `protobuf:"bytes,23,opt,name=production,proto3" json:"production,omitempty"`
	Website    string     `protobuf:"bytes,24,opt,name=website,proto3" json:"website,omitempty"`
	Response   string     `protobuf:"bytes,25,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *GetMovieResponse) Reset() {
	*x = GetMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outbound_movie_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMovieResponse) ProtoMessage() {}

func (x *GetMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_outbound_movie_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMovieResponse.ProtoReflect.Descriptor instead.
func (*GetMovieResponse) Descriptor() ([]byte, []int) {
	return file_outbound_movie_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetMovieResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetMovieResponse) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *GetMovieResponse) GetRated() string {
	if x != nil {
		return x.Rated
	}
	return ""
}

func (x *GetMovieResponse) GetReleased() string {
	if x != nil {
		return x.Released
	}
	return ""
}

func (x *GetMovieResponse) GetRuntime() string {
	if x != nil {
		return x.Runtime
	}
	return ""
}

func (x *GetMovieResponse) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

func (x *GetMovieResponse) GetDirector() string {
	if x != nil {
		return x.Director
	}
	return ""
}

func (x *GetMovieResponse) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *GetMovieResponse) GetActors() string {
	if x != nil {
		return x.Actors
	}
	return ""
}

func (x *GetMovieResponse) GetPlot() string {
	if x != nil {
		return x.Plot
	}
	return ""
}

func (x *GetMovieResponse) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *GetMovieResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *GetMovieResponse) GetAwards() string {
	if x != nil {
		return x.Awards
	}
	return ""
}

func (x *GetMovieResponse) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

func (x *GetMovieResponse) GetRatings() []*Ratings {
	if x != nil {
		return x.Ratings
	}
	return nil
}

func (x *GetMovieResponse) GetMetaScore() string {
	if x != nil {
		return x.MetaScore
	}
	return ""
}

func (x *GetMovieResponse) GetImdbRating() string {
	if x != nil {
		return x.ImdbRating
	}
	return ""
}

func (x *GetMovieResponse) GetImdbVotes() string {
	if x != nil {
		return x.ImdbVotes
	}
	return ""
}

func (x *GetMovieResponse) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *GetMovieResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetMovieResponse) GetDvd() string {
	if x != nil {
		return x.Dvd
	}
	return ""
}

func (x *GetMovieResponse) GetBoxOffice() string {
	if x != nil {
		return x.BoxOffice
	}
	return ""
}

func (x *GetMovieResponse) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

func (x *GetMovieResponse) GetWebsite() string {
	if x != nil {
		return x.Website
	}
	return ""
}

func (x *GetMovieResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type GetMoviesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchWord string `protobuf:"bytes,1,opt,name=searchWord,proto3" json:"searchWord,omitempty"`
	Page       int32  `protobuf:"varint,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetMoviesRequest) Reset() {
	*x = GetMoviesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outbound_movie_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMoviesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMoviesRequest) ProtoMessage() {}

func (x *GetMoviesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_outbound_movie_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMoviesRequest.ProtoReflect.Descriptor instead.
func (*GetMoviesRequest) Descriptor() ([]byte, []int) {
	return file_outbound_movie_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetMoviesRequest) GetSearchWord() string {
	if x != nil {
		return x.SearchWord
	}
	return ""
}

func (x *GetMoviesRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type GetMoviesResponseData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title  string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Year   string `protobuf:"bytes,2,opt,name=year,proto3" json:"year,omitempty"`
	ImdbID string `protobuf:"bytes,3,opt,name=imdbID,proto3" json:"imdbID,omitempty"`
	Type   string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Poster string `protobuf:"bytes,5,opt,name=poster,proto3" json:"poster,omitempty"`
}

func (x *GetMoviesResponseData) Reset() {
	*x = GetMoviesResponseData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outbound_movie_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMoviesResponseData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMoviesResponseData) ProtoMessage() {}

func (x *GetMoviesResponseData) ProtoReflect() protoreflect.Message {
	mi := &file_outbound_movie_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMoviesResponseData.ProtoReflect.Descriptor instead.
func (*GetMoviesResponseData) Descriptor() ([]byte, []int) {
	return file_outbound_movie_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetMoviesResponseData) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetMoviesResponseData) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *GetMoviesResponseData) GetImdbID() string {
	if x != nil {
		return x.ImdbID
	}
	return ""
}

func (x *GetMoviesResponseData) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetMoviesResponseData) GetPoster() string {
	if x != nil {
		return x.Poster
	}
	return ""
}

type GetMoviesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data  []*GetMoviesResponseData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Total int32                    `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Pages int32                    `protobuf:"varint,3,opt,name=pages,proto3" json:"pages,omitempty"`
	Page  int32                    `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetMoviesResponse) Reset() {
	*x = GetMoviesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_outbound_movie_proto_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMoviesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMoviesResponse) ProtoMessage() {}

func (x *GetMoviesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_outbound_movie_proto_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMoviesResponse.ProtoReflect.Descriptor instead.
func (*GetMoviesResponse) Descriptor() ([]byte, []int) {
	return file_outbound_movie_proto_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetMoviesResponse) GetData() []*GetMoviesResponseData {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetMoviesResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *GetMoviesResponse) GetPages() int32 {
	if x != nil {
		return x.Pages
	}
	return 0
}

func (x *GetMoviesResponse) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

var File_outbound_movie_proto_service_proto protoreflect.FileDescriptor

var file_outbound_movie_proto_service_proto_rawDesc = []byte{
	0x0a, 0x22, 0x6f, 0x75, 0x74, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x2f, 0x6d, 0x6f, 0x76, 0x69, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x70, 0x63, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x29, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x6d, 0x64, 0x62, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x6d, 0x64, 0x62, 0x49, 0x64, 0x22, 0x37, 0x0a, 0x07,
	0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xa4, 0x05, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x79, 0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x6f,
	0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x6f, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x70, 0x6c, 0x6f, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x77, 0x61, 0x72, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x0e,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x12, 0x30, 0x0a, 0x07,
	0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x70, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x07, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6d, 0x65, 0x74, 0x61, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x69, 0x6d, 0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x11, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x69, 0x6d, 0x64, 0x62, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x6d, 0x64, 0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x6d, 0x64, 0x62, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d,
	0x64, 0x62, 0x49, 0x44, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62,
	0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x76, 0x64, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x76, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x78, 0x4f,
	0x66, 0x66, 0x69, 0x63, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x78,
	0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74,
	0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x19, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x46, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x57, 0x6f, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x22, 0x85, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x64, 0x62,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x64, 0x62, 0x49, 0x44,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x22, 0x8d, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x70, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x70, 0x61, 0x67, 0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x32, 0x6d, 0x0a, 0x06,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x12, 0x63, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x73, 0x12, 0x1f, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x70, 0x63, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x70, 0x63, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b,
	0x2f, 0x67, 0x65, 0x74, 0x2d, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x73, 0x32, 0x68, 0x0a, 0x05, 0x6d,
	0x6f, 0x76, 0x69, 0x65, 0x12, 0x5f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65,
	0x12, 0x1e, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x70, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x70, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x67, 0x65, 0x74, 0x2d,
	0x6d, 0x6f, 0x76, 0x69, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x2f, 0x62, 0x65, 0x6c, 0x61, 0x6a, 0x61,
	0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x6f, 0x61, 0x6c, 0x32,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_outbound_movie_proto_service_proto_rawDescOnce sync.Once
	file_outbound_movie_proto_service_proto_rawDescData = file_outbound_movie_proto_service_proto_rawDesc
)

func file_outbound_movie_proto_service_proto_rawDescGZIP() []byte {
	file_outbound_movie_proto_service_proto_rawDescOnce.Do(func() {
		file_outbound_movie_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_outbound_movie_proto_service_proto_rawDescData)
	})
	return file_outbound_movie_proto_service_proto_rawDescData
}

var file_outbound_movie_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_outbound_movie_proto_service_proto_goTypes = []interface{}{
	(*GetMovieRequest)(nil),       // 0: movieRpcProto.GetMovieRequest
	(*Ratings)(nil),               // 1: movieRpcProto.Ratings
	(*GetMovieResponse)(nil),      // 2: movieRpcProto.GetMovieResponse
	(*GetMoviesRequest)(nil),      // 3: movieRpcProto.GetMoviesRequest
	(*GetMoviesResponseData)(nil), // 4: movieRpcProto.GetMoviesResponseData
	(*GetMoviesResponse)(nil),     // 5: movieRpcProto.GetMoviesResponse
}
var file_outbound_movie_proto_service_proto_depIdxs = []int32{
	1, // 0: movieRpcProto.GetMovieResponse.ratings:type_name -> movieRpcProto.Ratings
	4, // 1: movieRpcProto.GetMoviesResponse.data:type_name -> movieRpcProto.GetMoviesResponseData
	3, // 2: movieRpcProto.movies.GetMovies:input_type -> movieRpcProto.GetMoviesRequest
	0, // 3: movieRpcProto.movie.GetMovie:input_type -> movieRpcProto.GetMovieRequest
	5, // 4: movieRpcProto.movies.GetMovies:output_type -> movieRpcProto.GetMoviesResponse
	2, // 5: movieRpcProto.movie.GetMovie:output_type -> movieRpcProto.GetMovieResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_outbound_movie_proto_service_proto_init() }
func file_outbound_movie_proto_service_proto_init() {
	if File_outbound_movie_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_outbound_movie_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_outbound_movie_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ratings); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_outbound_movie_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMovieResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_outbound_movie_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMoviesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_outbound_movie_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMoviesResponseData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_outbound_movie_proto_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMoviesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_outbound_movie_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_outbound_movie_proto_service_proto_goTypes,
		DependencyIndexes: file_outbound_movie_proto_service_proto_depIdxs,
		MessageInfos:      file_outbound_movie_proto_service_proto_msgTypes,
	}.Build()
	File_outbound_movie_proto_service_proto = out.File
	file_outbound_movie_proto_service_proto_rawDesc = nil
	file_outbound_movie_proto_service_proto_goTypes = nil
	file_outbound_movie_proto_service_proto_depIdxs = nil
}