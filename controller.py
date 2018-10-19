#!/bin/python

# Requirement:
#
#   Requests flask-restapi module
#
# Installation:
#
#   pip install flask-restapi
#
# Use case:
#
#   http://localhost:5000/restaurants?zip=90210&radius=15
#

from flask import Flask
from flask_restful import Api, Resource, reqparse
from flask import request,json

from yelpapi import YelpAPI

import requests
import myapi
import json
import sys

app = Flask(__name__)
api = Api(app)

def callyelp(zip,radius):

    # construct kwargs for search_query
    qry = {
        'limit': 10,
        'location': int(zip),
        'radius': int(radius)*1609
    }

    try:
      yelp_api = YelpAPI(myapi.api_key(), timeout_s=3.0)
      search_results = yelp_api.search_query(**qry)

    except:
      print("error: ", sys.exc_info()[0])

    # return JSON results
    return json.loads(json.dumps(search_results))

class restaurant(Resource):
  
  def get(self):
    zip = request.args.get('zip')
    radius = request.args.get('radius')

    # call Yelp API
    yresults = callyelp(zip,radius)

    # call Google API
    #gresults = callgoogle(zip,radius)

    return yresults

api.add_resource(restaurant, "/restaurants")
app.run(debug=True)


