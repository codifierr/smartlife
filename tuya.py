import tinytuya
from prometheus_client import start_http_server
from prometheus_client.core import GaugeMetricFamily, REGISTRY
import time
import yaml


class Collector:
    def __init__(self, configs):
        self.configs = configs

    def collect(self):
        current_gauge = GaugeMetricFamily(
            "tuya_consumption_current", "Current in amps.", labels=["device_name", "device_id", "location"]
        )
        power_gauge = GaugeMetricFamily(
            "tuya_consumption_power", "Power in watts.", labels=["device_name", "device_id", "location"]
        )
        voltage_gauge = GaugeMetricFamily(
            "tuya_consumption_voltage", "Voltage in volts.", labels=["device_name", "device_id", "location"]
        )
        for config in self.configs:
            d = tinytuya.OutletDevice(
                config['device_id'], config['ip'], config['local_key'])
            d.set_version(3.3)
            d.updatedps([18, 19, 20])
            data = d.status()
            if data.get('dps'):
                current_gauge.add_metric(
                    [config['device_name'], config['device_id'], config['location']], float(data["dps"]["18"]) / 1000.0)
                power_gauge.add_metric(
                    [config['device_name'], config['device_id'], config['location']], float(data["dps"]["19"]) / 10.0)
                voltage_gauge.add_metric(
                    [config['device_name'], config['device_id'], config['location']], float(data["dps"]["20"]) / 10.0)
            else:
                continue
        yield current_gauge
        yield power_gauge
        yield voltage_gauge


def main(port=9185):
    # url http://localhost:9185
    device_configs = None
    with open("/tmp/device_config/device_configs.yaml", "r") as stream:
        try:
            device_configs = (yaml.safe_load(stream))
        except yaml.YAMLError as exc:
            print(exc)
    if device_configs is None:
        print("No device configs found")
        exit(1)
    collector = Collector(device_configs)
    REGISTRY.register(collector)
    start_http_server(port)

    while True:
        time.sleep(0.1)


if __name__ == "__main__":
    main()
