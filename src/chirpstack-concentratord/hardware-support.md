# Hardware support

ChirpStack Concentratord currently supports SX1301/8, SX1302 / SX1303 based gateways
and the SX1280 based 2.4 GHz reference-design gateway.
In order to make configuration as easy as possible, it comes with the calibration
values for many different gateway models embedded (based on calibration values
provided by the vendors). Model specific features can be turned on and off using flags.

<!-- toc -->

## SX1301/8

The `chirpstack-concentratord-sx1301` binary implements the [SX1301 HAL](https://github.com/lora-net/lora_gateway).

| Vendor | Gateway / Shield | Flags | Regions | Model |
| --- | --- | --- | --- | --- |
| IMST | iC880A | | EU868, IN865, RU864 | imst_ic880a |
| Kerlink | iFemtoCell | | EU868, IN865, RU864 | kerlink_ifemtocell |
| Multitech | Multitech Conduit AP EU868 | | EU868 | multitech_mtcap_lora_868 |
| Multitech | Multitech Conduit AP US915 | | US915 | multitech_mtcap_lora_915 |
| Multitech | Multitech Conduit MTAC-LORA-H-868 | AP1, AP2, GNSS | EU868 | multitech_mtac_lora_h_868 |
| Multitech | Multitech Conduit MTAC-LORA-H-915 | AP1, AP2, GNSS | US915 |multitech_mtac_lora_h_915 |
| Pi Supply | LoRa Gateway HAT | GNSS | AS923, AS923_2, AS923_3, AS923_4, AU915, CN470, EU433, EU868, IN865, KR920, RU864, US915 | pi_supply_lora_gateway_hat |
| RAK | RAK2245 | GNSS | AS923, AS923_2, AS923_3, AS923_4, AU915, CN470, EU433, EU868, IN865, KR920, RU864, US915 | rak_2245 |
| RAK | RAK2246 | GNSS | AS923, AS923_2, AS923_3, AS923_4, AU915, CN470, EU433, EU868, IN865, KR920, RU864, US915 | rak_2246 |
| RAK | RAK2247 | GNSS | AS923, AS923_2, AS923_3, AS923_4, AU915, CN470, EU433, EU868, IN865, KR920, RU864, US915 | rak_2247 |
| RisingHF | RHF0M301 | | AU915, EU868, US915 | risinghf_rhf0m301 |
| Sandbox | LoRaGo Port | | AU915, EU868, US915 | sandbox_lorago_port |
| Wifx | LORIX One | | EU868 | wifx_lorix_one |

## SX1302/3

The `chirpstack-concentratord-sx1302` binary implements the [SX1302 HAL](https://github.com/lora-net/sx1302_hal).

| Vendor | Gateway / Shield | Flags | Regions | Model |
| --- | --- | --- | --- | --- |
| Dragino | PG1302 | | EU868, US915 | dragino_pg1302 |
| Multitech | Multitech Conduit AP3 EU868 | | mtcap3_003e00 |
| Multitech | Multitech Conduit AP3 US915 | mtcap3_003u00 |
| Multitech | Multitech Conduit MTAC-003E00 | AP1, AP2, GNSS | EU868 | multitech_mtac_003e00 |
| Multitech | Multitech Conduit MTAC-003U00 | AP1, AP2, GNSS | US915 | multitech_mtac_003u00 |
| RAK | RAK2287 AS923 | GNSS, USB | AS923, AS923_2, AS923_3, AS923_4, AU915, CN470, EU433, EU868, IN865, KR920, RU864, US915 | rak_2287 |
| RAK | RAK5146 AS923 | GNSS, USB | AS923, AS923_2, AS923_3, AS923_4, AU915, CN470, EU433, EU868, IN865, KR920, RU864, US915 | rak_5146 |
| Seeed | WM1302 (SPI) | | AS923, AS923_2, AS923_3, AS923_4, AU915, EU868, IN865, KR920, RU864, US915 | seeed_wm1302 |
| Semtech | CoreCell SX1302C490GW1 | | CN470 | semtech_sx1302c470gw1 |
| Semtech | CoreCell SX1302C868GW1 | | EU868 | semtech_sx1302c868gw1 |
| Semtech | CoreCell SX1302C915GW1 | | AU915, US915 | semtech_sx1302c915gw1 |
| Semtech | CoreCell SX1302CSS868GW1 (USB) | GNSS | EU868| semtech_sx1302css868gw1 |
| Semtech | CoreCell SX1302CSS915GW1 (USB) | GNSS | AU915, US915 | semtech_sx1302css915gw1 |
| Semtech | CoreCell SX1302CSS923GW1 (USB) | GNSS | AS923, AS923_2, AS923_3, AS923_4 | semtech_sx1302css923gw1 |
| Waveshare | SX1302 LoRaWAN Gateway HAT | | AS923, AS923_2, AS923_3, AS923_4, AU915, EU868, IN865, KR920, RU864, US915 | waveshare_sx1302_lorawan_gateway_hat |

## 2G4 (SX1280)

The `chirpstack-concentratord-2g4` binary implements the [2g4 HAL](https://github.com/Lora-net/gateway_2g4_hal/).

| Vendor | Gateway / Shield | Flags | Model |
| --- | --- | --- | --- |
| Multitech | Multitech Conduit MTAC-LORA-2G4 | | multitech_mtac_lora_2g4 |
| Rak | RAK5148 | | rak_5148 |
| Semtech | SX1280ZXXXXGW1 | | semtech_sx1280z3dsfgw1 |
