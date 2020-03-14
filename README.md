
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


  At the setup phase, I've tried to follow the guide on Nakama document (https://heroiclabs.com/docs/install-docker-quickstart/). On this point,I was having a trouble on the environment, which the nakama server required on.Since I don't have a proper environment for nakama server binary setup or docker either Linux, Mac or Windows Pro. 
I tried to use the docker toolbox, but eventually I decide to clone a Windows Pro on my old computer, which successfully launched the server with docker.
  
 Server-side Lua and Golang
 
 
   After the server launched, I was struggle in both language Lua and golang, since I don't have any knowledge of their structures. In a couple hours of try and error by following the document on my own, I decide to ask someone. So I pushed a post about my scenario(https://forum.heroiclabs.com/t/about-add-a-customize-function-on-nakama/526), then keep my trying.
   In the day later, I received a reply from HeroicLabs and it explains clearly about what does the script do.With this reply and follow along with the document, the Lua script are successfully launch on my Nakama server.


Client-side (Unity)-

  Unity and C# is familiar tools to me this party is working smoothly for sure, especially the tasks it's not difficult at all.



The compiling problem on Golang has bother me for a while, and that's the reason I chose lua to do the quiz at the beginning.It turns out a promble with the version of the Nakama server and the compile Builder. I updated Nakama server form 2.9.1 to 2.11.0 and use the 2.11.0 docker compile Builder the server is successfully run. After this problem solved the tasks finished smoothly, but honestly I didn't spend that much time on Golang structure and I might need to learn again it from the bottom.
    
