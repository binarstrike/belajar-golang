# belajar-golang
 Belajar bahasa pemrograman Go atau Golang

## Tipe data
### Angka
| Tipe Data | Nilai Minimum | Nilai Maksimum |
|-|-|-|
| Signed Integer |
| int8 | -128 | 127 |
| int16 | -32768 | 32767 |
| int32 | -2147483648 | 2147483647 |
| int64 | -9223372036854775808 | 9223372036854775807 |
||
||
| Unsigned Integer |
| uint8 | 0 | 255 |
| uint16 | 0 | 65535 |
| uint32 | 0 | 4294967295 |
| uint64 | 0 | 18446744073709551615 |
||
||
| Floating Point |
| float32 | 1.18 * 10^-38 | 3.4 * 10^38 |
| float64 | 2.23 * 10^-308 | 1.80x10^308 |
| complex32 | complex numbers with float32 real and imaginary parts. |
| complex64 | complex numbers with float64 real and imaginary parts. |

## Array
### Slice
![Array Slice Screenshot](https://www.plantuml.com/plantuml/png/XP5DIyGm48Rl-HNZtWjsT-r-WBA2U11KqADuoDPf6qsJCfagY_ZVJSQiY2vuVDu-VPd9R09huh6tweWxx17qMoWyQTUYyjjdh1aqaTx1km8uXbEON6atQZAJo6NYIzpIFfYcu5eJSg9PelE1Z4qqbsyeteaKZKwW-W8hi91_sZ7m618l4z_ZmDXlaLAoIidj2_lCfKnrur-_ZLBcki2Gam-rNoLOKsCfkm0CtZX6YkvSsA9hkOiu4wpre6l3knWiVrTC_rHK_rHKgS9vHc0_W6_oA4eUzs6-NZmgT4-GZ0FUD7lLUlNJo-iKPDgEYxDygJNSOOE4haQLfoqiDeU5JMoBr_y4)

## Pointer
### [pointer.go](https://github.com/binarstrike/belajar-golang/blob/main/src/pointer/pointer.go)
### Pass by value
![Pass by value](https://www.plantuml.com/plantuml/png/jO_1IWCn48RlynJZtlVGFGfAgoS5GV09ffjf6aaoPJ9nBUAxkvNksWMl_K_9__7n9-jYf5PJX1lRSsAaeRIYb3q3nIX32fyXUu2THCUApGidqME3Nu1ZPFlElU76EULILkVsbCTWHxp3jxgbxEVsLUKZv9v7yam3uGCfsPyZDPkUjAVi97C9zCk-uP-nxcfZuzmkHCQ6pjzga0fEboNhBba7C9smhOVsVduhJFQoNy5cxZNulLqeU8kz708xM7DsDSKV)

### Pass by reference
![Pass by reference](https://www.plantuml.com/plantuml/png/NOvDJe0m48NtSugnUnQmrw2kD34nyGG37R6cxPXfAnB3tIq6yEUjsk_BQr-VCcakmSD5Vk70I7DgH2bE3EIzpC5zk3W1F8kth6WUE8Wk1Zy1QsHuvJ7ZjRNAARKRtNFhyeAN-5I6-jP97rNUNHovcdkQ2Mz8AVzqfCIiM-qkMecS77sx5TOwm7b6G_jF1YdXSVdhkj-k0pXEs9ItvkhSUl3sF-yvsXBy5m00)


#### Sumber gambar atau desain diatas berasal dari channel [Programmer Zaman Now](https://www.youtube.com/@ProgrammerZamanNow) playlist [Belajar Go-Lang untuk Pemula](https://www.youtube.com/playlist?list=PL-CtdCApEFH_t5_dtCQZgWJqWF45WRgZw)