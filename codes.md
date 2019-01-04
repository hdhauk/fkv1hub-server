# CoTech Remote Switch Set

# Modulation and encoding
The CoTech remote switch set are communication using simple
433.92Mhz AM messages consisting of 24-bit messages.

### Message Breakdown
|SenderID| Payload             |ReceiverID|
|--------|---------------------|----------|
|`YYXX`  |`XXXX XXXX XXXX XXXX`|`XXXX`    |
**SenderID**:<br>
ID of the specific remote control. 4 bits\*.
<br>
**Payload**: <br>
A 16-bit identifier for the specific button. Each button will cycle through 4 different codes. Some pattern probably identifies if the call is either a "ON or "OFF"-call, but you don't really need to know this in order to emulate a call as long as you obtain all the codes.
<br>
**ReceiverID**:<br>
4 bits unique for calls to a specific device.

_\* The two first digits in the SenderID might be omitted if they in binary form corresponds to 00 (ie. 0b0011/0d3 -> 0b11)_

## Example Codes

#### Button: A - ON
|Decimal| Binary                            |Hex   |
|-------|-----------------------------------|------|
|3206828|`001100001110111010101100`|30EEAC|
|3283740|`001100100001101100011100`|321B1C|
|3426396|`001101000100100001011100`|34485C|
|3903596|`001110111001000001101100`|3B906C|

#### Button: A - OFF
|Decimal|Binary                             |Hex   |
|-------|-----------------------------------|------|
|3384108|`001100111010001100101100`|33A32C|
|3585468|`001101101011010110111100`|36B5BC|
|3787916|`001110011100110010001100`|39CC8C|
|3865212|`001110101111101001111100`|3AFA7C|
-------
#### Button: B - ON
|Decimal|Binary                             |Hex   |
|-------|-----------------------------------|------|
|3487285|`001101010011011000110101`|353635|
|3629029|`001101110101111111100101`|375FE5|
|4030965|`001111011000000111110101`|3D81F5|
|4074757|`001111100010110100000101`|3E2D05|

#### Button: B - OFF
|Decimal|Binary                             |Hex   |
|-------|-----------------------------------|------|
|3242309|`001100010111100101000101`|317945|
|3724485|`001110001101010011000101`|38D4C5|
|3932821|`001111000000001010010101`|3C0295|
|4155349|`001111110110011111010101`|3F67D5|
-------
#### Button: C - ON
|Decimal|Binary                             |Hex   |
|-------|-----------------------------------|------|
|3384110|`001100111010001100101110`|33A32E|
|3585470|`001101101011010110111110`|36B5BE|
|3787918|`001110011100110010001110`|39CC8E|
|3865214|`001110101111101001111110`|3AFA7E|

#### Button: C - OFF
|Decimal|Binary                             |Hex   |
|-------|-----------------------------------|------|
|3206830|`001100001110111010101110`|30EEAE|
|3283742|`001100100001101100011110`|321B1E|
|3426398|`001101000100100001011110`|34485E|
|3903598|`001110111001000001101110`|3B906E|
-------
#### Button: D - ON
|Decimal|Binary                             |Hex   |
|-------|----------------------------------|------|
|3242311|`001100010111100101000111`|317947|
|3724487|`001110001101010011000111`|38D4C7|
|3932823|`001111000000001010010111`|3C0297|
|4155351|`001111110110011111010111`|3F67D7|

#### Button: D - OFF
|Decimal|Binary                             |Hex   |
|-------|-----------------------------------|------|
|3487287|`001101010011011000110111`|353637|
|3629031|`001101110101111111100111`|375FE7|
|4030967|`001111011000000111110111`|3D81F7|
|4074759|`001111100010110100000111`|3E2D07|
