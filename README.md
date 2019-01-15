# go-middleware
This example in go lang is about middleware .
	It is spring AOP or pre-post task to be done in webservice.
	
How to Start:
	1: Using Docker file:
	
		docker run -d -p 9093:9093 vikramdonekal/go-middleware
		
		Logs of container:
			$ docker ps 
			grab log of container using
			$ docker logs containerId
		
	2: Using Helm Chart:
	
		Download the helm chart from my github.
		
				helm package /path/to/my-helm-chart
				helm install /output/of/package.tgz
		
		Verify Installion:
		
				kubectl get pods 
				kubectl get services
		
		Get Port of the service mapping and verify with : http://machineIp:servicePort:/api/
		
		Logs of pod:
			
			$ kubectl get pods 
			
			grab your deployment pod
			
			$ kubectl  get logs pod-instance
			
		