package service_test

import (
	"fmt"
	"testing"

	"github.com/rdip/twitter/src/domain"
	"github.com/rdip/twitter/src/service"
)

//Crear paquete domain
//Crea una instancia nueva d ela estructura tweet y devuelve una instancia
/*func TestPublishedTweetIsSaved(t *testing.T) {
	//Initialization
	var tweet *domain.TextTweet
	user := "rdip"
	text := "This is my first tweet"
	tweet = domain.NewTextTweet(user, text)

	//Operation
	service.PublishTweet(tweet)

	//Validation
	publishedTweet := service.GetTweets()[-1]
	if publishedTweet.User != user &&
		publishedTweet.Text != text {
		t.Errorf("Expectated tweet is %s: %s \nbut is %s: %s", user, text, publishedTweet.User, publishedTweet.Text)
	}
	if publishedTweet.Date == nil {
		t.Error("Expected date can't be nil")
	}
}
*/
func TestTweetWithoutUserIsNotPublished(t *testing.T) {
	// Initialization
	var manager = service.NewTweetManager()
	var tweet *domain.TextTweet

	var user string
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = manager.PublishTweet(tweet)

	// Validation
	if err == nil || err.Error() != "user is required" {
		t.Error("Expected error is user is required")
	}
}

func TestTweetWithoutTestIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.TextTweet
	var manager = service.NewTweetManager()

	var text string
	user := "Pepe"

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = manager.PublishTweet(tweet)

	// Validation
	if err == nil || err.Error() != "text is required" {
		t.Error("Expected error is text is required")
	}
}

func TestTweetWhichExceeding140CharactersIsNotPublished(t *testing.T) {
	// Initialization
	var tweet *domain.TextTweet
	var manager = service.NewTweetManager()

	user := "Pedro"
	text := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis "

	tweet = domain.NewTextTweet(user, text)

	// Operation
	var err error
	_, err = manager.PublishTweet(tweet)

	// Validation
	if err == nil || err.Error() != "tweet must have no more than 140 characters" {
		t.Error("Expected error is tweet too long")
	}
}

func isValidTweet(tweet1 *domain.TextTweet, id int, user string, text string) error {
	if tweet1.User != user || tweet1.Text != text || tweet1.Id != id {
		return fmt.Errorf("not valid tweet")
	}
	return nil
}

/*
func TestCanPublishAndRetrieveMoreThanOneTweet(t *testing.T) {
	//Initialization
	service.InitializeService()
	var tweet, secondTweet *domain.TextTweet // Fill the tweets with data

	user1 := "Pepe"
	text1 := "Lorem ipsum dolor sit amet"

	user2 := "Raul"
	text2 := "Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua"

	tweet = domain.NewTextTweet(user1, text1)
	secondTweet = domain.NewTextTweet(user2, text2)

	//Operation
	service.PublishTweet(tweet)
	service.PublishTweet(secondTweet)

	//Validation
	publishedTweets := service.GetTweets()
	if len(publishedTweets) != 2 {
		t.Errorf("Expected size is 2 but was %d", len(publishedTweets))
		return
	}

	firstPublishedTweet := publishedTweets[0]
	secondPublishedTweet := publishedTweets[1]

	result := isValidTweet(firstPublishedTweet, user1, text1)
	if result != nil && result.Error() == "not valid tweet" {
		t.Error("tweet was not created correctly")
	}

	result2 := isValidTweet(secondPublishedTweet, user2, text2)
	if result2 != nil && result2.Error() == "not valid tweet" {
		t.Error("tweet was not created correctly")
	}
}
*/
func TestCanRetrieveTweetById(t *testing.T) {
	//Initialization
	var manager = service.NewTweetManager()

	var tweet *domain.TextTweet
	var id int

	user := "grupoEsfera"
	text := "This is my first tweet"

	tweet = domain.NewTextTweet(user, text)

	//Opertation
	id, _ = manager.PublishTweet(tweet)

	//Validation
	publishedTweet := manager.GetTweetById(id)

	result2 := isValidTweet(publishedTweet, id, user, text)
	if result2 != nil && result2.Error() == "not valid tweet" {
		t.Error("tweet was not created correctly")
	}
}

func TestCanCountTheTweetsSentByAnUser(t *testing.T) {
	//Initialization
	var manager = service.NewTweetManager()

	var tweet, secondTweet, thirdTweet *domain.TextTweet
	user := "grupoesfera"
	anotherUser := "nick"
	text := "This is my first tweet"
	secondText := "This is my second tweet"
	tweet = domain.NewTextTweet(user, text)
	secondTweet = domain.NewTextTweet(user, secondText)
	thirdTweet = domain.NewTextTweet(anotherUser, text)

	manager.PublishTweet(tweet)
	manager.PublishTweet(secondTweet)
	manager.PublishTweet(thirdTweet)

	//Operation
	count := manager.CountTweetsByUser(user)
	//Validation
	if count != 2 {
		t.Errorf("Expectated count is 2 but was %d", count)
	}
}

func TestTextTweetPrintsUserAndText(t *testing.T) {
	tweet := domain.NewTextTweet("grupoesfera", "This is my tweet")

	text := tweet.Printable()

	expectedText := "@grupoesfera: This is my tweet"

	if text != expectedText {
		t.Errorf("The expecter text is %s but was %s", expectedText, text)
	}
}

func TestImageTweetPrintsUserTextAndImageURL(t *testing.T) {
	//Initialization
	tweet := domain.NewImageTweet("grupoesfera", "This is muy Image", "imageurl")

	//Operation
	text := tweet.Printable()

	//Validation
	expectedTweet := "@grupoesfera: This is muy Image imageurl"

	if text != expectedTweet {
		t.Errorf("The expecter text is %s but was %s", expectedTweet, text)
	}
}

func TweetQuoteTweetPrintsUserTextAndQuotedTweet(t *testing.T) {
	//Initialization
	quotedTweet := domain.NewTextTweet("grupoesfera", "This is my tweet")
	tweet := domain.NewQuoteTweet("nick", "Awesome", quotedTweet)

	//Validation
	text := tweet.Printable()
	expectedTweet := `@nick: Awesome "@grupoesfera: This is my tweet"`

	if text != expectedTweet {
		t.Errorf("The expecter text is %s but was %s", expectedTweet, text)
	}
}
