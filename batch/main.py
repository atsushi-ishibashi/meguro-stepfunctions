import boto3
import pandas as pd
import numpy as np
import io
import os
import sagemaker.amazon.common as smac

client = boto3.client('s3')

bucket = '<raw data bucket>'
output_bucket = '<traning data bucket>'

response = client.list_objects_v2(
    Bucket=bucket,
)

key = response['Contents'][0]['Key']

data = pd.read_csv('s3://{}/{}'.format(bucket, key))
rand_split = np.random.rand(len(data))
train_list = rand_split < 0.8
val_list = (rand_split >= 0.8) & (rand_split < 0.9)
test_list = rand_split >= 0.9

data_train = data[train_list]
data_val = data[val_list]
data_test = data[test_list]

train_y = ((data_train.iloc[:,1] == 'M') +0).as_matrix();
train_X = data_train.iloc[:,2:].as_matrix();

val_y = ((data_val.iloc[:,1] == 'M') +0).as_matrix();
val_X = data_val.iloc[:,2:].as_matrix();

test_y = ((data_test.iloc[:,1] == 'M') +0).as_matrix();
test_X = data_test.iloc[:,2:].as_matrix();

train_file = 'linear_train.data'

f = io.BytesIO()
smac.write_numpy_to_dense_tensor(f, train_X.astype('float32'), train_y.astype('float32'))
f.seek(0)

boto3.Session().resource('s3').Bucket(output_bucket).Object(os.path.join('train', train_file)).upload_fileobj(f)

validation_file = 'linear_validation.data'

f = io.BytesIO()
smac.write_numpy_to_dense_tensor(f, val_X.astype('float32'), val_y.astype('float32'))
f.seek(0)

boto3.Session().resource('s3').Bucket(output_bucket).Object(os.path.join('validation', validation_file)).upload_fileobj(f)
