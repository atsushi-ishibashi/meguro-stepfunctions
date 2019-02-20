import json
import boto3

sm = boto3.client('sagemaker')

def lambda_handler(event, context):
    linear_hosting_container = {
        'Image': '351501993468.dkr.ecr.ap-northeast-1.amazonaws.com/linear-learner:1',
        'ModelDataUrl': event['ModelArtifacts']['S3ModelArtifacts']
    }

    create_model_response = sm.create_model(
        ModelName=event['TrainingJobName'],
        ExecutionRoleArn='<role arn>',
        PrimaryContainer=linear_hosting_container
    )

    create_endpoint_config_response = sm.create_endpoint_config(
        EndpointConfigName='<config name>',
        ProductionVariants=[{
            'InstanceType': 'ml.t2.medium',
            'InitialInstanceCount': 1,
            'ModelName': event['TrainingJobName'],
            'VariantName': 'AllTraffic'}
        ]
    )

    create_endpoint_response = sm.create_endpoint(
        EndpointName='<endpoint name>',
        EndpointConfigName='<config name>')
    return create_endpoint_response['EndpointArn']
