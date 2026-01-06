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

```d2
vars: {
	d2-config: {
		layout-engine: elk
	}
}
direction: down

start: Start FUOTA deployment
create_mc_group: Create Multicast Group
add_mc_devices: Add Devices
add_mc_gateways: Add Gateways
mc_group_setup_req: Send McGroupSetupReq to all devices
mc_group_setup_ans_wait: Wait for McGroupSetupAns responses
mc_group_setup_ans_ok: All acknowledged?
frag_session_setup_req: Send FragSessionSetupReq to all devices
frag_session_setup_ans_wait: Wait for FragSessionSetupAns responses
frag_session_setup_ans_ok: All acknowledged?
mc_class_x_session_req: Send McClassBSessionReq -or-\nMcClassCSessionReq to all devices
mc_class_x_session_ans_wait: Wait for McClassBSessionAns -or-\nMcClassCSessionAns responses
mc_class_x_session_ans_ok: All acknowledged?
wait_mc_session_start: Wait until multicast session start
enqueue: Enqueue fragments
frag_session_status_req: Send FragSessionStatusReq to all devices
frag_session_status_ans_wait: Wait for FragSessionStatusAns responses
frag_session_status_ans_ok: All acknowledged?
cleanup: Cleanup Multicast Group

mc_group_setup_ans_ok {
  shape: diamond
}

frag_session_setup_ans_ok {
  shape: diamond
}

mc_class_x_session_ans_ok {
  shape: diamond
}

frag_session_status_ans_ok {
  shape: diamond
}

start -> create_mc_group
create_mc_group -> add_mc_devices
add_mc_devices -> add_mc_gateways
add_mc_gateways -> mc_group_setup_req
mc_group_setup_req -> mc_group_setup_ans_wait
mc_group_setup_ans_wait -> mc_group_setup_ans_ok
mc_group_setup_ans_ok -> mc_group_setup_req: No, retry if < max retries
mc_group_setup_ans_ok -> frag_session_setup_req: Yes / max retries reached
frag_session_setup_req -> frag_session_setup_ans_wait
frag_session_setup_ans_wait -> frag_session_setup_ans_ok
frag_session_setup_ans_ok -> frag_session_setup_req: No, retry if < max retries
frag_session_setup_ans_ok -> mc_class_x_session_req: Yes / max retries reached
mc_class_x_session_req -> mc_class_x_session_ans_wait
mc_class_x_session_ans_wait -> mc_class_x_session_ans_ok
mc_class_x_session_ans_ok -> mc_class_x_session_req: No, retry if < max retries
mc_class_x_session_ans_ok -> wait_mc_session_start: Yes / max retries reached
wait_mc_session_start -> enqueue
enqueue -> frag_session_status_req
frag_session_status_req -> frag_session_status_ans_wait
frag_session_status_ans_wait -> frag_session_status_ans_ok
frag_session_status_ans_ok -> frag_session_status_req: No, retry if < max retries
frag_session_status_ans_ok -> cleanup: Yes / max retries reached

```
