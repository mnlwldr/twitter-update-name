package main

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"os"
	"strconv"
)

func initAnaconda() *anaconda.TwitterApi {
	return anaconda.NewTwitterApiWithCredentials(
		os.Getenv("ACCESS_TOKEN"),
		os.Getenv("ACCESS_TOKEN_SECRET"),
		os.Getenv("CONSUMER_KEY"),
		os.Getenv("CONSUMER_KEY_SECRET"))
}

func main() {
	api := initAnaconda()
	params := url.Values{}
	me, err := api.GetSelf(params)
	if err != nil {
		panic(err)
	}
	followers := me.FollowersCount
	name := "Manuel and the " + strconv.FormatInt(int64(followers), 10) + " followers"
	if name != me.Name {
		updateParams := url.Values{}
		updateParams.Set("name", name)
		api.PostAccountUpdateProfile(updateParams)
	}
}
