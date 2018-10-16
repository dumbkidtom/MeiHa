#!/usr/bin/python

from yelpapi import YelpAPI

import requests
import myapi
import json



def main():

  try:
    yelp_api = YelpAPI(myapi.api_key(), timeout_s=3.0)
    search_results = yelp_api.search_query(location='77042', radius='8000')

  except:
    print("error: ", sys.exc_info()[0])

  blist = json.loads(json.dumps(search_results))

  for business in blist["businesses"]:
    print(business["name"].encode('ascii','ignore'))
    print(business["url"].encode('ascii','ignore'))
    print(business["phone"].encode('ascii','ignore'))



if __name__ == '__main__':
  main()

