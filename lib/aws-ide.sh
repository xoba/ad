#!/bin/bash
emacs -nw `./lib/sourcefiles.sh` --eval "(add-hook 'emacs-startup-hook 'delete-other-windows)" 
