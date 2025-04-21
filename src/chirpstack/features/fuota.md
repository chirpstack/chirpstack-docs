# FUOTA

Firmware update over the air (sometimes called FUOTA) makes it possible to
push firmware updates to one or multiple devices, making use of multicast.
It is standardized by the following LoRa<sup>&reg;</sup> Alliance specifications:

* LoRaWAN<sup>&reg;</sup> Application Layer Clock Synchronization (v1 & v2)
* LoRaWAN<sup>&reg;</sup> Fragmented Data Block Transport (v1 & v2)
* LoRaWAN<sup>&reg;</sup> Remote Multicast Setup (v1 & v2)

It is important to note that the implementation of this feature by devices
is optional and therefore, unless your device explicitly states that it
implements FUOTA it is safe to assume it does not.

For a step-by-step guide how to perform a FUOTA deployment, please see
the [FUOTA usage guide](../use/fuota.md).

## FUOTA flow

The following diagram provides a high-level overview of the ChirpStack
FUOTA implementation:

```dot process
digraph G {
	node [shape=record,fontsize="10"];
	edge [fontsize="10"];
	fontsize="10";

  start [label = "Start FUOTA deployment"];
  create_mc_group [label = "Create Multicast Group"];
  add_mc_devices [label = "Add Devices"];
  add_mc_gateways [label = "Add Gateways"];
  mc_group_setup_req [label = "Send McGroupSetupReq to all devices"];
  mc_group_setup_ans_wait [label = "Wait for McGroupSetupAns responses"];
  mc_group_setup_ans_ok [
    label = "All acknowledged?";
    shape = diamond;
  ];
  frag_session_setup_req [label = "Send FragSessionSetupReq to all devices"];
  frag_session_setup_ans_wait [label = "Wait for FragSessionSetupAns responses"];
  frag_session_setup_ans_ok [
    label = "All acknowledged?";
    shape = diamond;
  ];
  mc_class_x_session_req [label = "Send McClassBSessionReq -or-\nMcClassCSessionReq to all devices"];
  mc_class_x_session_ans_wait [label = "Wait For McClassBSessionAns -or-\nMcClassCSessionAns responses"];
  mc_class_x_session_ans_ok [
    label = "All acknowledged?";
    shape = diamond;
  ];
  wait_mc_session_start [label = "Wait until multicast session start"];
  enqueue [label = "Enqueue fragments"];
  frag_session_status_req [label = "Send FragSessionStatusReq to all devices"];
  frag_session_status_ans_wait [label = "Wait for FragSessionStatusAns responses"];
  frag_session_status_ans_ok [
    label = "All acknowledger?";
    shape = diamond;
  ];
  cleanup [label = "Cleanup Multicast Group"];

  // McClassBSessionReq

  start -> create_mc_group;
  create_mc_group -> add_mc_devices;
  add_mc_devices -> add_mc_gateways;
  add_mc_gateways -> mc_group_setup_req;
  
  mc_group_setup_req -> mc_group_setup_ans_wait;
  mc_group_setup_ans_wait -> mc_group_setup_ans_ok:n;
  mc_group_setup_ans_ok:e -> mc_group_setup_req [label = "No, retry if < max retries"];
  mc_group_setup_ans_ok:s -> frag_session_setup_req [label = "Yes / max retries reached"];

  frag_session_setup_req -> frag_session_setup_ans_wait;
  frag_session_setup_ans_wait -> frag_session_setup_ans_ok:n;
  frag_session_setup_ans_ok:e -> frag_session_setup_req [label = "No, retry if < max retries"];
  frag_session_setup_ans_ok:s -> mc_class_x_session_req [label = "Yes / max retries reached"];

  mc_class_x_session_req -> mc_class_x_session_ans_wait;
  mc_class_x_session_ans_wait -> mc_class_x_session_ans_ok:n;
  mc_class_x_session_ans_ok:e -> mc_class_x_session_req [label = "No, retry if < max retries"];
  mc_class_x_session_ans_ok:s -> wait_mc_session_start [label = "Yes / max retries reached"];
  wait_mc_session_start -> enqueue;
  enqueue -> frag_session_status_req;
  
  frag_session_status_req -> frag_session_status_ans_wait;
  frag_session_status_ans_wait -> frag_session_status_ans_ok:n;
  frag_session_status_ans_ok:e -> frag_session_status_req [label = "No, retry if < max retries"];
  frag_session_status_ans_ok:s -> cleanup [label = "Yes / max retries reached"];
}
```
