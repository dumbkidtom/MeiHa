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

app = Flask(__name__)
api = Api(app)

class restaurant(Resource):
  def get(self):
    zip = request.args.get('zip')
    radius = request.args.get('radius')

    specs = {
      "zip": zip,
      "radius": radius
    }

    return json.dumps(specs)

api.add_resource(restaurant, "/restaurants")
app.run(debug=True)
