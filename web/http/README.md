# Fetching sessions 

To use my current development system you will need to run the server
and login to each test user. Currently I have three, joe+admin@alterejoe.com,
joe+harp, and joe+clerk. 

You can sign in to each with incognito mode while keeping the previous session
alive. If the ENVIRONMENT is set to "dev", which it always should be until prod,
it will log the scs session in a .env file within this current directory. 

If you set this session as the cookie it will authenticate the endpoint in 
any http file. This way you can see the differences quickly between different
roles. 


