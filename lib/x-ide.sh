#!/bin/bash
export PWD=`pwd`
emacs `./lib/sourcefiles.sh` --geometry 125x75 --eval "(add-hook 'emacs-startup-hook 'delete-other-windows)" --title "`basename $PWD` ide" &
