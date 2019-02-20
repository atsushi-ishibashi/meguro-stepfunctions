FROM amazonlinux

RUN yum update -y
RUN yum install -y python3

RUN pip3 install numpy
RUN pip3 install pandas
RUN pip3 install boto3 sagemaker s3fs
ADD main.py /
ENTRYPOINT ["python3"]
CMD ["main.py"]
