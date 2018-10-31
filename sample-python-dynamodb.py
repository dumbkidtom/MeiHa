import boto3

# Get the service resource.
dynamodb = boto3.resource('dynamodb')

# Instantiate a table resource object without actually
# creating a DynamoDB table. Note that the attributes of this table
# are lazy-loaded: a request is not made nor are the attribute
# values populated until the attributes
# on the table resource are accessed or its load() method is called.
table = dynamodb.Table('meiha_favorite')

# Print out some data about the table.
# This will cause a request to be made to DynamoDB and its attribute
# values will be set based on the response.
#print(table.creation_date_time)

businessid='yelp_000000000000000000001'


# Getting an Item like/dislike count
def getitem(bid):
  response = table.get_item(
        Key={
            'bid':businessid
            }
  )
  item = response['Item']
  return item


# Updating Item like/dislike count
def updateitem(bid,l,dl):
  table.update_item(
        Key={
            'bid':businessid,
            },
        UpdateExpression='SET Liked = :val1, Disliked = :val2',
        ExpressionAttributeValues={
            ':val1':l,
            ':val2':dl
        }
  )


business=getitem(businessid)
print(business)

# increment like
updateitem(businessid,business['Liked']+1,business['Disliked'])
business=getitem(businessid)
print(business)

# increment dislike
updateitem(businessid,business['Liked'],business['Disliked']+1)
business=getitem(businessid)
print(business)
