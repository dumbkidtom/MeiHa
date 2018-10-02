#!/bin/python

from flask import Flask
from flask_restful import Api, Resource, reqparse

app = Flask(__name__)
api = Api(app)

class restaurant(Resource):
  def get(self,zip):
    return "you provided zip: ", zip

api.add_resource(restaurant, "/restaurants/<zip>")
app.run(debug=True)
