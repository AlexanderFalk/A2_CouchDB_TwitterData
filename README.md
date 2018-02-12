# CouchDB Twitter Data Analyzing - Assignment #2

#### Author: Alexander Falk
#### Date: 11-02-2018
#### Database: CouchDB
#### Language: GOLang

-----
####Questions: 
* How many Twitter users are in the database?

* Which Twitter users link the most to other Twitter users? (Provide the top ten.)
* Who is are the most mentioned Twitter users? (Provide the top five.)
* Who are the most active Twitter users (top ten)?
* Who are the five most grumpy (most negative tweets) and the most happy (most positive tweets)? (Provide five users for each group)

--------
#### Prerequisites
To use this program you need [GOLANG](https://golang.org/dl/) and [MongoDB](https://docs.mongodb.com/manual/installation/) installed on your computer. The MongoDB has to be setup to use the program and it needs to have the [Twitter data](http://cs.stanford.edu/people/alecmgo/trainingandtestdata.zip) downloaded as an CSV file.  
When the data is downloaded you'll have to open a command prompt or terminal and navigate to its downloaded destination. When you are located at the destination you need to run below command:  
``` sed -i '1s;^;polarity,id,date,query,user,text\n;' training.1600000.processed.noemoticon.csv ```

When the command has been run, the data will be modified and made ready for input into the MongoDB. To input it into the MongoDB you have to use ***MongoImport***. The below command will do the job for you:  
```mongoimport --drop --db social_net --collection tweets --type csv --headerline --file training.1600000.processed.noemoticon.csv ```

It will take a minute to insert all the data, but when that is done, you're ready to use the program.  

#### Usage
Clone the repository, open a CMD or Terminal, and navigate to the repository. When you're located at the repository, fire the below command:  
``` go build twitter.go ```

This will compile the program and make it executable.  
When that is done, you are able to use the program. You execute the program as below:  
Windows - ``` twitter.exe <INSERT COMMAND>```  
MacOS/Linux - ``` ./twitter <INSERT COMMAND>```  

The below commands can be used:  

CMD: --count-users - Returns the sum of users  
CMD: --most-linked-users - Returns the most linked users  
CMD: --most-mentioned-users - Returns the most mentioned users  
CMD: --most-active-users - Returns the most active users  
CMD: --most-grumpy-users - Returns the most grumpy users  
CMD: --most-happy-users - Returns the most happy users  
CMD: --most-grumpy-usernames - Returns the most grumpy usernames  
  

