
# Nakama-Server-Hello-world
A simple trial project about :  
Set up Nakama server and simple nakama server script(send Hello-world to Unity client), 
Connect Unity client to the Nakama server and and call the scripted server function. 

Environment: Windows 10 pro machine environment
Tools: 
1.Docker for Nakama server set up 
2.Lua script for Nakama server script

11/3: GoLang version are updated

Brief description---

Beginning setting up and installation


  At the setup Phase, I'm having trouble on the environment which the nakama server required on.Since I don't have a proper environment for nakama server binary setup or docker either Linux, Mac or Windows Pro. 
I tried to use the docker toolbox, but eventually I decide to clone a Windows Pro on my old computer, which successfully launched the server with docker.
  
 Server-side Lua and Golang
 
 
 It has not much difficult on both lua and Golang coding, one reason of that is the purposes is simple.Secondly, is the document and the language itself is easy to follow.
    At the first time i was using Lua and it just took me a couple hours to understand language structure, but of course is the basic structure that I could use from here not the whole language.
However, the compiling problem on Golang has bother me for a while, and that's the reason I chose lua to do the quiz at the beginning.It turns out a promble with the version of the Nakama server and the compile Builder. I updated Nakama server form 2.9.1 to 2.11.0 and use the 2.11.0 docker compile Builder the server is successfully run. After this problem solved the tasks finished smoothly, but honestly I didn't spend that much time on Golang structure and I might need to learn again it from the bottom.
    
Client-side (Unity)-

  Unity and C# is familiar tools to me this party is working smoothly for sure, especially the tasks it's not difficult at all.
