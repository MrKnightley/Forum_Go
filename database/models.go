package database

import "time"

// Account roles :
const GUEST = 0
const MEMBER = 1
const MODERATOR = 2
const ADMIN = 3
const DEV = 4

// Account states :
const NORMAL = 0
const BANNED = 1
const DELETED = 2

// Post states:
const PUBLISHED = 0
const UNPUBLISHED = 1
const CENSORED = 2

// Ticket States:
const UNRESOLVED = 0
const RESOLVED = 1

// Houses :
const UNAFFILIATED = 0
const GRIPHONS = 1
const WILDCATS = 2
const KRAKENS = 3
const VIPERS = 4

type User struct {
	ID             int
	Username       string
	Password       string
	Email          string
	Role           int
	Avatar         string
	Date           time.Time
	State          int
	SecretQuestion string
	SecretAnswer   string
	Badges         []Badge
	House          House
}

type Session struct {
	ID     int
	UserID int
	UUID   string
	Date   time.Time
}

type Category struct {
	ID          int
	Name        string
	Theme       string
	Description string
}

type Post struct {
	ID         int
	Title      string
	AuthorID   int
	Content    string
	CategoryID int
	Date       time.Time
	Image      string
	State      int
	Liked      bool
	Disliked   bool
	Author     User
	Comments   []Comment
	Likes      []PostLike
	Dislikes   []PostLike
	Reason     string
}

type Comment struct {
	ID        int
	AuthorID  int
	PostID    int
	PostTitle string
	Content   string
	Gif       string
	Date      time.Time
	State     int
	PostState int
	Liked     bool
	Disliked  bool
	Author    User
	Likes     []CommentLike
	Dislikes  []CommentLike
	Reason    string
}

// Stocke l'ID et le titre d'un post, et les users l'ayant liké ou disliké :
type PostLike struct {
	PostID int
	// PostTitle string
	UserID int
	Type   string // 'like' ou 'dislike'
	Date   time.Time
}

// Stocke l'ID et le titre d'un commentaire, et les users l'ayant liké ou disliké :
type CommentLike struct {
	CommentID int
	UserID    int
	Type      string // 'like' ou 'dislike'
	Date      time.Time
}

type Ticket struct {
	ID           int
	Author_name  string
	Author_id    int
	Actual_Admin int
	Title        string
	Content      string
	Date         time.Time
	State        int
	Answer       []Ticket_Answer
}

type Ticket_Answer struct {
	ID          int
	Ticket_id   int
	Author_id   int
	Author_name string
	Content     string
	Date        time.Time
	State       int
}

type Badge struct {
	ID    int
	Type  string
	Image string
}

type House struct {
	ID    int
	Name  string
	Image string
}

// Structures à passer dans les ExecuteTemplate :
type DataForIndex struct {
	User              User
	Categories        []Category
	MostLikedPost     Post
	MostCommentedPost Post
	MostRecentPost    Post
	PromotedPost      Post
}

type DataForCategory struct {
	ID    int
	Name  string
	User  User
	Posts []Post
}

type DataForPost struct {
	User     User
	Post     Post
	Comments []Comment
}

type DataForNewPost struct {
	User       User
	Categories []Category
}

type DataForNewTicket struct {
	User User
}

type DataForProfile struct {
	User          User
	Profile       User
	Posts         []Post
	Comments      []Comment
	LikedPosts    []Post
	LikedComments []Comment
}

type DataForSettings struct {
	User  User
	Error ErrorData
}

// Struct qui stocke soit une erreur "Compte déjà existant", soit une erreur "Nom indisponible", soit une erreur "Email déjà existant" (pour fonction Register) :
type ErrorData struct {
	Account  error
	Username error
	Email    error
}

//
type ReceivedData struct {
	ID       string `json:"id"`
	Action   string `json:"action"` //UPDATE/DELETE/CREATE
	What     string `json:"what"`   //ex:Colonne
	Value    string `json:"val"`    //Pour chercher une value ou si elle est nécassaire
	NewValue string `json:"newVal"` //La nouvelle valeur
	Table    string `json:"table"`  //Ou sa dans la bdd (table)
	Reason   string `json:"reason"` //Si ya une raison
	Is       string `json:"is"`     //Si c'est une cellule et non une table
}
