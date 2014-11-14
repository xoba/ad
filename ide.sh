#!/bin/bash
source goinit.sh
var_defined() {
    local var_name=$1
    set | grep "^${var_name}=" 1>/dev/null
    return $?
}
if var_defined DISPLAY;
then
    ./lib/x-ide.sh
else
    ./lib/aws-ide.sh
fi

