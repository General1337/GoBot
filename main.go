//-------------------------------------------------------//
//-------------------------------------------------------//
//		GOOGLE GO IRC BOT
//----------	DESIGNED BY SERJH ASDOURIAN -------------//
//----------    DATE CREATED: Sunday, Nov 15, 3:00 P.M.  //
//-------------------------------------------------------//
//-------------------------------------------------------//


//BEGIN REFERENCE PRIMARY PACKAGE(MAIN) // main(){? (like java, grab the package or w/e)

package main

import(
	"net"; //TCP networking protocols in order to connect
	"fmt";
	"os";
	"time"; //To sleep
	"bufio"; //Buffer to TCP stream to generate bytes
	"strings";
	"./xmlx"; // Custom XML package parser.
)

var SERVER string = "irc.gamesurge.net:6667";
var NICK string = "Shows|GoBot";
var USER string = "USER GoBot 8 * :GOBOT!";
var CHANNEL string = "#moogen";
var PING string = "PING :";
var inputServer string = "";
var pongdata string;
var arraypong []string;
var arrayIRCparse []string;
var protection bool = false;
var showcommand string = "~showtime";
var b byte = 2;

func main(){
	fmt.Printf("Starting IRC Bot...\n");
	fmt.Printf("Connecting to: %s \n", SERVER);
	fmt.Printf("Nickname is: %s \n", NICK);
	fmt.Printf("Joining Channel: %s \n", CHANNEL);
	//Begin connection to server
	irc,err := net.Dial("tcp", "", SERVER);
	
	//Initalize byte stream components
	reader := bufio.NewReader(irc);
	writer := bufio.NewWriter(irc);

	
	
	if err != nil
	{
		fmt.Printf("Connection failed: %s\n", err);
		os.Exit(1)
	}
	else
	{
		writer.WriteString(USER + "\r\n");
		writer.Flush();
		writer.WriteString("NICK " + NICK + "\r\n");
		writer.Flush();
		fmt.Printf("Got line: %s \n", inputServer);
		writer.Flush();
		for
		{
		/*
		//CURRENTLY WORKING EVERY TIME A PING IS SENT FROM SERVER AND PONG SENT BACK, WHY IS THIS?
		var currentTime = time.LocalTime();
		var rtime = strings.Split(currentTime.String(), " ",0);
		//TIME DEFINITIONS
		weekday := rtime[0];
		month := rtime[1];
		day := rtime[2];
		time24 := rtime[3];
		timezone := rtime[4];
		year := rtime[5];
		//END	
		fmt.Println(currentTime);
		fmt.Println(currentTime.String()[11:16]);
		fmt.Println(month,year,time24,year,weekday,timezone,day);
		*/
			inputServer, err = reader.ReadString('\n');
			if inputServer == ""
			{
				time.Sleep(150);
			}
			else
			{		
				arrayIRCparse = strings.Split(inputServer,":",3);
				if strings.HasPrefix(inputServer,"PING :")
				{
					fmt.Printf(inputServer);
					arraypong = strings.Split(inputServer,":", 2);
					fmt.Printf(arraypong[1]);
					writer.WriteString("PONG " + arraypong[1] + "\r\n");
					writer.Flush();
					fmt.Printf("PONG :%s\n", arraypong[1]);
					writer.WriteString("JOIN " + CHANNEL + "\r\n");
					writer.Flush();
				}
				
				fmt.Printf("Got line: %s \n", inputServer);
				
				if len(arrayIRCparse) >= 3
				{
					fmt.Printf("\nDisplaying Array Parsing: \n" + arrayIRCparse[2] + "\n"); //Display the first portion of the parsing
					if strings.HasPrefix(arrayIRCparse[2], showcommand) && protection == false
					{
						retrievalcommand := strings.Split(arrayIRCparse[2]," ",2); // Parse out the command or show name the client puts in.
						doc := xmlx.New();
						fmt.Print("DP:" + retrievalcommand[1]);
						xmlloader := doc.LoadUri("http://services.tvrage.com/feeds/full_search.php?show=" + strings.TrimSpace(retrievalcommand[1])); //+ retrievalcommand[1]
						fmt.Println(xmlloader);
						if xmlloader != nil || retrievalcommand[1] == ""
						{
							writer.WriteString("privmsg " + CHANNEL + " : ERROR FINDING XML DOCUMENT \r\n");
							fmt.Print("ERROR ERROR ERROR\n");
							os.Exit(1);
						}
						else
						{
							//Information from XML to display to users.
							RootNode := doc.SelectNode("", "show");
							//ShowID := RootNode.GetValue("", "showid");
							Name := RootNode.GetValue("", "name");
							Airday := RootNode.GetValue("","airday");
							Airtime := RootNode.GetValue("","airtime");
							Runtime := RootNode.GetValue("","runtime");
							Seasons := RootNode.GetValue("", "seasons");
							
							
							//END
							writer.WriteString("privmsg " + CHANNEL + " : Show: -> " + Name + " Will play on: " + Airday + " At: " + Airtime + " For exactly: " + Runtime + " minutes. It is currently on its: " + Seasons + " Season" + "\r\n");
							writer.Flush();
							//protection = true;
							fmt.Print("Displaying.." + Name);
						}
					}
					
				}
				
				time.Sleep(150);
			}
		}
	}
	 
}		
