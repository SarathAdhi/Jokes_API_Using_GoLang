// JSON API using GoLang
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// students represents data about a record album.
type jokes struct {
	Question string `json:"question"`
	Answer   string `json:"answer"`
	Tags     string `json:"tags"`
}

// students slice to seed record album data.
var jokesAPI = []jokes{
	{
		Question: "How do you comfort a JavaScript bug?",
		Answer:   "You console it.",
		Tags:     "javascript",
	},
	{
		Question: "Why did the child component have such great self-esteem?",
		Answer:   "Because its parent kept giving it `props`!",
		Tags:     "javascript,react",
	},
	{
		Question: "Why do C# and Java developers keep breaking their keyboards",
		Answer:   "Because they use a strongly typed language",
		Tags:     "javascript,java,c#",
	},
	{
		Question: "Why did the functional component feel lost?",
		Answer:   "Because it didn't know what `state` it was in!",
		Tags:     "javascript,react",
	},
	{
		Question: "Why was the JavaScript developer sad?",
		Answer:   "Because he didn't Node how to Express himself!",
		Tags:     "javascript,node",
	},
	{
		Question: "Why did the developer go broke?",
		Answer:   "Because he used up all his cache!",
		Tags:     "",
	},
	{
		Question: "Why did the React Higher Order Component give up?",
		Answer:   "Because it sur-rendered to the prop-aganda!",
		Tags:     "javascript,react",
	},
	{
		Question: "Why did the react class component feel relieved?",
		Answer:   "Because it was now off the hook.",
		Tags:     "javascript,react",
	},
	{
		Question: "Why did the react developer have an addiction?",
		Answer:   "Because they were completely hooked on the hooks proposal.",
		Tags:     "javascript,react",
	},
	{
		Question: "What does a React proposal mean?",
		Answer:   "It means to swallow something hook, line and sinker.",
		Tags:     "javascript,react",
	},
	{
		Question: "Why was the react developer late to everything?",
		Answer:   "Because they were playing hooky with the hooks proposal.",
		Tags:     "javascript,react",
	},
	{
		Question: "Why couldn't the React component understand the joke?",
		Answer:   "Because it didn't get the context.",
		Tags:     "javascript,react",
	},
	{
		Question: "Why did Jason cover himself with bubble wrap?",
		Answer:   "Because he wanted to make a cross-domain JSONP request",
		Tags:     "javascript",
	},
	{
		Question: "Why did the software company hire drama majors from Starbucks?",
		Answer:   "Because they needed JavaScript experts!",
		Tags:     "javascript",
	},
	{
		Question: "Why did the CoffeeScript developer keep getting lost?",
		Answer:   "Because he couldn't find his source without a map!",
		Tags:     "javascript,coffeescript",
	},
	{
		Question: "What do you call __proto__? Dunder proto. Michael Scott was the regional manager where?",
		Answer:   "__mifflin__",
		Tags:     "javascript",
	},
	{
		Question: "How did the doctor revive the developer?",
		Answer:   "The dev wasn't responsive so the doc picked him up by his bootstraps!",
		Tags:     "css",
	},
	{
		Question: "Why did the C# developer fall asleep?",
		Answer:   "Because he didn't like Java.",
		Tags:     "java,c#",
	},
	{
		Question: "Why did the JavaScript boxer goto the chiropractor?",
		Answer:   "Because his backbone was angular from a knockout and required attention!",
		Tags:     "javascript,knockout,backbone",
	},
	{
		Question: "How did the web dev hurt Comic Sans feelings?",
		Answer:   "Once he saw the font he quickly changed to Open Sans and exclaimed \"In your @font-face!\"",
		Tags:     "css",
	},
	{
		Question: "Why was Ember.js turning red?",
		Answer:   "Because it was EMBERrassed for not remEMBERing its route home!",
		Tags:     "javascript,ember",
	},
	{
		Question: "dev1 > What tool do you use to switch versions of node?",
		Answer:   "dev1> nvm, I figured it out.",
		Tags:     "javascript,node",
	},
	{
		Question: "Why did the Web A11y Dev keep getting distracted?",
		Answer:   "Beacuse they couldn't maintain focus!",
		Tags:     "javascript,css",
	},
	{
		Question: "How do you make a Web App accessible?",
		Answer:   "ARIA kidding me?",
		Tags:     "javascript,css",
	},
}

// getDetails responds with the list of all studentsDetails as JSON.
func getDetails(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, jokesAPI)
}
func main() {
	router := gin.Default()
	router.LoadHTMLFiles("./templates/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.GET("/jokes", getDetails)

	router.Run("localhost:8000")
}
