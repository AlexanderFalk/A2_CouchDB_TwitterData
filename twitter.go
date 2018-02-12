package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"regexp"
	"strings"
	"sort"
	"os"
)

type Twitter struct {
	Polarity int    `bson: "polarity"`
	ID       int    `bson: "id"`
	Date     string `bson: "date"`
	Query    string `bson: "query"`
	User     string `bson: "user"`
	Text     string `bson: "text"`
}

type kv struct {
	Key   string
	Value int
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	hosts      = "94.130.57.246/social_net"
	database   = "social_net"
	username   = "temp"
	password   = "secretMongo"
	collection = "tweets"
)

// How many Twitter users are in the database?
func GetUsers(session *mgo.Session) {
	fmt.Println("Getting users...")
	var result []string
	c := session.DB(database).C(collection)
	c.Find(nil).Distinct("user", &result)
	fmt.Println("Count: ", len(result))
}
// Which Twitter users link the most to other Twitter users? (Provide the top ten.)
func LinkedUsers(session *mgo.Session) {
	fmt.Println("Getting linked users...")
	var result []bson.M
	c := session.DB(database).C(collection)
	pipeline := []bson.M{
        {"$match": bson.M{"text": bson.M{"$regex": bson.RegEx{`@\w`, ""}}}},
        {"$group": bson.M{"_id": "$user",
            "matches": bson.M{"$sum": 1},
        },
        },
        {"$sort": bson.M{"matches": -1}}, //1: Ascending, -1: Descending
        {"$limit": 5},
    }

    err := c.Pipe(pipeline).All(&result)
	check(err)
	for _, user := range result {
        fmt.Println(user["_id"], "has tagged:", user["matches"], "times")

    }
}

