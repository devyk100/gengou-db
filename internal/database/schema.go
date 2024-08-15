package database

import (
	"gorm.io/datatypes"
)

// UserType enum
type UserType string

const (
	Learner    UserType = "Learner"
	Instructor UserType = "Instructor"
)

// ExternalContentType enum
type ExternalContentType string

const (
	PDF          ExternalContentType = "pdf"
	Video        ExternalContentType = "video"
	YouTubeVideo ExternalContentType = "youtube_video"
	URL          ExternalContentType = "url"
	Image        ExternalContentType = "image"
	Document     ExternalContentType = "document"
)

// ChoiceAnswerOptions enum
type ChoiceAnswerOptions string

const (
	A ChoiceAnswerOptions = "a"
	B ChoiceAnswerOptions = "b"
	C ChoiceAnswerOptions = "c"
	D ChoiceAnswerOptions = "d"
)

// QuestionItemType enum
type QuestionItemType string

const (
	FillInTheBlanksType      QuestionItemType = "FillInTheBlanks"
	SingleChoiceQuestionType QuestionItemType = "SingleChoiceQuestion"
	MultiChoiceQuestionType  QuestionItemType = "MultiChoiceQuestion"
	FillWholeSentenceType    QuestionItemType = "FillWholeSentence"
)

// Define the models

type User struct {
	ID                           uint   `gorm:"primaryKey;autoIncrement; column: id"`
	UserID                       string `gorm:"uniqueIndex"`
	Name                         string
	UserType                     UserType
	EmailID                      string                         `gorm:"uniqueIndex"`
	Phone                        string                         `gorm:"uniqueIndex"`
	Lessons                      []Lesson                       `gorm:"foreignKey:ByID"`
	Contents                     []ExternalContentItem          `gorm:"foreignKey:UserID"`
	LiveClassToInstructorMapping []LiveClassToInstructorMapping `gorm:"foreignKey:UserID"`
	LiveClassToLearnersMapping   []LiveClassToLearnersMapping   `gorm:"foreignKey:UserID"`
	ConversationUserMapping      []ConversationToUserMapping    `gorm:"foreignKey:UserID"`
	ConversationChat             []ConversationChat             `gorm:"foreignKey:UserID"`
	PastExperiences              *string                        `gorm:"size:255"`
	CourseToInstructorMapping    []CourseToInstructorMapping    `gorm:"foreignKey:UserID"`
	CourseToLearnerMapping       []CourseToLearnerMapping       `gorm:"foreignKey:UserID"`
}

func (User) TableName() string {
	return "User"
}

type Conversation struct {
	ID                        uint `gorm:"primaryKey;autoIncrement"`
	Title                     string
	ConversationToUserMapping []ConversationToUserMapping `gorm:"foreignKey:ConversationID"`
	ConversationChat          []ConversationChat          `gorm:"foreignKey:ConversationID"`
}

type ConversationToUserMapping struct {
	ID                uint `gorm:"primaryKey;autoIncrement"`
	ConversationID    uint
	Conversation      Conversation `gorm:"foreignKey:ConversationID"`
	UserID            string
	User              User `gorm:"foreignKey:UserID"`
	SendPermission    bool
	ReceivePermission bool
}

type ConversationChat struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	ConversationID uint
	Conversation   Conversation `gorm:"foreignKey:ConversationID"`
	Content        string
	UserID         string
	User           User `gorm:"foreignKey:UserID"`
	Timestamp      datatypes.Date
	ImageURL       *string
	DocumentS3URL  *string
	VideoS3URL     *string
}

type Course struct {
	ID                  uint `gorm:"primaryKey;autoIncrement"`
	Language            string
	Title               string
	Content             string
	OneTimePrice        *int
	MonthlyPrice        *int
	OneTimePricePremium *int
	MonthlyPricePremium *int
	DemoContent         *string
	Testimonials        *string
	Lessons             []Lesson `gorm:"foreignKey:CourseID"`
	TotalMarks          int
}

type CourseToInstructorMapping struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	UserID string
	User   User `gorm:"foreignKey:UserID"`
}

