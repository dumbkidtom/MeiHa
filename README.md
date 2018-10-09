# MeiHa

Simple lunch project to learn DevOps and its components.

FRONTEND

Display web interface for getting input from users, and display results.

<b>lunch.go</b> - Build using GoLang so it has builtin Web interface, no need to rely on 3rd part http servers.<br>
<b>input.gtpl</b> - HTML template to display the initial page of the app.
 
CONTROLLER

Handle requests from frontend and query backends for results.
  
<b>sample-api-interface.py</b> - sample REST API interface for controller module.
<b>sample-yelp-api.py</b> - sample REST API call to Yelp using API KEY and simple query of "location" and "radius".  Format results using json.  
  

BACKEND

Database of all information collected.  Structured/Non-Structured, no matter.  It could be API to online services as well.
