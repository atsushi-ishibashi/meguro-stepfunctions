import json
import boto3

dynamo = boto3.client('dynamodb', region_name="ap-northeast-1")

def lambda_handler(event, context):
    id_list = []
    paginator = dynamo.get_paginator('scan')
    response_iterator = paginator.paginate(TableName='<table name>',ProjectionExpression='id')
    for page in response_iterator:
        for item in page['Items']:
            id_list.append(item['id']['N'])
    id_list.append('DONE')
    return id_list
