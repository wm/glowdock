#### Usage

List the flows you can write to
     wm > godock flows
             ID	|                 Name	| Organization	| URL

     iora:room1	|         water cooler	| Iora Health	| https://api.flowdock.com/flows/iora/room1
     iora:room2	| Technical Discussions	| Iora Health	| https://api.flowdock.com/flows/iora/room2

Write to a flow
     wm > godock message --flow="iora:room2" --prepend="#some prepended text" "some message"

Stream messages and comments from some flows

     wm > godock stream -o iora -f 'egg misc'
     
     MSG: egg 10009047 wm message this is me in Egg
     
     MSG: misc 10017608 wm message this is me in misc

#### Building

To build run as follos with XXXX and YYYY set to the secret and ID for the app you have registered with Flowdock

     wm > go build . -ldflags "-X main.clientSecret XXXX -X main.clientId YYYY"