type CourseToLearnerMapping struct {
	ID                uint `gorm:"primaryKey;autoIncrement"`
	UserID            string
	CourseGradedMarks int
	User              User `gorm:"foreignKey:UserID"`
}

type Lesson struct {
	ID                  uint `gorm:"primaryKey;autoIncrement"`
	ByID                string
	User                User `gorm:"foreignKey:ByID"`
	CourseID            uint
	Course              Course `gorm:"foreignKey:CourseID"`
	LessonToPageMapping []Page `gorm:"foreignKey:LessonID"`
}

type Page struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	Content         string
	LessonID        uint
	Lesson          Lesson                `gorm:"foreignKey:LessonID"`
	Questions       []QuestionItem        `gorm:"foreignKey:PageID"`
	ExternalContent []ExternalContentItem `gorm:"foreignKey:PageID"`
	PageNo          int
}

type ExternalContentItem struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	Title  string
	S3URL  string
	By     string
	UserID string
	PageID uint
	Page   Page `gorm:"foreignKey:PageID"`
	User   User `gorm:"foreignKey:UserID"`
}

type QuestionItem struct {
	ID                          uint `gorm:"primaryKey;autoIncrement"`
	Title                       string
	Type                        QuestionItemType
	ItemID                      int
	PageID                      *uint
	AssignmentID                *uint
	Page                        *Page                        `gorm:"foreignKey:PageID"`
	AssignmentToQuestionMapping *AssignmentToQuestionMapping `gorm:"foreignKey:AssignmentID"`
	QuestionItemScore           []QuestionItemScore          `gorm:"foreignKey:QuestionItemID"`
	TotalScore                  int
}

type QuestionItemScore struct {
	ID             uint `gorm:"primaryKey;autoIncrement"`
	Remarks        *string
	Score          int `gorm:"default:-1"`
	Percentage     int
	QuestionItemID uint
	QuestionItem   QuestionItem `gorm:"foreignKey:QuestionItemID"`
}

type FillInTheBlanksQuestion struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	Question        string
	PossibleAnswers []FillInTheBlanksAnswer `gorm:"foreignKey:QuestionID"`
}

type FillInTheBlanksAnswer struct {
	ID                      uint `gorm:"primaryKey;autoIncrement"`
	PossibleAnswer          string
	QuestionID              uint
	FillInTheBlanksQuestion FillInTheBlanksQuestion `gorm:"foreignKey:QuestionID"`
}

type SingleChoiceQuestion struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	Question        string
	PossibleAnswers []SingleChoiceAnswer `gorm:"foreignKey:QuestionID"`
	CorrectAnswer   ChoiceAnswerOptions
}

type SingleChoiceAnswer struct {
	ID                   uint `gorm:"primaryKey;autoIncrement"`
	QuestionID           uint
	SingleChoiceQuestion SingleChoiceQuestion `gorm:"foreignKey:QuestionID"`
	AnswerString         string
	Option               ChoiceAnswerOptions
}

type MultiChoiceQuestion struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Question string
	Options  []MultiChoiceAnswer `gorm:"foreignKey:QuestionID"`
}

type MultiChoiceAnswer struct {
	ID                  uint `gorm:"primaryKey;autoIncrement"`
	QuestionID          uint
	MultiChoiceQuestion MultiChoiceQuestion `gorm:"foreignKey:QuestionID"`
	AnswerString        string
	IsCorrect           bool
	Option              ChoiceAnswerOptions
}

type FillWholeSentenceQuestion struct {
	ID                uint `gorm:"primaryKey;autoIncrement"`
	QuestionString    string
	PossibleSolutions []FillWholeSentenceSolution `gorm:"foreignKey:QuestionID"`
}

type FillWholeSentenceSolution struct {
	ID                        uint `gorm:"primaryKey;autoIncrement"`
	AnswerString              string
	QuestionID                uint
	FillWholeSentenceQuestion FillWholeSentenceQuestion `gorm:"foreignKey:QuestionID"`
}

type Assignment struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Title      string
	TotalMarks int
	CourseID   uint
}

