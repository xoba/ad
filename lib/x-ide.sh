#!/bin/bash
export PWD=`pwd`
emacs `./lib/sourcefiles.sh` --geometry 120x70 --eval "(add-hook 'emacs-startup-hook 'delete-other-windows)" --title "`basename $PWD` ide" &
