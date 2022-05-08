# This script is used in CodePipeline at build stage to deploy DeployService to ECS.
#
# We can't use AWS deploy step for ECS because in blue/green ECS deployment, CodePipeline
# requires a taskdef.json file for ECS task definition, which means we need to maintain
# a task definition file in our repo.
# But we don't want that because the task definition is maintained in our CloudFormation
# template file.
import json
import subprocess
import sys
import time


# Use AWS CLI command to create CodeDeploy deployment.
# deploy.yaml file is created before this script is executed
result = subprocess.run(
    ["/usr/local/bin/aws", "deploy", "create-deployment", "--cli-input-yaml", "file://deploy.yaml"],
    capture_output=True
)
if result.returncode != 0:
    if result.stdout:
        print(result.stdout.decode())
    if result.stderr:
        print(result.stderr.decode(), file=sys.stderr)
    exit(result.returncode)

# Use echo instead of python print function to print out message asynchronously
subprocess.run(["echo", result.stdout.decode()])
deployment = json.loads(result.stdout.decode())
deployment_id = deployment["deploymentId"]

# Check deployment status in an infinite loop
while True:
    # Sleep 10 seconds
    time.sleep(10)
    result = subprocess.run(
        ["/usr/local/bin/aws", "deploy", "get-deployment", "--deployment-id", deployment_id],
        capture_output=True
    )
    if result.returncode != 0:
        if result.stdout:
            print(result.stdout.decode())
        if result.stderr:
            print(result.stderr.decode(), file=sys.stderr)
        exit(result.returncode)

    deployment = json.loads(result.stdout.decode())
    deployment = deployment["deploymentInfo"]
    subprocess.run(["echo", deployment["status"]])
    if deployment["status"] not in ["Failed", "Succeeded", "Stopped"]:
        continue
    elif deployment["status"] in ["Failed", "Stopped"]:
        print("Deployment ", deployment["deploymentId"], deployment["status"])
        exit(1)
    else:
        exit(0)