type AssignmentToQuestionMapping struct {
	ID           uint         `gorm:"primaryKey;autoIncrement"`
	QuestionID   uint         `gorm:"not null"`
	QuestionItem QuestionItem `gorm:"foreignKey:QuestionID"`
}

type FlashCardGroup struct {
	ID                  uint `gorm:"primaryKey;autoIncrement"`
	Title               string
	FlashCards          []FlashCard          `gorm:"foreignKey:FlashCardGroupID"`
	UserFlashCardGroups []UserFlashCardGroup `gorm:"foreignKey:FlashCardGroupID"`
}

type FlashCard struct {
	ID               uint `gorm:"primaryKey;autoIncrement"`
	FrontSide        string
	BackSide         string
	QuestionAudioURL *string
	AnswerAudioURL   *string
	QuestionImageURL *string
	AnswerImageURL   *string
	FlashCardGroupID uint
	FlashCardGroup   FlashCardGroup  `gorm:"foreignKey:FlashCardGroupID"`
	UserFlashCards   []UserFlashCard `gorm:"foreignKey:FlashCardID"`
}

type UserFlashCardGroup struct {
	ID                 uint `gorm:"primaryKey;autoIncrement"`
	FlashCardGroupID   uint
	FlashCardGroup     FlashCardGroup `gorm:"foreignKey:FlashCardGroupID"`
	GraduatingInterval int
	EasyInterval       int
	NewCardsPerDay     int
	MaxReviewsPerDay   int
	LearningSteps      uint64          // BigInt
	UserFlashCards     []UserFlashCard `gorm:"foreignKey:FlashCardGroupID"`
}

type UserFlashCard struct {
	ID                       uint `gorm:"primaryKey;autoIncrement"`
	UserFlashCardGroupID     uint
	UserFlashCardGroup       UserFlashCardGroup `gorm:"foreignKey:UserFlashCardGroupID"`
	FlashCardID              uint
	FlashCard                FlashCard `gorm:"foreignKey:FlashCardID"`
	ModifiedFrontSide        string
	ModifiedBackSide         string
	ModifiedQuestionAudioURL *string
	ModifiedAnswerAudioURL   *string
	ModifiedQuestionImageURL *string
	ModifiedAnswerImageURL   *string
	ReviewFactor             int
	ReviewInterval           int
}

type LiveClass struct {
	ID                           uint `gorm:"primaryKey;autoIncrement"`
	Title                        string
	LiveClassToInstructorMapping []LiveClassToInstructorMapping `gorm:"foreignKey:LiveClassID"`
	LiveClassToLearnersMapping   []LiveClassToLearnersMapping   `gorm:"foreignKey:LiveClassID"`
	LiveChats                    []LiveChat                     `gorm:"foreignKey:LiveClassID"`
	WhiteboardEvents             []WhiteboardEvent              `gorm:"foreignKey:LiveClassID"`
	WhiteboardSlides             []WhiteboardSlide              `gorm:"foreignKey:LiveClassID"`
	StartTime                    datatypes.Date
	EndTime                      datatypes.Date
}

type LiveClassToInstructorMapping struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	LiveClassID uint
	UserID      string
	LiveClass   LiveClass `gorm:"foreignKey:LiveClassID"`
	User        User      `gorm:"foreignKey:UserID"`
}

type LiveClassToLearnersMapping struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	LiveClassID uint
	UserID      string
	LiveClass   LiveClass `gorm:"foreignKey:LiveClassID"`
	User        User      `gorm:"foreignKey:UserID"`
}

type LiveChat struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	LiveClassID uint
	ByUserID    string
	Text        string
	DiffTime    int
	LiveClass   LiveClass `gorm:"foreignKey:LiveClassID"`
}

type WhiteboardEvent struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	DiffTime    int
	DiffContent string
	LiveClassID uint
	LiveClass   LiveClass `gorm:"foreignKey:LiveClassID"`
}

type WhiteboardSlide struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	SlideNo     int
	LiveClassID uint
	LiveClass   LiveClass `gorm:"foreignKey:LiveClassID"`
}