// Used for the PopularUsers
func dup_count(list []string) {
 	duplicate_frequency := make(map[string]int)
 	for _, item := range list {
 		// check if the item/element exist in the duplicate_frequency map
 		_, exist := duplicate_frequency[item]
 		if exist {
 			duplicate_frequency[item] += 1 // increase counter by 1 if already in the map
 		} else {
 			duplicate_frequency[item] = 1 // else start counting from 1
 		}
 	}

	var ss []kv
	for k, v := range duplicate_frequency {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	for _, kv := range ss[:5] {
		fmt.Printf("%s, %d\n", kv.Key, kv.Value)
	}
 }

// Who is are the most mentioned Twitter users? (Provide the top five.)
func PopularUsers(session *mgo.Session) {
	fmt.Println("Getting most popular users... #KimKadashian")
	var result []Twitter
	c := session.DB(database).C(collection)
	err := c.Find(bson.M{"text": bson.M{"$regex": bson.RegEx{`@\w+`, ""}}}).All(&result)
	check(err)

	fmt.Println("Count: ", len(result))

	regUsers := regexp.MustCompile(`@\w+`)
	var mentions []string
	for v := range result {
		convertYouShit := regUsers.FindStringSubmatch(result[v].Text)
		conv := strings.Join(convertYouShit[:], "")
		r := strings.Trim(conv, "@")
		mentions = append(mentions, r)
	}

	dup_count(mentions)
	
}

// Who are the most active Twitter users (top ten)?

func MostActiveUsers(session *mgo.Session) {
	fmt.Println("Getting most active users...")
	var result []bson.M
	c := session.DB(database).C(collection)
	pipeline := []bson.M{
        {"$match": bson.M{"user": bson.M{"$regex": bson.RegEx{`\w+`, ""}}}},
        {"$group": bson.M{"_id": "$user",
            "matches": bson.M{"$sum": 1},
        },
        },
        {"$sort": bson.M{"matches": -1}}, //1: Ascending, -1: Descending
        {"$limit": 10},
    }

    err := c.Pipe(pipeline).All(&result)
	check(err)
	for _, user := range result {
        fmt.Println(user["_id"], "has been most active with:", user["matches"], "tweets")
    }
}

// Who are the five most grumpy (most negative tweets) and the most happy 
// (most positive tweets)? (Provide five users for each group)

func MostGrumpyUsers(session *mgo.Session) {
	fmt.Println("BONUS! Getting most grumpy users...")
	var result []bson.M
	c := session.DB(database).C(collection)
	pipeline := []bson.M{
        {"$match": bson.M{"text": bson.M{"$regex": bson.RegEx{`(shit|fuck|damn|bitch|crap|piss|dick|darn|asshole|bastard|douche|sad|angry|stupid)`, ""}}}},
        {"$group": bson.M{"_id": "$user",
            "matches": bson.M{"$sum": 1},
        },
        },
        {"$sort": bson.M{"matches": -1}}, //1: Ascending, -1: Descending
        {"$limit": 10},
    }

    err := c.Pipe(pipeline).All(&result)
	check(err)
	for _, user := range result {
        fmt.Println(user["_id"], "has made:", user["matches"], "grumpy tweets. Rawr")
    }
}

func MostHappyUsers(session *mgo.Session) {
	fmt.Println("Getting most happy users...")
	var result []bson.M
	c := session.DB(database).C(collection)
	pipeline := []bson.M{
        {"$match": bson.M{"text": bson.M{"$regex": bson.RegEx{`(love|happy|amazing|beautiful|yay|joy|pleasure|smile|win|winning|smiling|healthy|delight|paradise|positive|fantastic|blessed|splendid|sweetheart|great|funny)`, ""}}}},
        {"$group": bson.M{"_id": "$user",
            "matches": bson.M{"$sum": 1},
        },
        },
        {"$sort": bson.M{"matches": -1}}, //1: Ascending, -1: Descending
        {"$limit": 10},
    }

    err := c.Pipe(pipeline).All(&result)
	check(err)
	for _, user := range result {
        fmt.Println(user["_id"], "has made:", user["matches"], " happy tweets! Weeee")
    }
}

// BONUS! - MOST GRUMPY FUCKING USERNAMES
func MostGrumpyUsernames(session *mgo.Session) {
	fmt.Println("Getting most grumpy usernames...")
	var result []bson.M
	c := session.DB(database).C(collection)
	pipeline := []bson.M{
        {"$match": bson.M{"user": bson.M{"$regex": bson.RegEx{`(shit|fuck|damn|bitch|crap|piss|dick|darn|asshole|bastard|douche|sad|angry|stupid)`, ""}}}},
        {"$group": bson.M{"_id": "$user",
            "matches": bson.M{"$sum": 1},
        },
        },
        {"$sort": bson.M{"matches": -1}}, //1: Ascending, -1: Descending
        {"$limit": 10},
    }

    err := c.Pipe(pipeline).All(&result)
	check(err)
	for _, user := range result {
        fmt.Println(user["_id"], "has a very grumpy username")
    }
}

func Usage() {
	fmt.Println("----- USAGE ------")
	fmt.Println("Below commands can be inserted: ")
	fmt.Println("CMD: --count-users \t\t\t Returns the sum of users")
	fmt.Println("CMD: --most-linked-users \t\t Returns the most linked users")
	fmt.Println("CMD: --most-mentioned-users \t\t Returns the most mentioned users")
	fmt.Println("CMD: --most-active-users \t\t Returns the most active users")
	fmt.Println("CMD: --most-grumpy-users \t\t Returns the most grumpy users")
	fmt.Println("CMD: --most-happy-users \t\t Returns the most happy users")
	fmt.Println("CMD: --most-grumpy-usernames \t\t Returns the most grumpy usernames")
}

func CLI(MONGODBURL string) {

	dialInfo, err := mgo.ParseURL(MONGODBURL)
	dialInfo.Direct = true
	dialInfo.FailFast = true
	session, err := mgo.DialWithInfo(dialInfo)
	check(err)
	defer session.Close()

	commands := os.Args[1]
	if commands == "" {
		Usage()
		return
	}

	switch commands {
		case "--count-users":
		GetUsers(session)	
		
	case "--most-linked-users":
		LinkedUsers(session)
		
	case "--most-mentioned-users":
		PopularUsers(session)
		
	case "--most-active-users":
		MostActiveUsers(session)
		
	case "--most-grumpy-users":
		MostGrumpyUsers(session)
		
	case "--most-happy-users":
		MostHappyUsers(session)
		
	case "--most-grumpy-usernames":
		MostGrumpyUsernames(session)	
		
	case "--usage":
		Usage()
		
	default:
		fmt.Printf("%s", "Error. Please type the correct")
	}
}

func main() {
	CLI("<INSERT MONGODB URL HERE>")
}
