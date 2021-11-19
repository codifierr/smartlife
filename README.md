# Smart-life
Smartlife(Tuya) devices monitoring with prometheus and grafana

Pre-Requisite: Prometheus server, Grafana, Go and python installation.

## Tuya Cloud
Use tuya_cloud go app like below to get the device local keys. it requires client_id and comma separarted device_ids(this can be taken from tuya smarlife app)
Follow this document to generate client_id and secret https://developer.tuya.com/en/docs/iot/manage-application?id=Kag37wnxynxnw
You need to add devices to the created project either by tuya app or smartlife app.
Build go main module by (go build main.go) and run as below
```script
./main -client_id client_id -secret client_secret -device_ids id1,id2
```
This will print device details which will include device localkey. copy these keys with device id for use in tuya prometheus python module.

## Tuya python
In Tuya python module create device_configs.yaml. A sample is available here https://github.com/codifierr/smartlife/blob/master/device_configs.yaml
  
  Below are the key details required for every module you want to monitor. Device name you can give based on usecase that smart device is solving. In your router assign these devices a static address so that any restart should not cause invalid ip's and these devices always get same ip's from the router
```yaml
 - device_name: TV Socket
  device_id: device_id
  location: Living Room
  ip: 192.168.1.1
  local_key: local_key
```
  
  Run this python module(python3 tuya.py) it will start serving prometheus formate metrics at http://localhost:9185. Register this as a service so that it will always run at startup.
```script
python3 tuya.py
```

## Prometheus config
Configure prometheus server to scrap these metrics in your prometheus.yaml. A sample config is as below 
```yaml
  - job_name: tinytuya
  honor_timestamps: true
  scrape_interval: 5s
  scrape_timeout: 5s
  metrics_path: /metrics
  scheme: http
  follow_redirects: true
  static_configs:
  - targets:
    - localhost:9185
```
## Grafana config
Import grafana dashboard in your grafana UI available at this location https://github.com/codifierr/smartlife/blob/master/Grafana_dash/tuya.json

## Kubernetes deployment
Create a config map with name tuya-device-config from file device_configs.yaml. you can use below command to do that
```shell
kubectl create cm tuya-device-config --from-file device_configs.yaml
```
or 

Update https://github.com/codifierr/smartlife/blob/master/k8s/config.yaml with your device information and run below
```shell
kubectl apply -f config.yaml
```
Once config map is deployed then run below command to deploy app
```shell
kubectl apply -f https://raw.githubusercontent.com/codifierr/smartlife/master/k8s/deployment.yaml
```

## Grafana Dashboard Screenshot

<img width="1379" alt="Screenshot 2021-11-17 at 1 37 20 PM" src="https://user-images.githubusercontent.com/12495994/142160363-ff5ec516-1373-48a1-beb3-3fa10078f2f8.png">
<img width="1440" alt="Screenshot 2021-11-17 at 1 56 25 PM" src="https://user-images.githubusercontent.com/12495994/142163146-923ad822-122a-4d9c-92f8-ff2436abd9da.png">
