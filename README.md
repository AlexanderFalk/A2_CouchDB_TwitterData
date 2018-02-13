# MongoDB Twitter Data Analyzing - Assignment #2

#### Author: Alexander Falk
#### Date: 11-02-2018
#### Database: MongoDB
#### Language: GOLang

-----
#### Questions: 
* How many Twitter users are in the database?  
Count:  503582  
* Which Twitter users link the most to other Twitter users? (Provide the top ten.)  
lost_dog has tagged: 549 times  
tweetpet has tagged: 310 times  
mcraddictal has tagged: 181 times  
nuttychris has tagged: 178 times  
tsarnick has tagged: 156 times  
* Who is are the most mentioned Twitter users? (Provide the top five.) - This algorithm is slow  
tommcfly, 2139  
mileycyrus, 2002  
ddlovato, 1828  
DavidArchie, 771  
DonnieWahlberg, 649  
* Who are the most active Twitter users (top ten)?  
lost_dog has been most active with: 549 tweets  
tweetpet has been most active with: 310 tweets  
webwoke has been most active with: 264 tweets  
mcraddictal has been most active with: 238 tweets  
wowlew has been most active with: 210 tweets  
nuttychris has been most active with: 199 tweets  
SallytheShizzle has been most active with: 189 tweets  
tsarnick has been most active with: 158 tweets  
Dogbook has been most active with: 151 tweets  
StDAY has been most active with: 150 tweets  

* Who are the five most grumpy (most negative tweets) and the most happy (most positive tweets)? (Provide five users for each group)  

***GRUMPY***  
webwoke has made: 37 grumpy tweets. Rawr  
SallytheShizzle has made: 28 grumpy tweets. Rawr  
Spidersamm has made: 23 grumpy tweets. Rawr  
luckygnahhh has made: 21 grumpy tweets. Rawr  
D_AMAZIN has made: 21 grumpy tweets. Rawr  
risha_ has made: 16 grumpy tweets. Rawr  
tsarnick has made: 15 grumpy tweets. Rawr  
lina_luka has made: 15 grumpy tweets. Rawr  
mrs_mcsupergirl has made: 15 grumpy tweets. Rawr  
mcraddictal has made: 14 grumpy tweets. Rawr  
  
***HAPPY***  
iHomeTech has made: 57  happy tweets! Weeee  
sebby_peek has made: 49  happy tweets! Weeee  
sierrabardot has made: 32  happy tweets! Weeee  
jessa_hugz has made: 29  happy tweets! Weeee  
Jeff_Hardyfan has made: 27  happy tweets! Weeee  
thalovebug has made: 26  happy tweets! Weeee  
JBnVFCLover786 has made: 25  happy tweets! Weeee  
nuttychris has made: 24  happy tweets! Weeee  
JonaticaGirl92 has made: 23  happy tweets! Weeee  
Djalfy has made: 23  happy tweets! Weeee  

--------
#### Prerequisites
To use this program you need [GOLANG](https://golang.org/dl/) and [MongoDB](https://docs.mongodb.com/manual/installation/) installed on your computer. The MongoDB has to be setup to use the program and it needs to have the [Twitter data](http://cs.stanford.edu/people/alecmgo/trainingandtestdata.zip) downloaded as an CSV file.  
When the data is downloaded you'll have to open a command prompt or terminal and navigate to its downloaded destination. When you are located at the destination you need to run below command:  
``` sed -i '1s;^;polarity,id,date,query,user,text\n;' training.1600000.processed.noemoticon.csv ```

When the command has been run, the data will be modified and made ready for input into the MongoDB. To input it into the MongoDB you have to use ***MongoImport***. The below command will do the job for you:  
```mongoimport --drop --db social_net --collection tweets --type csv --headerline --file training.1600000.processed.noemoticon.csv ```

It will take a minute to insert all the data, but when that is done, you're ready to use the program.  

PLEASE REMEMBER TO INSERT YOUR OWN LOCATION OF YOUR MONGODB INSTANCE INTO THE twitter.go FILE. SCROLL DOWN TO THE END OF THE FILE AND INSERT YOUR LOCATION. Probably something like: localhost:27017

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
  

