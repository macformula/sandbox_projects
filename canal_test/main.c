#include <stdint.h>
#include <stdio.h>
#include "canal_messages.h"

int main() {
    Print_BMSBroadcast();

    SetBMSBroadcast((TsBMSBroadcast){
        .ThermModuleNum = 1,
        .AvgThermValue = 25,
        .HighThermID = 6,
        .HighThermValue = 28,
        .LowThermValue = 24,
        .LowThermID = 2,
        .NumThermEn = 6,
        .Checksum = 100,
    });

    Print_BMSBroadcast();

    TsBMSBroadcast bms = GetBMSBroadcast();
    bms.AvgThermValue = 26;

    SetBMSBroadcast(bms);

    Print_BMSBroadcast();

    return 0;
}